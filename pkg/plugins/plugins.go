package plugins

type PluginType interface {
	Name() string
	String() string
	Versions() ([]Version, error)
	Download(version string)
	Extract()
}

type Plugin struct {
	Name         string
	Dependencies []PluginType
}

type Version struct {
	Tag         string
	CreatedDate string
	TarBallUrl  string
	ZipBallUrl  string
}
