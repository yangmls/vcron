// Code generated by protoc-gen-go.
// source: request.proto
// DO NOT EDIT!

/*
Package vcron is a generated protocol buffer package.

It is generated from these files:
	request.proto
	response.proto

It has these top-level messages:
	Request
	Job
	Response
*/
package vcron

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Request struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Jobs []*Job `protobuf:"bytes,2,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type Job struct {
	Command    string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Expression string `protobuf:"bytes,2,opt,name=expression" json:"expression,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "vcron.Request")
	proto.RegisterType((*Job)(nil), "vcron.Job")
}

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x4b, 0x2e, 0xca, 0xcf,
	0x53, 0x32, 0xe4, 0x62, 0x0f, 0x82, 0x88, 0x0b, 0xf1, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x0a, 0x49, 0x70, 0xb1, 0x64, 0xe5, 0x27, 0x15, 0x4b, 0x30, 0x29,
	0x30, 0x6b, 0x70, 0x1b, 0x71, 0xe9, 0x81, 0x95, 0xeb, 0x79, 0xe5, 0x27, 0x29, 0x69, 0x71, 0x31,
	0x03, 0x29, 0x21, 0x7e, 0x2e, 0xf6, 0xe4, 0xfc, 0xdc, 0xdc, 0xc4, 0xbc, 0x14, 0xa8, 0x0e, 0x21,
	0x2e, 0xae, 0xd4, 0x8a, 0x82, 0xa2, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0xa0, 0x3e, 0xa0, 0x58,
	0x12, 0x1b, 0xd8, 0x32, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x6e, 0x3e, 0x9d, 0x7d,
	0x00, 0x00, 0x00,
}
