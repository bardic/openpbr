package gen

import (
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var HeightCmd = &cobra.Command{
	Use:   "height",
	Short: "create heightmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if b, err := utils.CheckForOverride(args[1]); err != nil || b {
			return err
		}

		return utils.CreateHeightMap(args[0], args[1])
	},
}
