package generator

import (
	"github.com/satori/go.uuid"
	"fmt"
	"semver/model"
	"strings"
)

// ISemverGenerator generate version
type ISemverGenerator interface {
	GenerateVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) string
}

// SemverGenerator implements `ISemverGenerator` interface
type SemverGenerator struct {
	VersionRepository model.IVersionRepository `inject:""`
}

func (s *SemverGenerator) generateCommonVersionTag(version *model.Version) string {
	commonTag := fmt.Sprintf("v%d.%d.%d", version.Major, version.Minor, 0)

	return commonTag
}

func (s *SemverGenerator) addRevisionPostfix(version *model.Version, tag string) string {
	if version.Revision > 0 {
		tag += fmt.Sprintf(".%d", version.Revision)
	}

	return tag
}

func (s *SemverGenerator) getStoredVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) *model.Version {

	rep := s.VersionRepository
	ver, isEmpty := rep.Select(projectUUID, major, minor, branch)

	if isEmpty {
		ver.ProjectUUID = projectUUID
		ver.Revision = 0
		ver.Major = major
		ver.Minor = minor
		ver.Branch = branch
	} else {
		ver.Revision++
	}

	return ver
}

// GenerateVersion function generate version for project
func (s *SemverGenerator) GenerateVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) string {
	//todo: move branch names to config DB table
	branchMaster := "master"
	branchDev := "dev"
	branchRelease := "release"

	ver := s.getStoredVersion(projectUUID, major, minor, branch)

	var versionName string

	//TODO: if current minor version is stable - disable generating of unstable versions
	if branchMaster == ver.Branch {
		versionName = fmt.Sprintf("v%d.%d.%d", ver.Major, ver.Minor, ver.Revision)
	} else if branchDev == ver.Branch {
		commonVersion := s.generateCommonVersionTag(ver)
		versionName = commonVersion + "-dev"
		versionName = s.addRevisionPostfix(ver, versionName)
	} else if branchRelease == ver.Branch {
		commonVersion := s.generateCommonVersionTag(ver)
		versionName = commonVersion + "-rc"
		versionName = s.addRevisionPostfix(ver, versionName)
	} else {
		//feature branch with any name
		//and hotfix branch
		commonVersion := s.generateCommonVersionTag(ver)
		//todo: use branch name normalizer
		versionName = fmt.Sprintf("%s-%s", commonVersion, strings.Replace(ver.Branch, "/", "-", -1))
		versionName = s.addRevisionPostfix(ver, versionName)
	}

	rep := s.VersionRepository

	if ver.UUID.Valid {
		rep.UpdateRevision(ver)
	} else {
		rep.Insert(ver)
	}

	return versionName
}
