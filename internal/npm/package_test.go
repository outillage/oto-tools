package npm

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {
	_, nameError := GeneratePackageInfo(Package{})

	assert.Error(t, nameError)

	_, versionError := GeneratePackageInfo(Package{Name: "test-package"})

	assert.Error(t, versionError)

	generated, err := GeneratePackageInfo(Package{"test-package", "0.0.1"})

	assert.NoError(t, err)

	file, err := ioutil.ReadFile("./testdata/package.json")

	assert.NoError(t, err)
	assert.Equal(t, file, generated)
}
