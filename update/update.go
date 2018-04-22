package update

import (
	"path/filepath"
	"os"
	"fmt"
)

func (u *Updater) Update() (error) {
	fmt.Println("Download newer version")
	err := u.provider.GetBinary(u.release)
	if err != nil {
		return err
	}

	ex, err := os.Executable()
	if err != nil {
		return err
	}
	currentLocation := filepath.Dir(ex)

	if err := moveAsset(u.release.Location, currentLocation); err != nil {
		return err
	}
	return nil
}

func moveAsset(downloaded, current string) (error) {
	return nil
}
