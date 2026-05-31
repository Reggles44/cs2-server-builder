package plugins

import (
	"fmt"
)

type Version struct {
	Tag         string `json:"tag"`
	CreatedDate string `json:"created_date"`
	TarBallURL  string `json:"tarball_url"`
	ZipBallURL  string `json:"zipball_url"`
}

func (v *Version) String() string {
	return fmt.Sprintf("%s", v.Tag)
}

type VersionSlice []*Version

func (vs VersionSlice) Latest() *Version {
	if len(vs) == 0 {
		return &Version{}
	}

	return vs[0]
}
