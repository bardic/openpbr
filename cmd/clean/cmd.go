package clean

import (
	"os"

	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "clean",
	Short: "delete and regenerates release workspace",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if e := os.RemoveAll(utils.BaseAssets); e != nil {
			return e
		}

		if e := os.RemoveAll(utils.BuildDir); e != nil {
			return e
		}

		if e := os.MkdirAll(utils.BaseAssets, os.ModePerm); e != nil {
			return e
		}
		if e := os.MkdirAll(utils.BuildDir, os.ModePerm); e != nil {
			return e
		}

		for _, s := range utils.TargetAssets {
			if e := os.MkdirAll(utils.BuildDir+"/textures/"+s, os.ModePerm); e != nil {
				return e
			}

		}

		return nil
	},
}
