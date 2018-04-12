package handler

import (
	"golang.org/x/net/context"
	"semver/generator"
	proto "semver/proto/semver"
	"github.com/satori/go.uuid"
)

//SemverHandler is a handler for generate version
type SemverHandler struct {
	Generator generator.ISemverGenerator `inject:""`
}

//Generate new version tag for project
func (h *SemverHandler) Generate(ctx context.Context, req *proto.GenerateRequest, rsp *proto.GenerateResponse) error {
	projectUUID := uuid.FromStringOrNil(req.ProjectUuid)
	rsp.Version = h.Generator.GenerateVersion(projectUUID, req.Major, req.Minor, req.Branch)

	return nil
}
