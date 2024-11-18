package export

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/bardic/openpbr/utils"
)

type Manifest struct {
	Templates   embed.FS
	Out         string
	Name        string
	Header_uuid string
	Module_uuid string
	Description string
	Version     string
	Author      string
	License     string
	URL         string
	Capibility  string
}

func (cmd *Manifest) SetOut(out string) {
	cmd.Out = out
}

func (cmd *Manifest) GetOut() string {
	return cmd.Out
}

func (cmd *Manifest) Perform() error {
	utils.AppendLoadOut("--- Create manifest")

	b, err := utils.Templates.ReadFile((path.Join("templates", "manifest.tmpl")))
	t, err := template.New("a").Parse(string(b))

	if err != nil {
		fmt.Println(err)
		return err
	}

	f, err := os.Create(utils.LocalPath(utils.OutDir + string(os.PathSeparator) + "manifest.json"))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	err = t.Execute(f, cmd)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (cmd *Manifest) Save() error {
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
