package versioning

import "strings"

func cleanVersion(version string) string {
	version = strings.TrimPrefix(version, "/refs/tags/")

	version = strings.TrimLeft(version, "v")

	return version
}
