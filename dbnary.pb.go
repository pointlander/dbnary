// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dbnary.proto

/*
Package dbnary is a generated protocol buffer package.

It is generated from these files:
	dbnary.proto

It has these top-level messages:
	Term
	Triple
	Entry
	Translation
	Compressed
*/
package dbnary

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

type Term_Type int32

const (
	Term_Blank   Term_Type = 0
	Term_IRI     Term_Type = 1
	Term_Literal Term_Type = 2
)

var Term_Type_name = map[int32]string{
	0: "Blank",
	1: "IRI",
	2: "Literal",
}
var Term_Type_value = map[string]int32{
	"Blank":   0,
	"IRI":     1,
	"Literal": 2,
}

func (x Term_Type) String() string {
	return proto.EnumName(Term_Type_name, int32(x))
}
func (Term_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Term struct {
	Type    Term_Type `protobuf:"varint,1,opt,name=type,enum=dbnary.Term_Type" json:"type,omitempty"`
	Prefix  uint64    `protobuf:"varint,2,opt,name=prefix" json:"prefix,omitempty"`
	Suffix  uint64    `protobuf:"varint,3,opt,name=suffix" json:"suffix,omitempty"`
	Key     string    `protobuf:"bytes,4,opt,name=key" json:"key,omitempty"`
	Literal string    `protobuf:"bytes,5,opt,name=literal" json:"literal,omitempty"`
}

func (m *Term) Reset()                    { *m = Term{} }
func (m *Term) String() string            { return proto.CompactTextString(m) }
func (*Term) ProtoMessage()               {}
func (*Term) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Term) GetType() Term_Type {
	if m != nil {
		return m.Type
	}
	return Term_Blank
}

func (m *Term) GetPrefix() uint64 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *Term) GetSuffix() uint64 {
	if m != nil {
		return m.Suffix
	}
	return 0
}

func (m *Term) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Term) GetLiteral() string {
	if m != nil {
		return m.Literal
	}
	return ""
}

type Triple struct {
	Predicate *Term `protobuf:"bytes,1,opt,name=predicate" json:"predicate,omitempty"`
	Object    *Term `protobuf:"bytes,2,opt,name=object" json:"object,omitempty"`
}

func (m *Triple) Reset()                    { *m = Triple{} }
func (m *Triple) String() string            { return proto.CompactTextString(m) }
func (*Triple) ProtoMessage()               {}
func (*Triple) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Triple) GetPredicate() *Term {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *Triple) GetObject() *Term {
	if m != nil {
		return m.Object
	}
	return nil
}

type Entry struct {
	Triples []*Triple `protobuf:"bytes,1,rep,name=triples" json:"triples,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetTriples() []*Triple {
	if m != nil {
		return m.Triples
	}
	return nil
}

type Translation struct {
	Keys []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *Translation) Reset()                    { *m = Translation{} }
func (m *Translation) String() string            { return proto.CompactTextString(m) }
func (*Translation) ProtoMessage()               {}
func (*Translation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Translation) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type Compressed struct {
	Size uint64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Compressed) Reset()                    { *m = Compressed{} }
func (m *Compressed) String() string            { return proto.CompactTextString(m) }
func (*Compressed) ProtoMessage()               {}
func (*Compressed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Compressed) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Compressed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Term)(nil), "dbnary.Term")
	proto.RegisterType((*Triple)(nil), "dbnary.Triple")
	proto.RegisterType((*Entry)(nil), "dbnary.Entry")
	proto.RegisterType((*Translation)(nil), "dbnary.Translation")
	proto.RegisterType((*Compressed)(nil), "dbnary.Compressed")
	proto.RegisterEnum("dbnary.Term_Type", Term_Type_name, Term_Type_value)
}

func init() { proto.RegisterFile("dbnary.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0xdd, 0x66, 0x9b, 0x90, 0x97, 0x52, 0xe2, 0x3b, 0x48, 0x8e, 0x31, 0x28, 0x06, 0x0f,
	0x05, 0xab, 0xbf, 0x40, 0xf1, 0x50, 0xf0, 0xb4, 0xe4, 0xe4, 0x6d, 0xdb, 0xbc, 0x42, 0x6c, 0x9a,
	0x2c, 0x9b, 0x15, 0x8c, 0x3f, 0xcb, 0x5f, 0x28, 0xbb, 0x9b, 0xaa, 0x78, 0x9b, 0xf9, 0x66, 0x96,
	0x37, 0xb0, 0xb0, 0xa8, 0xb7, 0x9d, 0xd4, 0xe3, 0x4a, 0xe9, 0xde, 0xf4, 0x18, 0x7a, 0x57, 0x7c,
	0x31, 0xe0, 0x15, 0xe9, 0x23, 0x5e, 0x03, 0x37, 0xa3, 0xa2, 0x8c, 0xe5, 0xac, 0x5c, 0xae, 0xcf,
	0x57, 0x53, 0xdb, 0x66, 0xab, 0x6a, 0x54, 0x24, 0x5c, 0x8c, 0x17, 0x10, 0x2a, 0x4d, 0xfb, 0xe6,
	0x23, 0x9b, 0xe5, 0xac, 0xe4, 0x62, 0x72, 0x96, 0x0f, 0xef, 0x7b, 0xcb, 0x03, 0xcf, 0xbd, 0xc3,
	0x14, 0x82, 0x03, 0x8d, 0x19, 0xcf, 0x59, 0x19, 0x0b, 0x2b, 0x31, 0x83, 0xa8, 0x6d, 0x0c, 0x69,
	0xd9, 0x66, 0x73, 0x47, 0x4f, 0xb6, 0xb8, 0x01, 0x6e, 0x2f, 0x61, 0x0c, 0xf3, 0xc7, 0x56, 0x76,
	0x87, 0xf4, 0x0c, 0x23, 0x08, 0x36, 0x62, 0x93, 0x32, 0x4c, 0x20, 0x7a, 0xf1, 0xb5, 0x74, 0x56,
	0xbc, 0x42, 0x58, 0xe9, 0x46, 0xb5, 0x84, 0xb7, 0x10, 0x2b, 0x4d, 0x75, 0xb3, 0x93, 0xc6, 0x4f,
	0x4f, 0xd6, 0x8b, 0xbf, 0xd3, 0xc5, 0x6f, 0x8c, 0x57, 0x10, 0xf6, 0xdb, 0x37, 0xda, 0x19, 0x37,
	0xfd, 0x7f, 0x71, 0xca, 0x8a, 0x3b, 0x98, 0x3f, 0x77, 0x46, 0x8f, 0x58, 0x42, 0x64, 0xdc, 0x91,
	0x21, 0x63, 0x79, 0x50, 0x26, 0xeb, 0xe5, 0x4f, 0xdf, 0x61, 0x71, 0x8a, 0x8b, 0x4b, 0x48, 0x2a,
	0x2d, 0xbb, 0xa1, 0x95, 0xa6, 0xe9, 0x3b, 0x44, 0xe0, 0x07, 0x1a, 0xfd, 0xab, 0x58, 0x38, 0x5d,
	0x3c, 0x00, 0x3c, 0xf5, 0x47, 0xa5, 0x69, 0x18, 0xa8, 0xb6, 0x8d, 0xa1, 0xf9, 0xf4, 0x83, 0xb9,
	0x70, 0xda, 0xb2, 0x5a, 0x1a, 0xe9, 0xb6, 0x2d, 0x84, 0xd3, 0xdb, 0xd0, 0xfd, 0xd5, 0xfd, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x41, 0x46, 0x0c, 0xbb, 0x01, 0x00, 0x00,
}