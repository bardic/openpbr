package gen

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "generate files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	Cmd.AddCommand(HeightCmd)
	Cmd.AddCommand(JsonCmd)
	Cmd.AddCommand(UpscaleCmd)
	Cmd.AddCommand(ManifestCmd)
	Cmd.AddCommand(PackageCmd)
	Cmd.AddCommand(ConvertPsdCmd)
	Cmd.AddCommand(CreateCSVCmd)
	Cmd.AddCommand(ReadCSVCmd)
	Cmd.AddCommand(ConfigCmd)
}
