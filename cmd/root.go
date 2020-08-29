package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/FalcoSuessgott/gitag/git"
	"github.com/FalcoSuessgott/gitag/semver"
	"github.com/FalcoSuessgott/gitag/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "gitag",
	Short: "Interactively bump git tags using SemVer",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		gitag()
	},
}

// Execute invokes the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("push", "p", false, "pushes to remote after tag has been created")
	rootCmd.PersistentFlags().StringP("remote", "r", "origin", "name of the remote")
	rootCmd.PersistentFlags().StringP("msg", "m", "", "tag message")

	rootCmd.PersistentFlags().Bool("major", false, "bump major version part")
	rootCmd.PersistentFlags().Bool("minor", false, "bump minor version part")
	rootCmd.PersistentFlags().Bool("patch", false, "bump patch version part")

	viper.SetDefault("remote", "origin")
}

func gitag() {
	dir, err := git.GetRootDir()
	if err != nil {
		ui.ErrorMsg(nil, "%s is not a git repository. Exiting", dir)
	}

	tags := git.GetTags(dir)
	latestTag := tags[len(tags)-1]

	if len(tags) == 0 {
		ui.InfoMsg("No tags found.")
	} else {
		ui.InfoMsg("Found %s tags.", strconv.Itoa(len(tags)))
		ui.SuccessMsg("Latest SemVer tag: %s", latestTag)
	}

	v, err := semver.NewSemVer("1.0.0")
	if err != nil {
		ui.ErrorMsg(err, "%s is not a valid SemVer-version number.", latestTag)
	}

	parts := semver.BuildBumpedOptions(v)
	answer := ui.PromptList("Which part to increment?", string(v.Minor), parts)
	msg := ui.PromptMsg("Message (optional):")

	switch answer {
	case 1: // major
		git.AddTag(dir, semver.IncrementMajor(v), msg)
	case 2: // minor
		git.AddTag(dir, semver.IncrementMinor(v), msg)
	case 3: // patch
		git.AddTag(dir, semver.IncrementPatch(v), msg)
	}

}
