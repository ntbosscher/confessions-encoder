// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proMask.proto

package protos

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

type ProMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseSlide *Slide `protobuf:"bytes,1,opt,name=base_slide,json=baseSlide,proto3" json:"base_slide,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProMask) Reset() {
	*x = ProMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proMask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProMask) ProtoMessage() {}

func (x *ProMask) ProtoReflect() protoreflect.Message {
	mi := &file_proMask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProMask.ProtoReflect.Descriptor instead.
func (*ProMask) Descriptor() ([]byte, []int) {
	return file_proMask_proto_rawDescGZIP(), []int{0}
}

func (x *ProMask) GetBaseSlide() *Slide {
	if x != nil {
		return x.BaseSlide
	}
	return nil
}

func (x *ProMask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proMask_proto protoreflect.FileDescriptor

var file_proMask_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x4d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x2d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53,
	0x6c, 0x69, 0x64, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proMask_proto_rawDescOnce sync.Once
	file_proMask_proto_rawDescData = file_proMask_proto_rawDesc
)

func file_proMask_proto_rawDescGZIP() []byte {
	file_proMask_proto_rawDescOnce.Do(func() {
		file_proMask_proto_rawDescData = protoimpl.X.CompressGZIP(file_proMask_proto_rawDescData)
	})
	return file_proMask_proto_rawDescData
}

var file_proMask_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proMask_proto_goTypes = []interface{}{
	(*ProMask)(nil), // 0: rv.data.ProMask
	(*Slide)(nil),   // 1: rv.data.Slide
}
var file_proMask_proto_depIdxs = []int32{
	1, // 0: rv.data.ProMask.base_slide:type_name -> rv.data.Slide
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proMask_proto_init() }
func file_proMask_proto_init() {
	if File_proMask_proto != nil {
		return
	}
	file_slide_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proMask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProMask); i {
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
			RawDescriptor: file_proMask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proMask_proto_goTypes,
		DependencyIndexes: file_proMask_proto_depIdxs,
		MessageInfos:      file_proMask_proto_msgTypes,
	}.Build()
	File_proMask_proto = out.File
	file_proMask_proto_rawDesc = nil
	file_proMask_proto_goTypes = nil
	file_proMask_proto_depIdxs = nil
}