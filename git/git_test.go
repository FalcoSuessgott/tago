package git

import (
	"testing"
)

func TestIsGitRepo1(t *testing.T) {
	dir, _ := GetRootDir()
	if !IsGitRepository(dir) {
		t.Fail()
	}
}

func TestIsGitRepo2(t *testing.T) {
	if IsGitRepository("/tmp") {
		t.Fail()
	}
}
