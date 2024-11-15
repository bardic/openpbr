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
	horizonBlendStopsMinVBox      *vo.EntryView
	horizonBlendStopsStartVBox    *vo.EntryView
	horizonBlendStopsMieStartVBox *vo.EntryView
	mieStartVBox                  *vo.EntryView
	horizonBlendMaxVBox           *vo.EntryView
	rayleighStrengthVBox          *vo.EntryView
	sunMieStrengthVBox            *vo.EntryView
	moonMieStrengthVBox           *vo.EntryView
	sunGlareShapeVBox             *vo.EntryView
	skyZenithColorVBox            *vo.EntryView
	skyHorizonColorVBox           *vo.EntryView
}

func (v *Atmospherics) Build(refresh func(), popupErr func(error)) *fyne.Container {
	//
	// Horizon Blend Stops Min
	//

	v.horizonBlendStopsMinVBox = utils.CreateEntryView("Horizon Blend Stops Min", id)

	//
	// Horizon Blend Stops Start
	//

	v.horizonBlendStopsStartVBox = utils.CreateEntryView("Horizon Blend Stops Start", id)

	//
	// Horizon Blend Stops Mie Start
	//

	v.horizonBlendStopsMieStartVBox = utils.CreateEntryView("Horizon Blend Stops Mie Start", id)

	//
	// Mie Start
	//

	v.mieStartVBox = utils.CreateEntryView("Mie Start", id)

	//
	// Horizon Blend Max
	//

	v.horizonBlendMaxVBox = utils.CreateEntryView("Horizon Blend Max", id)

	//
	// Rayleigh Strength
	//

	v.rayleighStrengthVBox = utils.CreateEntryView("Rayleigh Strength", id)

	//
	// Sun Mie Strength
	//

	v.sunMieStrengthVBox = utils.CreateEntryView("Sun Mie Strength", id)

	//
	// Moon Mie Strength
	//

	v.moonMieStrengthVBox = utils.CreateEntryView("Moon Mie Strength", id)

	//
	// Sun Glare Shape
	//

	v.sunGlareShapeVBox = utils.CreateEntryView("Sun Glare Shape", id)

	//
	// Sky Zenith Color
	//

	v.skyZenithColorVBox = utils.CreateEntryView("Sky Zenith Color", id)

	//
	// Sky Horizon Color
	//

	v.skyHorizonColorVBox = utils.CreateEntryView("Sky Horizon Color", id)

	accItem1 := widget.NewAccordionItem("Horizon Blend Stops Min", v.horizonBlendStopsMinVBox.C)
	accItem2 := widget.NewAccordionItem("Horizon Blend Stops Start", v.horizonBlendStopsStartVBox.C)
	accItem3 := widget.NewAccordionItem("Horizon Blend Stops Mie Start", v.horizonBlendStopsMieStartVBox.C)
	accItem4 := widget.NewAccordionItem("Mie Start", v.mieStartVBox.C)
	accItem5 := widget.NewAccordionItem("Horizon Blend Max", v.horizonBlendMaxVBox.C)
	accItem6 := widget.NewAccordionItem("Rayleigh Strength", v.rayleighStrengthVBox.C)
	accItem7 := widget.NewAccordionItem("Sun Mie Strength", v.sunMieStrengthVBox.C)
	accItem8 := widget.NewAccordionItem("Moon Mie Strength", v.moonMieStrengthVBox.C)
	accItem9 := widget.NewAccordionItem("Sun Glare Shape", v.sunGlareShapeVBox.C)
	accItem10 := widget.NewAccordionItem("Sky Zenith Color", v.skyZenithColorVBox.C)
	accItem11 := widget.NewAccordionItem("Sky Horizon Color", v.skyHorizonColorVBox.C)

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
			Out: "./example/settings/shared/atmospherics/atmospherics.json",
			Atmospherics: vo.Atmospherics{
				HorizonBlendStopsMin:      utils.StepsToVO(v.horizonBlendStopsMinVBox.Steps),
				HorizonBlendStopsStart:    utils.StepsToVO(v.horizonBlendStopsStartVBox.Steps),
				HorizonBlendStopsMieStart: utils.StepsToVO(v.horizonBlendStopsMieStartVBox.Steps),
				MieStart:                  utils.StepsToVO(v.mieStartVBox.Steps),
				HorizonBlendMax:           utils.StepsToVO(v.horizonBlendMaxVBox.Steps),
				RayleighStrength:          utils.StepsToVO(v.rayleighStrengthVBox.Steps),
				SunMieStrength:            utils.StepsToVO(v.sunMieStrengthVBox.Steps),
				MoonMieStrength:           utils.StepsToVO(v.moonMieStrengthVBox.Steps),
				SunGlareShape:             utils.StepsToVO(v.sunGlareShapeVBox.Steps),
				SkyZenithColor:            utils.StepsToVO(v.skyZenithColorVBox.Steps),
				SkyHorizonColor:           utils.StepsToVO(v.skyHorizonColorVBox.Steps),
			},
		}

		cmd.Perform()
	})

	c := container.NewVBox(save, acc)
	return c
}

func (v *Atmospherics) Defaults() {
	for range 3 {
		p := utils.CreateEntryViewHolder()
		p.KeyEntry.SetText("Key")
		p.ValueEntry.SetText("Value")
		v.horizonBlendStopsMinVBox.Steps = append(v.horizonBlendStopsMinVBox.Steps, p)
		v.horizonBlendStopsMinVBox.C.Add(v.horizonBlendStopsMinVBox.Steps[len(v.horizonBlendStopsMinVBox.Steps)-1].HBox)
	}

}
