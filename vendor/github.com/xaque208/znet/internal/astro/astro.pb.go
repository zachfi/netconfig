// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: internal/astro/astro.proto

package astro

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_astro_astro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_astro_astro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_astro_astro_proto_rawDescGZIP(), []int{0}
}

var File_internal_astro_astro_proto protoreflect.FileDescriptor

var file_internal_astro_astro_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x73, 0x74, 0x72, 0x6f,
	0x2f, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x73,
	0x74, 0x72, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x7d, 0x0a, 0x05,
	0x41, 0x73, 0x74, 0x72, 0x6f, 0x12, 0x25, 0x0a, 0x07, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65,
	0x12, 0x0c, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c,
	0x2e, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x06,
	0x53, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x27, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x53, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x12,
	0x0c, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e,
	0x61, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x73, 0x74, 0x72, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_astro_astro_proto_rawDescOnce sync.Once
	file_internal_astro_astro_proto_rawDescData = file_internal_astro_astro_proto_rawDesc
)

func file_internal_astro_astro_proto_rawDescGZIP() []byte {
	file_internal_astro_astro_proto_rawDescOnce.Do(func() {
		file_internal_astro_astro_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_astro_astro_proto_rawDescData)
	})
	return file_internal_astro_astro_proto_rawDescData
}

var file_internal_astro_astro_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_astro_astro_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: astro.Empty
}
var file_internal_astro_astro_proto_depIdxs = []int32{
	0, // 0: astro.Astro.Sunrise:input_type -> astro.Empty
	0, // 1: astro.Astro.Sunset:input_type -> astro.Empty
	0, // 2: astro.Astro.PreSunset:input_type -> astro.Empty
	0, // 3: astro.Astro.Sunrise:output_type -> astro.Empty
	0, // 4: astro.Astro.Sunset:output_type -> astro.Empty
	0, // 5: astro.Astro.PreSunset:output_type -> astro.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_astro_astro_proto_init() }
func file_internal_astro_astro_proto_init() {
	if File_internal_astro_astro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_astro_astro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_internal_astro_astro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_astro_astro_proto_goTypes,
		DependencyIndexes: file_internal_astro_astro_proto_depIdxs,
		MessageInfos:      file_internal_astro_astro_proto_msgTypes,
	}.Build()
	File_internal_astro_astro_proto = out.File
	file_internal_astro_astro_proto_rawDesc = nil
	file_internal_astro_astro_proto_goTypes = nil
	file_internal_astro_astro_proto_depIdxs = nil
}