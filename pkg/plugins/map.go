package plugins

import "strings"

var Plugins = []*PluginType{

}

func FindPlugin(name string) *Plugin {
	for pluginName, plugin := range Map {
		if strings.ToLower(pluginName) == strings.ToLower(name) {
			return plugin
		}
	}

	return nil
}
