// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: cheque.proto

package pb

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

type Cheque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value           string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	TokenAddress    string `protobuf:"bytes,2,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"` // token
	Nonce           string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	From            string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`                                              // user
	To              string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`                                                  // storage
	OperatorAddress string `protobuf:"bytes,6,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"` // operator
	ContractAddress string `protobuf:"bytes,7,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"` //运营合约地址
}

func (x *Cheque) Reset() {
	*x = Cheque{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheque_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cheque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cheque) ProtoMessage() {}

func (x *Cheque) ProtoReflect() protoreflect.Message {
	mi := &file_cheque_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cheque.ProtoReflect.Descriptor instead.
func (*Cheque) Descriptor() ([]byte, []int) {
	return file_cheque_proto_rawDescGZIP(), []int{0}
}

func (x *Cheque) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Cheque) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *Cheque) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Cheque) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Cheque) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Cheque) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

func (x *Cheque) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

type PayCheque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cheque    *Cheque `protobuf:"bytes,1,opt,name=cheque,proto3" json:"cheque,omitempty"`
	ChequeSig []byte  `protobuf:"bytes,2,opt,name=cheque_sig,json=chequeSig,proto3" json:"cheque_sig,omitempty"` //运营商对cheque的签名
	PayValue  string  `protobuf:"bytes,3,opt,name=pay_value,json=payValue,proto3" json:"pay_value,omitempty"`    //支付给存储节点的数额必须小于等于cheque.max_amount
}

func (x *PayCheque) Reset() {
	*x = PayCheque{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheque_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayCheque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayCheque) ProtoMessage() {}

func (x *PayCheque) ProtoReflect() protoreflect.Message {
	mi := &file_cheque_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayCheque.ProtoReflect.Descriptor instead.
func (*PayCheque) Descriptor() ([]byte, []int) {
	return file_cheque_proto_rawDescGZIP(), []int{1}
}

func (x *PayCheque) GetCheque() *Cheque {
	if x != nil {
		return x.Cheque
	}
	return nil
}

func (x *PayCheque) GetChequeSig() []byte {
	if x != nil {
		return x.ChequeSig
	}
	return nil
}

func (x *PayCheque) GetPayValue() string {
	if x != nil {
		return x.PayValue
	}
	return ""
}

type PayChequeBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChequeBatchTo    string `protobuf:"bytes,1,opt,name=cheque_batch_to,json=chequeBatchTo,proto3" json:"cheque_batch_to,omitempty"`           //存储节点号
	ChequeBatchValue int64  `protobuf:"varint,2,opt,name=cheque_batch_value,json=chequeBatchValue,proto3" json:"cheque_batch_value,omitempty"` //聚合后的支票面额
	MinNonce         int64  `protobuf:"varint,3,opt,name=min_nonce,json=minNonce,proto3" json:"min_nonce,omitempty"`                           //聚合的nonce最小值
	MaxNonce         int64  `protobuf:"varint,4,opt,name=max_nonce,json=maxNonce,proto3" json:"max_nonce,omitempty"`                           //聚合的nonce最大值
}

func (x *PayChequeBatch) Reset() {
	*x = PayChequeBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheque_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayChequeBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayChequeBatch) ProtoMessage() {}

func (x *PayChequeBatch) ProtoReflect() protoreflect.Message {
	mi := &file_cheque_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayChequeBatch.ProtoReflect.Descriptor instead.
func (*PayChequeBatch) Descriptor() ([]byte, []int) {
	return file_cheque_proto_rawDescGZIP(), []int{2}
}

func (x *PayChequeBatch) GetChequeBatchTo() string {
	if x != nil {
		return x.ChequeBatchTo
	}
	return ""
}

func (x *PayChequeBatch) GetChequeBatchValue() int64 {
	if x != nil {
		return x.ChequeBatchValue
	}
	return 0
}

func (x *PayChequeBatch) GetMinNonce() int64 {
	if x != nil {
		return x.MinNonce
	}
	return 0
}

func (x *PayChequeBatch) GetMaxNonce() int64 {
	if x != nil {
		return x.MaxNonce
	}
	return 0
}

var File_cheque_proto protoreflect.FileDescriptor

var file_cheque_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xd3, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x65, 0x71, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6b, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x43,
	0x68, 0x65, 0x71, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x71, 0x75,
	0x65, 0x52, 0x06, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x65,
	0x71, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63,
	0x68, 0x65, 0x71, 0x75, 0x65, 0x53, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x65,
	0x71, 0x75, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x71,
	0x75, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x6f,
	0x12, 0x2c, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x68,
	0x65, 0x71, 0x75, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheque_proto_rawDescOnce sync.Once
	file_cheque_proto_rawDescData = file_cheque_proto_rawDesc
)

func file_cheque_proto_rawDescGZIP() []byte {
	file_cheque_proto_rawDescOnce.Do(func() {
		file_cheque_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheque_proto_rawDescData)
	})
	return file_cheque_proto_rawDescData
}

var file_cheque_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cheque_proto_goTypes = []interface{}{
	(*Cheque)(nil),         // 0: pb.Cheque
	(*PayCheque)(nil),      // 1: pb.PayCheque
	(*PayChequeBatch)(nil), // 2: pb.PayChequeBatch
}
var file_cheque_proto_depIdxs = []int32{
	0, // 0: pb.PayCheque.cheque:type_name -> pb.Cheque
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cheque_proto_init() }
func file_cheque_proto_init() {
	if File_cheque_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheque_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cheque); i {
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
		file_cheque_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayCheque); i {
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
		file_cheque_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayChequeBatch); i {
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
			RawDescriptor: file_cheque_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheque_proto_goTypes,
		DependencyIndexes: file_cheque_proto_depIdxs,
		MessageInfos:      file_cheque_proto_msgTypes,
	}.Build()
	File_cheque_proto = out.File
	file_cheque_proto_rawDesc = nil
	file_cheque_proto_goTypes = nil
	file_cheque_proto_depIdxs = nil
}
