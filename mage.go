//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Test() error {
	return sh.RunV("go", "test", "./...", "-v", "-race")
}

func GoModTidy() error {
	err := sh.RunV("go", "mod", "tidy", "-v")
	if err != nil {
		return err
	}
	return sh.RunV("git", "diff-index", "--quiet", "HEAD")
}

func Install() error {
	return sh.RunV("go", "install", "github.com/pacedotdev/oto")
}
