package gen

import (
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var NormalCmd = &cobra.Command{
	Use:   "normal",
	Short: "create normalmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.CreateNormalMap(args[0], args[1])
	},
}
