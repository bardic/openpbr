package cmd

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
)

type TextureSet struct {
	Root       string
	SubRoot    string
	In         string
	Out        string
	Capibility string
}

func (cmd *TextureSet) Perform() error {
	utils.AppendLoadOut("--- Export Texture JSON ")

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

	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
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
			return err
		}

		cmd.Out = filepath.Join(utils.LocalPath(utils.OutDir), "textures", subpath)

		in = strings.TrimSuffix(in, ".png")
		merPath := in + utils.MerMapNameSuffix
		heightPath := in + utils.HeightMapNameSuffix
		out := strings.Replace(cmd.Out, ".png", ".texture_set.json", -1)

		useMerFile := true
		if _, err := os.Stat(merPath + ".png"); err != nil {
			useMerFile = false
		}

		err = (&export.PBR{
			Out:        out,
			Colour:     in,
			MerArr:     "#0000FF",
			MerFile:    merPath,
			Height:     heightPath,
			UseMerFile: useMerFile,
			Capibility: cmd.Capibility,
		}).Perform()

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
