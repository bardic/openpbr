package cmd

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/bardic/openpbr/cmd/data"
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/google/uuid"
)

func UI(templates embed.FS) {

	a := app.New()
	w := a.NewWindow("OpenPBR Config Creator")

	CheckEnv()

	manifestName := widget.NewEntry()
	manifestNameContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Name"), manifestName)

	authorEntry := widget.NewEntry()
	authorEntryContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Author Name"), authorEntry)

	licenseURL := widget.NewEntry()
	licenseURLContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("License URL"), licenseURL)

	packageURL := widget.NewEntry()
	packageURLContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Package URL"), packageURL)

	capibility := widget.NewSelect([]string{"pbr", "rtx"}, func(value string) {
		log.Println("Select set to", value)
	})
	capibilityContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Capibility"), capibility)

	manifestDescription := widget.NewEntry()
	manifestDescriptionContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Description"), manifestDescription)

	manifestHeaderUUID := widget.NewEntry()
	manifestHeaderUUIDBtn := widget.NewButton("<", func() {
		manifestHeaderUUID.Text = uuid.New().String()
		w.Canvas().Content().Refresh()
	})

	manifestHeaderUUIDGroup := container.New(layout.NewAdaptiveGridLayout(2), manifestHeaderUUID, manifestHeaderUUIDBtn)
	manifestHeaderUUIDBtn.Resize(fyne.NewSize(25, 25))
	manifestHeaderUUIDContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Header Guid"), manifestHeaderUUIDGroup)

	manifestModuleUUID := widget.NewEntry()
	manifestModuleUUIDBtn := widget.NewButton("<", func() {
		manifestModuleUUID.Text = uuid.New().String()
		w.Canvas().Content().Refresh()
	})
	manifestModuleUUIDBtn.Resize(fyne.NewSize(25, 25))
	manifestModuleUUIDGroup := container.New(layout.NewAdaptiveGridLayout(2), manifestModuleUUID, manifestModuleUUIDBtn)
	manifestModuleUUIDContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Module Guid"), manifestModuleUUIDGroup)

	manifestVersion := widget.NewEntry()
	manifestVersion.SetPlaceHolder("ex: [1, 0, 5]")
	manifestVersionContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Version"), manifestVersion)

	heightTemplateEntry := widget.NewEntry()
	heightTemplateEntry.SetPlaceHolder("ex: _height")
	heightTemplateEntryContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Height Template"), heightTemplateEntry)

	normalTemplateEntry := widget.NewEntry()
	normalTemplateEntry.SetPlaceHolder("ex: _normal")
	normalTemplateEntryContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Normal Template"), normalTemplateEntry)

	merTemplateEntry := widget.NewEntry()
	merTemplateEntry.SetPlaceHolder("ex: _mer")
	merTemplateEntryContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("MER Template"), merTemplateEntry)

	texturesetSelector := widget.NewSelect([]string{"1.16.100", "1.21.30"}, func(value string) {
		log.Println("Select set to", value)
	})
	texturesetSelectorContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Texture Set Version"), texturesetSelector)

	defaultMERArrEntry := widget.NewEntry()
	defaultMERArrEntry.SetPlaceHolder("ex: [255, 0, 255, 200]")
	defaultMERArrEntryContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Default MER Array"), defaultMERArrEntry)

	exportMERCSVCheck := widget.NewCheck("Export MER Override CSV", func(b bool) {

	})
	exportMERCSVCheckContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel(""), exportMERCSVCheck)

	manifestSectionHeader := widget.NewLabel("Manifest")
	manifestSectionHeader.TextStyle.Bold = true
	manifestSectionHeader.TextStyle.Underline = true

	pbrSectionHeader := widget.NewLabel("PBR Settings")
	pbrSectionHeader.TextStyle.Bold = true
	pbrSectionHeader.TextStyle.Underline = true

	v := container.New(
		layout.NewVBoxLayout(),
		manifestSectionHeader,
		manifestNameContainer,
		authorEntryContainer,
		licenseURLContainer,
		packageURLContainer,
		capibilityContainer,
		manifestDescriptionContainer,
		manifestHeaderUUIDContainer,
		manifestModuleUUIDContainer,
		manifestVersionContainer,
		pbrSectionHeader,
		texturesetSelectorContainer,
		defaultMERArrEntryContainer,
		heightTemplateEntryContainer,
		normalTemplateEntryContainer,
		merTemplateEntryContainer,
		exportMERCSVCheckContainer,
		widget.NewButton("Save", func() {
			dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {

				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if f == nil {
					return
				}
				var saveFile = f.URI().Path()

				//cmd.CreateCSV(saveFile, defaultMERArrEntry.Text)
				CreateManifest([]string{
					saveFile,
					manifestName.Text,
					manifestDescription.Text,
					manifestHeaderUUID.Text,
					manifestModuleUUID.Text,
					texturesetSelector.Selected,
					defaultMERArrEntry.Text,
					manifestVersion.Text,
					authorEntry.Text,
					licenseURL.Text,
					packageURL.Text,
					capibility.Selected,
					heightTemplateEntry.Text,
					normalTemplateEntry.Text,
					merTemplateEntry.Text,
					strconv.FormatBool(exportMERCSVCheck.Checked),
				})
			}, w)
		}))

	loadConfigContainer := container.New(layout.NewVBoxLayout())
	loadBtn := widget.NewButton("Load Config", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if f == nil {
				return
			}
			var saveFile = f.URI().Path()
			utils.Basedir = filepath.Dir(saveFile)

			dir, err := templates.ReadDir("templates")

			if err != nil {
				fmt.Println(err)
				return
			}

			tempDir := utils.LocalPath("templates")
			os.MkdirAll(tempDir, os.ModePerm)
			os.MkdirAll(utils.LocalPath(utils.OutDir), os.ModePerm)
			os.MkdirAll(utils.LocalPath(utils.Psds), os.ModePerm)
			os.MkdirAll(utils.LocalPath(utils.Overrides), os.ModePerm)
			os.MkdirAll(utils.LocalPath(utils.SettingDir), os.ModePerm)

			for _, v := range dir {

				filePath := tempDir + string(os.PathSeparator) + v.Name()
				out, err := os.Create(filePath)

				if err != nil {
					fmt.Println(err)
					return
				}
				defer out.Close()

				b, err := templates.ReadFile("templates/" + v.Name())

				if err != nil {
					fmt.Println(err)
					return
				}

				_, err = io.WriteString(out, string(b))

				if err != nil {
					fmt.Println(err)
					return
				}
			}
			pb := widget.NewProgressBarInfinite()
			loadConfigContainer.Add(pb)
			utils.LoadStdOut = widget.NewTextGrid()
			loadConfigContainer.Add(utils.LoadStdOut)

			w.Canvas().Content().Refresh()
			err = Build([]string{
				saveFile,
			})
			pb.Stop()
			if err != nil {
				pb.Theme().Color(fyne.ThemeColorName("red"), fyne.ThemeVariant(1))
				return
			}

			pb.Theme().Color(fyne.ThemeColorName("green"), fyne.ThemeVariant(1))
			w.Canvas().Content().Refresh()

		}, w)
	})

	loadBtn.Resize(fyne.NewSize(25, 25))
	loadConfigContainer.Add(loadBtn)

	tabs := container.NewAppTabs(
		container.NewTabItem("Create Config", v),
		container.NewTabItem("Build Package", loadConfigContainer),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(800, 600))

	w.SetMainMenu(fyne.NewMainMenu(&fyne.Menu{
		Label: "Actions",
		Items: []*fyne.MenuItem{
			fyne.NewMenuItem(
				"Load Config",
				func() {
					dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {

						if err != nil {
							dialog.ShowError(err, w)
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

						var jsonConfig data.Targets
						json.Unmarshal(byteValue, &jsonConfig)

						if len(jsonConfig.Targets) == 0 {
							utils.AppendLoadOut("Fatal error: no targets configured in config")
							return
						}

						manifestName.Text = jsonConfig.Targets[0].Name
						manifestDescription.Text = jsonConfig.Targets[0].Description
						manifestHeaderUUID.Text = jsonConfig.Targets[0].Header_uuid
						manifestModuleUUID.Text = jsonConfig.Targets[0].Module_uuid
						if jsonConfig.Targets[0].Textureset_format == "1.16.100" {
							texturesetSelector.SetSelectedIndex(0)
						} else {
							texturesetSelector.SetSelectedIndex(1)
						}
						defaultMERArrEntry.Text = jsonConfig.Targets[0].Default_mer
						manifestVersion.Text = jsonConfig.Targets[0].Version
						authorEntry.Text = jsonConfig.Targets[0].Author
						licenseURL.Text = jsonConfig.Targets[0].License
						packageURL.Text = jsonConfig.Targets[0].URL
						if jsonConfig.Targets[0].Capibility == "pbr" {
							capibility.SetSelectedIndex(0)
						} else {
							capibility.SetSelectedIndex(1)
						}
						heightTemplateEntry.Text = jsonConfig.Targets[0].HeightTemplate
						normalTemplateEntry.Text = jsonConfig.Targets[0].NormalTemplate
						merTemplateEntry.Text = jsonConfig.Targets[0].MerTemplate

						exportMERCSVCheck.Checked = jsonConfig.Targets[0].ExportMer

						w.Canvas().Content().Refresh()

					}, w)

				}),
		},
	}))

	w.Show()

	a.Run()
}
