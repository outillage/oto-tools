package versioning

import (
	"io/ioutil"
	"log"

	"github.com/outillage/git/v2"
)

// FindVersion tries to find the version from Git. Non-tagged commits get version next.
func FindVersion() string {
	gitRepo, err := git.OpenGit(".", log.New(ioutil.Discard, "", 0))

	if err != nil {
		return ""
	}

	currentTag, err := gitRepo.CurrentTag()

	if err == git.ErrCommitNotOnTag {
		return "next"
	}
	if err != nil {
		return ""
	}

	return currentTag.Name
}
