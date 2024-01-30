// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: api/cluster/cluster.proto

package cluster

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	core "kmesh.net/kmesh/api/v2/core"
	endpoint "kmesh.net/kmesh/api/v2/endpoint"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cluster_LbPolicy int32

const (
	Cluster_ROUND_ROBIN   Cluster_LbPolicy = 0
	Cluster_LEAST_REQUEST Cluster_LbPolicy = 1
	Cluster_RANDOM        Cluster_LbPolicy = 3
)

// Enum value maps for Cluster_LbPolicy.
var (
	Cluster_LbPolicy_name = map[int32]string{
		0: "ROUND_ROBIN",
		1: "LEAST_REQUEST",
		3: "RANDOM",
	}
	Cluster_LbPolicy_value = map[string]int32{
		"ROUND_ROBIN":   0,
		"LEAST_REQUEST": 1,
		"RANDOM":        3,
	}
)

func (x Cluster_LbPolicy) Enum() *Cluster_LbPolicy {
	p := new(Cluster_LbPolicy)
	*p = x
	return p
}

func (x Cluster_LbPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_LbPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_api_cluster_cluster_proto_enumTypes[0].Descriptor()
}

func (Cluster_LbPolicy) Type() protoreflect.EnumType {
	return &file_api_cluster_cluster_proto_enumTypes[0]
}

func (x Cluster_LbPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_LbPolicy.Descriptor instead.
func (Cluster_LbPolicy) EnumDescriptor() ([]byte, []int) {
	return file_api_cluster_cluster_proto_rawDescGZIP(), []int{0, 0}
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiStatus       core.ApiStatus                  `protobuf:"varint,128,opt,name=api_status,json=apiStatus,proto3,enum=core.ApiStatus" json:"api_status,omitempty"`
	Name            string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ConnectTimeout  uint32                          `protobuf:"varint,4,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	LbPolicy        Cluster_LbPolicy                `protobuf:"varint,6,opt,name=lb_policy,json=lbPolicy,proto3,enum=cluster.Cluster_LbPolicy" json:"lb_policy,omitempty"`
	LoadAssignment  *endpoint.ClusterLoadAssignment `protobuf:"bytes,33,opt,name=load_assignment,json=loadAssignment,proto3" json:"load_assignment,omitempty"`
	CircuitBreakers *CircuitBreakers                `protobuf:"bytes,10,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_api_cluster_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetApiStatus() core.ApiStatus {
	if x != nil {
		return x.ApiStatus
	}
	return core.ApiStatus(0)
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetConnectTimeout() uint32 {
	if x != nil {
		return x.ConnectTimeout
	}
	return 0
}

func (x *Cluster) GetLbPolicy() Cluster_LbPolicy {
	if x != nil {
		return x.LbPolicy
	}
	return Cluster_ROUND_ROBIN
}

func (x *Cluster) GetLoadAssignment() *endpoint.ClusterLoadAssignment {
	if x != nil {
		return x.LoadAssignment
	}
	return nil
}

func (x *Cluster) GetCircuitBreakers() *CircuitBreakers {
	if x != nil {
		return x.CircuitBreakers
	}
	return nil
}

var File_api_cluster_cluster_proto protoreflect.FileDescriptor

var file_api_cluster_cluster_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a, 0x07, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x80, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x41, 0x70, 0x69, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x61, 0x70, 0x69,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x6c, 0x62, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x62, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x08, 0x6c, 0x62, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x48, 0x0a, 0x0f, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x10, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x0f, 0x63, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x08, 0x4c, 0x62,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f,
	0x52, 0x4f, 0x42, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x41, 0x53, 0x54,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x41,
	0x4e, 0x44, 0x4f, 0x4d, 0x10, 0x03, 0x42, 0x25, 0x5a, 0x23, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x6e, 0x65, 0x74, 0x2f, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x3b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cluster_cluster_proto_rawDescOnce sync.Once
	file_api_cluster_cluster_proto_rawDescData = file_api_cluster_cluster_proto_rawDesc
)

func file_api_cluster_cluster_proto_rawDescGZIP() []byte {
	file_api_cluster_cluster_proto_rawDescOnce.Do(func() {
		file_api_cluster_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cluster_cluster_proto_rawDescData)
	})
	return file_api_cluster_cluster_proto_rawDescData
}

var file_api_cluster_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_cluster_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_cluster_cluster_proto_goTypes = []interface{}{
	(Cluster_LbPolicy)(0),                  // 0: cluster.Cluster.LbPolicy
	(*Cluster)(nil),                        // 1: cluster.Cluster
	(core.ApiStatus)(0),                    // 2: core.ApiStatus
	(*endpoint.ClusterLoadAssignment)(nil), // 3: endpoint.ClusterLoadAssignment
	(*CircuitBreakers)(nil),                // 4: cluster.CircuitBreakers
}
var file_api_cluster_cluster_proto_depIdxs = []int32{
	2, // 0: cluster.Cluster.api_status:type_name -> core.ApiStatus
	0, // 1: cluster.Cluster.lb_policy:type_name -> cluster.Cluster.LbPolicy
	3, // 2: cluster.Cluster.load_assignment:type_name -> endpoint.ClusterLoadAssignment
	4, // 3: cluster.Cluster.circuit_breakers:type_name -> cluster.CircuitBreakers
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_cluster_cluster_proto_init() }
func file_api_cluster_cluster_proto_init() {
	if File_api_cluster_cluster_proto != nil {
		return
	}
	file_api_cluster_circuit_breaker_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_cluster_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
			RawDescriptor: file_api_cluster_cluster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cluster_cluster_proto_goTypes,
		DependencyIndexes: file_api_cluster_cluster_proto_depIdxs,
		EnumInfos:         file_api_cluster_cluster_proto_enumTypes,
		MessageInfos:      file_api_cluster_cluster_proto_msgTypes,
	}.Build()
	File_api_cluster_cluster_proto = out.File
	file_api_cluster_cluster_proto_rawDesc = nil
	file_api_cluster_cluster_proto_goTypes = nil
	file_api_cluster_cluster_proto_depIdxs = nil
}
