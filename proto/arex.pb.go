// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: proto/arex.proto

package proto

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

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Denom  string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Asking string `protobuf:"bytes,4,opt,name=asking,proto3" json:"asking,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_arex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_arex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_proto_arex_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invoice) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Invoice) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Invoice) GetAsking() string {
	if x != nil {
		return x.Asking
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_arex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_arex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_arex_proto_rawDescGZIP(), []int{1}
}

type GetAllInvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Invoice `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllInvoicesResponse) Reset() {
	*x = GetAllInvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_arex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllInvoicesResponse) ProtoMessage() {}

func (x *GetAllInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_arex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllInvoicesResponse.ProtoReflect.Descriptor instead.
func (*GetAllInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_proto_arex_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllInvoicesResponse) GetData() []*Invoice {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_arex_proto protoreflect.FileDescriptor

var file_proto_arex_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x41, 0x0a, 0x0e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72,
	0x63, 0x6f, 0x73, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x61, 0x73, 0x74, 0x61, 0x73, 0x69, 0x2f, 0x61,
	0x72, 0x65, 0x78, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_arex_proto_rawDescOnce sync.Once
	file_proto_arex_proto_rawDescData = file_proto_arex_proto_rawDesc
)

func file_proto_arex_proto_rawDescGZIP() []byte {
	file_proto_arex_proto_rawDescOnce.Do(func() {
		file_proto_arex_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_arex_proto_rawDescData)
	})
	return file_proto_arex_proto_rawDescData
}

var file_proto_arex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_arex_proto_goTypes = []interface{}{
	(*Invoice)(nil),                // 0: v1.Invoice
	(*Empty)(nil),                  // 1: v1.Empty
	(*GetAllInvoicesResponse)(nil), // 2: v1.GetAllInvoicesResponse
}
var file_proto_arex_proto_depIdxs = []int32{
	0, // 0: v1.GetAllInvoicesResponse.data:type_name -> v1.Invoice
	1, // 1: v1.InvoiceService.GetAll:input_type -> v1.Empty
	2, // 2: v1.InvoiceService.GetAll:output_type -> v1.GetAllInvoicesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_arex_proto_init() }
func file_proto_arex_proto_init() {
	if File_proto_arex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_arex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_proto_arex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_arex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllInvoicesResponse); i {
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
			RawDescriptor: file_proto_arex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_arex_proto_goTypes,
		DependencyIndexes: file_proto_arex_proto_depIdxs,
		MessageInfos:      file_proto_arex_proto_msgTypes,
	}.Build()
	File_proto_arex_proto = out.File
	file_proto_arex_proto_rawDesc = nil
	file_proto_arex_proto_goTypes = nil
	file_proto_arex_proto_depIdxs = nil
}
