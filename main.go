package main

import (
	"embed"
	"github.com/bardic/openpbr/ui"
)

//go:embed templates/*.tmpl
var templates embed.FS

func main() {
	ui := ui.UI{}
	ui.Build(templates)
}
