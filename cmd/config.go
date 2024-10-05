package cmd

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "create normalmaps based on colour image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		exportMer, err := strconv.ParseBool(args[15])

		if err != nil {
			return err
		}

		c := Targets{
			Targets: []Target{
				{
					Buildname:         "",
					Name:              args[1],
					Description:       args[2],
					Header_uuid:       args[3],
					Module_uuid:       args[4],
					Textureset_format: args[5],
					Default_mer:       args[6],
					Version:           args[7],
					Author:            args[8],
					License:           args[9],
					URL:               args[10],
					Capibility:        args[11],
					HeightTemplate:    args[12],
					NormalTemplate:    args[13],
					MerTemplate:       args[14],
					ExportMer:         exportMer,
				},
			},
		}

		rankingsJson, _ := json.Marshal(c)
		return os.WriteFile(args[0], rankingsJson, 0644)
	},
}
