package plugins

import "fmt"

type PluginType interface {
	String() string
	Key() string
	Match(arg string) bool
	Versions() VersionSlice
	Download(version string)
	Extract()
}

type Plugin struct {
	Repo     string
	Author   string
	versions VersionSlice
}

func (p *Plugin) String() string {
	return fmt.Sprintf("%v (by %v) [%s]", p.Repo, p.Author, p.versions.Latest().String())
}
func (p *Plugin) Key() string             { return p.Repo }
func (p *Plugin) Versions() VersionSlice  { return p.versions }
func (p *Plugin) Match(arg string) bool   { return arg == p.Repo }
func (p *Plugin) Download(version string) {}
func (p *Plugin) Extract()                {}

var Plugins = parsePlugin()

func FindPlugin(arg string) (PluginType, error) {
	for _, plug := range Plugins {
		if plug.Match(arg) {
			return plug, nil
		}
	}
	return &Plugin{}, fmt.Errorf("%s does not match any plugin", arg)
}
