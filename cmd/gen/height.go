package gen

import (
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var HeightCmd = &cobra.Command{
	Use:   "height",
	Short: "create heightmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.CreateHeightMap(args[0], args[1])
	},
}
