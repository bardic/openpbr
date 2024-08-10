package img

import (
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var CrushCmd = &cobra.Command{
	Use:   "crush",
	Short: "crushes image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.CrushFiles(args[0])
		return nil
	},
}
