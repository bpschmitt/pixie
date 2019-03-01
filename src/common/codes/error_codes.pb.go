// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/common/codes/error_codes.proto

package codepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Code int32

const (
	OK                   Code = 0
	CANCELLED            Code = 1
	UNKNOWN              Code = 2
	INVALID_ARGUMENT     Code = 3
	DEADLINE_EXCEEDED    Code = 4
	NOT_FOUND            Code = 5
	ALREADY_EXISTS       Code = 6
	PERMISSION_DENIED    Code = 7
	UNAUTHENTICATED      Code = 8
	INTERNAL             Code = 9
	UNIMPLEMENTED        Code = 10
	RESOURCE_UNAVAILABLE Code = 11
	DO_NOT_USE_          Code = 100
)

var Code_name = map[int32]string{
	0:   "OK",
	1:   "CANCELLED",
	2:   "UNKNOWN",
	3:   "INVALID_ARGUMENT",
	4:   "DEADLINE_EXCEEDED",
	5:   "NOT_FOUND",
	6:   "ALREADY_EXISTS",
	7:   "PERMISSION_DENIED",
	8:   "UNAUTHENTICATED",
	9:   "INTERNAL",
	10:  "UNIMPLEMENTED",
	11:  "RESOURCE_UNAVAILABLE",
	100: "DO_NOT_USE_",
}
var Code_value = map[string]int32{
	"OK":                   0,
	"CANCELLED":            1,
	"UNKNOWN":              2,
	"INVALID_ARGUMENT":     3,
	"DEADLINE_EXCEEDED":    4,
	"NOT_FOUND":            5,
	"ALREADY_EXISTS":       6,
	"PERMISSION_DENIED":    7,
	"UNAUTHENTICATED":      8,
	"INTERNAL":             9,
	"UNIMPLEMENTED":        10,
	"RESOURCE_UNAVAILABLE": 11,
	"DO_NOT_USE_":          100,
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_error_codes_998aea5103b251fc, []int{0}
}

func init() {
	proto.RegisterEnum("pl.error.Code", Code_name, Code_value)
}
func (x Code) String() string {
	s, ok := Code_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("src/common/codes/error_codes.proto", fileDescriptor_error_codes_998aea5103b251fc)
}

var fileDescriptor_error_codes_998aea5103b251fc = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x1c, 0x90, 0xbf, 0x4e, 0x2a, 0x41,
	0x18, 0xc5, 0x77, 0xb8, 0xdc, 0x05, 0x86, 0xcb, 0x65, 0x18, 0x31, 0xb1, 0x9a, 0xc2, 0xd2, 0x02,
	0x0a, 0x5b, 0x9b, 0x61, 0xe7, 0x53, 0x27, 0x0c, 0xdf, 0x92, 0xdd, 0x1d, 0x44, 0x9b, 0x49, 0xf8,
	0xd3, 0x89, 0x4b, 0x16, 0x1f, 0xc0, 0x47, 0xf0, 0x31, 0x7c, 0x14, 0x4b, 0x4a, 0x4a, 0x19, 0x1a,
	0x4b, 0xe2, 0x13, 0x98, 0xa5, 0x3b, 0x39, 0xc9, 0x39, 0xf9, 0xe5, 0x47, 0x2f, 0x37, 0xc5, 0xbc,
	0x3f, 0xcf, 0x57, 0xab, 0xfc, 0xa5, 0x3f, 0xcf, 0x17, 0xcb, 0x4d, 0x7f, 0x59, 0x14, 0x79, 0xe1,
	0x4e, 0xb9, 0xb7, 0x2e, 0xf2, 0xd7, 0x9c, 0xd7, 0xd7, 0xcf, 0xbd, 0x53, 0x7b, 0xf5, 0x43, 0x68,
	0x35, 0xca, 0x17, 0x4b, 0x1e, 0xd2, 0x4a, 0x3c, 0x64, 0x01, 0x6f, 0xd1, 0x46, 0x24, 0x31, 0x02,
	0x63, 0x40, 0x31, 0xc2, 0x9b, 0xb4, 0x66, 0x71, 0x88, 0xf1, 0x03, 0xb2, 0x0a, 0xef, 0x52, 0xa6,
	0x71, 0x22, 0x8d, 0x56, 0x4e, 0x26, 0x77, 0x76, 0x04, 0x98, 0xb1, 0x3f, 0xfc, 0x9c, 0x76, 0x14,
	0x48, 0x65, 0x34, 0x82, 0x83, 0x69, 0x04, 0xa0, 0x40, 0xb1, 0x6a, 0x79, 0x84, 0x71, 0xe6, 0x6e,
	0x63, 0x8b, 0x8a, 0xfd, 0xe5, 0x9c, 0xfe, 0x97, 0x26, 0x01, 0xa9, 0x1e, 0x1d, 0x4c, 0x75, 0x9a,
	0xa5, 0x2c, 0x2c, 0x97, 0x63, 0x48, 0x46, 0x3a, 0x4d, 0x75, 0x8c, 0x4e, 0x01, 0x6a, 0x50, 0xac,
	0xc6, 0xcf, 0x68, 0xdb, 0xa2, 0xb4, 0xd9, 0x3d, 0x60, 0xa6, 0x23, 0x99, 0x81, 0x62, 0x75, 0xfe,
	0x8f, 0xd6, 0x35, 0x66, 0x90, 0xa0, 0x34, 0xac, 0xc1, 0x3b, 0xb4, 0x65, 0x51, 0x8f, 0xc6, 0x06,
	0x4a, 0x08, 0x50, 0x8c, 0xf2, 0x0b, 0xda, 0x4d, 0x20, 0x8d, 0x6d, 0x12, 0x81, 0xb3, 0x28, 0x27,
	0x52, 0x1b, 0x39, 0x30, 0xc0, 0x9a, 0xbc, 0x4d, 0x9b, 0x2a, 0x76, 0x25, 0x8c, 0x4d, 0xc1, 0xb1,
	0xc5, 0xe0, 0x66, 0xbb, 0x17, 0xc1, 0x6e, 0x2f, 0x82, 0xe3, 0x5e, 0x90, 0x37, 0x2f, 0xc8, 0x87,
	0x17, 0xe4, 0xd3, 0x0b, 0xb2, 0xf5, 0x82, 0x7c, 0x79, 0x41, 0xbe, 0xbd, 0x08, 0x8e, 0x5e, 0x90,
	0xf7, 0x83, 0x08, 0xb6, 0x07, 0x11, 0xec, 0x0e, 0x22, 0x78, 0x0a, 0x4b, 0x83, 0xeb, 0xd9, 0x2c,
	0x3c, 0x39, 0xbc, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x38, 0x77, 0xfc, 0x69, 0x01, 0x00,
	0x00,
}
