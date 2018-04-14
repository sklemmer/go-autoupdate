package update

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sklemmer/go-autoupdate/provider"
	"errors"
)

var errFake = errors.New("fake error while getting latest release")

type fakeWithError struct{}

func (*fakeWithError) GetLatestRelease() (*provider.Release, error) {
	return nil, errFake
}

func (*fakeWithError) GetBinary(*provider.Release) (error) {
	panic("implement me")
}

type fakeProvider struct{}

func (*fakeProvider) GetLatestRelease() (*provider.Release, error) {
	return &provider.Release{
		Version: "0.0.2",
	}, nil
}

func (*fakeProvider) GetBinary(*provider.Release) (error) {
	panic("implement me")
}

func TestUpdater_Check_IsNewer(t *testing.T) {
	updater := NewUpdater(&fakeProvider{}, "0.0.1")
	ok, err := updater.Check()
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestUpdater_Check_WithError(t *testing.T) {
	updater := NewUpdater(&fakeWithError{}, "0.0.1")
	ok, err := updater.Check()
	assert.Error(t, err)
	assert.EqualError(t, err, errFake.Error())
	assert.False(t, ok)
}

func TestUpdater_CheckIsNotNewer(t *testing.T) {
	updater := NewUpdater(&fakeProvider{}, "0.0.3")
	ok, err := updater.Check()
	assert.NoError(t, err)
	assert.False(t, ok)
}
