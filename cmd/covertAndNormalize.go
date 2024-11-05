package cmd

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/utils"
)

type CovertAndNormalize struct {
	Root    string
	SubRoot string
	In      string
	Out     string
}

func (cmd *CovertAndNormalize) Perform() error {
	utils.AppendLoadOut("--- Normalize source images")
	entries, _ := os.ReadDir(cmd.Root)
	f := entries[0]
	for _, s := range utils.TargetAssets {
		utils.AppendLoadOut("--- --- " + s)
		cmd.SubRoot = filepath.Join(cmd.Root, f.Name(), "resource_pack", "textures", s)
		err := cmd.createTGA()

		if err != nil {
			utils.AppendLoadOut("Fatal error: Failed to build item in pack - " + s)
			return err
		}
	}

	return nil
}

func (cmd *CovertAndNormalize) createTGA() error {
	root := cmd.SubRoot
	fileSystem := os.DirFS(root)

	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		cmd.In = filepath.Join(root, path)
		subpath, err := utils.GetTextureSubpath(cmd.In, "textures")
		if err != nil {
			return err
		}
		cmd.Out = filepath.Join(utils.LocalPath(utils.OutDir), subpath)

		if filepath.Ext(path) == ".tga" {
			cmd.Out = strings.Replace(cmd.Out, ".tga", ".png", 1)
		}

		err = os.MkdirAll(filepath.Dir(cmd.Out), os.ModePerm)

		if err != nil {
			return err
		}

		return cmd.Exec()
	})

	return err
}

func (cmd *CovertAndNormalize) Exec() error {
	c := exec.Command(
		utils.ImCmd,
		cmd.In,
		"png32:"+cmd.Out,
	)

	return utils.RunCmd(c)
}
