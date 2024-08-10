package gen

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

type Manifest struct {
	Name        string
	Header_uuid string
	Module_uuid string
	Description string
	Version     string
}

var ManifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "generate's a manifest file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		manifest := Manifest{
			Name:        args[0],
			Description: args[1],
			Header_uuid: args[2],
			Module_uuid: args[3],
			Version:     args[4],
		}

		var tmplFile = "./templates/manifest.tmpl"

		t, err := template.ParseFiles(tmplFile)
		if err != nil {
			return err
		}

		f, err := os.Create(utils.OutDir + "/manifest.json")
		if err != nil {
			return err
		}

		defer f.Close()

		err = t.Execute(f, manifest)
		if err != nil {
			return err
		}

		return nil
	},
}
