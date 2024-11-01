package cmd

import (
	"encoding/json"
	"os"
)

type Config struct {
	Buildname         string
	Name              string
	Description       string
	Header_uuid       string
	Module_uuid       string
	Textureset_format string
	Default_mer       string
	Version           string
	Author            string
	License           string
	URL               string
	Capibility        string
	HeightTemplate    string
	NormalTemplate    string
	MerTemplate       string
	ExportMer         string
}

func (cmd *Config) Perform() error {
	rankingsJson, _ := json.Marshal(cmd)
	return os.WriteFile(cmd.Buildname, rankingsJson, 0644)
}
