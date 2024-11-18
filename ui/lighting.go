package ui

import (
	"encoding/json"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

var id int = 0

type Lighting struct {
	sunIlluminanceVBox      *vo.EntryView
	sunColourVBox           *vo.EntryView
	moonIlluminanceVBox     *vo.EntryView
	moonColourEntry         *widget.Entry
	orbitalOffsetEntry      *widget.Entry
	desaturationEntry       *widget.Entry
	ambientIlluminanceEntry *widget.Entry
	ambientColourEntry      *widget.Entry
}

func (v *Lighting) Build(p fyne.Window) *fyne.Container {

	//
	// Sun Illuminance
	//

	v.sunIlluminanceVBox = utils.CreateEntryView("Sun Illuminance", id)

	//
	// Sun Colour
	//

	v.sunColourVBox = utils.CreateEntryView("Sun Colour", id)

	//
	// Moon Illuminance
	//

	v.moonIlluminanceVBox = utils.CreateEntryView("Moon Illuminance", id)

	//
	// Moon Colour
	//

	moonColourLabel := widget.NewLabel("Moon Colour")
	v.moonColourEntry = widget.NewEntry()
	moonHBox := container.NewHBox(moonColourLabel, v.moonColourEntry)

	//
	// Orbital Offset
	//

	orbitalOffsetLabel := widget.NewLabel("Moon Colour")
	v.orbitalOffsetEntry = widget.NewEntry()
	orbitalOffsetHBox := container.NewHBox(orbitalOffsetLabel, v.orbitalOffsetEntry)

	//
	// Desaturation
	//

	desaturationLabel := widget.NewLabel("Desaturation")
	v.desaturationEntry = widget.NewEntry()
	desaturationHBox := container.NewHBox(desaturationLabel, v.desaturationEntry)

	//
	// Ambient Illuminance
	//

	ambientIlluminanceLabel := widget.NewLabel("Ambient Illuminance")
	v.ambientIlluminanceEntry = widget.NewEntry()
	ambientIlluminanceHBox := container.NewHBox(ambientIlluminanceLabel, v.ambientIlluminanceEntry)

	//
	// Ambient Colour
	//

	ambientColourLabel := widget.NewLabel("Ambient Colour")
	v.ambientColourEntry = widget.NewEntry()
	ambientColourHBox := container.NewHBox(ambientColourLabel, v.ambientColourEntry)

	accItem1 := widget.NewAccordionItem("Sun Illuminance", v.sunIlluminanceVBox.C)
	accItem2 := widget.NewAccordionItem("Sun Colour", v.sunColourVBox.C)
	accItem3 := widget.NewAccordionItem("Moon Illuminance", v.moonIlluminanceVBox.C)
	accItem4 := widget.NewAccordionItem("Moon Colour", moonHBox)
	accItem5 := widget.NewAccordionItem("Orbital Offset", orbitalOffsetHBox)
	accItem6 := widget.NewAccordionItem("Desaturation", desaturationHBox)
	accItem7 := widget.NewAccordionItem("Ambient Illuminance", ambientIlluminanceHBox)
	accItem8 := widget.NewAccordionItem("Ambient Colour", ambientColourHBox)

	acc := widget.NewAccordion(
		accItem1,
		accItem2,
		accItem3,
		accItem4,
		accItem5,
		accItem6,
		accItem7,
		accItem8,
	)

	c := container.NewVBox(acc)
	return c
}

func (v *Lighting) Defaults(b []byte) {
	var d vo.Lighting
	json.Unmarshal(b, &d)

	utils.PopulateKeysWithFloat(d.SunIlluminance, v.sunIlluminanceVBox)
	utils.PopulateKeysWithFloat(d.SunColour, v.sunColourVBox)
	utils.PopulateKeysWithFloat(d.MoonIlluminance, v.moonIlluminanceVBox)
	v.moonColourEntry.SetText(d.MoonColour)
	v.orbitalOffsetEntry.SetText(utils.FloatToString(d.OrbitalOffset))
	v.desaturationEntry.SetText(utils.FloatToString(d.Desaturation))
	v.ambientIlluminanceEntry.SetText(utils.FloatToString(d.AmbientIlluminance))
	v.ambientColourEntry.SetText(d.AmbientColour)
}

func (v *Lighting) Save() {
	cmd := export.Lighting{
		Lighting: vo.Lighting{
			BaseConf: vo.BaseConf{
				Out: path.Join(store.PackageStore, "global"),
			},
			SunIlluminance:     utils.StepsToVO(v.sunIlluminanceVBox.Steps),
			SunColour:          utils.StepsToVO(v.sunColourVBox.Steps),
			MoonIlluminance:    utils.StepsToVO(v.moonIlluminanceVBox.Steps),
			MoonColour:         v.moonColourEntry.Text,
			OrbitalOffset:      utils.ToFloat64(v.orbitalOffsetEntry),
			Desaturation:       utils.ToFloat64(v.desaturationEntry),
			AmbientIlluminance: utils.ToFloat64(v.ambientIlluminanceEntry),
			AmbientColour:      v.ambientColourEntry.Text,
		},
	}

	cmd.Save()
}
