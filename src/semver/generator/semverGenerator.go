package generator

import (
	"fmt"
	"github.com/google/uuid"
	"semver/model"
	"strings"
)

// ISemverGenerator generate version
type ISemverGenerator interface {
	GenerateVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) model.TagSet
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

// GenerateVersion function generate set of tags for project
func (s *SemverGenerator) GenerateVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) model.TagSet {
	//todo: move branch names to config DB table
	branchMaster := "master"
	branchDev := "dev"
	branchRelease := "release"

	if 0 == strings.Index(branch, branchRelease) {
		branch = branchRelease
	}

	ver := s.getStoredVersion(projectUUID, major, minor, branch)

	var fullVersionName string
	var tagSet = model.TagSet{
		IsLatest: false,
	}

	//TODO: if current minor version is stable - disable generating of unstable versions
	if branchMaster == ver.Branch {
		majorVersionName := fmt.Sprintf("v%d", ver.Major)
		minorVersionName := fmt.Sprintf("v%d.%d", ver.Major, ver.Minor)

		tagSet.IsLatest = true
		tagSet.Major = &majorVersionName
		tagSet.Minor = &minorVersionName
		tagSet.Full = fmt.Sprintf("v%d.%d.%d", ver.Major, ver.Minor, ver.Revision)
	} else if branchDev == ver.Branch {
		commonVersion := s.generateCommonVersionTag(ver)
		branchTag := commonVersion + "-dev"
		fullVersionName = s.addRevisionPostfix(ver, branchTag)

		tagSet.Full = fullVersionName
		tagSet.Branch = &branchTag
	} else if branchRelease == ver.Branch {
		commonVersion := s.generateCommonVersionTag(ver)
		branchTag := commonVersion + "-rc"
		fullVersionName = s.addRevisionPostfix(ver, branchTag)

		tagSet.Full = fullVersionName
		tagSet.Branch = &branchTag
	} else {
		//feature branch with any name
		//and hotfix branch
		commonVersion := s.generateCommonVersionTag(ver)
		//todo: use branch name normalizer
		replacer := strings.NewReplacer("/", "-", "#", "-")
		branchTag := fmt.Sprintf("%s-%s", commonVersion, replacer.Replace(ver.Branch))
		fullVersionName = s.addRevisionPostfix(ver, branchTag)

		tagSet.Full = fullVersionName
		tagSet.Branch = &branchTag
	}

	rep := s.VersionRepository

	if nil != ver.UUID {
		rep.UpdateRevision(ver)
	} else {
		rep.Insert(ver)
	}

	return tagSet
}
