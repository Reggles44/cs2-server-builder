package lock

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/reggles44/cs2-server-builder/pkg/plugins"
)

type Lock map[string]*PluginLock

type PluginLock struct {
	Repo    string           `json:"repo"`
	Author  string           `json:"author"`
	Version *plugins.Version `json:"version"`
}

var path = getLockFilePath()

func getLockFilePath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	return filepath.Join(dir, "Plugins.json")
}

func ReadLock() *Lock {
	lock := make(Lock)

	dat, err := os.ReadFile(path)
	if err != nil {
		return &lock
	}

	err = json.Unmarshal(dat, &lock)
	if err != nil {
		log.Panic(err)
	}

	return &lock
}

func (l *Lock) WriteLock() error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	dat, err := json.MarshalIndent(l, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(dat)
	return err
}

func (l Lock) AddPlugin(plugin *plugins.Plugin, version *plugins.Version) {
	l[plugin.Key()] = &PluginLock{
		Repo:    plugin.Repo,
		Author:  plugin.Author,
		Version: version,
	}
}

func (l Lock) RemovePlugin(plugin *plugins.Plugin) {
	delete(l, plugin.Key())
}
