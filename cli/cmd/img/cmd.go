package img

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "img",
	Short: "manipulates image files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete note")
	},
}

func init() {
	Cmd.AddCommand(AdjustColorCmd)
	Cmd.AddCommand(TgaPngCmd)
}
