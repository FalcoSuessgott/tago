package semver

import (
	"fmt"
	"log"

	"github.com/blang/semver/v4"
)

// NewSemVer returns a new SemVer object of the provided version
func NewSemVer(version string) (semver.Version, error) {
	return semver.Make(version)
}

// IncrementMajor increments the major Version
func IncrementMajor(v semver.Version) string {
	err := v.IncrementMajor()
	if err != nil {
		log.Print(err)
	}

	return v.String()
}

// IncrementMinor increments the minor Version
func IncrementMinor(v semver.Version) string {
	err := v.IncrementMinor()
	if err != nil {
		log.Print(err)
	}

	return v.String()
}

// IncrementPatch increments the patch Version
func IncrementPatch(v semver.Version) string {
	err := v.IncrementPatch()
	if err != nil {
		log.Print(err)
	}

	return v.String()
}

// BuildBumpedOptions returns the options for bump dialog
func BuildBumpedOptions(v semver.Version) []string {
	opts := []string{}
	strfmt := "v%s => v%s (%s)"
	opts = append(opts, fmt.Sprintf(strfmt, v.String(), IncrementMajor(v), "Major"))
	opts = append(opts, fmt.Sprintf(strfmt, v.String(), IncrementMinor(v), "Minor"))
	opts = append(opts, fmt.Sprintf(strfmt, v.String(), IncrementPatch(v), "Patch"))

	return opts
}
