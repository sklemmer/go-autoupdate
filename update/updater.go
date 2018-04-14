package update

import (
	"github.com/sklemmer/go-autoupdate/provider"
)

type Updater struct {
	provider provider.UpdateProvider
	Version  string
	release  *provider.Release
}

func NewUpdater(updateProvider provider.UpdateProvider, version string) (*Updater) {
	return &Updater{provider: updateProvider, Version: version}
}
