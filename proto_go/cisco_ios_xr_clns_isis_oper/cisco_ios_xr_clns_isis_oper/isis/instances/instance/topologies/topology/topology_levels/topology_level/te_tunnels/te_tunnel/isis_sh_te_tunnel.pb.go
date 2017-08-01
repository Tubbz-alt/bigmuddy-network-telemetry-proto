// Code generated by protoc-gen-go.
// source: isis_sh_te_tunnel.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_tunnels_te_tunnel is a generated protocol buffer package.

It is generated from these files:
	isis_sh_te_tunnel.proto

It has these top-level messages:
	IsisShTeTunnel_KEYS
	IsisShTeTunnel
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_tunnels_te_tunnel

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

// MPLS TE tunnel
type IsisShTeTunnel_KEYS struct {
	InstanceName  string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName        string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName       string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	TopologyName  string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
	Level         string `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
	SystemId      string `protobuf:"bytes,6,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	InterfaceName string `protobuf:"bytes,7,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *IsisShTeTunnel_KEYS) Reset()                    { *m = IsisShTeTunnel_KEYS{} }
func (m *IsisShTeTunnel_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeTunnel_KEYS) ProtoMessage()               {}
func (*IsisShTeTunnel_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShTeTunnel_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IsisShTeTunnel struct {
	// Destination system ID
	TeSystemId string `protobuf:"bytes,50,opt,name=te_system_id,json=teSystemId" json:"te_system_id,omitempty"`
	// Tunnel interface
	TeInterface string `protobuf:"bytes,51,opt,name=te_interface,json=teInterface" json:"te_interface,omitempty"`
	// Tunnel bandwidth
	TeBandwidth uint32 `protobuf:"varint,52,opt,name=te_bandwidth,json=teBandwidth" json:"te_bandwidth,omitempty"`
	// Tunnel metric
	TeigpMetric int32 `protobuf:"zigzag32,53,opt,name=teigp_metric,json=teigpMetric" json:"teigp_metric,omitempty"`
	// Tunnel next-hop IP address
	TeNextHopIpAddress string `protobuf:"bytes,54,opt,name=te_next_hop_ip_address,json=teNextHopIpAddress" json:"te_next_hop_ip_address,omitempty"`
	// Tunnel metric mode
	TeModeType string `protobuf:"bytes,55,opt,name=te_mode_type,json=teModeType" json:"te_mode_type,omitempty"`
	// Indicates whether MPLS TE IPv4 forwarding adjacency is enabled
	TeIpv4FaEnabled bool `protobuf:"varint,56,opt,name=te_ipv4_fa_enabled,json=teIpv4FaEnabled" json:"te_ipv4_fa_enabled,omitempty"`
	// Indicates whether MPLS TE IPv6 forwarding adjacency is enabled
	TeIpv6FaEnabled bool `protobuf:"varint,57,opt,name=te_ipv6_fa_enabled,json=teIpv6FaEnabled" json:"te_ipv6_fa_enabled,omitempty"`
	// Indicates whether MPLS TE IPv4 autoroute announce is enabled
	TeIpv4AaEnabled bool `protobuf:"varint,58,opt,name=te_ipv4_aa_enabled,json=teIpv4AaEnabled" json:"te_ipv4_aa_enabled,omitempty"`
	// Indicates whether MPLS TE IPv6 autoroute announce is enabled
	TeIpv6AaEnabled bool `protobuf:"varint,59,opt,name=te_ipv6_aa_enabled,json=teIpv6AaEnabled" json:"te_ipv6_aa_enabled,omitempty"`
	// Tunnel checkpoint object ID
	TeCheckpointObjectId uint32 `protobuf:"varint,60,opt,name=te_checkpoint_object_id,json=teCheckpointObjectId" json:"te_checkpoint_object_id,omitempty"`
	// Indicates whether MPLS TE segment routing is enabled
	TeSegmentRoutingEnabled bool `protobuf:"varint,61,opt,name=te_segment_routing_enabled,json=teSegmentRoutingEnabled" json:"te_segment_routing_enabled,omitempty"`
	// Indicates whether MPLS TE segment routing strict SPF is enabled
	TeSegmentRoutingStrictSpf bool `protobuf:"varint,62,opt,name=te_segment_routing_strict_spf,json=teSegmentRoutingStrictSpf" json:"te_segment_routing_strict_spf,omitempty"`
	// Indicates whether Segment routing labeled traffic exclusion is enabled
	TeSegmentRoutingExclude bool `protobuf:"varint,63,opt,name=te_segment_routing_exclude,json=teSegmentRoutingExclude" json:"te_segment_routing_exclude,omitempty"`
}

func (m *IsisShTeTunnel) Reset()                    { *m = IsisShTeTunnel{} }
func (m *IsisShTeTunnel) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeTunnel) ProtoMessage()               {}
func (*IsisShTeTunnel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShTeTunnel) GetTeSystemId() string {
	if m != nil {
		return m.TeSystemId
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeInterface() string {
	if m != nil {
		return m.TeInterface
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeBandwidth() uint32 {
	if m != nil {
		return m.TeBandwidth
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeigpMetric() int32 {
	if m != nil {
		return m.TeigpMetric
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeNextHopIpAddress() string {
	if m != nil {
		return m.TeNextHopIpAddress
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeModeType() string {
	if m != nil {
		return m.TeModeType
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeIpv4FaEnabled() bool {
	if m != nil {
		return m.TeIpv4FaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeIpv6FaEnabled() bool {
	if m != nil {
		return m.TeIpv6FaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeIpv4AaEnabled() bool {
	if m != nil {
		return m.TeIpv4AaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeIpv6AaEnabled() bool {
	if m != nil {
		return m.TeIpv6AaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeCheckpointObjectId() uint32 {
	if m != nil {
		return m.TeCheckpointObjectId
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingEnabled() bool {
	if m != nil {
		return m.TeSegmentRoutingEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingStrictSpf() bool {
	if m != nil {
		return m.TeSegmentRoutingStrictSpf
	}
	return false
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingExclude() bool {
	if m != nil {
		return m.TeSegmentRoutingExclude
	}
	return false
}

func init() {
	proto.RegisterType((*IsisShTeTunnel_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_tunnels.te_tunnel.isis_sh_te_tunnel_KEYS")
	proto.RegisterType((*IsisShTeTunnel)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_tunnels.te_tunnel.isis_sh_te_tunnel")
}

func init() { proto.RegisterFile("isis_sh_te_tunnel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x5f, 0x6b, 0x13, 0x41,
	0x14, 0xc5, 0x59, 0xb5, 0x49, 0x3a, 0x36, 0x4a, 0x97, 0xd2, 0x6c, 0x15, 0x21, 0xb6, 0x08, 0x01,
	0x21, 0x0f, 0x6d, 0x1a, 0xff, 0xd4, 0x7f, 0x55, 0x2a, 0x06, 0x69, 0x85, 0xc4, 0x17, 0x9f, 0x2e,
	0x93, 0xdd, 0xbb, 0xc9, 0xe8, 0xee, 0xcc, 0xb0, 0x73, 0x13, 0x13, 0xfc, 0x28, 0x7e, 0x49, 0x3f,
	0x82, 0xec, 0x6c, 0x66, 0x37, 0xb4, 0xe9, 0xdb, 0x9d, 0x73, 0x7e, 0x73, 0xce, 0x4d, 0x06, 0x96,
	0xb5, 0x84, 0x11, 0x06, 0xcc, 0x14, 0x08, 0x81, 0x66, 0x52, 0x62, 0xd2, 0xd5, 0x99, 0x22, 0xe5,
	0xff, 0x09, 0x85, 0x09, 0x15, 0x08, 0x65, 0x60, 0x91, 0x41, 0x98, 0x48, 0x03, 0x16, 0x55, 0x1a,
	0xb3, 0x6e, 0x3e, 0x75, 0x85, 0x34, 0xc4, 0x65, 0x88, 0xd5, 0xd4, 0x25, 0xa5, 0x55, 0xa2, 0x26,
	0x02, 0x8d, 0x1b, 0x97, 0xe5, 0x00, 0x09, 0xce, 0x31, 0x31, 0xd7, 0xce, 0xdd, 0xb2, 0xd7, 0x54,
	0xe3, 0xe1, 0x3f, 0x8f, 0xed, 0xdf, 0x58, 0x0c, 0xbe, 0x5e, 0xfc, 0x18, 0xf9, 0x47, 0xac, 0xe9,
	0xea, 0x40, 0xf2, 0x14, 0x03, 0xaf, 0xed, 0x75, 0xb6, 0x87, 0x3b, 0x4e, 0xbc, 0xe2, 0x29, 0xfa,
	0x2d, 0x56, 0xe7, 0x71, 0x61, 0xdf, 0xb1, 0x76, 0x8d, 0xc7, 0xd6, 0x38, 0x60, 0x0d, 0xe3, 0x9c,
	0xbb, 0xd6, 0xa9, 0x9b, 0x95, 0x75, 0xc4, 0x9a, 0xe5, 0x6a, 0xd6, 0xbf, 0x57, 0x04, 0x3b, 0xd1,
	0x42, 0x7b, 0x6c, 0xcb, 0xae, 0x1d, 0x6c, 0x59, 0xb3, 0x38, 0xf8, 0x8f, 0xd9, 0xb6, 0x59, 0x1a,
	0xc2, 0x14, 0x44, 0x14, 0xd4, 0xac, 0xd3, 0x28, 0x84, 0x41, 0xe4, 0x3f, 0x63, 0x0f, 0x84, 0x24,
	0xcc, 0x62, 0xee, 0x36, 0xae, 0x5b, 0xa2, 0x59, 0xaa, 0x79, 0xf2, 0xe1, 0xdf, 0x2d, 0xb6, 0x7b,
	0xe3, 0x27, 0xfb, 0x6d, 0xb6, 0x43, 0x08, 0x55, 0xf8, 0xb1, 0xbd, 0xca, 0x08, 0x47, 0x2e, 0xfe,
	0xa9, 0x25, 0xca, 0xac, 0xe0, 0xc4, 0x12, 0xf7, 0x09, 0x07, 0x4e, 0x5a, 0x21, 0x63, 0x2e, 0xa3,
	0xdf, 0x22, 0xa2, 0x69, 0xd0, 0x6b, 0x7b, 0x9d, 0x66, 0x8e, 0x7c, 0x74, 0x52, 0x81, 0x88, 0x89,
	0x86, 0x14, 0x29, 0x13, 0x61, 0x70, 0xda, 0xf6, 0x3a, 0xbb, 0x39, 0x22, 0x26, 0xfa, 0xd2, 0x4a,
	0xfe, 0x31, 0xdb, 0x27, 0x04, 0x89, 0x0b, 0x82, 0xa9, 0xd2, 0x20, 0x34, 0xf0, 0x28, 0xca, 0xd0,
	0x98, 0xa0, 0x6f, 0x2b, 0x7d, 0xc2, 0x2b, 0x5c, 0xd0, 0x17, 0xa5, 0x07, 0xfa, 0xbc, 0x70, 0x56,
	0xeb, 0xa7, 0x2a, 0x42, 0xa0, 0xa5, 0xc6, 0xe0, 0x85, 0x5b, 0xff, 0x52, 0x45, 0xf8, 0x7d, 0xa9,
	0xd1, 0x7f, 0xce, 0xfc, 0x7c, 0x7d, 0x3d, 0xef, 0x41, 0xcc, 0x01, 0x25, 0x1f, 0x27, 0x18, 0x05,
	0x2f, 0xdb, 0x5e, 0xa7, 0x31, 0x7c, 0x48, 0x38, 0xd0, 0xf3, 0xde, 0x67, 0x7e, 0x51, 0xc8, 0x15,
	0xdc, 0x5f, 0x87, 0x5f, 0xad, 0xc1, 0xfd, 0x0d, 0x70, 0x0f, 0x78, 0x05, 0xbf, 0x5e, 0x4f, 0x3e,
	0xdf, 0x94, 0xbc, 0x06, 0x9f, 0xad, 0x27, 0x57, 0xf0, 0x29, 0x6b, 0x11, 0x42, 0x38, 0xc5, 0xf0,
	0x97, 0x56, 0x42, 0x12, 0xa8, 0xf1, 0x4f, 0x0c, 0x29, 0x7f, 0x9f, 0x37, 0xf6, 0xaf, 0xdd, 0x23,
	0xfc, 0x54, 0xba, 0xdf, 0xac, 0x39, 0x88, 0xfc, 0x33, 0xf6, 0x28, 0x7f, 0x4b, 0x9c, 0xa4, 0x28,
	0x09, 0x32, 0x35, 0x23, 0x21, 0x27, 0x65, 0xd7, 0x5b, 0xdb, 0xd5, 0x22, 0x1c, 0x15, 0xc0, 0xb0,
	0xf0, 0x5d, 0xe7, 0x07, 0xf6, 0x64, 0xc3, 0x65, 0x93, 0xbf, 0x0c, 0x81, 0xd1, 0x71, 0xf0, 0xce,
	0xde, 0x3f, 0xb8, 0x7e, 0x7f, 0x64, 0x89, 0x91, 0x8e, 0x6f, 0xab, 0x5f, 0x84, 0xc9, 0x2c, 0xc2,
	0xe0, 0xfd, 0x2d, 0xf5, 0x85, 0x3d, 0xae, 0xd9, 0x8f, 0xc2, 0xc9, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x49, 0x07, 0x1f, 0xb0, 0x2f, 0x04, 0x00, 0x00,
}