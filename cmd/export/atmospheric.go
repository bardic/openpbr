package export

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/vo"
)

type Atmospherics struct {
	Out string
	vo.Atmospherics
}

func (cmd *Atmospherics) SetOut(out string) {
	cmd.Out = out
}

func (cmd *Atmospherics) GetOut() string {
	return cmd.Out
}

func (cmd *Atmospherics) Perform() error {
	tmplFile := "./templates/atmospherics.tmpl"

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
