// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: template-manager.proto

package template_manager

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TemplateConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateID         string `protobuf:"bytes,1,opt,name=templateID,proto3" json:"templateID,omitempty"`
	BuildID            string `protobuf:"bytes,2,opt,name=buildID,proto3" json:"buildID,omitempty"`
	MemoryMB           int32  `protobuf:"varint,3,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	VCpuCount          int32  `protobuf:"varint,4,opt,name=vCpuCount,proto3" json:"vCpuCount,omitempty"`
	DiskSizeMB         int32  `protobuf:"varint,5,opt,name=diskSizeMB,proto3" json:"diskSizeMB,omitempty"`
	KernelVersion      string `protobuf:"bytes,6,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	FirecrackerVersion string `protobuf:"bytes,7,opt,name=firecrackerVersion,proto3" json:"firecrackerVersion,omitempty"`
	StartCommand       string `protobuf:"bytes,8,opt,name=startCommand,proto3" json:"startCommand,omitempty"`
	HugePages          bool   `protobuf:"varint,9,opt,name=hugePages,proto3" json:"hugePages,omitempty"`
}

func (x *TemplateConfig) Reset() {
	*x = TemplateConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateConfig) ProtoMessage() {}

func (x *TemplateConfig) ProtoReflect() protoreflect.Message {
	mi := &file_template_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateConfig.ProtoReflect.Descriptor instead.
func (*TemplateConfig) Descriptor() ([]byte, []int) {
	return file_template_manager_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateConfig) GetTemplateID() string {
	if x != nil {
		return x.TemplateID
	}
	return ""
}

func (x *TemplateConfig) GetBuildID() string {
	if x != nil {
		return x.BuildID
	}
	return ""
}

func (x *TemplateConfig) GetMemoryMB() int32 {
	if x != nil {
		return x.MemoryMB
	}
	return 0
}

func (x *TemplateConfig) GetVCpuCount() int32 {
	if x != nil {
		return x.VCpuCount
	}
	return 0
}

func (x *TemplateConfig) GetDiskSizeMB() int32 {
	if x != nil {
		return x.DiskSizeMB
	}
	return 0
}

func (x *TemplateConfig) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *TemplateConfig) GetFirecrackerVersion() string {
	if x != nil {
		return x.FirecrackerVersion
	}
	return ""
}

func (x *TemplateConfig) GetStartCommand() string {
	if x != nil {
		return x.StartCommand
	}
	return ""
}

func (x *TemplateConfig) GetHugePages() bool {
	if x != nil {
		return x.HugePages
	}
	return false
}

type TemplateCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *TemplateConfig `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *TemplateCreateRequest) Reset() {
	*x = TemplateCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateCreateRequest) ProtoMessage() {}

func (x *TemplateCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateCreateRequest.ProtoReflect.Descriptor instead.
func (*TemplateCreateRequest) Descriptor() ([]byte, []int) {
	return file_template_manager_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateCreateRequest) GetTemplate() *TemplateConfig {
	if x != nil {
		return x.Template
	}
	return nil
}

// Data required for deleting a template.
type TemplateDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateID string `protobuf:"bytes,1,opt,name=templateID,proto3" json:"templateID,omitempty"`
}

func (x *TemplateDeleteRequest) Reset() {
	*x = TemplateDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateDeleteRequest) ProtoMessage() {}

func (x *TemplateDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateDeleteRequest.ProtoReflect.Descriptor instead.
func (*TemplateDeleteRequest) Descriptor() ([]byte, []int) {
	return file_template_manager_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateDeleteRequest) GetTemplateID() string {
	if x != nil {
		return x.TemplateID
	}
	return ""
}

// Logs from template build
type TemplateBuildLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *TemplateBuildLog) Reset() {
	*x = TemplateBuildLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateBuildLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateBuildLog) ProtoMessage() {}

func (x *TemplateBuildLog) ProtoReflect() protoreflect.Message {
	mi := &file_template_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateBuildLog.ProtoReflect.Descriptor instead.
func (*TemplateBuildLog) Descriptor() ([]byte, []int) {
	return file_template_manager_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateBuildLog) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_template_manager_proto protoreflect.FileDescriptor

var file_template_manager_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x0e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x12, 0x1c,
	0x0a, 0x09, 0x76, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x76, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x42, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x42, 0x12, 0x24, 0x0a, 0x0d,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x65, 0x63, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x66, 0x69, 0x72, 0x65, 0x63, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x75, 0x67, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x75, 0x67, 0x65, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x15, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x15, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x44, 0x22, 0x24, 0x0a, 0x10, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x32, 0x92, 0x01, 0x0a, 0x0f, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x0e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x36,
	0x5a, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x75, 0x6c, 0x6e, 0x61, 0x73, 0x6f, 0x66, 0x74, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_manager_proto_rawDescOnce sync.Once
	file_template_manager_proto_rawDescData = file_template_manager_proto_rawDesc
)

func file_template_manager_proto_rawDescGZIP() []byte {
	file_template_manager_proto_rawDescOnce.Do(func() {
		file_template_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_manager_proto_rawDescData)
	})
	return file_template_manager_proto_rawDescData
}

var file_template_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_template_manager_proto_goTypes = []interface{}{
	(*TemplateConfig)(nil),        // 0: TemplateConfig
	(*TemplateCreateRequest)(nil), // 1: TemplateCreateRequest
	(*TemplateDeleteRequest)(nil), // 2: TemplateDeleteRequest
	(*TemplateBuildLog)(nil),      // 3: TemplateBuildLog
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_template_manager_proto_depIdxs = []int32{
	0, // 0: TemplateCreateRequest.template:type_name -> TemplateConfig
	1, // 1: TemplateService.TemplateCreate:input_type -> TemplateCreateRequest
	2, // 2: TemplateService.TemplateDelete:input_type -> TemplateDeleteRequest
	3, // 3: TemplateService.TemplateCreate:output_type -> TemplateBuildLog
	4, // 4: TemplateService.TemplateDelete:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_template_manager_proto_init() }
func file_template_manager_proto_init() {
	if File_template_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateConfig); i {
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
		file_template_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateCreateRequest); i {
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
		file_template_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateDeleteRequest); i {
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
		file_template_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateBuildLog); i {
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
			RawDescriptor: file_template_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_manager_proto_goTypes,
		DependencyIndexes: file_template_manager_proto_depIdxs,
		MessageInfos:      file_template_manager_proto_msgTypes,
	}.Build()
	File_template_manager_proto = out.File
	file_template_manager_proto_rawDesc = nil
	file_template_manager_proto_goTypes = nil
	file_template_manager_proto_depIdxs = nil
}
