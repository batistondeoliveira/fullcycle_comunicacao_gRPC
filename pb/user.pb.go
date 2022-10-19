// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserResultStream struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResultStream) Reset()         { *m = UserResultStream{} }
func (m *UserResultStream) String() string { return proto.CompactTextString(m) }
func (*UserResultStream) ProtoMessage()    {}
func (*UserResultStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResultStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResultStream.Unmarshal(m, b)
}
func (m *UserResultStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResultStream.Marshal(b, m, deterministic)
}
func (m *UserResultStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResultStream.Merge(m, src)
}
func (m *UserResultStream) XXX_Size() int {
	return xxx_messageInfo_UserResultStream.Size(m)
}
func (m *UserResultStream) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResultStream.DiscardUnknown(m)
}

var xxx_messageInfo_UserResultStream proto.InternalMessageInfo

func (m *UserResultStream) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UserResultStream) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UserResultStream)(nil), "pb.UserResultStream")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4b, 0xc0, 0x30,
	0x14, 0x84, 0x69, 0x8c, 0x55, 0x5f, 0xa1, 0xc8, 0xa3, 0x48, 0x11, 0x05, 0xe9, 0xe4, 0x14, 0xa5,
	0xfe, 0x01, 0x75, 0x72, 0x6e, 0xd1, 0xc1, 0xc9, 0xc4, 0xbc, 0x21, 0xd0, 0x9a, 0x90, 0xa4, 0xfe,
	0x7e, 0xe9, 0xb3, 0x94, 0x6e, 0x77, 0x97, 0xcb, 0x97, 0x23, 0x00, 0x4b, 0xa2, 0xa8, 0x42, 0xf4,
	0xd9, 0xa3, 0x08, 0xa6, 0x7b, 0x06, 0xf9, 0x9e, 0x28, 0x62, 0x0d, 0xc2, 0xd9, 0xb6, 0xb8, 0x2b,
	0xee, 0x2f, 0x06, 0xe1, 0x2c, 0x22, 0xc8, 0x1f, 0x3d, 0x53, 0x2b, 0x38, 0x61, 0x8d, 0x0d, 0x9c,
	0xd2, 0xac, 0xdd, 0xd4, 0x9e, 0x70, 0xf8, 0x6f, 0xba, 0x37, 0xb8, 0x5c, 0x09, 0x03, 0xa5, 0x65,
	0xca, 0x63, 0x8e, 0xa4, 0x67, 0xbc, 0x82, 0x32, 0x65, 0x9d, 0x97, 0xb4, 0x11, 0x37, 0x87, 0x37,
	0x20, 0xd7, 0xf7, 0x99, 0x5a, 0xf5, 0xe7, 0x2a, 0x18, 0xc5, 0x77, 0x39, 0xed, 0xbf, 0xa0, 0x5a,
	0xdd, 0x48, 0xf1, 0xd7, 0x7d, 0x13, 0xde, 0xc2, 0xd9, 0x8b, 0xb5, 0xbc, 0x6e, 0x6f, 0x5e, 0xef,
	0x0a, 0x7b, 0xa8, 0xb7, 0xe3, 0x0f, 0x8a, 0xc6, 0x27, 0x7f, 0x68, 0x35, 0x3b, 0xf9, 0xb0, 0xea,
	0xb1, 0x78, 0x2d, 0x3f, 0xa5, 0x7a, 0x08, 0xc6, 0x94, 0xfc, 0x01, 0x4f, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x54, 0x37, 0x5b, 0x0e, 0x01, 0x00, 0x00,
}
