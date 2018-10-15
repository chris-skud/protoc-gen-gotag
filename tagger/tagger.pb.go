// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tagger/tagger.proto

package tagger // import "github.com/chris-skud/protoc-gen-gotag/tagger"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Tags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         847939,
	Name:          "tagger.tags",
	Tag:           "bytes,847939,opt,name=tags",
	Filename:      "tagger/tagger.proto",
}

var E_OneofTags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.OneofOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         847939,
	Name:          "tagger.oneof_tags",
	Tag:           "bytes,847939,opt,name=oneof_tags,json=oneofTags",
	Filename:      "tagger/tagger.proto",
}

func init() {
	proto.RegisterExtension(E_Tags)
	proto.RegisterExtension(E_OneofTags)
}

func init() { proto.RegisterFile("tagger/tagger.proto", fileDescriptor_tagger_405e48db8be7f212) }

var fileDescriptor_tagger_405e48db8be7f212 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x49, 0x4c, 0x4f,
	0x4f, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x42, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x34, 0xa9, 0x34, 0x4d, 0x3f, 0x25, 0xb5,
	0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0x1f, 0xaa, 0xd2, 0xca, 0x98, 0x8b, 0xa5, 0x24, 0x31, 0xbd,
	0x58, 0x48, 0x56, 0x0f, 0xa2, 0x54, 0x0f, 0xa6, 0x54, 0xcf, 0x2d, 0x33, 0x35, 0x27, 0xc5, 0xbf,
	0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xe2, 0xf0, 0x03, 0x63, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xb0,
	0x62, 0x2b, 0x3b, 0x2e, 0xae, 0xfc, 0xbc, 0xd4, 0xfc, 0xb4, 0x78, 0x1c, 0x5a, 0xfd, 0x41, 0x92,
	0xe8, 0x5a, 0x39, 0xc1, 0x5a, 0x42, 0x12, 0xd3, 0x8b, 0x9d, 0x4c, 0xa2, 0x8c, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0x8b, 0x32, 0xb3, 0x8b, 0x8a, 0xf3, 0x12,
	0x21, 0xae, 0x4c, 0xd6, 0x4d, 0x4f, 0xcd, 0xd3, 0x4d, 0xcf, 0x2f, 0x49, 0x4c, 0x87, 0xfa, 0xc9,
	0x1a, 0x42, 0x25, 0xb1, 0x81, 0xe5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x4b, 0x5c,
	0x8f, 0xf2, 0x00, 0x00, 0x00,
}
