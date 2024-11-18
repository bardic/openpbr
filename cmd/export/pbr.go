package export

import (
	"embed"
	"encoding/json"
	"html/template"
	"os"

	"github.com/bardic/openpbr/utils"
)

type PBR struct {
	templates  embed.FS
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
	tmplFile := "templates/pbr.tmpl"

	if cmd.Capibility == "pbr" {
		tmplFile = "templates/pbr2.tmpl"
		cmd.MerArr = cmd.MerArr + "FF"
	}

	b, err := utils.Templates.ReadFile(tmplFile)
	t, err := template.New("a").Parse(string(b))

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

func (cmd *PBR) Save() error {
	f, err := os.Create(cmd.Out + ".dat")
	if err != nil {
		return err
	}

	defer f.Close()

	b, err := json.Marshal(cmd)

	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	return nil
}
