package export

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/vo"
)

type Water struct {
	Out string
	vo.Water
}

func (cmd *Water) Perform() error {
	tmplFile := "./templates/water.tmpl"

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
