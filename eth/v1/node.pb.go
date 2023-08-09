// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: eth/v1/node.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	github_com_prysmaticlabs_go_bitfield "github.com/prysmaticlabs/go-bitfield"
	_ "github.com/prysmaticlabs/prysm/v4/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PeerDirection int32

const (
	PeerDirection_INBOUND  PeerDirection = 0
	PeerDirection_OUTBOUND PeerDirection = 1
)

// Enum value maps for PeerDirection.
var (
	PeerDirection_name = map[int32]string{
		0: "INBOUND",
		1: "OUTBOUND",
	}
	PeerDirection_value = map[string]int32{
		"INBOUND":  0,
		"OUTBOUND": 1,
	}
)

func (x PeerDirection) Enum() *PeerDirection {
	p := new(PeerDirection)
	*p = x
	return p
}

func (x PeerDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_eth_v1_node_proto_enumTypes[0].Descriptor()
}

func (PeerDirection) Type() protoreflect.EnumType {
	return &file_eth_v1_node_proto_enumTypes[0]
}

func (x PeerDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerDirection.Descriptor instead.
func (PeerDirection) EnumDescriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{0}
}

type ConnectionState int32

const (
	ConnectionState_DISCONNECTED  ConnectionState = 0
	ConnectionState_CONNECTING    ConnectionState = 1
	ConnectionState_CONNECTED     ConnectionState = 2
	ConnectionState_DISCONNECTING ConnectionState = 3
)

// Enum value maps for ConnectionState.
var (
	ConnectionState_name = map[int32]string{
		0: "DISCONNECTED",
		1: "CONNECTING",
		2: "CONNECTED",
		3: "DISCONNECTING",
	}
	ConnectionState_value = map[string]int32{
		"DISCONNECTED":  0,
		"CONNECTING":    1,
		"CONNECTED":     2,
		"DISCONNECTING": 3,
	}
)

func (x ConnectionState) Enum() *ConnectionState {
	p := new(ConnectionState)
	*p = x
	return p
}

func (x ConnectionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionState) Descriptor() protoreflect.EnumDescriptor {
	return file_eth_v1_node_proto_enumTypes[1].Descriptor()
}

func (ConnectionState) Type() protoreflect.EnumType {
	return &file_eth_v1_node_proto_enumTypes[1]
}

func (x ConnectionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionState.Descriptor instead.
func (ConnectionState) EnumDescriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{1}
}

type IdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Identity `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IdentityResponse) Reset() {
	*x = IdentityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityResponse) ProtoMessage() {}

func (x *IdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityResponse.ProtoReflect.Descriptor instead.
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityResponse) GetData() *Identity {
	if x != nil {
		return x.Data
	}
	return nil
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId             string    `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Enr                string    `protobuf:"bytes,2,opt,name=enr,proto3" json:"enr,omitempty"`
	P2PAddresses       []string  `protobuf:"bytes,3,rep,name=p2p_addresses,json=p2pAddresses,proto3" json:"p2p_addresses,omitempty"`
	DiscoveryAddresses []string  `protobuf:"bytes,4,rep,name=discovery_addresses,json=discoveryAddresses,proto3" json:"discovery_addresses,omitempty"`
	Metadata           *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *Identity) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Identity) GetEnr() string {
	if x != nil {
		return x.Enr
	}
	return ""
}

func (x *Identity) GetP2PAddresses() []string {
	if x != nil {
		return x.P2PAddresses
	}
	return nil
}

func (x *Identity) GetDiscoveryAddresses() []string {
	if x != nil {
		return x.DiscoveryAddresses
	}
	return nil
}

func (x *Identity) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNumber uint64                                           `protobuf:"varint,1,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
	Attnets   github_com_prysmaticlabs_go_bitfield.Bitvector64 `protobuf:"bytes,2,opt,name=attnets,proto3" json:"attnets,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitvector64" ssz-size:"8"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetSeqNumber() uint64 {
	if x != nil {
		return x.SeqNumber
	}
	return 0
}

func (x *Metadata) GetAttnets() github_com_prysmaticlabs_go_bitfield.Bitvector64 {
	if x != nil {
		return x.Attnets
	}
	return github_com_prysmaticlabs_go_bitfield.Bitvector64(nil)
}

type PeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (x *PeerRequest) Reset() {
	*x = PeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerRequest) ProtoMessage() {}

func (x *PeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerRequest.ProtoReflect.Descriptor instead.
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *PeerRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

type PeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State     []ConnectionState `protobuf:"varint,1,rep,packed,name=state,proto3,enum=ethereum.eth.v1.ConnectionState" json:"state,omitempty"`
	Direction []PeerDirection   `protobuf:"varint,2,rep,packed,name=direction,proto3,enum=ethereum.eth.v1.PeerDirection" json:"direction,omitempty"`
}

func (x *PeersRequest) Reset() {
	*x = PeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersRequest) ProtoMessage() {}

func (x *PeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersRequest.ProtoReflect.Descriptor instead.
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *PeersRequest) GetState() []ConnectionState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *PeersRequest) GetDirection() []PeerDirection {
	if x != nil {
		return x.Direction
	}
	return nil
}

type PeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Peer              `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Meta *PeerResponse_Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *PeerResponse) Reset() {
	*x = PeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerResponse) ProtoMessage() {}

func (x *PeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerResponse.ProtoReflect.Descriptor instead.
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *PeerResponse) GetData() *Peer {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PeerResponse) GetMeta() *PeerResponse_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type PeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Peer `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PeersResponse) Reset() {
	*x = PeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersResponse) ProtoMessage() {}

func (x *PeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersResponse.ProtoReflect.Descriptor instead.
func (*PeersResponse) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{6}
}

func (x *PeersResponse) GetData() []*Peer {
	if x != nil {
		return x.Data
	}
	return nil
}

type PeerCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *PeerCountResponse_PeerCount `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PeerCountResponse) Reset() {
	*x = PeerCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerCountResponse) ProtoMessage() {}

func (x *PeerCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerCountResponse.ProtoReflect.Descriptor instead.
func (*PeerCountResponse) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{7}
}

func (x *PeerCountResponse) GetData() *PeerCountResponse_PeerCount {
	if x != nil {
		return x.Data
	}
	return nil
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId             string          `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Enr                string          `protobuf:"bytes,2,opt,name=enr,proto3" json:"enr,omitempty"`
	LastSeenP2PAddress string          `protobuf:"bytes,3,opt,name=last_seen_p2p_address,json=lastSeenP2pAddress,proto3" json:"last_seen_p2p_address,omitempty"`
	State              ConnectionState `protobuf:"varint,4,opt,name=state,proto3,enum=ethereum.eth.v1.ConnectionState" json:"state,omitempty"`
	Direction          PeerDirection   `protobuf:"varint,5,opt,name=direction,proto3,enum=ethereum.eth.v1.PeerDirection" json:"direction,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{8}
}

func (x *Peer) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Peer) GetEnr() string {
	if x != nil {
		return x.Enr
	}
	return ""
}

func (x *Peer) GetLastSeenP2PAddress() string {
	if x != nil {
		return x.LastSeenP2PAddress
	}
	return ""
}

func (x *Peer) GetState() ConnectionState {
	if x != nil {
		return x.State
	}
	return ConnectionState_DISCONNECTED
}

func (x *Peer) GetDirection() PeerDirection {
	if x != nil {
		return x.Direction
	}
	return PeerDirection_INBOUND
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Version `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{9}
}

func (x *VersionResponse) GetData() *Version {
	if x != nil {
		return x.Data
	}
	return nil
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{10}
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type PeerResponse_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PeerResponse_Meta) Reset() {
	*x = PeerResponse_Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerResponse_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerResponse_Meta) ProtoMessage() {}

func (x *PeerResponse_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerResponse_Meta.ProtoReflect.Descriptor instead.
func (*PeerResponse_Meta) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{5, 0}
}

func (x *PeerResponse_Meta) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PeerCountResponse_PeerCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disconnected  uint64 `protobuf:"varint,1,opt,name=disconnected,proto3" json:"disconnected,omitempty"`
	Connecting    uint64 `protobuf:"varint,2,opt,name=connecting,proto3" json:"connecting,omitempty"`
	Connected     uint64 `protobuf:"varint,3,opt,name=connected,proto3" json:"connected,omitempty"`
	Disconnecting uint64 `protobuf:"varint,4,opt,name=disconnecting,proto3" json:"disconnecting,omitempty"`
}

func (x *PeerCountResponse_PeerCount) Reset() {
	*x = PeerCountResponse_PeerCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1_node_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerCountResponse_PeerCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerCountResponse_PeerCount) ProtoMessage() {}

func (x *PeerCountResponse_PeerCount) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1_node_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerCountResponse_PeerCount.ProtoReflect.Descriptor instead.
func (*PeerCountResponse_PeerCount) Descriptor() ([]byte, []int) {
	return file_eth_v1_node_proto_rawDescGZIP(), []int{7, 0}
}

func (x *PeerCountResponse_PeerCount) GetDisconnected() uint64 {
	if x != nil {
		return x.Disconnected
	}
	return 0
}

func (x *PeerCountResponse_PeerCount) GetConnecting() uint64 {
	if x != nil {
		return x.Connecting
	}
	return 0
}

func (x *PeerCountResponse_PeerCount) GetConnected() uint64 {
	if x != nil {
		return x.Connected
	}
	return 0
}

func (x *PeerCountResponse_PeerCount) GetDisconnecting() uint64 {
	if x != nil {
		return x.Disconnecting
	}
	return 0
}

var File_eth_v1_node_proto protoreflect.FileDescriptor

var file_eth_v1_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc2, 0x01, 0x0a,
	0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x6e, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x32, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x32, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x7e, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x07,
	0x61, 0x74, 0x74, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x39, 0x82,
	0xb5, 0x18, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x62,
	0x69, 0x74, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x36, 0x34, 0x8a, 0xb5, 0x18, 0x01, 0x38, 0x52, 0x07, 0x61, 0x74, 0x74, 0x6e, 0x65, 0x74,
	0x73, 0x22, 0x26, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xeb,
	0x01, 0x0a, 0x11, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x93, 0x01, 0x0a, 0x09, 0x50, 0x65, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xda, 0x01, 0x0a,
	0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x72,
	0x12, 0x31, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x70, 0x32,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x50, 0x32, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x0f, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a,
	0x2a, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x2a, 0x55, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x11, 0x0a, 0x0d, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x42, 0x7c, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x0f, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1_node_proto_rawDescOnce sync.Once
	file_eth_v1_node_proto_rawDescData = file_eth_v1_node_proto_rawDesc
)

func file_eth_v1_node_proto_rawDescGZIP() []byte {
	file_eth_v1_node_proto_rawDescOnce.Do(func() {
		file_eth_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1_node_proto_rawDescData)
	})
	return file_eth_v1_node_proto_rawDescData
}

var file_eth_v1_node_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eth_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_eth_v1_node_proto_goTypes = []interface{}{
	(PeerDirection)(0),                  // 0: ethereum.eth.v1.PeerDirection
	(ConnectionState)(0),                // 1: ethereum.eth.v1.ConnectionState
	(*IdentityResponse)(nil),            // 2: ethereum.eth.v1.IdentityResponse
	(*Identity)(nil),                    // 3: ethereum.eth.v1.Identity
	(*Metadata)(nil),                    // 4: ethereum.eth.v1.Metadata
	(*PeerRequest)(nil),                 // 5: ethereum.eth.v1.PeerRequest
	(*PeersRequest)(nil),                // 6: ethereum.eth.v1.PeersRequest
	(*PeerResponse)(nil),                // 7: ethereum.eth.v1.PeerResponse
	(*PeersResponse)(nil),               // 8: ethereum.eth.v1.PeersResponse
	(*PeerCountResponse)(nil),           // 9: ethereum.eth.v1.PeerCountResponse
	(*Peer)(nil),                        // 10: ethereum.eth.v1.Peer
	(*VersionResponse)(nil),             // 11: ethereum.eth.v1.VersionResponse
	(*Version)(nil),                     // 12: ethereum.eth.v1.Version
	(*PeerResponse_Meta)(nil),           // 13: ethereum.eth.v1.PeerResponse.Meta
	(*PeerCountResponse_PeerCount)(nil), // 14: ethereum.eth.v1.PeerCountResponse.PeerCount
}
var file_eth_v1_node_proto_depIdxs = []int32{
	3,  // 0: ethereum.eth.v1.IdentityResponse.data:type_name -> ethereum.eth.v1.Identity
	4,  // 1: ethereum.eth.v1.Identity.metadata:type_name -> ethereum.eth.v1.Metadata
	1,  // 2: ethereum.eth.v1.PeersRequest.state:type_name -> ethereum.eth.v1.ConnectionState
	0,  // 3: ethereum.eth.v1.PeersRequest.direction:type_name -> ethereum.eth.v1.PeerDirection
	10, // 4: ethereum.eth.v1.PeerResponse.data:type_name -> ethereum.eth.v1.Peer
	13, // 5: ethereum.eth.v1.PeerResponse.meta:type_name -> ethereum.eth.v1.PeerResponse.Meta
	10, // 6: ethereum.eth.v1.PeersResponse.data:type_name -> ethereum.eth.v1.Peer
	14, // 7: ethereum.eth.v1.PeerCountResponse.data:type_name -> ethereum.eth.v1.PeerCountResponse.PeerCount
	1,  // 8: ethereum.eth.v1.Peer.state:type_name -> ethereum.eth.v1.ConnectionState
	0,  // 9: ethereum.eth.v1.Peer.direction:type_name -> ethereum.eth.v1.PeerDirection
	12, // 10: ethereum.eth.v1.VersionResponse.data:type_name -> ethereum.eth.v1.Version
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_eth_v1_node_proto_init() }
func file_eth_v1_node_proto_init() {
	if File_eth_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerResponse_Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1_node_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerCountResponse_PeerCount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eth_v1_node_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1_node_proto_goTypes,
		DependencyIndexes: file_eth_v1_node_proto_depIdxs,
		EnumInfos:         file_eth_v1_node_proto_enumTypes,
		MessageInfos:      file_eth_v1_node_proto_msgTypes,
	}.Build()
	File_eth_v1_node_proto = out.File
	file_eth_v1_node_proto_rawDesc = nil
	file_eth_v1_node_proto_goTypes = nil
	file_eth_v1_node_proto_depIdxs = nil
}
