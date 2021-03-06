// Code generated by protoc-gen-go. DO NOT EDIT.
// source: execution.proto

package execution

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JobParams struct {
	Method               string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Params               map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JobParams) Reset()         { *m = JobParams{} }
func (m *JobParams) String() string { return proto.CompactTextString(m) }
func (*JobParams) ProtoMessage()    {}
func (*JobParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_776e2c5022e94aef, []int{0}
}

func (m *JobParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobParams.Unmarshal(m, b)
}
func (m *JobParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobParams.Marshal(b, m, deterministic)
}
func (m *JobParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobParams.Merge(m, src)
}
func (m *JobParams) XXX_Size() int {
	return xxx_messageInfo_JobParams.Size(m)
}
func (m *JobParams) XXX_DiscardUnknown() {
	xxx_messageInfo_JobParams.DiscardUnknown(m)
}

var xxx_messageInfo_JobParams proto.InternalMessageInfo

func (m *JobParams) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *JobParams) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type JobStatus struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobStatus) Reset()         { *m = JobStatus{} }
func (m *JobStatus) String() string { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()    {}
func (*JobStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_776e2c5022e94aef, []int{1}
}

func (m *JobStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatus.Unmarshal(m, b)
}
func (m *JobStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatus.Marshal(b, m, deterministic)
}
func (m *JobStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatus.Merge(m, src)
}
func (m *JobStatus) XXX_Size() int {
	return xxx_messageInfo_JobStatus.Size(m)
}
func (m *JobStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatus.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatus proto.InternalMessageInfo

func (m *JobStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*JobParams)(nil), "JobParams")
	proto.RegisterMapType((map[string]string)(nil), "JobParams.ParamsEntry")
	proto.RegisterType((*JobStatus)(nil), "JobStatus")
}

func init() {
	proto.RegisterFile("execution.proto", fileDescriptor_776e2c5022e94aef)
}

var fileDescriptor_776e2c5022e94aef = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xad, 0x48, 0x4d,
	0x2e, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21,
	0x92, 0x4a, 0x7d, 0x8c, 0x5c, 0x9c, 0x5e, 0xf9, 0x49, 0x01, 0x89, 0x45, 0x89, 0xb9, 0xc5, 0x42,
	0x62, 0x5c, 0x6c, 0xb9, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x50, 0x9e, 0x90, 0x1e, 0x17, 0x5b, 0x01, 0x58, 0x85, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x98, 0x1e, 0x5c, 0x8f, 0x1e, 0x84, 0x72, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xaa, 0x92, 0xb2,
	0xe4, 0xe2, 0x46, 0x12, 0x16, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x84, 0x9a, 0x09, 0x62, 0x0a,
	0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x81, 0xc5, 0x20, 0x1c, 0x2b, 0x26,
	0x0b, 0x46, 0x25, 0x55, 0xb0, 0x7b, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xa1, 0x9a, 0x61, 0x5c, 0x23, 0x63, 0x2e, 0x4e, 0x57,
	0x98, 0x3f, 0x85, 0xd4, 0xb8, 0xb8, 0x20, 0x9c, 0x54, 0xaf, 0xfc, 0x24, 0x21, 0x2e, 0x84, 0xe3,
	0xa4, 0xc0, 0x6c, 0x88, 0x61, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x3f, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xa4, 0x4b, 0x3d, 0x23, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExecutionClient is the client API for Execution service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutionClient interface {
	ExecuteJob(ctx context.Context, in *JobParams, opts ...grpc.CallOption) (*JobStatus, error)
}

type executionClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionClient(cc grpc.ClientConnInterface) ExecutionClient {
	return &executionClient{cc}
}

func (c *executionClient) ExecuteJob(ctx context.Context, in *JobParams, opts ...grpc.CallOption) (*JobStatus, error) {
	out := new(JobStatus)
	err := c.cc.Invoke(ctx, "/Execution/ExecuteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionServer is the server API for Execution service.
type ExecutionServer interface {
	ExecuteJob(context.Context, *JobParams) (*JobStatus, error)
}

// UnimplementedExecutionServer can be embedded to have forward compatible implementations.
type UnimplementedExecutionServer struct {
}

func (*UnimplementedExecutionServer) ExecuteJob(ctx context.Context, req *JobParams) (*JobStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteJob not implemented")
}

func RegisterExecutionServer(s *grpc.Server, srv ExecutionServer) {
	s.RegisterService(&_Execution_serviceDesc, srv)
}

func _Execution_ExecuteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).ExecuteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Execution/ExecuteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).ExecuteJob(ctx, req.(*JobParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Execution_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Execution",
	HandlerType: (*ExecutionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteJob",
			Handler:    _Execution_ExecuteJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "execution.proto",
}
