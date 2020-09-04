package semver

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
)

const (
	defaultPrefix = "v"
)

// SemVer represents a SemVer struct
type SemVer struct {
	Version *semver.Version
	Prefix  string
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

	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), majorBump.String(), "Major"))
	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), minorBump.String(), "Minor"))
	opts = append(opts, fmt.Sprintf(strfmt, sv.Version.String(), patchBump.String(), "Patch"))

	return opts
}

// NewSemVer returns a new SemVer struct
func NewSemVer(t string, prefix bool) (*SemVer, error) {
	var v *semver.Version
	var e error
	p := ""

	if strings.HasPrefix(t, defaultPrefix) || prefix {
		t = strings.Replace(t, "v", "", -1)
		v, e = semver.NewVersion(t)
		p = defaultPrefix
	}

	if !strings.HasPrefix(t, defaultPrefix) && !prefix {
		v, e = semver.NewVersion(t)
	}

	return &SemVer{v, p}, e
}
