// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_UserInfoSrv

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

// 用户注册的请求
type UserRegistRequest struct {
	// 手机号
	Mobile string `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	// 短信验证码
	SmsCode              string   `protobuf:"bytes,3,opt,name=SmsCode,proto3" json:"SmsCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegistRequest) Reset()         { *m = UserRegistRequest{} }
func (m *UserRegistRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegistRequest) ProtoMessage()    {}
func (*UserRegistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *UserRegistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegistRequest.Unmarshal(m, b)
}
func (m *UserRegistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegistRequest.Marshal(b, m, deterministic)
}
func (m *UserRegistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegistRequest.Merge(m, src)
}
func (m *UserRegistRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegistRequest.Size(m)
}
func (m *UserRegistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegistRequest proto.InternalMessageInfo

func (m *UserRegistRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserRegistRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRegistRequest) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

// 用户注册的响应
type UserRegistResponse struct {
	// 错误码
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	// 将session id返回
	SessionId            string   `protobuf:"bytes,3,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegistResponse) Reset()         { *m = UserRegistResponse{} }
func (m *UserRegistResponse) String() string { return proto.CompactTextString(m) }
func (*UserRegistResponse) ProtoMessage()    {}
func (*UserRegistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *UserRegistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegistResponse.Unmarshal(m, b)
}
func (m *UserRegistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegistResponse.Marshal(b, m, deterministic)
}
func (m *UserRegistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegistResponse.Merge(m, src)
}
func (m *UserRegistResponse) XXX_Size() int {
	return xxx_messageInfo_UserRegistResponse.Size(m)
}
func (m *UserRegistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegistResponse proto.InternalMessageInfo

func (m *UserRegistResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UserRegistResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserRegistResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 获取Session的请求
type SessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (m *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(m, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 获取Session的响应
type SessionResponse struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//返回用户名
	UserName             string   `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionResponse) Reset()         { *m = SessionResponse{} }
func (m *SessionResponse) String() string { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()    {}
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *SessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionResponse.Unmarshal(m, b)
}
func (m *SessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionResponse.Marshal(b, m, deterministic)
}
func (m *SessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionResponse.Merge(m, src)
}
func (m *SessionResponse) XXX_Size() int {
	return xxx_messageInfo_SessionResponse.Size(m)
}
func (m *SessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SessionResponse proto.InternalMessageInfo

func (m *SessionResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *SessionResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *SessionResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 用户登录的请求
type UserLoginRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 用户登录的响应
type UserLoginResponse struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	SessionId            string   `protobuf:"bytes,3,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UserLoginResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserLoginResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 用户退出的请求
type DelSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelSessionRequest) Reset()         { *m = DelSessionRequest{} }
func (m *DelSessionRequest) String() string { return proto.CompactTextString(m) }
func (*DelSessionRequest) ProtoMessage()    {}
func (*DelSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{7}
}

func (m *DelSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelSessionRequest.Unmarshal(m, b)
}
func (m *DelSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelSessionRequest.Marshal(b, m, deterministic)
}
func (m *DelSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelSessionRequest.Merge(m, src)
}
func (m *DelSessionRequest) XXX_Size() int {
	return xxx_messageInfo_DelSessionRequest.Size(m)
}
func (m *DelSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelSessionRequest proto.InternalMessageInfo

func (m *DelSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 用户退出的响应
type DelSessionResponse struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelSessionResponse) Reset()         { *m = DelSessionResponse{} }
func (m *DelSessionResponse) String() string { return proto.CompactTextString(m) }
func (*DelSessionResponse) ProtoMessage()    {}
func (*DelSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{8}
}

func (m *DelSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelSessionResponse.Unmarshal(m, b)
}
func (m *DelSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelSessionResponse.Marshal(b, m, deterministic)
}
func (m *DelSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelSessionResponse.Merge(m, src)
}
func (m *DelSessionResponse) XXX_Size() int {
	return xxx_messageInfo_DelSessionResponse.Size(m)
}
func (m *DelSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelSessionResponse proto.InternalMessageInfo

func (m *DelSessionResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *DelSessionResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// 获取用户信息的请求
type UserInfoRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{9}
}

func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(m, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// 获取用户信息的响应
type UserInfoResponse struct {
	// "errno": "0",
	Errno string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	// "errmsg": "成功",
	ErrMsg string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	// 用户ID: 1,
	UserId string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// 用户名: "Panda",
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	// 用户手机号: "110",
	Mobile string `protobuf:"bytes,5,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	// 真实姓名: "熊猫",
	RealName string `protobuf:"bytes,6,opt,name=RealName,proto3" json:"RealName,omitempty"`
	// 用户身份证号码: "210112244556677",
	IdCard string `protobuf:"bytes,7,opt,name=IdCard,proto3" json:"IdCard,omitempty"`
	// 用户头像:"http://101.200.170.171:9998/group1/M00/00/00/Zciqq1n7It2ANn1dAADexS5wJKs808.png"
	AvatarUrl            string   `protobuf:"bytes,8,opt,name=AvatarUrl,proto3" json:"AvatarUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{10}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UserInfoResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserInfoResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfoResponse) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UserInfoResponse) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *UserInfoResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

// 上传用户头像的请求
type UserAvatarRequest struct {
	// 图片二进制
	Avatar []byte `protobuf:"bytes,1,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	// 用户Sessionid
	SessionId string `protobuf:"bytes,2,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	// 文件大小
	FileSize int64 `protobuf:"varint,3,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	// 文件名
	FileName             string   `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAvatarRequest) Reset()         { *m = UserAvatarRequest{} }
func (m *UserAvatarRequest) String() string { return proto.CompactTextString(m) }
func (*UserAvatarRequest) ProtoMessage()    {}
func (*UserAvatarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{11}
}

func (m *UserAvatarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAvatarRequest.Unmarshal(m, b)
}
func (m *UserAvatarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAvatarRequest.Marshal(b, m, deterministic)
}
func (m *UserAvatarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAvatarRequest.Merge(m, src)
}
func (m *UserAvatarRequest) XXX_Size() int {
	return xxx_messageInfo_UserAvatarRequest.Size(m)
}
func (m *UserAvatarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAvatarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserAvatarRequest proto.InternalMessageInfo

func (m *UserAvatarRequest) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *UserAvatarRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *UserAvatarRequest) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *UserAvatarRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

// 上传用户头像的响应
type UserAvatarResponse struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	// 返回的头像Url
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=AvatarUrl,proto3" json:"AvatarUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAvatarResponse) Reset()         { *m = UserAvatarResponse{} }
func (m *UserAvatarResponse) String() string { return proto.CompactTextString(m) }
func (*UserAvatarResponse) ProtoMessage()    {}
func (*UserAvatarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{12}
}

func (m *UserAvatarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAvatarResponse.Unmarshal(m, b)
}
func (m *UserAvatarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAvatarResponse.Marshal(b, m, deterministic)
}
func (m *UserAvatarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAvatarResponse.Merge(m, src)
}
func (m *UserAvatarResponse) XXX_Size() int {
	return xxx_messageInfo_UserAvatarResponse.Size(m)
}
func (m *UserAvatarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAvatarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserAvatarResponse proto.InternalMessageInfo

func (m *UserAvatarResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UserAvatarResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserAvatarResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

// 更新用户名的请求
type UpdateUserNameRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserNameRequest) Reset()         { *m = UpdateUserNameRequest{} }
func (m *UpdateUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserNameRequest) ProtoMessage()    {}
func (*UpdateUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{13}
}

func (m *UpdateUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserNameRequest.Unmarshal(m, b)
}
func (m *UpdateUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserNameRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserNameRequest.Merge(m, src)
}
func (m *UpdateUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserNameRequest.Size(m)
}
func (m *UpdateUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserNameRequest proto.InternalMessageInfo

func (m *UpdateUserNameRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *UpdateUserNameRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 更新用户名的响应
type UpdateUserNameResponse struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserNameResponse) Reset()         { *m = UpdateUserNameResponse{} }
func (m *UpdateUserNameResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserNameResponse) ProtoMessage()    {}
func (*UpdateUserNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{14}
}

func (m *UpdateUserNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserNameResponse.Unmarshal(m, b)
}
func (m *UpdateUserNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserNameResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserNameResponse.Merge(m, src)
}
func (m *UpdateUserNameResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserNameResponse.Size(m)
}
func (m *UpdateUserNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserNameResponse proto.InternalMessageInfo

func (m *UpdateUserNameResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UpdateUserNameResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UpdateUserNameResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 更新用户实名认证的请求
type UpdateUserAuthRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	// 真实姓名
	RealName string `protobuf:"bytes,2,opt,name=RealName,proto3" json:"RealName,omitempty"`
	// 身份证号
	IdCard               string   `protobuf:"bytes,3,opt,name=IdCard,proto3" json:"IdCard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserAuthRequest) Reset()         { *m = UpdateUserAuthRequest{} }
func (m *UpdateUserAuthRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserAuthRequest) ProtoMessage()    {}
func (*UpdateUserAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{15}
}

func (m *UpdateUserAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserAuthRequest.Unmarshal(m, b)
}
func (m *UpdateUserAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserAuthRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserAuthRequest.Merge(m, src)
}
func (m *UpdateUserAuthRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserAuthRequest.Size(m)
}
func (m *UpdateUserAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserAuthRequest proto.InternalMessageInfo

func (m *UpdateUserAuthRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *UpdateUserAuthRequest) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UpdateUserAuthRequest) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

// 更新用户实名认证的响应
type UpdateUserAuthResponse struct {
	Errno                string   `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserAuthResponse) Reset()         { *m = UpdateUserAuthResponse{} }
func (m *UpdateUserAuthResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserAuthResponse) ProtoMessage()    {}
func (*UpdateUserAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{16}
}

func (m *UpdateUserAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserAuthResponse.Unmarshal(m, b)
}
func (m *UpdateUserAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserAuthResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserAuthResponse.Merge(m, src)
}
func (m *UpdateUserAuthResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserAuthResponse.Size(m)
}
func (m *UpdateUserAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserAuthResponse proto.InternalMessageInfo

func (m *UpdateUserAuthResponse) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *UpdateUserAuthResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.UserInfoSrv.Message")
	proto.RegisterType((*UserRegistRequest)(nil), "go.micro.srv.UserInfoSrv.UserRegistRequest")
	proto.RegisterType((*UserRegistResponse)(nil), "go.micro.srv.UserInfoSrv.UserRegistResponse")
	proto.RegisterType((*SessionRequest)(nil), "go.micro.srv.UserInfoSrv.SessionRequest")
	proto.RegisterType((*SessionResponse)(nil), "go.micro.srv.UserInfoSrv.SessionResponse")
	proto.RegisterType((*UserLoginRequest)(nil), "go.micro.srv.UserInfoSrv.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "go.micro.srv.UserInfoSrv.UserLoginResponse")
	proto.RegisterType((*DelSessionRequest)(nil), "go.micro.srv.UserInfoSrv.DelSessionRequest")
	proto.RegisterType((*DelSessionResponse)(nil), "go.micro.srv.UserInfoSrv.DelSessionResponse")
	proto.RegisterType((*UserInfoRequest)(nil), "go.micro.srv.UserInfoSrv.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "go.micro.srv.UserInfoSrv.UserInfoResponse")
	proto.RegisterType((*UserAvatarRequest)(nil), "go.micro.srv.UserInfoSrv.UserAvatarRequest")
	proto.RegisterType((*UserAvatarResponse)(nil), "go.micro.srv.UserInfoSrv.UserAvatarResponse")
	proto.RegisterType((*UpdateUserNameRequest)(nil), "go.micro.srv.UserInfoSrv.UpdateUserNameRequest")
	proto.RegisterType((*UpdateUserNameResponse)(nil), "go.micro.srv.UserInfoSrv.UpdateUserNameResponse")
	proto.RegisterType((*UpdateUserAuthRequest)(nil), "go.micro.srv.UserInfoSrv.UpdateUserAuthRequest")
	proto.RegisterType((*UpdateUserAuthResponse)(nil), "go.micro.srv.UserInfoSrv.UpdateUserAuthResponse")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x49, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0xd2, 0x26, 0xed, 0xeb, 0x3e, 0x82, 0xca, 0x72, 0x39, 0xa0, 0x39, 0x75, 0x41, 0x2e,
	0xcb, 0x2f, 0x28, 0x6d, 0x53, 0x45, 0x22, 0x28, 0x38, 0xca, 0x89, 0x03, 0x38, 0xf5, 0xd4, 0x35,
	0xb2, 0x33, 0x61, 0xc6, 0x09, 0xcb, 0x8d, 0x3f, 0xc9, 0x5f, 0xe1, 0x8a, 0xec, 0x59, 0x32, 0x76,
	0x88, 0x6b, 0x82, 0x38, 0xc5, 0xdf, 0xf8, 0xbd, 0x79, 0xef, 0xfb, 0xde, 0xe2, 0xc0, 0xd1, 0x84,
	0xd1, 0x84, 0x9e, 0x93, 0xaf, 0x5e, 0x3c, 0x89, 0x88, 0xfa, 0x75, 0xb2, 0x53, 0x64, 0x05, 0xd4,
	0x89, 0xc3, 0x5b, 0x46, 0x1d, 0xce, 0x66, 0xce, 0x90, 0x13, 0xd6, 0x1d, 0xdf, 0xd1, 0x01, 0x9b,
	0xe1, 0x23, 0x68, 0xf7, 0x08, 0xe7, 0x5e, 0x40, 0xd0, 0x3e, 0x34, 0xb9, 0xf7, 0xcd, 0xaa, 0x3f,
	0xad, 0x1f, 0x6f, 0xba, 0xe9, 0x23, 0xf6, 0xe0, 0x20, 0xb5, 0x75, 0x49, 0x10, 0xf2, 0xc4, 0x25,
	0x9f, 0xa7, 0x84, 0x27, 0xe8, 0x10, 0x5a, 0x3d, 0x3a, 0x0a, 0x23, 0x22, 0x2d, 0x25, 0x42, 0x36,
	0x6c, 0xf4, 0x3d, 0xce, 0xbf, 0x50, 0xe6, 0x5b, 0x8d, 0xec, 0x8d, 0xc6, 0xc8, 0x82, 0xf6, 0x20,
	0xe6, 0x97, 0xd4, 0x27, 0x56, 0x33, 0x7b, 0xa5, 0x20, 0xfe, 0x08, 0xc8, 0x0c, 0xc1, 0x27, 0x74,
	0xcc, 0x09, 0x7a, 0x04, 0xeb, 0xd7, 0x8c, 0x8d, 0xa9, 0x0c, 0x21, 0x40, 0x1a, 0xf9, 0x9a, 0xb1,
	0x1e, 0x0f, 0xe4, 0xfd, 0x12, 0xa1, 0x27, 0xb0, 0x39, 0x20, 0x9c, 0x87, 0x74, 0xdc, 0xf5, 0xe5,
	0xfd, 0xf3, 0x03, 0xec, 0xc0, 0xae, 0x04, 0x8a, 0x41, 0xce, 0xbe, 0x5e, 0xb4, 0x7f, 0x0f, 0x7b,
	0xda, 0xbe, 0x42, 0x3a, 0x71, 0x2e, 0x9d, 0x98, 0x07, 0xa9, 0x10, 0x29, 0xa5, 0xb7, 0x5e, 0xac,
	0xd8, 0x6a, 0x8c, 0x3b, 0xb0, 0x9f, 0x3e, 0xbf, 0xa1, 0x41, 0x38, 0xfe, 0x07, 0x41, 0xf1, 0x07,
	0x51, 0x19, 0x79, 0xcf, 0x7f, 0x50, 0xed, 0x05, 0x1c, 0x5c, 0x91, 0xe8, 0xaf, 0x84, 0x7b, 0x0d,
	0xc8, 0x74, 0x59, 0x25, 0x29, 0x7c, 0x0e, 0x7b, 0xaa, 0x3b, 0xab, 0x05, 0xfd, 0x59, 0x17, 0x8a,
	0x0a, 0x8f, 0x95, 0x84, 0x38, 0x84, 0x56, 0x76, 0x83, 0x52, 0x41, 0x22, 0x84, 0x60, 0x2d, 0xab,
	0xe1, 0x5a, 0x76, 0x9a, 0x3d, 0x1b, 0xb5, 0x5a, 0x2f, 0xd6, 0xca, 0x25, 0x5e, 0x94, 0xd9, 0xb7,
	0x44, 0xad, 0x14, 0x4e, 0x7d, 0xba, 0xfe, 0xa5, 0xc7, 0x7c, 0xab, 0x2d, 0x7c, 0x04, 0x4a, 0x89,
	0x5d, 0xcc, 0xbc, 0xc4, 0x63, 0x43, 0x16, 0x59, 0x1b, 0x82, 0x98, 0x3e, 0xc0, 0x3f, 0xea, 0xa2,
	0xc4, 0xe2, 0xc4, 0xe8, 0x15, 0x71, 0x90, 0x51, 0xdb, 0x76, 0x25, 0xca, 0x8b, 0xd4, 0x28, 0x88,
	0x94, 0x66, 0xd7, 0x09, 0x23, 0x32, 0x08, 0xbf, 0x8b, 0x8e, 0x6c, 0xba, 0x1a, 0xab, 0x77, 0x06,
	0x53, 0x8d, 0xd5, 0x70, 0xaa, 0x14, 0x56, 0x6d, 0xb3, 0x39, 0xcb, 0x66, 0x91, 0xe5, 0x3b, 0x78,
	0x3c, 0x9c, 0xf8, 0x5e, 0x42, 0xd4, 0x84, 0x54, 0xaa, 0x7a, 0x6e, 0xc4, 0x1a, 0x85, 0x11, 0x1b,
	0xc1, 0x61, 0xf1, 0xca, 0x95, 0x12, 0x2f, 0x1b, 0xe3, 0xd0, 0x4c, 0xfb, 0x62, 0x9a, 0xdc, 0x57,
	0x4e, 0x5b, 0x77, 0x49, 0x63, 0x69, 0x97, 0x34, 0xcd, 0x2e, 0xc1, 0x1d, 0x93, 0x8e, 0x08, 0xb5,
	0x0a, 0x9d, 0x97, 0xbf, 0x5a, 0xd0, 0xbe, 0x16, 0x1f, 0x05, 0x14, 0xc3, 0x6e, 0x9f, 0xf2, 0x64,
	0xbe, 0x78, 0xd1, 0x99, 0xb3, 0xec, 0x0b, 0xe1, 0x2c, 0x7c, 0x01, 0xec, 0x67, 0xd5, 0x8c, 0x45,
	0x9a, 0xb8, 0x86, 0x6e, 0x01, 0x6e, 0x48, 0x22, 0x65, 0x40, 0xc7, 0xcb, 0xbd, 0xf3, 0xeb, 0xc6,
	0x3e, 0xa9, 0x60, 0xa9, 0x83, 0x7c, 0x82, 0x1d, 0xc5, 0x29, 0xdb, 0x8a, 0xe8, 0xb4, 0x3c, 0x4b,
	0x73, 0x05, 0xdb, 0x67, 0x95, 0x6c, 0x75, 0xac, 0x08, 0x76, 0xae, 0x48, 0x44, 0x12, 0xa2, 0x38,
	0x95, 0xf8, 0x2f, 0x6c, 0xd1, 0x32, 0xf9, 0x16, 0xf7, 0x27, 0xae, 0xa1, 0x3b, 0xd8, 0xba, 0x21,
	0x89, 0x32, 0x43, 0x27, 0xe5, 0xb9, 0x1a, 0xab, 0xd3, 0x3e, 0xad, 0x62, 0xaa, 0xe3, 0x18, 0x5d,
	0x21, 0xb7, 0xca, 0x03, 0xb2, 0xe4, 0x56, 0xd3, 0x43, 0x5d, 0x91, 0x5f, 0x22, 0xb8, 0x86, 0x18,
	0x6c, 0xf5, 0xa7, 0x73, 0x5a, 0xe7, 0x25, 0xee, 0x7f, 0xda, 0x10, 0xf6, 0xf3, 0xea, 0x0e, 0x3a,
	0x26, 0x87, 0x6d, 0x4d, 0x71, 0x9a, 0xdc, 0x57, 0x0b, 0x6a, 0xcc, 0x77, 0xb5, 0xa0, 0xe6, 0x94,
	0xe2, 0xda, 0xa8, 0x95, 0xfd, 0x07, 0x7b, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xc7, 0xdc,
	0x05, 0xa2, 0x09, 0x00, 0x00,
}
