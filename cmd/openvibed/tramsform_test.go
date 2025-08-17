package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var transformTests = []struct {
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
			{param: "--target-assets", arg: "blocks"},
			{param: "--in", arg: "./test_data/sample"},
			//{param: "--out", arg: "/tmp"},
			{param: "--capability", arg: "pbr"},
			{param: "--modifiers", arg: "test.png test.jpg"},
		},
		expected: "test.jpg",
	},
}

func TestTransforms(t *testing.T) {
	actual := new(bytes.Buffer)
	cmdArgs := []string{"openvibe", "transform"}

	Cmd.Root().SetOut(actual)
	Cmd.Root().SetErr(actual)

	for _, test := range transformTests {
		for _, args := range test.args {
			cmdArgs = append(cmdArgs, args.param, args.arg)
		}

		tmpOut, err := os.MkdirTemp("", "*")

		if err != nil {
			fmt.Println(err)
		}

		cmdArgs = append(cmdArgs, "--out", tmpOut)
		Cmd.Root().SetArgs(cmdArgs)
		Cmd.Root().Execute()

		c := filepath.Join(tmpOut, test.expected)

		if err != nil {
			fmt.Println(err)
		}

		assert.Equal(t, actual.String(), c, "actual is not expected")
	}
}
