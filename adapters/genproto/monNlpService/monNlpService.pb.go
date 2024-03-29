// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: monNlpService.proto

package monNlp

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

type TokenizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TokenizeRequest) Reset() {
	*x = TokenizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monNlpService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizeRequest) ProtoMessage() {}

func (x *TokenizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monNlpService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizeRequest.ProtoReflect.Descriptor instead.
func (*TokenizeRequest) Descriptor() ([]byte, []int) {
	return file_monNlpService_proto_rawDescGZIP(), []int{0}
}

func (x *TokenizeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TokenizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *TokenizeResponse) Reset() {
	*x = TokenizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monNlpService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizeResponse) ProtoMessage() {}

func (x *TokenizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monNlpService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizeResponse.ProtoReflect.Descriptor instead.
func (*TokenizeResponse) Descriptor() ([]byte, []int) {
	return file_monNlpService_proto_rawDescGZIP(), []int{1}
}

func (x *TokenizeResponse) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Orth_   string `protobuf:"bytes,2,opt,name=orth_,json=orth,proto3" json:"orth_,omitempty"`
	Lemma_  string `protobuf:"bytes,3,opt,name=lemma_,json=lemma,proto3" json:"lemma_,omitempty"`
	Norm_   string `protobuf:"bytes,4,opt,name=norm_,json=norm,proto3" json:"norm_,omitempty"`
	Lower_  string `protobuf:"bytes,5,opt,name=lower_,json=lower,proto3" json:"lower_,omitempty"`
	Shape_  string `protobuf:"bytes,6,opt,name=shape_,json=shape,proto3" json:"shape_,omitempty"`
	Prefix_ string `protobuf:"bytes,7,opt,name=prefix_,json=prefix,proto3" json:"prefix_,omitempty"`
	Suffix_ string `protobuf:"bytes,8,opt,name=suffix_,json=suffix,proto3" json:"suffix_,omitempty"`
	Pos_    string `protobuf:"bytes,9,opt,name=pos_,json=pos,proto3" json:"pos_,omitempty"`
	Tag_    string `protobuf:"bytes,10,opt,name=tag_,json=tag,proto3" json:"tag_,omitempty"`
	Dep_    string `protobuf:"bytes,11,opt,name=dep_,json=dep,proto3" json:"dep_,omitempty"`
	Lang_   string `protobuf:"bytes,12,opt,name=lang_,json=lang,proto3" json:"lang_,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monNlpService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_monNlpService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_monNlpService_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Token) GetOrth_() string {
	if x != nil {
		return x.Orth_
	}
	return ""
}

func (x *Token) GetLemma_() string {
	if x != nil {
		return x.Lemma_
	}
	return ""
}

func (x *Token) GetNorm_() string {
	if x != nil {
		return x.Norm_
	}
	return ""
}

func (x *Token) GetLower_() string {
	if x != nil {
		return x.Lower_
	}
	return ""
}

func (x *Token) GetShape_() string {
	if x != nil {
		return x.Shape_
	}
	return ""
}

func (x *Token) GetPrefix_() string {
	if x != nil {
		return x.Prefix_
	}
	return ""
}

func (x *Token) GetSuffix_() string {
	if x != nil {
		return x.Suffix_
	}
	return ""
}

func (x *Token) GetPos_() string {
	if x != nil {
		return x.Pos_
	}
	return ""
}

func (x *Token) GetTag_() string {
	if x != nil {
		return x.Tag_
	}
	return ""
}

func (x *Token) GetDep_() string {
	if x != nil {
		return x.Dep_
	}
	return ""
}

func (x *Token) GetLang_() string {
	if x != nil {
		return x.Lang_
	}
	return ""
}

var File_monNlpService_proto protoreflect.FileDescriptor

var file_monNlpService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x6e, 0x4e, 0x6c, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x6e, 0x4e, 0x6c, 0x70, 0x22, 0x25, 0x0a,
	0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x39, 0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x6e, 0x4e, 0x6c,
	0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22,
	0x8a, 0x02, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x13, 0x0a,
	0x05, 0x6f, 0x72, 0x74, 0x68, 0x5f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x72,
	0x74, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x65, 0x6d, 0x6d, 0x61, 0x5f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x6d, 0x6d, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x6f, 0x72,
	0x6d, 0x5f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x72, 0x6d, 0x12, 0x15,
	0x0a, 0x06, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x5f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x11,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x5f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f,
	0x73, 0x12, 0x11, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x5f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x11, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x5f, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x70, 0x12, 0x13, 0x0a, 0x05, 0x6c, 0x61, 0x6e, 0x67, 0x5f,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x32, 0x4e, 0x0a, 0x0d,
	0x4d, 0x6f, 0x6e, 0x4e, 0x6c, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x4e,
	0x6c, 0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x4e, 0x6c, 0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12,
	0x6d, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x6f, 0x6e, 0x4e,
	0x6c, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monNlpService_proto_rawDescOnce sync.Once
	file_monNlpService_proto_rawDescData = file_monNlpService_proto_rawDesc
)

func file_monNlpService_proto_rawDescGZIP() []byte {
	file_monNlpService_proto_rawDescOnce.Do(func() {
		file_monNlpService_proto_rawDescData = protoimpl.X.CompressGZIP(file_monNlpService_proto_rawDescData)
	})
	return file_monNlpService_proto_rawDescData
}

var file_monNlpService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_monNlpService_proto_goTypes = []interface{}{
	(*TokenizeRequest)(nil),  // 0: monNlp.TokenizeRequest
	(*TokenizeResponse)(nil), // 1: monNlp.TokenizeResponse
	(*Token)(nil),            // 2: monNlp.Token
}
var file_monNlpService_proto_depIdxs = []int32{
	2, // 0: monNlp.TokenizeResponse.tokens:type_name -> monNlp.Token
	0, // 1: monNlp.MonNlpService.tokenize:input_type -> monNlp.TokenizeRequest
	1, // 2: monNlp.MonNlpService.tokenize:output_type -> monNlp.TokenizeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_monNlpService_proto_init() }
func file_monNlpService_proto_init() {
	if File_monNlpService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monNlpService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizeRequest); i {
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
		file_monNlpService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizeResponse); i {
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
		file_monNlpService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_monNlpService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monNlpService_proto_goTypes,
		DependencyIndexes: file_monNlpService_proto_depIdxs,
		MessageInfos:      file_monNlpService_proto_msgTypes,
	}.Build()
	File_monNlpService_proto = out.File
	file_monNlpService_proto_rawDesc = nil
	file_monNlpService_proto_goTypes = nil
	file_monNlpService_proto_depIdxs = nil
}
