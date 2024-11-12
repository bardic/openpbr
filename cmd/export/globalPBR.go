package export

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/vo"
)

type GlobalPBR struct {
	Out string
	vo.PBR
}

func (cmd *GlobalPBR) Perform() error {
	tmplFile := "./templates/globalpbr.tmpl"

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
