package generaterunner

import (
	"errors"
	"strings"

	"github.com/outillage/oto-tools/internal/npm"
)

func (runner *Runner) Run(path string, packageOptions npm.Package, otoOptions OtoOptions) error {
	if path == "" {
		return errors.New("path can't be empty")
	}
	path = strings.TrimRight(path, "/")

	err := runner.packageJSON(path, packageOptions)

	if err != nil {
		return err
	}

	err = runner.generateOtoFiles(path, otoOptions)

	if err != nil {
		return err
	}

	return nil
}
