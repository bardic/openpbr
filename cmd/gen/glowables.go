package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var ConvertPsdCmd = &cobra.Command{
	Use:   "psds",
	Short: "processess psds folders and if needed copies to overrides",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cats")
		build(args[0])

		return nil
	},
}

func build(in string) error {
	fmt.Println(in)
	items, err := os.ReadDir(in)

	fmt.Println(err)
	if err != nil {
		return nil
	}

	for _, item := range items {
		fmt.Println(item.Name())
		if item.IsDir() {
			if err := build(in + "/" + item.Name()); err != nil {
				return err
			}
		} else {
			if filepath.Ext(item.Name()) != ".psd" {
				continue
			}

			fmt.Println(in + "/" + item.Name())
			err := utils.PsdPng(in+"/"+item.Name(), utils.Overrides+"/"+strings.Replace(item.Name(), ".psd", ".png", 1))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
