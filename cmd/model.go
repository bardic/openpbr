package cmd

type PBR struct {
	Colour  string
	MerType bool
	MerFile string
	MerArr  string
	Height  string
}

type GithubRelease struct {
	Zipball_url string
}

type Target struct {
	Buildname         string
	Name              string
	Header_uuid       string
	Module_uuid       string
	Description       string
	Textureset_format string
	Default_mer       string
	Version           string
	Author            string
	License           string
	URL               string
	Capibility        string
	HeightTemplate    string
	NormalTemplate    string
	MerTemplate       string
	ExportMer         bool
}

func (t *Target) Perform() error {
	return nil
}
