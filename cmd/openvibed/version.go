package main

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewVersionCmd(fs afero.Fs) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("openvibe -- v0.0.1")
		},
	}
}
