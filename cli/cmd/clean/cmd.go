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
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	os.RemoveAll(utils.BaseAssets)
	os.RemoveAll(utils.BuildDir)
	os.RemoveAll(utils.Temp)

	os.MkdirAll(utils.BaseAssets, os.ModePerm)
	os.MkdirAll(utils.BuildDir, os.ModePerm)

	for _, s := range utils.TaretAssets {
		os.MkdirAll(utils.BuildDir+"/textures/"+s, os.ModePerm)
	}
}
