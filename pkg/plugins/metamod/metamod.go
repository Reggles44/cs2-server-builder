package metamod

import "github.com/reggles44/cs2-server-builder/pkg/plugins"

var (
	MetaModPlugin = plugins.GithubPlugin{
		RepoName:   "metamod-source",
		RepoAuthor: "alliedmodders",
		Plugin: plugins.Plugin{
			Name:         "MetaMod",
			Dependencies: []plugins.PluginType{},
		},
	}
	CS2FixesPlugin = plugins.GithubPlugin{
		RepoName:   "CS2Fixes",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "CS2Fixes",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	MultiAddonManagerPlugin = plugins.GithubPlugin{
		RepoName:   "MultiAddonManager",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "MultiAddonManager",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	AcceleratorCS2Plugin = plugins.GithubPlugin{
		RepoName:   "AcceleratorCS2",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "AcceleratorCS2",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	CS2ServerGUIPlugin = plugins.GithubPlugin{
		RepoName:   "CS2ServerGUI",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "CS2ServerGUI",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	Source2SchemaDumperPlugin = plugins.GithubPlugin{
		RepoName:   "Source2SchemaDumper",
		RepoAuthor: "GAMMACASE",
		Plugin: plugins.Plugin{
			Name:         "Source2SchemaDumper",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	CleanerCS2Plugin = plugins.GithubPlugin{
		RepoName:   "CleanerCS2",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "CleanerCS2",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	StripperCS2Plugin = plugins.GithubPlugin{
		RepoName:   "StripperCS2",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "StripperCS2",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	MovementUnlockerPlugin = plugins.GithubPlugin{
		RepoName:   "MovementUnlocker",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "MovementUnlocker",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	ServerListPlayersFixPlugin = plugins.GithubPlugin{
		RepoName:   "ServerListPlayersFix",
		RepoAuthor: "Source2ZE",
		Plugin: plugins.Plugin{
			Name:         "ServerListPlayersFix",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	AudioPlugin = plugins.GithubPlugin{
		RepoName:   "Audio",
		RepoAuthor: "samyycX",
		Plugin: plugins.Plugin{
			Name:         "Audio",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	GameBanFixPlugin = plugins.GithubPlugin{
		RepoName:   "GameBanFix",
		RepoAuthor: "Cruze03",
		Plugin: plugins.Plugin{
			Name:         "GameBanFix",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	FakeRanksRevealAllPlugin = plugins.GithubPlugin{
		RepoName:   "FakeRanks-RevealAll",
		RepoAuthor: "Cruze03",
		Plugin: plugins.Plugin{
			Name:         "FakeRanks-RevealAll",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	cs2kzmetamodPlugin = plugins.GithubPlugin{
		RepoName:   "cs2kz-metamod",
		RepoAuthor: "KZGlobalTeam",
		Plugin: plugins.Plugin{
			Name:         "cs2kz-metamod",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	CS2FixesRampbugFixPlugin = plugins.GithubPlugin{
		RepoName:   "CS2Fixes-RampbugFix",
		RepoAuthor: "Interesting-exe",
		Plugin: plugins.Plugin{
			Name:         "CS2Fixes-RampbugFix",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
	cs2fakerconPlugin = plugins.GithubPlugin{
		RepoName:   "cs2-fake-rcon",
		RepoAuthor: "Salvatore-Als",
		Plugin: plugins.Plugin{
			Name:         "cs2-fake-rcon",
			Dependencies: []plugins.PluginType{&MetaModPlugin},
		},
	}
)
