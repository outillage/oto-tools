package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanVersion(t *testing.T) {
	tests := map[string]string{
		"/refs/tags/v0.0.1": "0.0.1",
		"refs/tags/v0.0.1":  "0.0.1",
		"v0.0.2":            "0.0.2",
		"0.0.3":             "0.0.3",
	}

	for test, expected := range tests {
		assert.Equal(t, expected, cleanVersion(test))
	}

}
