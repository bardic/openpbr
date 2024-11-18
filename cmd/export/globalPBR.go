package export

import (
	"html/template"
	"os"

	"github.com/bardic/openpbr/vo"
)

type GlobalPBR struct {
	vo.PBR
}

func (cmd *GlobalPBR) SetOut(out string) {
	cmd.Out = out
}

func (cmd *GlobalPBR) GetOut() string {
	return cmd.Out
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
