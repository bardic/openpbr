package cmd

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/bardic/openpbr/utils"
)

type CovertAndNormalize struct {
	Root    string
	SubRoot string
	In      string
	Out     string
}

func (cmd *CovertAndNormalize) Perform() error {
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

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		cmd.In = filepath.Join(root, path)
		subpath, err := utils.GetTextureSubpath(cmd.In, "textures")
		cmd.Out = filepath.Join(utils.LocalPath(utils.OutDir), subpath)

		if filepath.Ext(path) == ".tga" {
			cmd.Out = strings.Replace(cmd.Out, ".tga", ".png", -1)
		}

		os.MkdirAll(filepath.Dir(cmd.Out), os.ModePerm)

		if err != nil {
			return nil
		}

		cmd.Exec()

		return nil
	})

	return nil
}

func (cmd *CovertAndNormalize) Exec() error {
	c := exec.Command(
		utils.IM_CMD,
		cmd.In,
		"png32:"+cmd.Out,
	)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
	return nil
}
