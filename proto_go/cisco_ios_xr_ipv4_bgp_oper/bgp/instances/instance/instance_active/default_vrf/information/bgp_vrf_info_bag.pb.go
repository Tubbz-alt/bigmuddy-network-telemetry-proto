// Code generated by protoc-gen-go.
// source: bgp_vrf_info_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_information is a generated protocol buffer package.

It is generated from these files:
	bgp_vrf_info_bag.proto

It has these top-level messages:
	BgpVrfInfoBag_KEYS
	BgpVrfInfoBag
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_information

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

// BGP VRF information bag
type BgpVrfInfoBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *BgpVrfInfoBag_KEYS) Reset()                    { *m = BgpVrfInfoBag_KEYS{} }
func (m *BgpVrfInfoBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpVrfInfoBag_KEYS) ProtoMessage()               {}
func (*BgpVrfInfoBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpVrfInfoBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpVrfInfoBag struct {
	// Route Distinguisher
	RouteDistinguisher string `protobuf:"bytes,50,opt,name=route_distinguisher,json=routeDistinguisher" json:"route_distinguisher,omitempty"`
}

func (m *BgpVrfInfoBag) Reset()                    { *m = BgpVrfInfoBag{} }
func (m *BgpVrfInfoBag) String() string            { return proto.CompactTextString(m) }
func (*BgpVrfInfoBag) ProtoMessage()               {}
func (*BgpVrfInfoBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpVrfInfoBag) GetRouteDistinguisher() string {
	if m != nil {
		return m.RouteDistinguisher
	}
	return ""
}

func init() {
	proto.RegisterType((*BgpVrfInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.information.bgp_vrf_info_bag_KEYS")
	proto.RegisterType((*BgpVrfInfoBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.information.bgp_vrf_info_bag")
}

func init() { proto.RegisterFile("bgp_vrf_info_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0xf1, 0x52, 0xa8, 0x68, 0xa1, 0xa8, 0xb4, 0x78, 0x2c, 0xee, 0xd2, 0x49, 0x85, 0x36,
	0x63, 0xb6, 0x24, 0x53, 0x20, 0x43, 0x32, 0x79, 0x3a, 0x24, 0xf9, 0xac, 0x1c, 0xc4, 0x3a, 0x21,
	0xc9, 0x26, 0x3f, 0x3f, 0xd8, 0x10, 0x13, 0xbc, 0x3d, 0xee, 0xdd, 0xf7, 0x1d, 0x27, 0x3e, 0x8d,
	0x0b, 0x30, 0xc4, 0x16, 0xc8, 0xb7, 0x0c, 0x46, 0x3b, 0x15, 0x22, 0x67, 0x96, 0xb5, 0xa5, 0x64,
	0x19, 0x88, 0x13, 0x5c, 0x23, 0x50, 0x18, 0x56, 0x30, 0x6e, 0x72, 0xc0, 0xa8, 0x8c, 0x0b, 0x8a,
	0x7c, 0xca, 0xda, 0x5b, 0x4c, 0x73, 0x9a, 0x03, 0x68, 0x9b, 0x69, 0x40, 0xd5, 0x60, 0xab, 0xfb,
	0x4b, 0x1e, 0xed, 0x6a, 0xb4, 0xc7, 0x4e, 0x67, 0x62, 0x5f, 0xad, 0xc5, 0xc7, 0xf2, 0x28, 0xec,
	0x77, 0xf5, 0x49, 0x7e, 0x8b, 0xd7, 0xd9, 0xe1, 0x75, 0x87, 0x65, 0xf1, 0x55, 0xfc, 0x3c, 0x1f,
	0x5f, 0xee, 0xc3, 0x83, 0xee, 0xb0, 0xda, 0x88, 0xb7, 0x25, 0x2d, 0x7f, 0xc5, 0x7b, 0xe4, 0x3e,
	0x23, 0x34, 0x94, 0x32, 0x79, 0xd7, 0x53, 0x3a, 0x63, 0x2c, 0xff, 0x26, 0x5c, 0x4e, 0xd5, 0xf6,
	0xb1, 0x31, 0x4f, 0xd3, 0x93, 0xff, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x62, 0xd4, 0x8d,
	0xfe, 0x00, 0x00, 0x00,
}