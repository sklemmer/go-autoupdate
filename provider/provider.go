package provider

import "context"

type Release struct {
	Version     string
	CommitID    int64
	Location    string
	DownloadUrl string
}

type UpdateProvider interface {
	GetLatestRelease() (*Release, error)
	GetBinary(*Release) (error)
}

type GenericOptions struct {
	ctx context.Context
}

type GitOptions struct {
	*GenericOptions
	Owner string
	Repo  string
}
