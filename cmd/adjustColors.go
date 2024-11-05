package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/utils"
)

type AdjustColor struct {
	Root    string
	SubRoot string
	In      string
	Out     string
	ROffset int
	GOffset int
	BOffset int
}

func (cmd *AdjustColor) Perform() error {
	utils.AppendLoadOut("--- Adjust colours")
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

func (cmd *AdjustColor) walkDir() error {
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

func (cmd *AdjustColor) Exec() error {
	r := cmd.ROffset + 100
	g := cmd.GOffset + 100
	b := cmd.BOffset + 100

	rgb := fmt.Sprintf("%d,%d,%d", r, g, b)
	c := exec.Command(utils.ImCmd, cmd.In, "-modulate", rgb, "png32:"+cmd.In)

	return utils.RunCmd(c)
}
