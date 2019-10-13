// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/conversion_adjustment_type.proto

package enums

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

// The different actions advertisers can take to adjust the conversions that
// they already reported. Retractions negate a conversion. Restatements change
// the value of a conversion.
type ConversionAdjustmentTypeEnum_ConversionAdjustmentType int32

const (
	// Not specified.
	ConversionAdjustmentTypeEnum_UNSPECIFIED ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 0
	// Represents value unknown in this version.
	ConversionAdjustmentTypeEnum_UNKNOWN ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 1
	// Negates a conversion so that its total value and count are both zero.
	ConversionAdjustmentTypeEnum_RETRACTION ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 2
	// Changes the value of a conversion.
	ConversionAdjustmentTypeEnum_RESTATEMENT ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 3
)

var ConversionAdjustmentTypeEnum_ConversionAdjustmentType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "RETRACTION",
	3: "RESTATEMENT",
}

var ConversionAdjustmentTypeEnum_ConversionAdjustmentType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"RETRACTION":  2,
	"RESTATEMENT": 3,
}

func (x ConversionAdjustmentTypeEnum_ConversionAdjustmentType) String() string {
	return proto.EnumName(ConversionAdjustmentTypeEnum_ConversionAdjustmentType_name, int32(x))
}

func (ConversionAdjustmentTypeEnum_ConversionAdjustmentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_407eae25c297e5ef, []int{0, 0}
}

// Container for enum describing conversion adjustment types.
type ConversionAdjustmentTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionAdjustmentTypeEnum) Reset()         { *m = ConversionAdjustmentTypeEnum{} }
func (m *ConversionAdjustmentTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionAdjustmentTypeEnum) ProtoMessage()    {}
func (*ConversionAdjustmentTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_407eae25c297e5ef, []int{0}
}

func (m *ConversionAdjustmentTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAdjustmentTypeEnum.Unmarshal(m, b)
}
func (m *ConversionAdjustmentTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAdjustmentTypeEnum.Marshal(b, m, deterministic)
}
func (m *ConversionAdjustmentTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAdjustmentTypeEnum.Merge(m, src)
}
func (m *ConversionAdjustmentTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionAdjustmentTypeEnum.Size(m)
}
func (m *ConversionAdjustmentTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAdjustmentTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAdjustmentTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.ConversionAdjustmentTypeEnum_ConversionAdjustmentType", ConversionAdjustmentTypeEnum_ConversionAdjustmentType_name, ConversionAdjustmentTypeEnum_ConversionAdjustmentType_value)
	proto.RegisterType((*ConversionAdjustmentTypeEnum)(nil), "google.ads.googleads.v1.enums.ConversionAdjustmentTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/conversion_adjustment_type.proto", fileDescriptor_407eae25c297e5ef)
}

var fileDescriptor_407eae25c297e5ef = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xf3, 0x30,
	0x1c, 0xfd, 0xd6, 0xc1, 0x27, 0x64, 0xa0, 0xa5, 0x57, 0x22, 0xdb, 0xc5, 0xf6, 0x00, 0x29, 0xc5,
	0xbb, 0x08, 0x42, 0x36, 0xe3, 0x18, 0x62, 0x37, 0xb6, 0x6c, 0xa2, 0x14, 0x46, 0x5c, 0x42, 0xa8,
	0xac, 0x49, 0x59, 0xb2, 0xc1, 0x5e, 0xc7, 0x4b, 0x1f, 0xc5, 0x47, 0xf1, 0xd2, 0x27, 0x90, 0x26,
	0xae, 0x77, 0xf5, 0xa6, 0x1c, 0x7a, 0xfe, 0xfc, 0x4e, 0x0e, 0xb8, 0x95, 0x5a, 0xcb, 0xad, 0x88,
	0x19, 0x37, 0xb1, 0x87, 0x15, 0x3a, 0x24, 0xb1, 0x50, 0xfb, 0xc2, 0xc4, 0x1b, 0xad, 0x0e, 0x62,
	0x67, 0x72, 0xad, 0xd6, 0x8c, 0xbf, 0xed, 0x8d, 0x2d, 0x84, 0xb2, 0x6b, 0x7b, 0x2c, 0x05, 0x2c,
	0x77, 0xda, 0xea, 0xa8, 0xe7, 0x4d, 0x90, 0x71, 0x03, 0x6b, 0x3f, 0x3c, 0x24, 0xd0, 0xf9, 0xaf,
	0xba, 0xa7, 0xf8, 0x32, 0x8f, 0x99, 0x52, 0xda, 0x32, 0x9b, 0x6b, 0x65, 0xbc, 0x79, 0x70, 0x04,
	0xdd, 0x51, 0x7d, 0x00, 0xd7, 0xf9, 0xf4, 0x58, 0x0a, 0xa2, 0xf6, 0xc5, 0xe0, 0x19, 0x5c, 0x36,
	0xf1, 0xd1, 0x05, 0xe8, 0x2c, 0xd3, 0xc5, 0x8c, 0x8c, 0x26, 0xf7, 0x13, 0x72, 0x17, 0xfe, 0x8b,
	0x3a, 0xe0, 0x6c, 0x99, 0x3e, 0xa4, 0xd3, 0xa7, 0x34, 0x6c, 0x45, 0xe7, 0x00, 0xcc, 0x09, 0x9d,
	0xe3, 0x11, 0x9d, 0x4c, 0xd3, 0x30, 0xa8, 0xd4, 0x73, 0xb2, 0xa0, 0x98, 0x92, 0x47, 0x92, 0xd2,
	0xb0, 0x3d, 0xfc, 0x6e, 0x81, 0xfe, 0x46, 0x17, 0xf0, 0xcf, 0xfa, 0xc3, 0x5e, 0xd3, 0xf9, 0x59,
	0xd5, 0x7f, 0xd6, 0x7a, 0x19, 0xfe, 0xfa, 0xa5, 0xde, 0x32, 0x25, 0xa1, 0xde, 0xc9, 0x58, 0x0a,
	0xe5, 0x5e, 0x77, 0x9a, 0xb3, 0xcc, 0x4d, 0xc3, 0xba, 0x37, 0xee, 0xfb, 0x1e, 0xb4, 0xc7, 0x18,
	0x7f, 0x04, 0xbd, 0xb1, 0x8f, 0xc2, 0xdc, 0x40, 0x0f, 0x2b, 0xb4, 0x4a, 0x60, 0x35, 0x85, 0xf9,
	0x3c, 0xf1, 0x19, 0xe6, 0x26, 0xab, 0xf9, 0x6c, 0x95, 0x64, 0x8e, 0xff, 0x0a, 0xfa, 0xfe, 0x27,
	0x42, 0x98, 0x1b, 0x84, 0x6a, 0x05, 0x42, 0xab, 0x04, 0x21, 0xa7, 0x79, 0xfd, 0xef, 0x8a, 0x5d,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xca, 0xf8, 0xcb, 0xf5, 0x01, 0x00, 0x00,
}
