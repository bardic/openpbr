package ui

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
	"github.com/google/uuid"
)

type Create struct {
	parent                       fyne.Window
	manifestName                 *widget.Entry
	manifestNameContainer        *fyne.Container
	authorEntry                  *widget.Entry
	authorEntryContainer         *fyne.Container
	licenseURL                   *widget.Entry
	licenseURLContainer          *fyne.Container
	packageURL                   *widget.Entry
	packageURLContainer          *fyne.Container
	capibility                   *widget.Select
	capibilityContainer          *fyne.Container
	manifestDescription          *widget.Entry
	manifestDescriptionContainer *fyne.Container
	manifestHeaderUUID           *widget.Entry
	manifestHeaderUUIDBtn        *widget.Button
	manifestHeaderUUIDGroup      *fyne.Container
	manifestHeaderUUIDContainer  *fyne.Container
	manifestModuleUUID           *widget.Entry
	manifestModuleUUIDBtn        *widget.Button
	manifestModuleUUIDGroup      *fyne.Container
	manifestModuleUUIDContainer  *fyne.Container
	manifestVersion              *widget.Entry
	manifestVersionContainer     *fyne.Container
	heightTemplateEntry          *widget.Entry
	heightTemplateEntryContainer *fyne.Container
	merTemplateEntry             *widget.Entry
	merTemplateEntryContainer    *fyne.Container
	defaultMERArrEntry           *widget.Entry
	defaultMERArrEntryContainer  *fyne.Container
	manifestSectionHeader        *widget.Label
	pbrSectionHeader             *widget.Label
	rbgInfoHeader                *widget.Label
	rgbContainer                 *fyne.Container
	rText                        *widget.Entry
	gText                        *widget.Entry
	bText                        *widget.Entry
}

func (c *Create) Build(p fyne.Window) *fyne.Container {
	c.manifestName = widget.NewEntry()
	c.manifestNameContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Name"), c.manifestName)

	c.authorEntry = widget.NewEntry()
	c.authorEntryContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Author Name"), c.authorEntry)

	c.licenseURL = widget.NewEntry()
	c.licenseURLContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("License URL"), c.licenseURL)

	c.packageURL = widget.NewEntry()
	c.packageURLContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Package URL"), c.packageURL)

	c.capibility = widget.NewSelect([]string{"pbr", "plain"}, func(value string) {
		log.Println("Select set to", value)
	})
	c.capibilityContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Capibility"), c.capibility)

	c.manifestDescription = widget.NewEntry()
	c.manifestDescriptionContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Description"), c.manifestDescription)

	c.manifestHeaderUUID = widget.NewEntry()
	c.manifestHeaderUUIDBtn = widget.NewButton("<", func() {
		c.manifestHeaderUUID.Text = uuid.New().String()
		p.Canvas().Content().Refresh()
	})

	c.manifestHeaderUUIDGroup = container.New(layout.NewAdaptiveGridLayout(2), c.manifestHeaderUUID, c.manifestHeaderUUIDBtn)
	c.manifestHeaderUUIDBtn.Resize(fyne.NewSize(25, 25))
	c.manifestHeaderUUIDContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Header Guid"), c.manifestHeaderUUIDGroup)

	c.manifestModuleUUID = widget.NewEntry()
	c.manifestModuleUUIDBtn = widget.NewButton("<", func() {
		c.manifestModuleUUID.Text = uuid.New().String()
		p.Canvas().Content().Refresh()
	})
	c.manifestModuleUUIDBtn.Resize(fyne.NewSize(25, 25))
	c.manifestModuleUUIDGroup = container.New(layout.NewAdaptiveGridLayout(2), c.manifestModuleUUID, c.manifestModuleUUIDBtn)
	c.manifestModuleUUIDContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Module Guid"), c.manifestModuleUUIDGroup)

	c.rText = widget.NewEntry()
	c.rText.SetText("0")
	c.gText = widget.NewEntry()
	c.gText.SetText("0")
	c.bText = widget.NewEntry()
	c.bText.SetText("0")

	c.rbgInfoHeader = widget.NewLabel("Reasonable RGB offset values are between -15 and 15")
	c.rgbContainer = container.New(layout.NewAdaptiveGridLayout(3), c.rText, c.gText, c.bText)

	c.manifestVersion = widget.NewEntry()
	c.manifestVersion.SetPlaceHolder("ex: [1, 0, 5]")
	c.manifestVersionContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Version"), c.manifestVersion)

	c.heightTemplateEntry = widget.NewEntry()
	c.heightTemplateEntry.SetText("_height")
	c.heightTemplateEntryContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Height Template"), c.heightTemplateEntry)

	c.merTemplateEntry = widget.NewEntry()
	c.merTemplateEntry.SetText("_mer")
	c.merTemplateEntryContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("MER Template"), c.merTemplateEntry)

	c.defaultMERArrEntry = widget.NewEntry()
	c.defaultMERArrEntry.SetPlaceHolder("ex: [255, 0, 255, 200]")
	c.defaultMERArrEntryContainer = container.New(layout.NewAdaptiveGridLayout(2), widget.NewLabel("Default MER Array"), c.defaultMERArrEntry)

	c.manifestSectionHeader = widget.NewLabel("Manifest")
	c.manifestSectionHeader.TextStyle.Bold = true
	c.manifestSectionHeader.TextStyle.Underline = true

	c.pbrSectionHeader = widget.NewLabel("PBR Settings")
	c.pbrSectionHeader.TextStyle.Bold = true
	c.pbrSectionHeader.TextStyle.Underline = true

	v := container.New(
		layout.NewVBoxLayout(),
		c.manifestSectionHeader,
		c.manifestNameContainer,
		c.authorEntryContainer,
		c.licenseURLContainer,
		c.packageURLContainer,
		c.capibilityContainer,
		c.manifestDescriptionContainer,
		c.manifestHeaderUUIDContainer,
		c.manifestModuleUUIDContainer,
		c.manifestVersionContainer,
		c.pbrSectionHeader,
		c.rbgInfoHeader,
		c.rgbContainer,
		c.defaultMERArrEntryContainer,
		c.heightTemplateEntryContainer,
		c.merTemplateEntryContainer,
	)

	return v
}

func (c *Create) Defaults(b []byte) {
	if b == nil {
		c.manifestName.SetText("")
		c.manifestDescription.SetText("")
		c.manifestHeaderUUID.SetText("")
		c.manifestModuleUUID.SetText("")
		c.defaultMERArrEntry.SetText("")
		c.manifestVersion.SetText("")
		c.authorEntry.SetText("")
		c.licenseURL.SetText("")
		c.packageURL.SetText("")
		c.capibility.SetSelectedIndex(0)
		c.heightTemplateEntry.SetText("")
		c.merTemplateEntry.SetText("")
		c.rText.SetText("")
		c.gText.SetText("")
		c.bText.SetText("")

		return
	}

	var vo cmd.Config
	json.Unmarshal(b, &vo)

	c.manifestName.SetText(vo.Name)
	c.manifestDescription.SetText(vo.Description)
	c.manifestHeaderUUID.SetText(vo.Header_uuid)
	c.manifestModuleUUID.SetText(vo.Module_uuid)
	c.defaultMERArrEntry.SetText(vo.Default_mer)
	c.manifestVersion.SetText(vo.Version)
	c.authorEntry.SetText(vo.Author)
	c.licenseURL.SetText(vo.License)
	c.packageURL.SetText(vo.URL)
	if vo.Capibility == "pbr" {
		c.capibility.SetSelectedIndex(0)
	} else {
		c.capibility.SetSelectedIndex(1)
	}

	c.heightTemplateEntry.SetText(vo.HeightTemplate)
	c.merTemplateEntry.SetText(vo.MerTemplate)
	c.rText.SetText(vo.ROffset)
	c.gText.SetText(vo.GOffset)
	c.bText.SetText(vo.BOffset)
}

func (c *Create) Update(t cmd.Config) {

	c.manifestName.Text = t.Name
	c.manifestDescription.Text = t.Description
	c.manifestHeaderUUID.Text = t.Header_uuid
	c.manifestModuleUUID.Text = t.Module_uuid
	c.defaultMERArrEntry.Text = t.Default_mer
	c.manifestVersion.Text = t.Version
	c.authorEntry.Text = t.Author
	c.licenseURL.Text = t.License
	c.packageURL.Text = t.URL
	if t.Capibility == "pbr" {
		c.capibility.SetSelectedIndex(0)
	} else {
		c.capibility.SetSelectedIndex(1)
	}
	c.heightTemplateEntry.Text = t.HeightTemplate
	c.merTemplateEntry.Text = t.MerTemplate
	c.rText.Text = t.ROffset
	c.gText.Text = t.GOffset
	c.bText.Text = t.BOffset
}

func (c *Create) Save() {
	dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, c.parent)
			return
		}

		if f == nil {
			return
		}

		store.PackageStore = filepath.Dir(f.URI().Path())
		store.Output = path.Join(store.PackageStore, "export")

		os.MkdirAll(path.Join(store.PackageStore, utils.Psds), os.ModePerm)
		os.MkdirAll(path.Join(store.PackageStore, utils.Overrides), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "atmospherics"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "color_grading"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "fogs"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "lighting"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "point_lights"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "textures"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "shared", "water"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, utils.SettingDir, "pbr", "pbr"), os.ModePerm)
		os.MkdirAll(path.Join(store.Output, "overrides"), os.ModePerm)

		os.MkdirAll(store.Output, os.ModePerm)

		config := &cmd.Config{
			BaseConf: vo.BaseConf{
				Out: f.URI().Path(),
			},
			Name:           c.manifestName.Text,
			Header_uuid:    c.manifestHeaderUUID.Text,
			Module_uuid:    c.manifestModuleUUID.Text,
			Description:    c.manifestDescription.Text,
			Default_mer:    c.defaultMERArrEntry.Text,
			Version:        c.manifestVersion.Text,
			Author:         c.authorEntry.Text,
			License:        c.licenseURL.Text,
			URL:            c.packageURL.Text,
			Capibility:     c.capibility.Selected,
			HeightTemplate: c.heightTemplateEntry.Text,
			MerTemplate:    c.merTemplateEntry.Text,
			ROffset:        c.rText.Text,
			GOffset:        c.gText.Text,
			BOffset:        c.bText.Text,
		}

		config.Perform()
	}, c.parent)

}
