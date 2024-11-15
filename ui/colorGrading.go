package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type ColorGrading struct {
	RGBs map[string]vo.RGB
}

func (v *ColorGrading) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {

	rgbEntryiesTitles := []string{
		"Highlights Contrast",
		"Highlights Gain",
		"Highlights Gamma",
		"Highlights Offset",
		"Highlights Saturation",
		"Midtones Contrast",
		"Midtones Gain",
		"Midtones Gamma",
		"Midtones Offset",
		"Midtones Saturation",
		"Shadows Max",
		"Shadows Contrast",
		"Shadows Gain",
		"Shadows Gamma",
		"Shadows Offset",
		"Shadows Saturation",
	}

	// rgbEntryies := make([]fyne.CanvasObject, 0)

	save := widget.NewButton("Save", func() {
		colorGrading := export.ColorGrading{
			Out: "./example/settings/shared/color_grading/color_grading.json",
			ColorGrading: vo.ColorGrading{
				HighlightsContrastG:   utils.ToFloat64(v.RGBs[rgbEntryiesTitles[0]].G),
				HighlightsContrastB:   utils.ToFloat64(v.RGBs[rgbEntryiesTitles[0]].B),
				HighlightsContrastR:   utils.ToFloat64(v.RGBs[rgbEntryiesTitles[0]].R),
				HighlightsGainR:       utils.ToFloat64(v.RGBs[rgbEntryiesTitles[1]].R),
				HighlightsGainG:       utils.ToFloat64(v.RGBs[rgbEntryiesTitles[1]].G),
				HighlightsGainB:       utils.ToFloat64(v.RGBs[rgbEntryiesTitles[1]].B),
				HighlightsGammaR:      utils.ToFloat64(v.RGBs[rgbEntryiesTitles[2]].R),
				HighlightsGammaG:      utils.ToFloat64(v.RGBs[rgbEntryiesTitles[2]].G),
				HighlightsGammaB:      utils.ToFloat64(v.RGBs[rgbEntryiesTitles[2]].B),
				HighlightsOffsetR:     utils.ToFloat64(v.RGBs[rgbEntryiesTitles[3]].R),
				HighlightsOffsetG:     utils.ToFloat64(v.RGBs[rgbEntryiesTitles[3]].G),
				HighlightsOffsetB:     utils.ToFloat64(v.RGBs[rgbEntryiesTitles[3]].B),
				HighlightsSaturationR: utils.ToFloat64(v.RGBs[rgbEntryiesTitles[4]].R),
				HighlightsSaturationG: utils.ToFloat64(v.RGBs[rgbEntryiesTitles[4]].G),
				HighlightsSaturationB: utils.ToFloat64(v.RGBs[rgbEntryiesTitles[4]].B),
				MidtonesContrastR:     utils.ToFloat64(v.RGBs[rgbEntryiesTitles[5]].R),
				MidtonesContrastG:     utils.ToFloat64(v.RGBs[rgbEntryiesTitles[5]].G),
				MidtonesContrastB:     utils.ToFloat64(v.RGBs[rgbEntryiesTitles[5]].B),
				MidtonesGainR:         utils.ToFloat64(v.RGBs[rgbEntryiesTitles[6]].R),
				MidtonesGainG:         utils.ToFloat64(v.RGBs[rgbEntryiesTitles[6]].G),
				MidtonesGainB:         utils.ToFloat64(v.RGBs[rgbEntryiesTitles[6]].B),
				MidtonesGammaR:        utils.ToFloat64(v.RGBs[rgbEntryiesTitles[7]].R),
				MidtonesGammaG:        utils.ToFloat64(v.RGBs[rgbEntryiesTitles[7]].G),
				MidtonesGammaB:        utils.ToFloat64(v.RGBs[rgbEntryiesTitles[7]].B),
				MidtonesOffsetR:       utils.ToFloat64(v.RGBs[rgbEntryiesTitles[8]].R),
				MidtonesOffsetG:       utils.ToFloat64(v.RGBs[rgbEntryiesTitles[8]].G),
				MidtonesOffsetB:       utils.ToFloat64(v.RGBs[rgbEntryiesTitles[8]].B),
				MidtonesSaturationR:   utils.ToFloat64(v.RGBs[rgbEntryiesTitles[9]].R),
				MidtonesSaturationG:   utils.ToFloat64(v.RGBs[rgbEntryiesTitles[9]].G),
				MidtonesSaturationB:   utils.ToFloat64(v.RGBs[rgbEntryiesTitles[9]].B),
				ShadowsMax:            utils.ToFloat64(v.RGBs[rgbEntryiesTitles[10]].R),
				ShadowsContrastR:      utils.ToFloat64(v.RGBs[rgbEntryiesTitles[11]].R),
				ShadowsContrastG:      utils.ToFloat64(v.RGBs[rgbEntryiesTitles[11]].G),
				ShadowsContrastB:      utils.ToFloat64(v.RGBs[rgbEntryiesTitles[11]].B),
				ShadowsGainR:          utils.ToFloat64(v.RGBs[rgbEntryiesTitles[12]].R),
				ShadowsGainG:          utils.ToFloat64(v.RGBs[rgbEntryiesTitles[12]].G),
				ShadowsGainB:          utils.ToFloat64(v.RGBs[rgbEntryiesTitles[12]].B),
				ShadowsGammaR:         utils.ToFloat64(v.RGBs[rgbEntryiesTitles[13]].R),
				ShadowsGammaG:         utils.ToFloat64(v.RGBs[rgbEntryiesTitles[13]].G),
				ShadowsGammaB:         utils.ToFloat64(v.RGBs[rgbEntryiesTitles[13]].B),
				ShadowsOffsetR:        utils.ToFloat64(v.RGBs[rgbEntryiesTitles[14]].R),
				ShadowsOffsetG:        utils.ToFloat64(v.RGBs[rgbEntryiesTitles[14]].G),
				ShadowsOffsetB:        utils.ToFloat64(v.RGBs[rgbEntryiesTitles[14]].B),
				ShadowsSaturationR:    utils.ToFloat64(v.RGBs[rgbEntryiesTitles[15]].R),
				ShadowsSaturationG:    utils.ToFloat64(v.RGBs[rgbEntryiesTitles[15]].G),
				ShadowsSaturationB:    utils.ToFloat64(v.RGBs[rgbEntryiesTitles[15]].B),
				ToneMappingOperator:   "Hable",
			},
		}

		colorGrading.Perform()
	})

	c := container.New(layout.NewFormLayout())

	for _, title := range rgbEntryiesTitles {

		label := widget.NewLabel(title)
		rLabel := widget.NewLabel("R")
		rEntry := widget.NewEntry()
		gLabel := widget.NewLabel("G")
		gEntry := widget.NewEntry()
		bLabel := widget.NewLabel("B")
		bEntry := widget.NewEntry()
		hbox := container.NewHBox(rLabel, rEntry, gLabel, gEntry, bLabel, bEntry)

		c.Add(label)
		c.Add(hbox)

		rgb := vo.RGB{
			R: rEntry,
			G: gEntry,
			B: bEntry,
		}

		if v.RGBs == nil {
			v.RGBs = make(map[string]vo.RGB)
		}

		v.RGBs[title] = rgb
	}

	c.Add(save)
	c.Add(layout.NewSpacer())

	return c
}
