// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: protos/userproto/user.proto

package userproto

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

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_userproto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_userproto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_protos_userproto_user_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Admin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*User `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_userproto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_protos_userproto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_protos_userproto_user_proto_rawDescGZIP(), []int{1}
}

func (x *Users) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

type PayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId int32 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Price   int32 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PayRequest) Reset() {
	*x = PayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_userproto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayRequest) ProtoMessage() {}

func (x *PayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_userproto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayRequest.ProtoReflect.Descriptor instead.
func (*PayRequest) Descriptor() ([]byte, []int) {
	return file_protos_userproto_user_proto_rawDescGZIP(), []int{2}
}

func (x *PayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PayRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Budget   int32  `protobuf:"varint,3,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_userproto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_protos_userproto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_protos_userproto_user_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetBudget() int32 {
	if x != nil {
		return x.Budget
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK string `protobuf:"bytes,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_userproto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protos_userproto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_protos_userproto_user_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetOK() string {
	if x != nil {
		return x.OK
	}
	return ""
}

var File_protos_userproto_user_proto protoreflect.FileDescriptor

var file_protos_userproto_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x4a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x1a,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x4b, 0x32, 0xbc, 0x01, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x4c, 0x6f,
	0x67, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x62, 0x61, 0x72, 0x73, 0x68, 0x6f,
	0x68, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x4f, 0x4c, 0x58, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_userproto_user_proto_rawDescOnce sync.Once
	file_protos_userproto_user_proto_rawDescData = file_protos_userproto_user_proto_rawDesc
)

func file_protos_userproto_user_proto_rawDescGZIP() []byte {
	file_protos_userproto_user_proto_rawDescOnce.Do(func() {
		file_protos_userproto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_userproto_user_proto_rawDescData)
	})
	return file_protos_userproto_user_proto_rawDescData
}

var file_protos_userproto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_userproto_user_proto_goTypes = []interface{}{
	(*Admin)(nil),      // 0: proto.Admin
	(*Users)(nil),      // 1: proto.Users
	(*PayRequest)(nil), // 2: proto.payRequest
	(*User)(nil),       // 3: proto.User
	(*Response)(nil),   // 4: proto.Response
}
var file_protos_userproto_user_proto_depIdxs = []int32{
	3, // 0: proto.Users.List:type_name -> proto.User
	3, // 1: proto.UserService.Log:input_type -> proto.User
	2, // 2: proto.UserService.Payment:input_type -> proto.payRequest
	0, // 3: proto.UserService.UserList:input_type -> proto.Admin
	0, // 4: proto.UserService.LogAdmin:input_type -> proto.Admin
	4, // 5: proto.UserService.Log:output_type -> proto.Response
	4, // 6: proto.UserService.Payment:output_type -> proto.Response
	1, // 7: proto.UserService.UserList:output_type -> proto.Users
	4, // 8: proto.UserService.LogAdmin:output_type -> proto.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_userproto_user_proto_init() }
func file_protos_userproto_user_proto_init() {
	if File_protos_userproto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_userproto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_protos_userproto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_protos_userproto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayRequest); i {
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
		file_protos_userproto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_protos_userproto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_protos_userproto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_userproto_user_proto_goTypes,
		DependencyIndexes: file_protos_userproto_user_proto_depIdxs,
		MessageInfos:      file_protos_userproto_user_proto_msgTypes,
	}.Build()
	File_protos_userproto_user_proto = out.File
	file_protos_userproto_user_proto_rawDesc = nil
	file_protos_userproto_user_proto_goTypes = nil
	file_protos_userproto_user_proto_depIdxs = nil
}
