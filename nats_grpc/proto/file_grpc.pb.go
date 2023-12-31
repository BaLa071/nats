// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/file.proto

package BaLa071

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

// FileTransferServiceClient is the client API for FileTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileTransferServiceClient interface {
	TransferFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type fileTransferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileTransferServiceClient(cc grpc.ClientConnInterface) FileTransferServiceClient {
	return &fileTransferServiceClient{cc}
}

func (c *fileTransferServiceClient) TransferFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/filetransfer.FileTransferService/transferFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileTransferServiceServer is the server API for FileTransferService service.
// All implementations must embed UnimplementedFileTransferServiceServer
// for forward compatibility
type FileTransferServiceServer interface {
	TransferFile(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedFileTransferServiceServer()
}

// UnimplementedFileTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileTransferServiceServer struct {
}

func (UnimplementedFileTransferServiceServer) TransferFile(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferFile not implemented")
}
func (UnimplementedFileTransferServiceServer) mustEmbedUnimplementedFileTransferServiceServer() {}

// UnsafeFileTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileTransferServiceServer will
// result in compilation errors.
type UnsafeFileTransferServiceServer interface {
	mustEmbedUnimplementedFileTransferServiceServer()
}

func RegisterFileTransferServiceServer(s grpc.ServiceRegistrar, srv FileTransferServiceServer) {
	s.RegisterService(&FileTransferService_ServiceDesc, srv)
}

func _FileTransferService_TransferFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTransferServiceServer).TransferFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filetransfer.FileTransferService/transferFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTransferServiceServer).TransferFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// FileTransferService_ServiceDesc is the grpc.ServiceDesc for FileTransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileTransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filetransfer.FileTransferService",
	HandlerType: (*FileTransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "transferFile",
			Handler:    _FileTransferService_TransferFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/file.proto",
}
