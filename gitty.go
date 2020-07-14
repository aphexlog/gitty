package main

import (
	"os"

	"gopkg.in/src-d/go-git.v4"
)

// GitClone the given repository to the given directory
func GitClone(srcURL string, dstDir string) {

	Info("git clone " + srcURL)
	_, err := git.PlainClone(dstDir, false, &git.CloneOptions{
		URL:      srcURL,
		Progress: os.Stdout,
	})

	CheckIfError(err)
}
