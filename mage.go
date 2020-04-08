//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Test() error {
	return sh.Run("go", "test", "./...", "-v", "-race")
}

func GoModTidy() error {
	err := sh.Run("go", "mod", "tidy", "-v")
	if err != nil {
		return err
	}
	return sh.Run("git", "diff-index", "--quiet", "HEAD")
}
