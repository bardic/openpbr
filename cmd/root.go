package cmd

import (
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
