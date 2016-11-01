package generator


import (
	"semver/model"
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
)

type semanticStrategyTestSuite struct {
	suite.Suite
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForFeatureBranch() {
	projectId := uint64(1)
	major := uint32(1)
	minor := uint32(2)
	branch := "feature/22"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionString := strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.0-feature-22", versionString)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "feature/22",
	}

	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionString = strategy.GenerateVersion(projectId, major, minor, branch)
	fmt.Printf("%v", versionString)
	suite.Equal("v1.2.0-feature-22.1", versionString)
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForDevBranch() {
	projectId := uint64(1)
	major := uint32(1)
	minor := uint32(2)
	branch := "dev"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionString := strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.0-dev", versionString)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "dev",
	}
	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionString = strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.0-dev.1", versionString)
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForReleaseBranch() {
	projectId := uint64(1)
	major := uint32(1)
	minor := uint32(2)
	branch := "release"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionString := strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.0-rc", versionString)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "release",
	}
	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionString = strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.0-rc.1", versionString)
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForProductionBranch() {
	projectId := uint64(1)
	major := uint32(1)
	minor := uint32(2)
	branch := "master"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionString := strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.0", versionString)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "master",
	}
	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionString = strategy.GenerateVersion(projectId, major, minor, branch)
	suite.Equal("v1.2.1", versionString)
}

type versionGatewayMock struct {
	VersionEmpty  bool
	StoredVersion *model.Version
}

func (mock *versionGatewayMock) Insert(ver *model.Version) *model.Version {
	mock.StoredVersion = ver
	mock.StoredVersion.Id = 1

	return ver
}

func (mock *versionGatewayMock) Select(uint64, uint32, uint32, string) (*model.Version, bool) {
	return mock.StoredVersion, mock.VersionEmpty
}

func (mock *versionGatewayMock) UpdateRevision(ver *model.Version) *model.Version {
	mock.StoredVersion = ver

	return ver
}

func TestSemanticStrategySuite(t *testing.T) {
	suite.Run(t, new(semanticStrategyTestSuite))
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

