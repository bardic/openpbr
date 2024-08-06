package gen

import (
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var NormalCmd = &cobra.Command{
	Use:   "normal",
	Short: "create normalmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if b, err := utils.CheckForOverride(args[1]); err != nil || b {
			return err
		}

		return utils.CreateNormalMap(args[0], args[1])
	},
}
