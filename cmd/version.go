package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of ob",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dat, err := os.ReadFile("VERSION")
		if err != nil {
			return
		}
		fmt.Println("Release Version: " + string(dat))
	},
}
