package main

// import (
// 	"bytes"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var cleanTests = []struct {
// 	in  string
// 	out string
// }{
// 	{"%a", "[%a]"},
// 	{"%-a", "[%-a]"},
// 	{"%+a", "[%+a]"},
// 	{"%#a", "[%#a]"},
// 	{"% a", "[% a]"},
// 	{"%0a", "[%0a]"},
// 	{"%1.2a", "[%1.2a]"},
// 	{"%-1.2a", "[%-1.2a]"},
// 	{"%+1.2a", "[%+1.2a]"},
// 	{"%-+1.2a", "[%+-1.2a]"},
// 	{"%-+1.2abc", "[%+-1.2a]bc"},
// 	{"%-1.2abc", "[%-1.2a]bc"},
// }

// func TestHelloName(t *testing.T) {

// 	actual := new(bytes.Buffer)
// 	Cmd.Root().SetOut(actual)
// 	Cmd.Root().SetErr(actual)
// 	Cmd.Root().SetArgs([]string{"A", "a1"})
// 	Cmd.Root().Execute()

// 	expected := "This-is-command-a1"

// 	assert.Equal(t, actual.String(), expected, "actual is not expected")
// }
