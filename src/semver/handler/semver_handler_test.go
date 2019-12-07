package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"semver/model"
	proto "semver/proto/semver"
	"testing"
)

type semverHandlerTestSuite struct {
	suite.Suite
}

func (suite *semverHandlerTestSuite) TestGenerate() {
	ctx := context.TODO()
	req := &proto.GenerateRequest{ProjectUuid: uuid.New().String()}
	rsp := new(proto.GenerateResponse)

	generator := new(versionGeneratorMock)
	handler := SemverHandler{
		Generator: generator,
	}

	handler.Generate(ctx, req, rsp)

	//suite.Equal(rsp.Version, generator.StoredVersion)
	suite.Equal("v1.0.1-dev.2", rsp.GetFull())
	suite.Equal("v1.0.1-dev", rsp.GetBranch())
	suite.False(rsp.GetIsLatest())
	suite.Equal("", rsp.Major)
	suite.Equal("", rsp.Minor)
}

type versionGeneratorMock struct {
	StoredVersion model.TagSet
}

func (mock *versionGeneratorMock) GenerateVersion(projectUUID uuid.UUID, major uint32, minor uint32, branch string) model.TagSet {
	branchTag := "v1.0.1-dev"
	mock.StoredVersion = model.TagSet{
		IsLatest: false,
		Full:     branchTag + ".2",
		Branch:   &branchTag,
	}

	return mock.StoredVersion
}

func TestSemverHandlerSuite(t *testing.T) {
	suite.Run(t, new(semverHandlerTestSuite))
}
