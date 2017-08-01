// Code generated by protoc-gen-go.
// source: isis_sh_nbr.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_neighbors_neighbor is a generated protocol buffer package.

It is generated from these files:
	isis_sh_nbr.proto

It has these top-level messages:
	IsisShNbr_KEYS
	IsisShNbr
	IsisIpv6AddressType
	IsisTopoIdType
	OsiAreaAddressType
	IsisShIntfDet
	IsisIpv4AdjSidBackupInfo
	IsisIpv4AdjSid
	IsisIpv6AdjSidBackupInfo
	IsisIpv6AdjSid
	IsisShAdjIpv4
	IsisShAdjIpv6
	IsisShAdjAf
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_neighbors_neighbor

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

// A neighbor IS
type IsisShNbr_KEYS struct {
	InstanceName  string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	SystemId      string `protobuf:"bytes,2,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *IsisShNbr_KEYS) Reset()                    { *m = IsisShNbr_KEYS{} }
func (m *IsisShNbr_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShNbr_KEYS) ProtoMessage()               {}
func (*IsisShNbr_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShNbr_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShNbr_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *IsisShNbr_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IsisShNbr struct {
	// Neighbor system ID
	NeighborSystemId string `protobuf:"bytes,50,opt,name=neighbor_system_id,json=neighborSystemId" json:"neighbor_system_id,omitempty"`
	// Neighbor SNPA
	NeighborSnpa string `protobuf:"bytes,51,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Local interface
	LocalInterface string `protobuf:"bytes,52,opt,name=local_interface,json=localInterface" json:"local_interface,omitempty"`
	// Neighbor (adjacency) state
	NeighborState string `protobuf:"bytes,53,opt,name=neighbor_state,json=neighborState" json:"neighbor_state,omitempty"`
	// Circuit type
	NeighborCircuitType string `protobuf:"bytes,54,opt,name=neighbor_circuit_type,json=neighborCircuitType" json:"neighbor_circuit_type,omitempty"`
	// TRUE if neighbor is IETF-NSF capable
	NeighborIetfNsfCapableFlag uint32 `protobuf:"varint,55,opt,name=neighbor_ietf_nsf_capable_flag,json=neighborIetfNsfCapableFlag" json:"neighbor_ietf_nsf_capable_flag,omitempty"`
	// Link media type
	NeighborMediaType string `protobuf:"bytes,56,opt,name=neighbor_media_type,json=neighborMediaType" json:"neighbor_media_type,omitempty"`
	// Time (s) until neighbor declared down if no IIH received
	NeighborHoldtime uint32 `protobuf:"varint,57,opt,name=neighbor_holdtime,json=neighborHoldtime" json:"neighbor_holdtime,omitempty"`
	// Active area addresses
	NeighborActiveAreaAddresses []*OsiAreaAddressType `protobuf:"bytes,58,rep,name=neighbor_active_area_addresses,json=neighborActiveAreaAddresses" json:"neighbor_active_area_addresses,omitempty"`
	// TRUE if NeighborUptime is set
	NeighborUptimeValidFlag bool `protobuf:"varint,59,opt,name=neighbor_uptime_valid_flag,json=neighborUptimeValidFlag" json:"neighbor_uptime_valid_flag,omitempty"`
	// How long the neighbor has been up (s)
	NeighborUptime uint32 `protobuf:"varint,60,opt,name=neighbor_uptime,json=neighborUptime" json:"neighbor_uptime,omitempty"`
	// Topologies supported by both neighbor and local system
	TopologiesSupported []*IsisTopoIdType `protobuf:"bytes,61,rep,name=topologies_supported,json=topologiesSupported" json:"topologies_supported,omitempty"`
	// Per address-family data
	NeighborPerAddressFamilyData []*IsisShAdjAf `protobuf:"bytes,62,rep,name=neighbor_per_address_family_data,json=neighborPerAddressFamilyData" json:"neighbor_per_address_family_data,omitempty"`
	// ISIS NSR STANDBY
	NsrStandby bool `protobuf:"varint,63,opt,name=nsr_standby,json=nsrStandby" json:"nsr_standby,omitempty"`
}

func (m *IsisShNbr) Reset()                    { *m = IsisShNbr{} }
func (m *IsisShNbr) String() string            { return proto.CompactTextString(m) }
func (*IsisShNbr) ProtoMessage()               {}
func (*IsisShNbr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShNbr) GetNeighborSystemId() string {
	if m != nil {
		return m.NeighborSystemId
	}
	return ""
}

func (m *IsisShNbr) GetNeighborSnpa() string {
	if m != nil {
		return m.NeighborSnpa
	}
	return ""
}

func (m *IsisShNbr) GetLocalInterface() string {
	if m != nil {
		return m.LocalInterface
	}
	return ""
}

func (m *IsisShNbr) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *IsisShNbr) GetNeighborCircuitType() string {
	if m != nil {
		return m.NeighborCircuitType
	}
	return ""
}

func (m *IsisShNbr) GetNeighborIetfNsfCapableFlag() uint32 {
	if m != nil {
		return m.NeighborIetfNsfCapableFlag
	}
	return 0
}

func (m *IsisShNbr) GetNeighborMediaType() string {
	if m != nil {
		return m.NeighborMediaType
	}
	return ""
}

func (m *IsisShNbr) GetNeighborHoldtime() uint32 {
	if m != nil {
		return m.NeighborHoldtime
	}
	return 0
}

func (m *IsisShNbr) GetNeighborActiveAreaAddresses() []*OsiAreaAddressType {
	if m != nil {
		return m.NeighborActiveAreaAddresses
	}
	return nil
}

func (m *IsisShNbr) GetNeighborUptimeValidFlag() bool {
	if m != nil {
		return m.NeighborUptimeValidFlag
	}
	return false
}

func (m *IsisShNbr) GetNeighborUptime() uint32 {
	if m != nil {
		return m.NeighborUptime
	}
	return 0
}

func (m *IsisShNbr) GetTopologiesSupported() []*IsisTopoIdType {
	if m != nil {
		return m.TopologiesSupported
	}
	return nil
}

func (m *IsisShNbr) GetNeighborPerAddressFamilyData() []*IsisShAdjAf {
	if m != nil {
		return m.NeighborPerAddressFamilyData
	}
	return nil
}

func (m *IsisShNbr) GetNsrStandby() bool {
	if m != nil {
		return m.NsrStandby
	}
	return false
}

type IsisIpv6AddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisIpv6AddressType) Reset()                    { *m = IsisIpv6AddressType{} }
func (m *IsisIpv6AddressType) String() string            { return proto.CompactTextString(m) }
func (*IsisIpv6AddressType) ProtoMessage()               {}
func (*IsisIpv6AddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisIpv6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Identification of an IS-IS topology
type IsisTopoIdType struct {
	// AF name
	AfName string `protobuf:"bytes,1,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	// Sub-AF name
	SafName string `protobuf:"bytes,2,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	// VRF Name
	VrfName string `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Topology Name
	TopologyName string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
}

func (m *IsisTopoIdType) Reset()                    { *m = IsisTopoIdType{} }
func (m *IsisTopoIdType) String() string            { return proto.CompactTextString(m) }
func (*IsisTopoIdType) ProtoMessage()               {}
func (*IsisTopoIdType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisTopoIdType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisTopoIdType) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisTopoIdType) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IsisTopoIdType) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

type OsiAreaAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *OsiAreaAddressType) Reset()                    { *m = OsiAreaAddressType{} }
func (m *OsiAreaAddressType) String() string            { return proto.CompactTextString(m) }
func (*OsiAreaAddressType) ProtoMessage()               {}
func (*OsiAreaAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OsiAreaAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Interface Detail
type IsisShIntfDet struct {
	// Local interface handle
	InterfaceHandle string `protobuf:"bytes,1,opt,name=interface_handle,json=interfaceHandle" json:"interface_handle,omitempty"`
	// Interface name
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *IsisShIntfDet) Reset()                    { *m = IsisShIntfDet{} }
func (m *IsisShIntfDet) String() string            { return proto.CompactTextString(m) }
func (*IsisShIntfDet) ProtoMessage()               {}
func (*IsisShIntfDet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsisShIntfDet) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
}

func (m *IsisShIntfDet) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

// IPv4 Adjacency SID backup information
type IsisIpv4AdjSidBackupInfo struct {
	// Number of labels in the backup path label stack
	BackupLabelStackSize uint32 `protobuf:"varint,1,opt,name=backup_label_stack_size,json=backupLabelStackSize" json:"backup_label_stack_size,omitempty"`
	// Backup path label stack
	BackupLabelStack []uint32 `protobuf:"varint,2,rep,packed,name=backup_label_stack,json=backupLabelStack" json:"backup_label_stack,omitempty"`
	// Neighbor address used as adjacency backup target
	BackupNodeAddress string `protobuf:"bytes,3,opt,name=backup_node_address,json=backupNodeAddress" json:"backup_node_address,omitempty"`
	// Backup path nexthop address
	BackupNexthop string `protobuf:"bytes,4,opt,name=backup_nexthop,json=backupNexthop" json:"backup_nexthop,omitempty"`
	// Backup path interface
	BackupInterface string `protobuf:"bytes,5,opt,name=backup_interface,json=backupInterface" json:"backup_interface,omitempty"`
}

func (m *IsisIpv4AdjSidBackupInfo) Reset()                    { *m = IsisIpv4AdjSidBackupInfo{} }
func (m *IsisIpv4AdjSidBackupInfo) String() string            { return proto.CompactTextString(m) }
func (*IsisIpv4AdjSidBackupInfo) ProtoMessage()               {}
func (*IsisIpv4AdjSidBackupInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IsisIpv4AdjSidBackupInfo) GetBackupLabelStackSize() uint32 {
	if m != nil {
		return m.BackupLabelStackSize
	}
	return 0
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupLabelStack() []uint32 {
	if m != nil {
		return m.BackupLabelStack
	}
	return nil
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupNodeAddress() string {
	if m != nil {
		return m.BackupNodeAddress
	}
	return ""
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupNexthop() string {
	if m != nil {
		return m.BackupNexthop
	}
	return ""
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupInterface() string {
	if m != nil {
		return m.BackupInterface
	}
	return ""
}

// IPv4 Adjacency SID
type IsisIpv4AdjSid struct {
	// Adjacency SID value
	AdjacencySidValue uint32 `protobuf:"varint,1,opt,name=adjacency_sid_value,json=adjacencySidValue" json:"adjacency_sid_value,omitempty"`
	// Adjacency SID Backup Info
	AdjacencySidBackup *IsisIpv4AdjSidBackupInfo `protobuf:"bytes,2,opt,name=adjacency_sid_backup,json=adjacencySidBackup" json:"adjacency_sid_backup,omitempty"`
	// Adjacency SID Backup Info TE
	AdjacencySidBackupTe *IsisIpv4AdjSidBackupInfo `protobuf:"bytes,3,opt,name=adjacency_sid_backup_te,json=adjacencySidBackupTe" json:"adjacency_sid_backup_te,omitempty"`
}

func (m *IsisIpv4AdjSid) Reset()                    { *m = IsisIpv4AdjSid{} }
func (m *IsisIpv4AdjSid) String() string            { return proto.CompactTextString(m) }
func (*IsisIpv4AdjSid) ProtoMessage()               {}
func (*IsisIpv4AdjSid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsisIpv4AdjSid) GetAdjacencySidValue() uint32 {
	if m != nil {
		return m.AdjacencySidValue
	}
	return 0
}

func (m *IsisIpv4AdjSid) GetAdjacencySidBackup() *IsisIpv4AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackup
	}
	return nil
}

func (m *IsisIpv4AdjSid) GetAdjacencySidBackupTe() *IsisIpv4AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackupTe
	}
	return nil
}

// IPv4 Adjacency SID backup information
type IsisIpv6AdjSidBackupInfo struct {
	// Number of labels in the backup path label stack
	BackupLabelStackSize uint32 `protobuf:"varint,1,opt,name=backup_label_stack_size,json=backupLabelStackSize" json:"backup_label_stack_size,omitempty"`
	// Backup path label stack
	BackupLabelStack []uint32 `protobuf:"varint,2,rep,packed,name=backup_label_stack,json=backupLabelStack" json:"backup_label_stack,omitempty"`
	// Neighbor address used as adjacency backup target
	BackupNodeAddress string `protobuf:"bytes,3,opt,name=backup_node_address,json=backupNodeAddress" json:"backup_node_address,omitempty"`
	// Backup path nexthop address
	BackupNexthop string `protobuf:"bytes,4,opt,name=backup_nexthop,json=backupNexthop" json:"backup_nexthop,omitempty"`
	// Backup path interface
	BackupInterface string `protobuf:"bytes,5,opt,name=backup_interface,json=backupInterface" json:"backup_interface,omitempty"`
}

func (m *IsisIpv6AdjSidBackupInfo) Reset()                    { *m = IsisIpv6AdjSidBackupInfo{} }
func (m *IsisIpv6AdjSidBackupInfo) String() string            { return proto.CompactTextString(m) }
func (*IsisIpv6AdjSidBackupInfo) ProtoMessage()               {}
func (*IsisIpv6AdjSidBackupInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IsisIpv6AdjSidBackupInfo) GetBackupLabelStackSize() uint32 {
	if m != nil {
		return m.BackupLabelStackSize
	}
	return 0
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupLabelStack() []uint32 {
	if m != nil {
		return m.BackupLabelStack
	}
	return nil
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupNodeAddress() string {
	if m != nil {
		return m.BackupNodeAddress
	}
	return ""
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupNexthop() string {
	if m != nil {
		return m.BackupNexthop
	}
	return ""
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupInterface() string {
	if m != nil {
		return m.BackupInterface
	}
	return ""
}

// IPv6 Adjacency SID
type IsisIpv6AdjSid struct {
	// Adjacency SID value
	AdjacencySidValue uint32 `protobuf:"varint,1,opt,name=adjacency_sid_value,json=adjacencySidValue" json:"adjacency_sid_value,omitempty"`
	// Adjacency SID Backup Info
	AdjacencySidBackup *IsisIpv6AdjSidBackupInfo `protobuf:"bytes,2,opt,name=adjacency_sid_backup,json=adjacencySidBackup" json:"adjacency_sid_backup,omitempty"`
	// Adjacency SID Backup Info TE
	AdjacencySidBackupTe *IsisIpv6AdjSidBackupInfo `protobuf:"bytes,3,opt,name=adjacency_sid_backup_te,json=adjacencySidBackupTe" json:"adjacency_sid_backup_te,omitempty"`
}

func (m *IsisIpv6AdjSid) Reset()                    { *m = IsisIpv6AdjSid{} }
func (m *IsisIpv6AdjSid) String() string            { return proto.CompactTextString(m) }
func (*IsisIpv6AdjSid) ProtoMessage()               {}
func (*IsisIpv6AdjSid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IsisIpv6AdjSid) GetAdjacencySidValue() uint32 {
	if m != nil {
		return m.AdjacencySidValue
	}
	return 0
}

func (m *IsisIpv6AdjSid) GetAdjacencySidBackup() *IsisIpv6AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackup
	}
	return nil
}

func (m *IsisIpv6AdjSid) GetAdjacencySidBackupTe() *IsisIpv6AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackupTe
	}
	return nil
}

// IPv4 Specific Per-Adjacency Data
type IsisShAdjIpv4 struct {
	// Adjacency next hop
	NextHop string `protobuf:"bytes,1,opt,name=next_hop,json=nextHop" json:"next_hop,omitempty"`
	// Adjacency interface addresses
	InterfaceAddresses [][]byte `protobuf:"bytes,2,rep,name=interface_addresses,json=interfaceAddresses,proto3" json:"interface_addresses,omitempty"`
	// Adjacency SID
	AdjacencySid *IsisIpv4AdjSid `protobuf:"bytes,3,opt,name=adjacency_sid,json=adjacencySid" json:"adjacency_sid,omitempty"`
	// Adjacency SID not eligible for FRR protection
	NonFrrAdjacencySid *IsisIpv4AdjSid `protobuf:"bytes,4,opt,name=non_frr_adjacency_sid,json=nonFrrAdjacencySid" json:"non_frr_adjacency_sid,omitempty"`
	// Underlying interface list for bundle interfaces
	UnderlyingInterfaceList []*IsisShIntfDet `protobuf:"bytes,5,rep,name=underlying_interface_list,json=underlyingInterfaceList" json:"underlying_interface_list,omitempty"`
	// Per bundle member Adjacency sid
	UnderlyingAdjacencySidList []uint32 `protobuf:"varint,6,rep,packed,name=underlying_adjacency_sid_list,json=underlyingAdjacencySidList" json:"underlying_adjacency_sid_list,omitempty"`
}

func (m *IsisShAdjIpv4) Reset()                    { *m = IsisShAdjIpv4{} }
func (m *IsisShAdjIpv4) String() string            { return proto.CompactTextString(m) }
func (*IsisShAdjIpv4) ProtoMessage()               {}
func (*IsisShAdjIpv4) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IsisShAdjIpv4) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *IsisShAdjIpv4) GetInterfaceAddresses() [][]byte {
	if m != nil {
		return m.InterfaceAddresses
	}
	return nil
}

func (m *IsisShAdjIpv4) GetAdjacencySid() *IsisIpv4AdjSid {
	if m != nil {
		return m.AdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv4) GetNonFrrAdjacencySid() *IsisIpv4AdjSid {
	if m != nil {
		return m.NonFrrAdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv4) GetUnderlyingInterfaceList() []*IsisShIntfDet {
	if m != nil {
		return m.UnderlyingInterfaceList
	}
	return nil
}

func (m *IsisShAdjIpv4) GetUnderlyingAdjacencySidList() []uint32 {
	if m != nil {
		return m.UnderlyingAdjacencySidList
	}
	return nil
}

// IPv6 Specific Per-Adjacency Data
type IsisShAdjIpv6 struct {
	// Adjacency next hop
	NextHop string `protobuf:"bytes,1,opt,name=next_hop,json=nextHop" json:"next_hop,omitempty"`
	// Adjacency interface addresses
	InterfaceAddresses []*IsisIpv6AddressType `protobuf:"bytes,2,rep,name=interface_addresses,json=interfaceAddresses" json:"interface_addresses,omitempty"`
	// Adjacency SID
	AdjacencySid *IsisIpv6AdjSid `protobuf:"bytes,3,opt,name=adjacency_sid,json=adjacencySid" json:"adjacency_sid,omitempty"`
	// Adjacency SID not eligible for FRR protection
	NonFrrAdjacencySid *IsisIpv6AdjSid `protobuf:"bytes,4,opt,name=non_frr_adjacency_sid,json=nonFrrAdjacencySid" json:"non_frr_adjacency_sid,omitempty"`
	// Underlying interface list for bundle interfaces
	UnderlyingInterfaceList []*IsisShIntfDet `protobuf:"bytes,5,rep,name=underlying_interface_list,json=underlyingInterfaceList" json:"underlying_interface_list,omitempty"`
	// Per bundle member Adjacency sid
	UnderlyingAdjacencySidList []uint32 `protobuf:"varint,6,rep,packed,name=underlying_adjacency_sid_list,json=underlyingAdjacencySidList" json:"underlying_adjacency_sid_list,omitempty"`
}

func (m *IsisShAdjIpv6) Reset()                    { *m = IsisShAdjIpv6{} }
func (m *IsisShAdjIpv6) String() string            { return proto.CompactTextString(m) }
func (*IsisShAdjIpv6) ProtoMessage()               {}
func (*IsisShAdjIpv6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IsisShAdjIpv6) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *IsisShAdjIpv6) GetInterfaceAddresses() []*IsisIpv6AddressType {
	if m != nil {
		return m.InterfaceAddresses
	}
	return nil
}

func (m *IsisShAdjIpv6) GetAdjacencySid() *IsisIpv6AdjSid {
	if m != nil {
		return m.AdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv6) GetNonFrrAdjacencySid() *IsisIpv6AdjSid {
	if m != nil {
		return m.NonFrrAdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv6) GetUnderlyingInterfaceList() []*IsisShIntfDet {
	if m != nil {
		return m.UnderlyingInterfaceList
	}
	return nil
}

func (m *IsisShAdjIpv6) GetUnderlyingAdjacencySidList() []uint32 {
	if m != nil {
		return m.UnderlyingAdjacencySidList
	}
	return nil
}

type IsisShAdjAf struct {
	AfName string `protobuf:"bytes,1,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	// IPV4 neighbor info
	Ipv4 *IsisShAdjIpv4 `protobuf:"bytes,2,opt,name=ipv4" json:"ipv4,omitempty"`
	// IPV6 neighbor info
	Ipv6 *IsisShAdjIpv6 `protobuf:"bytes,3,opt,name=ipv6" json:"ipv6,omitempty"`
}

func (m *IsisShAdjAf) Reset()                    { *m = IsisShAdjAf{} }
func (m *IsisShAdjAf) String() string            { return proto.CompactTextString(m) }
func (*IsisShAdjAf) ProtoMessage()               {}
func (*IsisShAdjAf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IsisShAdjAf) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShAdjAf) GetIpv4() *IsisShAdjIpv4 {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *IsisShAdjAf) GetIpv6() *IsisShAdjIpv6 {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShNbr_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_nbr_KEYS")
	proto.RegisterType((*IsisShNbr)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_nbr")
	proto.RegisterType((*IsisIpv6AddressType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv6_address_type")
	proto.RegisterType((*IsisTopoIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_topo_id_type")
	proto.RegisterType((*OsiAreaAddressType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.osi_area_address_type")
	proto.RegisterType((*IsisShIntfDet)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_intf_det")
	proto.RegisterType((*IsisIpv4AdjSidBackupInfo)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv4_adj_sid_backup_info")
	proto.RegisterType((*IsisIpv4AdjSid)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv4_adj_sid")
	proto.RegisterType((*IsisIpv6AdjSidBackupInfo)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv6_adj_sid_backup_info")
	proto.RegisterType((*IsisIpv6AdjSid)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv6_adj_sid")
	proto.RegisterType((*IsisShAdjIpv4)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_adj_ipv4")
	proto.RegisterType((*IsisShAdjIpv6)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_adj_ipv6")
	proto.RegisterType((*IsisShAdjAf)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_adj_af")
}

func init() { proto.RegisterFile("isis_sh_nbr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x4b, 0x6f, 0x1c, 0x35,
	0x1c, 0xd7, 0x6c, 0x9e, 0x75, 0xb2, 0x79, 0x38, 0x09, 0x99, 0xb6, 0x14, 0x56, 0x1b, 0xa1, 0x06,
	0x01, 0x8b, 0x94, 0xb6, 0xc3, 0xa3, 0x3c, 0x94, 0x16, 0xa2, 0x44, 0x94, 0x08, 0xcd, 0x96, 0xa8,
	0x39, 0x80, 0xe5, 0x1d, 0x7b, 0xb2, 0x6e, 0x67, 0xed, 0xd1, 0xd8, 0xbb, 0xea, 0x16, 0x89, 0x1b,
	0x70, 0x41, 0xe2, 0x82, 0xc4, 0x05, 0x71, 0xe0, 0xc4, 0x37, 0xe1, 0x73, 0xf0, 0x3d, 0x38, 0x20,
	0xdb, 0xe3, 0x99, 0xd9, 0x3c, 0x90, 0x90, 0x76, 0x21, 0x12, 0xbd, 0xcd, 0xfc, 0x9f, 0x3f, 0xff,
	0xfd, 0xf3, 0x6f, 0xbc, 0x0b, 0x56, 0x99, 0x64, 0x12, 0xc9, 0x2e, 0xe2, 0x9d, 0xac, 0x95, 0x66,
	0x42, 0x09, 0xb8, 0x17, 0x31, 0x19, 0x09, 0xc4, 0x84, 0x44, 0x4f, 0x33, 0x14, 0x25, 0x5c, 0x22,
	0x13, 0x24, 0x52, 0x9a, 0xb5, 0xf4, 0x53, 0x8b, 0x71, 0xa9, 0x30, 0x8f, 0x68, 0xf9, 0xd4, 0xe2,
	0x94, 0x9d, 0x74, 0x3b, 0x22, 0x93, 0xc5, 0x53, 0xf3, 0x2b, 0xb0, 0x52, 0x29, 0x8e, 0x3e, 0xf9,
	0xf8, 0xb8, 0x0d, 0xb7, 0x40, 0xdd, 0xa5, 0x20, 0x8e, 0x7b, 0xd4, 0xf7, 0x1a, 0xde, 0xf6, 0x95,
	0x70, 0xd1, 0x19, 0x0f, 0x71, 0x8f, 0xc2, 0xeb, 0xe0, 0x8a, 0x1c, 0x4a, 0x45, 0x7b, 0x88, 0x11,
	0xbf, 0x66, 0x02, 0xe6, 0xad, 0xe1, 0x80, 0xc0, 0x57, 0xc0, 0x12, 0xe3, 0x8a, 0x66, 0x31, 0x76,
	0x25, 0xa6, 0x4c, 0x44, 0xbd, 0xb0, 0xea, 0x1a, 0xcd, 0x3f, 0xe6, 0xc0, 0x42, 0xa5, 0x3b, 0x7c,
	0x1d, 0x40, 0x07, 0x0c, 0x95, 0xc5, 0x77, 0x4c, 0xea, 0x8a, 0xf3, 0xb4, 0x5d, 0x93, 0x2d, 0x50,
	0x2f, 0xa3, 0x79, 0x8a, 0xfd, 0x5b, 0x16, 0x66, 0x11, 0xc8, 0x53, 0x0c, 0x6f, 0x82, 0xe5, 0x44,
	0x44, 0x38, 0x41, 0x45, 0x67, 0xff, 0xb6, 0x09, 0x5b, 0x32, 0xe6, 0x03, 0x67, 0xd5, 0x90, 0xcb,
	0x6a, 0x0a, 0x2b, 0xea, 0xdf, 0xb1, 0x90, 0x8b, 0x72, 0xda, 0x08, 0x77, 0xc0, 0x46, 0x11, 0x16,
	0xb1, 0x2c, 0xea, 0x33, 0x85, 0xd4, 0x30, 0xa5, 0x7e, 0x60, 0xa2, 0xd7, 0x9c, 0xf3, 0xbe, 0xf5,
	0x3d, 0x1c, 0xa6, 0x14, 0xde, 0x03, 0x2f, 0x15, 0x39, 0x8c, 0xaa, 0x18, 0x71, 0x19, 0xa3, 0x08,
	0xa7, 0xb8, 0x93, 0x50, 0x14, 0x27, 0xf8, 0xc4, 0x7f, 0xab, 0xe1, 0x6d, 0xd7, 0xc3, 0x6b, 0x2e,
	0xea, 0x80, 0xaa, 0xf8, 0x50, 0xc6, 0xf7, 0x6d, 0xc8, 0x5e, 0x82, 0x4f, 0x60, 0x0b, 0x14, 0xa5,
	0x51, 0x8f, 0x12, 0x86, 0x6d, 0xd7, 0xb7, 0x4d, 0xd7, 0x55, 0xe7, 0xfa, 0x54, 0x7b, 0x4c, 0xcf,
	0xd7, 0x40, 0x61, 0x44, 0x5d, 0x91, 0x10, 0xc5, 0x7a, 0xd4, 0x7f, 0xc7, 0xb4, 0x29, 0x26, 0xb9,
	0x9f, 0xdb, 0xe1, 0xaf, 0x5e, 0x05, 0x21, 0x8e, 0x14, 0x1b, 0x50, 0x84, 0x33, 0x8a, 0x11, 0x26,
	0x24, 0xa3, 0x52, 0x52, 0xe9, 0xbf, 0xdb, 0x98, 0xda, 0x5e, 0xd8, 0xf9, 0xa2, 0x35, 0x1e, 0xda,
	0xb5, 0x84, 0x64, 0x23, 0x1d, 0xcc, 0x6a, 0xc2, 0xeb, 0xce, 0xbf, 0x6b, 0x30, 0xec, 0x66, 0x14,
	0xef, 0x3a, 0x04, 0xf0, 0x2e, 0x28, 0xe6, 0x83, 0xfa, 0xa9, 0xc6, 0x8d, 0x06, 0x38, 0x61, 0xc4,
	0x4e, 0xf0, 0x6e, 0xc3, 0xdb, 0x9e, 0x0f, 0x37, 0x5d, 0xc4, 0xe7, 0x26, 0xe0, 0x48, 0xfb, 0xcd,
	0xf8, 0x6e, 0x82, 0xe5, 0x53, 0xc9, 0xfe, 0x7b, 0x66, 0x18, 0x4b, 0xa3, 0x19, 0xf0, 0x7b, 0x0f,
	0xac, 0x2b, 0x91, 0x8a, 0x44, 0x9c, 0x30, 0x2a, 0x91, 0xec, 0xa7, 0xa9, 0xc8, 0x14, 0x25, 0xfe,
	0xfb, 0x66, 0x00, 0xc7, 0xe3, 0x1a, 0x80, 0x49, 0xd2, 0x8d, 0x10, 0x23, 0x76, 0xf1, 0x6b, 0x65,
	0xdb, 0xb6, 0xeb, 0x0a, 0x7f, 0xf1, 0x40, 0xa3, 0x00, 0x9e, 0xd2, 0xac, 0x98, 0x57, 0x8c, 0x7b,
	0x2c, 0x19, 0x22, 0x82, 0x15, 0xf6, 0x3f, 0x30, 0xd0, 0x8e, 0xc6, 0x0a, 0x4d, 0x76, 0x11, 0x26,
	0x8f, 0x11, 0x8e, 0xc3, 0x17, 0x9d, 0xe3, 0x33, 0x9a, 0xe5, 0xbb, 0xb1, 0x67, 0x9a, 0x7f, 0x84,
	0x15, 0x86, 0x2f, 0x83, 0x05, 0x2e, 0xcd, 0x89, 0xe1, 0xa4, 0x33, 0xf4, 0x3f, 0x34, 0xdb, 0x00,
	0xb8, 0xd4, 0xc7, 0x45, 0x5b, 0x9a, 0x2d, 0xf0, 0x82, 0x29, 0xc8, 0xd2, 0x41, 0x30, 0xb2, 0xdb,
	0x70, 0x1d, 0xcc, 0x0c, 0x70, 0xd2, 0x77, 0xf2, 0x62, 0x5f, 0x9a, 0xdf, 0x79, 0xb9, 0xdc, 0x55,
	0x87, 0x03, 0x37, 0xc1, 0x1c, 0x8e, 0xab, 0x62, 0x34, 0x8b, 0x63, 0x23, 0x43, 0x57, 0xc1, 0xbc,
	0x74, 0x1e, 0xab, 0x42, 0x73, 0xb2, 0x74, 0x0d, 0xb2, 0xb8, 0x2a, 0x3f, 0x73, 0x83, 0xcc, 0xba,
	0xb6, 0x40, 0x3d, 0x9f, 0xf6, 0xd0, 0xfa, 0xa7, 0xad, 0x74, 0x38, 0xa3, 0x51, 0xa7, 0x37, 0xc0,
	0xc6, 0xb9, 0x34, 0xbd, 0x00, 0x38, 0x29, 0x95, 0x94, 0x71, 0x15, 0x23, 0x42, 0x15, 0x7c, 0x15,
	0xac, 0x94, 0x3a, 0xd8, 0xc5, 0x9c, 0x24, 0x2e, 0x69, 0xb9, 0xb0, 0xef, 0x1b, 0xf3, 0x39, 0x92,
	0x59, 0x3b, 0x4f, 0x32, 0xbf, 0xad, 0x81, 0x1b, 0x6e, 0x9e, 0xb7, 0xcd, 0x16, 0x49, 0x46, 0x50,
	0x07, 0x47, 0x4f, 0xfa, 0x29, 0x62, 0x3c, 0x16, 0xf0, 0x0e, 0xd8, 0xcc, 0x5f, 0x13, 0xdc, 0xa1,
	0x89, 0xde, 0x9a, 0xe8, 0x09, 0x92, 0xec, 0x99, 0x6d, 0x5d, 0x0f, 0xd7, 0xad, 0xfb, 0x81, 0xf6,
	0xb6, 0xb5, 0xb3, 0xcd, 0x9e, 0x51, 0xad, 0xbd, 0x67, 0xd3, 0xfc, 0x5a, 0x63, 0x4a, 0x2b, 0xc6,
	0xe9, 0x0c, 0x2d, 0x47, 0x79, 0x34, 0x17, 0x84, 0xba, 0xf1, 0xe4, 0x63, 0x5e, 0xb5, 0xae, 0x43,
	0x41, 0x68, 0x4e, 0x18, 0xbd, 0x3a, 0x17, 0x4f, 0x9f, 0xaa, 0xae, 0x48, 0xf3, 0x89, 0xd7, 0xf3,
	0x50, 0x6b, 0xd4, 0xf3, 0x2a, 0x96, 0xe2, 0xe4, 0x7a, 0xc6, 0xce, 0xcb, 0xda, 0x0b, 0xbd, 0x6e,
	0xfe, 0x59, 0xcb, 0x79, 0x52, 0x1d, 0x84, 0xc6, 0x85, 0xc9, 0x63, 0x1c, 0x51, 0x1e, 0x0d, 0xcd,
	0x64, 0xca, 0x8d, 0xaa, 0x87, 0xab, 0x85, 0xab, 0xcd, 0xc8, 0x91, 0x76, 0xc0, 0x9f, 0x3c, 0xb0,
	0x3e, 0x9a, 0x60, 0xfb, 0x98, 0xe1, 0x2f, 0xec, 0xd0, 0xb1, 0x9e, 0xa9, 0x8b, 0xb6, 0x2c, 0x84,
	0x55, 0x60, 0xf7, 0x8c, 0x03, 0xfe, 0xec, 0x81, 0xcd, 0xf3, 0x90, 0x21, 0x65, 0xd9, 0xfc, 0xaf,
	0x81, 0x5b, 0x3f, 0x0b, 0xee, 0xe1, 0x28, 0x0f, 0x83, 0xe7, 0x3c, 0x3c, 0x35, 0x88, 0x4b, 0xcf,
	0xc3, 0xe0, 0x32, 0xf3, 0x30, 0xf8, 0x07, 0x3c, 0xfc, 0x7d, 0xba, 0x94, 0x5d, 0x9d, 0xa5, 0x69,
	0xac, 0x95, 0x5f, 0x6f, 0x2f, 0xd2, 0xfb, 0x6b, 0xe5, 0x76, 0x4e, 0xbf, 0xef, 0x8b, 0x14, 0xbe,
	0x09, 0xd6, 0x4a, 0x99, 0x2d, 0xaf, 0x37, 0x9a, 0x5f, 0x8b, 0x21, 0x2c, 0x5c, 0xe5, 0xb5, 0xe3,
	0x6b, 0x50, 0x1f, 0x59, 0x7e, 0xbe, 0xe8, 0xe3, 0x89, 0x1d, 0xbe, 0x70, 0xb1, 0xba, 0x50, 0x7d,
	0x21, 0xd9, 0xe0, 0x82, 0xa3, 0x38, 0xd3, 0x1f, 0xff, 0x2a, 0x90, 0xe9, 0x49, 0x03, 0x81, 0x5c,
	0xf0, 0xbd, 0x2c, 0xdb, 0xad, 0xc2, 0xf9, 0xd1, 0x03, 0x57, 0xfb, 0x9c, 0xd0, 0x2c, 0x19, 0x32,
	0x7e, 0x52, 0x1e, 0x0f, 0x94, 0x30, 0xa9, 0xfc, 0x19, 0x73, 0x13, 0x79, 0x34, 0xee, 0x9b, 0x88,
	0xfb, 0x9e, 0x86, 0x9b, 0x65, 0xeb, 0xe2, 0x04, 0x3e, 0x60, 0x52, 0xc1, 0x5d, 0x70, 0xa3, 0x82,
	0x6a, 0x94, 0xaf, 0x06, 0xd9, 0xac, 0x11, 0x90, 0x6b, 0x65, 0x50, 0x75, 0x55, 0xba, 0x44, 0xf3,
	0xb7, 0x99, 0x33, 0x4c, 0x0a, 0xfe, 0x8e, 0x49, 0x3f, 0x78, 0x17, 0x53, 0x69, 0x61, 0xe7, 0xcb,
	0x49, 0x1c, 0x8a, 0xca, 0x55, 0xf9, 0xbf, 0xa4, 0x6a, 0x70, 0x59, 0xa8, 0x1a, 0xfc, 0x4f, 0xa9,
	0xfa, 0x4d, 0x0d, 0x2c, 0x8d, 0xde, 0xd2, 0x2f, 0xbe, 0x20, 0x27, 0x60, 0x5a, 0x1f, 0xea, 0xfc,
	0x43, 0xf2, 0x68, 0x12, 0x3f, 0x12, 0x74, 0xfd, 0xd0, 0x74, 0xc9, 0xbb, 0x05, 0x39, 0xf3, 0x26,
	0xd5, 0x2d, 0x30, 0xdd, 0x82, 0xce, 0xac, 0xf9, 0x2f, 0xe4, 0xd6, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x15, 0xc6, 0x80, 0x68, 0x20, 0x11, 0x00, 0x00,
}