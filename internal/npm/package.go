package npm

import (
	"encoding/json"
	"errors"
)

// Package contains all the fields required for creating an npm package
type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Main    string `json:"main"`
}

// GeneratePackageInfo validates that all Package fields are present and returns the marshalled bytes of the package json
func GeneratePackageInfo(npmPackage Package) ([]byte, error) {
	emptyReturn := []byte("")

	if npmPackage.Name == "" {
		return emptyReturn, errors.New("missing npm package name")
	}

	if npmPackage.Version == "" {
		return emptyReturn, errors.New("missing npm package version")
	}

	if npmPackage.Main == "" {
		npmPackage.Main = NPMPackageName
	}

	marshalled, err := json.Marshal(npmPackage)

	if err != nil {
		return emptyReturn, err
	}

	return marshalled, nil
}
