package handler

import (
	"context"
	"github.com/stretchr/testify/suite"
	proto "semver/proto/semver"
	"testing"
)

type semverHandlerTestSuite struct {
	suite.Suite
}

func (suite *semverHandlerTestSuite) TestGenerate() {
	ctx := context.TODO()
	req := &proto.GenerateRequest{}
	rsp := new(proto.GenerateResponse)

	generator := new(versioGeneratorMock)
	handler := SemverHandler{
		Generator: generator,
	}

	handler.Generate(ctx, req, rsp)

	//suite.Equal(rsp.Version, generator.StoredVersion)
	suite.Equal(rsp.Version, "1.0.1")
}

type versioGeneratorMock struct {
	StoredVersion string
}

func (mock *versioGeneratorMock) GenerateVersion(projectID uint64, major uint32, minor uint32, branch string) string {
	mock.StoredVersion = "1.0.1"

	return mock.StoredVersion
}

func TestSemverHandlerSuite(t *testing.T) {
	suite.Run(t, new(semverHandlerTestSuite))
}
