package generator

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"io"
	"semver/model"
	"testing"
)

type semanticStrategyTestSuite struct {
	suite.Suite
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForFeatureBranch() {
	projectUUID, _ := uuid.FromString("7df6fe94-4f84-4803-8846-4b05b8baafd2")
	major := uint32(1)
	minor := uint32(2)
	branch := "feature/22"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionTag := strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.False(versionTag.IsLatest)
	suite.Equal("v1.2.0-feature-22", versionTag.Full)
	suite.Equal("v1.2.0-feature-22", *versionTag.Branch)
	suite.Nil(versionTag.Major)
	suite.Nil(versionTag.Minor)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "feature/22",
	}

	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionTag = strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.False(versionTag.IsLatest)
	suite.Equal("v1.2.0-feature-22.1", versionTag.Full)
	suite.Equal("v1.2.0-feature-22", *versionTag.Branch)
	suite.Nil(versionTag.Major)
	suite.Nil(versionTag.Minor)
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForDevBranch() {
	projectUUID, _ := uuid.FromString("6df6fe94-4f84-4803-8846-4b05b8baafd2")
	major := uint32(1)
	minor := uint32(2)
	branch := "dev"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionTag := strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.False(versionTag.IsLatest)
	suite.Equal("v1.2.0-dev", versionTag.Full)
	suite.Equal("v1.2.0-dev", *versionTag.Branch)
	suite.Nil(versionTag.Major)
	suite.Nil(versionTag.Minor)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "dev",
	}
	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionTag = strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.False(versionTag.IsLatest)
	suite.Equal("v1.2.0-dev.1", versionTag.Full)
	suite.Equal("v1.2.0-dev", *versionTag.Branch)
	suite.Nil(versionTag.Major)
	suite.Nil(versionTag.Minor)
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForReleaseBranch() {
	projectUUID, _ := uuid.FromString("5df6fe94-4f84-4803-8846-4b05b8baafd2")
	major := uint32(1)
	minor := uint32(2)
	branch := "release"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionTag := strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.Equal("v1.2.0-rc", versionTag.Full)
	suite.Equal("v1.2.0-rc", *versionTag.Branch)
	suite.False(versionTag.IsLatest)
	suite.Nil(versionTag.Major)
	suite.Nil(versionTag.Minor)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "release",
	}
	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionTag = strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.Equal("v1.2.0-rc.1", versionTag.Full)
	suite.Equal("v1.2.0-rc", *versionTag.Branch)
	suite.False(versionTag.IsLatest)
	suite.Nil(versionTag.Major)
	suite.Nil(versionTag.Minor)
}

func (suite *semanticStrategyTestSuite) TestGenerateVersionForProductionBranch() {
	projectUUID, _ := uuid.FromString("4df6fe94-4f84-4803-8846-4b05b8baafd2")
	major := uint32(1)
	minor := uint32(2)
	branch := "master"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := SemverGenerator{VersionRepository: rep}
	versionTag := strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.True(versionTag.IsLatest)
	suite.Equal("v1.2.0", versionTag.Full)
	suite.Equal("v1.2", *versionTag.Minor)
	suite.Equal("v1", *versionTag.Major)
	suite.Nil(versionTag.Branch)

	//revision already saved
	version := &model.Version{
		Major:    1,
		Minor:    2,
		Revision: 0,
		Branch:   "master",
	}
	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = SemverGenerator{VersionRepository: rep}
	versionTag = strategy.GenerateVersion(projectUUID, major, minor, branch)
	suite.True(versionTag.IsLatest)
	suite.Equal("v1.2.1", versionTag.Full)
	suite.Equal("v1.2", *versionTag.Minor)
	suite.Equal("v1", *versionTag.Major)
	suite.Nil(versionTag.Branch)
}

type versionGatewayMock struct {
	VersionEmpty  bool
	StoredVersion *model.Version
}

func (mock *versionGatewayMock) Insert(ver *model.Version) *model.Version {
	mock.StoredVersion = ver
	id, err := uuid.FromString("8df6fe94-4f84-4803-8846-4b05b8baafd2")
	mock.StoredVersion.UUID = uuid.NullUUID{UUID: id, Valid: nil == err}

	return ver
}

func (mock *versionGatewayMock) Select(uuid.UUID, uint32, uint32, string) (*model.Version, bool) {
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
