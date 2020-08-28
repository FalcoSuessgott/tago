package main

import (
	"testing"
	"os"
)

func TestIsGitRepo1(t *testing.T){
	if !isGitRepository(){
		t.Fail()
	}
}

func TestIsGitRepo2(t *testing.T){
	err := os.Chdir("/tmp")
	if err != nil {
		panic(err)
		
	}
	if isGitRepository(){
		t.Fail()
	}
}