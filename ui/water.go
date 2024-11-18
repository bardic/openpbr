package ui

import (
	"encoding/json"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type Water struct {
	chlorophyllEntry           *widget.Entry
	suspendedSedimentEntry     *widget.Entry
	cdomEntry                  *widget.Entry
	wavesEnabledEntry          *widget.Check
	wavesDepthEntry            *widget.Entry
	wavesFrequencyEntry        *widget.Entry
	wavesFrequencyScalingEntry *widget.Entry
	wavesMixEntry              *widget.Entry
	wavesOctavesEntry          *widget.Entry
	wavesPullEntry             *widget.Entry
	wavesSampleWidthEntry      *widget.Entry
	wavesShapeEntry            *widget.Entry
	wavesSpeedEntry            *widget.Entry
	wavesSpeedScalingEntry     *widget.Entry
}

func (v *Water) Build(p fyne.Window) *fyne.Container {
	//
	// Chlorophyll
	//

	chlorophyllLabel := widget.NewLabel("Chlorophyll")
	v.chlorophyllEntry = widget.NewEntry()

	//
	// Suspended Sediment
	//

	suspendedSedimentLabel := widget.NewLabel("Suspended Sediment")
	v.suspendedSedimentEntry = widget.NewEntry()

	//
	// CDOM
	//

	cdomLabel := widget.NewLabel("CDOM")
	v.cdomEntry = widget.NewEntry()

	//
	// Waves Enabled
	//

	wavesEnabledLabel := widget.NewLabel("Waves Enabled")
	v.wavesEnabledEntry = widget.NewCheck("Waves Enabled", func(bool) {})

	//
	// Waves Depth
	//

	wavesDepthLabel := widget.NewLabel("Waves Depth")
	v.wavesDepthEntry = widget.NewEntry()

	//
	// Waves Frequency
	//

	wavesFrequencyLabel := widget.NewLabel("Waves Frequency")
	v.wavesFrequencyEntry = widget.NewEntry()

	//
	// Waves Frequency Scaling
	//

	wavesFrequencyScalingLabel := widget.NewLabel("Waves Frequency Scaling")
	v.wavesFrequencyScalingEntry = widget.NewEntry()

	//
	// Waves Mix
	//

	wavesMixLabel := widget.NewLabel("Waves Mix")
	v.wavesMixEntry = widget.NewEntry()

	//
	// Waves Octaves
	//

	wavesOctavesLabel := widget.NewLabel("Waves Octaves")
	v.wavesOctavesEntry = widget.NewEntry()

	//
	// Waves Pull
	//

	wavesPullLabel := widget.NewLabel("Waves Pull")
	v.wavesPullEntry = widget.NewEntry()

	//
	// Waves Sample Width
	//

	wavesSampleWidthLabel := widget.NewLabel("Waves Sample Width")
	v.wavesSampleWidthEntry = widget.NewEntry()

	//
	// Waves Shape
	//

	wavesShapeLabel := widget.NewLabel("Waves Shape")
	v.wavesShapeEntry = widget.NewEntry()

	//
	// Waves Speed
	//

	wavesSpeedLabel := widget.NewLabel("Waves Speed")
	v.wavesSpeedEntry = widget.NewEntry()

	//
	// Waves Speed Scaling
	//

	wavesSpeedScalingLabel := widget.NewLabel("Waves Speed Scaling")
	v.wavesSpeedScalingEntry = widget.NewEntry()

	c := container.New(layout.NewFormLayout(),
		chlorophyllLabel, v.chlorophyllEntry,
		suspendedSedimentLabel, v.suspendedSedimentEntry,
		cdomLabel, v.cdomEntry,
		wavesEnabledLabel, v.wavesEnabledEntry,
		wavesDepthLabel, v.wavesDepthEntry,
		wavesFrequencyLabel, v.wavesFrequencyEntry,
		wavesFrequencyScalingLabel, v.wavesFrequencyScalingEntry,
		wavesMixLabel, v.wavesMixEntry,
		wavesOctavesLabel, v.wavesOctavesEntry,
		wavesPullLabel, v.wavesPullEntry,
		wavesSampleWidthLabel, v.wavesSampleWidthEntry,
		wavesShapeLabel, v.wavesShapeEntry,
		wavesSpeedLabel, v.wavesSpeedEntry,
		wavesSpeedScalingLabel, v.wavesSpeedScalingEntry,
	)
	return c
}

func (v *Water) Defaults(b []byte) {

	var c vo.Water
	json.Unmarshal(b, &c)

	v.chlorophyllEntry.SetText(utils.FloatToString(c.Chlorophyll))
	v.suspendedSedimentEntry.SetText(utils.FloatToString(c.SuspendedSediment))
	v.cdomEntry.SetText(utils.FloatToString(c.CDOM))
	v.wavesEnabledEntry.Checked = c.WavesEnabled
	v.wavesDepthEntry.SetText(utils.FloatToString(c.WavesDepth))
	v.wavesFrequencyEntry.SetText(utils.FloatToString(c.WavesFrequency))
	v.wavesFrequencyScalingEntry.SetText(utils.FloatToString(c.WavesFrequencyScaling))
	v.wavesMixEntry.SetText(utils.FloatToString(c.WavesMix))
	v.wavesOctavesEntry.SetText(utils.FloatToString(c.WavesOctaves))
	v.wavesPullEntry.SetText(utils.FloatToString(c.WavesPull))
	v.wavesSampleWidthEntry.SetText(utils.FloatToString(c.WavesSampleWidth))
	v.wavesShapeEntry.SetText(utils.FloatToString(c.WavesShape))
	v.wavesSpeedEntry.SetText(utils.FloatToString(c.WavesSpeed))
	v.wavesSpeedScalingEntry.SetText(utils.FloatToString(c.WavesSpeedScaling))
}

func (v *Water) Save() {
	water := export.Water{
		Water: vo.Water{
			BaseConf: vo.BaseConf{
				Out: path.Join(store.PackageStore, "water"),
			},
			Chlorophyll:           utils.ToFloat64(v.chlorophyllEntry),
			SuspendedSediment:     utils.ToFloat64(v.suspendedSedimentEntry),
			CDOM:                  utils.ToFloat64(v.cdomEntry),
			WavesEnabled:          v.wavesEnabledEntry.Checked,
			WavesDepth:            utils.ToFloat64(v.wavesDepthEntry),
			WavesFrequency:        utils.ToFloat64(v.wavesFrequencyEntry),
			WavesFrequencyScaling: utils.ToFloat64(v.wavesFrequencyScalingEntry),
			WavesMix:              utils.ToFloat64(v.wavesMixEntry),
			WavesOctaves:          utils.ToFloat64(v.wavesOctavesEntry),
			WavesPull:             utils.ToFloat64(v.wavesPullEntry),
			WavesSampleWidth:      utils.ToFloat64(v.wavesSampleWidthEntry),
			WavesShape:            utils.ToFloat64(v.wavesShapeEntry),
			WavesSpeed:            utils.ToFloat64(v.wavesSpeedEntry),
			WavesSpeedScaling:     utils.ToFloat64(v.wavesSpeedScalingEntry),
		},
	}

	water.Save()

}
