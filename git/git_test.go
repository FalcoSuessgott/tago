package git

import (
	"fmt"
	"testing"

	"github.com/integralist/go-findroot/find"
)

func TestIsGitRepo1(t *testing.T) {
	root, err := find.Repo()
	if err != nil {
		t.Fail()
	}

	fmt.Print(root)
	if !IsGitRepository(root.Path) {
		t.Fail()
	}
}

func TestIsGitRepo2(t *testing.T) {
	if IsGitRepository("/tmp") {
		t.Fail()
	}
}
