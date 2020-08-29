package git

import (
	"fmt"
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
func GetTags(dir string) []string {

	tags := []string{}
	r, _ := git.PlainOpen(dir)
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

// AddTag adds a tag
func AddTag(dir, tag string, message ...string) {
	opts := &git.CreateTagOptions{
		Tagger: defaultSignature(),
	}

	r, _ := git.PlainOpen(dir)

	head, err := r.Head()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(head)
	if len(message) > 0 {
		opts.Message = message[0]
	}

	_, err = r.CreateTag(tag, head.Hash(), opts)
	if err != nil {
		fmt.Println(err)
	}

}
