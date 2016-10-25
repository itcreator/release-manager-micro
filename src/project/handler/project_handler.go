package handler

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"project/model"
	proto "project/proto/project"
)

//ProjectHandler is a CRUD handler for project (with out delete)
type ProjectHandler struct {
	Gateway model.IProjectGateway `inject:""`
}

//Create new project
func (h *ProjectHandler) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	project := &model.Project{
		Name:        req.Name,
		Description: req.Description,
	}

	h.Gateway.Insert(project)

	rsp.Status = uint32(codes.OK)
	rsp.Id = project.Id

	return nil
}

//Read project by id
func (h *ProjectHandler) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	project, isEmpty := h.Gateway.SelectById(req.Id)

	if isEmpty {
		rsp.Status = uint32(codes.NotFound)
	} else {
		rsp.Status = uint32(codes.OK)
		rsp.Project = &proto.ProjectItem{
			Id:          project.Id,
			Name:        project.Name,
			Description: project.Description,
		}
	}

	return nil
}

//Update projects
func (h *ProjectHandler) Update(ctx context.Context, req *proto.UpdateRequest, rsp *proto.UpdateResponse) error {
	project, notFound := h.Gateway.SelectById(req.Id)

	if notFound {
		rsp.Status = uint32(codes.NotFound)
	} else {
		project.Id = req.Id
		project.Name = req.Name
		project.Description = req.Description

		if h.Gateway.Update(project) {
			rsp.Status = uint32(codes.OK)
		} else {
			rsp.Status = uint32(codes.NotFound)
		}
	}

	return nil
}

//List all elements
func (h *ProjectHandler) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	projects := h.Gateway.SelectAll()
	for _, project := range projects {
		readRsp := &proto.ProjectItem{
			Id:          project.Id,
			Name:        project.Name,
			Description: project.Description,
		}

		rsp.Projects = append(rsp.Projects, readRsp)
	}

	return nil
}
