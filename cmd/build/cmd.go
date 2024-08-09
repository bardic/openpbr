package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/cmd/img"
	"github.com/bardic/openpbr/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

// Cmd represents the build command
var Cmd = &cobra.Command{
	Use:              "build",
	Short:            "build pack",
	Long:             ``,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(time.Now().String())

		if !utils.ZipOnly {

			if utils.DeleteAutoGen {
				fmt.Println("--- Cleaning workspace")
				clean.Cmd.RunE(cmd, nil)

				fmt.Println("--- Download latest base assets")

				if utils.SkipDownload {
					download.Cmd.RunE(cmd, []string{"skip"})
				} else {
					download.Cmd.RunE(cmd, nil)
				}
			}

			fmt.Println("--- Prcoess PSDs")
			gen.ConvertPsdCmd.RunE(cmd, []string{"./psds/blocks"})

			fmt.Println("--- Copy custom configs")
			cp.Copy(utils.SettingDIr, utils.OutDir)

			entries, _ := os.ReadDir(utils.BaseAssets)
			f := entries[0]

			for _, s := range utils.TargetAssets {
				fmt.Println("--- Create json, mer and height files for " + s)
				p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)
				build(cmd, s, p)
			}

			fmt.Println("--- Copy Overrides")
			cp.Copy(utils.Overrides, utils.OutDir+string(os.PathSeparator)+"textures")

			for _, s := range utils.TargetAssets {
				p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)
				createMers(cmd, p)
			}

			fmt.Println("--- Create manifest")
			gen.ManifestCmd.RunE(cmd, []string{})
		}

		if utils.Crush {
			fmt.Println("--- Crush images")
			img.CrushCmd.RunE(cmd, []string{utils.OutDir})
		}

		gen.PackageCmd.RunE(cmd, []string{utils.OutDir})

		dat, err := os.ReadFile("VERSION")
		if err != nil {
			return
		}

		fmt.Println("Release Version: " + string(dat))
	},
}

// build is a recursive function that processes the images and generates json, mer, and height files.
func build(cmd *cobra.Command, target string, imgPath string) error {
	subPaths := strings.Split(imgPath, string(os.PathSeparator))
	items, _ := os.ReadDir(imgPath)

	for _, item := range items {
		outPath := utils.OutDir + string(os.PathSeparator) + strings.Join(subPaths[3:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name()
		if item.IsDir() {
			if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
				return err
			}
			p := filepath.Join(imgPath, item.Name())
			build(cmd, target, p)
		} else {
			if !strings.Contains(item.Name(), ".tga") && !strings.Contains(item.Name(), ".png") {
				continue
			}

			in := filepath.Join(imgPath, item.Name())

			if strings.Contains(item.Name(), ".tga") {
				outPath = strings.ReplaceAll(outPath, ".tga", ".png")
				img.TgaPngCmd.RunE(cmd, []string{in, outPath})
			} else {
				err := utils.CopyF(in, outPath)
				if err != nil {
					return err
				}
			}

			if strings.Contains(outPath, ".png") {
				img.AdjustColorCmd.Run(cmd, []string{outPath})

				if utils.NormalMaps {
					err := gen.NormalCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", "_normal.png")})
					if err != nil {
						return err
					}
				} else {
					err := gen.HeightCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", "_height.png")})
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
			}
		}
	}

	return nil
}

func createMers(cmd *cobra.Command, inputPath string) error {
	subPaths := strings.Split(inputPath, string(os.PathSeparator))
	items, _ := os.ReadDir(inputPath)

	for _, item := range items {
		if item.IsDir() {
			createMers(cmd, inputPath+string(os.PathSeparator)+item.Name())
		} else {
			outPath := utils.OutDir + string(os.PathSeparator) + strings.Join(subPaths[3:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name()
			merPath := strings.ReplaceAll(outPath, ".png", "_mer.png")

			useMerFile := true
			if _, err := os.Stat(merPath); err != nil {
				useMerFile = false
			}

			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

			merArr := "[0, 0, 255]"
			if utils.TexturesetVersion == "1.21.30" {
				merArr = "[0, 0, 255, 255]"
			}

			merFile := fn + "_mer"

			heightNormalFile := fn + "_height"
			if utils.NormalMaps {
				heightNormalFile = fn + "_normal"
			}

			err := gen.JsonCmd.RunE(cmd, []string{
				strings.ReplaceAll(outPath, ".png", ".texture_set.json"),
				fn,
				merArr,
				merFile,
				heightNormalFile,
				strconv.FormatBool(useMerFile),
				utils.TexturesetVersion,
			})

			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}
