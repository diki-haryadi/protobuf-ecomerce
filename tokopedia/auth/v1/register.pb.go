// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tokopedia/auth/v1/register.proto

package authv1

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	RoleId   string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokopedia_auth_v1_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokopedia_auth_v1_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_tokopedia_auth_v1_register_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokopedia_auth_v1_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokopedia_auth_v1_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_tokopedia_auth_v1_register_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_tokopedia_auth_v1_register_proto protoreflect.FileDescriptor

var file_tokopedia_auth_v1_register_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x22, 0x62, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32, 0x66, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x74,
	0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd1, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f,
	0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6b,
	0x69, 0x2d, 0x68, 0x61, 0x72, 0x79, 0x61, 0x64, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2d, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x6f, 0x6b,
	0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x41, 0x58, 0xaa, 0x02, 0x11, 0x54, 0x6f,
	0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x11, 0x54, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x54, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x5c,
	0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x54, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x3a,
	0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tokopedia_auth_v1_register_proto_rawDescOnce sync.Once
	file_tokopedia_auth_v1_register_proto_rawDescData = file_tokopedia_auth_v1_register_proto_rawDesc
)

func file_tokopedia_auth_v1_register_proto_rawDescGZIP() []byte {
	file_tokopedia_auth_v1_register_proto_rawDescOnce.Do(func() {
		file_tokopedia_auth_v1_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_tokopedia_auth_v1_register_proto_rawDescData)
	})
	return file_tokopedia_auth_v1_register_proto_rawDescData
}

var file_tokopedia_auth_v1_register_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tokopedia_auth_v1_register_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),  // 0: tokopedia.auth.v1.RegisterRequest
	(*RegisterResponse)(nil), // 1: tokopedia.auth.v1.RegisterResponse
}
var file_tokopedia_auth_v1_register_proto_depIdxs = []int32{
	0, // 0: tokopedia.auth.v1.RegisterService.Register:input_type -> tokopedia.auth.v1.RegisterRequest
	1, // 1: tokopedia.auth.v1.RegisterService.Register:output_type -> tokopedia.auth.v1.RegisterResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tokopedia_auth_v1_register_proto_init() }
func file_tokopedia_auth_v1_register_proto_init() {
	if File_tokopedia_auth_v1_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tokopedia_auth_v1_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_tokopedia_auth_v1_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_tokopedia_auth_v1_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tokopedia_auth_v1_register_proto_goTypes,
		DependencyIndexes: file_tokopedia_auth_v1_register_proto_depIdxs,
		MessageInfos:      file_tokopedia_auth_v1_register_proto_msgTypes,
	}.Build()
	File_tokopedia_auth_v1_register_proto = out.File
	file_tokopedia_auth_v1_register_proto_rawDesc = nil
	file_tokopedia_auth_v1_register_proto_goTypes = nil
	file_tokopedia_auth_v1_register_proto_depIdxs = nil
}
