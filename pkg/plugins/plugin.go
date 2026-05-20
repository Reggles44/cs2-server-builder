package plugins

type PluginType interface {
	Versions() []Version
	Download(version string)
	Extract()
	String() string
}

type Plugin struct {
	Name         string
	Dependencies []*Plugin
}

type Version struct {
	Tag         string
	CreatedDate string
	TarBallUrl  string
	ZipBallUrl  string
}
