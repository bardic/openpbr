package gen

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/data"
	"github.com/spf13/cobra"
)

var JsonCmd = &cobra.Command{
	Use:   "book",
	Short: "delete book",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		out := args[0]
		color := args[1]
		merArr := args[2]
		merFile := args[3]
		height := args[4]
		var tmplFile = "pbr.tmpl"

		pbr := data.PBR{
			Colour:  color,
			MerArr:  merArr,
			MerFile: merFile,
			Height:  height,
		}

		t, err := template.ParseFiles(tmplFile)
		if err != nil {
			return
		}

		f, err := os.Create(out)
		if err != nil {
			return
		}

		defer f.Close()

		t.Execute(f, pbr)
	},
}
