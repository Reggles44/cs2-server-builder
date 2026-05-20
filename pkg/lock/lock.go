package lock

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/reggles44/cs2-server-builder/pkg/plugins"
)

type Lock struct {
	Plugins []PluginLock `json:"plugins"`
}

type PluginLock struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Version string `json:"version"`
}

func getPluginPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "Plugins"), nil
}

func ReadLock() (*Lock, error) {
	var lock Lock

	path, err := getPluginPath()
	if err != nil {
		return &lock, err
	}

	dat, err := os.ReadFile(path)
	if err != nil {
		return &lock, err
	}

	err = json.Unmarshal(dat, &lock)
	if err != nil {
		return &lock, err
	}

	return &lock, nil
}

func (l *Lock) WriteLock() error {
	path, err := getPluginPath()
	if err != nil {
		return err
	}

	dat, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(path, dat, os.ModePerm)
}

func (l *Lock) AddPlugin(plugin plugins.Plugin, version string) {
	l.Plugins = append(l.Plugins, PluginLock{plugin.Name, plugin.Repo, version})
}

func (pl *PluginLock) GetPlugin() (*plugins.Plugin, bool) {
	p, ok := plugins.Map[pl.Name]
	return p, ok
}
