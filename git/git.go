package git

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

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

// AddTag adds a tag
func AddTag(dir, tag string, message ...string) {
	opts := &git.CreateTagOptions{}
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
