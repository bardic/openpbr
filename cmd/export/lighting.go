package export

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/vo"
)

type Lighting struct {
	Out string
	vo.Lighting
}

func (cmd *Lighting) Perform() error {
	tmplFile := "./templates/pbr.tmpl"

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
