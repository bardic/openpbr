package main

import (
	"embed"

	"github.com/bardic/openpbr/cmd"
)

//go:embed templates/*.tmpl
var templates embed.FS

func main() {
	cmd.Execute(templates)
}
