// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: set_rules.proto

package set_pb

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

type SetRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set             string   `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Hosts           []string `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	DataShards      uint64   `protobuf:"varint,3,opt,name=data_shards,json=dataShards,proto3" json:"data_shards,omitempty"`
	ParityShards    uint64   `protobuf:"varint,4,opt,name=parity_shards,json=parityShards,proto3" json:"parity_shards,omitempty"`
	MAXShardSize    uint64   `protobuf:"varint,5,opt,name=MAX_shard_size,json=MAXShardSize,proto3" json:"MAX_shard_size,omitempty"`
	ECMode          bool     `protobuf:"varint,6,opt,name=EC_mode,json=ECMode,proto3" json:"EC_mode,omitempty"`
	ReplicationMode bool     `protobuf:"varint,7,opt,name=replication_mode,json=replicationMode,proto3" json:"replication_mode,omitempty"`
}

func (x *SetRules) Reset() {
	*x = SetRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRules) ProtoMessage() {}

func (x *SetRules) ProtoReflect() protoreflect.Message {
	mi := &file_set_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRules.ProtoReflect.Descriptor instead.
func (*SetRules) Descriptor() ([]byte, []int) {
	return file_set_rules_proto_rawDescGZIP(), []int{0}
}

func (x *SetRules) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *SetRules) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *SetRules) GetDataShards() uint64 {
	if x != nil {
		return x.DataShards
	}
	return 0
}

func (x *SetRules) GetParityShards() uint64 {
	if x != nil {
		return x.ParityShards
	}
	return 0
}

func (x *SetRules) GetMAXShardSize() uint64 {
	if x != nil {
		return x.MAXShardSize
	}
	return 0
}

func (x *SetRules) GetECMode() bool {
	if x != nil {
		return x.ECMode
	}
	return false
}

func (x *SetRules) GetReplicationMode() bool {
	if x != nil {
		return x.ReplicationMode
	}
	return false
}

var File_set_rules_proto protoreflect.FileDescriptor

var file_set_rules_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x22, 0xe2, 0x01, 0x0a, 0x08, 0x53, 0x65,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x4d, 0x41, 0x58, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x4d, 0x41,
	0x58, 0x53, 0x68, 0x61, 0x72, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x45, 0x43,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x45, 0x43, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_set_rules_proto_rawDescOnce sync.Once
	file_set_rules_proto_rawDescData = file_set_rules_proto_rawDesc
)

func file_set_rules_proto_rawDescGZIP() []byte {
	file_set_rules_proto_rawDescOnce.Do(func() {
		file_set_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_set_rules_proto_rawDescData)
	})
	return file_set_rules_proto_rawDescData
}

var file_set_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_set_rules_proto_goTypes = []interface{}{
	(*SetRules)(nil), // 0: set_pb.SetRules
}
var file_set_rules_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_set_rules_proto_init() }
func file_set_rules_proto_init() {
	if File_set_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_set_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRules); i {
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
			RawDescriptor: file_set_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_set_rules_proto_goTypes,
		DependencyIndexes: file_set_rules_proto_depIdxs,
		MessageInfos:      file_set_rules_proto_msgTypes,
	}.Build()
	File_set_rules_proto = out.File
	file_set_rules_proto_rawDesc = nil
	file_set_rules_proto_goTypes = nil
	file_set_rules_proto_depIdxs = nil
}
