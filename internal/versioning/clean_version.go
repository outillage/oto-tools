package versioning

import "strings"

func cleanVersion(version string) string {
	version = strings.TrimLeft(version, "refs/tag")

	version = strings.TrimLeft(version, "v")

	return version
}
