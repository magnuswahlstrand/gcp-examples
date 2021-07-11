// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0, 0}
}

type Person struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber  `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamppb.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=main.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "main.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "main.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "main.AddressBook")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x9a, 0x06, 0x3b, 0x29, 0xa5, 0x0c, 0x52, 0x42, 0x2f, 0x86, 0xe2, 0x21, 0x28,
	0x6c, 0xa5, 0x3d, 0x7b, 0xb0, 0x50, 0x50, 0xb4, 0xb6, 0x2e, 0x15, 0x8f, 0x92, 0x92, 0xb1, 0x06,
	0xb3, 0xd9, 0x25, 0x9b, 0x1c, 0xfa, 0x5c, 0xbe, 0xa0, 0xec, 0x26, 0x95, 0x82, 0xb7, 0x7f, 0x67,
	0x3e, 0x66, 0xbe, 0x59, 0xe8, 0x09, 0xbd, 0x67, 0xaa, 0x94, 0x95, 0x44, 0x4f, 0x24, 0x59, 0x31,
	0xbe, 0xdc, 0x4b, 0xb9, 0xcf, 0x69, 0x6a, 0x6b, 0xbb, 0xfa, 0x73, 0x5a, 0x65, 0x82, 0x74, 0x95,
	0x08, 0xd5, 0x60, 0x93, 0x1f, 0x17, 0xfc, 0x0d, 0x95, 0x5a, 0x16, 0x88, 0xe0, 0x15, 0x89, 0xa0,
	0xd0, 0x89, 0x9c, 0xb8, 0xc7, 0x6d, 0xc6, 0x01, 0xb8, 0x59, 0x1a, 0xba, 0x91, 0x13, 0x77, 0xb9,
	0x9b, 0xa5, 0x78, 0x01, 0x5d, 0x12, 0x49, 0x96, 0x87, 0x1d, 0x0b, 0x35, 0x0f, 0xbc, 0x05, 0x5f,
	0x7d, 0xc9, 0x82, 0x74, 0xe8, 0x45, 0x9d, 0x38, 0x98, 0x85, 0xcc, 0x2c, 0x67, 0xcd, 0x5c, 0xb6,
	0x31, 0xad, 0x97, 0x5a, 0xec, 0xa8, 0xe4, 0x2d, 0x87, 0x77, 0xd0, 0xcf, 0x13, 0x5d, 0x7d, 0xd4,
	0x2a, 0x4d, 0x2a, 0x4a, 0xc3, 0x6e, 0xe4, 0xc4, 0xc1, 0x6c, 0xcc, 0x1a, 0x5d, 0x76, 0xd4, 0x65,
	0xdb, 0xa3, 0x2e, 0x0f, 0x0c, 0xff, 0xd6, 0xe0, 0xe3, 0x57, 0x08, 0x4e, 0xa6, 0xe2, 0x08, 0xfc,
	0xc2, 0xa6, 0xd6, 0xbd, 0x7d, 0xe1, 0x35, 0x78, 0xd5, 0x41, 0x91, 0xf5, 0x1f, 0xcc, 0x46, 0xff,
	0xad, 0xb6, 0x07, 0x45, 0xdc, 0x32, 0x93, 0x1b, 0xe8, 0xfd, 0x95, 0x10, 0xc0, 0x5f, 0xad, 0x17,
	0x8f, 0xcf, 0xcb, 0xe1, 0x19, 0x9e, 0x83, 0xf7, 0xb0, 0x5e, 0x2d, 0x87, 0x8e, 0x49, 0xef, 0x6b,
	0xfe, 0x34, 0x74, 0x27, 0x73, 0x08, 0xee, 0xd3, 0xb4, 0x24, 0xad, 0x17, 0x52, 0x7e, 0xe3, 0x15,
	0xf8, 0x8a, 0xa4, 0xca, 0xcd, 0xdf, 0x99, 0xfb, 0xfb, 0xa7, 0x9b, 0x78, 0xdb, 0xdb, 0xf9, 0xf6,
	0xaa, 0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0xe7, 0xfc, 0xcf, 0xa5, 0x01, 0x00, 0x00,
}