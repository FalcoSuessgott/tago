package main

import (
	"os"
	"fmt"

	"github.com/FalcoSuessgott/gitag/ui"
	"github.com/FalcoSuessgott/gitag/git"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(git.IsGitRepository())
}
