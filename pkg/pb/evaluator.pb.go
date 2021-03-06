// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/evaluator.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EvaluateRequest struct {
	// A Matches proposed by the Match Function representing a candidate of the final results.
	Match                *Match   `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateRequest) Reset()         { *m = EvaluateRequest{} }
func (m *EvaluateRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateRequest) ProtoMessage()    {}
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c58cb7dff9acb0f, []int{0}
}

func (m *EvaluateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateRequest.Unmarshal(m, b)
}
func (m *EvaluateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateRequest.Merge(m, src)
}
func (m *EvaluateRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateRequest.Size(m)
}
func (m *EvaluateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateRequest proto.InternalMessageInfo

func (m *EvaluateRequest) GetMatch() *Match {
	if m != nil {
		return m.Match
	}
	return nil
}

type EvaluateResponse struct {
	// A Match ID representing a shortlisted match returned by the evaluator as the final result.
	MatchId              string   `protobuf:"bytes,2,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateResponse) Reset()         { *m = EvaluateResponse{} }
func (m *EvaluateResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateResponse) ProtoMessage()    {}
func (*EvaluateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c58cb7dff9acb0f, []int{1}
}

func (m *EvaluateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateResponse.Unmarshal(m, b)
}
func (m *EvaluateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateResponse.Merge(m, src)
}
func (m *EvaluateResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateResponse.Size(m)
}
func (m *EvaluateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateResponse proto.InternalMessageInfo

func (m *EvaluateResponse) GetMatchId() string {
	if m != nil {
		return m.MatchId
	}
	return ""
}

func init() {
	proto.RegisterType((*EvaluateRequest)(nil), "openmatch.EvaluateRequest")
	proto.RegisterType((*EvaluateResponse)(nil), "openmatch.EvaluateResponse")
}

func init() { proto.RegisterFile("api/evaluator.proto", fileDescriptor_8c58cb7dff9acb0f) }

var fileDescriptor_8c58cb7dff9acb0f = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xf4, 0xb8, 0x6b, 0xcd, 0x40, 0x65, 0x24, 0x54, 0x0a, 0x42, 0xa6, 0x27, 0xa1,
	0x52, 0xd1, 0xb8, 0xd7, 0xeb, 0x42, 0x11, 0xd2, 0x1d, 0xd0, 0xa1, 0xe8, 0x00, 0xa9, 0x48, 0x0c,
	0x2c, 0xc8, 0x49, 0x1e, 0x49, 0xa0, 0xf1, 0x33, 0x7e, 0x4e, 0x8f, 0x0d, 0x89, 0x99, 0x09, 0x36,
	0x3e, 0x02, 0x2b, 0x1f, 0x85, 0x8d, 0x99, 0x0f, 0x82, 0x92, 0x50, 0x5a, 0xdd, 0xdd, 0x92, 0xc8,
	0xfe, 0x3d, 0xff, 0xff, 0xef, 0xfd, 0x6d, 0x76, 0x55, 0x99, 0x4c, 0xc2, 0x4a, 0x2d, 0x0b, 0xe5,
	0xd0, 0x06, 0xc6, 0xa2, 0x43, 0xde, 0x42, 0x03, 0x3a, 0x57, 0x2e, 0x4a, 0xbb, 0xbc, 0xe4, 0x39,
	0x10, 0xa9, 0x04, 0xa8, 0xc6, 0xdd, 0x9b, 0x09, 0x62, 0xb2, 0x04, 0x59, 0x22, 0xa5, 0x35, 0x3a,
	0xe5, 0x32, 0xd4, 0x6b, 0x7a, 0xaf, 0xfa, 0x45, 0xc3, 0x04, 0xf4, 0x90, 0x4e, 0x55, 0x92, 0x80,
	0x95, 0x68, 0xaa, 0x8a, 0xf3, 0xd5, 0xbd, 0xfb, 0xec, 0xca, 0xac, 0x76, 0x87, 0x05, 0x7c, 0x28,
	0x80, 0x1c, 0xbf, 0xc3, 0x2e, 0x55, 0xde, 0x1d, 0x4f, 0x78, 0xfd, 0xcb, 0xe3, 0x76, 0xf0, 0xbf,
	0x9b, 0xe0, 0x59, 0xf9, 0x5d, 0xd4, 0xb8, 0x77, 0xc8, 0xda, 0x9b, 0xa3, 0x64, 0x50, 0x13, 0xf0,
	0xeb, 0xac, 0x59, 0xc1, 0x37, 0x59, 0xdc, 0xf1, 0x85, 0xd7, 0x6f, 0x2d, 0xf6, 0xaa, 0xf5, 0x3c,
	0x7e, 0xba, 0xd3, 0xf4, 0xda, 0xfe, 0xf8, 0x13, 0x6b, 0xcd, 0xd6, 0xd3, 0x72, 0xcb, 0x9a, 0x6b,
	0x05, 0xde, 0xdd, 0xb2, 0x39, 0xd3, 0x51, 0xf7, 0xc6, 0x85, 0xac, 0xb6, 0xec, 0xdd, 0xfd, 0xfc,
	0xeb, 0xcf, 0x37, 0x7f, 0xbf, 0x77, 0x4b, 0xae, 0x0e, 0x36, 0x49, 0xca, 0xaa, 0x1a, 0x68, 0xfa,
	0x6f, 0x07, 0xa6, 0xde, 0xa0, 0xef, 0x8d, 0xbc, 0x47, 0x5f, 0x1a, 0x5f, 0x8f, 0x7f, 0xfb, 0xfc,
	0xa7, 0xb7, 0xd5, 0x48, 0x6f, 0xce, 0xd8, 0x0b, 0x03, 0x5a, 0x54, 0xf3, 0xf1, 0x6b, 0xa9, 0x73,
	0x86, 0xa6, 0x52, 0x96, 0xae, 0xc3, 0xda, 0x36, 0x86, 0x55, 0x77, 0x7f, 0xb3, 0x1e, 0xc6, 0x19,
	0x45, 0x05, 0xd1, 0x51, 0x7d, 0x15, 0x89, 0xc5, 0xc2, 0x50, 0x10, 0x61, 0x3e, 0x78, 0xc5, 0xf8,
	0xb1, 0x51, 0x51, 0x0a, 0x62, 0x1c, 0x8c, 0xc4, 0x49, 0x16, 0x41, 0x99, 0xcb, 0xd1, 0x5a, 0x32,
	0xc9, 0x5c, 0x5a, 0x84, 0x65, 0xa5, 0xac, 0x8f, 0xbe, 0x45, 0x9b, 0xa8, 0x1c, 0x68, 0xcb, 0x4c,
	0x86, 0x4b, 0x0c, 0x65, 0xae, 0xc8, 0x81, 0x95, 0x27, 0xf3, 0xc7, 0xb3, 0xe7, 0x2f, 0x67, 0xe3,
	0xc6, 0x41, 0x30, 0x1a, 0xf8, 0x9e, 0x3f, 0x6e, 0x2b, 0x63, 0x96, 0x59, 0x54, 0xdd, 0xa2, 0x7c,
	0x47, 0xa8, 0xa7, 0xe7, 0x76, 0x16, 0x0f, 0x58, 0x63, 0x32, 0x9a, 0xf0, 0x09, 0x1b, 0x2c, 0xc0,
	0x15, 0x56, 0x43, 0x2c, 0x4e, 0x53, 0xd0, 0xc2, 0xa5, 0x20, 0x2c, 0x10, 0x16, 0x36, 0x02, 0x11,
	0x23, 0x90, 0xd0, 0xe8, 0x04, 0x7c, 0xcc, 0xc8, 0x05, 0x7c, 0x97, 0xed, 0x7c, 0xf7, 0xbd, 0x3d,
	0xfb, 0x90, 0x75, 0x36, 0x61, 0x88, 0x27, 0x18, 0x15, 0x39, 0xe8, 0xfa, 0xd5, 0xf0, 0xdb, 0x17,
	0x47, 0x23, 0x29, 0x73, 0x20, 0x63, 0x8c, 0x48, 0xbe, 0x16, 0x67, 0xd0, 0xd6, 0x5c, 0xe6, 0x7d,
	0x22, 0x4d, 0xf8, 0xc3, 0x6f, 0x95, 0xfa, 0x95, 0x7c, 0xb8, 0x5b, 0x3d, 0xc3, 0xc3, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0x55, 0x7c, 0x81, 0x08, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EvaluatorClient is the client API for Evaluator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvaluatorClient interface {
	// Evaluate evaluates a list of proposed matches based on quality, collision status, and etc, then shortlist the matches and returns the final results.
	Evaluate(ctx context.Context, opts ...grpc.CallOption) (Evaluator_EvaluateClient, error)
}

type evaluatorClient struct {
	cc *grpc.ClientConn
}

func NewEvaluatorClient(cc *grpc.ClientConn) EvaluatorClient {
	return &evaluatorClient{cc}
}

func (c *evaluatorClient) Evaluate(ctx context.Context, opts ...grpc.CallOption) (Evaluator_EvaluateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Evaluator_serviceDesc.Streams[0], "/openmatch.Evaluator/Evaluate", opts...)
	if err != nil {
		return nil, err
	}
	x := &evaluatorEvaluateClient{stream}
	return x, nil
}

type Evaluator_EvaluateClient interface {
	Send(*EvaluateRequest) error
	Recv() (*EvaluateResponse, error)
	grpc.ClientStream
}

type evaluatorEvaluateClient struct {
	grpc.ClientStream
}

func (x *evaluatorEvaluateClient) Send(m *EvaluateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *evaluatorEvaluateClient) Recv() (*EvaluateResponse, error) {
	m := new(EvaluateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EvaluatorServer is the server API for Evaluator service.
type EvaluatorServer interface {
	// Evaluate evaluates a list of proposed matches based on quality, collision status, and etc, then shortlist the matches and returns the final results.
	Evaluate(Evaluator_EvaluateServer) error
}

// UnimplementedEvaluatorServer can be embedded to have forward compatible implementations.
type UnimplementedEvaluatorServer struct {
}

func (*UnimplementedEvaluatorServer) Evaluate(srv Evaluator_EvaluateServer) error {
	return status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}

func RegisterEvaluatorServer(s *grpc.Server, srv EvaluatorServer) {
	s.RegisterService(&_Evaluator_serviceDesc, srv)
}

func _Evaluator_Evaluate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EvaluatorServer).Evaluate(&evaluatorEvaluateServer{stream})
}

type Evaluator_EvaluateServer interface {
	Send(*EvaluateResponse) error
	Recv() (*EvaluateRequest, error)
	grpc.ServerStream
}

type evaluatorEvaluateServer struct {
	grpc.ServerStream
}

func (x *evaluatorEvaluateServer) Send(m *EvaluateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *evaluatorEvaluateServer) Recv() (*EvaluateRequest, error) {
	m := new(EvaluateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Evaluator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.Evaluator",
	HandlerType: (*EvaluatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Evaluate",
			Handler:       _Evaluator_Evaluate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/evaluator.proto",
}
