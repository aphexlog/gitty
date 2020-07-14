package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

// GitClone the given repository to the given directory
func GitClone(srcURL string, dstDir string) {

	fmt.Println("git clone " + srcURL)
	_, err := git.PlainClone(dstDir, false, &git.CloneOptions{
		URL:      srcURL,
		Progress: os.Stdout,
	})

	CheckIfError(err)
}
