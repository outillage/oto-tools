package npm

import (
	"errors"

	"github.com/magefile/mage/sh"
)

type PublishOptions struct {
	Token       string
	DryRun      bool
	Private     bool
	RegistryURL string
	// Owner is used only for Github
	Owner string
}

func Publish(path string, options PublishOptions) error {
	var command []string

	command = append(command, "publish", path)

	if options.Token == "" {
		return errors.New("missing NPM token")
	}

	if options.DryRun {
		command = append(command, "--dry-run")
	}

	env := map[string]string{
		"NPM_TOKEN": options.Token,
	}

	if !options.Private {
		command = append(command, "--access", "public")
	}

	return sh.RunWith(env, "npm", command...)
}
