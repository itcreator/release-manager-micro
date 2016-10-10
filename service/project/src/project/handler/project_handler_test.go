package handler

import (
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
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

	suite.Equal(rsp.Success, true)
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

	suite.Equal(rsp.Id, uint64(1))
	suite.Equal(gateway.StoredProject.Name, rsp.Name)
	suite.Equal(gateway.StoredProject.Description, rsp.Description)
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

	suite.Equal(rsp.Success, true)
	suite.Equal(gateway.StoredProject.Id, uint64(1))
	suite.Equal(gateway.StoredProject.Name, req.Name)
	suite.Equal(gateway.StoredProject.Description, req.Description)
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

func (mock *projectGatewayMock) SelectById(id uint64) (*model.Project, bool) {
	return mock.StoredProject, nil == mock.StoredProject
}

func (mock *projectGatewayMock) Update(p *model.Project) *model.Project {
	mock.StoredProject = p

	return p
}

func (mock *projectGatewayMock) SelectAll() []*model.Project {
	var projects = []*model.Project{}
	projects = append(projects, mock.StoredProject)

	return projects
}

func TestProjectHandlerSuite(t *testing.T) {
	suite.Run(t, new(projectHandlerTestSuite))
}
