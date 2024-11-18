package cmd

import (
	"embed"
	"encoding/json"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
)

type Build struct {
	ConfigPath string
	Templates  embed.FS
}

func (cmd *Build) Perform() error {

	utils.AppendLoadOut(time.Now().String())

	jsonFile, err := os.Open(cmd.ConfigPath)
	logE(err)

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	logE(err)

	var jsonConfig Config
	err = json.Unmarshal(byteValue, &jsonConfig)
	logE(err)

	cmds := []ICmd{}
	cmds = append(cmds, &Clean{})
	cmds = append(cmds, &Download{})
	cmds = append(cmds, &ConvertPSD{
		Path: utils.LocalPath(utils.Psds),
	})
	cmds = append(cmds, &Copy{
		Target: filepath.Join("export", utils.SettingDir, "shared"),
		Dest:   utils.OutDir,
	})
	cmds = append(cmds, &Copy{
		Target: filepath.Join("export", utils.SettingDir, jsonConfig.Capibility),
		Dest:   utils.OutDir,
	})
	cmds = append(cmds, &CovertAndNormalize{
		Root: utils.LocalPath(utils.BaseAssets),
	})
	if jsonConfig.Capibility == "pbr" {
		cmds = append(cmds, &HeightMap{
			Root: utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
		})
	}
	cmds = append(cmds, &AdjustColor{
		Root: utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
	})
	cmds = append(cmds, &Copy{
		Target: utils.Overrides,
		Dest:   filepath.Join(utils.OutDir, "textures"),
	})
	cmds = append(cmds, &TextureSet{
		Root:       path.Join(utils.Basedir, utils.OutDir, "textures"),
		Capibility: jsonConfig.Capibility,
	})
	cmds = append(cmds, &export.Manifest{
		Templates:   cmd.Templates,
		Name:        jsonConfig.Name,
		Description: jsonConfig.Description,
		Header_uuid: jsonConfig.Header_uuid,
		Module_uuid: jsonConfig.Module_uuid,
		Version:     jsonConfig.Version,
		Author:      jsonConfig.Author,
		License:     jsonConfig.License,
		URL:         jsonConfig.URL,
		Capibility:  jsonConfig.Capibility,
	})
	cmds = append(cmds, &PackBundle{
		InDir:  utils.LocalPath(utils.OutDir),
		OutDir: utils.OutDir,
	})

	logE(Exec(cmds))

	return nil
}

func Exec(cmds []ICmd) error {
	for _, c := range cmds {
		err := c.Perform()
		if err != nil {
			return err
		}
	}

	utils.AppendLoadOut("--- OpenPBR complete")
	return nil
}

func logE(err error) {
	if err != nil {
		utils.AppendLoadOut(err.Error())
	}
}
