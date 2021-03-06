// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/date_range_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible date range errors.
type DateRangeErrorEnum_DateRangeError int32

const (
	// Enum unspecified.
	DateRangeErrorEnum_UNSPECIFIED DateRangeErrorEnum_DateRangeError = 0
	// The received error code is not known in this version.
	DateRangeErrorEnum_UNKNOWN DateRangeErrorEnum_DateRangeError = 1
	// Invalid date.
	DateRangeErrorEnum_INVALID_DATE DateRangeErrorEnum_DateRangeError = 2
	// The start date was after the end date.
	DateRangeErrorEnum_START_DATE_AFTER_END_DATE DateRangeErrorEnum_DateRangeError = 3
	// Cannot set date to past time
	DateRangeErrorEnum_CANNOT_SET_DATE_TO_PAST DateRangeErrorEnum_DateRangeError = 4
	// A date was used that is past the system "last" date.
	DateRangeErrorEnum_AFTER_MAXIMUM_ALLOWABLE_DATE DateRangeErrorEnum_DateRangeError = 5
	// Trying to change start date on a campaign that has started.
	DateRangeErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED DateRangeErrorEnum_DateRangeError = 6
)

var DateRangeErrorEnum_DateRangeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_DATE",
	3: "START_DATE_AFTER_END_DATE",
	4: "CANNOT_SET_DATE_TO_PAST",
	5: "AFTER_MAXIMUM_ALLOWABLE_DATE",
	6: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
}
var DateRangeErrorEnum_DateRangeError_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"INVALID_DATE":                 2,
	"START_DATE_AFTER_END_DATE":    3,
	"CANNOT_SET_DATE_TO_PAST":      4,
	"AFTER_MAXIMUM_ALLOWABLE_DATE": 5,
	"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED": 6,
}

func (x DateRangeErrorEnum_DateRangeError) String() string {
	return proto.EnumName(DateRangeErrorEnum_DateRangeError_name, int32(x))
}
func (DateRangeErrorEnum_DateRangeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_date_range_error_46426a282e457061, []int{0, 0}
}

// Container for enum describing possible date range errors.
type DateRangeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateRangeErrorEnum) Reset()         { *m = DateRangeErrorEnum{} }
func (m *DateRangeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DateRangeErrorEnum) ProtoMessage()    {}
func (*DateRangeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_date_range_error_46426a282e457061, []int{0}
}
func (m *DateRangeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRangeErrorEnum.Unmarshal(m, b)
}
func (m *DateRangeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRangeErrorEnum.Marshal(b, m, deterministic)
}
func (dst *DateRangeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRangeErrorEnum.Merge(dst, src)
}
func (m *DateRangeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DateRangeErrorEnum.Size(m)
}
func (m *DateRangeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRangeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DateRangeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DateRangeErrorEnum)(nil), "google.ads.googleads.v0.errors.DateRangeErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.DateRangeErrorEnum_DateRangeError", DateRangeErrorEnum_DateRangeError_name, DateRangeErrorEnum_DateRangeError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/date_range_error.proto", fileDescriptor_date_range_error_46426a282e457061)
}

var fileDescriptor_date_range_error_46426a282e457061 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xe3, 0x40,
	0x1c, 0xc6, 0x37, 0xed, 0x6e, 0x17, 0xa6, 0xcb, 0x6e, 0x98, 0x3d, 0x88, 0xa8, 0x45, 0x7a, 0x15,
	0x26, 0x01, 0xf1, 0xe4, 0xe9, 0xdf, 0xce, 0xa4, 0x04, 0x93, 0x49, 0x48, 0xd2, 0x56, 0x25, 0x30,
	0x44, 0x13, 0x82, 0xd0, 0x66, 0x4a, 0x52, 0xfb, 0x40, 0x1e, 0x7d, 0x13, 0xbd, 0xf8, 0x16, 0x3e,
	0x87, 0x24, 0x93, 0x16, 0x7b, 0xd0, 0x53, 0xbe, 0x7c, 0xf3, 0xfd, 0xe6, 0xcf, 0x7c, 0x7f, 0x74,
	0x91, 0x4b, 0x99, 0x2f, 0x32, 0x23, 0x49, 0x2b, 0x43, 0xc9, 0x5a, 0x6d, 0x4c, 0x23, 0x2b, 0x4b,
	0x59, 0x56, 0x46, 0x9a, 0xac, 0x33, 0x51, 0x26, 0x45, 0x9e, 0x89, 0xc6, 0x21, 0xab, 0x52, 0xae,
	0x25, 0x1e, 0xa8, 0x2c, 0x49, 0xd2, 0x8a, 0xec, 0x30, 0xb2, 0x31, 0x89, 0xc2, 0x86, 0xef, 0x1a,
	0xc2, 0x34, 0x59, 0x67, 0x41, 0x4d, 0xb2, 0xda, 0x63, 0xc5, 0xe3, 0x72, 0xf8, 0xa6, 0xa1, 0xbf,
	0xfb, 0x36, 0xfe, 0x87, 0xfa, 0x53, 0x1e, 0xfa, 0x6c, 0x6c, 0x5b, 0x36, 0xa3, 0xfa, 0x0f, 0xdc,
	0x47, 0xbf, 0xa7, 0xfc, 0x8a, 0x7b, 0x73, 0xae, 0x6b, 0x58, 0x47, 0x7f, 0x6c, 0x3e, 0x03, 0xc7,
	0xa6, 0x82, 0x42, 0xc4, 0xf4, 0x0e, 0x3e, 0x41, 0x87, 0x61, 0x04, 0x41, 0xd4, 0xfc, 0x0b, 0xb0,
	0x22, 0x16, 0x08, 0xc6, 0xdb, 0xe3, 0x2e, 0x3e, 0x42, 0x07, 0x63, 0xe0, 0xdc, 0x8b, 0x44, 0xc8,
	0xda, 0x4c, 0xe4, 0x09, 0x1f, 0xc2, 0x48, 0xff, 0x89, 0x4f, 0xd1, 0xb1, 0x02, 0x5c, 0xb8, 0xb6,
	0xdd, 0xa9, 0x2b, 0xc0, 0x71, 0xbc, 0x39, 0x8c, 0x1c, 0xa6, 0xf0, 0x5f, 0xd8, 0x40, 0x67, 0x2d,
	0xee, 0x7a, 0xd4, 0xb6, 0x6e, 0xc4, 0xa7, 0x59, 0xb6, 0x25, 0xc0, 0x09, 0x18, 0xd0, 0xd6, 0x65,
	0x54, 0xef, 0x8d, 0x5e, 0x34, 0x34, 0xbc, 0x97, 0x4b, 0xf2, 0x7d, 0x1f, 0xa3, 0xff, 0xfb, 0xaf,
	0xf6, 0xeb, 0x12, 0x7d, 0xed, 0x96, 0xb6, 0x58, 0x2e, 0x17, 0x49, 0x91, 0x13, 0x59, 0xe6, 0x46,
	0x9e, 0x15, 0x4d, 0xc5, 0xdb, 0x6d, 0xac, 0x1e, 0xaa, 0xaf, 0x96, 0x73, 0xa9, 0x3e, 0x4f, 0x9d,
	0xee, 0x04, 0xe0, 0xb9, 0x33, 0x98, 0xa8, 0xcb, 0x20, 0xad, 0x88, 0x92, 0xb5, 0x9a, 0x99, 0xa4,
	0x19, 0x59, 0xbd, 0x6e, 0x03, 0x31, 0xa4, 0x55, 0xbc, 0x0b, 0xc4, 0x33, 0x33, 0x56, 0x81, 0xbb,
	0x5e, 0x33, 0xf8, 0xfc, 0x23, 0x00, 0x00, 0xff, 0xff, 0x76, 0x21, 0xc0, 0x89, 0x14, 0x02, 0x00,
	0x00,
}
