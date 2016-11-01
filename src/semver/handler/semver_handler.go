package handler

import (
	proto "semver/proto/semver"
	"golang.org/x/net/context"
	"semver/generator"
)

type SemverHandler struct {
	Generator generator.ISemverGenerator `inject:""`
}

//Generate new version tag for project
func (h *SemverHandler) Generate(ctx context.Context, req *proto.GenerateRequest, rsp *proto.GenerateResponse) error {
	rsp.Version = h.Generator.GenerateVersion(req.ProjectId, req.Major, req.Minor, req.Branch)

	return nil
}
