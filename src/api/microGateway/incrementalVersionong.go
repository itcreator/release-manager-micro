package microGateway

import (
	"api/models"
	proto "api/proto/incremental"
	incremental "api/restapi/operations/version_incremental"
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/micro/go-micro"
	"google.golang.org/grpc/codes"
)

//IIncrementalVersioningGateway is go-micro gateway for incremental versioning
type IIncrementalVersioningGateway interface {
	GenerateVersionAction(params incremental.IncrementalGenerateParams) middleware.Responder
	DeleteVersionAction(params incremental.IncrementalDeleteParams) middleware.Responder
	UpdateVersionAction(params incremental.IncrementalUpdateParams) middleware.Responder
}

type incrementalVersioningGateway struct {
	incrementalClient proto.VersionIncrementalClient
}

//NewIncrementalVersioningGateway returns go-micro gateway for semver
func NewIncrementalVersioningGateway() IIncrementalVersioningGateway {
	service := micro.NewService(micro.Name("incremental_version.client"))

	return &incrementalVersioningGateway{
		incrementalClient: proto.NewVersionIncrementalClient("version_incremental", service.Client()),
	}
}

//GenerateVersionAction sends generate request to micro-service
func (g *incrementalVersioningGateway) GenerateVersionAction(params incremental.IncrementalGenerateParams) middleware.Responder {
	//TODO: read project by name from project service
	//Then use projectID

	rsp, err := g.incrementalClient.Generate(context.TODO(), &proto.GenerateRequest{
		ProjectName: params.ProjectName,
	})
	if err != nil {
		fmt.Println(err)
		return incremental.NewIncrementalGenerateInternalServerError()
	}

	fmt.Println(fmt.Sprintf("Version was generated: projectName = %v, version = %v", params.ProjectName, rsp.Version))

	s := &models.IncrementalVersionNumber{
		Version: rsp.Version,
	}

	return incremental.NewIncrementalGenerateCreated().WithPayload(s)
}

//DeleteVersionAction sends delete request to micro-service
func (g *incrementalVersioningGateway) DeleteVersionAction(params incremental.IncrementalDeleteParams) middleware.Responder {
	//TODO: read project by name from project service
	//Then use projectID

	rsp, err := g.incrementalClient.Delete(context.TODO(), &proto.DeleteRequest{
		ProjectName: params.ProjectName,
	})

	if err != nil || rsp.Status != uint32(codes.OK) {
		fmt.Println(err)
		return incremental.NewIncrementalDeleteInternalServerError()
	}

	fmt.Println(fmt.Sprintf("Version was deleted: projectName = %v", params.ProjectName))

	return incremental.NewIncrementalDeleteNoContent()
}

//UpdateVersionAction sends update request to micro-service
func (g *incrementalVersioningGateway) UpdateVersionAction(params incremental.IncrementalUpdateParams) middleware.Responder {
	//TODO: read project by name from project service
	//Then use projectID

	rsp, err := g.incrementalClient.Update(context.TODO(), &proto.UpdateRequest{
		ProjectName: params.ProjectName,
		Version: params.Body.Version,
	})

	if err != nil || rsp.Status != uint32(codes.OK) {
		fmt.Println(err)
		return incremental.NewIncrementalUpdateInternalServerError()
	}

	fmt.Println(fmt.Sprintf("Version was updated: projectName = %v, revision = %v", params.ProjectName, params.Body.Version))

	s := &models.IncrementalVersionNumber{
		Version: params.Body.Version,
	}

	return incremental.NewIncrementalUpdateOK().WithPayload(s)
}
