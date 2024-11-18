package export

import (
	"html/template"
	"os"
)

type PBR struct {
	Out        string
	Colour     string
	MerArr     string
	MerFile    string
	Height     string
	UseMerFile bool
	Capibility string
}

func (cmd *PBR) SetOut(out string) {
	cmd.Out = out
}

func (cmd *PBR) GetOut() string {
	return cmd.Out
}

func (cmd *PBR) Perform() error {
	tmplFile := "./templates/pbr.tmpl"

	if cmd.Capibility == "pbr" {
		tmplFile = "./templates/pbr2.tmpl"
		cmd.MerArr = cmd.MerArr + "FF"
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

	if err := t.Execute(f, cmd); err != nil {
		return err
	}

	return nil
}
