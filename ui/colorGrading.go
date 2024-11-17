package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type ColorGrading struct {
	RGBs                map[string]vo.RGB
	RBGEntryiesTitles   []string
	ShadowsMaxEntry     *widget.Entry
	ToneMappingOperator *widget.Entry
}

func (v *ColorGrading) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {

	v.RBGEntryiesTitles = []string{
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
		// "Shadows Max",
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
				HighlightsContrastR:   utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[0]].R),
				HighlightsContrastG:   utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[0]].G),
				HighlightsContrastB:   utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[0]].B),
				HighlightsGainR:       utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[1]].R),
				HighlightsGainG:       utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[1]].G),
				HighlightsGainB:       utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[1]].B),
				HighlightsGammaR:      utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[2]].R),
				HighlightsGammaG:      utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[2]].G),
				HighlightsGammaB:      utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[2]].B),
				HighlightsOffsetR:     utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[3]].R),
				HighlightsOffsetG:     utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[3]].G),
				HighlightsOffsetB:     utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[3]].B),
				HighlightsSaturationR: utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[4]].R),
				HighlightsSaturationG: utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[4]].G),
				HighlightsSaturationB: utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[4]].B),
				MidtonesContrastR:     utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[5]].R),
				MidtonesContrastG:     utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[5]].G),
				MidtonesContrastB:     utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[5]].B),
				MidtonesGainR:         utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[6]].R),
				MidtonesGainG:         utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[6]].G),
				MidtonesGainB:         utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[6]].B),
				MidtonesGammaR:        utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[7]].R),
				MidtonesGammaG:        utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[7]].G),
				MidtonesGammaB:        utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[7]].B),
				MidtonesOffsetR:       utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[8]].R),
				MidtonesOffsetG:       utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[8]].G),
				MidtonesOffsetB:       utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[8]].B),
				MidtonesSaturationR:   utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[9]].R),
				MidtonesSaturationG:   utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[9]].G),
				MidtonesSaturationB:   utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[9]].B),
				ShadowsMax:            utils.ToFloat64(v.ShadowsMaxEntry),
				ShadowsContrastR:      utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[11]].R),
				ShadowsContrastG:      utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[11]].G),
				ShadowsContrastB:      utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[11]].B),
				ShadowsGainR:          utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[12]].R),
				ShadowsGainG:          utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[12]].G),
				ShadowsGainB:          utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[12]].B),
				ShadowsGammaR:         utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[13]].R),
				ShadowsGammaG:         utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[13]].G),
				ShadowsGammaB:         utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[13]].B),
				ShadowsOffsetR:        utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[14]].R),
				ShadowsOffsetG:        utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[14]].G),
				ShadowsOffsetB:        utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[14]].B),
				ShadowsSaturationR:    utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[15]].R),
				ShadowsSaturationG:    utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[15]].G),
				ShadowsSaturationB:    utils.ToFloat64(v.RGBs[v.RBGEntryiesTitles[15]].B),
				ToneMappingOperator:   v.ToneMappingOperator.Text,
			},
		}

		colorGrading.Perform()
	})

	c := container.New(layout.NewFormLayout())

	for _, title := range v.RBGEntryiesTitles {

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

	c.Add(widget.NewLabel("Shadows Max"))
	v.ShadowsMaxEntry = widget.NewEntry()
	c.Add(v.ShadowsMaxEntry)

	c.Add(widget.NewLabel("Tone Mapping Operator"))
	v.ToneMappingOperator = widget.NewEntry()
	c.Add(v.ToneMappingOperator)

	c.Add(save)
	c.Add(layout.NewSpacer())

	return c
}

func (v *ColorGrading) Defaults(c *vo.ColorGrading) {
	v.RGBs[v.RBGEntryiesTitles[0]].R.Text = strconv.FormatFloat(c.HighlightsContrastR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[0]].G.Text = strconv.FormatFloat(c.HighlightsContrastG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[0]].B.Text = strconv.FormatFloat(c.HighlightsContrastB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[1]].R.Text = strconv.FormatFloat(c.HighlightsGainR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[1]].G.Text = strconv.FormatFloat(c.HighlightsGainG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[1]].B.Text = strconv.FormatFloat(c.HighlightsGainB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[2]].R.Text = strconv.FormatFloat(c.HighlightsGammaR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[2]].G.Text = strconv.FormatFloat(c.HighlightsGammaG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[2]].B.Text = strconv.FormatFloat(c.HighlightsGammaB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[3]].R.Text = strconv.FormatFloat(c.HighlightsOffsetR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[3]].G.Text = strconv.FormatFloat(c.HighlightsOffsetG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[3]].B.Text = strconv.FormatFloat(c.HighlightsOffsetB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[4]].R.Text = strconv.FormatFloat(c.HighlightsSaturationR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[4]].G.Text = strconv.FormatFloat(c.HighlightsSaturationG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[4]].B.Text = strconv.FormatFloat(c.HighlightsSaturationB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[5]].R.Text = strconv.FormatFloat(c.MidtonesContrastR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[5]].G.Text = strconv.FormatFloat(c.MidtonesContrastG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[5]].B.Text = strconv.FormatFloat(c.MidtonesContrastB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[6]].R.Text = strconv.FormatFloat(c.MidtonesGainR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[6]].G.Text = strconv.FormatFloat(c.MidtonesGainG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[6]].B.Text = strconv.FormatFloat(c.MidtonesGainB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[7]].R.Text = strconv.FormatFloat(c.MidtonesGammaR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[7]].G.Text = strconv.FormatFloat(c.MidtonesGammaG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[7]].B.Text = strconv.FormatFloat(c.MidtonesGammaB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[8]].R.Text = strconv.FormatFloat(c.MidtonesOffsetR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[8]].G.Text = strconv.FormatFloat(c.MidtonesOffsetG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[8]].B.Text = strconv.FormatFloat(c.MidtonesOffsetB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[9]].R.Text = strconv.FormatFloat(c.MidtonesSaturationR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[9]].G.Text = strconv.FormatFloat(c.MidtonesSaturationG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[9]].B.Text = strconv.FormatFloat(c.MidtonesSaturationB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[10]].R.Text = strconv.FormatFloat(c.ShadowsContrastR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[10]].G.Text = strconv.FormatFloat(c.ShadowsContrastG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[10]].B.Text = strconv.FormatFloat(c.ShadowsContrastB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[11]].R.Text = strconv.FormatFloat(c.ShadowsGainR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[11]].G.Text = strconv.FormatFloat(c.ShadowsGainG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[11]].B.Text = strconv.FormatFloat(c.ShadowsGainB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[12]].R.Text = strconv.FormatFloat(c.ShadowsGammaR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[12]].G.Text = strconv.FormatFloat(c.ShadowsGammaG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[12]].B.Text = strconv.FormatFloat(c.ShadowsGammaB, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[13]].R.Text = strconv.FormatFloat(c.ShadowsOffsetR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[13]].G.Text = strconv.FormatFloat(c.ShadowsOffsetB, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[13]].B.Text = strconv.FormatFloat(c.ShadowsOffsetG, 'f', -1, 64)

	v.RGBs[v.RBGEntryiesTitles[14]].R.Text = strconv.FormatFloat(c.ShadowsSaturationR, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[14]].G.Text = strconv.FormatFloat(c.ShadowsSaturationG, 'f', -1, 64)
	v.RGBs[v.RBGEntryiesTitles[14]].B.Text = strconv.FormatFloat(c.ShadowsSaturationB, 'f', -1, 64)

	v.ShadowsMaxEntry.Text = strconv.FormatFloat(c.ShadowsMax, 'f', -1, 64)
	v.ToneMappingOperator.Text = c.ToneMappingOperator
}
