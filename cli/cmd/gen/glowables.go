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
	Short: "delete book",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		build(args[0])
	},
}

func build(in string) {
	items, _ := os.ReadDir(in)
	for _, item := range items {
		if item.IsDir() {
			build(in + "/" + item.Name())
		} else {
			if filepath.Ext(item.Name()) != ".psd" {
				continue
			}
			c := exec.Command("convert", in+"/"+item.Name()+"[0]", utils.Overrides+strings.Replace(item.Name(), ".psd", ".png", 1))
			c.Run()
		}
	}
}
