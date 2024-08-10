package main

import (
	_ "embed"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/bardic/openpbr/cmd"
	"github.com/google/uuid"
)

func main() {

	a := app.New()
	w := a.NewWindow("OpenPBR Config Creator")

	manifestName := widget.NewEntry()
	item1 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Name"), manifestName)

	homeDirPath := widget.NewEntry()
	homeBrowseBtn := widget.NewButton("^", func() {
		dialog.ShowFolderOpen(func(f fyne.ListableURI, err error) {

			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if f == nil {
				return
			}

			l, err := f.List()
			var saveFile = l[0].Path()
			homeDirPath.Text = saveFile
		}, w)
	})

	homeBrowseBtn.Resize(fyne.NewSize(25, 25))
	homeGroup := container.New(layout.NewAdaptiveGridLayout(2), homeDirPath, homeBrowseBtn)
	homeLayout := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Home Dir"), homeGroup)

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

	combo := widget.NewSelect([]string{"1.16.100", "1.21.30"}, func(value string) {
		log.Println("Select set to", value)
	})
	item6 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Texture Set Version"), combo)

	textureSetVersion := widget.NewEntry()
	item7 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Texture Set Version"), textureSetVersion)

	defaultMERArr := widget.NewEntry()
	defaultMERArr.SetPlaceHolder("ex: [255, 0, 255, 200]")
	item8 := container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Default MER Array"), defaultMERArr)

	v := container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Manifest"),
		item1,
		homeLayout,
		item2,
		item3,
		item4,
		item5,
		widget.NewLabel("PBR Settings"),
		item6,
		item7,
		item8,
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
					textureSetVersion.Text,
					defaultMERArr.Text,
					manifestVersion.Text,
				})
			}, w)
		}))

	tabs := container.NewAppTabs(
		container.NewTabItem("Create Config", v),
		container.NewTabItem("Build Package", widget.NewButton("Load Config", func() {
			dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
				// saveFile := "NoFileYet"
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if f == nil {
					return
				}
				var saveFile = f.URI().Path()
				fmt.Println(saveFile)
				cmd.Build([]string{
					saveFile,
					manifestName.Text,
					manifestDescription.Text,
					manifestHeaderUUID.Text,
					manifestModuleUUID.Text,
					manifestVersion.Text,
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
