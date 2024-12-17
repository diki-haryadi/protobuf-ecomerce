// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: go_micro_template/article/v1/article.proto

package articlev1

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

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_micro_template_article_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_micro_template_article_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_go_micro_template_article_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArticleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateArticleRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_micro_template_article_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_micro_template_article_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_go_micro_template_article_v1_article_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateArticleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateArticleResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type GetArticleByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleByIdRequest) Reset() {
	*x = GetArticleByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_micro_template_article_v1_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleByIdRequest) ProtoMessage() {}

func (x *GetArticleByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_micro_template_article_v1_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleByIdRequest.ProtoReflect.Descriptor instead.
func (*GetArticleByIdRequest) Descriptor() ([]byte, []int) {
	return file_go_micro_template_article_v1_article_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticleByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetArticleByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *GetArticleByIdResponse) Reset() {
	*x = GetArticleByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_micro_template_article_v1_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleByIdResponse) ProtoMessage() {}

func (x *GetArticleByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_micro_template_article_v1_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleByIdResponse.ProtoReflect.Descriptor instead.
func (*GetArticleByIdResponse) Descriptor() ([]byte, []int) {
	return file_go_micro_template_article_v1_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleByIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetArticleByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetArticleByIdResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

var File_go_micro_template_article_v1_article_proto protoreflect.FileDescriptor

var file_go_micro_template_article_v1_article_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f,
	0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x4f, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x27, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x32, 0x87, 0x02, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x5f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x5f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x8c, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x6b, 0x69, 0x2d, 0x68, 0x61, 0x72, 0x79, 0x61, 0x64, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x58, 0xaa, 0x02, 0x1a,
	0x47, 0x6f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x47, 0x6f, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1c, 0x47, 0x6f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_micro_template_article_v1_article_proto_rawDescOnce sync.Once
	file_go_micro_template_article_v1_article_proto_rawDescData = file_go_micro_template_article_v1_article_proto_rawDesc
)

func file_go_micro_template_article_v1_article_proto_rawDescGZIP() []byte {
	file_go_micro_template_article_v1_article_proto_rawDescOnce.Do(func() {
		file_go_micro_template_article_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_micro_template_article_v1_article_proto_rawDescData)
	})
	return file_go_micro_template_article_v1_article_proto_rawDescData
}

var file_go_micro_template_article_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_micro_template_article_v1_article_proto_goTypes = []interface{}{
	(*CreateArticleRequest)(nil),   // 0: go_micro_template.article.v1.CreateArticleRequest
	(*CreateArticleResponse)(nil),  // 1: go_micro_template.article.v1.CreateArticleResponse
	(*GetArticleByIdRequest)(nil),  // 2: go_micro_template.article.v1.GetArticleByIdRequest
	(*GetArticleByIdResponse)(nil), // 3: go_micro_template.article.v1.GetArticleByIdResponse
}
var file_go_micro_template_article_v1_article_proto_depIdxs = []int32{
	0, // 0: go_micro_template.article.v1.ArticleService.CreateArticle:input_type -> go_micro_template.article.v1.CreateArticleRequest
	2, // 1: go_micro_template.article.v1.ArticleService.GetArticleById:input_type -> go_micro_template.article.v1.GetArticleByIdRequest
	1, // 2: go_micro_template.article.v1.ArticleService.CreateArticle:output_type -> go_micro_template.article.v1.CreateArticleResponse
	3, // 3: go_micro_template.article.v1.ArticleService.GetArticleById:output_type -> go_micro_template.article.v1.GetArticleByIdResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_micro_template_article_v1_article_proto_init() }
func file_go_micro_template_article_v1_article_proto_init() {
	if File_go_micro_template_article_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_micro_template_article_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_go_micro_template_article_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_go_micro_template_article_v1_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleByIdRequest); i {
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
		file_go_micro_template_article_v1_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleByIdResponse); i {
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
			RawDescriptor: file_go_micro_template_article_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_micro_template_article_v1_article_proto_goTypes,
		DependencyIndexes: file_go_micro_template_article_v1_article_proto_depIdxs,
		MessageInfos:      file_go_micro_template_article_v1_article_proto_msgTypes,
	}.Build()
	File_go_micro_template_article_v1_article_proto = out.File
	file_go_micro_template_article_v1_article_proto_rawDesc = nil
	file_go_micro_template_article_v1_article_proto_goTypes = nil
	file_go_micro_template_article_v1_article_proto_depIdxs = nil
}