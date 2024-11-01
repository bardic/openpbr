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

type HeightMap struct {
	Root    string
	SubRoot string
	In      string
	Out     string
}

func (cmd *HeightMap) Perform() error {
	for _, s := range utils.TargetAssets {
		utils.AppendLoadOut("--- --- " + s)
		cmd.SubRoot = filepath.Join(cmd.Root, s)
		err := cmd.walkDir()

		if err != nil {
			utils.AppendLoadOut("Fatal error: Failed to build item in pack - " + s)
			return err
		}
	}

	return nil
}

func (cmd *HeightMap) walkDir() error {
	root := cmd.SubRoot
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() {
			return nil
		}

		cmd.In = filepath.Join(root, path)
		cmd.Out = strings.ReplaceAll(cmd.In, ".png", utils.HeightMapNameSuffix+".png")

		cmd.Exec()

		return nil
	})

	return nil
}

func (cmd *HeightMap) Exec() error {

	c := exec.Command(
		utils.IM_CMD,
		cmd.In,
		"-channel",
		"RGB",
		"-negate",
		"-set",
		"colorspace",
		"Gray",
		"png32:"+cmd.Out,
	)

	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()

	return nil

}
