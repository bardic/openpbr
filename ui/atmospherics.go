package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type Atmospherics struct {
}

func (v *Atmospherics) Build(refresh func(), popupErr func(error)) *fyne.Container {
	//
	// Horizon Blend Stops Min
	//

	horizonBlendStopsMinVBox := utils.CreateEntryView("Horizon Blend Stops Min", id)

	//
	// Horizon Blend Stops Start
	//

	horizonBlendStopsStartVBox := utils.CreateEntryView("Horizon Blend Stops Start", id)

	//
	// Horizon Blend Stops Mie Start
	//

	horizonBlendStopsMieStartVBox := utils.CreateEntryView("Horizon Blend Stops Mie Start", id)

	//
	// Mie Start
	//

	mieStartVBox := utils.CreateEntryView("Mie Start", id)

	//
	// Horizon Blend Max
	//

	horizonBlendMaxVBox := utils.CreateEntryView("Horizon Blend Max", id)

	//
	// Rayleigh Strength
	//

	rayleighStrengthVBox := utils.CreateEntryView("Rayleigh Strength", id)

	//
	// Sun Mie Strength
	//

	sunMieStrengthVBox := utils.CreateEntryView("Sun Mie Strength", id)

	//
	// Moon Mie Strength
	//

	moonMieStrengthVBox := utils.CreateEntryView("Moon Mie Strength", id)

	//
	// Sun Glare Shape
	//

	sunGlareShapeVBox := utils.CreateEntryView("Sun Glare Shape", id)

	//
	// Sky Zenith Color
	//

	skyZenithColorVBox := utils.CreateEntryView("Sky Zenith Color", id)

	//
	// Sky Horizon Color
	//

	skyHorizonColorVBox := utils.CreateEntryView("Sky Horizon Color", id)

	accItem1 := widget.NewAccordionItem("Horizon Blend Stops Min", horizonBlendStopsMinVBox.C)
	accItem2 := widget.NewAccordionItem("Horizon Blend Stops Start", horizonBlendStopsStartVBox.C)
	accItem3 := widget.NewAccordionItem("Horizon Blend Stops Mie Start", horizonBlendStopsMieStartVBox.C)
	accItem4 := widget.NewAccordionItem("Mie Start", mieStartVBox.C)
	accItem5 := widget.NewAccordionItem("Horizon Blend Max", horizonBlendMaxVBox.C)
	accItem6 := widget.NewAccordionItem("Rayleigh Strength", rayleighStrengthVBox.C)
	accItem7 := widget.NewAccordionItem("Sun Mie Strength", sunMieStrengthVBox.C)
	accItem8 := widget.NewAccordionItem("Moon Mie Strength", moonMieStrengthVBox.C)
	accItem9 := widget.NewAccordionItem("Sun Glare Shape", sunGlareShapeVBox.C)
	accItem10 := widget.NewAccordionItem("Sky Zenith Color", skyZenithColorVBox.C)
	accItem11 := widget.NewAccordionItem("Sky Horizon Color", skyHorizonColorVBox.C)

	acc := widget.NewAccordion(
		accItem1,
		accItem2,
		accItem3,
		accItem4,
		accItem5,
		accItem6,
		accItem7,
		accItem8,
		accItem9,
		accItem10,
		accItem11,
	)

	save := widget.NewButton("Save", func() {
		cmd := export.Atmospherics{
			Out: "./openpbr_out/atmospherics/atmospherics.json",
			Atmospherics: vo.Atmospherics{
				HorizonBlendStopsMin:      utils.StepsToVO(horizonBlendStopsMinVBox.Steps),
				HorizonBlendStopsStart:    utils.StepsToVO(horizonBlendStopsStartVBox.Steps),
				HorizonBlendStopsMieStart: utils.StepsToVO(horizonBlendStopsMieStartVBox.Steps),
				MieStart:                  utils.StepsToVO(mieStartVBox.Steps),
				HorizonBlendMax:           utils.StepsToVO(horizonBlendMaxVBox.Steps),
				RayleighStrength:          utils.StepsToVO(rayleighStrengthVBox.Steps),
				SunMieStrength:            utils.StepsToVO(sunMieStrengthVBox.Steps),
				MoonMieStrength:           utils.StepsToVO(moonMieStrengthVBox.Steps),
				SunGlareShape:             utils.StepsToVO(sunGlareShapeVBox.Steps),
				SkyZenithColor:            utils.StepsToVO(skyZenithColorVBox.Steps),
				SkyHorizonColor:           utils.StepsToVO(skyHorizonColorVBox.Steps),
			},
		}

		cmd.Perform()
	})

	c := container.NewVBox(save, acc)
	return c
}
