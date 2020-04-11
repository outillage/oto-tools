package publishrunner

import (
	"fmt"
	"io/ioutil"
)

func createConfigFile(path, registry, token string) error {
	contents := fmt.Sprintf("%s:_authToken=%s", registry, token)

	return ioutil.WriteFile("./.npmrc", []byte(contents), 0777)
}
