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
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type UI struct {
	activeTabView *vo.TabInfo
	defaults      embed.FS
	app           fyne.App
	window        fyne.Window
}

func (ui *UI) Build(templates, defaults embed.FS) {
	ui.app = app.New()
	ui.window = ui.app.NewWindow("OpenPBR Config Creator")

	ui.defaults = defaults
	utils.Templates = templates

	store.Tabs = []*vo.TabInfo{
		{
			TabName: "Config",
			View: &Create{
				parent: ui.window,
			},
			TemplateSettings: &vo.TemplateSettings{
				DefaultData: "defaults/config.json",
			},
		},
		{
			TabName: "PBR",
			View:    &PBR{},
			TemplateSettings: &vo.TemplateSettings{
				TemplatePath: "pbr_global.tmpl",
				Output:       "pbr/pbr/global.json",
				DefaultData:  "defaults/pbr_global.json",
			},
		},
		{
			TabName: "Atmospheric",
			View:    &Atmospherics{},
			TemplateSettings: &vo.TemplateSettings{
				TemplatePath: "atmospherics.tmpl",
				Output:       "shared/atmospherics/atmospherics.json",
				DefaultData:  "defaults/atmospherics.json",
			},
		},
		{
			TabName: "Fog",
			View:    &Fog{},
			TemplateSettings: &vo.TemplateSettings{
				TemplatePath: "default_fog_settings.tmpl",
				Output:       "shared/fogs/default_fog_settings.json",
				DefaultData:  "defaults/default_fog_settings.json",
			},
		},
		{
			TabName: "Lighting",
			View:    &Lighting{},
			TemplateSettings: &vo.TemplateSettings{
				TemplatePath: "lighting_global.tmpl",
				Output:       "shared/lighting/global.json",
				DefaultData:  "defaults/lighting_global.json",
			},
		},
		{
			TabName: "Color Grading",
			View:    &ColorGrading{},
			TemplateSettings: &vo.TemplateSettings{
				TemplatePath: "color_grading.tmpl",
				Output:       "shared/color_grading/color_grading.json",
				DefaultData:  "defaults/color_grading.json",
			}},
		{
			TabName: "Water",
			View:    &Water{},
			TemplateSettings: &vo.TemplateSettings{
				TemplatePath: "water.tmpl",
				Output:       "shared/water/water.json",
				DefaultData:  "defaults/water.json",
			},
		},
		{TabName: "Build Package", View: &Pack{
			templates: templates,
		}, TemplateSettings: nil},
	}

	ui.activeTabView = store.Tabs[0]

	tabBar := container.NewAppTabs()
	for _, t := range store.Tabs {
		tabBar.Append(container.NewTabItem(t.TabName, t.View.Build(ui.window)))
	}

	tabBar.OnSelected = func(ti *container.TabItem) {
		for _, t := range store.Tabs {
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
				f, err := fs.ReadFile(ui.defaults, ui.activeTabView.DefaultData)

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
