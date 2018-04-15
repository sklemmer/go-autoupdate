package provider

import (
	"github.com/google/go-github/github"
	"net/http"
	"strings"
	"context"
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
	//if options != nil && options.GitOptions != nil && options.GitOptions.ctx == nil {
	//	options.GitOptions.ctx = context.TODO()
	//}
	return &GithubProvider{options}
}

func (gp *GithubProvider) GetLatestRelease() (*Release, error) {
	rr, _, err := getClient(gp.options).Repositories.GetLatestRelease(context.TODO(), gp.options.Owner, gp.options.Repo)
	if err != nil {
		return nil, err
	}
	return &Release{
		Version:  rr.GetTagName(),
		CommitID: rr.GetID(),
	}, nil
}

func (gp *GithubProvider) GetBinary(release *Release) (*Release, error) {
	ra, _, _ := getClient(gp.options).Repositories.GetReleaseAsset(context.TODO(), gp.options.Owner, gp.options.Repo, release.CommitID)
	release.DownloadUrl = ra.GetBrowserDownloadURL()
	return downloadBinary(release)
}
