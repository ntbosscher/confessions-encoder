// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: template.proto

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

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0}
}

type Template_Slide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseSlide *Slide `protobuf:"bytes,1,opt,name=base_slide,json=baseSlide,proto3" json:"base_slide,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Template_Slide) Reset() {
	*x = Template_Slide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template_Slide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template_Slide) ProtoMessage() {}

func (x *Template_Slide) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template_Slide.ProtoReflect.Descriptor instead.
func (*Template_Slide) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Template_Slide) GetBaseSlide() *Slide {
	if x != nil {
		return x.BaseSlide
	}
	return nil
}

func (x *Template_Slide) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Template_Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationInfo *ApplicationInfo  `protobuf:"bytes,1,opt,name=application_info,json=applicationInfo,proto3" json:"application_info,omitempty"`
	Slides          []*Template_Slide `protobuf:"bytes,3,rep,name=slides,proto3" json:"slides,omitempty"`
}

func (x *Template_Document) Reset() {
	*x = Template_Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template_Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template_Document) ProtoMessage() {}

func (x *Template_Document) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template_Document.ProtoReflect.Descriptor instead.
func (*Template_Document) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Template_Document) GetApplicationInfo() *ApplicationInfo {
	if x != nil {
		return x.ApplicationInfo
	}
	return nil
}

func (x *Template_Document) GetSlides() []*Template_Slide {
	if x != nil {
		return x.Slides
	}
	return nil
}

type Template_Identification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       *UUID  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SlideUuid  *UUID  `protobuf:"bytes,3,opt,name=slide_uuid,json=slideUuid,proto3" json:"slide_uuid,omitempty"`
	SlideName  string `protobuf:"bytes,4,opt,name=slide_name,json=slideName,proto3" json:"slide_name,omitempty"`
	SlideIndex uint32 `protobuf:"varint,5,opt,name=slide_index,json=slideIndex,proto3" json:"slide_index,omitempty"`
}

func (x *Template_Identification) Reset() {
	*x = Template_Identification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template_Identification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template_Identification) ProtoMessage() {}

func (x *Template_Identification) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template_Identification.ProtoReflect.Descriptor instead.
func (*Template_Identification) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Template_Identification) GetUuid() *UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *Template_Identification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template_Identification) GetSlideUuid() *UUID {
	if x != nil {
		return x.SlideUuid
	}
	return nil
}

func (x *Template_Identification) GetSlideName() string {
	if x != nil {
		return x.SlideName
	}
	return ""
}

func (x *Template_Identification) GetSlideIndex() uint32 {
	if x != nil {
		return x.SlideIndex
	}
	return 0
}

var File_template_proto protoreflect.FileDescriptor

var file_template_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x6c, 0x69,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x08, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x4a, 0x0a, 0x05, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x12, 0x2d,
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x6c, 0x69,
	0x64, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x1a, 0x80, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x43,
	0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x06, 0x73, 0x6c,
	0x69, 0x64, 0x65, 0x73, 0x1a, 0xb5, 0x01, 0x0a, 0x0e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x0a, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x09, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x6c, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x1c, 0x5a, 0x1a,
	0x63, 0x6f, 0x6e, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_template_proto_rawDescOnce sync.Once
	file_template_proto_rawDescData = file_template_proto_rawDesc
)

func file_template_proto_rawDescGZIP() []byte {
	file_template_proto_rawDescOnce.Do(func() {
		file_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_proto_rawDescData)
	})
	return file_template_proto_rawDescData
}

var file_template_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_template_proto_goTypes = []interface{}{
	(*Template)(nil),                // 0: rv.data.Template
	(*Template_Slide)(nil),          // 1: rv.data.Template.Slide
	(*Template_Document)(nil),       // 2: rv.data.Template.Document
	(*Template_Identification)(nil), // 3: rv.data.Template.Identification
	(*Slide)(nil),                   // 4: rv.data.Slide
	(*ApplicationInfo)(nil),         // 5: rv.data.ApplicationInfo
	(*UUID)(nil),                    // 6: rv.data.UUID
}
var file_template_proto_depIdxs = []int32{
	4, // 0: rv.data.Template.Slide.base_slide:type_name -> rv.data.Slide
	5, // 1: rv.data.Template.Document.application_info:type_name -> rv.data.ApplicationInfo
	1, // 2: rv.data.Template.Document.slides:type_name -> rv.data.Template.Slide
	6, // 3: rv.data.Template.Identification.uuid:type_name -> rv.data.UUID
	6, // 4: rv.data.Template.Identification.slide_uuid:type_name -> rv.data.UUID
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_template_proto_init() }
func file_template_proto_init() {
	if File_template_proto != nil {
		return
	}
	file_basicTypes_proto_init()
	file_slide_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template_Slide); i {
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
		file_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template_Document); i {
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
		file_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template_Identification); i {
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
			RawDescriptor: file_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_proto_goTypes,
		DependencyIndexes: file_template_proto_depIdxs,
		MessageInfos:      file_template_proto_msgTypes,
	}.Build()
	File_template_proto = out.File
	file_template_proto_rawDesc = nil
	file_template_proto_goTypes = nil
	file_template_proto_depIdxs = nil
}
