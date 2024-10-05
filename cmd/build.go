package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/data"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/cmd/img"
	"github.com/bardic/openpbr/cmd/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "build project based on json config",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonFile, err := os.Open(args[0])
		if err != nil {
			utils.AppendLoadOut("Fatal error: config.json missing")
			return err
		}

		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)

		if err != nil {
			utils.AppendLoadOut("Fatal error: failed to read config.json")
			return err
		}

		var jsonConfig data.Targets
		json.Unmarshal(byteValue, &jsonConfig)

		if len(jsonConfig.Targets) == 0 {
			utils.AppendLoadOut("Fatal error: no targets configured in config")
			return errors.New("no targets configured")
		}

		utils.TexturesetVersion = jsonConfig.Targets[0].Textureset_format

		utils.AppendLoadOut(time.Now().String())

		utils.AppendLoadOut("--- Cleaning workspace")
		err = clean.Cmd.RunE(cmd, nil)

		if err != nil {
			utils.AppendLoadOut("Fatal error: Failed to clean")
			return err
		}

		utils.AppendLoadOut("--- Download latest base assets")
		err = download.Cmd.RunE(cmd, []string{})

		if err != nil {
			utils.AppendLoadOut("Fatal error: Failed to download assets")
			return err
		}

		utils.AppendLoadOut("--- Prcoess PSDs")
		err = gen.ConvertPsdCmd.RunE(cmd, []string{utils.LocalPath(utils.Psds)})

		if err != nil {
			utils.AppendLoadOut("Warning: Failed to convert PSDs")
		}

		utils.AppendLoadOut("--- Copy custom configs")
		err = cp.Copy(utils.LocalPath(utils.SettingDir), utils.LocalPath(utils.OutDir))

		if err != nil {
			utils.AppendLoadOut("Warning: Failed to copy custom configs")
		}

		utils.AppendLoadOut("--- Create height files for ")
		entries, _ := os.ReadDir(utils.LocalPath(utils.BaseAssets))
		f := entries[0]
		for _, s := range utils.TargetAssets {
			utils.AppendLoadOut("--- --- " + s)
			p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)
			err = CreateHeightNormalFile(cmd, s, utils.LocalPath(p))

			if err != nil {
				utils.AppendLoadOut("Fatal error: Failed to build item in pack - " + s)
				return err
			}
		}

		utils.AppendLoadOut("--- Copy Overrides")
		err = cp.Copy(utils.LocalPath(utils.Overrides), utils.LocalPath(utils.OutDir+string(os.PathSeparator)+"textures"))

		if err != nil {
			utils.AppendLoadOut("Warning: Failed to copy overrides")
		}

		for _, s := range utils.TargetAssets {
			utils.AppendLoadOut("--- Create JSON files")
			p := utils.LocalPath(filepath.Join(utils.OutDir, "textures", s))
			utils.AppendLoadOut("--- --- " + p)
			err = CreateTextureSets(cmd, p)

			if err != nil {
				utils.AppendLoadOut("Fatal error: Failed to create JSON for item in pack - " + s)
				return err
			}
		}

		utils.AppendLoadOut("--- Create manifest")
		err = gen.ManifestCmd.RunE(cmd, []string{
			jsonConfig.Targets[0].Name,
			jsonConfig.Targets[0].Description,
			jsonConfig.Targets[0].Header_uuid,
			jsonConfig.Targets[0].Module_uuid,
			jsonConfig.Targets[0].Version,
			jsonConfig.Targets[0].Author,
			jsonConfig.Targets[0].License,
			jsonConfig.Targets[0].URL,
			jsonConfig.Targets[0].Capibility,
		})

		if err != nil {
			utils.AppendLoadOut("Fatal error: failed to create manifest")
			return err
		}

		utils.AppendLoadOut("--- Package Release")

		err = gen.PackageCmd.RunE(cmd, []string{utils.LocalPath(utils.OutDir)})

		if err != nil {
			utils.AppendLoadOut("Warning : packaging failed ")
			return err
		}

		utils.AppendLoadOut("--- OpenPBR complete")

		return nil
	},
}

func CreateHeightNormalFile(cmd *cobra.Command, target string, imgPath string) error {
	items, err := os.ReadDir(imgPath)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {
		if item.IsDir() {
			p := filepath.Join(imgPath, item.Name())
			CreateHeightNormalFile(cmd, target, p)
		} else {
			subpath, err := utils.GetTextureSubpath(imgPath, "textures")

			if err != nil {
				fmt.Println(err)
				return err
			}

			in := filepath.Join(imgPath, item.Name())
			outPath := utils.LocalPath(utils.OutDir + string(os.PathSeparator) + subpath + string(os.PathSeparator) + item.Name())
			os.MkdirAll(filepath.Dir(outPath), os.ModePerm)

			if strings.Contains(outPath, ".tga") {
				out := strings.ReplaceAll(in, ".tga", ".png")
				err := img.TgaPngCmd.RunE(cmd, []string{in, out})
				if err != nil {
					fmt.Println(err)
					return err
				}

				in = out
			}

			go func(in, out string) {
				err := cp.Copy(in, out)

				if err != nil {
					utils.AppendLoadOut("Warning: Failed to copy custom configs")
				}
			}(utils.LocalPath(utils.SettingDir), utils.LocalPath(utils.OutDir))

			err = gen.HeightCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", utils.HeightMapNameSuffix+".png")})
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}

func CreateTextureSets(cmd *cobra.Command, out string) error {
	items, _ := os.ReadDir(out)

	for _, item := range items {
		if item.IsDir() {
			CreateTextureSets(cmd, out+string(os.PathSeparator)+item.Name())
		} else {
			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

			if strings.Contains(fn, "_height") || strings.Contains(fn, "texture_set") || strings.Contains(fn, "_mer") {
				continue
			}

			q, e := utils.GetTextureSubpath(out, "textures")

			if e != nil {
				continue
			}

			outPath := utils.LocalPath(utils.OutDir + string(os.PathSeparator) + q + string(os.PathSeparator) + item.Name())
			merPath := strings.ReplaceAll(outPath, ".png", utils.MerMapNameSuffix+".png")

			useMerFile := true
			if _, err := os.Stat(merPath); err != nil {
				useMerFile = false
			}

			heightNormalFile := fn + utils.HeightMapNameSuffix
			if utils.NormalMaps {
				heightNormalFile = fn + utils.NormalMapNameSuffix
			}

			err := gen.JsonCmd.RunE(cmd, []string{
				strings.ReplaceAll(outPath, ".png", ".texture_set.json"),
				fn,
				"[0,0,255]",
				fn + utils.MerMapNameSuffix,
				heightNormalFile,
				strconv.FormatBool(useMerFile),
				utils.TexturesetVersion,
			})

			if err != nil {
				continue
			}
		}
	}

	return nil
}
