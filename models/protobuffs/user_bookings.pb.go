// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: user_bookings.proto

package protobuffs

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

type UserBookingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserBookingsRequest) Reset() {
	*x = UserBookingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_bookings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBookingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBookingsRequest) ProtoMessage() {}

func (x *UserBookingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_bookings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBookingsRequest.ProtoReflect.Descriptor instead.
func (*UserBookingsRequest) Descriptor() ([]byte, []int) {
	return file_user_bookings_proto_rawDescGZIP(), []int{0}
}

func (x *UserBookingsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserBookingsRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UserBookingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*Booking `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
	ErrorMsg string     `protobuf:"bytes,2,opt,name=ErrorMsg,proto3" json:"ErrorMsg,omitempty"`
}

func (x *UserBookingsResponse) Reset() {
	*x = UserBookingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_bookings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBookingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBookingsResponse) ProtoMessage() {}

func (x *UserBookingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_bookings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBookingsResponse.ProtoReflect.Descriptor instead.
func (*UserBookingsResponse) Descriptor() ([]byte, []int) {
	return file_user_bookings_proto_rawDescGZIP(), []int{1}
}

func (x *UserBookingsResponse) GetBookings() []*Booking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

func (x *UserBookingsResponse) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

var File_user_bookings_proto protoreflect.FileDescriptor

var file_user_bookings_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66,
	0x73, 0x1a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x45, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x63, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x32, 0x66, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x56, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_bookings_proto_rawDescOnce sync.Once
	file_user_bookings_proto_rawDescData = file_user_bookings_proto_rawDesc
)

func file_user_bookings_proto_rawDescGZIP() []byte {
	file_user_bookings_proto_rawDescOnce.Do(func() {
		file_user_bookings_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_bookings_proto_rawDescData)
	})
	return file_user_bookings_proto_rawDescData
}

var file_user_bookings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_bookings_proto_goTypes = []interface{}{
	(*UserBookingsRequest)(nil),  // 0: protobuffs.UserBookingsRequest
	(*UserBookingsResponse)(nil), // 1: protobuffs.UserBookingsResponse
	(*Booking)(nil),              // 2: protobuffs.Booking
}
var file_user_bookings_proto_depIdxs = []int32{
	2, // 0: protobuffs.UserBookingsResponse.bookings:type_name -> protobuffs.Booking
	0, // 1: protobuffs.UserBookings.GetUserBookings:input_type -> protobuffs.UserBookingsRequest
	1, // 2: protobuffs.UserBookings.GetUserBookings:output_type -> protobuffs.UserBookingsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_bookings_proto_init() }
func file_user_bookings_proto_init() {
	if File_user_bookings_proto != nil {
		return
	}
	file_booking_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_bookings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBookingsRequest); i {
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
		file_user_bookings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBookingsResponse); i {
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
			RawDescriptor: file_user_bookings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_bookings_proto_goTypes,
		DependencyIndexes: file_user_bookings_proto_depIdxs,
		MessageInfos:      file_user_bookings_proto_msgTypes,
	}.Build()
	File_user_bookings_proto = out.File
	file_user_bookings_proto_rawDesc = nil
	file_user_bookings_proto_goTypes = nil
	file_user_bookings_proto_depIdxs = nil
}
