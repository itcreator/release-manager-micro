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

	gateway := new(projectGatewayMock)
	handler := &ProjectHandler{
		Gateway: gateway,
	}

	handler.Create(ctx, req, rsp)

	suite.Equal(rsp.Status, uint32(codes.OK))
	suite.Equal(rsp.Id, uint64(1))
	suite.Equal(gateway.StoredProject.Name, req.Name)
	suite.Equal(gateway.StoredProject.Description, req.Description)
}

func (suite *projectHandlerTestSuite) TestRead() {
	gateway := new(projectGatewayMock)
	gateway.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.ReadRequest{
		Id: 1,
	}
	rsp := new(proto.ReadResponse)

	handler := &ProjectHandler{
		Gateway: gateway,
	}

	handler.Read(ctx, req, rsp)

	suite.Equal(rsp.Status, uint32(codes.OK))
	suite.Equal(rsp.Project.Id, uint64(1))
	suite.Equal(gateway.StoredProject.Name, rsp.Project.Name)
	suite.Equal(gateway.StoredProject.Description, rsp.Project.Description)
}

func (suite *projectHandlerTestSuite) TestReadNotFound() {
	gateway := new(projectGatewayMock)
	gateway.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.ReadRequest{
		Id: 17636356,
	}
	rsp := new(proto.ReadResponse)

	handler := &ProjectHandler{
		Gateway: gateway,
	}

	handler.Read(ctx, req, rsp)

	suite.Equal(uint32(codes.NotFound), rsp.Status)
	suite.Nil(rsp.Project)
}

func (suite *projectHandlerTestSuite) TestUpdate() {
	gateway := new(projectGatewayMock)
	gateway.Insert(&model.Project{
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
		Gateway: gateway,
	}

	handler.Update(ctx, req, rsp)

	suite.Equal(rsp.Status, uint32(codes.OK))
	suite.Equal(gateway.StoredProject.Id, uint64(1))
	suite.Equal(gateway.StoredProject.Name, req.Name)
	suite.Equal(gateway.StoredProject.Description, req.Description)
}

func (suite *projectHandlerTestSuite) TestUpdateNotFound() {
	gateway := new(projectGatewayMock)

	ctx := context.TODO()
	req := &proto.UpdateRequest{
		Id:          uint64(13435322),
		Name:        "N4",
		Description: "D4",
	}
	rsp := new(proto.UpdateResponse)

	handler := &ProjectHandler{
		Gateway: gateway,
	}

	handler.Update(ctx, req, rsp)

	suite.Equal(uint32(codes.NotFound), rsp.Status)
	suite.Nil(gateway.StoredProject)
}

func (suite *projectHandlerTestSuite) TestList() {
	gateway := new(projectGatewayMock)
	gateway.Insert(&model.Project{
		Name:        "N",
		Description: "D",
	})

	ctx := context.TODO()
	req := &proto.ListRequest{}
	rsp := new(proto.ListResponse)

	handler := &ProjectHandler{
		Gateway: gateway,
	}

	handler.List(ctx, req, rsp)

	suite.Len(rsp.Projects, 1)
	suite.Equal(rsp.Projects[0].Id, gateway.StoredProject.Id)
	suite.Equal(rsp.Projects[0].Name, gateway.StoredProject.Name)
	suite.Equal(rsp.Projects[0].Description, gateway.StoredProject.Description)
}

type projectGatewayMock struct {
	StoredProject *model.Project
}

func (mock *projectGatewayMock) Insert(p *model.Project) {
	mock.StoredProject = p
	mock.StoredProject.Id = 1
}

func (mock *projectGatewayMock) isNotFound(id uint64) bool {
	var notFound bool
	if nil == mock.StoredProject || mock.StoredProject.Id != id {
		notFound = true
	} else {
		notFound = false
	}

	return notFound
}

func (mock *projectGatewayMock) SelectById(id uint64) (*model.Project, bool) {
	if mock.isNotFound(id) {
		return &model.Project{}, true
	}

	return mock.StoredProject, false
}

func (mock *projectGatewayMock) Update(p *model.Project) bool {
	var result bool
	if mock.isNotFound(p.Id) {
		result = false
	} else {
		mock.StoredProject = p
		result = true
	}

	return result
}

func (mock *projectGatewayMock) SelectAll() []*model.Project {
	var projects = []*model.Project{}
	projects = append(projects, mock.StoredProject)

	return projects
}

func TestProjectHandlerSuite(t *testing.T) {
	suite.Run(t, new(projectHandlerTestSuite))
}
