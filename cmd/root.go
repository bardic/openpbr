package cmd

import (
	"github.com/bardic/openpbr/cmd/build"
	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/utils"
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

	build.Cmd.Flags().BoolVarP(&utils.Beta, "beta", "b", false, "Perform image manipulation")
	build.Cmd.Flags().BoolVarP(&utils.DeleteAutoGen, "delete", "d", false, "Delete all autogenerated files from working dir")
	build.Cmd.Flags().BoolVarP(&utils.SkipDownload, "skip", "s", false, "Skip downloading assets zip")
	build.Cmd.Flags().BoolVarP(&utils.NormalMaps, "normal", "n", false, "Generate normals instead of height maps")
	build.Cmd.Flags().BoolVarP(&utils.ZipOnly, "zip", "z", false, "Creates a zip file of the output")
	build.Cmd.Flags().BoolVarP(&utils.Crush, "crush", "c", false, "Uses PNGCrush to reduce file size")
	build.Cmd.Flags().StringVarP(&utils.TexturesetVersion, "version", "v", "", "set to 1.21.30 for new format")

	RootCmd.AddCommand(clean.Cmd)
	RootCmd.AddCommand(download.Cmd)
	RootCmd.AddCommand(gen.Cmd)
	RootCmd.AddCommand(build.Cmd)
	RootCmd.AddCommand(versionCmd)
}
