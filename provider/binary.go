package provider

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io"
	"errors"
)

const (
	errOsCreateText = "error while creating %s\n%s"
	errDownloadText = "error while downloading %s\n%s"
)

var (
	errNoDownloadUrlFound = errors.New("could not find any downloads")
)

func downloadBinary(release *Release) (*Release, error) {
	if release.DownloadUrl == "" {
		return nil, errNoDownloadUrlFound
	}

	fileName := nameFromDownloadUrl(release.DownloadUrl)
	fmt.Println("Downloading", release.DownloadUrl, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf(errOsCreateText, fileName, err)
	}
	defer output.Close()

	response, err := http.Get(release.DownloadUrl)
	if err != nil {
		return nil, fmt.Errorf(errDownloadText, release.DownloadUrl, err)
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return nil, fmt.Errorf(errDownloadText, release.DownloadUrl, err)
	}

	release.Location = fileName

	fmt.Println(n, "bytes downloaded.")
	return release, nil
}

func nameFromDownloadUrl(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1];
}
