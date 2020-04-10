package generaterunner

import (
	"errors"

	"github.com/outillage/oto-tools/internal/npm"
)

func (runner *Runner) Run(path string, packageOptions npm.Package) error {
	if path == "" {
		return errors.New("path can't be empty")
	}

	err := runner.packageJSON(path, packageOptions)

	if err != nil {
		return err
	}

	return nil
}
