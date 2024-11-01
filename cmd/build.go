package cmd

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bardic/openpbr/utils"
	cp "github.com/otiai10/copy"
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

	var jsonConfig Target
	err = json.Unmarshal(byteValue, &jsonConfig)
	logE(err)

	// utils.TexturesetVersion = jsonConfig.Textureset_format

	utils.AppendLoadOut("--- Cleaning workspace")
	err = (&Clean{}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Download latest base assets")
	err = (&Download{}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Prcoess PSDs")
	err = (&ConvertPSD{
		Path: utils.LocalPath(utils.Psds),
	}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Copy custom configs")
	err = cp.Copy(
		utils.LocalPath(utils.SettingDir),
		utils.LocalPath(utils.OutDir),
	)
	logE(err)

	utils.AppendLoadOut("--- Convert TGA to PNG")
	err = (&CovertAndNormalize{
		Root: utils.LocalPath(utils.BaseAssets),
	}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Create height files for ")
	err = (&HeightMap{
		Root: utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
	}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Copy Overrides")
	err = cp.Copy(utils.LocalPath(utils.Overrides), utils.LocalPath(filepath.Join(utils.OutDir, "textures")))
	logE(err)

	utils.AppendLoadOut("--- Export Texture JSON ")
	err = (&TextureSet{
		Root: utils.LocalPath(filepath.Join(utils.OutDir, "textures")),
	}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Create manifest")
	err = (&Manifest{
		Name:        jsonConfig.Name,
		Description: jsonConfig.Description,
		Header_uuid: jsonConfig.Header_uuid,
		Module_uuid: jsonConfig.Module_uuid,
		Version:     jsonConfig.Version,
		Author:      jsonConfig.Author,
		License:     jsonConfig.License,
		URL:         jsonConfig.URL,
		Capibility:  jsonConfig.Capibility,
	}).Perform()
	logE(err)

	utils.AppendLoadOut("--- Package Release")
	err = (&PackBundle{
		InDir:  utils.LocalPath(utils.OutDir),
		OutDir: utils.OutDir,
	}).Perform()
	logE(err)

	utils.AppendLoadOut("--- OpenPBR complete")

	return nil
}

func logE(err error) {
	if err != nil {
		utils.AppendLoadOut(err.Error())
	}
}
