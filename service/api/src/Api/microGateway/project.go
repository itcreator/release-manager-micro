package microGateway

import (
	"Api/models"
	proto "Api/proto/project"
	"Api/restapi/operations/project"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
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
	service micro.Service
}

//NewProjectGateway returns go-micro gateway for project
func NewProjectGateway() IProjectGateway {
	return &projectGateway{
		service: micro.NewService(micro.Name("project.client")),
	}
}

//CreateProjectAction sends create request to micro-service
func (g *projectGateway) CreateProjectAction(params project.CreateProjectParams) middleware.Responder {
	service := micro.NewService(micro.Name("project.client"))

	proj := proto.NewProjectClient("project", service.Client())

	rsp, err := proj.Create(context.TODO(), &proto.CreateRequest{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	if uint32(codes.OK) == rsp.Status {
		fmt.Println(fmt.Sprintf("project.client: ok. Id = %v", rsp.Id))
	} else {
		fmt.Println("project.client: create fail. ")
	}

	readRsp, err := proj.Read(context.TODO(), &proto.ReadRequest{
		Id: rsp.Id,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	pr := &models.Project{
		ID:          readRsp.Project.Id,
		Name:        readRsp.Project.Name,
		Description: readRsp.Project.Description,
	}

	return project.NewCreateProjectCreated().WithPayload(pr)
}

//ReadProjectAction read project from micro-service
func (g *projectGateway) ReadProjectAction(params project.ReadProjectParams) middleware.Responder {
	service := micro.NewService(micro.Name("project.client"))

	proj := proto.NewProjectClient("project", service.Client())
	readRsp, err := proj.Read(context.TODO(), &proto.ReadRequest{
		Id: params.ID,
	})

	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	if uint32(codes.OK) == readRsp.Status {
		fmt.Println(fmt.Sprintf("project.client read: ok. Id = %v", params.ID))
	} else if uint32(codes.NotFound) == readRsp.Status {
		return project.NewReadProjectNotFound()
	}

	pr := &models.Project{
		ID:          readRsp.Project.Id,
		Name:        readRsp.Project.Name,
		Description: readRsp.Project.Description,
	}

	return project.NewReadProjectsOK().WithPayload(pr)
}

//ReadProjectAction sends update request to micro-service
func (g *projectGateway) UpdateProjectAction(params project.UpdateProjectParams) middleware.Responder {
	service := micro.NewService(micro.Name("project.client"))

	proj := proto.NewProjectClient("project", service.Client())

	// request the Hello method on the Greeter handler
	rsp, err := proj.Update(context.TODO(), &proto.UpdateRequest{
		Id:          params.ID,
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	if uint32(codes.OK) == rsp.Status {
		fmt.Println(fmt.Sprintf("project.client update: ok. Id = %v", params.ID))
	} else if uint32(codes.NotFound) == rsp.Status {
		return project.NewUpdateProjectNotFound()
	} else {
		fmt.Println(fmt.Sprintf("project.client: update fail. Id = %v, status = %v", params.ID, rsp.Status))
		return project.NewCreateProjectInternalServerError()
	}

	readRsp, err := proj.Read(context.TODO(), &proto.ReadRequest{
		Id: params.ID,
	})
	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	pr := &models.Project{
		ID:          readRsp.Project.Id,
		Name:        readRsp.Project.Name,
		Description: readRsp.Project.Description,
	}

	return project.NewUpdateProjectOK().WithPayload(pr)
}

//ListProjectsAction get all of projects from micro-service
func (g *projectGateway) ListProjectsAction(params project.ListProjectsParams) middleware.Responder {
	service := micro.NewService(micro.Name("project.client"))

	proj := proto.NewProjectClient("project", service.Client())
	listRsp, err := proj.List(context.TODO(), &proto.ListRequest{})

	if err != nil {
		fmt.Println(err)
		return project.NewCreateProjectInternalServerError()
	}

	var projects = []*models.Project{}
	for _, projResp := range listRsp.Projects {
		p := &models.Project{
			ID:          projResp.Id,
			Name:        projResp.Name,
			Description: projResp.Description,
		}
		projects = append(projects, p)
	}

	return project.NewListProjectsOK().WithPayload(projects)
}
