package provider

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_NameFromDownloadUrl(t *testing.T) {
	tests := []struct {
		url      string
		fileName string
	}{
		{url: "", fileName: ""},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example", fileName: "example"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example.zip", fileName: "example.zip"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example.tar.gz", fileName: "example.tar.gz"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example.sh", fileName: "example.sh"},

		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1/example-with_versions", fileName: "example-with_versions"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1/example-with_versions.zip", fileName: "example-with_versions.zip"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1/example-with_versions.tar.gz", fileName: "example-with_versions.tar.gz"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1/example-with_versions.sh", fileName: "example-with_versions.sh"},

		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1-beta/example-with_versions", fileName: "example-with_versions"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1-beta/example-with_versions.zip", fileName: "example-with_versions.zip"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1-beta/example-with_versions.tar.gz", fileName: "example-with_versions.tar.gz"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/v0.0.1-beta/example-with_versions.sh", fileName: "example-with_versions.sh"},

		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example-with_different_dashes", fileName: "example-with_different_dashes"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example-with_different_dashes.zip", fileName: "example-with_different_dashes.zip"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example-with_different_dashes.tar.gz", fileName: "example-with_different_dashes.tar.gz"},
		{url: "https://github.com/sklemmer/go-autoupdate-example/releases/download/0.0.1/example-with_different_dashes.sh", fileName: "example-with_different_dashes.sh"},
	}
	for _, test := range tests {
		result := nameFromDownloadUrl(test.url)
		assert.Equal(t, test.fileName, result)
		fmt.Println("Pass: ", result)
	}
}
