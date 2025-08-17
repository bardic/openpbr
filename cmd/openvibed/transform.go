package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type transformCmd struct {
	appfs        afero.Fs
	targetAssets []string
	samplesIn    string
	texturesOut  string
	transforms   []transformConfig
	capability   string
	modifiers    []string
}

type transformConfig struct {
	// subRoot string
	in     string
	out    string
	params []string
}

func NewTransformCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transform",
		Short: "parse input to generate mers and maps",
		RunE: func(cmd *cobra.Command, args []string) error {
			targets := viper.GetStringSlice("target-assets")
			in := viper.GetString("in")
			out := viper.GetString("out")
			capability := viper.GetString("capability")
			modifiers := viper.GetStringSlice("modifiers")

			c := transformCmd{
				appfs:        fs,
				targetAssets: targets,
				samplesIn:    in,
				texturesOut:  out,
				capability:   capability,
				modifiers:    modifiers,
			}

			if err := c.parse(); err != nil {
				return err
			}

			fmt.Println(c)

			if err := c.exec(); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringSlice("target-assets", []string{"textures", "entity", "particle"}, "Entites to parse")
	cmd.Flags().String("in", "./", "samples dir in")
	cmd.Flags().String("out", "./out", "out dir for transformed textures")
	cmd.Flags().String("capability", "pbr", "-")
	cmd.Flags().StringSlice("modifiers", nil, "CSV of arg+value that relates to ImageMagick")

	viper.BindPFlag("target-assets", cmd.Flags().Lookup("target-assets"))
	viper.BindPFlag("in", cmd.Flags().Lookup("in"))
	viper.BindPFlag("out", cmd.Flags().Lookup("out"))
	viper.BindPFlag("capability", cmd.Flags().Lookup("capability"))
	viper.BindPFlag("modifiers", cmd.Flags().Lookup("modifiers"))

	return cmd
}

func (c *transformCmd) parse() error {
	for _, assetSubdir := range c.targetAssets {
		subdirPath := filepath.Join(c.samplesIn, "resource_pack", "textures", assetSubdir)

		err := afero.Walk(c.appfs, subdirPath, func(path string, d fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			rel, err := filepath.Rel(c.capability, path)

			if err != nil {
				return nil
			}

			out := filepath.Join(c.texturesOut, rel)

			params := []string{path}
			params = append(params, c.modifiers...)

			conf := transformConfig{
				in:     path,
				out:    out,
				params: params,
			}

			c.transforms = append(c.transforms, conf)

			return nil
		})

		return err
	}
	return nil
}

func (c *transformConfig) replaceEnvInModifiers() error {
	replacementMap := c.replacementEnvs()

	for i, modifier := range c.params {
		for _, rp := range replacementMap {
			modifier = strings.ReplaceAll(modifier, modifier, rp)
			c.params[i] = modifier
		}
	}

	return nil
}

func (c *transformConfig) replacementEnvs() map[string]string {
	fn := strings.TrimSuffix(filepath.Base(c.out), filepath.Ext(c.out))
	fe := filepath.Ext(c.out)

	return map[string]string{
		"%IN":       c.in,
		"%OUT":      c.out,
		"%FILENAME": fn,
		"%FILEEXT":  fe,
	}
}

func (c *transformCmd) exec() error {
	for _, transform := range c.transforms {
		err := os.MkdirAll(filepath.Dir(transform.out), os.ModePerm)

		if err != nil {
			return err
		}
	}
	return RunCmd(exec.Command(
		ImCmd,
		c.modifiers...,
	))
}
