package semver

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
)

// SemVer represents a SemVer struct
type SemVer struct {
	Version *semver.Version
	Prefix  bool
	String  string
}

// BumpMajor bumps major part
func (sv *SemVer) BumpMajor() string {
	v := sv.Version.IncMajor()
	return v.String()
}

// BumpMinor bumps minor part
func (sv *SemVer) BumpMinor() string {
	v := sv.Version.IncMinor()
	return v.String()
}

// BumpPatch bumps patch part
func (sv *SemVer) BumpPatch() string {
	v := sv.Version.IncPatch()
	return v.String()
}

// HighestSemVer returns the highest SemVer
func HighestSemVer(vs []*SemVer) *SemVer {
	sv := make([]*semver.Version, len(vs))

	for i, v := range vs {
		sv[i] = v.Version
	}

	sort.Sort(semver.Collection(sv))

	for _, v := range vs {
		if sv[len(sv)-1] == v.Version {
			return v
		}
	}

	return vs[0]
}

// BuildBumpedOptions returns the options for bump dialog
func (sv *SemVer) BuildBumpedOptions() []string {
	opts := []string{}
	strfmt := "%s => %s (%s)"

	majorBump := sv.Version.IncMajor()
	minorBump := sv.Version.IncMinor()
	patchBump := sv.Version.IncPatch()

	if sv.Prefix {
		strfmt = "v%s => v%s (%s)"
	}

	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), majorBump.String(), "Major: breaking API change"))
	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), minorBump.String(), "Minor: feature add"))
	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), patchBump.String(), "Patch: bug fix"))

	return opts
}

// NewSemVer returns a new SemVer struct
func NewSemVer(t string) (*SemVer, error) {
	v, e := semver.NewVersion(t)
	hasPrefix := false

	if strings.HasPrefix(t, "v") {
		hasPrefix = true
	}

	return &SemVer{v, hasPrefix, t}, e
}
