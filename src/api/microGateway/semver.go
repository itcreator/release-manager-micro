package microGateway

import (
	"api/models"
	protoProject "api/proto/project"
	proto "api/proto/semver"
	"api/restapi/operations/project"
	semver "api/restapi/operations/version_semantic"
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/micro/go-micro"
	"google.golang.org/grpc/codes"
)

//ISemverGateway is go-micro gateway for semver
type ISemverGateway interface {
	GenerateVersionAction(params semver.SemverGenerateParams) middleware.Responder
}

type semverGateway struct {
	projectClient protoProject.ProjectClient
	semverClient  proto.VersionSemverClient
}

//NewSemverGateway returns go-micro gateway for semver
func NewSemverGateway() ISemverGateway {
	service := micro.NewService(micro.Name("semver.client"))

	return &semverGateway{
		projectClient: protoProject.NewProjectClient("project", service.Client()),
		semverClient:  proto.NewVersionSemverClient("semver", service.Client()),
	}
}

//GenerateVersionAction sends generate request to micro-service
func (g *semverGateway) GenerateVersionAction(params semver.SemverGenerateParams) middleware.Responder {
	readRsp, err := g.projectClient.Read(context.TODO(), &protoProject.ReadRequest{
		Uuid: string(params.ProjectUUID),
	})

	if err != nil {
		fmt.Println(err)
		return project.NewReadProjectInternalServerError()
	}

	if uint32(codes.NotFound) == readRsp.Status {
		return semver.NewSemverGenerateNotFound()
	} else if uint32(codes.OK) == readRsp.Status {
		fmt.Println(fmt.Sprintf("Project exist: Id = %v", params.ProjectUUID))
	}

	rsp, err := g.semverClient.Generate(context.TODO(), &proto.GenerateRequest{
		ProjectUuid: string(params.ProjectUUID),
		Major:       params.Body.Major,
		Minor:       params.Body.Minor,
		Branch:      params.Body.Branch,
	})
	if err != nil {
		fmt.Println(err)
		return semver.NewSemverGenerateInternalServerError()
	}

	fmt.Println(fmt.Sprintf("Version was generated: projectUUID = %v, version = %v", params.ProjectUUID, rsp.Full))

	all := []string{rsp.Full}

	if "" != rsp.Minor {
		all = append(all, rsp.Minor)
	}

	if "" != rsp.Major {
		all = append(all, rsp.Major)
	}

	if "" != rsp.Branch && rsp.Branch != rsp.Full {
		all = append(all, rsp.Branch)
	}

	if rsp.IsLatest {
		all = append(all, "latest")
	}

	s := &models.SemverTagSet{
		All:      all,
		IsLatest: rsp.IsLatest,
		Full:     rsp.Full,
		Major:    rsp.Major,
		Minor:    rsp.Minor,
		Branch:   rsp.Branch,
	}

	return semver.NewSemverGenerateCreated().WithPayload(s)
}
