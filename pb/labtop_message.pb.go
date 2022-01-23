// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: labtop_message.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Labtop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory   *Memory    `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are assignable to Weight:
	//	*Labtop_WeightKg
	//	*Labtop_WeightLb
	Weight      isLabtop_Weight      `protobuf_oneof:"weight"`
	PriceUsd    float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Labtop) Reset() {
	*x = Labtop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labtop_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labtop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labtop) ProtoMessage() {}

func (x *Labtop) ProtoReflect() protoreflect.Message {
	mi := &file_labtop_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labtop.ProtoReflect.Descriptor instead.
func (*Labtop) Descriptor() ([]byte, []int) {
	return file_labtop_message_proto_rawDescGZIP(), []int{0}
}

func (x *Labtop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Labtop) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Labtop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Labtop) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Labtop) GetMemory() *Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *Labtop) GetGpus() []*GPU {
	if x != nil {
		return x.Gpus
	}
	return nil
}

func (x *Labtop) GetStorages() []*Storage {
	if x != nil {
		return x.Storages
	}
	return nil
}

func (x *Labtop) GetScreen() *Screen {
	if x != nil {
		return x.Screen
	}
	return nil
}

func (x *Labtop) GetKeyboard() *Keyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

func (m *Labtop) GetWeight() isLabtop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (x *Labtop) GetWeightKg() float64 {
	if x, ok := x.GetWeight().(*Labtop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (x *Labtop) GetWeightLb() float64 {
	if x, ok := x.GetWeight().(*Labtop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (x *Labtop) GetPriceUsd() float64 {
	if x != nil {
		return x.PriceUsd
	}
	return 0
}

func (x *Labtop) GetReleaseYear() uint32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *Labtop) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type isLabtop_Weight interface {
	isLabtop_Weight()
}

type Labtop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Labtop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Labtop_WeightKg) isLabtop_Weight() {}

func (*Labtop_WeightLb) isLabtop_Weight() {}

var File_labtop_message_proto protoreflect.FileDescriptor

var file_labtop_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x62, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65, 0x63, 0x6b, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2,
	0x04, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x65, 0x63, 0x6b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x31, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x74, 0x65, 0x63, 0x6b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x12, 0x2a, 0x0a, 0x04, 0x67, 0x70, 0x75, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x74, 0x65, 0x63, 0x6b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x47, 0x50, 0x55, 0x52, 0x04, 0x67, 0x70, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x65, 0x63, 0x6b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x63, 0x6b, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52,
	0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x63, 0x6b,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4b, 0x65,
	0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x12, 0x1d, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6b, 0x67, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x4b, 0x67, 0x12,
	0x1d, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6c, 0x62, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x62, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x55, 0x73, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x42, 0x1a, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x79, 0x6f, 0x75,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x62, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_labtop_message_proto_rawDescOnce sync.Once
	file_labtop_message_proto_rawDescData = file_labtop_message_proto_rawDesc
)

func file_labtop_message_proto_rawDescGZIP() []byte {
	file_labtop_message_proto_rawDescOnce.Do(func() {
		file_labtop_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_labtop_message_proto_rawDescData)
	})
	return file_labtop_message_proto_rawDescData
}

var file_labtop_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_labtop_message_proto_goTypes = []interface{}{
	(*Labtop)(nil),              // 0: teckschool.pcbook.Labtop
	(*CPU)(nil),                 // 1: teckschool.pcbook.CPU
	(*Memory)(nil),              // 2: teckschool.pcbook.Memory
	(*GPU)(nil),                 // 3: teckschool.pcbook.GPU
	(*Storage)(nil),             // 4: teckschool.pcbook.Storage
	(*Screen)(nil),              // 5: teckschool.pcbook.Screen
	(*Keyboard)(nil),            // 6: teckschool.pcbook.Keyboard
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_labtop_message_proto_depIdxs = []int32{
	1, // 0: teckschool.pcbook.Labtop.cpu:type_name -> teckschool.pcbook.CPU
	2, // 1: teckschool.pcbook.Labtop.memory:type_name -> teckschool.pcbook.Memory
	3, // 2: teckschool.pcbook.Labtop.gpus:type_name -> teckschool.pcbook.GPU
	4, // 3: teckschool.pcbook.Labtop.storages:type_name -> teckschool.pcbook.Storage
	5, // 4: teckschool.pcbook.Labtop.screen:type_name -> teckschool.pcbook.Screen
	6, // 5: teckschool.pcbook.Labtop.keyboard:type_name -> teckschool.pcbook.Keyboard
	7, // 6: teckschool.pcbook.Labtop.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_labtop_message_proto_init() }
func file_labtop_message_proto_init() {
	if File_labtop_message_proto != nil {
		return
	}
	file_memory_message_proto_init()
	file_keyboard_message_proto_init()
	file_processor_message_proto_init()
	file_storage_message_proto_init()
	file_screen_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_labtop_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labtop); i {
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
	file_labtop_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Labtop_WeightKg)(nil),
		(*Labtop_WeightLb)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_labtop_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_labtop_message_proto_goTypes,
		DependencyIndexes: file_labtop_message_proto_depIdxs,
		MessageInfos:      file_labtop_message_proto_msgTypes,
	}.Build()
	File_labtop_message_proto = out.File
	file_labtop_message_proto_rawDesc = nil
	file_labtop_message_proto_goTypes = nil
	file_labtop_message_proto_depIdxs = nil
}
