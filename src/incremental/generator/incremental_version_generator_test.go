package generator

import (
	"github.com/stretchr/testify/suite"
	//"io"
	"incremental/model"
	"testing"
)

type incrementalStrategyTestSuite struct {
	suite.Suite
}

func (suite *incrementalStrategyTestSuite) TestGenerateVersion() {
	projectName := "test"

	//revision not exist in database
	rep := &versionGatewayMock{VersionEmpty: true, StoredVersion: new(model.Version)}
	strategy := IncrementalVersionGenerator{VersionRepository: rep}
	versionString := strategy.GenerateVersion(projectName)
	suite.Equal(uint64(1), versionString)

	//revision already saved
	version := &model.Version{
		Revision: 1,
	}

	rep = &versionGatewayMock{VersionEmpty: false, StoredVersion: version}
	strategy = IncrementalVersionGenerator{VersionRepository: rep}
	versionString = strategy.GenerateVersion(projectName)
	suite.Equal(uint64(2), versionString)
}

type versionGatewayMock struct {
	VersionEmpty  bool
	StoredVersion *model.Version
}

func (mock *versionGatewayMock) Insert(ver *model.Version) *model.Version {
	mock.StoredVersion = ver
	mock.StoredVersion.ID = 1

	return ver
}

func (mock *versionGatewayMock) Select(string) (*model.Version, bool) {
	return mock.StoredVersion, mock.VersionEmpty
}

func (mock *versionGatewayMock) Update(ver *model.Version) *model.Version {
	mock.StoredVersion = ver

	return ver
}

func TestIncrementalStrategySuite(t *testing.T) {
	suite.Run(t, new(incrementalStrategyTestSuite))
}
