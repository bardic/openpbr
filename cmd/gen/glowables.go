package gen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var GlowablesCmd = &cobra.Command{
	Use:   "glowables",
	Short: "processess glowables folders and if needed copies to overrides",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		build(args[0])

		return nil
	},
}

func build(in string) error {
	items, err := os.ReadDir(in)

	if err != nil {
		return nil
	}

	for _, item := range items {
		if item.IsDir() {
			if err := build(in + "/" + item.Name()); err != nil {
				return err
			}
		} else {
			if filepath.Ext(item.Name()) != ".psd" {
				continue
			}
			c := exec.Command("convert", in+"/"+item.Name()+"[0]/", utils.Overrides+strings.Replace(item.Name(), ".psd", ".png", 1))
			if e := c.Run(); e != nil {
				return nil
			}
		}
	}

	return nil
}
