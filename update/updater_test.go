package update

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUpdater(t *testing.T) {
	updater := NewUpdater(&fakeProvider{}, "v0.0.1")
	ok, err := updater.Check()
	assert.True(t, ok)
	assert.NoError(t, err)
}
