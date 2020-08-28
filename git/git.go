package git

import (
	"github.com/go-git/go-git/v5"
)

//IsGitRepository returns true if the current working directory is a Git repository
func IsGitRepository() bool {
	dir, err := os.Getwd()
	if err != nil {
		return false
	}

	_, err = git.PlainOpen(dir)
	if err != nil {
		return false
	}

	return true
}