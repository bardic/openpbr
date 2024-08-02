package data

type Manifest struct {
	VersionStr string
	VersionArr string
}

type PBR struct {
	Colour  string
	MerType int
	MerFile string
	MerArr  string
	Height  string
}

type GithubRelease struct {
	Zipball_url string
}
