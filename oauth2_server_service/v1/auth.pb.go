// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: oauth2_server_service/v1/auth.proto

package oauth2_server_servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_oauth2_server_service_v1_auth_proto protoreflect.FileDescriptor

var file_oauth2_server_service_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x27, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcf, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x46,
	0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x73, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x2f, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xff, 0x01, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6b, 0x69, 0x2d, 0x68, 0x61, 0x72, 0x79, 0x61, 0x64,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x32, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x16, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x32, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22,
	0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x17, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_oauth2_server_service_v1_auth_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),        // 0: oauth2_server_service.v1.RegisterRequest
	(*ForgotPasswordRequest)(nil),  // 1: oauth2_server_service.v1.ForgotPasswordRequest
	(*ChangePasswordRequest)(nil),  // 2: oauth2_server_service.v1.ChangePasswordRequest
	(*UpdateUsernameRequest)(nil),  // 3: oauth2_server_service.v1.UpdateUsernameRequest
	(*RegisterResponse)(nil),       // 4: oauth2_server_service.v1.RegisterResponse
	(*ForgotPasswordResponse)(nil), // 5: oauth2_server_service.v1.ForgotPasswordResponse
	(*ChangePasswordResponse)(nil), // 6: oauth2_server_service.v1.ChangePasswordResponse
	(*UpdateUsernameResponse)(nil), // 7: oauth2_server_service.v1.UpdateUsernameResponse
}
var file_oauth2_server_service_v1_auth_proto_depIdxs = []int32{
	0, // 0: oauth2_server_service.v1.AuthService.Register:input_type -> oauth2_server_service.v1.RegisterRequest
	1, // 1: oauth2_server_service.v1.AuthService.ForgotPassword:input_type -> oauth2_server_service.v1.ForgotPasswordRequest
	2, // 2: oauth2_server_service.v1.AuthService.ChangePassword:input_type -> oauth2_server_service.v1.ChangePasswordRequest
	3, // 3: oauth2_server_service.v1.AuthService.UpdateUsername:input_type -> oauth2_server_service.v1.UpdateUsernameRequest
	4, // 4: oauth2_server_service.v1.AuthService.Register:output_type -> oauth2_server_service.v1.RegisterResponse
	5, // 5: oauth2_server_service.v1.AuthService.ForgotPassword:output_type -> oauth2_server_service.v1.ForgotPasswordResponse
	6, // 6: oauth2_server_service.v1.AuthService.ChangePassword:output_type -> oauth2_server_service.v1.ChangePasswordResponse
	7, // 7: oauth2_server_service.v1.AuthService.UpdateUsername:output_type -> oauth2_server_service.v1.UpdateUsernameResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oauth2_server_service_v1_auth_proto_init() }
func file_oauth2_server_service_v1_auth_proto_init() {
	if File_oauth2_server_service_v1_auth_proto != nil {
		return
	}
	file_oauth2_server_service_v1_register_proto_init()
	file_oauth2_server_service_v1_forgot_password_proto_init()
	file_oauth2_server_service_v1_change_password_proto_init()
	file_oauth2_server_service_v1_update_username_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oauth2_server_service_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oauth2_server_service_v1_auth_proto_goTypes,
		DependencyIndexes: file_oauth2_server_service_v1_auth_proto_depIdxs,
	}.Build()
	File_oauth2_server_service_v1_auth_proto = out.File
	file_oauth2_server_service_v1_auth_proto_rawDesc = nil
	file_oauth2_server_service_v1_auth_proto_goTypes = nil
	file_oauth2_server_service_v1_auth_proto_depIdxs = nil
}