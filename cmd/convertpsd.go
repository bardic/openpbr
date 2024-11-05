package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/utils"
)

type ConvertPSD struct {
	Path string
}

func (cmd *ConvertPSD) Perform() error {
	utils.AppendLoadOut("--- Prcoess PSDs")

	return build(cmd.Path)
}

func build(in string) error {
	utils.AppendLoadOut("Convert PSD: " + in)
	items, err := os.ReadDir(in)

	if err != nil {
		return nil
	}

	for _, item := range items {
		newIn := filepath.Join(in, item.Name())
		out := strings.Replace(newIn, ".psd", ".png", 1)
		out = strings.Replace(out, "psds", "overrides", 1)

		if item.IsDir() {
			os.MkdirAll(out, os.ModePerm)
			if err := build(newIn); err != nil {
				return err
			}
		} else {
			if filepath.Ext(item.Name()) != ".psd" {
				continue
			}

			c := exec.Command(utils.ImCmd, newIn+"[0]", "png32:"+out)

			err := utils.RunCmd(c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
