package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/FalcoSuessgott/tago/git"
	"github.com/FalcoSuessgott/tago/semver"
	"github.com/FalcoSuessgott/tago/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultPrefix = "v"
)

// Tago project struct
type Tago struct {
	LatestTag                         *semver.SemVer
	Repository                        *git.Repository
	Message, Remote                   string
	Prefix, Major, Minor, Patch, Push bool
}

var (
	errParseError = errors.New("cannot parse flag")
	rootCmd       = &cobra.Command{
		Use:   "tago",
		Short: "bumping semantic versioning git tags",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			var g Tago
			var newTag string

			g.parseFlags(cmd.Flags())
			g.repo()
			g.tags()
			if !g.Major && !g.Minor && !g.Patch {
				newTag = g.prompt()
			}

			if g.Major {
				newTag = g.LatestTag.BumpMajor()
			}

			if g.Minor {
				newTag = g.LatestTag.BumpMinor()
			}

			if g.Patch {
				newTag = g.LatestTag.BumpPatch()
			}

			if g.Prefix || g.LatestTag.Prefix {
				newTag = defaultPrefix + newTag
			}

			err := g.Repository.AddTag(newTag, g.Message)
			if err != nil {
				ui.ErrorMsg(err, "could not create tag: %s", newTag)
			}

			ui.SuccessMsg("successfully created tag: %s", newTag)

			if g.Push {
				err := g.Repository.PushTags(g.Remote)
				if err != nil {
					ui.ErrorMsg(err, "cannot push to remote %s", g.Remote)
				}
				ui.SuccessMsg("pushed tags to %s", g.Remote)
			}
		},
	}
)

// Execute invokes the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("prefix", false, "use \"v\" as prefix")
	rootCmd.PersistentFlags().BoolP("push", "p", false, "push tag afterwards")
	rootCmd.PersistentFlags().StringP("remote", "r", "origin", "remote")
	rootCmd.PersistentFlags().StringP("msg", "m", "", "tag annotation message")

	rootCmd.PersistentFlags().BoolP("major", "x", false, "bump major version part")
	rootCmd.PersistentFlags().BoolP("minor", "y", false, "bump minor version part")
	rootCmd.PersistentFlags().BoolP("patch", "z", false, "bump patch version part")

	viper.SetDefault("remote", "origin")
}

func (t *Tago) parseFlags(flags *pflag.FlagSet) {
	var err error

	t.Prefix, err = flags.GetBool("prefix")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "prefix", errParseError))
	}

	t.Push, err = flags.GetBool("push")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "push", errParseError))
	}

	t.Remote, err = flags.GetString("remote")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "remote", errParseError))
	}

	t.Message, err = flags.GetString("msg")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "msg", errParseError))
	}

	t.Major, err = flags.GetBool("major")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "major", errParseError))
	}

	t.Minor, err = flags.GetBool("minor")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "minor", errParseError))
	}

	t.Patch, err = flags.GetBool("patch")
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", "patch", errParseError))
	}
}

func (t *Tago) repo() {
	var err error
	dir, err := git.GetRootDir()
	if err != nil {
		ui.ErrorMsg(err, "%s is not a git repository. Exiting", dir)
	}

	t.Repository, err = git.NewRepository(dir)
	if err != nil {
		ui.ErrorMsg(err, "can not open repository. Exiting", dir)

	}
}

func (t *Tago) tags() {
	tags := t.Repository.GetTags()
	if len(tags) == 0 {
		ui.WarningMsg(nil, "no tags found")
		newTag := ui.PromptMsg("new tag (e.g: v1.1.0):")

		_, err := semver.NewSemVer(newTag)
		if err != nil {
			ui.ErrorMsg(err, "%s is not a valid SemVer-version number.", newTag)
		}

		if t.Message == "" {
			t.Message = ui.PromptMsg("message (optional):")
		}

		err = t.Repository.AddTag(newTag, t.Message)
		if err != nil {
			ui.ErrorMsg(err, "could not create tag %s", newTag)
		}

		ui.SuccessMsg("successfully added tag: %s", newTag)
		os.Exit(0)
	}

	semVers := []*semver.SemVer{}
	invalid := 0
	for _, t := range tags {
		v, err := semver.NewSemVer(t)
		if err != nil {
			ui.WarningMsg(nil, "%s is not a valid SemVer-version number. Skipping.", t)
			invalid++
			continue
		}
		semVers = append(semVers, v)
	}

	ui.InfoMsg("found %s valid semVer tags. Invalid: %s tags", strconv.Itoa(len(semVers)), strconv.Itoa(invalid))

	t.LatestTag = semver.HighestSemVer(semVers)
	ui.SuccessMsg("latest SemVer tag: %s", t.LatestTag.String)
}

func (t *Tago) prompt() string {
	parts := t.LatestTag.BuildBumpedOptions()
	answer := ui.PromptList("which part to increment?", parts[1], parts)
	if t.Message == "" {
		t.Message = ui.PromptMsg("message (optional):")
	}

	switch answer {
	case 0:
		return t.LatestTag.BumpMajor()
	case 1:
		return t.LatestTag.BumpMinor()
	case 2:
		return t.LatestTag.BumpPatch()
	default:
		ui.ErrorMsg(nil, "Invalid Option.")
	}

	return ""
}
