// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PartyUIClient is the client API for PartyUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartyUIClient interface {
	// ListEngineSpecs returns a list of Party Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (PartyUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type partyUIClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyUIClient(cc grpc.ClientConnInterface) PartyUIClient {
	return &partyUIClient{cc}
}

func (c *partyUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (PartyUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PartyUI_ServiceDesc.Streams[0], "/v1.PartyUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &partyUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PartyUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type partyUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *partyUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *partyUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.PartyUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyUIServer is the server API for PartyUI service.
// All implementations must embed UnimplementedPartyUIServer
// for forward compatibility
type PartyUIServer interface {
	// ListEngineSpecs returns a list of Party Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, PartyUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedPartyUIServer()
}

// UnimplementedPartyUIServer must be embedded to have forward compatible implementations.
type UnimplementedPartyUIServer struct {
}

func (UnimplementedPartyUIServer) ListEngineSpecs(*ListEngineSpecsRequest, PartyUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedPartyUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedPartyUIServer) mustEmbedUnimplementedPartyUIServer() {}

// UnsafePartyUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyUIServer will
// result in compilation errors.
type UnsafePartyUIServer interface {
	mustEmbedUnimplementedPartyUIServer()
}

func RegisterPartyUIServer(s grpc.ServiceRegistrar, srv PartyUIServer) {
	s.RegisterService(&PartyUI_ServiceDesc, srv)
}

func _PartyUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PartyUIServer).ListEngineSpecs(m, &partyUIListEngineSpecsServer{stream})
}

type PartyUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type partyUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *partyUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PartyUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PartyUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyUI_ServiceDesc is the grpc.ServiceDesc for PartyUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PartyUI",
	HandlerType: (*PartyUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _PartyUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _PartyUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "party-ui.proto",
}
