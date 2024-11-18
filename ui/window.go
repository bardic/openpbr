package ui

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type UI struct {
	activeTabView *TabInfo
	templates     embed.FS
	defaults      embed.FS
	app           fyne.App
	window        fyne.Window
}

type TabInfo struct {
	TabName string
	View    vo.IBaseView
	Default string
}

func (ui *UI) Build(templates, defaults embed.FS) {
	ui.app = app.New()
	ui.window = ui.app.NewWindow("OpenPBR Config Creator")

	tabs := []*TabInfo{
		{"Config", &Create{
			parent: ui.window,
		}, "defaults/config.json"},
		{"PBR", &PBR{}, "defaults/pbr_global.json"},
		{"Atmospheric", &Atmospherics{}, "defaults/atmospherics.json"},
		{"Fog", &Fog{}, "defaults/default_fog_settings.json"},
		{"Lighting", &Lighting{}, "defaults/lighting_global.json"},
		{"Color Grading", &ColorGrading{}, "defaults/color_grading.json"},
		{"Water", &Water{}, "defaults/water.json"},
		{"Build Package", &Pack{}, ""},
	}

	ui.activeTabView = tabs[0]

	ui.templates = templates
	ui.defaults = defaults

	tabBar := container.NewAppTabs()
	for _, t := range tabs {
		tabBar.Append(container.NewTabItem(t.TabName, t.View.Build(ui.window)))
	}

	tabBar.OnSelected = func(ti *container.TabItem) {
		for _, t := range tabs {
			if t.TabName == ti.Text {
				ui.activeTabView = t
				break
			}
		}
	}

	tabBar.SetTabLocation(container.TabLocationTop)

	tb := widget.NewToolbar(

		widget.NewToolbarAction(
			theme.DocumentIcon(),
			func() {
				ui.activeTabView.View.Defaults(nil)
				ui.window.Canvas().Content().Refresh()
			},
		),
		widget.NewToolbarAction(
			theme.FolderOpenIcon(),
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

					ui.activeTabView.View.Defaults(byteValue)

				}, ui.window)
			},
		),
		widget.NewToolbarAction(
			theme.DocumentSaveIcon(),
			func() {
				ui.activeTabView.View.Save()
			}),
		widget.NewToolbarAction(
			theme.HistoryIcon(),
			func() {
				f, err := fs.ReadFile(ui.defaults, ui.activeTabView.Default)

				if err != nil {
					dialog.ShowError(err, ui.window)
					return
				}
				ui.activeTabView.View.Defaults(f)

				ui.window.Canvas().Content().Refresh()
			},
		),
	)

	ui.window.SetContent(container.NewVBox(tb, tabBar))
	ui.window.Resize(fyne.NewSize(800, 600))

	ui.window.SetMainMenu(fyne.NewMainMenu(&fyne.Menu{
		Label: "Actions",
		Items: []*fyne.MenuItem{
			fyne.NewMenuItem(
				"About",
				func() {
					fmt.Println("About")
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
