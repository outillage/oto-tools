package generaterunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOto(t *testing.T) {
	runner := Runner{}

	missingTemplate := runner.generateOtoFiles("./jstest", OtoOptions{})

	assert.Error(t, missingTemplate)

	missingDefinitions := runner.generateOtoFiles("./jstest", OtoOptions{TemplatePath: "./testdata/template/client.js.plush"})

	assert.Error(t, missingDefinitions)

	err := runner.generateOtoFiles("./jstest", OtoOptions{TemplatePath: "./testdata/template/client.js.plush", DefinitionPath: "./testdata/definitions"})

	assert.NoError(t, err)
	assert.FileExists(t, "./jstest/client.js")
}
