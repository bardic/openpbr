package gen

import (
	"html/template"
	"os"
	"strings"

	"github.com/bardic/openpbr/data"
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var ManifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "generate's a manifest file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		dat, err := os.ReadFile("VERSION")
		if err != nil {
			return err
		}

		vals := strings.Split(string(dat)[1:], ".")

		m := &data.Manifest{
			VersionStr: string(dat),
			VersionArr: "[" + vals[0] + "," + vals[1] + "," + vals[2] + "]",
		}

		err = t.Execute(f, m)
		if err != nil {
			return err
		}

		return nil
	},
}
