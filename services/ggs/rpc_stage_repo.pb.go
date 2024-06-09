// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: rpc_stage_repo.proto

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

type AppFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Mode     int32  `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	ModeTime string `protobuf:"bytes,4,opt,name=modeTime,proto3" json:"modeTime,omitempty"`
	CTime    string `protobuf:"bytes,5,opt,name=cTime,proto3" json:"cTime,omitempty"`
}

func (x *AppFile) Reset() {
	*x = AppFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stage_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppFile) ProtoMessage() {}

func (x *AppFile) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stage_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppFile.ProtoReflect.Descriptor instead.
func (*AppFile) Descriptor() ([]byte, []int) {
	return file_rpc_stage_repo_proto_rawDescGZIP(), []int{0}
}

func (x *AppFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AppFile) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AppFile) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *AppFile) GetModeTime() string {
	if x != nil {
		return x.ModeTime
	}
	return ""
}

func (x *AppFile) GetCTime() string {
	if x != nil {
		return x.CTime
	}
	return ""
}

type StageToRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId string     `protobuf:"bytes,1,opt,name=repoId,proto3" json:"repoId,omitempty"`
	Files  []*AppFile `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *StageToRepoRequest) Reset() {
	*x = StageToRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stage_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageToRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageToRepoRequest) ProtoMessage() {}

func (x *StageToRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stage_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageToRepoRequest.ProtoReflect.Descriptor instead.
func (*StageToRepoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_stage_repo_proto_rawDescGZIP(), []int{1}
}

func (x *StageToRepoRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *StageToRepoRequest) GetFiles() []*AppFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type StageToRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StageToRepoResponse) Reset() {
	*x = StageToRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stage_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageToRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageToRepoResponse) ProtoMessage() {}

func (x *StageToRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stage_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageToRepoResponse.ProtoReflect.Descriptor instead.
func (*StageToRepoResponse) Descriptor() ([]byte, []int) {
	return file_rpc_stage_repo_proto_rawDescGZIP(), []int{2}
}

func (x *StageToRepoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_stage_repo_proto protoreflect.FileDescriptor

var file_rpc_stage_repo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x67, 0x73, 0x22, 0x7d, 0x0a, 0x07, 0x41,
	0x70, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x67, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x13,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x69, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_stage_repo_proto_rawDescOnce sync.Once
	file_rpc_stage_repo_proto_rawDescData = file_rpc_stage_repo_proto_rawDesc
)

func file_rpc_stage_repo_proto_rawDescGZIP() []byte {
	file_rpc_stage_repo_proto_rawDescOnce.Do(func() {
		file_rpc_stage_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_stage_repo_proto_rawDescData)
	})
	return file_rpc_stage_repo_proto_rawDescData
}

var file_rpc_stage_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_stage_repo_proto_goTypes = []interface{}{
	(*AppFile)(nil),             // 0: ggs.AppFile
	(*StageToRepoRequest)(nil),  // 1: ggs.StageToRepoRequest
	(*StageToRepoResponse)(nil), // 2: ggs.StageToRepoResponse
}
var file_rpc_stage_repo_proto_depIdxs = []int32{
	0, // 0: ggs.StageToRepoRequest.files:type_name -> ggs.AppFile
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_stage_repo_proto_init() }
func file_rpc_stage_repo_proto_init() {
	if File_rpc_stage_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_stage_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppFile); i {
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
		file_rpc_stage_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageToRepoRequest); i {
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
		file_rpc_stage_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageToRepoResponse); i {
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
			RawDescriptor: file_rpc_stage_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_stage_repo_proto_goTypes,
		DependencyIndexes: file_rpc_stage_repo_proto_depIdxs,
		MessageInfos:      file_rpc_stage_repo_proto_msgTypes,
	}.Build()
	File_rpc_stage_repo_proto = out.File
	file_rpc_stage_repo_proto_rawDesc = nil
	file_rpc_stage_repo_proto_goTypes = nil
	file_rpc_stage_repo_proto_depIdxs = nil
}
