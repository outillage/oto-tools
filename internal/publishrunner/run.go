package publishrunner

import (
	"log"

	"github.com/aevea/oto-tools/internal/npm"
)

// Run executes the full publish command. It will create a .npmrc file, publish to NPM and then delete the .npmrc file.
func (runner *Runner) Run(path string, options npm.PublishOptions) error {
	var registry string

	if options.RegistryURL == "" {
		registry = npm.NPMRegistry
	}

	if options.RegistryURL == "npm" {
		registry = npm.NPMRegistry
	}

	if options.RegistryURL == "github" {
		registry = npm.GHRegistry
	}

	err := createConfigFile(path, registry, options.Token, options.Owner)

	if err != nil {
		return err
	}

	err = npm.Publish(path, options)

	if err != nil {
		log.Printf("Failed to publish to npm %v", err)
	}

	cleanup()

	return err
}
