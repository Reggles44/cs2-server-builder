package plugins

import (
	_ "embed"
	"encoding/json"
	"log"
	"strings"
)

//go:embed plugins.json
var pluginsData []byte

type pluginData struct {
	Repo     string     `json:"repo"`
	Versions []*Version `json:"releases"`
}

func parsePlugin() []*Plugin {
	var raw []*pluginData
	err := json.Unmarshal(pluginsData, &raw)
	if err != nil {
		log.Panic(err)
	}

	var plugins []*Plugin

	for _, p := range raw {
		parts := strings.Split(p.Repo, "/")
		author, repo := parts[0], parts[1]
		plugins = append(plugins, &Plugin{
			Repo:     repo,
			Author:   author,
			versions: p.Versions,
		})
	}

	return plugins
}
