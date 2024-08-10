package build

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/common"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/cmd/img"
	"github.com/bardic/openpbr/utils"
	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

type Targets struct {
	Targets []Target
}

type Target struct {
	Buildname         string
	Name              string
	Header_uuid       string
	Module_uuid       string
	Description       string
	Textureset_format string
	Default_mer       string
	Version           string
}

var Cmd = &cobra.Command{
	Use:   "build",
	Short: "build project based on json config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Open our jsonFile
		jsonFile, err := os.Open("config.json")
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		var targets Targets
		json.Unmarshal(byteValue, &targets)

		fmt.Println(targets)

		fmt.Println(time.Now().String())

		fmt.Println("--- Cleaning workspace")
		clean.Cmd.RunE(cmd, nil)

		fmt.Println("--- Download latest base assets")
		download.Cmd.RunE(cmd, []string{"skip"})

		fmt.Println("--- Prcoess PSDs")
		gen.ConvertPsdCmd.RunE(cmd, []string{utils.Psds})

		fmt.Println("--- Copy custom configs")
		cp.Copy(utils.SettingDIr, utils.OutDir)

		entries, _ := os.ReadDir(utils.BaseAssets)
		f := entries[0]

		for _, s := range utils.TargetAssets {
			fmt.Println("--- Create height files for " + s)
			p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)
			common.Build(cmd, s, p)
		}

		fmt.Println("--- Copy Overrides")
		cp.Copy(utils.Overrides, utils.OutDir+string(os.PathSeparator)+"textures")

		for _, s := range utils.TargetAssets {
			fmt.Println("--- Create JSON files")
			p := filepath.Join(utils.BaseAssets, f.Name(), "resource_pack", "textures", s)
			common.CreateMers(cmd, p)
		}

		fmt.Println("--- Create manifest")
		gen.ManifestCmd.RunE(cmd, []string{
			targets.Targets[0].Name,
			targets.Targets[0].Description,
			targets.Targets[0].Header_uuid,
			targets.Targets[0].Module_uuid,
			targets.Targets[0].Version,
		})

		fmt.Println("--- Crush images")
		img.CrushCmd.RunE(cmd, []string{utils.OutDir})

		gen.PackageCmd.RunE(cmd, []string{utils.OutDir})

		dat, err := os.ReadFile("VERSION")
		if err != nil {
			return
		}

		fmt.Println("Release Version: " + string(dat))

	},
}
