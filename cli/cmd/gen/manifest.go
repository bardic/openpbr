package gen

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/bardic/openpbr/data"
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var ManifestCmd = &cobra.Command{
	Use:   "book",
	Short: "delete book",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var tmplFile = "manifest.tmpl"

		t, err := template.ParseFiles(tmplFile)
		if err != nil {
			return
		}

		f, err := os.Create(utils.BuildDir + "/manifest.json")
		if err != nil {
			return
		}

		defer f.Close()

		dat, err := os.ReadFile("VERSION")
		if err != nil {
			return
		}
		fmt.Println("Release Version: " + string(dat))
		vals := strings.Split(string(dat)[1:], ".")

		m := &data.Manifest{
			VersionStr: string(dat),
			VersionArr: "[" + vals[0] + "," + vals[1] + "," + vals[2] + "]",
		}

		err = t.Execute(f, m)
		if err != nil {
			log.Fatal(err)
		}
	},
}
