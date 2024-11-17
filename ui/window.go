package ui

import (
	"embed"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type UI struct {
	templates    embed.FS
	defaults     embed.FS
	app          fyne.App
	window       fyne.Window
	createView   *Create
	packView     *Pack
	lightingView *Lighting
	water        *Water
	atmospherics *Atmospherics
	fog          *Fog
	colorGrading *ColorGrading
	pbr          *PBR
}

func (ui *UI) Build(templates, defaults embed.FS) {
	ui.templates = templates
	ui.defaults = defaults

	ui.app = app.New()
	ui.window = ui.app.NewWindow("OpenPBR Config Creator")

	ui.createView = &Create{}
	ui.packView = &Pack{
		templates: templates,
		window:    ui.window,
	}
	ui.lightingView = &Lighting{}
	ui.water = &Water{}
	ui.atmospherics = &Atmospherics{}
	ui.fog = &Fog{}
	ui.colorGrading = &ColorGrading{}
	ui.pbr = &PBR{}

	tabs := container.NewAppTabs(
		container.NewTabItem("Config",
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

						config.Buildname = f.URI().Path()

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
		container.NewTabItem("PBR", ui.pbr.BuildLightingView(
			ui.window.Canvas().Content().Refresh,
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
		container.NewTabItem("Atmospheric", ui.atmospherics.Build(
			ui.window.Canvas().Content().Refresh,
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
		container.NewTabItem("Fog", ui.fog.BuildLightingView(
			ui.window.Canvas().Content().Refresh,
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
		container.NewTabItem("Lighting", ui.lightingView.BuildLightingView(
			ui.window.Canvas().Content().Refresh,
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
		container.NewTabItem("Color Grading", ui.colorGrading.BuildLightingView(
			ui.window.Canvas().Content().Refresh,
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
		container.NewTabItem("Water", ui.water.BuildLightingView(
			ui.window.Canvas().Content().Refresh,
			func(err error) {
				dialog.ShowError(err, ui.window)
			},
		)),
		container.NewTabItem("Build Package", ui.packView.BuildPackageView(
			ui.window.Canvas().Content().Refresh,
			func(manfiest *export.Manifest, err error) {
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

	tb := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.HistoryIcon(),
			func() {
				switch tabs.Selected().Text {
				case "Config":

				case "PBR":

				case "Atmospheric":
					f, err := fs.ReadFile(ui.defaults, "defaults/atmospherics.json")

					if err != nil {
						dialog.ShowError(err, ui.window)
						return
					}

					var vo *vo.Atmospherics
					json.Unmarshal(f, &vo)

					ui.atmospherics.Defaults(vo)
				case "Fog":
					f, err := fs.ReadFile(ui.defaults, "defaults/default_fog_settings.json")

					if err != nil {
						dialog.ShowError(err, ui.window)
						return
					}

					var vo *vo.Fog
					json.Unmarshal(f, &vo)

					ui.fog.Defaults(vo)
				case "Lighting":

				case "Color Grading":
					f, err := fs.ReadFile(ui.defaults, "defaults/color_grading.json")

					if err != nil {
						dialog.ShowError(err, ui.window)
						return
					}

					var vo *vo.ColorGrading
					json.Unmarshal(f, &vo)

					ui.colorGrading.Defaults(vo)
				case "Water":
					f, err := fs.ReadFile(ui.defaults, "defaults/water.json")

					if err != nil {
						dialog.ShowError(err, ui.window)
						return
					}

					var vo *vo.Water
					json.Unmarshal(f, &vo)

					ui.water.Defaults(vo)
				case "Build Package":

				}

				ui.window.Canvas().Content().Refresh()
			},
		),
	)

	ui.window.SetContent(container.NewVBox(tb, tabs))
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

						var jsonConfig cmd.Config
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

	err := utils.StartUpCheck()

	if err != nil {
		dialog.ShowError(errors.New("meow"), ui.window)
	}

	ui.app.Run()

}
