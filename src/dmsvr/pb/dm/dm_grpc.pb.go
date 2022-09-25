// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/dm.proto

package dm

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

// DeviceAuthClient is the client API for DeviceAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceAuthClient interface {
	//设备登录认证
	LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error)
	//设备操作认证
	AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error)
	//鉴定是否是root账号
	RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error)
}

type deviceAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceAuthClient(cc grpc.ClientConnInterface) DeviceAuthClient {
	return &deviceAuthClient{cc}
}

func (c *deviceAuthClient) LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.DeviceAuth/loginAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAuthClient) AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.DeviceAuth/accessAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceAuthClient) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.DeviceAuth/rootCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceAuthServer is the server API for DeviceAuth service.
// All implementations must embed UnimplementedDeviceAuthServer
// for forward compatibility
type DeviceAuthServer interface {
	//设备登录认证
	LoginAuth(context.Context, *LoginAuthReq) (*Response, error)
	//设备操作认证
	AccessAuth(context.Context, *AccessAuthReq) (*Response, error)
	//鉴定是否是root账号
	RootCheck(context.Context, *RootCheckReq) (*Response, error)
	mustEmbedUnimplementedDeviceAuthServer()
}

// UnimplementedDeviceAuthServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceAuthServer struct {
}

func (UnimplementedDeviceAuthServer) LoginAuth(context.Context, *LoginAuthReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAuth not implemented")
}
func (UnimplementedDeviceAuthServer) AccessAuth(context.Context, *AccessAuthReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessAuth not implemented")
}
func (UnimplementedDeviceAuthServer) RootCheck(context.Context, *RootCheckReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RootCheck not implemented")
}
func (UnimplementedDeviceAuthServer) mustEmbedUnimplementedDeviceAuthServer() {}

// UnsafeDeviceAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceAuthServer will
// result in compilation errors.
type UnsafeDeviceAuthServer interface {
	mustEmbedUnimplementedDeviceAuthServer()
}

func RegisterDeviceAuthServer(s grpc.ServiceRegistrar, srv DeviceAuthServer) {
	s.RegisterService(&DeviceAuth_ServiceDesc, srv)
}

func _DeviceAuth_LoginAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAuthServer).LoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceAuth/loginAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAuthServer).LoginAuth(ctx, req.(*LoginAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAuth_AccessAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAuthServer).AccessAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceAuth/accessAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAuthServer).AccessAuth(ctx, req.(*AccessAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceAuth_RootCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceAuthServer).RootCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceAuth/rootCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceAuthServer).RootCheck(ctx, req.(*RootCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceAuth_ServiceDesc is the grpc.ServiceDesc for DeviceAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dm.DeviceAuth",
	HandlerType: (*DeviceAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "loginAuth",
			Handler:    _DeviceAuth_LoginAuth_Handler,
		},
		{
			MethodName: "accessAuth",
			Handler:    _DeviceAuth_AccessAuth_Handler,
		},
		{
			MethodName: "rootCheck",
			Handler:    _DeviceAuth_RootCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dm.proto",
}

// DeviceManageClient is the client API for DeviceManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceManageClient interface {
	//新增设备
	DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
	//更新设备
	DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
	//删除设备
	DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	//获取设备信息列表
	DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error)
	//获取设备信息详情
	DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error)
}

type deviceManageClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceManageClient(cc grpc.ClientConnInterface) DeviceManageClient {
	return &deviceManageClient{cc}
}

func (c *deviceManageClient) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.DeviceManage/DeviceInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManageClient) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.DeviceManage/DeviceInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManageClient) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.DeviceManage/DeviceInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManageClient) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	out := new(DeviceInfoIndexResp)
	err := c.cc.Invoke(ctx, "/dm.DeviceManage/DeviceInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManageClient) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	out := new(DeviceInfo)
	err := c.cc.Invoke(ctx, "/dm.DeviceManage/DeviceInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceManageServer is the server API for DeviceManage service.
// All implementations must embed UnimplementedDeviceManageServer
// for forward compatibility
type DeviceManageServer interface {
	//新增设备
	DeviceInfoCreate(context.Context, *DeviceInfo) (*Response, error)
	//更新设备
	DeviceInfoUpdate(context.Context, *DeviceInfo) (*Response, error)
	//删除设备
	DeviceInfoDelete(context.Context, *DeviceInfoDeleteReq) (*Response, error)
	//获取设备信息列表
	DeviceInfoIndex(context.Context, *DeviceInfoIndexReq) (*DeviceInfoIndexResp, error)
	//获取设备信息详情
	DeviceInfoRead(context.Context, *DeviceInfoReadReq) (*DeviceInfo, error)
	mustEmbedUnimplementedDeviceManageServer()
}

// UnimplementedDeviceManageServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceManageServer struct {
}

func (UnimplementedDeviceManageServer) DeviceInfoCreate(context.Context, *DeviceInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoCreate not implemented")
}
func (UnimplementedDeviceManageServer) DeviceInfoUpdate(context.Context, *DeviceInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoUpdate not implemented")
}
func (UnimplementedDeviceManageServer) DeviceInfoDelete(context.Context, *DeviceInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoDelete not implemented")
}
func (UnimplementedDeviceManageServer) DeviceInfoIndex(context.Context, *DeviceInfoIndexReq) (*DeviceInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoIndex not implemented")
}
func (UnimplementedDeviceManageServer) DeviceInfoRead(context.Context, *DeviceInfoReadReq) (*DeviceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoRead not implemented")
}
func (UnimplementedDeviceManageServer) mustEmbedUnimplementedDeviceManageServer() {}

// UnsafeDeviceManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceManageServer will
// result in compilation errors.
type UnsafeDeviceManageServer interface {
	mustEmbedUnimplementedDeviceManageServer()
}

func RegisterDeviceManageServer(s grpc.ServiceRegistrar, srv DeviceManageServer) {
	s.RegisterService(&DeviceManage_ServiceDesc, srv)
}

func _DeviceManage_DeviceInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManageServer).DeviceInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceManage/DeviceInfoCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManageServer).DeviceInfoCreate(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManage_DeviceInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManageServer).DeviceInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceManage/DeviceInfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManageServer).DeviceInfoUpdate(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManage_DeviceInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManageServer).DeviceInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceManage/DeviceInfoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManageServer).DeviceInfoDelete(ctx, req.(*DeviceInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManage_DeviceInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManageServer).DeviceInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceManage/DeviceInfoIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManageServer).DeviceInfoIndex(ctx, req.(*DeviceInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManage_DeviceInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfoReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManageServer).DeviceInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.DeviceManage/DeviceInfoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManageServer).DeviceInfoRead(ctx, req.(*DeviceInfoReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceManage_ServiceDesc is the grpc.ServiceDesc for DeviceManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dm.DeviceManage",
	HandlerType: (*DeviceManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceInfoCreate",
			Handler:    _DeviceManage_DeviceInfoCreate_Handler,
		},
		{
			MethodName: "DeviceInfoUpdate",
			Handler:    _DeviceManage_DeviceInfoUpdate_Handler,
		},
		{
			MethodName: "DeviceInfoDelete",
			Handler:    _DeviceManage_DeviceInfoDelete_Handler,
		},
		{
			MethodName: "DeviceInfoIndex",
			Handler:    _DeviceManage_DeviceInfoIndex_Handler,
		},
		{
			MethodName: "DeviceInfoRead",
			Handler:    _DeviceManage_DeviceInfoRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dm.proto",
}

// ProductManageClient is the client API for ProductManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductManageClient interface {
	//新增产品
	ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
	//更新产品
	ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
	//删除产品
	ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	//获取产品信息列表
	ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error)
	//获取产品信息详情
	ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error)
	//更新产品物模型
	ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error)
	//新增产品
	ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Response, error)
	//删除产品
	ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Response, error)
	//获取产品信息列表
	ProductSchemaIndex(ctx context.Context, in *ProductSchemaIndexReq, opts ...grpc.CallOption) (*ProductSchemaIndexResp, error)
	//删除产品
	ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Response, error)
	//获取产品信息列表
	ProductSchemaTslRead(ctx context.Context, in *ProductSchemaTslReadReq, opts ...grpc.CallOption) (*ProductSchemaTslReadResp, error)
}

type productManageClient struct {
	cc grpc.ClientConnInterface
}

func NewProductManageClient(cc grpc.ClientConnInterface) ProductManageClient {
	return &productManageClient{cc}
}

func (c *productManageClient) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	out := new(ProductInfoIndexResp)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductSchemaUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductSchemaCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductSchemaDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductSchemaIndex(ctx context.Context, in *ProductSchemaIndexReq, opts ...grpc.CallOption) (*ProductSchemaIndexResp, error) {
	out := new(ProductSchemaIndexResp)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductSchemaIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductSchemaTslImport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManageClient) ProductSchemaTslRead(ctx context.Context, in *ProductSchemaTslReadReq, opts ...grpc.CallOption) (*ProductSchemaTslReadResp, error) {
	out := new(ProductSchemaTslReadResp)
	err := c.cc.Invoke(ctx, "/dm.ProductManage/ProductSchemaTslRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductManageServer is the server API for ProductManage service.
// All implementations must embed UnimplementedProductManageServer
// for forward compatibility
type ProductManageServer interface {
	//新增产品
	ProductInfoCreate(context.Context, *ProductInfo) (*Response, error)
	//更新产品
	ProductInfoUpdate(context.Context, *ProductInfo) (*Response, error)
	//删除产品
	ProductInfoDelete(context.Context, *ProductInfoDeleteReq) (*Response, error)
	//获取产品信息列表
	ProductInfoIndex(context.Context, *ProductInfoIndexReq) (*ProductInfoIndexResp, error)
	//获取产品信息详情
	ProductInfoRead(context.Context, *ProductInfoReadReq) (*ProductInfo, error)
	//更新产品物模型
	ProductSchemaUpdate(context.Context, *ProductSchemaUpdateReq) (*Response, error)
	//新增产品
	ProductSchemaCreate(context.Context, *ProductSchemaCreateReq) (*Response, error)
	//删除产品
	ProductSchemaDelete(context.Context, *ProductSchemaDeleteReq) (*Response, error)
	//获取产品信息列表
	ProductSchemaIndex(context.Context, *ProductSchemaIndexReq) (*ProductSchemaIndexResp, error)
	//删除产品
	ProductSchemaTslImport(context.Context, *ProductSchemaTslImportReq) (*Response, error)
	//获取产品信息列表
	ProductSchemaTslRead(context.Context, *ProductSchemaTslReadReq) (*ProductSchemaTslReadResp, error)
	mustEmbedUnimplementedProductManageServer()
}

// UnimplementedProductManageServer must be embedded to have forward compatible implementations.
type UnimplementedProductManageServer struct {
}

func (UnimplementedProductManageServer) ProductInfoCreate(context.Context, *ProductInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoCreate not implemented")
}
func (UnimplementedProductManageServer) ProductInfoUpdate(context.Context, *ProductInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoUpdate not implemented")
}
func (UnimplementedProductManageServer) ProductInfoDelete(context.Context, *ProductInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoDelete not implemented")
}
func (UnimplementedProductManageServer) ProductInfoIndex(context.Context, *ProductInfoIndexReq) (*ProductInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoIndex not implemented")
}
func (UnimplementedProductManageServer) ProductInfoRead(context.Context, *ProductInfoReadReq) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoRead not implemented")
}
func (UnimplementedProductManageServer) ProductSchemaUpdate(context.Context, *ProductSchemaUpdateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaUpdate not implemented")
}
func (UnimplementedProductManageServer) ProductSchemaCreate(context.Context, *ProductSchemaCreateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaCreate not implemented")
}
func (UnimplementedProductManageServer) ProductSchemaDelete(context.Context, *ProductSchemaDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaDelete not implemented")
}
func (UnimplementedProductManageServer) ProductSchemaIndex(context.Context, *ProductSchemaIndexReq) (*ProductSchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaIndex not implemented")
}
func (UnimplementedProductManageServer) ProductSchemaTslImport(context.Context, *ProductSchemaTslImportReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaTslImport not implemented")
}
func (UnimplementedProductManageServer) ProductSchemaTslRead(context.Context, *ProductSchemaTslReadReq) (*ProductSchemaTslReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaTslRead not implemented")
}
func (UnimplementedProductManageServer) mustEmbedUnimplementedProductManageServer() {}

// UnsafeProductManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductManageServer will
// result in compilation errors.
type UnsafeProductManageServer interface {
	mustEmbedUnimplementedProductManageServer()
}

func RegisterProductManageServer(s grpc.ServiceRegistrar, srv ProductManageServer) {
	s.RegisterService(&ProductManage_ServiceDesc, srv)
}

func _ProductManage_ProductInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductInfoCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductInfoCreate(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductInfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductInfoUpdate(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductInfoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductInfoDelete(ctx, req.(*ProductInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductInfoIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductInfoIndex(ctx, req.(*ProductInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductInfoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductInfoRead(ctx, req.(*ProductInfoReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductSchemaUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchemaUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductSchemaUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductSchemaUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductSchemaUpdate(ctx, req.(*ProductSchemaUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductSchemaCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchemaCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductSchemaCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductSchemaCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductSchemaCreate(ctx, req.(*ProductSchemaCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductSchemaDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchemaDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductSchemaDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductSchemaDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductSchemaDelete(ctx, req.(*ProductSchemaDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductSchemaIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchemaIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductSchemaIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductSchemaIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductSchemaIndex(ctx, req.(*ProductSchemaIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductSchemaTslImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchemaTslImportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductSchemaTslImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductSchemaTslImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductSchemaTslImport(ctx, req.(*ProductSchemaTslImportReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManage_ProductSchemaTslRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchemaTslReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManageServer).ProductSchemaTslRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.ProductManage/ProductSchemaTslRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManageServer).ProductSchemaTslRead(ctx, req.(*ProductSchemaTslReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductManage_ServiceDesc is the grpc.ServiceDesc for ProductManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dm.ProductManage",
	HandlerType: (*ProductManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductInfoCreate",
			Handler:    _ProductManage_ProductInfoCreate_Handler,
		},
		{
			MethodName: "ProductInfoUpdate",
			Handler:    _ProductManage_ProductInfoUpdate_Handler,
		},
		{
			MethodName: "ProductInfoDelete",
			Handler:    _ProductManage_ProductInfoDelete_Handler,
		},
		{
			MethodName: "ProductInfoIndex",
			Handler:    _ProductManage_ProductInfoIndex_Handler,
		},
		{
			MethodName: "ProductInfoRead",
			Handler:    _ProductManage_ProductInfoRead_Handler,
		},
		{
			MethodName: "ProductSchemaUpdate",
			Handler:    _ProductManage_ProductSchemaUpdate_Handler,
		},
		{
			MethodName: "ProductSchemaCreate",
			Handler:    _ProductManage_ProductSchemaCreate_Handler,
		},
		{
			MethodName: "ProductSchemaDelete",
			Handler:    _ProductManage_ProductSchemaDelete_Handler,
		},
		{
			MethodName: "ProductSchemaIndex",
			Handler:    _ProductManage_ProductSchemaIndex_Handler,
		},
		{
			MethodName: "ProductSchemaTslImport",
			Handler:    _ProductManage_ProductSchemaTslImport_Handler,
		},
		{
			MethodName: "ProductSchemaTslRead",
			Handler:    _ProductManage_ProductSchemaTslRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dm.proto",
}
