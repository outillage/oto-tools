//+build mage

package main

import (
	"github.com/aevea/magefiles"
)

func Test() error {
	return magefiles.Test()
}

func GoModTidy() error {
	return magefiles.GoModTidy()
}

func Install() error {
	dependencies := []string{
		"github.com/pacedotdev/oto",
	}

	return magefiles.Install(dependencies)
}
