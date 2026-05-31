#!/usr/bin/env python
#####################################################
# Build using https://github.com/samyycX/awesome-cs2
#####################################################

import json
import logging
import time
import os
import requests
import dataclasses
import datetime


PLUGIN_FILE = "pkg/plugins/plugins.json"
PLUGIN_BACKUP_FILE = "pkg/plugins/plugins.backup.json"
AWESOME_CS2_GITHUB_URL = (
    "https://raw.githubusercontent.com/samyycX/awesome-cs2/refs/heads/master/data/{}"
)
RELEASES_GITHUB_URL = "https://api.github.com/repos/{}/releases"
REQUESTS_AUTH = os.getenv("GH_TOKEN")
REQUESTS_HEADERS = {
    "Authorization": f"Bearer {REQUESTS_AUTH}",
}


logging.basicConfig()
logging.getLogger().setLevel(logging.DEBUG)
requests_log = logging.getLogger("requests.packages.urllib3")
requests_log.setLevel(logging.DEBUG)
requests_log.propagate = True


def get_plugin_list(list: str, deps: list[str] = []) -> list[str]:
    list_resp = requests.get(
        AWESOME_CS2_GITHUB_URL.format(list), headers=REQUESTS_HEADERS
    )
    list_resp.raise_for_status()
    return [plug["repo"] for plug in list_resp.json()]


@dataclasses.dataclass()
class Version:
    url: str
    id: int
    tag: str
    created_at: datetime.datetime
    tarball_url: str
    zipball_url: str


def get_version_list(repo: str) -> list[Version]:
    attempt = 0
    while attempt < 5:
        resp = requests.get(RELEASES_GITHUB_URL.format(repo), headers=REQUESTS_HEADERS)
        if resp.status_code == 403:
            time.sleep(5)
            attempt += 1
            continue
        elif resp.status_code == 404:
            break
        else:
            return [
                Version(
                    url=obj["url"],
                    id=obj["id"],
                    tag=obj["tag_name"],
                    created_at=obj["created_at"],
                    tarball_url=obj["tarball_url"],
                    zipball_url=obj["zipball_url"],
                )
                for obj in resp.json()
            ]
    return []


@dataclasses.dataclass()
class Plugin:
    repo: str
    releases: list[Version]


def main():

    if os.path.exists(PLUGIN_FILE):
        os.rename(PLUGIN_FILE, PLUGIN_BACKUP_FILE)

    plugins = [
        "alliedmodders/metamod-source",
        "roflmuffin/CounterStrikeSharp",
        "swiftly-solution/swiftlys2",
        "Kxnrl/modsharp-public",
        "untrustedmodders/plugify",
        *get_plugin_list("metamod-plugins.stats.json"),
        *get_plugin_list("counterstrikesharp-plugins.stats.json"),
        *get_plugin_list("swiftly-plugins.stats.json"),
        *get_plugin_list("swiftlys2-plugins.stats.json"),
    ]

    final = []
    for p in plugins:
        print("\n\n", p)
        plugin = Plugin(repo=p, releases=get_version_list(p))
        final.append(dataclasses.asdict(plugin))

    with open(PLUGIN_FILE, "+w") as plugin_file:
        plugin_file.write(json.dumps(final, indent=4))

    if os.path.exists(PLUGIN_BACKUP_FILE):
        os.remove(PLUGIN_BACKUP_FILE)


if __name__ == "__main__":
    main()
