// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: rpc_get_repo.proto

package ggs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId string `protobuf:"bytes,1,opt,name=repoId,proto3" json:"repoId,omitempty"`
}

func (x *GetRepoRequest) Reset() {
	*x = GetRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepoRequest) ProtoMessage() {}

func (x *GetRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepoRequest.ProtoReflect.Descriptor instead.
func (*GetRepoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_repo_proto_rawDescGZIP(), []int{0}
}

func (x *GetRepoRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

type GetRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo *Repo `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *GetRepoResponse) Reset() {
	*x = GetRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepoResponse) ProtoMessage() {}

func (x *GetRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepoResponse.ProtoReflect.Descriptor instead.
func (*GetRepoResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_repo_proto_rawDescGZIP(), []int{1}
}

func (x *GetRepoResponse) GetRepo() *Repo {
	if x != nil {
		return x.Repo
	}
	return nil
}

var File_rpc_get_repo_proto protoreflect.FileDescriptor

var file_rpc_get_repo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x67, 0x73, 0x1a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x22,
	0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x67, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x04, 0x72, 0x65, 0x70,
	0x6f, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x69, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x69, 0x74, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_get_repo_proto_rawDescOnce sync.Once
	file_rpc_get_repo_proto_rawDescData = file_rpc_get_repo_proto_rawDesc
)

func file_rpc_get_repo_proto_rawDescGZIP() []byte {
	file_rpc_get_repo_proto_rawDescOnce.Do(func() {
		file_rpc_get_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_repo_proto_rawDescData)
	})
	return file_rpc_get_repo_proto_rawDescData
}

var file_rpc_get_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_get_repo_proto_goTypes = []interface{}{
	(*GetRepoRequest)(nil),  // 0: ggs.GetRepoRequest
	(*GetRepoResponse)(nil), // 1: ggs.GetRepoResponse
	(*Repo)(nil),            // 2: ggs.Repo
}
var file_rpc_get_repo_proto_depIdxs = []int32{
	2, // 0: ggs.GetRepoResponse.repo:type_name -> ggs.Repo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_get_repo_proto_init() }
func file_rpc_get_repo_proto_init() {
	if File_rpc_get_repo_proto != nil {
		return
	}
	file_repo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_get_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_get_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_repo_proto_goTypes,
		DependencyIndexes: file_rpc_get_repo_proto_depIdxs,
		MessageInfos:      file_rpc_get_repo_proto_msgTypes,
	}.Build()
	File_rpc_get_repo_proto = out.File
	file_rpc_get_repo_proto_rawDesc = nil
	file_rpc_get_repo_proto_goTypes = nil
	file_rpc_get_repo_proto_depIdxs = nil
}
