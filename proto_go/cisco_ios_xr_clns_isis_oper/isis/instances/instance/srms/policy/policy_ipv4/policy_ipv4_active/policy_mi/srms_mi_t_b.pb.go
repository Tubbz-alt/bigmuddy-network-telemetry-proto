// Code generated by protoc-gen-go.
// source: srms_mi_t_b.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi is a generated protocol buffer package.

It is generated from these files:
	srms_mi_t_b.proto

It has these top-level messages:
	SrmsMiTB_KEYS
	SrmsMiTB
	In6AddrTB
	Addr
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi

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

// SRMS show bag
type SrmsMiTB_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	MiId         string `protobuf:"bytes,2,opt,name=mi_id,json=miId" json:"mi_id,omitempty"`
}

func (m *SrmsMiTB_KEYS) Reset()                    { *m = SrmsMiTB_KEYS{} }
func (m *SrmsMiTB_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SrmsMiTB_KEYS) ProtoMessage()               {}
func (*SrmsMiTB_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SrmsMiTB_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *SrmsMiTB_KEYS) GetMiId() string {
	if m != nil {
		return m.MiId
	}
	return ""
}

type SrmsMiTB struct {
	Src string `protobuf:"bytes,50,opt,name=src" json:"src,omitempty"`
	// Router ID
	Router string `protobuf:"bytes,51,opt,name=router" json:"router,omitempty"`
	// Area (OSPF) or Level (ISIS)
	Area string `protobuf:"bytes,52,opt,name=area" json:"area,omitempty"`
	Addr *Addr  `protobuf:"bytes,53,opt,name=addr" json:"addr,omitempty"`
	// Prefix length
	Prefix uint32 `protobuf:"varint,54,opt,name=prefix" json:"prefix,omitempty"`
	// Starting SID
	SidStart uint32 `protobuf:"varint,55,opt,name=sid_start,json=sidStart" json:"sid_start,omitempty"`
	// SID range
	SidCount uint32 `protobuf:"varint,56,opt,name=sid_count,json=sidCount" json:"sid_count,omitempty"`
	// Last IP Prefix
	LastPrefix string `protobuf:"bytes,57,opt,name=last_prefix,json=lastPrefix" json:"last_prefix,omitempty"`
	// Last SID Index
	LastSidIndex uint32 `protobuf:"varint,58,opt,name=last_sid_index,json=lastSidIndex" json:"last_sid_index,omitempty"`
	// Attached flag
	FlagAttached string `protobuf:"bytes,59,opt,name=flag_attached,json=flagAttached" json:"flag_attached,omitempty"`
}

func (m *SrmsMiTB) Reset()                    { *m = SrmsMiTB{} }
func (m *SrmsMiTB) String() string            { return proto.CompactTextString(m) }
func (*SrmsMiTB) ProtoMessage()               {}
func (*SrmsMiTB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SrmsMiTB) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *SrmsMiTB) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *SrmsMiTB) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *SrmsMiTB) GetAddr() *Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *SrmsMiTB) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *SrmsMiTB) GetSidStart() uint32 {
	if m != nil {
		return m.SidStart
	}
	return 0
}

func (m *SrmsMiTB) GetSidCount() uint32 {
	if m != nil {
		return m.SidCount
	}
	return 0
}

func (m *SrmsMiTB) GetLastPrefix() string {
	if m != nil {
		return m.LastPrefix
	}
	return ""
}

func (m *SrmsMiTB) GetLastSidIndex() uint32 {
	if m != nil {
		return m.LastSidIndex
	}
	return 0
}

func (m *SrmsMiTB) GetFlagAttached() string {
	if m != nil {
		return m.FlagAttached
	}
	return ""
}

type In6AddrTB struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *In6AddrTB) Reset()                    { *m = In6AddrTB{} }
func (m *In6AddrTB) String() string            { return proto.CompactTextString(m) }
func (*In6AddrTB) ProtoMessage()               {}
func (*In6AddrTB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *In6AddrTB) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Addr struct {
	Af string `protobuf:"bytes,1,opt,name=af" json:"af,omitempty"`
	// IPv4
	Ipv4 string `protobuf:"bytes,2,opt,name=ipv4" json:"ipv4,omitempty"`
	// IPv6
	Ipv6 *In6AddrTB `protobuf:"bytes,3,opt,name=ipv6" json:"ipv6,omitempty"`
}

func (m *Addr) Reset()                    { *m = Addr{} }
func (m *Addr) String() string            { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()               {}
func (*Addr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Addr) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *Addr) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Addr) GetIpv6() *In6AddrTB {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

func init() {
	proto.RegisterType((*SrmsMiTB_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.srms_mi_t_b_KEYS")
	proto.RegisterType((*SrmsMiTB)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.srms_mi_t_b")
	proto.RegisterType((*In6AddrTB)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.in6_addr_t_b")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.addr")
}

func init() { proto.RegisterFile("srms_mi_t_b.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x26, 0x69, 0x45, 0x27, 0x69, 0x55, 0x0c, 0x42, 0x96, 0x38, 0x10, 0x85, 0x1e, 0x72,
	0xda, 0x43, 0x5b, 0xc2, 0xdf, 0x09, 0x21, 0x0e, 0x15, 0x08, 0xa1, 0xe4, 0xc4, 0x69, 0xe4, 0xda,
	0x0e, 0x1d, 0x69, 0xd7, 0x5e, 0xd9, 0x4e, 0x14, 0xae, 0x1c, 0x39, 0xf0, 0x3a, 0xbc, 0x1e, 0x1a,
	0xef, 0x2e, 0xdd, 0x27, 0xe8, 0x69, 0xbf, 0x1f, 0xeb, 0xb3, 0x67, 0x3e, 0x2d, 0x3c, 0x8e, 0xa1,
	0x8e, 0x58, 0x13, 0x26, 0xbc, 0x2d, 0x9b, 0xe0, 0x93, 0x17, 0x77, 0x9a, 0xa2, 0xf6, 0x48, 0x3e,
	0xe2, 0x21, 0xa0, 0xae, 0x5c, 0x44, 0x8a, 0x14, 0xd1, 0x37, 0x36, 0x94, 0x8c, 0x4a, 0x72, 0x31,
	0x29, 0xa7, 0xed, 0x3d, 0x2a, 0x39, 0xa6, 0x6c, 0x7c, 0x45, 0xfa, 0x67, 0xf7, 0x41, 0x6a, 0xf6,
	0xd7, 0x43, 0x8c, 0x4a, 0x27, 0xda, 0xdb, 0x5e, 0xaa, 0x69, 0xf1, 0x05, 0xce, 0x07, 0xd7, 0xe3,
	0xe7, 0x4f, 0xdf, 0x37, 0xe2, 0x25, 0x9c, 0xf6, 0xa1, 0xe8, 0x54, 0x6d, 0x65, 0x31, 0x2f, 0x96,
	0x27, 0xeb, 0x59, 0x2f, 0x7e, 0x55, 0xb5, 0x15, 0x4f, 0xe0, 0xa8, 0x26, 0x24, 0x23, 0x47, 0xd9,
	0x9c, 0xd4, 0x74, 0x63, 0x16, 0x7f, 0xc6, 0x30, 0x1d, 0xc4, 0x89, 0x73, 0x18, 0xc7, 0xa0, 0xe5,
	0x65, 0x3e, 0xc2, 0x50, 0x3c, 0x83, 0xe3, 0xe0, 0x77, 0xc9, 0x06, 0x79, 0x95, 0xc5, 0x8e, 0x09,
	0x01, 0x13, 0x15, 0xac, 0x92, 0xd7, 0x6d, 0x1a, 0x63, 0xf1, 0xab, 0x80, 0x89, 0x32, 0x26, 0xc8,
	0x57, 0xf3, 0x62, 0x39, 0xbd, 0x74, 0xe5, 0x43, 0x6d, 0xa5, 0xe4, 0x5b, 0xd7, 0xf9, 0x6e, 0x7e,
	0x70, 0x13, 0xec, 0x96, 0x0e, 0x72, 0x35, 0x2f, 0x96, 0xa7, 0xeb, 0x8e, 0x89, 0xe7, 0x70, 0x12,
	0xc9, 0x60, 0x4c, 0x2a, 0x24, 0xf9, 0x3a, 0x5b, 0x8f, 0x22, 0x99, 0x0d, 0xf3, 0xde, 0xd4, 0x7e,
	0xe7, 0x92, 0x7c, 0xf3, 0xdf, 0xfc, 0xc8, 0x5c, 0xbc, 0x80, 0x69, 0xa5, 0x62, 0xc2, 0x2e, 0xf6,
	0x6d, 0x9e, 0x18, 0x58, 0xfa, 0xd6, 0x46, 0x5f, 0xc0, 0x59, 0x3e, 0xc0, 0x11, 0xe4, 0x8c, 0x3d,
	0xc8, 0x77, 0x39, 0x62, 0xc6, 0xea, 0x86, 0xcc, 0x0d, 0x6b, 0xdc, 0xd2, 0xb6, 0x52, 0x3f, 0x50,
	0xa5, 0xa4, 0xf4, 0x9d, 0x35, 0xf2, 0x7d, 0xdb, 0x12, 0x8b, 0x1f, 0x3a, 0x6d, 0x71, 0x01, 0x33,
	0x72, 0x2b, 0xe4, 0x49, 0x72, 0x21, 0x4f, 0xe1, 0x68, 0xaf, 0xaa, 0x5d, 0x5f, 0x69, 0x4b, 0x16,
	0x7f, 0xbb, 0x45, 0x8b, 0x33, 0x18, 0xa9, 0x6d, 0xe7, 0x8d, 0xd4, 0x96, 0x5b, 0xe1, 0x05, 0xf5,
	0x1d, 0x33, 0x16, 0xbf, 0x8b, 0x2c, 0xae, 0xe4, 0x38, 0xb7, 0xb2, 0x7f, 0xb8, 0x56, 0x86, 0x93,
	0xe4, 0xc7, 0xac, 0x6e, 0x8f, 0xf3, 0xff, 0x72, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xd9,
	0x29, 0x6b, 0x44, 0x03, 0x00, 0x00,
}
