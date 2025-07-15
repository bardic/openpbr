package data

import "embed"

type BaseConf struct {
	Out     string
	Default string
}

type Lighting struct {
	BaseConf
	SunIlluminance     []EntryViewVO
	SunColour          []EntryViewVO
	MoonIlluminance    []EntryViewVO
	MoonColour         string
	OrbitalOffset      float64
	Desaturation       float64
	AmbientIlluminance float64
	AmbientColour      string
}

type EntryViewVO struct {
	Key   string
	Value float64
	Last  bool
}
type EntryViewStrVO struct {
	Key   string
	Value string
	Last  bool
}

type Atmospherics struct {
	BaseConf
	HorizonBlendStopsMin      []EntryViewVO
	HorizonBlendStopsStart    []EntryViewVO
	HorizonBlendStopsMieStart []EntryViewVO
	MieStart                  []EntryViewVO
	HorizonBlendMax           []EntryViewVO
	RayleighStrength          []EntryViewVO
	SunMieStrength            []EntryViewVO
	MoonMieStrength           []EntryViewVO
	SunGlareShape             []EntryViewVO
	SkyZenithColor            []EntryViewStrVO
	SkyHorizonColor           []EntryViewStrVO
}

type ColorGrading struct {
	BaseConf
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
	BaseConf
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
	BaseConf
	GlobalBlockR    float64
	GlobalBlockG    float64
	GlobalBlockB    float64
	GlobalBlockA    float64
	GlobalActorR    float64
	GlobalActorG    float64
	GlobalActorB    float64
	GlobalActorA    float64
	GlobalParticleR float64
	GlobalParticleG float64
	GlobalParticleB float64
	GlobalParticleA float64
	GlobalItemR     float64
	GlobalItemG     float64
	GlobalItemB     float64
	GlobalItemA     float64
}

type Water struct {
	BaseConf
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

type TemplateSettings struct {
	TemplatePath string
	Output       string
	DefaultData  string
}

type GithubRelease struct {
	ZipballUrl string
}

type Export struct {
	Out string
}

type PBRExport struct {
	templates  embed.FS
	Out        string
	Colour     string
	MerArr     string
	MerFile    string
	Height     string
	UseMerFile bool
	Capability string
}

type Config struct {
	BaseConf
	BuildName      string
	Name           string
	Description    string
	HeaderUuid     string
	ModuleUuid     string
	DefaultMer     string
	Version        string
	Author         string
	License        string
	URL            string
	Capability     string
	HeightTemplate string
	MerTemplate    string
	ROffset        string
	GOffset        string
	BOffset        string
}
