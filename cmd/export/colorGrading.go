package export

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type ColorGrading struct {
	vo.ColorGrading
}

func (cmd *ColorGrading) SetOut(out string) {
	cmd.Out = out
}

func (cmd *ColorGrading) GetOut() string {
	return cmd.Out
}

func (cmd *ColorGrading) Perform() error {
	tmplFile := "./templates/color_grading.tmpl"

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

func (cmd *ColorGrading) Save() error {
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
