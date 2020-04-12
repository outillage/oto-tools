package publishrunner

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/outillage/oto-tools/internal/npm"
)

func createConfigFile(path, registry, token, owner string) error {
	contents := fmt.Sprintf("%s:_authToken=%s", registry, token)

	if registry == npm.GHRegistry {
		if owner == "" {
			return errors.New("github package registry requires an owner")
		}

		contents += fmt.Sprintf("\nregistry=https://npm.pkg.github.com/%s", owner)
	}

	return ioutil.WriteFile("./.npmrc", []byte(contents), 0777)
}
