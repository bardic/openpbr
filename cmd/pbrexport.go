package cmd

import (
	"html/template"
	"os"
)

type PBRExport struct {
	Out           string
	Color         string
	MerArr        string
	MerFile       string
	Height        string
	UseMerFile    bool
	TextureSetVer string
}

func (cmd *PBRExport) Perform() error {

	tmplFile := "./templates/pbr.tmpl"

	if cmd.TextureSetVer == "1.21.30" {
		tmplFile = "./templates/pbr2.tmpl"

	}

	pbr := PBR{
		Colour:  cmd.Color,
		MerArr:  cmd.MerArr,
		MerFile: cmd.MerFile,
		Height:  cmd.Height,
		MerType: cmd.UseMerFile,
	}

	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	f, err := os.Create(cmd.Out)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := t.Execute(f, pbr); err != nil {
		return err
	}

	return nil

}
