package publishrunner

import (
	"log"

	"github.com/outillage/oto-tools/internal/npm"
)

const (
	// NPMRegistry is the default registry url for most JS projects.
	NPMRegistry = "//registry.npmjs.org/"
)

// Run executes the full publish command. It will create a .npmrc file, publish to NPM and then delete the .npmrc file.
func (runner *Runner) Run(path string, options npm.PublishOptions) error {
	err := createConfigFile(path, NPMRegistry, options.Token)

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
