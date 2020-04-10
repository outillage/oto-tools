package generaterunner

import (
	"errors"
	"os"

	"github.com/magefile/mage/sh"
)

type OtoOptions struct {
	TemplatePath   string
	Ignore         string
	DefinitionPath string
}

func (runner *Runner) generateOtoFiles(path string, options OtoOptions) error {
	var runOptions []string

	if options.TemplatePath == "" {
		return errors.New("missing template path")
	}

	runOptions = append(runOptions, "-template", options.TemplatePath)

	if path == "" {
		return errors.New("missing output path")
	}

	runOptions = append(runOptions, "-out", path+"/client.js")

	if options.DefinitionPath == "" {
		return errors.New("missing definition file path")
	}

	runOptions = append(runOptions, options.DefinitionPath)

	err := os.MkdirAll(path, 0777)

	if err != nil {
		return err
	}

	return sh.RunV("oto", runOptions...)
}
