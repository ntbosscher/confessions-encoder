// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: keymapping.proto

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

type KeyMapping_ComputerKeyboard_ModifierFlags int32

const (
	KeyMapping_ComputerKeyboard_MODIFIERFLAGS_COMMAND_KEY  KeyMapping_ComputerKeyboard_ModifierFlags = 0
	KeyMapping_ComputerKeyboard_MODIFIERFLAGS_SHIFT_KEY    KeyMapping_ComputerKeyboard_ModifierFlags = 1
	KeyMapping_ComputerKeyboard_MODIFIERFLAGS_OPTION_KEY   KeyMapping_ComputerKeyboard_ModifierFlags = 2
	KeyMapping_ComputerKeyboard_MODIFIERFLAGS_CONTROL_KEY  KeyMapping_ComputerKeyboard_ModifierFlags = 3
	KeyMapping_ComputerKeyboard_MODIFIERFLAGS_FUNCTION_KEY KeyMapping_ComputerKeyboard_ModifierFlags = 4
)

// Enum value maps for KeyMapping_ComputerKeyboard_ModifierFlags.
var (
	KeyMapping_ComputerKeyboard_ModifierFlags_name = map[int32]string{
		0: "MODIFIERFLAGS_COMMAND_KEY",
		1: "MODIFIERFLAGS_SHIFT_KEY",
		2: "MODIFIERFLAGS_OPTION_KEY",
		3: "MODIFIERFLAGS_CONTROL_KEY",
		4: "MODIFIERFLAGS_FUNCTION_KEY",
	}
	KeyMapping_ComputerKeyboard_ModifierFlags_value = map[string]int32{
		"MODIFIERFLAGS_COMMAND_KEY":  0,
		"MODIFIERFLAGS_SHIFT_KEY":    1,
		"MODIFIERFLAGS_OPTION_KEY":   2,
		"MODIFIERFLAGS_CONTROL_KEY":  3,
		"MODIFIERFLAGS_FUNCTION_KEY": 4,
	}
)

func (x KeyMapping_ComputerKeyboard_ModifierFlags) Enum() *KeyMapping_ComputerKeyboard_ModifierFlags {
	p := new(KeyMapping_ComputerKeyboard_ModifierFlags)
	*p = x
	return p
}

func (x KeyMapping_ComputerKeyboard_ModifierFlags) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyMapping_ComputerKeyboard_ModifierFlags) Descriptor() protoreflect.EnumDescriptor {
	return file_keymapping_proto_enumTypes[0].Descriptor()
}

func (KeyMapping_ComputerKeyboard_ModifierFlags) Type() protoreflect.EnumType {
	return &file_keymapping_proto_enumTypes[0]
}

func (x KeyMapping_ComputerKeyboard_ModifierFlags) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyMapping_ComputerKeyboard_ModifierFlags.Descriptor instead.
func (KeyMapping_ComputerKeyboard_ModifierFlags) EnumDescriptor() ([]byte, []int) {
	return file_keymapping_proto_rawDescGZIP(), []int{0, 0, 0}
}

type KeyMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyboard *KeyMapping_ComputerKeyboard `protobuf:"bytes,1,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	Midi     *KeyMapping_MIDIKeyboard     `protobuf:"bytes,2,opt,name=midi,proto3" json:"midi,omitempty"`
	// Types that are assignable to TargetIdentifier:
	//	*KeyMapping_MenuItem
	//	*KeyMapping_ClearGroupIdentifier
	//	*KeyMapping_CueIdentifier
	//	*KeyMapping_GroupIdentifier
	//	*KeyMapping_MacroIdentifier
	//	*KeyMapping_PropIdentifier
	//	*KeyMapping_TimerIdentifier
	TargetIdentifier isKeyMapping_TargetIdentifier `protobuf_oneof:"TargetIdentifier"`
}

func (x *KeyMapping) Reset() {
	*x = KeyMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyMapping) ProtoMessage() {}

func (x *KeyMapping) ProtoReflect() protoreflect.Message {
	mi := &file_keymapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyMapping.ProtoReflect.Descriptor instead.
func (*KeyMapping) Descriptor() ([]byte, []int) {
	return file_keymapping_proto_rawDescGZIP(), []int{0}
}

func (x *KeyMapping) GetKeyboard() *KeyMapping_ComputerKeyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

func (x *KeyMapping) GetMidi() *KeyMapping_MIDIKeyboard {
	if x != nil {
		return x.Midi
	}
	return nil
}

func (m *KeyMapping) GetTargetIdentifier() isKeyMapping_TargetIdentifier {
	if m != nil {
		return m.TargetIdentifier
	}
	return nil
}

func (x *KeyMapping) GetMenuItem() string {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_MenuItem); ok {
		return x.MenuItem
	}
	return ""
}

func (x *KeyMapping) GetClearGroupIdentifier() *CollectionElementType {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_ClearGroupIdentifier); ok {
		return x.ClearGroupIdentifier
	}
	return nil
}

func (x *KeyMapping) GetCueIdentifier() *CollectionElementType {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_CueIdentifier); ok {
		return x.CueIdentifier
	}
	return nil
}

func (x *KeyMapping) GetGroupIdentifier() *CollectionElementType {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_GroupIdentifier); ok {
		return x.GroupIdentifier
	}
	return nil
}

func (x *KeyMapping) GetMacroIdentifier() *CollectionElementType {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_MacroIdentifier); ok {
		return x.MacroIdentifier
	}
	return nil
}

func (x *KeyMapping) GetPropIdentifier() *CollectionElementType {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_PropIdentifier); ok {
		return x.PropIdentifier
	}
	return nil
}

func (x *KeyMapping) GetTimerIdentifier() *CollectionElementType {
	if x, ok := x.GetTargetIdentifier().(*KeyMapping_TimerIdentifier); ok {
		return x.TimerIdentifier
	}
	return nil
}

type isKeyMapping_TargetIdentifier interface {
	isKeyMapping_TargetIdentifier()
}

type KeyMapping_MenuItem struct {
	MenuItem string `protobuf:"bytes,100,opt,name=menu_item,json=menuItem,proto3,oneof"`
}

type KeyMapping_ClearGroupIdentifier struct {
	ClearGroupIdentifier *CollectionElementType `protobuf:"bytes,101,opt,name=clear_group_identifier,json=clearGroupIdentifier,proto3,oneof"`
}

type KeyMapping_CueIdentifier struct {
	CueIdentifier *CollectionElementType `protobuf:"bytes,102,opt,name=cue_identifier,json=cueIdentifier,proto3,oneof"`
}

type KeyMapping_GroupIdentifier struct {
	GroupIdentifier *CollectionElementType `protobuf:"bytes,103,opt,name=group_identifier,json=groupIdentifier,proto3,oneof"`
}

type KeyMapping_MacroIdentifier struct {
	MacroIdentifier *CollectionElementType `protobuf:"bytes,104,opt,name=macro_identifier,json=macroIdentifier,proto3,oneof"`
}

type KeyMapping_PropIdentifier struct {
	PropIdentifier *CollectionElementType `protobuf:"bytes,105,opt,name=prop_identifier,json=propIdentifier,proto3,oneof"`
}

type KeyMapping_TimerIdentifier struct {
	TimerIdentifier *CollectionElementType `protobuf:"bytes,106,opt,name=timer_identifier,json=timerIdentifier,proto3,oneof"`
}

func (*KeyMapping_MenuItem) isKeyMapping_TargetIdentifier() {}

func (*KeyMapping_ClearGroupIdentifier) isKeyMapping_TargetIdentifier() {}

func (*KeyMapping_CueIdentifier) isKeyMapping_TargetIdentifier() {}

func (*KeyMapping_GroupIdentifier) isKeyMapping_TargetIdentifier() {}

func (*KeyMapping_MacroIdentifier) isKeyMapping_TargetIdentifier() {}

func (*KeyMapping_PropIdentifier) isKeyMapping_TargetIdentifier() {}

func (*KeyMapping_TimerIdentifier) isKeyMapping_TargetIdentifier() {}

type KeyMappingDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationInfo    *ApplicationInfo `protobuf:"bytes,1,opt,name=application_info,json=applicationInfo,proto3" json:"application_info,omitempty"`
	Keymappings        []*KeyMapping    `protobuf:"bytes,2,rep,name=keymappings,proto3" json:"keymappings,omitempty"`
	MacosKeymappings   []*KeyMapping    `protobuf:"bytes,3,rep,name=macos_keymappings,json=macosKeymappings,proto3" json:"macos_keymappings,omitempty"`
	WindowsKeymappings []*KeyMapping    `protobuf:"bytes,4,rep,name=windows_keymappings,json=windowsKeymappings,proto3" json:"windows_keymappings,omitempty"`
}

func (x *KeyMappingDocument) Reset() {
	*x = KeyMappingDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymapping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyMappingDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyMappingDocument) ProtoMessage() {}

func (x *KeyMappingDocument) ProtoReflect() protoreflect.Message {
	mi := &file_keymapping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyMappingDocument.ProtoReflect.Descriptor instead.
func (*KeyMappingDocument) Descriptor() ([]byte, []int) {
	return file_keymapping_proto_rawDescGZIP(), []int{1}
}

func (x *KeyMappingDocument) GetApplicationInfo() *ApplicationInfo {
	if x != nil {
		return x.ApplicationInfo
	}
	return nil
}

func (x *KeyMappingDocument) GetKeymappings() []*KeyMapping {
	if x != nil {
		return x.Keymappings
	}
	return nil
}

func (x *KeyMappingDocument) GetMacosKeymappings() []*KeyMapping {
	if x != nil {
		return x.MacosKeymappings
	}
	return nil
}

func (x *KeyMappingDocument) GetWindowsKeymappings() []*KeyMapping {
	if x != nil {
		return x.WindowsKeymappings
	}
	return nil
}

type KeyMapping_ComputerKeyboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyEquivalent              string                                      `protobuf:"bytes,1,opt,name=key_equivalent,json=keyEquivalent,proto3" json:"key_equivalent,omitempty"`
	KeyEquivalentModifierFlags []KeyMapping_ComputerKeyboard_ModifierFlags `protobuf:"varint,2,rep,packed,name=key_equivalent_modifier_flags,json=keyEquivalentModifierFlags,proto3,enum=rv.data.KeyMapping_ComputerKeyboard_ModifierFlags" json:"key_equivalent_modifier_flags,omitempty"`
}

func (x *KeyMapping_ComputerKeyboard) Reset() {
	*x = KeyMapping_ComputerKeyboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymapping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyMapping_ComputerKeyboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyMapping_ComputerKeyboard) ProtoMessage() {}

func (x *KeyMapping_ComputerKeyboard) ProtoReflect() protoreflect.Message {
	mi := &file_keymapping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyMapping_ComputerKeyboard.ProtoReflect.Descriptor instead.
func (*KeyMapping_ComputerKeyboard) Descriptor() ([]byte, []int) {
	return file_keymapping_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KeyMapping_ComputerKeyboard) GetKeyEquivalent() string {
	if x != nil {
		return x.KeyEquivalent
	}
	return ""
}

func (x *KeyMapping_ComputerKeyboard) GetKeyEquivalentModifierFlags() []KeyMapping_ComputerKeyboard_ModifierFlags {
	if x != nil {
		return x.KeyEquivalentModifierFlags
	}
	return nil
}

type KeyMapping_MIDIKeyboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel  int32 `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Pitch    int32 `protobuf:"varint,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Velocity int32 `protobuf:"varint,3,opt,name=velocity,proto3" json:"velocity,omitempty"`
}

func (x *KeyMapping_MIDIKeyboard) Reset() {
	*x = KeyMapping_MIDIKeyboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymapping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyMapping_MIDIKeyboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyMapping_MIDIKeyboard) ProtoMessage() {}

func (x *KeyMapping_MIDIKeyboard) ProtoReflect() protoreflect.Message {
	mi := &file_keymapping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyMapping_MIDIKeyboard.ProtoReflect.Descriptor instead.
func (*KeyMapping_MIDIKeyboard) Descriptor() ([]byte, []int) {
	return file_keymapping_proto_rawDescGZIP(), []int{0, 1}
}

func (x *KeyMapping_MIDIKeyboard) GetChannel() int32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *KeyMapping_MIDIKeyboard) GetPitch() int32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *KeyMapping_MIDIKeyboard) GetVelocity() int32 {
	if x != nil {
		return x.Velocity
	}
	return 0
}

var File_keymapping_proto protoreflect.FileDescriptor

var file_keymapping_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x08,
	0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x34,
	0x0a, 0x04, 0x6d, 0x69, 0x64, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72,
	0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x49, 0x44, 0x49, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x04,
	0x6d, 0x69, 0x64, 0x69, 0x12, 0x1d, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x56, 0x0a, 0x16, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0e, 0x63,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00,
	0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x10, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x76,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x6d,
	0x61, 0x63, 0x72, 0x6f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x49,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x10, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x6a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0xdb, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6b,
	0x65, 0x79, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x45, 0x71, 0x75, 0x69, 0x76, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x12, 0x75, 0x0a, 0x1d, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x76, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x72, 0x76, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x1a, 0x6b,
	0x65, 0x79, 0x45, 0x71, 0x75, 0x69, 0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x4d,
	0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x46, 0x4c, 0x41, 0x47, 0x53, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f,
	0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x46, 0x4c, 0x41, 0x47, 0x53, 0x5f, 0x53, 0x48, 0x49, 0x46,
	0x54, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x49, 0x45, 0x52, 0x46, 0x4c, 0x41, 0x47, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4b, 0x45, 0x59, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x46, 0x4c, 0x41, 0x47, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4b,
	0x45, 0x59, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52,
	0x46, 0x4c, 0x41, 0x47, 0x53, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4b,
	0x45, 0x59, 0x10, 0x04, 0x1a, 0x5a, 0x0a, 0x0c, 0x4d, 0x49, 0x44, 0x49, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x69, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x42, 0x12, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0x98, 0x02, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x35, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x11, 0x6d, 0x61, 0x63, 0x6f, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x6d, 0x61, 0x63, 0x6f, 0x73, 0x4b, 0x65,
	0x79, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x13, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x12, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x73, 0x4b, 0x65, 0x79, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x42,
	0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keymapping_proto_rawDescOnce sync.Once
	file_keymapping_proto_rawDescData = file_keymapping_proto_rawDesc
)

func file_keymapping_proto_rawDescGZIP() []byte {
	file_keymapping_proto_rawDescOnce.Do(func() {
		file_keymapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_keymapping_proto_rawDescData)
	})
	return file_keymapping_proto_rawDescData
}

var file_keymapping_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_keymapping_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_keymapping_proto_goTypes = []interface{}{
	(KeyMapping_ComputerKeyboard_ModifierFlags)(0), // 0: rv.data.KeyMapping.ComputerKeyboard.ModifierFlags
	(*KeyMapping)(nil),                  // 1: rv.data.KeyMapping
	(*KeyMappingDocument)(nil),          // 2: rv.data.KeyMappingDocument
	(*KeyMapping_ComputerKeyboard)(nil), // 3: rv.data.KeyMapping.ComputerKeyboard
	(*KeyMapping_MIDIKeyboard)(nil),     // 4: rv.data.KeyMapping.MIDIKeyboard
	(*CollectionElementType)(nil),       // 5: rv.data.CollectionElementType
	(*ApplicationInfo)(nil),             // 6: rv.data.ApplicationInfo
}
var file_keymapping_proto_depIdxs = []int32{
	3,  // 0: rv.data.KeyMapping.keyboard:type_name -> rv.data.KeyMapping.ComputerKeyboard
	4,  // 1: rv.data.KeyMapping.midi:type_name -> rv.data.KeyMapping.MIDIKeyboard
	5,  // 2: rv.data.KeyMapping.clear_group_identifier:type_name -> rv.data.CollectionElementType
	5,  // 3: rv.data.KeyMapping.cue_identifier:type_name -> rv.data.CollectionElementType
	5,  // 4: rv.data.KeyMapping.group_identifier:type_name -> rv.data.CollectionElementType
	5,  // 5: rv.data.KeyMapping.macro_identifier:type_name -> rv.data.CollectionElementType
	5,  // 6: rv.data.KeyMapping.prop_identifier:type_name -> rv.data.CollectionElementType
	5,  // 7: rv.data.KeyMapping.timer_identifier:type_name -> rv.data.CollectionElementType
	6,  // 8: rv.data.KeyMappingDocument.application_info:type_name -> rv.data.ApplicationInfo
	1,  // 9: rv.data.KeyMappingDocument.keymappings:type_name -> rv.data.KeyMapping
	1,  // 10: rv.data.KeyMappingDocument.macos_keymappings:type_name -> rv.data.KeyMapping
	1,  // 11: rv.data.KeyMappingDocument.windows_keymappings:type_name -> rv.data.KeyMapping
	0,  // 12: rv.data.KeyMapping.ComputerKeyboard.key_equivalent_modifier_flags:type_name -> rv.data.KeyMapping.ComputerKeyboard.ModifierFlags
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_keymapping_proto_init() }
func file_keymapping_proto_init() {
	if File_keymapping_proto != nil {
		return
	}
	file_basicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_keymapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyMapping); i {
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
		file_keymapping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyMappingDocument); i {
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
		file_keymapping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyMapping_ComputerKeyboard); i {
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
		file_keymapping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyMapping_MIDIKeyboard); i {
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
	file_keymapping_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*KeyMapping_MenuItem)(nil),
		(*KeyMapping_ClearGroupIdentifier)(nil),
		(*KeyMapping_CueIdentifier)(nil),
		(*KeyMapping_GroupIdentifier)(nil),
		(*KeyMapping_MacroIdentifier)(nil),
		(*KeyMapping_PropIdentifier)(nil),
		(*KeyMapping_TimerIdentifier)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keymapping_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keymapping_proto_goTypes,
		DependencyIndexes: file_keymapping_proto_depIdxs,
		EnumInfos:         file_keymapping_proto_enumTypes,
		MessageInfos:      file_keymapping_proto_msgTypes,
	}.Build()
	File_keymapping_proto = out.File
	file_keymapping_proto_rawDesc = nil
	file_keymapping_proto_goTypes = nil
	file_keymapping_proto_depIdxs = nil
}
