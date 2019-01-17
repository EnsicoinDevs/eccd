// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Block struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Version              uint32   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Flags                []string `protobuf:"bytes,3,rep,name=flags,proto3" json:"flags,omitempty"`
	HashPrevBlock        string   `protobuf:"bytes,4,opt,name=hashPrevBlock,proto3" json:"hashPrevBlock,omitempty"`
	HashMerkleRoot       string   `protobuf:"bytes,5,opt,name=hashMerkleRoot,proto3" json:"hashMerkleRoot,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Height               uint32   `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Bits                 uint32   `protobuf:"varint,8,opt,name=bits,proto3" json:"bits,omitempty"`
	Nonce                uint64   `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Txs                  []*Tx    `protobuf:"bytes,10,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *Block) GetHashPrevBlock() string {
	if m != nil {
		return m.HashPrevBlock
	}
	return ""
}

func (m *Block) GetHashMerkleRoot() string {
	if m != nil {
		return m.HashMerkleRoot
	}
	return ""
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetBits() uint32 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *Block) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Block) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Tx struct {
	Hash                 string    `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Version              uint32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Flags                []string  `protobuf:"bytes,3,rep,name=flags,proto3" json:"flags,omitempty"`
	Inputs               []*Input  `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*Output `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{1}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Tx) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Tx) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *Tx) GetInputs() []*Input {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Tx) GetOutputs() []*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type Input struct {
	PreviousOutput       *Outpoint `protobuf:"bytes,1,opt,name=previousOutput,proto3" json:"previousOutput,omitempty"`
	Script               []byte    `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{2}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (dst *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(dst, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetPreviousOutput() *Outpoint {
	if m != nil {
		return m.PreviousOutput
	}
	return nil
}

func (m *Input) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

type Outpoint struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outpoint) Reset()         { *m = Outpoint{} }
func (m *Outpoint) String() string { return proto.CompactTextString(m) }
func (*Outpoint) ProtoMessage()    {}
func (*Outpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{3}
}
func (m *Outpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outpoint.Unmarshal(m, b)
}
func (m *Outpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outpoint.Marshal(b, m, deterministic)
}
func (dst *Outpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outpoint.Merge(dst, src)
}
func (m *Outpoint) XXX_Size() int {
	return xxx_messageInfo_Outpoint.Size(m)
}
func (m *Outpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Outpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Outpoint proto.InternalMessageInfo

func (m *Outpoint) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Outpoint) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type Output struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Script               []byte   `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{4}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (dst *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(dst, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Output) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

type GetBlockRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockRequest) Reset()         { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()    {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{5}
}
func (m *GetBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockRequest.Unmarshal(m, b)
}
func (m *GetBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockRequest.Merge(dst, src)
}
func (m *GetBlockRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockRequest.Size(m)
}
func (m *GetBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockRequest proto.InternalMessageInfo

func (m *GetBlockRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockReply struct {
	Block                *Block   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockReply) Reset()         { *m = GetBlockReply{} }
func (m *GetBlockReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockReply) ProtoMessage()    {}
func (*GetBlockReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{6}
}
func (m *GetBlockReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockReply.Unmarshal(m, b)
}
func (m *GetBlockReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockReply.Marshal(b, m, deterministic)
}
func (dst *GetBlockReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockReply.Merge(dst, src)
}
func (m *GetBlockReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockReply.Size(m)
}
func (m *GetBlockReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockReply proto.InternalMessageInfo

func (m *GetBlockReply) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetBestBlockHashRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBestBlockHashRequest) Reset()         { *m = GetBestBlockHashRequest{} }
func (m *GetBestBlockHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetBestBlockHashRequest) ProtoMessage()    {}
func (*GetBestBlockHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{7}
}
func (m *GetBestBlockHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBestBlockHashRequest.Unmarshal(m, b)
}
func (m *GetBestBlockHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBestBlockHashRequest.Marshal(b, m, deterministic)
}
func (dst *GetBestBlockHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBestBlockHashRequest.Merge(dst, src)
}
func (m *GetBestBlockHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetBestBlockHashRequest.Size(m)
}
func (m *GetBestBlockHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBestBlockHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBestBlockHashRequest proto.InternalMessageInfo

type GetBestBlockHashReply struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBestBlockHashReply) Reset()         { *m = GetBestBlockHashReply{} }
func (m *GetBestBlockHashReply) String() string { return proto.CompactTextString(m) }
func (*GetBestBlockHashReply) ProtoMessage()    {}
func (*GetBestBlockHashReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a9760e3eb7c37bd6, []int{8}
}
func (m *GetBestBlockHashReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBestBlockHashReply.Unmarshal(m, b)
}
func (m *GetBestBlockHashReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBestBlockHashReply.Marshal(b, m, deterministic)
}
func (dst *GetBestBlockHashReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBestBlockHashReply.Merge(dst, src)
}
func (m *GetBestBlockHashReply) XXX_Size() int {
	return xxx_messageInfo_GetBestBlockHashReply.Size(m)
}
func (m *GetBestBlockHashReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBestBlockHashReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBestBlockHashReply proto.InternalMessageInfo

func (m *GetBestBlockHashReply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Block)(nil), "rpc.Block")
	proto.RegisterType((*Tx)(nil), "rpc.Tx")
	proto.RegisterType((*Input)(nil), "rpc.Input")
	proto.RegisterType((*Outpoint)(nil), "rpc.Outpoint")
	proto.RegisterType((*Output)(nil), "rpc.Output")
	proto.RegisterType((*GetBlockRequest)(nil), "rpc.GetBlockRequest")
	proto.RegisterType((*GetBlockReply)(nil), "rpc.GetBlockReply")
	proto.RegisterType((*GetBestBlockHashRequest)(nil), "rpc.GetBestBlockHashRequest")
	proto.RegisterType((*GetBestBlockHashReply)(nil), "rpc.GetBestBlockHashReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainClient interface {
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockReply, error)
	GetBestBlockHash(ctx context.Context, in *GetBestBlockHashRequest, opts ...grpc.CallOption) (*GetBestBlockHashReply, error)
}

type blockchainClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainClient(cc *grpc.ClientConn) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockReply, error) {
	out := new(GetBlockReply)
	err := c.cc.Invoke(ctx, "/rpc.Blockchain/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBestBlockHash(ctx context.Context, in *GetBestBlockHashRequest, opts ...grpc.CallOption) (*GetBestBlockHashReply, error) {
	out := new(GetBestBlockHashReply)
	err := c.cc.Invoke(ctx, "/rpc.Blockchain/GetBestBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
type BlockchainServer interface {
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockReply, error)
	GetBestBlockHash(context.Context, *GetBestBlockHashRequest) (*GetBestBlockHashReply, error)
}

func RegisterBlockchainServer(s *grpc.Server, srv BlockchainServer) {
	s.RegisterService(&_Blockchain_serviceDesc, srv)
}

func _Blockchain_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Blockchain/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBestBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBestBlockHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBestBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Blockchain/GetBestBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBestBlockHash(ctx, req.(*GetBestBlockHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blockchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _Blockchain_GetBlock_Handler,
		},
		{
			MethodName: "GetBestBlockHash",
			Handler:    _Blockchain_GetBestBlockHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_a9760e3eb7c37bd6) }

var fileDescriptor_rpc_a9760e3eb7c37bd6 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x1c, 0xff, 0x24, 0x9e, 0x7c, 0x29, 0x68, 0x14, 0x60, 0x1b, 0xf5, 0xc2, 0x5a, 0x51,
	0x64, 0x09, 0xa9, 0x12, 0x01, 0xfa, 0x00, 0xdc, 0x00, 0x17, 0xfc, 0x68, 0x55, 0x71, 0xef, 0x98,
	0xa5, 0x5e, 0xd5, 0xf5, 0x1a, 0xef, 0x3a, 0x4a, 0x9f, 0x83, 0x3e, 0x30, 0xda, 0x59, 0xa7, 0x15,
	0x26, 0xbd, 0xe3, 0xce, 0x73, 0xe6, 0x9c, 0x39, 0x67, 0x46, 0x6b, 0x48, 0xbb, 0xb6, 0x3c, 0x6b,
	0x3b, 0x6d, 0x35, 0x86, 0x5d, 0x5b, 0xf2, 0xdb, 0x09, 0xc4, 0xef, 0x6a, 0x5d, 0x5e, 0x21, 0x42,
	0x54, 0x15, 0xa6, 0x62, 0x41, 0x16, 0xe4, 0xa9, 0xa0, 0x6f, 0x64, 0x30, 0xdd, 0xca, 0xce, 0x28,
	0xdd, 0xb0, 0x49, 0x16, 0xe4, 0x0b, 0xb1, 0x2f, 0x71, 0x09, 0xf1, 0x8f, 0xba, 0xb8, 0x34, 0x2c,
	0xcc, 0xc2, 0x3c, 0x15, 0xbe, 0xc0, 0xe7, 0xb0, 0x70, 0xba, 0xaf, 0x9d, 0xdc, 0xd2, 0x50, 0x16,
	0xd1, 0xb0, 0x3f, 0x41, 0x7c, 0x01, 0x47, 0x0e, 0xf8, 0x24, 0xbb, 0xab, 0x5a, 0x0a, 0xad, 0x2d,
	0x8b, 0x89, 0x36, 0x42, 0xf1, 0x04, 0x52, 0xab, 0xae, 0xa5, 0xb1, 0xc5, 0x75, 0xcb, 0x92, 0x2c,
	0xc8, 0x43, 0x71, 0x0f, 0xe0, 0x53, 0x48, 0x2a, 0xa9, 0x2e, 0x2b, 0xcb, 0xa6, 0x14, 0x6d, 0xa8,
	0xdc, 0x1e, 0x1b, 0x65, 0x0d, 0x9b, 0x11, 0x4a, 0xdf, 0x2e, 0x6d, 0xa3, 0x9b, 0x52, 0xb2, 0x34,
	0x0b, 0xf2, 0x48, 0xf8, 0x02, 0x8f, 0x21, 0xb4, 0x3b, 0xc3, 0x20, 0x0b, 0xf3, 0xf9, 0x7a, 0x7a,
	0xe6, 0x2e, 0x73, 0xb1, 0x13, 0x0e, 0xe3, 0xbf, 0x02, 0x98, 0x5c, 0xec, 0xfe, 0xc9, 0x4d, 0x38,
	0x24, 0xaa, 0x69, 0x7b, 0x6b, 0x58, 0x44, 0x46, 0x40, 0x46, 0x1f, 0x1d, 0x24, 0x86, 0x0e, 0x9e,
	0xc2, 0x54, 0xf7, 0x96, 0x48, 0x31, 0x91, 0xe6, 0x44, 0xfa, 0x42, 0x98, 0xd8, 0xf7, 0xf8, 0x37,
	0x88, 0x49, 0x87, 0x6f, 0xe1, 0xa8, 0xed, 0xe4, 0x56, 0xe9, 0xde, 0x78, 0x0e, 0x25, 0x9c, 0xaf,
	0x17, 0x77, 0x32, 0xad, 0x1a, 0x2b, 0x46, 0x24, 0x77, 0x32, 0x53, 0x76, 0xaa, 0xb5, 0x94, 0xfc,
	0x7f, 0x31, 0x54, 0xfc, 0x0d, 0xcc, 0xf6, 0x9a, 0x83, 0x2b, 0x2f, 0x21, 0x56, 0xcd, 0x77, 0xb9,
	0x1b, 0x16, 0xf6, 0x05, 0x3f, 0x87, 0x64, 0x98, 0xbb, 0x84, 0x78, 0x5b, 0xd4, 0xbd, 0x24, 0x51,
	0x24, 0x7c, 0xf1, 0xa0, 0xdb, 0x29, 0x3c, 0x7a, 0x2f, 0x2d, 0x3d, 0x05, 0x21, 0x7f, 0xf6, 0xd2,
	0x1c, 0x34, 0xe5, 0xaf, 0x60, 0x71, 0x4f, 0x6b, 0xeb, 0x1b, 0xcc, 0x20, 0xde, 0xd0, 0xa3, 0xf2,
	0xbb, 0xfa, 0x3b, 0xfa, 0xbe, 0x6f, 0xf0, 0x63, 0x78, 0xe6, 0x24, 0xd2, 0x78, 0xd9, 0x87, 0xc2,
	0x54, 0x83, 0x03, 0x7f, 0x09, 0x4f, 0xfe, 0x6e, 0xb9, 0xa9, 0x07, 0xac, 0xd7, 0xb7, 0x01, 0x00,
	0xd1, 0xca, 0xaa, 0x50, 0x0d, 0x9e, 0xc3, 0x6c, 0x9f, 0x04, 0x97, 0xe4, 0x3a, 0xca, 0xbf, 0xc2,
	0x11, 0xda, 0xd6, 0x37, 0xfc, 0x3f, 0xfc, 0x0c, 0x8f, 0xc7, 0x9e, 0x78, 0x72, 0xc7, 0x3c, 0x90,
	0x72, 0xb5, 0x7a, 0xa0, 0x4b, 0xf3, 0x36, 0x09, 0xfd, 0xb7, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x50, 0x91, 0x1d, 0xc4, 0x03, 0x00, 0x00,
}
