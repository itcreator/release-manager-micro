package handler

import (
	"context"
	"github.com/stretchr/testify/suite"
	proto "incremental/proto/incremental"
	"testing"
)

type incrementalHandlerTestSuite struct {
	suite.Suite
}

func (suite *incrementalHandlerTestSuite) TestGenerate() {
	ctx := context.TODO()
	req := &proto.GenerateRequest{}
	rsp := new(proto.GenerateResponse)

	generator := new(versionGeneratorMock)
	handler := IncrementalVersionHandler{
		Generator: generator,
	}

	handler.Generate(ctx, req, rsp)

	//suite.Equal(rsp.Version, generator.StoredVersion)
	suite.Equal(rsp.Version, uint64(22))
}

type versionGeneratorMock struct {
	StoredVersion uint64
}

func (mock *versionGeneratorMock) GenerateVersion(projectName string) uint64 {
	mock.StoredVersion = 22

	return mock.StoredVersion
}

func TestIncrementalHandlerSuite(t *testing.T) {
	suite.Run(t, new(incrementalHandlerTestSuite))
}
