package cmd

import (
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
	RootCmd.AddCommand(CleanCmd)
	RootCmd.AddCommand(DownloadCmd)
	RootCmd.AddCommand(BuildCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(HeightCmd)
	RootCmd.AddCommand(JsonCmd)
	RootCmd.AddCommand(ManifestCmd)
	RootCmd.AddCommand(PackageCmd)
	RootCmd.AddCommand(ConvertPsdCmd)
	RootCmd.AddCommand(ConfigCmd)
}

func Build(args []string) error {
	return BuildCmd.RunE(RootCmd, args)
}

func CreateManifest(args []string) {
	ConfigCmd.RunE(RootCmd, args)
}
