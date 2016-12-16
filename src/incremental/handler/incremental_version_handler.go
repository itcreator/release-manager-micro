package handler

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"incremental/generator"
	proto "incremental/proto/incremental"
)

//IncrementalVersionHandler is a handler for generate version
type IncrementalVersionHandler struct {
	Generator generator.IIncrementalVersionGenerator `inject:""`
}

//Generate new version tag for project
func (h *IncrementalVersionHandler) Generate(ctx context.Context, req *proto.GenerateRequest, rsp *proto.GenerateResponse) error {
	rsp.Version = h.Generator.GenerateVersion(req.ProjectName)

	return nil
}

//Delete generated version
func (h *IncrementalVersionHandler) Delete(ctx context.Context, req *proto.DeleteRequest, rsp *proto.DeleteResponse) error {
	h.Generator.DeleteVersion(req.ProjectName)
	rsp.Status = uint32(codes.OK)

	return nil
}

//Update generated version
func (h *IncrementalVersionHandler) Update(ctx context.Context, req *proto.UpdateRequest, rsp *proto.UpdateResponse) error {
	h.Generator.SetVersion(req.ProjectName, req.Version)
	rsp.Status = uint32(codes.OK)

	return nil
}
