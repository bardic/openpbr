package cmd

import (
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/widget"

	"embed"

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
func Execute(templates embed.FS) error {
	UI(templates)
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(clean.Cmd)
	RootCmd.AddCommand(download.Cmd)
	RootCmd.AddCommand(gen.Cmd)
	RootCmd.AddCommand(build.Cmd)
	RootCmd.AddCommand(versionCmd)
}

func Build(args []string) error {
	return build.Cmd.RunE(RootCmd, args)
}

func CreateManifest(args []string) {
	gen.ConfigCmd.RunE(RootCmd, args)
}

func CreateCSV(p string, defaultMer string) {
	gen.CreateCSVCmd.RunE(RootCmd, []string{p, defaultMer})
}
