package provider

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewGithubOptions(t *testing.T) {
	tests := []struct {
		repoStr string
		owner   string
		repo    string
	}{
		{"sklemmer/go-autoupdate", "sklemmer", "go-autoupdate"},
		{"sklemmer/go-autoupdate-example", "sklemmer", "go-autoupdate-example"},
	}

	for _, test := range tests {
		result := NewGithubOptions(test.repoStr)
		assert.Equal(t, test.owner, result.Owner)
		assert.Equal(t, test.repo, result.Repo)
	}
}

func TestGithubProvider_GetLatestRelease(t *testing.T) {
	expected := &Release{Version: "0.0.1", CommitID: 10547150}

	github := NewGithubProvider(&GithubOptions{GitOptions: &GitOptions{Owner: "sklemmer", Repo: "go-autoupdate-example"}})
	release, err := github.GetLatestRelease()
	assert.NoError(t, err)
	assert.Equal(t, expected, release)
}

func TestGithubProvider_GetBinary(t *testing.T) {
	expected := &Release{Version: "0.0.1", CommitID: 10547150}

	github := NewGithubProvider(&GithubOptions{GitOptions: &GitOptions{Owner: "sklemmer", Repo: "go-autoupdate-example"}})
	release, err := github.GetLatestRelease()
	assert.NoError(t, err)
	assert.Equal(t, expected, release)
	//TODO: download and verify
}

func TestGithubProvider_GetBinary_NoDownloads(t *testing.T) {
	expected := &Release{Version: "0.0.1", CommitID: 10547150}

	github := NewGithubProvider(&GithubOptions{GitOptions: &GitOptions{Owner: "sklemmer", Repo: "go-autoupdate-example"}})
	_, err := github.GetBinary(expected)

	assert.Error(t, err, errNoDownloadUrlFound)
}
