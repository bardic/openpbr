package gen

import (
	"encoding/json"
	"os"

	"github.com/bardic/openpbr/cmd/data"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "create normalmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := data.Targets{
			Targets: []data.Target{
				{
					Buildname:         "",
					Name:              args[1],
					Header_uuid:       args[2],
					Description:       args[3],
					Module_uuid:       args[4],
					Textureset_format: args[5],
					Default_mer:       args[6],
					Version:           args[7],
				},
			},
		}

		rankingsJson, _ := json.Marshal(c)
		return os.WriteFile(args[0], rankingsJson, 0644)
	},
}