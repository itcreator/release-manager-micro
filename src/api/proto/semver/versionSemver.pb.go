// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/semver/versionSemver.proto

/*
Package semver is a generated protocol buffer package.

It is generated from these files:
	proto/semver/versionSemver.proto

It has these top-level messages:
	GenerateRequest
	GenerateResponse
*/
package semver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GenerateRequest struct {
	ProjectUuid string `protobuf:"bytes,1,opt,name=projectUuid" json:"projectUuid,omitempty"`
	Major       uint32 `protobuf:"varint,2,opt,name=major" json:"major,omitempty"`
	Minor       uint32 `protobuf:"varint,3,opt,name=minor" json:"minor,omitempty"`
	Branch      string `protobuf:"bytes,4,opt,name=branch" json:"branch,omitempty"`
}

func (m *GenerateRequest) Reset()                    { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()               {}
func (*GenerateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GenerateRequest) GetProjectUuid() string {
	if m != nil {
		return m.ProjectUuid
	}
	return ""
}

func (m *GenerateRequest) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *GenerateRequest) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *GenerateRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type GenerateResponse struct {
	Full     string `protobuf:"bytes,1,opt,name=full" json:"full,omitempty"`
	Minor    string `protobuf:"bytes,2,opt,name=minor" json:"minor,omitempty"`
	Major    string `protobuf:"bytes,3,opt,name=major" json:"major,omitempty"`
	Branch   string `protobuf:"bytes,4,opt,name=branch" json:"branch,omitempty"`
	IsLatest bool   `protobuf:"varint,5,opt,name=isLatest" json:"isLatest,omitempty"`
}

func (m *GenerateResponse) Reset()                    { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()               {}
func (*GenerateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GenerateResponse) GetFull() string {
	if m != nil {
		return m.Full
	}
	return ""
}

func (m *GenerateResponse) GetMinor() string {
	if m != nil {
		return m.Minor
	}
	return ""
}

func (m *GenerateResponse) GetMajor() string {
	if m != nil {
		return m.Major
	}
	return ""
}

func (m *GenerateResponse) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *GenerateResponse) GetIsLatest() bool {
	if m != nil {
		return m.IsLatest
	}
	return false
}

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "semver.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "semver.GenerateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VersionSemver service

type VersionSemverClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error)
}

type versionSemverClient struct {
	c           client.Client
	serviceName string
}

func NewVersionSemverClient(serviceName string, c client.Client) VersionSemverClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "semver"
	}
	return &versionSemverClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *versionSemverClient) Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "VersionSemver.Generate", in)
	out := new(GenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VersionSemver service

type VersionSemverHandler interface {
	Generate(context.Context, *GenerateRequest, *GenerateResponse) error
}

func RegisterVersionSemverHandler(s server.Server, hdlr VersionSemverHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VersionSemver{hdlr}, opts...))
}

type VersionSemver struct {
	VersionSemverHandler
}

func (h *VersionSemver) Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error {
	return h.VersionSemverHandler.Generate(ctx, in, out)
}

func init() { proto.RegisterFile("proto/semver/versionSemver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0x71, 0x7f, 0xa2, 0xb0, 0xa8, 0x02, 0xad, 0x10, 0x58, 0x3d, 0x59, 0x39, 0xe5, 0x94,
	0x4a, 0xf0, 0x00, 0x1c, 0xb9, 0x70, 0x40, 0x41, 0x3c, 0x40, 0x5a, 0x16, 0xe1, 0xaa, 0xb5, 0xc3,
	0xda, 0x29, 0x6f, 0xc0, 0x73, 0x23, 0xd6, 0xa6, 0xe5, 0x47, 0xdc, 0x3c, 0xdf, 0x4a, 0x33, 0xe3,
	0x01, 0xd3, 0xb3, 0x8f, 0x7e, 0x11, 0x68, 0xbb, 0x23, 0x5e, 0xec, 0x88, 0x83, 0xf5, 0xee, 0x41,
	0x54, 0x23, 0x27, 0x2c, 0xd2, 0xad, 0x7a, 0x83, 0xd3, 0x5b, 0x72, 0xc4, 0x5d, 0xa4, 0x96, 0x5e,
	0x07, 0x0a, 0x11, 0x0d, 0x9c, 0xf4, 0xec, 0xd7, 0xb4, 0x8a, 0x8f, 0x83, 0x7d, 0xd2, 0xca, 0xa8,
	0xfa, 0xb8, 0xfd, 0x8e, 0xf0, 0x1c, 0xa6, 0xdb, 0x6e, 0xed, 0x59, 0x8f, 0x8c, 0xaa, 0x67, 0x6d,
	0x12, 0x42, 0xad, 0xf3, 0xac, 0xc7, 0x99, 0x7e, 0x0a, 0xbc, 0x80, 0x62, 0xc9, 0x9d, 0x5b, 0xbd,
	0xe8, 0x89, 0x18, 0x65, 0x55, 0xbd, 0x2b, 0x38, 0x3b, 0x24, 0x87, 0xde, 0xbb, 0x40, 0x88, 0x30,
	0x79, 0x1e, 0x36, 0x9b, 0x9c, 0x29, 0xef, 0x83, 0xed, 0x48, 0x60, 0xb6, 0xdd, 0x57, 0x18, 0x67,
	0x2a, 0x15, 0xfe, 0x09, 0xc3, 0x39, 0x94, 0x36, 0xdc, 0x75, 0x91, 0x42, 0xd4, 0x53, 0xa3, 0xea,
	0xb2, 0xdd, 0xeb, 0xab, 0x7b, 0x98, 0xfd, 0x18, 0x08, 0x6f, 0xa0, 0xfc, 0x2a, 0x86, 0x97, 0x4d,
	0xda, 0xa9, 0xf9, 0x35, 0xd2, 0x5c, 0xff, 0x3d, 0xa4, 0x3f, 0x54, 0x47, 0xcb, 0x42, 0x26, 0xbe,
	0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x72, 0x9c, 0xb2, 0x0e, 0x86, 0x01, 0x00, 0x00,
}
