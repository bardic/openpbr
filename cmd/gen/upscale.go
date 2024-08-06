package gen

import (
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var UpscaleCmd = &cobra.Command{
	Use:   "upscale",
	Short: "upscale image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Upscale(args[0], args[1])
	},
}
