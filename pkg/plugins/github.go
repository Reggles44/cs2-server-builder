package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubPlugin struct {
	Plugin
	RepoName   string
	RepoAuthor string
}

type githubReleases struct {
	URL        string `json:"url"`
	ID         int    `json:"id"`
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
	PreRelease bool   `json:"prerelease"`
	CreatedAt  string `json:"created_at"`
	TarBallUrl string `json:"tarball_url"`
	ZipBallUrl string `json:"zipball_url"`
}

func (p *GithubPlugin) Version() ([]Version, error) {
	versions := []Version{}

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", p.RepoAuthor, p.RepoName))
	if err != nil {
		return versions, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return versions, err
	}

	releases := []githubReleases{}
	err = json.Unmarshal(dat, &releases)
	if err != nil {
		return versions, err
	}

	for _, r := range releases {
		versions = append(versions, Version{
			Tag:         r.TagName,
			CreatedDate: r.CreatedAt,
			TarBallUrl:  r.TarBallUrl,
			ZipBallUrl:  r.ZipBallUrl,
		})
	}

	return versions, nil
}
