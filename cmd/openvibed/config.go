package main

import (
	"encoding/json"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	out, buildName, name, description,
	header_uuid, module_uuid, default_mer,
	version, author, license, url,
	capability, heightTemplate, merTemplate,
	rOffset, gOffset, bOffset string

	FS  afero.Fs
	Cmd = &cobra.Command{
		Use:   "config",
		Short: "creates the workspace config",
		Long:  `the data in the workspace config is used to generate the package config as well as offser defaults while processing.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := Config{
				BuildName:      buildName,
				Name:           name,
				Description:    description,
				HeaderUuid:     header_uuid,
				ModuleUuid:     module_uuid,
				DefaultMer:     default_mer,
				Version:        version,
				Author:         author,
				License:        license,
				URL:            url,
				Capability:     capability,
				HeightTemplate: heightTemplate,
				MerTemplate:    merTemplate,
				ROffset:        rOffset,
				GOffset:        gOffset,
				BOffset:        bOffset,
			}

			confStr, err := json.Marshal(config)

			if err != nil {
				return err
			}

			return os.WriteFile(out, confStr, 0644)
		},
	}
)

func init() {
	Cmd.Flags().StringVarP(&out, "out", "", "", "out dir")
	Cmd.Flags().StringVarP(&buildName, "build-name", "b", "", "build name")             //TODO: Is this the package or workspace name?
	Cmd.Flags().StringVarP(&name, "name", "n", "", "name")                              //TODO: How is this different than build name?
	Cmd.Flags().StringVarP(&description, "description", "d", "", "package description") //TODO: Is this for the mcpack conf?
	Cmd.Flags().StringVarP(&header_uuid, "header-uuid", "", "", "MCPack Header UUID")
	Cmd.Flags().StringVarP(&module_uuid, "module-uuid", "m", "", "MCPack Module UUID")
	Cmd.Flags().StringVarP(&default_mer, "default-mer", "", "", "default mer values") //TODO: Is this used?
	Cmd.Flags().StringVarP(&version, "version", "v", "", "set package version")
	Cmd.Flags().StringVarP(&author, "author", "a", "", "set mcpack author name")
	Cmd.Flags().StringVarP(&license, "license", "l", "", "mcpack license")
	Cmd.Flags().StringVarP(&url, "url", "u", "", "mcpack url")
	Cmd.Flags().StringVarP(&capability, "capability", "c", "", "if pbr files should be generated")
	Cmd.Flags().StringVarP(&heightTemplate, "height-template", "", "", "suffix used to denote height maps")
	Cmd.Flags().StringVarP(&merTemplate, "mer-template", "", "", "suffix used to denote mer maps")
	Cmd.Flags().StringVarP(&rOffset, "r-offset", "", "", "default red offset")
	Cmd.Flags().StringVarP(&bOffset, "b-offset", "", "", "default blue offset")
	Cmd.Flags().StringVarP(&gOffset, "g-offset", "", "", "default green offset")
}
