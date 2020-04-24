package generaterunner

import (
	"io/ioutil"
	"os"

	"github.com/aevea/oto-tools/internal/npm"
)

// PackageJSON creates and saves the package.json file for uploading to npm
func (runner *Runner) packageJSON(path string, packageOptions npm.Package) error {
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
