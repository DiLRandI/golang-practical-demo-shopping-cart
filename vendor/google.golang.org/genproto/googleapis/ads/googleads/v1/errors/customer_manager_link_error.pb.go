// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/customer_manager_link_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible CustomerManagerLink errors.
type CustomerManagerLinkErrorEnum_CustomerManagerLinkError int32

const (
	// Enum unspecified.
	CustomerManagerLinkErrorEnum_UNSPECIFIED CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 0
	// The received error code is not known in this version.
	CustomerManagerLinkErrorEnum_UNKNOWN CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 1
	// No pending invitation.
	CustomerManagerLinkErrorEnum_NO_PENDING_INVITE CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 2
	// Attempt to operate on the same client more than once in the same call.
	CustomerManagerLinkErrorEnum_SAME_CLIENT_MORE_THAN_ONCE_PER_CALL CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 3
	// Manager account has the maximum number of linked accounts.
	CustomerManagerLinkErrorEnum_MANAGER_HAS_MAX_NUMBER_OF_LINKED_ACCOUNTS CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 4
	// If no active user on account it cannot be unlinked from its manager.
	CustomerManagerLinkErrorEnum_CANNOT_UNLINK_ACCOUNT_WITHOUT_ACTIVE_USER CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 5
	// Account should have at least one active owner on it before being
	// unlinked.
	CustomerManagerLinkErrorEnum_CANNOT_REMOVE_LAST_CLIENT_ACCOUNT_OWNER CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 6
	// Only account owners may change their permission role.
	CustomerManagerLinkErrorEnum_CANNOT_CHANGE_ROLE_BY_NON_ACCOUNT_OWNER CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 7
	// When a client's link to its manager is not active, the link role cannot
	// be changed.
	CustomerManagerLinkErrorEnum_CANNOT_CHANGE_ROLE_FOR_NON_ACTIVE_LINK_ACCOUNT CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 8
	// Attempt to link a child to a parent that contains or will contain
	// duplicate children.
	CustomerManagerLinkErrorEnum_DUPLICATE_CHILD_FOUND CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 9
)

var CustomerManagerLinkErrorEnum_CustomerManagerLinkError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NO_PENDING_INVITE",
	3: "SAME_CLIENT_MORE_THAN_ONCE_PER_CALL",
	4: "MANAGER_HAS_MAX_NUMBER_OF_LINKED_ACCOUNTS",
	5: "CANNOT_UNLINK_ACCOUNT_WITHOUT_ACTIVE_USER",
	6: "CANNOT_REMOVE_LAST_CLIENT_ACCOUNT_OWNER",
	7: "CANNOT_CHANGE_ROLE_BY_NON_ACCOUNT_OWNER",
	8: "CANNOT_CHANGE_ROLE_FOR_NON_ACTIVE_LINK_ACCOUNT",
	9: "DUPLICATE_CHILD_FOUND",
}

var CustomerManagerLinkErrorEnum_CustomerManagerLinkError_value = map[string]int32{
	"UNSPECIFIED":                         0,
	"UNKNOWN":                             1,
	"NO_PENDING_INVITE":                   2,
	"SAME_CLIENT_MORE_THAN_ONCE_PER_CALL": 3,
	"MANAGER_HAS_MAX_NUMBER_OF_LINKED_ACCOUNTS":      4,
	"CANNOT_UNLINK_ACCOUNT_WITHOUT_ACTIVE_USER":      5,
	"CANNOT_REMOVE_LAST_CLIENT_ACCOUNT_OWNER":        6,
	"CANNOT_CHANGE_ROLE_BY_NON_ACCOUNT_OWNER":        7,
	"CANNOT_CHANGE_ROLE_FOR_NON_ACTIVE_LINK_ACCOUNT": 8,
	"DUPLICATE_CHILD_FOUND":                          9,
}

func (x CustomerManagerLinkErrorEnum_CustomerManagerLinkError) String() string {
	return proto.EnumName(CustomerManagerLinkErrorEnum_CustomerManagerLinkError_name, int32(x))
}

func (CustomerManagerLinkErrorEnum_CustomerManagerLinkError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb13b694be60d601, []int{0, 0}
}

// Container for enum describing possible CustomerManagerLink errors.
type CustomerManagerLinkErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerManagerLinkErrorEnum) Reset()         { *m = CustomerManagerLinkErrorEnum{} }
func (m *CustomerManagerLinkErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerManagerLinkErrorEnum) ProtoMessage()    {}
func (*CustomerManagerLinkErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb13b694be60d601, []int{0}
}

func (m *CustomerManagerLinkErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerManagerLinkErrorEnum.Unmarshal(m, b)
}
func (m *CustomerManagerLinkErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerManagerLinkErrorEnum.Marshal(b, m, deterministic)
}
func (m *CustomerManagerLinkErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerManagerLinkErrorEnum.Merge(m, src)
}
func (m *CustomerManagerLinkErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerManagerLinkErrorEnum.Size(m)
}
func (m *CustomerManagerLinkErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerManagerLinkErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerManagerLinkErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.CustomerManagerLinkErrorEnum_CustomerManagerLinkError", CustomerManagerLinkErrorEnum_CustomerManagerLinkError_name, CustomerManagerLinkErrorEnum_CustomerManagerLinkError_value)
	proto.RegisterType((*CustomerManagerLinkErrorEnum)(nil), "google.ads.googleads.v1.errors.CustomerManagerLinkErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/customer_manager_link_error.proto", fileDescriptor_bb13b694be60d601)
}

var fileDescriptor_bb13b694be60d601 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x69, 0x0a, 0x2d, 0x6c, 0x0f, 0x18, 0x4b, 0x95, 0x00, 0x95, 0x1e, 0xc2, 0xa1, 0x42,
	0x08, 0x5b, 0x81, 0x9b, 0xb9, 0xb0, 0x59, 0x4f, 0x12, 0xab, 0xf6, 0xac, 0xe5, 0x7f, 0x01, 0x14,
	0x69, 0x64, 0x9a, 0xc8, 0x8a, 0x9a, 0x78, 0x23, 0x3b, 0xed, 0xf3, 0x20, 0x8e, 0x3c, 0x0a, 0x8f,
	0xc2, 0x13, 0x70, 0x03, 0x39, 0x1b, 0x47, 0x15, 0x22, 0x9c, 0x76, 0xb4, 0xf3, 0x9b, 0xef, 0xfb,
	0xa4, 0x19, 0xf6, 0xa1, 0x50, 0xaa, 0x58, 0xcc, 0xec, 0x7c, 0x5a, 0xdb, 0xba, 0x6c, 0xaa, 0xdb,
	0x9e, 0x3d, 0xab, 0x2a, 0x55, 0xd5, 0xf6, 0xd5, 0x4d, 0xbd, 0x56, 0xcb, 0x59, 0x45, 0xcb, 0xbc,
	0xcc, 0x8b, 0x59, 0x45, 0x8b, 0x79, 0x79, 0x4d, 0x9b, 0xa6, 0xb5, 0xaa, 0xd4, 0x5a, 0x99, 0xe7,
	0x7a, 0xcc, 0xca, 0xa7, 0xb5, 0xb5, 0x53, 0xb0, 0x6e, 0x7b, 0x96, 0x56, 0x78, 0x7e, 0xd6, 0x3a,
	0xac, 0xe6, 0x76, 0x5e, 0x96, 0x6a, 0x9d, 0xaf, 0xe7, 0xaa, 0xac, 0xf5, 0x74, 0xf7, 0xeb, 0x21,
	0x3b, 0x13, 0x5b, 0x8f, 0x40, 0x5b, 0xf8, 0xf3, 0xf2, 0x1a, 0x9a, 0x59, 0x28, 0x6f, 0x96, 0xdd,
	0xdf, 0x1d, 0xf6, 0x74, 0x1f, 0x60, 0x3e, 0x66, 0x27, 0x29, 0xc6, 0x21, 0x08, 0x6f, 0xe0, 0x81,
	0x6b, 0xdc, 0x33, 0x4f, 0xd8, 0x71, 0x8a, 0x97, 0x28, 0xc7, 0x68, 0x1c, 0x98, 0xa7, 0xec, 0x09,
	0x4a, 0x0a, 0x01, 0x5d, 0x0f, 0x87, 0xe4, 0x61, 0xe6, 0x25, 0x60, 0x74, 0xcc, 0x0b, 0xf6, 0x32,
	0xe6, 0x01, 0x90, 0xf0, 0x3d, 0xc0, 0x84, 0x02, 0x19, 0x01, 0x25, 0x23, 0x8e, 0x24, 0x51, 0x00,
	0x85, 0x10, 0x91, 0xe0, 0xbe, 0x6f, 0x1c, 0x9a, 0x6f, 0xd8, 0xab, 0x80, 0x23, 0x1f, 0x42, 0x44,
	0x23, 0x1e, 0x53, 0xc0, 0x3f, 0x12, 0xa6, 0x41, 0x1f, 0x22, 0x92, 0x03, 0xf2, 0x3d, 0xbc, 0x04,
	0x97, 0xb8, 0x10, 0x32, 0xc5, 0x24, 0x36, 0xee, 0x37, 0xb8, 0xe0, 0x88, 0x32, 0xa1, 0x14, 0x9b,
	0x6e, 0xdb, 0xa3, 0xb1, 0x97, 0x8c, 0x64, 0x9a, 0x10, 0x17, 0x89, 0x97, 0x01, 0xa5, 0x31, 0x44,
	0xc6, 0x03, 0xf3, 0x35, 0xbb, 0xd8, 0xe2, 0x11, 0x04, 0x32, 0x03, 0xf2, 0x79, 0x9c, 0xb4, 0xa1,
	0xda, 0x51, 0x39, 0x46, 0x88, 0x8c, 0xa3, 0x3b, 0xb0, 0x18, 0x71, 0x1c, 0x02, 0x45, 0xd2, 0x07,
	0xea, 0x7f, 0x22, 0x94, 0xf8, 0x17, 0x7c, 0x6c, 0xbe, 0x65, 0xd6, 0x3f, 0xe0, 0x81, 0x8c, 0xb6,
	0xf4, 0x26, 0xc5, 0xdd, 0x84, 0xc6, 0x43, 0xf3, 0x19, 0x3b, 0x75, 0xd3, 0xd0, 0xf7, 0x04, 0x4f,
	0x80, 0xc4, 0xc8, 0xf3, 0x5d, 0x1a, 0xc8, 0x14, 0x5d, 0xe3, 0x51, 0xff, 0xd7, 0x01, 0xeb, 0x5e,
	0xa9, 0xa5, 0xf5, 0xff, 0x3d, 0xf7, 0x5f, 0xec, 0xdb, 0x52, 0xd8, 0x2c, 0x3a, 0x3c, 0xf8, 0xec,
	0x6e, 0x05, 0x0a, 0xb5, 0xc8, 0xcb, 0xc2, 0x52, 0x55, 0x61, 0x17, 0xb3, 0x72, 0x73, 0x06, 0xed,
	0xe9, 0xad, 0xe6, 0xf5, 0xbe, 0x4b, 0x7c, 0xaf, 0x9f, 0x6f, 0x9d, 0xc3, 0x21, 0xe7, 0xdf, 0x3b,
	0xe7, 0x43, 0x2d, 0xc6, 0xa7, 0xb5, 0xa5, 0xcb, 0xa6, 0xca, 0x7a, 0xd6, 0xc6, 0xb2, 0xfe, 0xd1,
	0x02, 0x13, 0x3e, 0xad, 0x27, 0x3b, 0x60, 0x92, 0xf5, 0x26, 0x1a, 0xf8, 0xd9, 0xe9, 0xea, 0x5f,
	0xc7, 0xe1, 0xd3, 0xda, 0x71, 0x76, 0x88, 0xe3, 0x64, 0x3d, 0xc7, 0xd1, 0xd0, 0x97, 0xa3, 0x4d,
	0xba, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xcb, 0x93, 0x3c, 0x26, 0x03, 0x00, 0x00,
}
