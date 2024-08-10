package img

import (
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var AdjustColorCmd = &cobra.Command{
	Use:   "adjust",
	Short: "color adjusts to base assets",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		utils.AdjustColor(args[0])
	},
}
