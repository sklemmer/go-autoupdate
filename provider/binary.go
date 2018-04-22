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
	errOsOpenText   = "error while opening %s\n%s"
	errDownloadText = "error while downloading %s\n%s"
)

var (
	errNoDownloadUrlFound = errors.New("could not find any downloads")
)

func downloadBinary(release *Release) error {
	if release.DownloadUrl == "" {
		return errNoDownloadUrlFound
	}

	fileName := nameFromDownloadUrl(release.DownloadUrl)
	fmt.Println("Downloading", release.DownloadUrl, "to", fileName)

	var output *os.File
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if output, err = os.Create(fileName); err != nil {
			return fmt.Errorf(errOsCreateText, fileName, err)
		}
	} else {
		if output, err = os.Open(fileName); err != nil {
			return fmt.Errorf(errOsOpenText, fileName, err)
		}
	}
	defer output.Close()

	response, err := http.Get(release.DownloadUrl)
	if err != nil {
		return fmt.Errorf(errDownloadText, release.DownloadUrl, err)
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf(errDownloadText, release.DownloadUrl, err)
	}

	release.Location = fileName

	fmt.Println(n, "bytes downloaded.")
	return nil
}

func nameFromDownloadUrl(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1];
}
