package handler

import (
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"project/model"
	proto "project/proto/project"
	"testing"
)

type projectHandlerTestSuite struct {
	suite.Suite
}

func (suite *projectHandlerTestSuite) TestCreate() {
	ctx := context.TODO()
	req := &proto.CreateRequest{
		Name:        "Test Project",
		Description: "Test Description",
	}
	rsp := new(proto.CreateResponse)

	repository := new(projectRepositoryMock)
	handler := &ProjectHandler{
		Repository: repository,
	}

	handler.Create(ctx, req, rsp)

	suite.Equal(rsp.Status, uint32(codes.OK))
	suite.Equal(rsp.Id, uint64(1))
	suite.Equal(repository.StoredProject.Name, req.Name)
	suite.Equal(repository.StoredProject.Description, req.Description)
}

func (suite *projectHandlerTestSuite) TestRead() {
	repository := new(projectRepositoryMock)
	repository.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.ReadRequest{
		Id: 1,
	}
	rsp := new(proto.ReadResponse)

	handler := &ProjectHandler{
		Repository: repository,
	}

	handler.Read(ctx, req, rsp)

	suite.Equal(rsp.Status, uint32(codes.OK))
	suite.Equal(rsp.Project.Id, uint64(1))
	suite.Equal(repository.StoredProject.Name, rsp.Project.Name)
	suite.Equal(repository.StoredProject.Description, rsp.Project.Description)
}

func (suite *projectHandlerTestSuite) TestReadNotFound() {
	repository := new(projectRepositoryMock)
	repository.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.ReadRequest{
		Id: 17636356,
	}
	rsp := new(proto.ReadResponse)

	handler := &ProjectHandler{
		Repository: repository,
	}

	handler.Read(ctx, req, rsp)

	suite.Equal(uint32(codes.NotFound), rsp.Status)
	suite.Nil(rsp.Project)
}

func (suite *projectHandlerTestSuite) TestUpdate() {
	repository := new(projectRepositoryMock)
	repository.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.UpdateRequest{
		Id:          1,
		Name:        "N2",
		Description: "D2",
	}
	rsp := new(proto.UpdateResponse)

	handler := &ProjectHandler{
		Repository: repository,
	}

	handler.Update(ctx, req, rsp)

	suite.Equal(rsp.Status, uint32(codes.OK))
	suite.Equal(repository.StoredProject.ID, uint64(1))
	suite.Equal(repository.StoredProject.Name, req.Name)
	suite.Equal(repository.StoredProject.Description, req.Description)
}

func (suite *projectHandlerTestSuite) TestUpdateNotFound() {
	repository := new(projectRepositoryMock)

	ctx := context.TODO()
	req := &proto.UpdateRequest{
		Id:          uint64(13435322),
		Name:        "N4",
		Description: "D4",
	}
	rsp := new(proto.UpdateResponse)

	handler := &ProjectHandler{
		Repository: repository,
	}

	handler.Update(ctx, req, rsp)

	suite.Equal(uint32(codes.NotFound), rsp.Status)
	suite.Nil(repository.StoredProject)
}

func (suite *projectHandlerTestSuite) TestList() {
	repository := new(projectRepositoryMock)
	repository.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.ListRequest{}
	rsp := new(proto.ListResponse)

	handler := &ProjectHandler{
		Repository: repository,
	}

	handler.List(ctx, req, rsp)

	suite.Len(rsp.Projects, 1)
	suite.Equal(rsp.Projects[0].Id, repository.StoredProject.ID)
	suite.Equal(rsp.Projects[0].Name, repository.StoredProject.Name)
	suite.Equal(rsp.Projects[0].Description, repository.StoredProject.Description)
}

type projectRepositoryMock struct {
	StoredProject *model.Project
}

func (mock *projectRepositoryMock) Insert(p *model.Project) {
	mock.StoredProject = p
	mock.StoredProject.ID = 1
}

func (mock *projectRepositoryMock) isNotFound(id uint64) bool {
	var notFound bool
	if nil == mock.StoredProject || mock.StoredProject.ID != id {
		notFound = true
	} else {
		notFound = false
	}

	return notFound
}

func (mock *projectRepositoryMock) SelectByID(id uint64) *model.Project {
	var p *model.Project
	if mock.isNotFound(id) {
		p = nil
	} else {
		p = mock.StoredProject
	}

	return p
}

func (mock *projectRepositoryMock) Update(p *model.Project) bool {
	var result bool
	if mock.isNotFound(p.ID) {
		result = false
	} else {
		mock.StoredProject = p
		result = true
	}

	return result
}

func (mock *projectRepositoryMock) SelectAll() []*model.Project {
	var projects = []*model.Project{}
	projects = append(projects, mock.StoredProject)

	return projects
}

func TestProjectHandlerSuite(t *testing.T) {
	suite.Run(t, new(projectHandlerTestSuite))
}
