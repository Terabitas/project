package version

import (
	"github.com/coreos/go-semver/semver"
)

// Version of the service.
const Version = "$VERSION"

// SemVersion (Semantic Version) of the service.
var SemVersion semver.Version

func init() {
	sv, err := semver.NewVersion(Version)
	if err != nil {
		panic("bad version string!")
	}
	SemVersion = *sv
}
