package microGateway

import (
	"api/models"
	proto "api/proto/project"
	"api/restapi/operations/project"
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/micro/go-micro"
	"google.golang.org/grpc/codes"
)

//IProjectGateway is go-micro gateway for project
type IProjectGateway interface {
	CreateProjectAction(params project.CreateProjectParams) middleware.Responder
	ReadProjectAction(params project.ReadProjectParams) middleware.Responder
	UpdateProjectAction(params project.UpdateProjectParams) middleware.Responder
	ListProjectsAction(params project.ListProjectsParams) middleware.Responder
}

type projectGateway struct {
	projectClient proto.ProjectClient
}

//NewProjectGateway returns go-micro gateway for project
func NewProjectGateway() IProjectGateway {
	service := micro.NewService(micro.Name("project.client"))

	return &projectGateway{
		projectClient: proto.NewProjectClient("project", service.Client()),
	}
}

//CreateProjectAction sends create request to micro-service
func (g *projectGateway) CreateProjectAction(params project.CreateProjectParams) middleware.Responder {
	rsp, err := g.projectClient.Create(context.TODO(), &proto.CreateRequest{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	if uint32(codes.OK) == rsp.Status {
		fmt.Println(fmt.Sprintf("project.client: ok. Id = %v", rsp.Uuid))
	} else {
		fmt.Println("project.client: create fail. ")
	}

	readRsp, err := g.projectClient.Read(context.TODO(), &proto.ReadRequest{
		Uuid: rsp.Uuid,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	pr := &models.Project{
		UUID:        strfmt.UUID(readRsp.Project.Uuid),
		Name:        readRsp.Project.Name,
		Description: readRsp.Project.Description,
	}

	return project.NewCreateProjectCreated().WithPayload(pr)
}

//ReadProjectAction read project from micro-service
func (g *projectGateway) ReadProjectAction(params project.ReadProjectParams) middleware.Responder {
	readRsp, err := g.projectClient.Read(context.TODO(), &proto.ReadRequest{
		Uuid: string(params.UUID),
	})

	if err != nil {
		fmt.Println(err)
		return project.NewReadProjectInternalServerError()
	}

	if uint32(codes.OK) == readRsp.Status {
		fmt.Println(fmt.Sprintf("project.client read: ok. Id = %v", params.UUID))
	} else if uint32(codes.NotFound) == readRsp.Status {
		return project.NewReadProjectNotFound()
	}

	pr := &models.Project{
		UUID:        strfmt.UUID(readRsp.Project.Uuid),
		Name:        readRsp.Project.Name,
		Description: readRsp.Project.Description,
	}

	return project.NewReadProjectOK().WithPayload(pr)
}

//ReadProjectAction sends update request to micro-service
func (g *projectGateway) UpdateProjectAction(params project.UpdateProjectParams) middleware.Responder {
	rsp, err := g.projectClient.Update(context.TODO(), &proto.UpdateRequest{
		Uuid:        string(params.UUID),
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	if uint32(codes.OK) == rsp.Status {
		fmt.Println(fmt.Sprintf("project.client update: ok. Id = %v", params.UUID))
	} else if uint32(codes.NotFound) == rsp.Status {
		return project.NewUpdateProjectNotFound()
	} else {
		fmt.Println(fmt.Sprintf("project.client: update fail. Id = %v, status = %v", params.UUID, rsp.Status))
		return project.NewCreateProjectInternalServerError()
	}

	readRsp, err := g.projectClient.Read(context.TODO(), &proto.ReadRequest{
		Uuid: string(params.UUID),
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	pr := &models.Project{
		UUID:        strfmt.UUID(readRsp.Project.Uuid),
		Name:        readRsp.Project.Name,
		Description: readRsp.Project.Description,
	}

	return project.NewUpdateProjectOK().WithPayload(pr)
}

//ListProjectsAction get all of projects from micro-service
func (g *projectGateway) ListProjectsAction(params project.ListProjectsParams) middleware.Responder {
	listRsp, err := g.projectClient.List(context.TODO(), &proto.ListRequest{})

	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	var projects = []*models.Project{}
	for _, listResp := range listRsp.Projects {
		p := &models.Project{
			UUID:        strfmt.UUID(listResp.Uuid),
			Name:        listResp.Name,
			Description: listResp.Description,
		}
		projects = append(projects, p)
	}

	return project.NewListProjectsOK().WithPayload(projects)
}
