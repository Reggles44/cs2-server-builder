package lookup

import (
	"strings"

	"github.com/reggles44/cs2-server-builder/pkg/plugins"
	"github.com/reggles44/cs2-server-builder/pkg/plugins/metamod"
)

var Plugins = []plugins.PluginType{
	&metamod.MetaModPlugin,
}

func FindPlugin(name string) *plugins.PluginType {
	for _, p := range Plugins {
		if strings.ToLower("") == strings.ToLower(name) {
			return &p
		}
	}

	return nil
}
