package cmd

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/utils"
)

type TextureSet struct {
	Root    string
	SubRoot string
	In      string
	Out     string
}

func (cmd *TextureSet) Perform() error {
	for _, s := range utils.TargetAssets {
		utils.AppendLoadOut("--- Create JSON files")
		cmd.SubRoot = filepath.Join(cmd.Root, s)
		err := cmd.CreateTextureSets()

		if err != nil {
			utils.AppendLoadOut("Fatal error: Failed to create JSON for item in pack - " + s)
			return err
		}
	}

	return nil
}

func (cmd *TextureSet) CreateTextureSets() error {

	root := cmd.SubRoot
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() {
			return nil
		}

		in := d.Name()

		if strings.Contains(in, "_height") || strings.Contains(in, "texture_set") || strings.Contains(in, "_mer") {
			return nil
		}

		cmd.In = filepath.Join(root, path)
		subpath, err := utils.GetTextureSubpath(cmd.In, "textures")

		if err != nil {
			return nil
		}

		cmd.Out = filepath.Join(utils.LocalPath(utils.OutDir), subpath)

		in = strings.TrimSuffix(in, ".png")
		merPath := in + utils.MerMapNameSuffix
		heightPath := in + utils.HeightMapNameSuffix
		out := strings.Replace(cmd.Out, ".png", ".texture_set.json", -1)

		useMerFile := true
		if _, err := os.Stat(merPath + ".png"); err != nil {
			useMerFile = false
		}

		err = (&PBRExport{
			Out:           out,
			Color:         in,
			MerArr:        "#0000FF11",
			MerFile:       merPath,
			Height:        heightPath,
			UseMerFile:    useMerFile,
			TextureSetVer: utils.TexturesetVersion,
		}).Perform()

		if err != nil {
			return err
		}

		return nil
	})

	return nil
}
