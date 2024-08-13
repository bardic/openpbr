package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var ConvertPsdCmd = &cobra.Command{
	Use:   "psds",
	Short: "processess psds folders and if needed copies to overrides",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		build(args[0])
		return nil
	},
}

func build(in string) error {
	utils.AppendLoadOut("Convert PSD: " + in)
	items, err := os.ReadDir(in)

	if err != nil {
		return nil
	}

	for _, item := range items {
		newIn := in + string(os.PathSeparator) + item.Name()
		out := strings.Replace(newIn, ".psd", ".png", 1)
		out = strings.Replace(out, "psds", "overrides", 1)

		if item.IsDir() {
			os.MkdirAll(out, os.ModePerm)
			if err := build(newIn); err != nil {
				return err
			}
		} else {
			if filepath.Ext(item.Name()) != ".psd" {
				continue
			}

			err := utils.PsdPng(newIn, out)
			if err != nil {
				fmt.Println("PSD-PNG :: " + err.Error())
			}
		}
	}

	return nil
}
