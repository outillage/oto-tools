package generaterunner

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/outillage/oto-tools/internal/npm"
)

func (runner *Runner) Run(path string, packageOptions npm.Package) error {
	if path == "" {
		return errors.New("path can't be empty")
	}

	npmPackage, err := npm.GeneratePackageInfo(packageOptions)

	if err != nil {
		return err
	}

	err = os.MkdirAll(path, 0777)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path+"/package.json", npmPackage, 0777)

	if err != nil {
		return err
	}

	return nil
}
