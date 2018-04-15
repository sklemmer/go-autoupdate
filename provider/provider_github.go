package provider

import (
	"github.com/google/go-github/github"
	"context"
	"net/http"
	"strings"
)

var client *github.Client

func getClient(options *GithubOptions) (*github.Client) {
	if client == nil {
		client = github.NewClient(options.HttpClient)
	}
	return client
}

type GithubOptions struct {
	*GitOptions
	HttpClient *http.Client
}

type GithubProvider struct {
	options *GithubOptions
}

func NewGithubOptions(repoStr string) (*GithubOptions) {
	tokens := strings.Split(repoStr, "/")
	return &GithubOptions{GitOptions: &GitOptions{Owner: tokens[0], Repo: tokens[1]}}
}

func NewGithubProvider(options *GithubOptions) (*GithubProvider) {
	if options.ctx == nil {
		options.ctx = context.Background()
	}
	return &GithubProvider{options}
}

func (gp *GithubProvider) GetLatestRelease() (*Release, error) {
	rr, _, err := getClient(gp.options).Repositories.GetLatestRelease(gp.options.ctx, gp.options.Owner, gp.options.Repo)
	if err != nil {
		return nil, err
	}
	return &Release{
		Version:  rr.GetTagName(),
		CommitID: rr.GetID(),
	}, nil
}

func (gp *GithubProvider) GetBinary(release *Release) (*Release, error) {
	ra, _, _ := getClient(gp.options).Repositories.GetReleaseAsset(gp.options.ctx, gp.options.Owner, gp.options.Repo, release.CommitID)
	release.DownloadUrl = ra.GetBrowserDownloadURL()
	return downloadBinary(release)
}
