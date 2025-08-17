package main

import (

	// "openvibe/pkg/cli/version"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type RootCmd struct {
}

func New(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openvibe",
		Short: "Create Vibrant Textures Packs for Minecraft Bedrock",
	}

	cmd.AddCommand(NewDownloadCmd(fs))
	cmd.AddCommand(NewTransformCmd(fs))
	cmd.AddCommand(NewVersionCmd(fs))

	return cmd

	// return RootCmd{
	// 	OVCommand: OVCommand{
	// 		AppFS: fs,
	// 		,
	// 		SubCmds: []OVRunner{
	// 			// build.Cmd,
	// 			// clean.Cmd,
	// 			// config.Cmd,
	// 			// download.Cmd,
	// 			// gen.Cmd,
	// 			// parse.Cmd,
	// 			NewVersionCmd(fs),
	// 		},
	// 	},
	// }
}
