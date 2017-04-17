// Code generated by protoc-gen-go.
// source: nflog/nflogpb/nflog.proto
// DO NOT EDIT!

/*
Package nflogpb is a generated protocol buffer package.

It is generated from these files:
	nflog/nflogpb/nflog.proto

It has these top-level messages:
	Receiver
	Entry
	MeshEntry
*/
package nflogpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Receiver struct {
	// Configured name of the receiver group.
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	// Name of the integration of the receiver.
	Integration string `protobuf:"bytes,2,opt,name=integration" json:"integration,omitempty"`
	// Index of the receiver with respect to the integration.
	// Every integration in a group may have 0..N configurations.
	Idx uint32 `protobuf:"varint,3,opt,name=idx" json:"idx,omitempty"`
}

func (m *Receiver) Reset()                    { *m = Receiver{} }
func (m *Receiver) String() string            { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()               {}
func (*Receiver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Receiver) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Receiver) GetIntegration() string {
	if m != nil {
		return m.Integration
	}
	return ""
}

func (m *Receiver) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

// Entry holds information about a successful notification
// sent to a receiver.
type Entry struct {
	// The key identifying the dispatching group.
	GroupKey []byte `protobuf:"bytes,1,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	// The receiver that was notified.
	Receiver *Receiver `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	// Hash over the state of the group at notification time.
	// Deprecated in favor of FiringAlerts field, but kept for compatibility.
	GroupHash []byte `protobuf:"bytes,3,opt,name=group_hash,json=groupHash,proto3" json:"group_hash,omitempty"`
	// Whether the notification was about a resolved alert.
	// Deprecated in favor of ResolvedAlerts field, but kept for compatibility.
	Resolved bool `protobuf:"varint,4,opt,name=resolved" json:"resolved,omitempty"`
	// Timestamp of the succeeding notification.
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
	// FiringAlerts list of hashes of firing alerts at last notification.
	FiringAlerts []uint64 `protobuf:"varint,6,rep,packed,name=firing_alerts,json=firingAlerts" json:"firing_alerts,omitempty"`
	// ResolvedAlerts list hashes of resolved alerts at last notification.
	ResolvedAlerts []uint64 `protobuf:"varint,7,rep,packed,name=resolved_alerts,json=resolvedAlerts" json:"resolved_alerts,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entry) GetGroupKey() []byte {
	if m != nil {
		return m.GroupKey
	}
	return nil
}

func (m *Entry) GetReceiver() *Receiver {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *Entry) GetGroupHash() []byte {
	if m != nil {
		return m.GroupHash
	}
	return nil
}

func (m *Entry) GetResolved() bool {
	if m != nil {
		return m.Resolved
	}
	return false
}

func (m *Entry) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetFiringAlerts() []uint64 {
	if m != nil {
		return m.FiringAlerts
	}
	return nil
}

func (m *Entry) GetResolvedAlerts() []uint64 {
	if m != nil {
		return m.ResolvedAlerts
	}
	return nil
}

// MeshEntry is a wrapper message to communicate a notify log
// entry through a mesh network.
type MeshEntry struct {
	// The original raw notify log entry.
	Entry *Entry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
	// A timestamp indicating when the mesh peer should evict
	// the log entry from its state.
	ExpiresAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
}

func (m *MeshEntry) Reset()                    { *m = MeshEntry{} }
func (m *MeshEntry) String() string            { return proto.CompactTextString(m) }
func (*MeshEntry) ProtoMessage()               {}
func (*MeshEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MeshEntry) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *MeshEntry) GetExpiresAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Receiver)(nil), "nflogpb.Receiver")
	proto.RegisterType((*Entry)(nil), "nflogpb.Entry")
	proto.RegisterType((*MeshEntry)(nil), "nflogpb.MeshEntry")
}

func init() { proto.RegisterFile("nflog/nflogpb/nflog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4f, 0xe3, 0x30,
	0x10, 0xc5, 0x95, 0xfe, 0x4d, 0xa6, 0x7f, 0x76, 0xd7, 0xa7, 0x6c, 0x57, 0x2b, 0xa2, 0x82, 0x44,
	0x2e, 0xa4, 0x52, 0xb9, 0xc0, 0xb1, 0x07, 0x24, 0x24, 0x04, 0x07, 0x8b, 0x2b, 0x8a, 0x5c, 0x3a,
	0x4d, 0x2c, 0x92, 0x38, 0x72, 0xdc, 0xaa, 0xfd, 0x22, 0x7c, 0x5e, 0xd4, 0x71, 0x92, 0x72, 0xe2,
	0x92, 0xd8, 0x6f, 0x7e, 0x7a, 0xf3, 0xfc, 0xe0, 0x6f, 0xb1, 0xcd, 0x54, 0xb2, 0xa0, 0x6f, 0xb9,
	0xb6, 0xff, 0xa8, 0xd4, 0xca, 0x28, 0x36, 0xac, 0xc5, 0xd9, 0x45, 0xa2, 0x54, 0x92, 0xe1, 0x82,
	0xe4, 0xf5, 0x6e, 0xbb, 0x30, 0x32, 0xc7, 0xca, 0x88, 0xbc, 0xb4, 0xe4, 0xfc, 0x0d, 0x5c, 0x8e,
	0xef, 0x28, 0xf7, 0xa8, 0xd9, 0x7f, 0x80, 0x44, 0xab, 0x5d, 0x19, 0x17, 0x22, 0x47, 0xdf, 0x09,
	0x9c, 0xd0, 0xe3, 0x1e, 0x29, 0x2f, 0x22, 0x47, 0x16, 0xc0, 0x48, 0x16, 0x06, 0x13, 0x2d, 0x8c,
	0x54, 0x85, 0xdf, 0xa1, 0xf9, 0x77, 0x89, 0xfd, 0x86, 0xae, 0xdc, 0x1c, 0xfc, 0x6e, 0xe0, 0x84,
	0x13, 0x7e, 0x3a, 0xce, 0x3f, 0x3b, 0xd0, 0x7f, 0x28, 0x8c, 0x3e, 0xb2, 0x7f, 0x60, 0xad, 0xe2,
	0x0f, 0x3c, 0x92, 0xf7, 0x98, 0xbb, 0x24, 0x3c, 0xe1, 0x91, 0xdd, 0x80, 0xab, 0xeb, 0x14, 0xe4,
	0x3b, 0x5a, 0xfe, 0x89, 0xea, 0x27, 0x44, 0x4d, 0x3c, 0xde, 0x22, 0xe7, 0xa0, 0xa9, 0xa8, 0x52,
	0x5a, 0x37, 0xae, 0x83, 0x3e, 0x8a, 0x2a, 0x65, 0xb3, 0x93, 0x5b, 0xa5, 0xb2, 0x3d, 0x6e, 0xfc,
	0x5e, 0xe0, 0x84, 0x2e, 0x6f, 0xef, 0xec, 0x0e, 0xbc, 0xb6, 0x02, 0xbf, 0x4f, 0xab, 0x66, 0x91,
	0x2d, 0x29, 0x6a, 0x4a, 0x8a, 0x5e, 0x1b, 0x82, 0x9f, 0x61, 0x76, 0x09, 0x93, 0xad, 0xd4, 0xb2,
	0x48, 0x62, 0x91, 0xa1, 0x36, 0x95, 0x3f, 0x08, 0xba, 0x61, 0x8f, 0x8f, 0xad, 0xb8, 0x22, 0x8d,
	0x5d, 0xc3, 0xaf, 0x66, 0x55, 0x83, 0x0d, 0x09, 0x9b, 0x36, 0xb2, 0x05, 0xe7, 0x19, 0x78, 0xcf,
	0x58, 0xa5, 0xb6, 0x9b, 0x2b, 0xe8, 0xe3, 0xe9, 0x40, 0xbd, 0x8c, 0x96, 0xd3, 0xf6, 0xed, 0x34,
	0xe6, 0x76, 0xc8, 0xee, 0x01, 0xf0, 0x50, 0x4a, 0x8d, 0x55, 0x2c, 0x4c, 0x5d, 0xd3, 0x8f, 0xd9,
	0x6b, 0x7a, 0x65, 0xd6, 0x03, 0x1a, 0xdf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x4a, 0xda,
	0xe6, 0x33, 0x02, 0x00, 0x00,
}