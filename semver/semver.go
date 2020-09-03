package semver

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver"
)

// SemVer represents a SemVer struct
type SemVer struct {
	Version *semver.Version
}

func (sv *SemVer) BumpMajor() string {
	v := sv.Version.IncMajor()
	return v.String()
}

func (sv *SemVer) BumpMinor() string {
	v := sv.Version.IncMinor()
	return v.String()
}

func (sv *SemVer) BumpPatch() string {
	v := sv.Version.IncPatch()
	return v.String()
}

// HighestSemVer returns the highest SemVer
func HighestSemVer(vs []*SemVer) *SemVer {
	sv := []*semver.Version{}

	for _, s := range vs {
		sv = append(sv, s.Version)
	}

	sort.Sort(semver.Collection(sv))
	return vs[0]
}

// BuildBumpedOptions returns the options for bump dialog
func (sv *SemVer) BuildBumpedOptions() []string {
	opts := []string{}
	strfmt := "%s => %s (%s)"

	majorBump := sv.Version.IncMajor()
	minorBump := sv.Version.IncMinor()
	patchBump := sv.Version.IncPatch()

	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), majorBump.String(), "Major"))
	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), minorBump.String(), "Minor"))
	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), patchBump.String(), "Patch"))

	return opts
}

// NewSemVer returns a new SemVer struct
func NewSemVer(t string) (*SemVer, error) {
	v, e := semver.NewVersion(t)
	return &SemVer{v}, e
}
