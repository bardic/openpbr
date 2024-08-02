package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "generate files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete note")
	},
}

func init() {
	Cmd.AddCommand(HeightCmd)
	Cmd.AddCommand(JsonCmd)
	Cmd.AddCommand(ManifestCmd)
}
