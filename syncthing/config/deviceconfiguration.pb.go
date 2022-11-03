// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: config/deviceconfiguration.proto

package config

import (
	protocol "github.com/zebfross/syncthing-mobile/syncthing/protocol"
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

type DeviceConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId                 []byte               `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name                     string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Addresses                []string             `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Compression              protocol.Compression `protobuf:"varint,4,opt,name=compression,proto3,enum=protocol.Compression" json:"compression,omitempty"`
	CertName                 string               `protobuf:"bytes,5,opt,name=cert_name,json=certName,proto3" json:"cert_name,omitempty"`
	Introducer               bool                 `protobuf:"varint,6,opt,name=introducer,proto3" json:"introducer,omitempty"`
	SkipIntroductionRemovals bool                 `protobuf:"varint,7,opt,name=skip_introduction_removals,json=skipIntroductionRemovals,proto3" json:"skip_introduction_removals,omitempty"`
	IntroducedBy             []byte               `protobuf:"bytes,8,opt,name=introduced_by,json=introducedBy,proto3" json:"introduced_by,omitempty"`
	Paused                   bool                 `protobuf:"varint,9,opt,name=paused,proto3" json:"paused,omitempty"`
	AllowedNetworks          []string             `protobuf:"bytes,10,rep,name=allowed_networks,json=allowedNetworks,proto3" json:"allowed_networks,omitempty"`
	AutoAcceptFolders        bool                 `protobuf:"varint,11,opt,name=auto_accept_folders,json=autoAcceptFolders,proto3" json:"auto_accept_folders,omitempty"`
	MaxSendKbps              int32                `protobuf:"varint,12,opt,name=max_send_kbps,json=maxSendKbps,proto3" json:"max_send_kbps,omitempty"`
	MaxRecvKbps              int32                `protobuf:"varint,13,opt,name=max_recv_kbps,json=maxRecvKbps,proto3" json:"max_recv_kbps,omitempty"`
	IgnoredFolders           []*ObservedFolder    `protobuf:"bytes,14,rep,name=ignored_folders,json=ignoredFolders,proto3" json:"ignored_folders,omitempty"`
	PendingFolders           []*ObservedFolder    `protobuf:"bytes,15,rep,name=pending_folders,json=pendingFolders,proto3" json:"pending_folders,omitempty"`
	MaxRequestKib            int32                `protobuf:"varint,16,opt,name=max_request_kib,json=maxRequestKib,proto3" json:"max_request_kib,omitempty"`
	Untrusted                bool                 `protobuf:"varint,17,opt,name=untrusted,proto3" json:"untrusted,omitempty"`
	RemoteGuiPort            int32                `protobuf:"varint,18,opt,name=remote_gui_port,json=remoteGuiPort,proto3" json:"remote_gui_port,omitempty"`
}

func (x *DeviceConfiguration) Reset() {
	*x = DeviceConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_deviceconfiguration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceConfiguration) ProtoMessage() {}

func (x *DeviceConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_config_deviceconfiguration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceConfiguration.ProtoReflect.Descriptor instead.
func (*DeviceConfiguration) Descriptor() ([]byte, []int) {
	return file_config_deviceconfiguration_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceConfiguration) GetDeviceId() []byte {
	if x != nil {
		return x.DeviceId
	}
	return nil
}

func (x *DeviceConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceConfiguration) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *DeviceConfiguration) GetCompression() protocol.Compression {
	if x != nil {
		return x.Compression
	}
	return protocol.Compression(0)
}

func (x *DeviceConfiguration) GetCertName() string {
	if x != nil {
		return x.CertName
	}
	return ""
}

func (x *DeviceConfiguration) GetIntroducer() bool {
	if x != nil {
		return x.Introducer
	}
	return false
}

func (x *DeviceConfiguration) GetSkipIntroductionRemovals() bool {
	if x != nil {
		return x.SkipIntroductionRemovals
	}
	return false
}

func (x *DeviceConfiguration) GetIntroducedBy() []byte {
	if x != nil {
		return x.IntroducedBy
	}
	return nil
}

func (x *DeviceConfiguration) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

func (x *DeviceConfiguration) GetAllowedNetworks() []string {
	if x != nil {
		return x.AllowedNetworks
	}
	return nil
}

func (x *DeviceConfiguration) GetAutoAcceptFolders() bool {
	if x != nil {
		return x.AutoAcceptFolders
	}
	return false
}

func (x *DeviceConfiguration) GetMaxSendKbps() int32 {
	if x != nil {
		return x.MaxSendKbps
	}
	return 0
}

func (x *DeviceConfiguration) GetMaxRecvKbps() int32 {
	if x != nil {
		return x.MaxRecvKbps
	}
	return 0
}

func (x *DeviceConfiguration) GetIgnoredFolders() []*ObservedFolder {
	if x != nil {
		return x.IgnoredFolders
	}
	return nil
}

func (x *DeviceConfiguration) GetPendingFolders() []*ObservedFolder {
	if x != nil {
		return x.PendingFolders
	}
	return nil
}

func (x *DeviceConfiguration) GetMaxRequestKib() int32 {
	if x != nil {
		return x.MaxRequestKib
	}
	return 0
}

func (x *DeviceConfiguration) GetUntrusted() bool {
	if x != nil {
		return x.Untrusted
	}
	return false
}

func (x *DeviceConfiguration) GetRemoteGuiPort() int32 {
	if x != nil {
		return x.RemoteGuiPort
	}
	return 0
}

var File_config_deviceconfiguration_proto protoreflect.FileDescriptor

var file_config_deviceconfiguration_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x62, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x05, 0x0a, 0x13, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x12, 0x3c, 0x0a, 0x1a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x6e, 0x74, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x64, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6b, 0x62, 0x70, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x62, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x6b, 0x62, 0x70, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x63, 0x76, 0x4b, 0x62, 0x70, 0x73,
	0x12, 0x3f, 0x0a, 0x0f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x0e, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6b, 0x69, 0x62, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x69, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x75,
	0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x47, 0x75, 0x69, 0x50, 0x6f, 0x72, 0x74,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x65, 0x62, 0x66, 0x72, 0x6f, 0x73, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_deviceconfiguration_proto_rawDescOnce sync.Once
	file_config_deviceconfiguration_proto_rawDescData = file_config_deviceconfiguration_proto_rawDesc
)

func file_config_deviceconfiguration_proto_rawDescGZIP() []byte {
	file_config_deviceconfiguration_proto_rawDescOnce.Do(func() {
		file_config_deviceconfiguration_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_deviceconfiguration_proto_rawDescData)
	})
	return file_config_deviceconfiguration_proto_rawDescData
}

var file_config_deviceconfiguration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_deviceconfiguration_proto_goTypes = []interface{}{
	(*DeviceConfiguration)(nil), // 0: config.DeviceConfiguration
	(protocol.Compression)(0),   // 1: protocol.Compression
	(*ObservedFolder)(nil),      // 2: config.ObservedFolder
}
var file_config_deviceconfiguration_proto_depIdxs = []int32{
	1, // 0: config.DeviceConfiguration.compression:type_name -> protocol.Compression
	2, // 1: config.DeviceConfiguration.ignored_folders:type_name -> config.ObservedFolder
	2, // 2: config.DeviceConfiguration.pending_folders:type_name -> config.ObservedFolder
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_deviceconfiguration_proto_init() }
func file_config_deviceconfiguration_proto_init() {
	if File_config_deviceconfiguration_proto != nil {
		return
	}
	file_config_observed_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_deviceconfiguration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceConfiguration); i {
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
			RawDescriptor: file_config_deviceconfiguration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_deviceconfiguration_proto_goTypes,
		DependencyIndexes: file_config_deviceconfiguration_proto_depIdxs,
		MessageInfos:      file_config_deviceconfiguration_proto_msgTypes,
	}.Build()
	File_config_deviceconfiguration_proto = out.File
	file_config_deviceconfiguration_proto_rawDesc = nil
	file_config_deviceconfiguration_proto_goTypes = nil
	file_config_deviceconfiguration_proto_depIdxs = nil
}
