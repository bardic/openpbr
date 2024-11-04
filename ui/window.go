package ui

import (
	"embed"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/utils"
)

type UI struct {
	templates  embed.FS
	app        fyne.App
	window     fyne.Window
	createView *Create
	packView   *Pack
}

func (ui *UI) Build(templates embed.FS) {
	ui.templates = templates

	ui.app = app.New()
	ui.window = ui.app.NewWindow("OpenPBR Config Creator")

	ui.createView = &Create{}
	ui.packView = &Pack{
		templates: templates,
		window:    ui.window,
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Create Config",
			ui.createView.BuildCreateView(
				ui.window.Canvas().Content().Refresh,
				func(config *cmd.Config, err error) {
					dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {
						if err != nil {
							dialog.ShowError(err, ui.window)
							return
						}

						if f == nil {
							return
						}

						err = config.Perform()

						if err != nil {
							dialog.ShowError(err, ui.window)
						}

					}, ui.window)

				},
				func(err error) {
					dialog.ShowError(err, ui.window)
				},
			)),
		container.NewTabItem("Build Package", ui.packView.BuildPackageView(
			ui.window.Canvas().Content().Refresh,
			func(manfiest *cmd.Manifest, err error) {
				dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, ui.window)
						return
					}

					if f == nil {
						return
					}

					err = manfiest.Perform()

					if err != nil {
						dialog.ShowError(err, ui.window)
					}

				}, ui.window)

			},
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	ui.window.SetContent(tabs)
	ui.window.Resize(fyne.NewSize(800, 600))

	ui.window.SetMainMenu(fyne.NewMainMenu(&fyne.Menu{
		Label: "Actions",
		Items: []*fyne.MenuItem{
			fyne.NewMenuItem(
				"Load Config",
				func() {
					dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {

						if err != nil {
							dialog.ShowError(err, ui.window)
							return
						}
						if f == nil {
							return
						}
						var saveFile = f.URI().Path()

						jsonFile, err := os.Open(saveFile)
						if err != nil {
							utils.AppendLoadOut("Fatal error: config.json missing")
							return
						}

						defer jsonFile.Close()

						byteValue, err := io.ReadAll(jsonFile)

						if err != nil {
							utils.AppendLoadOut("Fatal error: failed to read config.json")
							return
						}

						utils.Basedir = filepath.Dir(f.URI().Path())

						var jsonConfig cmd.Target
						err = json.Unmarshal(byteValue, &jsonConfig)

						if err != nil {
							utils.AppendLoadOut("Fatal error: failed to parse config.json")
							return
						}

						ui.createView.Update(jsonConfig)

						ui.window.Canvas().Content().Refresh()

					}, ui.window)

				}),
		},
	}))

	ui.window.Show()

	ui.app.Run()
}
