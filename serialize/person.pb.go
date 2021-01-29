// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package serialize

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// 人
type PersonProto3 struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64                `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  int64                `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Like                 []string             `protobuf:"bytes,4,rep,name=like,proto3" json:"like,omitempty"`
	Children             []*PersonProto3      `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	Address              []string             `protobuf:"bytes,6,rep,name=address,proto3" json:"address,omitempty"`
	Phone                []string             `protobuf:"bytes,7,rep,name=phone,proto3" json:"phone,omitempty"`
	Card                 string               `protobuf:"bytes,8,opt,name=card,proto3" json:"card,omitempty"`
	Qq                   string               `protobuf:"bytes,9,opt,name=qq,proto3" json:"qq,omitempty"`
	WeChat               string               `protobuf:"bytes,10,opt,name=weChat,proto3" json:"weChat,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,11,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Money                float64              `protobuf:"fixed64,12,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PersonProto3) Reset()         { *m = PersonProto3{} }
func (m *PersonProto3) String() string { return proto.CompactTextString(m) }
func (*PersonProto3) ProtoMessage()    {}
func (*PersonProto3) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *PersonProto3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonProto3.Unmarshal(m, b)
}
func (m *PersonProto3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonProto3.Marshal(b, m, deterministic)
}
func (m *PersonProto3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonProto3.Merge(m, src)
}
func (m *PersonProto3) XXX_Size() int {
	return xxx_messageInfo_PersonProto3.Size(m)
}
func (m *PersonProto3) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonProto3.DiscardUnknown(m)
}

var xxx_messageInfo_PersonProto3 proto.InternalMessageInfo

func (m *PersonProto3) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonProto3) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PersonProto3) GetSex() int64 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *PersonProto3) GetLike() []string {
	if m != nil {
		return m.Like
	}
	return nil
}

func (m *PersonProto3) GetChildren() []*PersonProto3 {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *PersonProto3) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PersonProto3) GetPhone() []string {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *PersonProto3) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *PersonProto3) GetQq() string {
	if m != nil {
		return m.Qq
	}
	return ""
}

func (m *PersonProto3) GetWeChat() string {
	if m != nil {
		return m.WeChat
	}
	return ""
}

func (m *PersonProto3) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *PersonProto3) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*PersonProto3)(nil), "serialize.PersonProto3")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xdd, 0x6a, 0xb4, 0x30,
	0x14, 0x24, 0xba, 0xbf, 0x51, 0x3e, 0x3e, 0x42, 0x69, 0x0f, 0x7b, 0xd3, 0xd0, 0x2b, 0xaf, 0xb2,
	0xb0, 0x42, 0x5f, 0xa0, 0x2f, 0xb0, 0x48, 0x5f, 0x20, 0xae, 0xa7, 0x1a, 0xaa, 0x46, 0x13, 0x4b,
	0x77, 0xfb, 0x86, 0x7d, 0xab, 0x92, 0xa4, 0x4a, 0xef, 0x66, 0x26, 0x33, 0x39, 0xc3, 0xd0, 0x74,
	0x40, 0x63, 0x75, 0x2f, 0x06, 0xa3, 0x27, 0xcd, 0xf6, 0x16, 0x8d, 0x92, 0xad, 0xfa, 0xc2, 0xc3,
	0x63, 0xad, 0x75, 0xdd, 0xe2, 0xd1, 0x3f, 0x94, 0x1f, 0x6f, 0xc7, 0x49, 0x75, 0x68, 0x27, 0xd9,
	0x0d, 0xc1, 0xfb, 0xf4, 0x1d, 0xd1, 0xf4, 0xec, 0xc3, 0x67, 0xc7, 0x73, 0xc6, 0xe8, 0xaa, 0x97,
	0x1d, 0x02, 0xe1, 0x24, 0xdb, 0x17, 0x1e, 0xb3, 0xff, 0x34, 0x96, 0x35, 0x42, 0xc4, 0x49, 0x16,
	0x17, 0x0e, 0x3a, 0xc5, 0xe2, 0x15, 0xe2, 0xa0, 0x58, 0xbc, 0xba, 0x5c, 0xab, 0xde, 0x11, 0x56,
	0x3c, 0x76, 0x39, 0x87, 0x59, 0x4e, 0x77, 0x97, 0x46, 0xb5, 0x95, 0xc1, 0x1e, 0xd6, 0x3c, 0xce,
	0x92, 0xd3, 0x83, 0x58, 0xba, 0x89, 0xbf, 0x67, 0x8b, 0xc5, 0xc8, 0x80, 0x6e, 0x65, 0x55, 0x19,
	0xb4, 0x16, 0x36, 0xfe, 0xaf, 0x99, 0xb2, 0x3b, 0xba, 0x1e, 0x1a, 0xdd, 0x23, 0x6c, 0xbd, 0x1e,
	0x88, 0x3b, 0x7c, 0x91, 0xa6, 0x82, 0x5d, 0x28, 0xec, 0x30, 0xfb, 0x47, 0xa3, 0x71, 0x84, 0xbd,
	0x57, 0xa2, 0x71, 0x64, 0xf7, 0x74, 0xf3, 0x89, 0x2f, 0x8d, 0x9c, 0x80, 0x7a, 0xed, 0x97, 0xb1,
	0x67, 0xba, 0x2b, 0x95, 0x99, 0x9a, 0x4a, 0xde, 0x20, 0xe1, 0x24, 0x4b, 0x4e, 0x07, 0x11, 0x16,
	0x13, 0xf3, 0x62, 0xe2, 0x75, 0x5e, 0xac, 0x58, 0xbc, 0xae, 0x49, 0xa7, 0x7b, 0xbc, 0x41, 0xca,
	0x49, 0x46, 0x8a, 0x40, 0xca, 0x8d, 0xcf, 0xe4, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x3e,
	0x8c, 0xeb, 0x8e, 0x01, 0x00, 0x00,
}