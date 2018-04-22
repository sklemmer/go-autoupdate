package update

import (
	"github.com/blang/semver"
	"fmt"
)

/**
checks for new releases in the repository
returns true if a new release is found

TODO: add ability to check for updates in gitlab, bitbucket, github and binary
 */
func (u *Updater) Check() (bool, error) {
	fmt.Println("Checking for updates")
	release, err := u.provider.GetLatestRelease()
	if err != nil {
		return false, err
	}

	if isNewer(release.Version, u.Version) {
		fmt.Printf("Found new version %s, old version %s", release.Version, u.Version)
		u.release = release
		return true, nil
	}

	fmt.Printf("Already using latest version %s", release.Version)
	return false, nil
}

func isNewer(v1 string, v2 string) (bool) {
	sem1, _ := semver.Make(v1)
	sem2, _ := semver.Make(v2)

	return sem1.GT(sem2)
}
