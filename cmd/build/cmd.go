package build

import (
	"encoding/json"
	"errors"
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
		utils.LoadStdOut.Text()
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

		entries, _ := os.ReadDir(utils.LocalPath(utils.BaseAssets))
		f := entries[0]

		for _, s := range utils.TargetAssets {
			utils.AppendLoadOut("--- Create height files for " + s)
			p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)

			utils.AppendLoadOut(p)

			err = common.Build(cmd, s, utils.LocalPath(p))

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
			p := utils.LocalPath(filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s))
			err = common.CreateMers(cmd, p)

			if err != nil {
				utils.AppendLoadOut("Fatal error: Failed to create JSON for item in pack - " + s)
				return err
			}
		}

		utils.AppendLoadOut("--- Create manifest")
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
			utils.AppendLoadOut("Fatal error: failed to create manifest")
			return err
		}

		utils.AppendLoadOut("--- Crush images")
		err = img.CrushCmd.RunE(cmd, []string{utils.LocalPath(utils.OutDir)})

		if err != nil {
			utils.AppendLoadOut("Warning: failed to crush")
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
