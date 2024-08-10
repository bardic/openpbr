package img

import (
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var PsdPngCmd = &cobra.Command{
	Use:   "psdpng",
	Short: "convert psds to pngs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PsdPng(args[0], args[1])
	},
}
