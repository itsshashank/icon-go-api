// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.20.1
// source: contract_list.proto

package models

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

type ContractList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	TransactionCount int64   `protobuf:"varint,2,opt,name=transaction_count,json=transactionCount,proto3" json:"transaction_count"`
	LogCount         int64   `protobuf:"varint,3,opt,name=log_count,json=logCount,proto3" json:"log_count"`
	Balance          float64 `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance"`
	Name             string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Status           string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	CreatedTimestamp int64   `protobuf:"varint,7,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp"`
	IsToken          bool    `protobuf:"varint,8,opt,name=is_token,json=isToken,proto3" json:"is_token"`
	TokenStandard    string  `protobuf:"bytes,9,opt,name=token_standard,json=tokenStandard,proto3" json:"token_standard"`
}

func (x *ContractList) Reset() {
	*x = ContractList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractList) ProtoMessage() {}

func (x *ContractList) ProtoReflect() protoreflect.Message {
	mi := &file_contract_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractList.ProtoReflect.Descriptor instead.
func (*ContractList) Descriptor() ([]byte, []int) {
	return file_contract_list_proto_rawDescGZIP(), []int{0}
}

func (x *ContractList) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContractList) GetTransactionCount() int64 {
	if x != nil {
		return x.TransactionCount
	}
	return 0
}

func (x *ContractList) GetLogCount() int64 {
	if x != nil {
		return x.LogCount
	}
	return 0
}

func (x *ContractList) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *ContractList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContractList) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ContractList) GetCreatedTimestamp() int64 {
	if x != nil {
		return x.CreatedTimestamp
	}
	return 0
}

func (x *ContractList) GetIsToken() bool {
	if x != nil {
		return x.IsToken
	}
	return false
}

func (x *ContractList) GetTokenStandard() string {
	if x != nil {
		return x.TokenStandard
	}
	return ""
}

var File_contract_list_proto protoreflect.FileDescriptor

var file_contract_list_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xa7, 0x02,
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_list_proto_rawDescOnce sync.Once
	file_contract_list_proto_rawDescData = file_contract_list_proto_rawDesc
)

func file_contract_list_proto_rawDescGZIP() []byte {
	file_contract_list_proto_rawDescOnce.Do(func() {
		file_contract_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_list_proto_rawDescData)
	})
	return file_contract_list_proto_rawDescData
}

var file_contract_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contract_list_proto_goTypes = []interface{}{
	(*ContractList)(nil), // 0: models.ContractList
}
var file_contract_list_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contract_list_proto_init() }
func file_contract_list_proto_init() {
	if File_contract_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractList); i {
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
			RawDescriptor: file_contract_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_list_proto_goTypes,
		DependencyIndexes: file_contract_list_proto_depIdxs,
		MessageInfos:      file_contract_list_proto_msgTypes,
	}.Build()
	File_contract_list_proto = out.File
	file_contract_list_proto_rawDesc = nil
	file_contract_list_proto_goTypes = nil
	file_contract_list_proto_depIdxs = nil
}
