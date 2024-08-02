package gen

import (
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var HeightCmd = &cobra.Command{
	Use:   "book",
	Short: "delete book",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err, b := utils.CheckForOverride(args[1]); err != nil || b {
			return
		}

		utils.CreateHeightMap(args[0], args[1])
	},
}
