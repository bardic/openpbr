package main

import (
	_ "embed"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/bardic/openpbr/cmd"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()

	cmd.Execute()
}
