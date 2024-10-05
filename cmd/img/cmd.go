package img

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "img",
	Short: "manipulates image files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(TgaPngCmd)
}
