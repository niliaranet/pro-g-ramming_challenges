package models

import (
	"strings"

	"carrega/models/errmodels"
)

type DownloadOptions struct {
	Url      string
	FileName string
	Progress int
}

func (ops *DownloadOptions) From(url string) *DownloadOptions {
	if url == "" {
		return ops
	}

	fileName, err := getUrlName(url)
	if err != nil {
		return ops
	}

	ops.Url = url
	ops.FileName = fileName
	return ops
}

func getUrlName(url string) (string, error) {
	urlParts := strings.Split(url, "/")
	name := ""
	for i := len(urlParts) - 1; i >= 0; i-- {
		name = urlParts[i]
		if name != "" {
			return name, nil
		}
	}

	return "", errmodels.TooFewArguments
}
