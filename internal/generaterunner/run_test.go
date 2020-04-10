package generaterunner

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/outillage/oto-tools/internal/npm"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	runner := &Runner{}

	emptyPathErr := runner.Run("", npm.Package{})

	assert.Error(t, emptyPathErr)

	err := runner.Run("./jstest", npm.Package{Name: "test-package", Version: "0.0.1"})

	assert.NoError(t, err)

	file, fileErr := ioutil.ReadFile("./jstest/package.json")

	assert.NoError(t, fileErr)

	generatedPackage := npm.Package{}

	jsonErr := json.Unmarshal(file, &generatedPackage)

	assert.NoError(t, jsonErr)

	assert.Equal(t, "test-package", generatedPackage.Name)
	assert.Equal(t, "0.0.1", generatedPackage.Version)
}
