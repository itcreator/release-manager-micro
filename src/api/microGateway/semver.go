package microGateway

import (
	"api/models"
	proto "api/proto/semver"
	semver "api/restapi/operations/version_semantic"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/micro/go-micro"
	"context"
)

//ISemverGateway is go-micro gateway for semver
type ISemverGateway interface {
	GenerateVersionAction(params semver.SemverGenerateParams) middleware.Responder
}

type semverGateway struct {
	semverClient proto.VersionSemverClient
}

//NewSemverGateway returns go-micro gateway for semver
func NewSemverGateway() ISemverGateway {
	service := micro.NewService(micro.Name("semver.client"))

	return &semverGateway{
		semverClient: proto.NewVersionSemverClient("semver", service.Client()),
	}
}

//GenerateVersionAction sends generate request to micro-service
func (g *semverGateway) GenerateVersionAction(params semver.SemverGenerateParams) middleware.Responder {
	rsp, err := g.semverClient.Generate(context.TODO(), &proto.GenerateRequest{
		ProjectId: params.ProjectID,
		Major: params.Body.Major,
		Minor: params.Body.Minor,
		Branch: params.Body.Branch,
	})
	if err != nil {
		fmt.Println(err)
		return semver.NewSemverGenerateInternalServerError()
	}

	fmt.Println(fmt.Sprintf("Version was generated: projectId = %v, version = %v", params.ProjectID, rsp.Version))

	s := &models.SemverNumber{
		Version: rsp.Version,
	}

	return semver.NewSemverGenerateCreated().WithPayload(s)
}
