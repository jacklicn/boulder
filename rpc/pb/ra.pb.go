// Code generated by protoc-gen-go.
// source: rpc/pb/ra.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	rpc/pb/ra.proto
	rpc/pb/va.proto

It has these top-level messages:
	Domain
	Valid
	PerformValidationRequest
	VAChallenge
	KeyAuthorization
	AuthzMeta
	ValidationRecords
	ValidationRecord
*/
package pb

import proto "github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

var fileDescriptor0 = []byte{
	// 50 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2a, 0x48, 0xd6,
	0x2f, 0x48, 0xd2, 0x2f, 0x4a, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x4a,
	0x62, 0x03, 0x33, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x19, 0x4e, 0xa1, 0x1d, 0x00,
	0x00, 0x00,
}