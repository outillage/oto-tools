package generaterunner

import (
	"testing"

	"github.com/aevea/oto-tools/internal/npm"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	runner := Runner{}

	pathErr := runner.Run("", npm.Package{}, OtoOptions{})

	assert.Error(t, pathErr)
}
