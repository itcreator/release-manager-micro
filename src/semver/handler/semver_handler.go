package handler

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"semver/generator"
	proto "semver/proto/semver"
)

//SemverHandler is a handler for generate version
type SemverHandler struct {
	Generator generator.ISemverGenerator `inject:""`
}

//Generate new version tag for project
func (h *SemverHandler) Generate(ctx context.Context, req *proto.GenerateRequest, rsp *proto.GenerateResponse) error {
	projectUUID := uuid.MustParse(req.ProjectUuid)
	tagSet := h.Generator.GenerateVersion(projectUUID, req.Major, req.Minor, req.Branch)
	rsp.IsLatest = tagSet.IsLatest
	rsp.Full = tagSet.Full
	if nil != tagSet.Major {
		rsp.Major = *tagSet.Major
	}
	if nil != tagSet.Minor {
		rsp.Minor = *tagSet.Minor
	}

	if nil != tagSet.Branch {
		rsp.Branch = *tagSet.Branch
	}

	return nil
}
