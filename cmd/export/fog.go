package export

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/vo"
)

type Fog struct {
	vo.Fog
}

func (cmd *Fog) SetOut(out string) {
	cmd.Out = out
}

func (cmd *Fog) GetOut() string {
	return cmd.Out
}

func (cmd *Fog) Perform() error {
	tmplFile := "./templates/default_fog_settings.tmpl"

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
