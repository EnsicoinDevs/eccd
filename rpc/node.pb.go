// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package ensicoin_rpc

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

type Error int32

const (
	Error_INVALID_DATA Error = 0
)

var Error_name = map[int32]string{
	0: "INVALID_DATA",
}
var Error_value = map[string]int32{
	"INVALID_DATA": 0,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{0}
}

type Reply struct {
	Error                Error    `protobuf:"varint,1,opt,name=error,proto3,enum=ensicoin_rpc.Error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{0}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_INVALID_DATA
}

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{1}
}
func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(dst, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type GetInfoReply struct {
	Implementation       string   `protobuf:"bytes,1,opt,name=implementation,proto3" json:"implementation,omitempty"`
	ProtocolVersion      uint32   `protobuf:"varint,2,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	BestBlockHash        []byte   `protobuf:"bytes,3,opt,name=best_block_hash,json=bestBlockHash,proto3" json:"best_block_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoReply) Reset()         { *m = GetInfoReply{} }
func (m *GetInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetInfoReply) ProtoMessage()    {}
func (*GetInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{2}
}
func (m *GetInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoReply.Unmarshal(m, b)
}
func (m *GetInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoReply.Marshal(b, m, deterministic)
}
func (dst *GetInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoReply.Merge(dst, src)
}
func (m *GetInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetInfoReply.Size(m)
}
func (m *GetInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoReply proto.InternalMessageInfo

func (m *GetInfoReply) GetImplementation() string {
	if m != nil {
		return m.Implementation
	}
	return ""
}

func (m *GetInfoReply) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func (m *GetInfoReply) GetBestBlockHash() []byte {
	if m != nil {
		return m.BestBlockHash
	}
	return nil
}

type PublishRawTxRequest struct {
	RawTx                []byte   `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRawTxRequest) Reset()         { *m = PublishRawTxRequest{} }
func (m *PublishRawTxRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRawTxRequest) ProtoMessage()    {}
func (*PublishRawTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{3}
}
func (m *PublishRawTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRawTxRequest.Unmarshal(m, b)
}
func (m *PublishRawTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRawTxRequest.Marshal(b, m, deterministic)
}
func (dst *PublishRawTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRawTxRequest.Merge(dst, src)
}
func (m *PublishRawTxRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRawTxRequest.Size(m)
}
func (m *PublishRawTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRawTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRawTxRequest proto.InternalMessageInfo

func (m *PublishRawTxRequest) GetRawTx() []byte {
	if m != nil {
		return m.RawTx
	}
	return nil
}

type PublishRawTxReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRawTxReply) Reset()         { *m = PublishRawTxReply{} }
func (m *PublishRawTxReply) String() string { return proto.CompactTextString(m) }
func (*PublishRawTxReply) ProtoMessage()    {}
func (*PublishRawTxReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{4}
}
func (m *PublishRawTxReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRawTxReply.Unmarshal(m, b)
}
func (m *PublishRawTxReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRawTxReply.Marshal(b, m, deterministic)
}
func (dst *PublishRawTxReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRawTxReply.Merge(dst, src)
}
func (m *PublishRawTxReply) XXX_Size() int {
	return xxx_messageInfo_PublishRawTxReply.Size(m)
}
func (m *PublishRawTxReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRawTxReply.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRawTxReply proto.InternalMessageInfo

type PublishRawBlockRequest struct {
	RawBlock             []byte   `protobuf:"bytes,1,opt,name=raw_block,json=rawBlock,proto3" json:"raw_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRawBlockRequest) Reset()         { *m = PublishRawBlockRequest{} }
func (m *PublishRawBlockRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRawBlockRequest) ProtoMessage()    {}
func (*PublishRawBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{5}
}
func (m *PublishRawBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRawBlockRequest.Unmarshal(m, b)
}
func (m *PublishRawBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRawBlockRequest.Marshal(b, m, deterministic)
}
func (dst *PublishRawBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRawBlockRequest.Merge(dst, src)
}
func (m *PublishRawBlockRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRawBlockRequest.Size(m)
}
func (m *PublishRawBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRawBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRawBlockRequest proto.InternalMessageInfo

func (m *PublishRawBlockRequest) GetRawBlock() []byte {
	if m != nil {
		return m.RawBlock
	}
	return nil
}

type PublishRawBlockReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRawBlockReply) Reset()         { *m = PublishRawBlockReply{} }
func (m *PublishRawBlockReply) String() string { return proto.CompactTextString(m) }
func (*PublishRawBlockReply) ProtoMessage()    {}
func (*PublishRawBlockReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{6}
}
func (m *PublishRawBlockReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRawBlockReply.Unmarshal(m, b)
}
func (m *PublishRawBlockReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRawBlockReply.Marshal(b, m, deterministic)
}
func (dst *PublishRawBlockReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRawBlockReply.Merge(dst, src)
}
func (m *PublishRawBlockReply) XXX_Size() int {
	return xxx_messageInfo_PublishRawBlockReply.Size(m)
}
func (m *PublishRawBlockReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRawBlockReply.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRawBlockReply proto.InternalMessageInfo

type GetBlockTemplateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockTemplateRequest) Reset()         { *m = GetBlockTemplateRequest{} }
func (m *GetBlockTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockTemplateRequest) ProtoMessage()    {}
func (*GetBlockTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{7}
}
func (m *GetBlockTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockTemplateRequest.Unmarshal(m, b)
}
func (m *GetBlockTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlockTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockTemplateRequest.Merge(dst, src)
}
func (m *GetBlockTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockTemplateRequest.Size(m)
}
func (m *GetBlockTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockTemplateRequest proto.InternalMessageInfo

type GetBlockTemplateReply struct {
	Template             *BlockTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetBlockTemplateReply) Reset()         { *m = GetBlockTemplateReply{} }
func (m *GetBlockTemplateReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockTemplateReply) ProtoMessage()    {}
func (*GetBlockTemplateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{8}
}
func (m *GetBlockTemplateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockTemplateReply.Unmarshal(m, b)
}
func (m *GetBlockTemplateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockTemplateReply.Marshal(b, m, deterministic)
}
func (dst *GetBlockTemplateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockTemplateReply.Merge(dst, src)
}
func (m *GetBlockTemplateReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockTemplateReply.Size(m)
}
func (m *GetBlockTemplateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockTemplateReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockTemplateReply proto.InternalMessageInfo

func (m *GetBlockTemplateReply) GetTemplate() *BlockTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

type GetBlockByHashRequest struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockByHashRequest) Reset()         { *m = GetBlockByHashRequest{} }
func (m *GetBlockByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashRequest) ProtoMessage()    {}
func (*GetBlockByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{9}
}
func (m *GetBlockByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashRequest.Unmarshal(m, b)
}
func (m *GetBlockByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlockByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashRequest.Merge(dst, src)
}
func (m *GetBlockByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashRequest.Size(m)
}
func (m *GetBlockByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashRequest proto.InternalMessageInfo

func (m *GetBlockByHashRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type GetBlockByHashReply struct {
	Block                *Block   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockByHashReply) Reset()         { *m = GetBlockByHashReply{} }
func (m *GetBlockByHashReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashReply) ProtoMessage()    {}
func (*GetBlockByHashReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{10}
}
func (m *GetBlockByHashReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashReply.Unmarshal(m, b)
}
func (m *GetBlockByHashReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashReply.Marshal(b, m, deterministic)
}
func (dst *GetBlockByHashReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashReply.Merge(dst, src)
}
func (m *GetBlockByHashReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashReply.Size(m)
}
func (m *GetBlockByHashReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashReply proto.InternalMessageInfo

func (m *GetBlockByHashReply) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetTxByHashRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTxByHashRequest) Reset()         { *m = GetTxByHashRequest{} }
func (m *GetTxByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxByHashRequest) ProtoMessage()    {}
func (*GetTxByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{11}
}
func (m *GetTxByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxByHashRequest.Unmarshal(m, b)
}
func (m *GetTxByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxByHashRequest.Marshal(b, m, deterministic)
}
func (dst *GetTxByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxByHashRequest.Merge(dst, src)
}
func (m *GetTxByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetTxByHashRequest.Size(m)
}
func (m *GetTxByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxByHashRequest proto.InternalMessageInfo

type GetTxByHashReply struct {
	Tx                   *Tx      `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTxByHashReply) Reset()         { *m = GetTxByHashReply{} }
func (m *GetTxByHashReply) String() string { return proto.CompactTextString(m) }
func (*GetTxByHashReply) ProtoMessage()    {}
func (*GetTxByHashReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_003585f98e6f42a2, []int{12}
}
func (m *GetTxByHashReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxByHashReply.Unmarshal(m, b)
}
func (m *GetTxByHashReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxByHashReply.Marshal(b, m, deterministic)
}
func (dst *GetTxByHashReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxByHashReply.Merge(dst, src)
}
func (m *GetTxByHashReply) XXX_Size() int {
	return xxx_messageInfo_GetTxByHashReply.Size(m)
}
func (m *GetTxByHashReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxByHashReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxByHashReply proto.InternalMessageInfo

func (m *GetTxByHashReply) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*Reply)(nil), "ensicoin_rpc.Reply")
	proto.RegisterType((*GetInfoRequest)(nil), "ensicoin_rpc.GetInfoRequest")
	proto.RegisterType((*GetInfoReply)(nil), "ensicoin_rpc.GetInfoReply")
	proto.RegisterType((*PublishRawTxRequest)(nil), "ensicoin_rpc.PublishRawTxRequest")
	proto.RegisterType((*PublishRawTxReply)(nil), "ensicoin_rpc.PublishRawTxReply")
	proto.RegisterType((*PublishRawBlockRequest)(nil), "ensicoin_rpc.PublishRawBlockRequest")
	proto.RegisterType((*PublishRawBlockReply)(nil), "ensicoin_rpc.PublishRawBlockReply")
	proto.RegisterType((*GetBlockTemplateRequest)(nil), "ensicoin_rpc.GetBlockTemplateRequest")
	proto.RegisterType((*GetBlockTemplateReply)(nil), "ensicoin_rpc.GetBlockTemplateReply")
	proto.RegisterType((*GetBlockByHashRequest)(nil), "ensicoin_rpc.GetBlockByHashRequest")
	proto.RegisterType((*GetBlockByHashReply)(nil), "ensicoin_rpc.GetBlockByHashReply")
	proto.RegisterType((*GetTxByHashRequest)(nil), "ensicoin_rpc.GetTxByHashRequest")
	proto.RegisterType((*GetTxByHashReply)(nil), "ensicoin_rpc.GetTxByHashReply")
	proto.RegisterEnum("ensicoin_rpc.Error", Error_name, Error_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error)
	PublishRawTx(ctx context.Context, opts ...grpc.CallOption) (Node_PublishRawTxClient, error)
	PublishRawBlock(ctx context.Context, opts ...grpc.CallOption) (Node_PublishRawBlockClient, error)
	GetBlockTemplate(ctx context.Context, in *GetBlockTemplateRequest, opts ...grpc.CallOption) (Node_GetBlockTemplateClient, error)
	GetBlockByHash(ctx context.Context, in *GetBlockByHashRequest, opts ...grpc.CallOption) (*GetBlockByHashReply, error)
	GetTxByHash(ctx context.Context, in *GetTxByHashRequest, opts ...grpc.CallOption) (*GetTxByHashReply, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error) {
	out := new(GetInfoReply)
	err := c.cc.Invoke(ctx, "/ensicoin_rpc.Node/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) PublishRawTx(ctx context.Context, opts ...grpc.CallOption) (Node_PublishRawTxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[0], "/ensicoin_rpc.Node/PublishRawTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodePublishRawTxClient{stream}
	return x, nil
}

type Node_PublishRawTxClient interface {
	Send(*PublishRawTxRequest) error
	CloseAndRecv() (*PublishRawTxReply, error)
	grpc.ClientStream
}

type nodePublishRawTxClient struct {
	grpc.ClientStream
}

func (x *nodePublishRawTxClient) Send(m *PublishRawTxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodePublishRawTxClient) CloseAndRecv() (*PublishRawTxReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishRawTxReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) PublishRawBlock(ctx context.Context, opts ...grpc.CallOption) (Node_PublishRawBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[1], "/ensicoin_rpc.Node/PublishRawBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodePublishRawBlockClient{stream}
	return x, nil
}

type Node_PublishRawBlockClient interface {
	Send(*PublishRawBlockRequest) error
	CloseAndRecv() (*PublishRawBlockReply, error)
	grpc.ClientStream
}

type nodePublishRawBlockClient struct {
	grpc.ClientStream
}

func (x *nodePublishRawBlockClient) Send(m *PublishRawBlockRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodePublishRawBlockClient) CloseAndRecv() (*PublishRawBlockReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishRawBlockReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) GetBlockTemplate(ctx context.Context, in *GetBlockTemplateRequest, opts ...grpc.CallOption) (Node_GetBlockTemplateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[2], "/ensicoin_rpc.Node/GetBlockTemplate", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeGetBlockTemplateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_GetBlockTemplateClient interface {
	Recv() (*GetBlockTemplateReply, error)
	grpc.ClientStream
}

type nodeGetBlockTemplateClient struct {
	grpc.ClientStream
}

func (x *nodeGetBlockTemplateClient) Recv() (*GetBlockTemplateReply, error) {
	m := new(GetBlockTemplateReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) GetBlockByHash(ctx context.Context, in *GetBlockByHashRequest, opts ...grpc.CallOption) (*GetBlockByHashReply, error) {
	out := new(GetBlockByHashReply)
	err := c.cc.Invoke(ctx, "/ensicoin_rpc.Node/GetBlockByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetTxByHash(ctx context.Context, in *GetTxByHashRequest, opts ...grpc.CallOption) (*GetTxByHashReply, error) {
	out := new(GetTxByHashReply)
	err := c.cc.Invoke(ctx, "/ensicoin_rpc.Node/GetTxByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error)
	PublishRawTx(Node_PublishRawTxServer) error
	PublishRawBlock(Node_PublishRawBlockServer) error
	GetBlockTemplate(*GetBlockTemplateRequest, Node_GetBlockTemplateServer) error
	GetBlockByHash(context.Context, *GetBlockByHashRequest) (*GetBlockByHashReply, error)
	GetTxByHash(context.Context, *GetTxByHashRequest) (*GetTxByHashReply, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ensicoin_rpc.Node/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_PublishRawTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).PublishRawTx(&nodePublishRawTxServer{stream})
}

type Node_PublishRawTxServer interface {
	SendAndClose(*PublishRawTxReply) error
	Recv() (*PublishRawTxRequest, error)
	grpc.ServerStream
}

type nodePublishRawTxServer struct {
	grpc.ServerStream
}

func (x *nodePublishRawTxServer) SendAndClose(m *PublishRawTxReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodePublishRawTxServer) Recv() (*PublishRawTxRequest, error) {
	m := new(PublishRawTxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_PublishRawBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).PublishRawBlock(&nodePublishRawBlockServer{stream})
}

type Node_PublishRawBlockServer interface {
	SendAndClose(*PublishRawBlockReply) error
	Recv() (*PublishRawBlockRequest, error)
	grpc.ServerStream
}

type nodePublishRawBlockServer struct {
	grpc.ServerStream
}

func (x *nodePublishRawBlockServer) SendAndClose(m *PublishRawBlockReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodePublishRawBlockServer) Recv() (*PublishRawBlockRequest, error) {
	m := new(PublishRawBlockRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_GetBlockTemplate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlockTemplateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).GetBlockTemplate(m, &nodeGetBlockTemplateServer{stream})
}

type Node_GetBlockTemplateServer interface {
	Send(*GetBlockTemplateReply) error
	grpc.ServerStream
}

type nodeGetBlockTemplateServer struct {
	grpc.ServerStream
}

func (x *nodeGetBlockTemplateServer) Send(m *GetBlockTemplateReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ensicoin_rpc.Node/GetBlockByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetBlockByHash(ctx, req.(*GetBlockByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ensicoin_rpc.Node/GetTxByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetTxByHash(ctx, req.(*GetTxByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ensicoin_rpc.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Node_GetInfo_Handler,
		},
		{
			MethodName: "GetBlockByHash",
			Handler:    _Node_GetBlockByHash_Handler,
		},
		{
			MethodName: "GetTxByHash",
			Handler:    _Node_GetTxByHash_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublishRawTx",
			Handler:       _Node_PublishRawTx_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PublishRawBlock",
			Handler:       _Node_PublishRawBlock_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetBlockTemplate",
			Handler:       _Node_GetBlockTemplate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_003585f98e6f42a2) }

var fileDescriptor_node_003585f98e6f42a2 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x6e, 0xd3, 0x40,
	0x10, 0xac, 0x43, 0x52, 0xda, 0xb5, 0x9b, 0x98, 0x4b, 0x5b, 0x5a, 0x17, 0x81, 0x7b, 0x40, 0x95,
	0x02, 0x8a, 0x90, 0x01, 0xf1, 0x4a, 0xaa, 0x56, 0x21, 0x12, 0xaa, 0x8a, 0x65, 0xe5, 0x81, 0x17,
	0xcb, 0x4e, 0x0f, 0xc5, 0xc2, 0xf1, 0x19, 0xfb, 0x4a, 0x9d, 0x4f, 0xe0, 0x13, 0xf8, 0x5b, 0x74,
	0x67, 0xbb, 0xf1, 0x39, 0x4d, 0xfa, 0x16, 0xcd, 0xce, 0xce, 0xee, 0xcd, 0x4e, 0x0c, 0x10, 0xd1,
	0x6b, 0xd2, 0x8f, 0x13, 0xca, 0x28, 0xd2, 0x48, 0x94, 0x06, 0x13, 0x1a, 0x44, 0x6e, 0x12, 0x4f,
	0x0c, 0x95, 0xcd, 0x63, 0x92, 0xe6, 0x25, 0x6c, 0x41, 0xcb, 0x26, 0x71, 0x38, 0x47, 0xa7, 0xd0,
	0x22, 0x49, 0x42, 0x93, 0x03, 0xc5, 0x54, 0x7a, 0x6d, 0xab, 0xdb, 0xaf, 0xf6, 0xf4, 0x2f, 0x78,
	0xc9, 0xce, 0x19, 0x58, 0x87, 0xf6, 0x90, 0xb0, 0x51, 0xf4, 0x93, 0xda, 0xe4, 0xf7, 0x0d, 0x49,
	0x19, 0xfe, 0xab, 0x80, 0x76, 0x07, 0x71, 0xb5, 0x13, 0x68, 0x07, 0xb3, 0x38, 0x24, 0x33, 0x12,
	0x31, 0x8f, 0x05, 0x34, 0x12, 0xb2, 0xdb, 0x76, 0x0d, 0x45, 0xa7, 0xa0, 0x8b, 0x3d, 0x26, 0x34,
	0x74, 0xff, 0x90, 0x24, 0xe5, 0xcc, 0x86, 0xa9, 0xf4, 0x76, 0xec, 0x4e, 0x89, 0x8f, 0x73, 0x18,
	0x9d, 0x40, 0xc7, 0x27, 0x29, 0x73, 0xfd, 0x90, 0x4e, 0x7e, 0xb9, 0x53, 0x2f, 0x9d, 0x1e, 0x3c,
	0x32, 0x95, 0x9e, 0x66, 0xef, 0x70, 0xf8, 0x8c, 0xa3, 0x5f, 0xbd, 0x74, 0x8a, 0xdf, 0x41, 0xf7,
	0xea, 0xc6, 0x0f, 0x83, 0x74, 0x6a, 0x7b, 0xb7, 0x4e, 0x56, 0xac, 0x88, 0xf6, 0x60, 0x33, 0xf1,
	0x6e, 0x5d, 0x96, 0x89, 0x4d, 0x34, 0xbb, 0x95, 0xf0, 0x2a, 0xee, 0xc2, 0x13, 0x99, 0x1d, 0x87,
	0x73, 0xfc, 0x09, 0xf6, 0x17, 0xa0, 0x50, 0x2e, 0x55, 0x8e, 0x60, 0x9b, 0xab, 0x88, 0x1d, 0x0a,
	0xa1, 0xad, 0xa4, 0xe0, 0xe0, 0x7d, 0xd8, 0x5d, 0x6a, 0xe3, 0x72, 0x87, 0xf0, 0x74, 0x48, 0xf2,
	0x0d, 0x1d, 0x32, 0x8b, 0x43, 0x8f, 0x91, 0xd2, 0xb8, 0x2b, 0xd8, 0x5b, 0x2e, 0x71, 0x03, 0x3f,
	0xc3, 0x16, 0x2b, 0x00, 0x31, 0x47, 0xb5, 0x8e, 0xe4, 0x8b, 0xc8, 0x3d, 0x77, 0x64, 0xfc, 0x76,
	0xa1, 0x78, 0x36, 0xe7, 0x86, 0x94, 0xab, 0x23, 0x68, 0x0a, 0xd3, 0xf2, 0xad, 0xc5, 0x6f, 0xfc,
	0x05, 0xba, 0x75, 0x72, 0x91, 0x85, 0xc5, 0x0b, 0xd5, 0x7a, 0x16, 0xf2, 0x97, 0xe5, 0x0c, 0xbc,
	0x0b, 0x68, 0x48, 0x98, 0x93, 0x49, 0xb3, 0xf0, 0x47, 0xd0, 0x25, 0x94, 0x8b, 0x9a, 0xd0, 0x28,
	0xcc, 0x57, 0x2d, 0x5d, 0x56, 0x74, 0x32, 0xbb, 0xc1, 0xb2, 0x37, 0x87, 0xd0, 0x12, 0x39, 0x43,
	0x3a, 0x68, 0xa3, 0xcb, 0xf1, 0xe0, 0xdb, 0xe8, 0xdc, 0x3d, 0x1f, 0x38, 0x03, 0x7d, 0xc3, 0xfa,
	0xd7, 0x84, 0xe6, 0x25, 0xbd, 0x26, 0xe8, 0x02, 0x1e, 0x17, 0x41, 0x43, 0xcf, 0x64, 0x11, 0x39,
	0x92, 0x86, 0xb1, 0xa2, 0xca, 0x0f, 0xb2, 0x81, 0xc6, 0xa0, 0x55, 0xcf, 0x8e, 0x8e, 0x65, 0xf6,
	0x3d, 0x01, 0x32, 0x5e, 0xac, 0xa3, 0x08, 0xd5, 0x9e, 0x82, 0x5c, 0xe8, 0xd4, 0x22, 0x80, 0x5e,
	0xad, 0xea, 0xab, 0x06, 0xcb, 0xc0, 0x0f, 0xb0, 0xca, 0x01, 0xbe, 0x70, 0x56, 0x3a, 0x3e, 0x7a,
	0xbd, 0xf4, 0xd4, 0xfb, 0xb2, 0x66, 0xbc, 0x7c, 0x88, 0x26, 0x66, 0xbc, 0x57, 0xd0, 0x0f, 0xf1,
	0xff, 0xae, 0xa4, 0x02, 0xad, 0x68, 0x95, 0x8e, 0x6e, 0x1c, 0xaf, 0x27, 0xe5, 0xc6, 0x7f, 0x07,
	0xb5, 0x92, 0x0c, 0x64, 0x2e, 0xf5, 0xd4, 0xa2, 0x64, 0x3c, 0x5f, 0xc3, 0x10, 0x92, 0xfe, 0xa6,
	0xf8, 0x52, 0x7c, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x61, 0x14, 0x3b, 0x1d, 0xf2, 0x04, 0x00,
	0x00,
}
