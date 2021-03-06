// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_OrderSrv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

// 发布订单的请求
type PostOrdersRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostOrdersRequest) Reset()         { *m = PostOrdersRequest{} }
func (m *PostOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*PostOrdersRequest) ProtoMessage()    {}
func (*PostOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *PostOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostOrdersRequest.Unmarshal(m, b)
}
func (m *PostOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostOrdersRequest.Marshal(b, m, deterministic)
}
func (m *PostOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostOrdersRequest.Merge(m, src)
}
func (m *PostOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_PostOrdersRequest.Size(m)
}
func (m *PostOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostOrdersRequest proto.InternalMessageInfo

func (m *PostOrdersRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *PostOrdersRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// 发布订单的响应
type PostOrdersResponse struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	// 返回订单id
	OrderId              int64    `protobuf:"varint,3,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostOrdersResponse) Reset()         { *m = PostOrdersResponse{} }
func (m *PostOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*PostOrdersResponse) ProtoMessage()    {}
func (*PostOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *PostOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostOrdersResponse.Unmarshal(m, b)
}
func (m *PostOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostOrdersResponse.Marshal(b, m, deterministic)
}
func (m *PostOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostOrdersResponse.Merge(m, src)
}
func (m *PostOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_PostOrdersResponse.Size(m)
}
func (m *PostOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostOrdersResponse proto.InternalMessageInfo

func (m *PostOrdersResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *PostOrdersResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PostOrdersResponse) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

// 查看房东/租客订单的请求
type UserOrderRequest struct {
	// 标识是房东还是租客可看，role=custom：租客查看，role=landlord房东查看
	Role                 string   `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
	SessionId            string   `protobuf:"bytes,2,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOrderRequest) Reset()         { *m = UserOrderRequest{} }
func (m *UserOrderRequest) String() string { return proto.CompactTextString(m) }
func (*UserOrderRequest) ProtoMessage()    {}
func (*UserOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *UserOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrderRequest.Unmarshal(m, b)
}
func (m *UserOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrderRequest.Marshal(b, m, deterministic)
}
func (m *UserOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrderRequest.Merge(m, src)
}
func (m *UserOrderRequest) XXX_Size() int {
	return xxx_messageInfo_UserOrderRequest.Size(m)
}
func (m *UserOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrderRequest proto.InternalMessageInfo

func (m *UserOrderRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UserOrderRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 查看房东/租客订单的响应
type UserOrderResponse struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	// 订单信息，字节流
	Orders               []byte   `protobuf:"bytes,3,opt,name=Orders,proto3" json:"Orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOrderResponse) Reset()         { *m = UserOrderResponse{} }
func (m *UserOrderResponse) String() string { return proto.CompactTextString(m) }
func (*UserOrderResponse) ProtoMessage()    {}
func (*UserOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *UserOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrderResponse.Unmarshal(m, b)
}
func (m *UserOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrderResponse.Marshal(b, m, deterministic)
}
func (m *UserOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrderResponse.Merge(m, src)
}
func (m *UserOrderResponse) XXX_Size() int {
	return xxx_messageInfo_UserOrderResponse.Size(m)
}
func (m *UserOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrderResponse proto.InternalMessageInfo

func (m *UserOrderResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UserOrderResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserOrderResponse) GetOrders() []byte {
	if m != nil {
		return m.Orders
	}
	return nil
}

// 房东同意/拒绝订单的请求
type PutOrdersRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	// 订单编号
	OrderId string `protobuf:"bytes,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	// 参数的合法性
	Action               string   `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutOrdersRequest) Reset()         { *m = PutOrdersRequest{} }
func (m *PutOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*PutOrdersRequest) ProtoMessage()    {}
func (*PutOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *PutOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutOrdersRequest.Unmarshal(m, b)
}
func (m *PutOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutOrdersRequest.Marshal(b, m, deterministic)
}
func (m *PutOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutOrdersRequest.Merge(m, src)
}
func (m *PutOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_PutOrdersRequest.Size(m)
}
func (m *PutOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutOrdersRequest proto.InternalMessageInfo

func (m *PutOrdersRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *PutOrdersRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *PutOrdersRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

// 房东同意/拒绝订单的响应
type PutOrdersResponse struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutOrdersResponse) Reset()         { *m = PutOrdersResponse{} }
func (m *PutOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*PutOrdersResponse) ProtoMessage()    {}
func (*PutOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *PutOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutOrdersResponse.Unmarshal(m, b)
}
func (m *PutOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutOrdersResponse.Marshal(b, m, deterministic)
}
func (m *PutOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutOrdersResponse.Merge(m, src)
}
func (m *PutOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_PutOrdersResponse.Size(m)
}
func (m *PutOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutOrdersResponse proto.InternalMessageInfo

func (m *PutOrdersResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *PutOrdersResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// 用户评论订单的请求
type UserCommentRequest struct {
	// 订单Id
	OrderId string `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	// 评价内容
	Comment              string   `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
	SessionId            string   `protobuf:"bytes,3,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCommentRequest) Reset()         { *m = UserCommentRequest{} }
func (m *UserCommentRequest) String() string { return proto.CompactTextString(m) }
func (*UserCommentRequest) ProtoMessage()    {}
func (*UserCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{7}
}

func (m *UserCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCommentRequest.Unmarshal(m, b)
}
func (m *UserCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCommentRequest.Marshal(b, m, deterministic)
}
func (m *UserCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCommentRequest.Merge(m, src)
}
func (m *UserCommentRequest) XXX_Size() int {
	return xxx_messageInfo_UserCommentRequest.Size(m)
}
func (m *UserCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCommentRequest proto.InternalMessageInfo

func (m *UserCommentRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *UserCommentRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *UserCommentRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 用户评论订单的响应
type UserCommentResponse struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCommentResponse) Reset()         { *m = UserCommentResponse{} }
func (m *UserCommentResponse) String() string { return proto.CompactTextString(m) }
func (*UserCommentResponse) ProtoMessage()    {}
func (*UserCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{8}
}

func (m *UserCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCommentResponse.Unmarshal(m, b)
}
func (m *UserCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCommentResponse.Marshal(b, m, deterministic)
}
func (m *UserCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCommentResponse.Merge(m, src)
}
func (m *UserCommentResponse) XXX_Size() int {
	return xxx_messageInfo_UserCommentResponse.Size(m)
}
func (m *UserCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserCommentResponse proto.InternalMessageInfo

func (m *UserCommentResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UserCommentResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.OrderSrv.Message")
	proto.RegisterType((*PostOrdersRequest)(nil), "go.micro.srv.OrderSrv.PostOrdersRequest")
	proto.RegisterType((*PostOrdersResponse)(nil), "go.micro.srv.OrderSrv.PostOrdersResponse")
	proto.RegisterType((*UserOrderRequest)(nil), "go.micro.srv.OrderSrv.UserOrderRequest")
	proto.RegisterType((*UserOrderResponse)(nil), "go.micro.srv.OrderSrv.UserOrderResponse")
	proto.RegisterType((*PutOrdersRequest)(nil), "go.micro.srv.OrderSrv.PutOrdersRequest")
	proto.RegisterType((*PutOrdersResponse)(nil), "go.micro.srv.OrderSrv.PutOrdersResponse")
	proto.RegisterType((*UserCommentRequest)(nil), "go.micro.srv.OrderSrv.UserCommentRequest")
	proto.RegisterType((*UserCommentResponse)(nil), "go.micro.srv.OrderSrv.UserCommentResponse")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x09, 0xa1, 0x44, 0x19, 0x71, 0x80, 0x69, 0x8b, 0x22, 0xe8, 0x01, 0xf9, 0xd2, 0xd0,
	0x43, 0x2a, 0xb5, 0x4f, 0x40, 0x69, 0x54, 0x71, 0x40, 0x45, 0x46, 0x3d, 0x54, 0xea, 0xa1, 0x90,
	0xb8, 0x08, 0x89, 0xc4, 0xd4, 0x0e, 0xa8, 0xbc, 0x4b, 0x1f, 0x76, 0x85, 0x63, 0x43, 0x12, 0x96,
	0x5d, 0x96, 0x53, 0xfc, 0x5b, 0xe3, 0xf9, 0x3f, 0x8f, 0x7f, 0x05, 0xfa, 0x5b, 0xc1, 0x33, 0xfe,
	0x91, 0xfd, 0x5b, 0x24, 0xdb, 0x0d, 0x33, 0xdf, 0x40, 0xed, 0xe2, 0xdb, 0x15, 0x0f, 0x92, 0x75,
	0x24, 0x78, 0x20, 0xc5, 0x3e, 0xf8, 0x2e, 0x62, 0x26, 0xe6, 0x62, 0x4f, 0xfa, 0xe0, 0x4c, 0x99,
	0x94, 0x8b, 0x15, 0xc3, 0x36, 0xd8, 0x72, 0x71, 0xf0, 0xac, 0x81, 0xe5, 0xbb, 0xf4, 0xb8, 0x24,
	0x21, 0x74, 0x66, 0x5c, 0x66, 0xaa, 0x58, 0x52, 0xf6, 0x77, 0xc7, 0x64, 0x86, 0xef, 0xc0, 0x9d,
	0x33, 0x29, 0xd7, 0x3c, 0x9d, 0xc4, 0xba, 0xf8, 0xbc, 0x81, 0x08, 0x8d, 0x2f, 0x3c, 0x3e, 0x78,
	0xf5, 0x81, 0xe5, 0xb7, 0xa8, 0x5a, 0x93, 0x5f, 0x80, 0xc5, 0x36, 0x72, 0xcb, 0x53, 0xc9, 0xf0,
	0x0d, 0xbc, 0x0a, 0x85, 0x48, 0xb9, 0xee, 0x91, 0x0b, 0xec, 0x42, 0x33, 0x14, 0x62, 0x2a, 0x57,
	0xaa, 0x83, 0x4b, 0xb5, 0x42, 0x0f, 0x1c, 0x75, 0x7e, 0x12, 0x7b, 0xf6, 0xc0, 0xf2, 0x6d, 0x6a,
	0x24, 0xf9, 0x0a, 0xed, 0x1f, 0x92, 0x09, 0x25, 0x0d, 0x23, 0x42, 0x83, 0xf2, 0x0d, 0xd3, 0xad,
	0xd5, 0xba, 0xcc, 0x5d, 0xaf, 0x70, 0x93, 0x9f, 0xd0, 0x29, 0x74, 0xb9, 0x0b, 0xb1, 0x0b, 0xcd,
	0xfc, 0x8a, 0x8a, 0xb0, 0x45, 0xb5, 0x22, 0x4b, 0x68, 0xcf, 0x76, 0x2f, 0x1a, 0x62, 0xe1, 0xb2,
	0xb9, 0x85, 0x91, 0x47, 0x8f, 0x51, 0x94, 0xad, 0x79, 0xaa, 0x3c, 0x5c, 0xaa, 0x15, 0x19, 0x41,
	0xa7, 0xe0, 0x71, 0x0f, 0x3e, 0xf9, 0x03, 0x78, 0x9c, 0xc0, 0x98, 0x27, 0x09, 0x4b, 0x33, 0x03,
	0x5a, 0x40, 0xb1, 0xca, 0x28, 0x1e, 0x38, 0xba, 0xd6, 0x40, 0x6a, 0x59, 0xbe, 0x9c, 0x5d, 0x9d,
	0xf4, 0x18, 0x5e, 0x97, 0x7c, 0xee, 0x81, 0xfd, 0xf4, 0xdf, 0x06, 0x27, 0xcc, 0xf3, 0x8d, 0x11,
	0xc0, 0x39, 0x5e, 0xe8, 0x07, 0x8f, 0x06, 0x3d, 0xb8, 0x08, 0x72, 0x6f, 0x78, 0x43, 0x65, 0x0e,
	0x47, 0x6a, 0x18, 0x41, 0xeb, 0x1b, 0xcb, 0x4e, 0x11, 0xc1, 0xf7, 0x57, 0x0e, 0x57, 0xa3, 0xd8,
	0xf3, 0x9f, 0x2f, 0x3c, 0x99, 0xfc, 0x06, 0xf7, 0xf4, 0x8a, 0x57, 0x1d, 0xaa, 0x59, 0xba, 0xea,
	0x70, 0x11, 0x08, 0x52, 0x43, 0x06, 0x30, 0xdb, 0x65, 0xe6, 0xa1, 0x86, 0x4f, 0xb0, 0x95, 0x73,
	0xd0, 0xfb, 0x70, 0x4b, 0xa9, 0xb1, 0x59, 0x36, 0xd5, 0x3f, 0xe7, 0xf3, 0x43, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x48, 0xfe, 0xdf, 0x92, 0x04, 0x00, 0x00,
}
