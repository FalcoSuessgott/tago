package git

import (
	"log"
	"strings"
	"time"

	"github.com/tcnksm/go-gitconfig"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/integralist/go-findroot/find"
)

// GetRootDir returns the root directory of the git repository
func GetRootDir() (string, error) {
	dir, err := find.Repo()
	return dir.Path, err

}

// IsGitRepository returns true if the directory is a Git repository
func IsGitRepository(dir string) bool {
	_, err := git.PlainOpen(dir)

	return err == nil
}

// GetTags returns a list of git tags
func GetTags(r *git.Repository) []string {
	tags := []string{}
	tagRefs, _ := r.Tags()

	err := tagRefs.ForEach(func(t *plumbing.Reference) error {
		tags = append(tags, strings.Split(t.String(), "/")[2])
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return tags
}

func getUsername() string {
	username, err := gitconfig.Username()
	if err != nil {
		log.Fatal(err)
	}

	return username
}

func getEmail() string {
	email, err := gitconfig.Email()
	if err != nil {
		log.Fatal(err)
	}

	return email
}

func defaultSignature() *object.Signature {
	return &object.Signature{
		Name:  getUsername(),
		Email: getEmail(),
		When:  time.Now(),
	}
}

// Repo returns the repository handle
func Repo(dir string) (*git.Repository, error) {
	return git.PlainOpen(dir)

}

// AddTag adds a tag
func AddTag(r *git.Repository, tag, message string) error {
	opts := &git.CreateTagOptions{
		Tagger:  defaultSignature(),
		Message: message,
	}

	head, err := r.Head()
	if err != nil {
		return err
	}

	_, err = r.CreateTag(tag, head.Hash(), opts)
	if err != nil {
		return err
	}

	return nil
}
