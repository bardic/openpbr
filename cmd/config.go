package cmd

import (
	"encoding/json"
	"os"

	"github.com/bardic/openpbr/vo"
)

type Config struct {
	vo.BaseConf
	Buildname      string
	Name           string
	Description    string
	Header_uuid    string
	Module_uuid    string
	Default_mer    string
	Version        string
	Author         string
	License        string
	URL            string
	Capibility     string
	HeightTemplate string
	MerTemplate    string
	ROffset        string
	GOffset        string
	BOffset        string
}

func (cmd *Config) Perform() error {
	conf, err := json.Marshal(cmd)

	if err != nil {
		return err
	}

	return os.WriteFile(cmd.Out, conf, 0644)
}

func (cmd *Config) SetOut(out string) {
	cmd.Out = out
}

func (cmd *Config) GetOut() string {
	return cmd.Out
}
