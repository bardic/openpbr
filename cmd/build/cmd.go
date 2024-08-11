package build

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/common"
	"github.com/bardic/openpbr/cmd/data"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/cmd/img"
	"github.com/bardic/openpbr/cmd/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "build",
	Short: "build project based on json config",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		jsonFile, err := os.Open(args[0])
		if err != nil {
			fmt.Println("Fatal error: config.json missing")
			return err
		}

		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)

		if err != nil {
			fmt.Println("Fatal error: failed to read config.json")
			return err
		}

		var jsonConfig data.Targets
		json.Unmarshal(byteValue, &jsonConfig)

		if len(jsonConfig.Targets) == 0 {
			fmt.Println("Fatal error: no targets configured in config")
			return errors.New("no targets configured")
		}

		utils.TexturesetVersion = jsonConfig.Targets[0].Textureset_format

		fmt.Println(time.Now().String())

		fmt.Println("--- Cleaning workspace")
		err = clean.Cmd.RunE(cmd, nil)

		if err != nil {
			fmt.Println("Fatal error: Failed to clean")
			return err
		}

		fmt.Println("--- Download latest base assets")
		err = download.Cmd.RunE(cmd, []string{})

		if err != nil {
			fmt.Println("Fatal error: Failed to download assets")
			return err
		}

		fmt.Println("--- Prcoess PSDs")
		err = gen.ConvertPsdCmd.RunE(cmd, []string{utils.LocalPath(utils.Psds)})

		if err != nil {
			fmt.Println("Warning: Failed to convert PSDs")
		}

		fmt.Println("--- Copy custom configs")
		err = cp.Copy(utils.LocalPath(utils.SettingDir), utils.LocalPath(utils.OutDir))

		if err != nil {
			fmt.Println("Warning: Failed to copy custom configs")
		}

		entries, _ := os.ReadDir(utils.LocalPath(utils.BaseAssets))
		f := entries[0]

		for _, s := range utils.TargetAssets {
			fmt.Println("--- Create height files for " + s)
			p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)

			fmt.Println(p)

			err = common.Build(cmd, s, utils.LocalPath(p))

			if err != nil {
				fmt.Println("Fatal error: Failed to build item in pack - " + s)
				return err
			}
		}

		fmt.Println("--- Copy Overrides")
		err = cp.Copy(utils.LocalPath(utils.Overrides), utils.LocalPath(utils.OutDir+string(os.PathSeparator)+"textures"))

		if err != nil {
			fmt.Println("Warning: Failed to copy overrides")
		}

		for _, s := range utils.TargetAssets {
			fmt.Println("--- Create JSON files")
			p := utils.LocalPath(filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s))
			err = common.CreateMers(cmd, p)

			if err != nil {
				fmt.Println("Fatal error: Failed to create JSON for item in pack - " + s)
				return err
			}
		}

		fmt.Println("--- Create manifest")
		err = gen.ManifestCmd.RunE(cmd, []string{
			"",
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
			fmt.Println("Fatal error: failed to create manifest")
			return err
		}

		fmt.Println("--- Crush images")
		err = img.CrushCmd.RunE(cmd, []string{utils.LocalPath(utils.OutDir)})

		if err != nil {
			fmt.Println("Warning: failed to crush")
		}

		fmt.Println("--- Package Release")
		err = gen.PackageCmd.RunE(cmd, []string{utils.LocalPath(utils.OutDir)})

		if err != nil {
			fmt.Println("Warning : packaging failed ")
			return err
		}

		fmt.Println("--- OpenPBR complete")

		return nil
	},
}
