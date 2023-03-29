// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobuf/claims.proto

package protobuf

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

type Claims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ClientId       int64  `protobuf:"varint,2,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	UnitId         int64  `protobuf:"varint,3,opt,name=UnitId,proto3" json:"UnitId,omitempty"`
	Username       string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Email          string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Fullname       string `protobuf:"bytes,6,opt,name=Fullname,proto3" json:"Fullname,omitempty"`
	Phone          string `protobuf:"bytes,7,opt,name=Phone,proto3" json:"Phone,omitempty"`
	IsAdmin        bool   `protobuf:"varint,8,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
	IsSystem       bool   `protobuf:"varint,9,opt,name=IsSystem,proto3" json:"IsSystem,omitempty"`
	Language       string `protobuf:"bytes,10,opt,name=Language,proto3" json:"Language,omitempty"`
	IsBaseLanguage bool   `protobuf:"varint,11,opt,name=IsBaseLanguage,proto3" json:"IsBaseLanguage,omitempty"`
}

func (x *Claims) Reset() {
	*x = Claims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_claims_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claims) ProtoMessage() {}

func (x *Claims) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_claims_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claims.ProtoReflect.Descriptor instead.
func (*Claims) Descriptor() ([]byte, []int) {
	return file_protobuf_claims_proto_rawDescGZIP(), []int{0}
}

func (x *Claims) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Claims) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Claims) GetUnitId() int64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *Claims) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Claims) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Claims) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Claims) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Claims) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *Claims) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

func (x *Claims) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Claims) GetIsBaseLanguage() bool {
	if x != nil {
		return x.IsBaseLanguage
	}
	return false
}

var File_protobuf_claims_proto protoreflect.FileDescriptor

var file_protobuf_claims_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x69, 0x64, 0x65, 0x61, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xaa, 0x02, 0x0a,
	0x06, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x64, 0x65, 0x61, 0x2d, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2d, 0x6c, 0x6c, 0x63, 0x2f, 0x6d, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_claims_proto_rawDescOnce sync.Once
	file_protobuf_claims_proto_rawDescData = file_protobuf_claims_proto_rawDesc
)

func file_protobuf_claims_proto_rawDescGZIP() []byte {
	file_protobuf_claims_proto_rawDescOnce.Do(func() {
		file_protobuf_claims_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_claims_proto_rawDescData)
	})
	return file_protobuf_claims_proto_rawDescData
}

var file_protobuf_claims_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_claims_proto_goTypes = []interface{}{
	(*Claims)(nil), // 0: mideamedia.protobuf.Claims
}
var file_protobuf_claims_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_claims_proto_init() }
func file_protobuf_claims_proto_init() {
	if File_protobuf_claims_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_claims_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claims); i {
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
			RawDescriptor: file_protobuf_claims_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_claims_proto_goTypes,
		DependencyIndexes: file_protobuf_claims_proto_depIdxs,
		MessageInfos:      file_protobuf_claims_proto_msgTypes,
	}.Build()
	File_protobuf_claims_proto = out.File
	file_protobuf_claims_proto_rawDesc = nil
	file_protobuf_claims_proto_goTypes = nil
	file_protobuf_claims_proto_depIdxs = nil
}
