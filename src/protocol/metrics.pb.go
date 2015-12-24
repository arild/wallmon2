// Code generated by protoc-gen-go.
// source: metrics.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	metrics.proto

It has these top-level messages:
	Metrics
	MetricsMessage
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Metrics struct {
	Pid              *uint32  `protobuf:"varint,1,req,name=pid" json:"pid,omitempty"`
	Cpu              *float32 `protobuf:"fixed32,2,req,name=cpu" json:"cpu,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Metrics) GetPid() uint32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *Metrics) GetCpu() float32 {
	if m != nil && m.Cpu != nil {
		return *m.Cpu
	}
	return 0
}

type MetricsMessage struct {
	Metricsmessage   []*Metrics `protobuf:"bytes,1,rep,name=metricsmessage" json:"metricsmessage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *MetricsMessage) Reset()                    { *m = MetricsMessage{} }
func (m *MetricsMessage) String() string            { return proto.CompactTextString(m) }
func (*MetricsMessage) ProtoMessage()               {}
func (*MetricsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MetricsMessage) GetMetricsmessage() []*Metrics {
	if m != nil {
		return m.Metricsmessage
	}
	return nil
}

func init() {
	proto.RegisterType((*Metrics)(nil), "protocol.Metrics")
	proto.RegisterType((*MetricsMessage)(nil), "protocol.MetricsMessage")
}

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x29,
	0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39,
	0x4a, 0xca, 0x5c, 0xec, 0xbe, 0x10, 0x29, 0x21, 0x6e, 0x2e, 0xe6, 0x82, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x26, 0x0d, 0x5e, 0x10, 0x27, 0xb9, 0xa0, 0x54, 0x82, 0x09, 0xc8, 0x61, 0x52, 0xb2, 0xe6,
	0xe2, 0x83, 0x2a, 0xf2, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xd2, 0xe4, 0xe2, 0x83, 0x9a,
	0x98, 0x0b, 0x11, 0x01, 0x6a, 0x63, 0xd6, 0xe0, 0x36, 0x12, 0xd4, 0x83, 0x99, 0xac, 0x07, 0xd5,
	0x01, 0x08, 0x00, 0x00, 0xff, 0xff, 0x66, 0x91, 0xf5, 0x89, 0x7b, 0x00, 0x00, 0x00,
}