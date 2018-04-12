package handler

import (
	"context"
	"github.com/satori/go.uuid"
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

	generator := new(versionGeneratorMock)
	handler := SemverHandler{
		Generator: generator,
	}

	handler.Generate(ctx, req, rsp)

	//suite.Equal(rsp.Version, generator.StoredVersion)
	suite.Equal(rsp.Version, "1.0.1")
}

type versionGeneratorMock struct {
	StoredVersion string
}

func (mock *versionGeneratorMock) GenerateVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) string {
	mock.StoredVersion = "1.0.1"

	return mock.StoredVersion
}

func TestSemverHandlerSuite(t *testing.T) {
	suite.Run(t, new(semverHandlerTestSuite))
}
