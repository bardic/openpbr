package cmd

import (
	"encoding/json"
	"os"
)

type Config struct {
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

	return os.WriteFile(cmd.Buildname, conf, 0644)
}
