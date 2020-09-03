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

//TODO: option for prefix, default prefix : v
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
		ui.ErrorMsg(err, "%s is not a git repository. Exiting", dir)
	}

	repo, err := git.Repo(dir)
	if err != nil {
		ui.ErrorMsg(err, "Can not open repository. Exiting", dir)

	}

	tags := git.GetTags(repo)
	if len(tags) == 0 {
		ui.InfoMsg("No tags found.")
		// create tag
	}

	semVers := []*semver.SemVer{}
	invalid := 0
	for _, t := range tags {
		v, err := semver.NewSemVer(t)
		if err != nil {
			ui.WarningMsg(err, "%s is not a valid SemVer-version number. Skipping", t)
			invalid++
		}
		semVers = append(semVers, v)
	}

	ui.InfoMsg("Found %s valid semVer tags. Invalid: %s tags.", strconv.Itoa(len(semVers)), strconv.Itoa(invalid))

	highestSemVer := semver.HighestSemVer(semVers)
	ui.SuccessMsg("Latest SemVer tag: %s", highestSemVer.Version.String())

	parts := highestSemVer.BuildBumpedOptions()
	answer := ui.PromptList("Which part to increment?", parts[1], parts)
	msg := ui.PromptMsg("Message (optional):")

	newSemVer := ""
	switch answer {
	case 0:
		newSemVer = highestSemVer.BumpMajor()
	case 1:
		newSemVer = highestSemVer.BumpMinor()
	case 2:
		newSemVer = highestSemVer.BumpPatch()
	default:
		ui.ErrorMsg(nil, "Invalid Option.")
	}

	err = git.AddTag(repo, newSemVer, msg)
	if err != nil {
		ui.ErrorMsg(nil, "Could not create tag %s. Error: %s", newSemVer, err.Error())
	}

	ui.SuccessMsg("Successfully created new Tag %s.", newSemVer)
	fmt.Println(git.GetTags(repo))
}
