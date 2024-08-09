package gen

import (
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/bardic/openpbr/data"
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var JsonCmd = &cobra.Command{
	Use:   "json",
	Short: "create deferred json files",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("json out" + args[0])

		out := args[0]
		color := args[1]
		merArr := args[2]
		merFile := args[3]
		height := args[4]
		useMerFile, _ := strconv.ParseBool(args[5])
		texturesetVersion := args[6]

		var tmplFile = "./templates/pbr.tmpl"

		if utils.NormalMaps {
			tmplFile = "./templates/pbr_normal.tmpl"
		}

		if texturesetVersion == "1.21.30" {
			tmplFile = "./templates/pbr2.tmpl"

			if utils.NormalMaps {
				tmplFile = "./templates/pbr2_normal.tmpl"
			}
		}

		pbr := data.PBR{
			Colour:  color,
			MerArr:  merArr,
			MerFile: merFile,
			Height:  height,
			MerType: useMerFile,
		}

		t, err := template.ParseFiles(tmplFile)
		if err != nil {
			return err
		}

		f, err := os.Create(out)
		if err != nil {
			return err
		}

		defer f.Close()

		if err := t.Execute(f, pbr); err != nil {
			return err
		}

		return nil
	},
}
