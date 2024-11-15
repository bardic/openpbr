package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

var id int = 0

type Lighting struct {
}

func (v *Lighting) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {

	//
	// Sun Illuminance
	//

	sunIlluminanceVBox := utils.CreateEntryView("Sun Illuminance", id)

	//
	// Sun Colour
	//

	sunColourVBox := utils.CreateEntryView("Sun Colour", id)

	//
	// Moon Illuminance
	//

	moonIlluminanceVBox := utils.CreateEntryView("Moon Illuminance", id)

	//
	// Moon Colour
	//

	moonColourLabel := widget.NewLabel("Moon Colour")
	moonColourEntry := widget.NewEntry()
	moonHBox := container.NewHBox(moonColourLabel, moonColourEntry)

	//
	// Orbital Offset
	//

	orbitalOffsetLabel := widget.NewLabel("Moon Colour")
	orbitalOffsetEntry := widget.NewEntry()
	orbitalOffsetHBox := container.NewHBox(orbitalOffsetLabel, orbitalOffsetEntry)

	//
	// Desaturation
	//

	desaturationLabel := widget.NewLabel("Desaturation")
	desaturationEntry := widget.NewEntry()
	desaturationHBox := container.NewHBox(desaturationLabel, desaturationEntry)

	//
	// Ambient Illuminance
	//

	ambientIlluminanceLabel := widget.NewLabel("Ambient Illuminance")
	ambientIlluminanceEntry := widget.NewEntry()
	ambientIlluminanceHBox := container.NewHBox(ambientIlluminanceLabel, ambientIlluminanceEntry)

	//
	// Ambient Colour
	//

	ambientColourLabel := widget.NewLabel("Ambient Colour")
	ambientColourEntry := widget.NewEntry()
	ambientColourHBox := container.NewHBox(ambientColourLabel, ambientColourEntry)

	accItem1 := widget.NewAccordionItem("Sun Illuminance", sunIlluminanceVBox.C)
	accItem2 := widget.NewAccordionItem("Sun Colour", sunColourVBox.C)
	accItem3 := widget.NewAccordionItem("Moon Illuminance", moonIlluminanceVBox.C)
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

	save := widget.NewButton("Save", func() {
		cmd := export.Lighting{
			Out: "./example/settings/shared/lighting/global.json",
			Lighting: vo.Lighting{
				SunIlluminance:     utils.StepsToVO(sunIlluminanceVBox.Steps),
				SunColour:          utils.StepsToVO(sunColourVBox.Steps),
				MoonIlluminance:    utils.StepsToVO(moonIlluminanceVBox.Steps),
				MoonColour:         moonColourEntry.Text,
				OrbitalOffset:      utils.ToFloat64(orbitalOffsetEntry),
				Desaturation:       utils.ToFloat64(desaturationEntry),
				AmbientIlluminance: utils.ToFloat64(ambientIlluminanceEntry),
				AmbientColour:      ambientColourEntry.Text,
			},
		}

		cmd.Perform()
	})

	c := container.NewVBox(save, acc)
	return c
}
