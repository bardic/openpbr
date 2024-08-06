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

		// if len(args) > 0 {
		// 	for _, a := range args {
		// 		if a == "clean" {
		// 			fmt.Println("--- Clean build dir")

		// 			utils.CleanDir = true
		// 		}

		// 		if a == "skip" {
		// 			fmt.Println("--- Skip Download")
		// 			utils.SkipDownload = true
		// 		}

		// 		if a == "beta" {
		// 			fmt.Println("--- Enable beta")
		// 			utils.Beta = true
		// 		}

		// 		if a == "normals" {
		// 			fmt.Println("--- Enable normals")
		// 			utils.NormalMaps = true
		// 		}

		// 		if a == "v=1.21.30" {
		// 			utils.TexturesetVersion = true
		// 		}
		// 	}
		// }

		if utils.CleanDir {
			fmt.Println("--- Cleaning workspace")
			clean.Cmd.RunE(cmd, nil)

			fmt.Println("--- Download latest base assets")

			if utils.SkipDownload {
				download.Cmd.RunE(cmd, []string{"skip"})
			} else {
				download.Cmd.RunE(cmd, nil)
			}
		}

		fmt.Println("--- Prcoess Glowables")
		gen.GlowablesCmd.RunE(cmd, []string{"./glowables/blocks"})

		entries, _ := os.ReadDir(utils.BaseAssets)
		f := entries[0]

		for _, s := range utils.TargetAssets {
			fmt.Println("--- Create json, mer and height files for " + s)
			p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)
			build(cmd, s, p)
		}
		fmt.Println("--- Copy custom configs")
		cp.Copy(utils.SettingDIr, utils.BuildDir)

		fmt.Println("--- Create manifest")
		gen.ManifestCmd.RunE(cmd, []string{})
	},
}

// build is a recursive function that processes the images and generates json, mer, and height files.
func build(cmd *cobra.Command, target string, imgPath string) error {
	items, _ := os.ReadDir(imgPath)

	for _, item := range items {
		if item.IsDir() {
			outputPath := strings.ReplaceAll(imgPath, utils.BaseAssets, utils.BuildDir)
			p := filepath.Join(outputPath, item.Name())
			if err := os.MkdirAll(p, 0770); err != nil {
				return err
			}
			p = filepath.Join(imgPath, item.Name())
			build(cmd, target, p)
		} else {
			if !strings.Contains(item.Name(), ".tga") && !strings.Contains(item.Name(), ".png") {
				continue
			}

			in := filepath.Join(imgPath, item.Name())
			out := filepath.Join(utils.BuildDir, "textures", target, item.Name())
			if strings.Contains(item.Name(), ".tga") {
				out = strings.ReplaceAll(out, ".tga", ".png")
				img.TgaPngCmd.RunE(cmd, []string{in, out})
			} else {
				err := utils.CopyF(in, out)
				if err != nil {
					return err
				}
			}

			if strings.Contains(out, ".png") {
				if utils.NormalMaps {
					fmt.Println("Create normalmaps")

					err := gen.NormalCmd.RunE(cmd, []string{out, strings.ReplaceAll(out, ".png", "_normal.png")})
					if err != nil {
						return err
					}
				} else {
					fmt.Println("Create heightmaps")
					err := gen.HeightCmd.RunE(cmd, []string{out, strings.ReplaceAll(out, ".png", "_height.png")})
					if err != nil {
						fmt.Println(err)
						return err
					}
				}

			}

			fmt.Println("Check overrides")

			b, err := utils.CheckForOverride(strings.ReplaceAll(out, ".png", "_mer.png"))

			if err != nil {
				fmt.Println(err)
				return err
			}

			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

			MerType := 1
			if b {
				MerType = 0
			}

			if utils.Beta {
				gen.UpscaleCmd.Run(cmd, []string{out, out})
				hmOut := strings.ReplaceAll(out, ".png", "_height.png")
				gen.UpscaleCmd.Run(cmd, []string{hmOut, hmOut})
				merOut := strings.ReplaceAll(out, ".png", "_mer.png")
				gen.UpscaleCmd.Run(cmd, []string{merOut, merOut})
			}

			fmt.Println("Create json")

			mer := "[0, 0, 255]"
			if utils.TexturesetVersion == "1.21.30" {
				mer = "[0, 0, 255, 255]"
			}

			err = gen.JsonCmd.RunE(cmd, []string{
				strings.ReplaceAll(out, ".png", ".texture_set.json"),
				fn,
				mer,
				fn + "_mer",
				fn + "_height",
				strconv.Itoa(MerType),
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
