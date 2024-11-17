package vo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Lighting struct {
	SunIlluminance     []EntryViewVO
	SunColour          []EntryViewVO
	MoonIlluminance    []EntryViewVO
	MoonColour         string
	OrbitalOffset      float64
	Desaturation       float64
	AmbientIlluminance float64
	AmbientColour      string
}

type EntryViewHolder struct {
	Id           int
	HBox         *fyne.Container
	KeyEntry     *widget.Entry
	ValueEntry   *widget.Entry
	DeleteButton *widget.Button
}

type EntryViewVO struct {
	Key   string
	Value float64
	Last  bool
}

type Atmospherics struct {
	HorizonBlendStopsMin      []EntryViewVO
	HorizonBlendStopsStart    []EntryViewVO
	HorizonBlendStopsMieStart []EntryViewVO
	MieStart                  []EntryViewVO
	HorizonBlendMax           []EntryViewVO
	RayleighStrength          []EntryViewVO
	SunMieStrength            []EntryViewVO
	MoonMieStrength           []EntryViewVO
	SunGlareShape             []EntryViewVO
	SkyZenithColor            []EntryViewVO
	SkyHorizonColor           []EntryViewVO
}

type ColorGrading struct {
	HighlightsContrastG   float64
	HighlightsContrastB   float64
	HighlightsContrastR   float64
	HighlightsGainR       float64
	HighlightsGainG       float64
	HighlightsGainB       float64
	HighlightsGammaR      float64
	HighlightsGammaG      float64
	HighlightsGammaB      float64
	HighlightsOffsetR     float64
	HighlightsOffsetG     float64
	HighlightsOffsetB     float64
	HighlightsSaturationR float64
	HighlightsSaturationG float64
	HighlightsSaturationB float64
	MidtonesContrastR     float64
	MidtonesContrastG     float64
	MidtonesContrastB     float64
	MidtonesGainR         float64
	MidtonesGainG         float64
	MidtonesGainB         float64
	MidtonesGammaR        float64
	MidtonesGammaG        float64
	MidtonesGammaB        float64
	MidtonesOffsetR       float64
	MidtonesOffsetG       float64
	MidtonesOffsetB       float64
	MidtonesSaturationR   float64
	MidtonesSaturationG   float64
	MidtonesSaturationB   float64
	ShadowsMax            float64
	ShadowsContrastR      float64
	ShadowsContrastG      float64
	ShadowsContrastB      float64
	ShadowsGainR          float64
	ShadowsGainG          float64
	ShadowsGainB          float64
	ShadowsGammaR         float64
	ShadowsGammaG         float64
	ShadowsGammaB         float64
	ShadowsOffsetR        float64
	ShadowsOffsetG        float64
	ShadowsOffsetB        float64
	ShadowsSaturationR    float64
	ShadowsSaturationG    float64
	ShadowsSaturationB    float64
	ToneMappingOperator   string
}

type Fog struct {
	WaterMaxDensity      float64
	WaterUniformDensity  bool
	AirMaxDensity        float64
	AirZeroDensityHeight float64
	AirMaxDensityHeight  float64
	WaterScatteringR     float64
	WaterScatteringG     float64
	WaterScatteringB     float64
	WaterAbsorptionR     float64
	WaterAbsorptionG     float64
	WaterAbsorptionB     float64
	AirScatteringG       float64
	AirScatteringB       float64
	AirScatteringR       float64
	AirAbsorptionR       float64
	AirAbsorptionG       float64
	AirAbsorptionB       float64
}

type PBR struct {
	GlobalBlockMERS    string
	GlobalActorMERS    string
	GlobalParticleMERS string
	GlobalItemMERS     string
}

type Water struct {
	Chlorophyll           float64
	SuspendedSediment     float64
	CDOM                  float64
	WavesEnabled          bool
	WavesDepth            float64
	WavesFrequency        float64
	WavesFrequencyScaling float64
	WavesMix              float64
	WavesOctaves          float64
	WavesPull             float64
	WavesSampleWidth      float64
	WavesShape            float64
	WavesSpeed            float64
	WavesSpeedScaling     float64
}

type RGB struct {
	R *widget.Entry
	G *widget.Entry
	B *widget.Entry
	C *fyne.Container
}

type RGBA struct {
	RGB
	A *widget.Entry
}

type EntryView struct {
	Steps []*EntryViewHolder
	C     *fyne.Container
}
