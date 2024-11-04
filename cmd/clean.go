package cmd

import (
	"os"

	"github.com/bardic/openpbr/utils"
)

type Clean struct {
}

func (c *Clean) Perform() error {
	utils.AppendLoadOut("--- Cleaning workspace")

	if e := os.RemoveAll(utils.BaseAssets); e != nil {
		return e
	}

	if e := os.RemoveAll(utils.OutDir); e != nil {
		return e
	}

	if e := os.RemoveAll(utils.Overrides); e != nil {
		return e
	}

	if e := os.MkdirAll(utils.BaseAssets, os.ModePerm); e != nil {
		return e
	}

	if e := os.MkdirAll(utils.OutDir, os.ModePerm); e != nil {
		return e
	}

	if e := os.MkdirAll(utils.Overrides, os.ModePerm); e != nil {
		return e
	}

	for _, s := range utils.TargetAssets {
		if e := os.MkdirAll(utils.OutDir+"/textures/"+s, os.ModePerm); e != nil {
			return e
		}

	}

	return nil

}
