package img

import (
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var TgaPngCmd = &cobra.Command{
	Use:   "tgapng",
	Short: "convert tgas to pngs",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.TgaPng(args[0], args[1])
	},
}
