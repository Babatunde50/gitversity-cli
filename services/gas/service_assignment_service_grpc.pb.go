// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: service_assignment_service.proto

package gas

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AssignmentService_CreateCourse_FullMethodName              = "/gas.AssignmentService/CreateCourse"
	AssignmentService_JoinAssignment_FullMethodName            = "/gas.AssignmentService/JoinAssignment"
	AssignmentService_CreateAssignment_FullMethodName          = "/gas.AssignmentService/CreateAssignment"
	AssignmentService_ListCourses_FullMethodName               = "/gas.AssignmentService/ListCourses"
	AssignmentService_UpdateAssignment_FullMethodName          = "/gas.AssignmentService/UpdateAssignment"
	AssignmentService_GetAssignment_FullMethodName             = "/gas.AssignmentService/GetAssignment"
	AssignmentService_GetUserAssignments_FullMethodName        = "/gas.AssignmentService/GetUserAssignments"
	AssignmentService_GetAssignmentByInviteCode_FullMethodName = "/gas.AssignmentService/GetAssignmentByInviteCode"
	AssignmentService_ListAssignments_FullMethodName           = "/gas.AssignmentService/ListAssignments"
	AssignmentService_DeleteAssignment_FullMethodName          = "/gas.AssignmentService/DeleteAssignment"
	AssignmentService_SubmitAssignment_FullMethodName          = "/gas.AssignmentService/SubmitAssignment"
	AssignmentService_ListSubmissions_FullMethodName           = "/gas.AssignmentService/ListSubmissions"
	AssignmentService_GetSubmission_FullMethodName             = "/gas.AssignmentService/GetSubmission"
)

// AssignmentServiceClient is the client API for AssignmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssignmentServiceClient interface {
	// Create an new course
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error)
	JoinAssignment(ctx context.Context, in *JoinAssignmentRequest, opts ...grpc.CallOption) (*JoinAssignmentResponse, error)
	// Create a new assignment
	CreateAssignment(ctx context.Context, in *CreateAssignmentRequest, opts ...grpc.CallOption) (*Assignment, error)
	// Get courses
	ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (*ListCoursesResponse, error)
	// Update Assignment
	UpdateAssignment(ctx context.Context, in *UpdateAssignmentRequest, opts ...grpc.CallOption) (*UpdateAssignmentResponse, error)
	// Get a specific assignment by ID
	GetAssignment(ctx context.Context, in *GetAssignmentRequest, opts ...grpc.CallOption) (*Assignment, error)
	GetUserAssignments(ctx context.Context, in *GetAssignmentRequest, opts ...grpc.CallOption) (*UserAssignments, error)
	// Get a specific assignment by inviteCode
	GetAssignmentByInviteCode(ctx context.Context, in *GetAssignmentByInviteCodeRequest, opts ...grpc.CallOption) (*Assignment, error)
	// List all assignments
	ListAssignments(ctx context.Context, in *ListAssignmentsRequest, opts ...grpc.CallOption) (*ListAssignmentsResponse, error)
	// Delete an assignment
	DeleteAssignment(ctx context.Context, in *DeleteAssignmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Submit an assignment
	SubmitAssignment(ctx context.Context, in *SubmitAssignmentRequest, opts ...grpc.CallOption) (*Submission, error)
	// List submissions for an assignment
	ListSubmissions(ctx context.Context, in *ListSubmissionsRequest, opts ...grpc.CallOption) (*ListSubmissionsResponse, error)
	// Get a specific submission
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*SubmissionDetail, error)
}

type assignmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssignmentServiceClient(cc grpc.ClientConnInterface) AssignmentServiceClient {
	return &assignmentServiceClient{cc}
}

func (c *assignmentServiceClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, AssignmentService_CreateCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) JoinAssignment(ctx context.Context, in *JoinAssignmentRequest, opts ...grpc.CallOption) (*JoinAssignmentResponse, error) {
	out := new(JoinAssignmentResponse)
	err := c.cc.Invoke(ctx, AssignmentService_JoinAssignment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) CreateAssignment(ctx context.Context, in *CreateAssignmentRequest, opts ...grpc.CallOption) (*Assignment, error) {
	out := new(Assignment)
	err := c.cc.Invoke(ctx, AssignmentService_CreateAssignment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (*ListCoursesResponse, error) {
	out := new(ListCoursesResponse)
	err := c.cc.Invoke(ctx, AssignmentService_ListCourses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) UpdateAssignment(ctx context.Context, in *UpdateAssignmentRequest, opts ...grpc.CallOption) (*UpdateAssignmentResponse, error) {
	out := new(UpdateAssignmentResponse)
	err := c.cc.Invoke(ctx, AssignmentService_UpdateAssignment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) GetAssignment(ctx context.Context, in *GetAssignmentRequest, opts ...grpc.CallOption) (*Assignment, error) {
	out := new(Assignment)
	err := c.cc.Invoke(ctx, AssignmentService_GetAssignment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) GetUserAssignments(ctx context.Context, in *GetAssignmentRequest, opts ...grpc.CallOption) (*UserAssignments, error) {
	out := new(UserAssignments)
	err := c.cc.Invoke(ctx, AssignmentService_GetUserAssignments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) GetAssignmentByInviteCode(ctx context.Context, in *GetAssignmentByInviteCodeRequest, opts ...grpc.CallOption) (*Assignment, error) {
	out := new(Assignment)
	err := c.cc.Invoke(ctx, AssignmentService_GetAssignmentByInviteCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) ListAssignments(ctx context.Context, in *ListAssignmentsRequest, opts ...grpc.CallOption) (*ListAssignmentsResponse, error) {
	out := new(ListAssignmentsResponse)
	err := c.cc.Invoke(ctx, AssignmentService_ListAssignments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) DeleteAssignment(ctx context.Context, in *DeleteAssignmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AssignmentService_DeleteAssignment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) SubmitAssignment(ctx context.Context, in *SubmitAssignmentRequest, opts ...grpc.CallOption) (*Submission, error) {
	out := new(Submission)
	err := c.cc.Invoke(ctx, AssignmentService_SubmitAssignment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) ListSubmissions(ctx context.Context, in *ListSubmissionsRequest, opts ...grpc.CallOption) (*ListSubmissionsResponse, error) {
	out := new(ListSubmissionsResponse)
	err := c.cc.Invoke(ctx, AssignmentService_ListSubmissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentServiceClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*SubmissionDetail, error) {
	out := new(SubmissionDetail)
	err := c.cc.Invoke(ctx, AssignmentService_GetSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignmentServiceServer is the server API for AssignmentService service.
// All implementations should embed UnimplementedAssignmentServiceServer
// for forward compatibility
type AssignmentServiceServer interface {
	// Create an new course
	CreateCourse(context.Context, *CreateCourseRequest) (*Course, error)
	JoinAssignment(context.Context, *JoinAssignmentRequest) (*JoinAssignmentResponse, error)
	// Create a new assignment
	CreateAssignment(context.Context, *CreateAssignmentRequest) (*Assignment, error)
	// Get courses
	ListCourses(context.Context, *ListCoursesRequest) (*ListCoursesResponse, error)
	// Update Assignment
	UpdateAssignment(context.Context, *UpdateAssignmentRequest) (*UpdateAssignmentResponse, error)
	// Get a specific assignment by ID
	GetAssignment(context.Context, *GetAssignmentRequest) (*Assignment, error)
	GetUserAssignments(context.Context, *GetAssignmentRequest) (*UserAssignments, error)
	// Get a specific assignment by inviteCode
	GetAssignmentByInviteCode(context.Context, *GetAssignmentByInviteCodeRequest) (*Assignment, error)
	// List all assignments
	ListAssignments(context.Context, *ListAssignmentsRequest) (*ListAssignmentsResponse, error)
	// Delete an assignment
	DeleteAssignment(context.Context, *DeleteAssignmentRequest) (*emptypb.Empty, error)
	// Submit an assignment
	SubmitAssignment(context.Context, *SubmitAssignmentRequest) (*Submission, error)
	// List submissions for an assignment
	ListSubmissions(context.Context, *ListSubmissionsRequest) (*ListSubmissionsResponse, error)
	// Get a specific submission
	GetSubmission(context.Context, *GetSubmissionRequest) (*SubmissionDetail, error)
}

// UnimplementedAssignmentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAssignmentServiceServer struct {
}

func (UnimplementedAssignmentServiceServer) CreateCourse(context.Context, *CreateCourseRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedAssignmentServiceServer) JoinAssignment(context.Context, *JoinAssignmentRequest) (*JoinAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinAssignment not implemented")
}
func (UnimplementedAssignmentServiceServer) CreateAssignment(context.Context, *CreateAssignmentRequest) (*Assignment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssignment not implemented")
}
func (UnimplementedAssignmentServiceServer) ListCourses(context.Context, *ListCoursesRequest) (*ListCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedAssignmentServiceServer) UpdateAssignment(context.Context, *UpdateAssignmentRequest) (*UpdateAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAssignment not implemented")
}
func (UnimplementedAssignmentServiceServer) GetAssignment(context.Context, *GetAssignmentRequest) (*Assignment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignment not implemented")
}
func (UnimplementedAssignmentServiceServer) GetUserAssignments(context.Context, *GetAssignmentRequest) (*UserAssignments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAssignments not implemented")
}
func (UnimplementedAssignmentServiceServer) GetAssignmentByInviteCode(context.Context, *GetAssignmentByInviteCodeRequest) (*Assignment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignmentByInviteCode not implemented")
}
func (UnimplementedAssignmentServiceServer) ListAssignments(context.Context, *ListAssignmentsRequest) (*ListAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssignments not implemented")
}
func (UnimplementedAssignmentServiceServer) DeleteAssignment(context.Context, *DeleteAssignmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssignment not implemented")
}
func (UnimplementedAssignmentServiceServer) SubmitAssignment(context.Context, *SubmitAssignmentRequest) (*Submission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAssignment not implemented")
}
func (UnimplementedAssignmentServiceServer) ListSubmissions(context.Context, *ListSubmissionsRequest) (*ListSubmissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubmissions not implemented")
}
func (UnimplementedAssignmentServiceServer) GetSubmission(context.Context, *GetSubmissionRequest) (*SubmissionDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}

// UnsafeAssignmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssignmentServiceServer will
// result in compilation errors.
type UnsafeAssignmentServiceServer interface {
	mustEmbedUnimplementedAssignmentServiceServer()
}

func RegisterAssignmentServiceServer(s grpc.ServiceRegistrar, srv AssignmentServiceServer) {
	s.RegisterService(&AssignmentService_ServiceDesc, srv)
}

func _AssignmentService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_CreateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_JoinAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).JoinAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_JoinAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).JoinAssignment(ctx, req.(*JoinAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_CreateAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).CreateAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_CreateAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).CreateAssignment(ctx, req.(*CreateAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_ListCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).ListCourses(ctx, req.(*ListCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_UpdateAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).UpdateAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_UpdateAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).UpdateAssignment(ctx, req.(*UpdateAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_GetAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).GetAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_GetAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).GetAssignment(ctx, req.(*GetAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_GetUserAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).GetUserAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_GetUserAssignments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).GetUserAssignments(ctx, req.(*GetAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_GetAssignmentByInviteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssignmentByInviteCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).GetAssignmentByInviteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_GetAssignmentByInviteCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).GetAssignmentByInviteCode(ctx, req.(*GetAssignmentByInviteCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_ListAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).ListAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_ListAssignments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).ListAssignments(ctx, req.(*ListAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_DeleteAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).DeleteAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_DeleteAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).DeleteAssignment(ctx, req.(*DeleteAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_SubmitAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).SubmitAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_SubmitAssignment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).SubmitAssignment(ctx, req.(*SubmitAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_ListSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubmissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).ListSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_ListSubmissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).ListSubmissions(ctx, req.(*ListSubmissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssignmentService_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServiceServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssignmentService_GetSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServiceServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssignmentService_ServiceDesc is the grpc.ServiceDesc for AssignmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssignmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gas.AssignmentService",
	HandlerType: (*AssignmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourse",
			Handler:    _AssignmentService_CreateCourse_Handler,
		},
		{
			MethodName: "JoinAssignment",
			Handler:    _AssignmentService_JoinAssignment_Handler,
		},
		{
			MethodName: "CreateAssignment",
			Handler:    _AssignmentService_CreateAssignment_Handler,
		},
		{
			MethodName: "ListCourses",
			Handler:    _AssignmentService_ListCourses_Handler,
		},
		{
			MethodName: "UpdateAssignment",
			Handler:    _AssignmentService_UpdateAssignment_Handler,
		},
		{
			MethodName: "GetAssignment",
			Handler:    _AssignmentService_GetAssignment_Handler,
		},
		{
			MethodName: "GetUserAssignments",
			Handler:    _AssignmentService_GetUserAssignments_Handler,
		},
		{
			MethodName: "GetAssignmentByInviteCode",
			Handler:    _AssignmentService_GetAssignmentByInviteCode_Handler,
		},
		{
			MethodName: "ListAssignments",
			Handler:    _AssignmentService_ListAssignments_Handler,
		},
		{
			MethodName: "DeleteAssignment",
			Handler:    _AssignmentService_DeleteAssignment_Handler,
		},
		{
			MethodName: "SubmitAssignment",
			Handler:    _AssignmentService_SubmitAssignment_Handler,
		},
		{
			MethodName: "ListSubmissions",
			Handler:    _AssignmentService_ListSubmissions_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _AssignmentService_GetSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_assignment_service.proto",
}
