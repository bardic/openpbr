package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configTests = []struct {
	args []struct {
		param string
		arg   string
	}
	expected string
}{
	{
		args: []struct {
			param string
			arg   string
		}{
			{param: "--out", arg: filepath.Join(os.TempDir(), "openvibe_test_config.json")},
			{param: "--build-name", arg: "open vibe"},
			{param: "--name", arg: "open vibe"},
			{param: "--description", arg: "vanilla pbr/vibrant texture pack"},
			{param: "--header-uuid", arg: "bd111387-103a-4ef4-b33b-b39af1270e50"},
			{param: "--module-uuid", arg: "5e867951-86e8-4c74-960c-1e7d1b8beedb"},
			{param: "--default-mer", arg: "[0,0,125,95]"},
			{param: "--version", arg: "[0,0,1]"},
			{param: "--author", arg: "judohippo"},
			{param: "--license", arg: "https://github.com/bardic/openpbr/blob/main/UNLICENSE"},
			{param: "--url", arg: "https://github.com/bardic/openpbr"},
			{param: "--capability", arg: "pbr"},
			{param: "--height-template", arg: "_height"},
			{param: "--mer-template", arg: "_mer"},
			{param: "--r-offset", arg: "1"},
			{param: "--g-offset", arg: "1"},
			{param: "--b-offset", arg: "1"},
		},
		expected: filepath.Join("./test_data", "config.json"),
	},
}

func TestCreateConfig(t *testing.T) {
	actual := new(bytes.Buffer)
	cmdArgs := []string{"openvibe", "config"}

	Cmd.Root().SetOut(actual)
	Cmd.Root().SetErr(actual)

	for _, test := range configTests {
		for _, args := range test.args {
			cmdArgs = append(cmdArgs, args.param, args.arg)
		}

		Cmd.Root().SetArgs(cmdArgs)
		Cmd.Root().Execute()

		c, err := PathToJson[string](test.expected)

		if err != nil {
			fmt.Println(err)
		}

		assert.Equal(t, actual.String(), c, "actual is not expected")
	}
}
