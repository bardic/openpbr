package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cleanTests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
	// {"%+a", "[%+a]"},
	// {"%#a", "[%#a]"},
	// {"% a", "[% a]"},
	// {"%0a", "[%0a]"},
	// {"%1.2a", "[%1.2a]"},
	// {"%-1.2a", "[%-1.2a]"},
	// {"%+1.2a", "[%+1.2a]"},
	// {"%-+1.2a", "[%+-1.2a]"},
	// {"%-+1.2abc", "[%+-1.2a]bc"},
	// {"%-1.2abc", "[%-1.2a]bc"},
}

func TestDownloand(t *testing.T) {
	baseDir := "./base"
	downloadDir := "./downloads"

	actual := new(bytes.Buffer)
	Cmd.Root().SetOut(actual)
	Cmd.Root().SetErr(actual)
	Cmd.Root().SetArgs([]string{"download", "--url", "https://api.github.com/repos/Mojang/bedrock-samples/releases/latest"})
	Cmd.Root().Execute()

	expected := ""

	assert.Equal(t, actual.String(), expected, "actual is not expected")

	_, err := os.Stat(filepath.Join(downloadDir, "latest.zip"))
	assert.False(t, os.IsNotExist(err), "File should exist")

	_, err = os.Stat(filepath.Join(baseDir, "Mojang-bedrock-samples-*"))
	assert.False(t, os.IsNotExist(err), "File should exist")

	err = os.RemoveAll(downloadDir)
	assert.ErrorIs(t, err, nil)

	err = os.RemoveAll(baseDir)
	assert.ErrorIs(t, err, nil)
}
