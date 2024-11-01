package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/bardic/openpbr/utils"
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

func (cmd *Manifest) Perform() error {
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

	err = t.Execute(f, cmd)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
