package gen

import (
	"fmt"
	"html/template"
	"os"

	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

type Manifest struct {
	Name        string
	Header_uuid string
	Module_uuid string
	Description string
	Version     string
	Author      string
	License     string
	URL         string
	Capibility  string
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
			Author:      args[5],
			License:     args[6],
			URL:         args[7],
			Capibility:  args[8],
		}

		var tmplFile = utils.LocalPath("templates" + string(os.PathSeparator) + "manifest.tmpl")

		t, err := template.ParseFiles(tmplFile)
		if err != nil {
			fmt.Println(err)
			return err
		}

		f, err := os.Create(utils.LocalPath(utils.OutDir + string(os.PathSeparator) + "manifest.json"))
		if err != nil {
			fmt.Println(err)
			return err
		}

		defer f.Close()

		err = t.Execute(f, manifest)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	},
}
