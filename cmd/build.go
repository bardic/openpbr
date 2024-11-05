package cmd

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bardic/openpbr/utils"
)

type Build struct {
	ConfigPath string
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
		Target: filepath.Join(utils.SettingDir, "shared"),
		Dest:   utils.OutDir,
	})
	cmds = append(cmds, &Copy{
		Target: filepath.Join(utils.SettingDir, jsonConfig.Textureset_format),
		Dest:   utils.OutDir,
	})
	cmds = append(cmds, &CovertAndNormalize{
		Root: utils.LocalPath(utils.BaseAssets),
	})
	cmds = append(cmds, &HeightMap{
		Root: utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
	})
	cmds = append(cmds, &AdjustColor{
		Root: utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
	})
	cmds = append(cmds, &Copy{
		Target: utils.Overrides,
		Dest:   filepath.Join(utils.OutDir, "textures"),
	})
	cmds = append(cmds, &TextureSet{
		Root:              utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
		TexturesetVersion: jsonConfig.Textureset_format,
	})
	cmds = append(cmds, &Manifest{
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
