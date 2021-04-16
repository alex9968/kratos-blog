// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: errors/shop.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/api/kratos/api"
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

type User int32

const (
	User_UNKNOWN_ERROR          User = 0
	User_USER_NOT_FOUND         User = 1
	User_PASSWORD_VERIFY_FAILED User = 2
)

// Enum value maps for User.
var (
	User_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "USER_NOT_FOUND",
		2: "PASSWORD_VERIFY_FAILED",
	}
	User_value = map[string]int32{
		"UNKNOWN_ERROR":          0,
		"USER_NOT_FOUND":         1,
		"PASSWORD_VERIFY_FAILED": 2,
	}
)

func (x User) Enum() *User {
	p := new(User)
	*p = x
	return p
}

func (x User) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_shop_proto_enumTypes[0].Descriptor()
}

func (User) Type() protoreflect.EnumType {
	return &file_errors_shop_proto_enumTypes[0]
}

func (x User) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User.Descriptor instead.
func (User) EnumDescriptor() ([]byte, []int) {
	return file_errors_shop_proto_rawDescGZIP(), []int{0}
}

var File_errors_shop_proto protoreflect.FileDescriptor

var file_errors_shop_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x53, 0x53, 0x57,
	0x4f, 0x52, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x1a, 0x03, 0xa0, 0x45, 0x01, 0x42, 0x20, 0x50, 0x01, 0x5a, 0x1c, 0x73, 0x68,
	0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_errors_shop_proto_rawDescOnce sync.Once
	file_errors_shop_proto_rawDescData = file_errors_shop_proto_rawDesc
)

func file_errors_shop_proto_rawDescGZIP() []byte {
	file_errors_shop_proto_rawDescOnce.Do(func() {
		file_errors_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_shop_proto_rawDescData)
	})
	return file_errors_shop_proto_rawDescData
}

var file_errors_shop_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_shop_proto_goTypes = []interface{}{
	(User)(0), // 0: shop.interface.errors.User
}
var file_errors_shop_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_shop_proto_init() }
func file_errors_shop_proto_init() {
	if File_errors_shop_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_shop_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_shop_proto_goTypes,
		DependencyIndexes: file_errors_shop_proto_depIdxs,
		EnumInfos:         file_errors_shop_proto_enumTypes,
	}.Build()
	File_errors_shop_proto = out.File
	file_errors_shop_proto_rawDesc = nil
	file_errors_shop_proto_goTypes = nil
	file_errors_shop_proto_depIdxs = nil
}
