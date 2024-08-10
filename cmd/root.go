package cmd

import (
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/widget"

	"github.com/bardic/openpbr/cmd/build"
	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "openpbr",
		Short: "A generator for Deferred Textured Packs",
		Long:  ``,
	}
)

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(clean.Cmd)
	RootCmd.AddCommand(download.Cmd)
	RootCmd.AddCommand(gen.Cmd)
	RootCmd.AddCommand(build.Cmd)
	RootCmd.AddCommand(versionCmd)
}

func Build(args []string) {
	build.Cmd.RunE(RootCmd, args)
}

func CreateManifest(args []string) {
	/*
		Name:        args[0],
					Description: args[1],
					Header_uuid: args[2],
					Module_uuid: args[3],
					Version:     args[4],
	*/
	gen.ConfigCmd.RunE(RootCmd, args)
}
