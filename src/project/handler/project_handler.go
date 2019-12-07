package handler

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"project/model"
	proto "project/proto/project"
)

//ProjectHandler is a CRUD handler for project (with out delete)
type ProjectHandler struct {
	Repository model.IProjectRepository `inject:""`
}

//Create new project
func (h *ProjectHandler) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	project := &model.Project{
		Name:        req.Name,
		Description: req.Description,
	}

	h.Repository.Insert(project)

	rsp.Status = uint32(codes.OK)

	rsp.Uuid = project.UUID.String()

	return nil
}

//Read project by uuid
func (h *ProjectHandler) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	project := h.Repository.SelectByUUID(uuid.MustParse(req.Uuid))

	if nil == project {
		rsp.Status = uint32(codes.NotFound)
	} else {
		rsp.Status = uint32(codes.OK)
		rsp.Project = &proto.ProjectItem{
			Uuid:        project.UUID.String(),
			Name:        project.Name,
			Description: project.Description,
		}
	}

	return nil
}

//Update projects
func (h *ProjectHandler) Update(ctx context.Context, req *proto.UpdateRequest, rsp *proto.UpdateResponse) error {
	project := h.Repository.SelectByUUID(uuid.MustParse(req.Uuid))

	if nil == project {
		rsp.Status = uint32(codes.NotFound)
	} else {
		project.UUID = uuid.MustParse(req.Uuid)
		project.Name = req.Name
		project.Description = req.Description

		if h.Repository.Update(project) {
			rsp.Status = uint32(codes.OK)
		} else {
			rsp.Status = uint32(codes.NotFound)
		}
	}

	return nil
}

//List all elements
func (h *ProjectHandler) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	projects := h.Repository.SelectAll()
	for _, project := range projects {
		readRsp := &proto.ProjectItem{
			Uuid:        project.UUID.String(),
			Name:        project.Name,
			Description: project.Description,
		}

		rsp.Projects = append(rsp.Projects, readRsp)
	}

	return nil
}
