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
		build(args[0])
		return nil
	},
}

func build(in string) error {
	fmt.Println("Convert PSD: " + in)
	subPaths := strings.Split(in, string(os.PathSeparator))
	items, err := os.ReadDir(in)

	if err != nil {
		return nil
	}

	for _, item := range items {
		outPath := utils.OutDir + string(os.PathSeparator) + strings.Join(subPaths[0:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name()
		itemPath := in + string(os.PathSeparator) + item.Name()

		fmt.Println("OUt path " + outPath)
		fmt.Println("Item : " + itemPath)

		if item.IsDir() {
			if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
				return err
			}

			if err := build(itemPath); err != nil {
				return err
			}
		} else {
			if filepath.Ext(item.Name()) != ".psd" {
				continue
			}

			err := utils.PsdPng(itemPath, strings.Replace(outPath, ".psd", ".png", 1))
			if err != nil {
				fmt.Println("PSD-PNG :: " + err.Error())
			}
		}
	}

	return nil
}
