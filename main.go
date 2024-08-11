package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/google/uuid"
)

//go:embed templates/*.tmpl
var templates embed.FS

func main() {

	a := app.New()
	w := a.NewWindow("OpenPBR Config Creator")

	manifestName := widget.NewEntry()
	item1 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Name"), manifestName)

	authorEntry := widget.NewEntry()
	authorContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Author Name"), authorEntry)

	licenseURL := widget.NewEntry()
	licenseContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("License URL"), licenseURL)

	packageURL := widget.NewEntry()
	packageContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Package URL"), packageURL)

	capibility := widget.NewSelect([]string{"PBR", "RTX"}, func(value string) {
		log.Println("Select set to", value)
	})
	capibilityContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Capibility"), capibility)

	manifestDescription := widget.NewEntry()
	item2 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Description"), manifestDescription)

	manifestHeaderUUID := widget.NewEntry()
	uuidBtn1 := widget.NewButton("<", func() {
		manifestHeaderUUID.Text = uuid.New().String()
	})

	group1 := container.New(layout.NewAdaptiveGridLayout(2), manifestHeaderUUID, uuidBtn1)
	uuidBtn1.Resize(fyne.NewSize(25, 25))
	item3 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Header Guid"), group1)

	manifestModuleUUID := widget.NewEntry()
	uuidBtn2 := widget.NewButton("<", func() {
		manifestModuleUUID.Text = uuid.New().String()
	})
	uuidBtn2.Resize(fyne.NewSize(25, 25))
	group2 := container.New(layout.NewAdaptiveGridLayout(2), manifestModuleUUID, uuidBtn2)
	item4 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Module Guid"), group2)

	manifestVersion := widget.NewEntry()
	manifestVersion.SetPlaceHolder("ex: [1, 0, 5]")
	item5 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Version"), manifestVersion)

	heightTemplateEntry := widget.NewEntry()
	heightTemplateEntry.SetPlaceHolder("ex: _height")
	heightContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Height Template"), heightTemplateEntry)

	normalTemplateEntry := widget.NewEntry()
	normalTemplateEntry.SetPlaceHolder("ex: _normal")
	normalContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Normal Template"), normalTemplateEntry)

	merTemplateEntry := widget.NewEntry()
	merTemplateEntry.SetPlaceHolder("ex: _mer")
	merContainer := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("MER Template"), merTemplateEntry)

	combo := widget.NewSelect([]string{"1.16.100", "1.21.30"}, func(value string) {
		log.Println("Select set to", value)
	})
	item6 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Texture Set Version"), combo)

	defaultMERArr := widget.NewEntry()
	defaultMERArr.SetPlaceHolder("ex: [255, 0, 255, 200]")
	item8 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Default MER Array"), defaultMERArr)

	v := container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Manifest"),
		item1,
		authorContainer,
		licenseContainer,
		packageContainer,
		capibilityContainer,
		item2,
		item3,
		item4,
		item5,
		widget.NewLabel("PBR Settings"),
		item6,
		item8,
		heightContainer,
		normalContainer,
		merContainer,
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

				cmd.CreateManifest([]string{
					saveFile,
					manifestName.Text,
					manifestDescription.Text,
					manifestHeaderUUID.Text,
					manifestModuleUUID.Text,
					combo.Selected,
					defaultMERArr.Text,
					manifestVersion.Text,
					authorEntry.Text,
					licenseURL.Text,
					packageURL.Text,
					capibility.Selected,
					heightTemplateEntry.Text,
					normalTemplateEntry.Text,
					merTemplateEntry.Text,
				})
			}, w)
		}))

	tabs := container.NewAppTabs(
		container.NewTabItem("Create Config", v),
		container.NewTabItem("Build Package", widget.NewButton("Load Config", func() {
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

				cmd.Build([]string{
					saveFile,
				})
			}, w)
		})),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(800, 600))

	w.Show()

	a.Run()

	cmd.Execute()
}
