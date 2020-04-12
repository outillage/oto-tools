package versioning

import (
	"fmt"
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

	if err == nil {
		return cleanVersion(currentTag.Name)
	}

	if err != nil && err != git.ErrCommitNotOnTag {
		return ""
	}

	currentCommit, err := gitRepo.CurrentCommit()

	if err != nil {
		return ""
	}

	previousTag, err := gitRepo.PreviousTag(currentCommit.Hash)

	if err != nil {
		return fmt.Sprintf("0.0.0-%s", currentCommit.Hash)
	}

	return fmt.Sprintf("%s-%s", cleanVersion(previousTag.Name), currentCommit.Hash)
}
