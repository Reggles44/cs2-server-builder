package plugins

import pluginType "github.com/reggles44/cs2-server-builder/pkg/plugins/types"

var (
	CounterStrikeSharp = &Plugin{
		"CounterStrikeSharp",
		"",
		"https://www.metamodsource.net/downloads.php?branch=stable",
		pluginType.MetaMod,
		[]*Plugin{MetaMod},
	}
	MatchZy = &Plugin{
		"MatchZy",
		"shobhit-pathak",
		"https://github.com/shobhit-pathak/MatchZy",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2WeaponPaints = &Plugin{
		"cs2-WeaponPaints",
		"Nereziel",
		"https://github.com/Nereziel/cs2-WeaponPaints",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2retakes = &Plugin{
		"cs2-retakes",
		"B3none",
		"https://github.com/B3none/cs2-retakes",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2SimpleAdmin = &Plugin{
		"CS2-SimpleAdmin",
		"daffyyyy",
		"https://github.com/daffyyyy/CS2-SimpleAdmin",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cssbans = &Plugin{
		"css-bans",
		"counterstrikesharp-panel",
		"https://github.com/counterstrikesharp-panel/css-bans",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2PlayerModelChanger = &Plugin{
		"CS2-PlayerModelChanger",
		"samyycX",
		"https://github.com/samyycX/CS2-PlayerModelChanger",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Deathmatch = &Plugin{
		"CS2-Deathmatch",
		"NockyCZ",
		"https://github.com/NockyCZ/CS2-Deathmatch",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	AFKManager = &Plugin{
		"AFKManager",
		"NiGHT757",
		"https://github.com/NiGHT757/AFKManager",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2VIPCore = &Plugin{
		"cs2-VIPCore",
		"partiusfabaa",
		"https://github.com/partiusfabaa/cs2-VIPCore",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	FortniteEmotesNDances = &Plugin{
		"FortniteEmotesNDances",
		"Cruze03",
		"https://github.com/Cruze03/FortniteEmotesNDances",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2ranks = &Plugin{
		"cs2-ranks",
		"partiusfabaa",
		"https://github.com/partiusfabaa/cs2-ranks",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2GameManagerGoldKingZ = &Plugin{
		"cs2-Game-Manager-GoldKingZ",
		"oqyh",
		"https://github.com/oqyh/cs2-Game-Manager-GoldKingZ",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2instadefuse = &Plugin{
		"cs2-instadefuse",
		"B3none",
		"https://github.com/B3none/cs2-instadefuse",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2rockthevote = &Plugin{
		"cs2-rockthevote",
		"Oz-Lin",
		"https://github.com/Oz-Lin/cs2-rockthevote",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	PugSharp = &Plugin{
		"PugSharp",
		"Lan2Play",
		"https://github.com/Lan2Play/PugSharp",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Tags = &Plugin{
		"CS2-Tags",
		"daffyyyy",
		"https://github.com/daffyyyy/CS2-Tags",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2tags = &Plugin{
		"cs2-tags",
		"CounterStrike2-Plugins-Archive",
		"https://github.com/CounterStrike2-Plugins-Archive/cs2-tags",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2MenuManager = &Plugin{
		"CS2MenuManager",
		"schwarper",
		"https://github.com/schwarper/CS2MenuManager",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	ResourcePrecacher = &Plugin{
		"ResourcePrecacher",
		"KillStr3aK",
		"https://github.com/KillStr3aK/ResourcePrecacher",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2executes = &Plugin{
		"cs2-executes",
		"zwolof",
		"https://github.com/zwolof/cs2-executes",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2TraceRay = &Plugin{
		"CS2TraceRay",
		"schwarper",
		"https://github.com/schwarper/CS2TraceRay",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Essentials = &Plugin{
		"CS2-Essentials",
		"HvH-gg",
		"https://github.com/HvH-gg/CS2-Essentials",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	EntWatchSharp = &Plugin{
		"EntWatchSharp",
		"darkerz7",
		"https://github.com/darkerz7/EntWatchSharp",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2gungame = &Plugin{
		"cs2-gungame",
		"ssypchenko",
		"https://github.com/ssypchenko/cs2-gungame",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2ChatLoggerGoldKingZ = &Plugin{
		"cs2-Chat-Logger-GoldKingZ",
		"oqyh",
		"https://github.com/oqyh/cs2-Chat-Logger-GoldKingZ",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2ConnectDisconnectSoundGoldKingZ = &Plugin{
		"cs2-Connect-Disconnect-Sound-GoldKingZ",
		"oqyh",
		"https://github.com/oqyh/cs2-Connect-Disconnect-Sound-GoldKingZ",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cssC4Timer = &Plugin{
		"css-C4-Timer",
		"R0mz1k",
		"https://github.com/R0mz1k/css-C4-Timer",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	ThirdPersonWIP = &Plugin{
		"ThirdPerson-WIP",
		"grrhn",
		"https://github.com/grrhn/ThirdPerson-WIP",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2HidePlayers = &Plugin{
		"CS2-HidePlayers",
		"qstage",
		"https://github.com/qstage/CS2-HidePlayers",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	Clientprefs = &Plugin{
		"Clientprefs",
		"Cruze03",
		"https://github.com/Cruze03/Clientprefs",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2updatemanager = &Plugin{
		"cs2-update-manager",
		"Kandru",
		"https://github.com/Kandru/cs2-update-manager",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	AutomaticAdsCS2 = &Plugin{
		"AutomaticAds-CS2",
		"wiruwiru",
		"https://github.com/wiruwiru/AutomaticAds-CS2",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CSSharpPatcher = &Plugin{
		"CSSharpPatcher",
		"samyycX",
		"https://github.com/samyycX/CSSharpPatcher",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_Duel = &Plugin{
		"SLAYER_Duel",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_Duel",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2SkyboxChanger = &Plugin{
		"CS2-SkyboxChanger",
		"samyycX",
		"https://github.com/samyycX/CS2-SkyboxChanger",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2AntiDLL = &Plugin{
		"CS2-AntiDLL",
		"KillStr3aK",
		"https://github.com/KillStr3aK/CS2-AntiDLL",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2FixRandomSpawn = &Plugin{
		"CS2-FixRandomSpawn",
		"qstage",
		"https://github.com/qstage/CS2-FixRandomSpawn",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2rollthedice = &Plugin{
		"cs2-roll-the-dice",
		"Kandru",
		"https://github.com/Kandru/cs2-roll-the-dice",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2quakesounds = &Plugin{
		"cs2-quake-sounds",
		"Kandru",
		"https://github.com/Kandru/cs2-quake-sounds",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	ThirdPersonRevamped = &Plugin{
		"ThirdPerson-Revamped",
		"KKNecmi",
		"https://github.com/KKNecmi/ThirdPerson-Revamped",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	advancedffcs2 = &Plugin{
		"advanced-ff-cs2",
		"mister-keno",
		"https://github.com/mister-keno/advanced-ff-cs2",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	TeleportFix = &Plugin{
		"TeleportFix",
		"HvH-gg",
		"https://github.com/HvH-gg/TeleportFix",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2_blockradiocommands = &Plugin{
		"cs2_blockradiocommands",
		"Cruze03",
		"https://github.com/Cruze03/cs2_blockradiocommands",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2_Speedometer = &Plugin{
		"CS2_Speedometer",
		"PhantomYopta",
		"https://github.com/PhantomYopta/CS2_Speedometer",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Playervotes = &Plugin{
		"CS2-Playervotes",
		"asapverneri",
		"https://github.com/asapverneri/CS2-Playervotes",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2BotAI = &Plugin{
		"CS2-BotAI",
		"Austinbots",
		"https://github.com/Austinbots/CS2-BotAI",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2PoorMapPropAds = &Plugin{
		"CS2-Poor-MapPropAds",
		"Letaryat",
		"https://github.com/Letaryat/CS2-Poor-MapPropAds",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	K4AlwaysWeaponSkins = &Plugin{
		"K4-AlwaysWeaponSkins",
		"K4ryuu",
		"https://github.com/K4ryuu/K4-AlwaysWeaponSkins",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2instaplant = &Plugin{
		"cs2-instaplant",
		"B3none",
		"https://github.com/B3none/cs2-instaplant",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2KillPlugin = &Plugin{
		"CS2-Kill-Plugin",
		"Quantor97",
		"https://github.com/Quantor97/CS2-Kill-Plugin",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	T3MenuAPI = &Plugin{
		"T3Menu-API",
		"T3Marius",
		"https://github.com/T3Marius/T3Menu-API",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_AntiCamp = &Plugin{
		"SLAYER_AntiCamp",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_AntiCamp",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2store = &Plugin{
		"cs2-store",
		"schwarper",
		"https://github.com/schwarper/cs2-store",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	RapidFireFix = &Plugin{
		"RapidFireFix",
		"HvH-gg",
		"https://github.com/HvH-gg/RapidFireFix",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2demorecorder = &Plugin{
		"cs2-demo-recorder",
		"Kandru",
		"https://github.com/Kandru/cs2-demo-recorder",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	MVPAnthem = &Plugin{
		"MVP-Anthem",
		"T3Marius",
		"https://github.com/T3Marius/MVP-Anthem",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2EntityFix = &Plugin{
		"CS2-EntityFix",
		"darkerz7",
		"https://github.com/darkerz7/CS2-EntityFix",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2challenges = &Plugin{
		"cs2-challenges",
		"Kandru",
		"https://github.com/Kandru/cs2-challenges",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_UnrestrictedFOV = &Plugin{
		"SLAYER_UnrestrictedFOV",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_UnrestrictedFOV",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_HeadshotOnly = &Plugin{
		"SLAYER_HeadshotOnly",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_HeadshotOnly",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2HideTeammates = &Plugin{
		"CS2-HideTeammates",
		"darkerz7",
		"https://github.com/darkerz7/CS2-HideTeammates",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CSSharpFixes = &Plugin{
		"CSSharp-Fixes",
		"darkerz7",
		"https://github.com/darkerz7/CSSharp-Fixes",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2TeamLimiter = &Plugin{
		"CS2-TeamLimiter",
		"Ferks-FK",
		"https://github.com/Ferks-FK/CS2-TeamLimiter",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2killfeedfilter = &Plugin{
		"cs2-killfeed-filter",
		"abnerfs",
		"https://github.com/abnerfs/cs2-killfeed-filter",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_Noscope = &Plugin{
		"SLAYER_Noscope",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_Noscope",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_Revive = &Plugin{
		"SLAYER_Revive",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_Revive",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	QuickDefuse = &Plugin{
		"QuickDefuse",
		"Interesting-exe",
		"https://github.com/Interesting-exe/QuickDefuse",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2ChatRelay = &Plugin{
		"CS2-ChatRelay",
		"asapverneri",
		"https://github.com/asapverneri/CS2-ChatRelay",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CSSKnockback = &Plugin{
		"CSS-Knockback",
		"R0mz1k",
		"https://github.com/R0mz1k/CSS-Knockback",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2WeaponRestrict = &Plugin{
		"cs2-WeaponRestrict",
		"TICHOJEBEC-SK",
		"https://github.com/TICHOJEBEC-SK/cs2-WeaponRestrict",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2FlashingHtmlHudFix = &Plugin{
		"CS2FlashingHtmlHudFix",
		"M-archand",
		"https://github.com/M-archand/CS2FlashingHtmlHudFix",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Gunsmenu = &Plugin{
		"CS2-Gunsmenu",
		"asapverneri",
		"https://github.com/asapverneri/CS2-Gunsmenu",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Baninfo = &Plugin{
		"CS2-Baninfo",
		"asapverneri",
		"https://github.com/asapverneri/CS2-Baninfo",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2BotQuotaManager = &Plugin{
		"CS2-BotQuotaManager",
		"asapverneri",
		"https://github.com/asapverneri/CS2-BotQuotaManager",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Legs = &Plugin{
		"CS2-Legs",
		"asapverneri",
		"https://github.com/asapverneri/CS2-Legs",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2Clantags = &Plugin{
		"CS2-Clantags",
		"asapverneri",
		"https://github.com/asapverneri/CS2-Clantags",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	SLAYER_1HitKill = &Plugin{
		"SLAYER_1HitKill",
		"zakriamansoor47",
		"https://github.com/zakriamansoor47/SLAYER_1HitKill",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2BotSlay = &Plugin{
		"CS2-BotSlay",
		"Dliix66",
		"https://github.com/Dliix66/CS2-BotSlay",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2ChangeMapRotationUsingTextFile = &Plugin{
		"cs2-Change-Map-Rotation-Using-Text-File",
		"Austinbots",
		"https://github.com/Austinbots/cs2-Change-Map-Rotation-Using-Text-File",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2ScoutzKnivez = &Plugin{
		"CS2-ScoutzKnivez",
		"asapverneri",
		"https://github.com/asapverneri/CS2-ScoutzKnivez",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2StaffList = &Plugin{
		"CS2-StaffList",
		"asapverneri",
		"https://github.com/asapverneri/CS2-StaffList",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2BotsNoKnife = &Plugin{
		"cs2-BotsNoKnife",
		"Austinbots",
		"https://github.com/Austinbots/cs2-BotsNoKnife",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2WarnSystem = &Plugin{
		"cs2-WarnSystem",
		"TICHOJEBEC-SK",
		"https://github.com/TICHOJEBEC-SK/cs2-WarnSystem",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	CS2ChatLogs = &Plugin{
		"CS2-ChatLogs",
		"asapverneri",
		"https://github.com/asapverneri/CS2-ChatLogs",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
	cs2anticheat = &Plugin{
		"cs2-anticheat",
		"schwarper",
		"https://github.com/schwarper/cs2-anticheat",
		pluginType.CounterStrikeSharp,
		[]*Plugin{CounterStrikeSharp},
	}
)
