// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graphqlc/compiler/plugin.proto

package compiler

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	graphqlc "github.com/samlitowitz/graphqlc/pkg/graphqlc"
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

// The version number of protocol compiler
type Version struct {
	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix               string   `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_2febc61c83511dd1, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

// An encoded CodeGeneratorRequest is written to the plugins stdin.
type CodeGeneratorRequest struct {
	// The .graphql files that were explicitly listed on the command-line.
	// The code generator should generate code only for these files.
	// Each files descriptor will be included in graphql_file, defined below.
	FileToGenerate []string `protobuf:"bytes,1,rep,name=file_to_generate,json=fileToGenerate,proto3" json:"file_to_generate,omitempty"`
	Parameter      string   `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	// FileDescriptorGraphql for all files in files_to_generate.
	GraphqlFile []*graphqlc.FileDescriptorGraphql `protobuf:"bytes,15,rep,name=graphql_file,json=graphqlFile,proto3" json:"graphql_file,omitempty"`
	// The version number of graphql compiler
	CompilerVersion      *Version `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion,proto3" json:"compiler_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeGeneratorRequest) Reset()         { *m = CodeGeneratorRequest{} }
func (m *CodeGeneratorRequest) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorRequest) ProtoMessage()    {}
func (*CodeGeneratorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2febc61c83511dd1, []int{1}
}

func (m *CodeGeneratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorRequest.Unmarshal(m, b)
}
func (m *CodeGeneratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorRequest.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorRequest.Merge(m, src)
}
func (m *CodeGeneratorRequest) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorRequest.Size(m)
}
func (m *CodeGeneratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorRequest proto.InternalMessageInfo

func (m *CodeGeneratorRequest) GetFileToGenerate() []string {
	if m != nil {
		return m.FileToGenerate
	}
	return nil
}

func (m *CodeGeneratorRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

func (m *CodeGeneratorRequest) GetGraphqlFile() []*graphqlc.FileDescriptorGraphql {
	if m != nil {
		return m.GraphqlFile
	}
	return nil
}

func (m *CodeGeneratorRequest) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded CodeGeneratorResponse to stdout.
type CodeGeneratorResponse struct {
	// Error message. If non-empty, code generation failed. The plugin
	// process should exist with status code zero even if it reports an
	// error in this way.
	//
	// This should be used to indicate errors in .graphql files which
	// prevent the code generator from generating correct code. Errors
	// which indicate a problem in graphqlc itself -- such as the input
	// CodeGeneratorRequest being unparseable -- should be reported by
	// writing a message to stderr and exiting with a non-zero status code.
	Error                string                        `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	File                 []*CodeGeneratorResponse_File `protobuf:"bytes,15,rep,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CodeGeneratorResponse) Reset()         { *m = CodeGeneratorResponse{} }
func (m *CodeGeneratorResponse) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse) ProtoMessage()    {}
func (*CodeGeneratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2febc61c83511dd1, []int{2}
}

func (m *CodeGeneratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse.Merge(m, src)
}
func (m *CodeGeneratorResponse) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse.Size(m)
}
func (m *CodeGeneratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse proto.InternalMessageInfo

func (m *CodeGeneratorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CodeGeneratorResponse) GetFile() []*CodeGeneratorResponse_File {
	if m != nil {
		return m.File
	}
	return nil
}

// Represents a single generated file.
type CodeGeneratorResponse_File struct {
	// The file name relative to the output directory. The name must
	// not contain "." or ".." components and must be relative. The
	// file cannot lie outside the output directory. "/" must used as
	// a path separtor, not "\".
	//
	// If the name is ommited the content will be appended to the
	// previous file. This allows the generator to break large files
	// into small chunks, and allows the generated text to be streamed
	// back to graphqlc so that large files need not reside completely
	// in memory at one time.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If non-empty, indicates that the named file should already exist,
	// and the content here is to be inserted into that file at a
	// defined insertion point. This feature allows a code generator to
	// extend the output produced by another code generator. The original
	// generator may provide insertion points by placing special
	// annotations in the file that look like:
	//  @@graphqlc_insertion_point(NAME)
	// The annotation can have arbitrary text before and after it on
	// the line which allows it to be placed in a comment. NAME should
	// be replaced with an identifier naming the point -- that is what
	// other generators will use as the insertion_point. Code inserted
	// at this point will be placed immediately above the line
	// containing the insertion point. This allows multiple insertions
	// to the same point to be added in order. The double-@ is intended
	// to make it unlikely that the generated code could contain things
	// that look like insertion points by accident.
	//
	// Note that if the line containing the insertion point begins with
	// whitespace the same whitespace will be added to every line of the
	// inserted text. This is useful for langauges like Python where
	// indentation matters. In these languages the insertion point
	// comment should be indented the same amount as any inserted code
	// will need to be in order to work correctly in that context.
	//
	// The code generator that generates the initial file and the one
	// which inserts into it must both run as part of a single invocation
	// of graphqlc. Code generators are executed in the order in which
	// they appear on the command line.
	//
	// If |insertion_point| is present |name| must also be present.
	InsertionPoint string `protobuf:"bytes,2,opt,name=insertion_point,json=insertionPoint,proto3" json:"insertion_point,omitempty"`
	// The file contents.
	Content              string   `protobuf:"bytes,15,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeGeneratorResponse_File) Reset()         { *m = CodeGeneratorResponse_File{} }
func (m *CodeGeneratorResponse_File) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse_File) ProtoMessage()    {}
func (*CodeGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return fileDescriptor_2febc61c83511dd1, []int{2, 0}
}

func (m *CodeGeneratorResponse_File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeGeneratorResponse_File.Unmarshal(m, b)
}
func (m *CodeGeneratorResponse_File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeGeneratorResponse_File.Marshal(b, m, deterministic)
}
func (m *CodeGeneratorResponse_File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeGeneratorResponse_File.Merge(m, src)
}
func (m *CodeGeneratorResponse_File) XXX_Size() int {
	return xxx_messageInfo_CodeGeneratorResponse_File.Size(m)
}
func (m *CodeGeneratorResponse_File) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeGeneratorResponse_File.DiscardUnknown(m)
}

var xxx_messageInfo_CodeGeneratorResponse_File proto.InternalMessageInfo

func (m *CodeGeneratorResponse_File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetInsertionPoint() string {
	if m != nil {
		return m.InsertionPoint
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "graphqlc.compiler.Version")
	proto.RegisterType((*CodeGeneratorRequest)(nil), "graphqlc.compiler.CodeGeneratorRequest")
	proto.RegisterType((*CodeGeneratorResponse)(nil), "graphqlc.compiler.CodeGeneratorResponse")
	proto.RegisterType((*CodeGeneratorResponse_File)(nil), "graphqlc.compiler.CodeGeneratorResponse.File")
}

func init() { proto.RegisterFile("graphqlc/compiler/plugin.proto", fileDescriptor_2febc61c83511dd1) }

var fileDescriptor_2febc61c83511dd1 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x15, 0x9a, 0xb6, 0xe4, 0x16, 0x35, 0xc5, 0x2a, 0xc8, 0x54, 0x08, 0xa2, 0x6e, 0xc8,
	0x86, 0x54, 0x2a, 0x4f, 0x40, 0xf9, 0xe9, 0x16, 0x59, 0x88, 0x05, 0x12, 0x8a, 0x42, 0x7a, 0xdb,
	0x7a, 0x94, 0xd8, 0xae, 0xe3, 0x8e, 0xe6, 0x2d, 0xe7, 0x39, 0xe6, 0x2d, 0x46, 0xb1, 0x9d, 0xcc,
	0x68, 0xa6, 0xbb, 0x9c, 0xef, 0xf8, 0x3a, 0xf7, 0x1c, 0x19, 0x3e, 0x1c, 0x74, 0xa1, 0x8e, 0xa7,
	0xaa, 0x5c, 0x95, 0xb2, 0x56, 0xbc, 0x42, 0xbd, 0x52, 0xd5, 0xf9, 0xc0, 0x45, 0xa6, 0xb4, 0x34,
	0x92, 0xbc, 0xee, 0xfc, 0xac, 0xf3, 0x17, 0xef, 0xfa, 0x91, 0x1d, 0x36, 0xa5, 0xe6, 0xca, 0x48,
	0xed, 0x4e, 0x2f, 0x4b, 0x18, 0xff, 0x41, 0xdd, 0x70, 0x29, 0xc8, 0x1c, 0x86, 0x75, 0x71, 0x25,
	0x35, 0x0d, 0x92, 0x20, 0x1d, 0x32, 0x27, 0x2c, 0xe5, 0x42, 0x6a, 0xfa, 0xc2, 0xd3, 0x56, 0xb4,
	0x54, 0x15, 0xa6, 0x3c, 0xd2, 0x81, 0xa3, 0x56, 0x90, 0xb7, 0x30, 0x6a, 0xce, 0xfb, 0x3d, 0xbf,
	0xa1, 0x61, 0x12, 0xa4, 0x11, 0xf3, 0x6a, 0x79, 0x17, 0xc0, 0xfc, 0x9b, 0xdc, 0xe1, 0x16, 0x05,
	0xea, 0xc2, 0x48, 0xcd, 0xf0, 0x74, 0xc6, 0xc6, 0x90, 0x14, 0x66, 0x7b, 0x5e, 0x61, 0x6e, 0x64,
	0x7e, 0x70, 0x1e, 0xd2, 0x20, 0x19, 0xa4, 0x11, 0x9b, 0xb6, 0xfc, 0xb7, 0xf4, 0x13, 0x48, 0xde,
	0x43, 0xa4, 0x0a, 0x5d, 0xd4, 0x68, 0xd0, 0xad, 0x12, 0xb1, 0x07, 0x40, 0x36, 0xf0, 0xca, 0x47,
	0xcc, 0xdb, 0x39, 0x1a, 0x27, 0x83, 0x74, 0xb2, 0xfe, 0x98, 0xf5, 0x55, 0xfc, 0xe4, 0x15, 0x7e,
	0xef, 0xb3, 0x6f, 0x1d, 0x66, 0x13, 0xef, 0xb7, 0x2e, 0xf9, 0x01, 0xb3, 0xae, 0xb0, 0xfc, 0xda,
	0x55, 0x62, 0xd3, 0x4d, 0xd6, 0x8b, 0xec, 0x59, 0xa5, 0x99, 0x2f, 0x8d, 0xc5, 0x1d, 0xf1, 0x60,
	0x79, 0x1b, 0xc0, 0x9b, 0x27, 0x59, 0x1b, 0x25, 0x45, 0x83, 0x6d, 0x67, 0xa8, 0xb5, 0xef, 0x37,
	0x62, 0x4e, 0x90, 0xaf, 0x10, 0x3e, 0x5a, 0xf9, 0xf3, 0x85, 0x5f, 0x5d, 0xbc, 0xcd, 0x26, 0x62,
	0x76, 0x74, 0xf1, 0x0f, 0x42, 0x9b, 0x80, 0x40, 0x28, 0x8a, 0x1a, 0xfd, 0xfd, 0xf6, 0x9b, 0x7c,
	0x82, 0x98, 0x8b, 0x06, 0xb5, 0xe1, 0x52, 0xe4, 0x4a, 0x72, 0x61, 0x7c, 0x7b, 0xd3, 0x1e, 0xff,
	0x6a, 0x29, 0xa1, 0x30, 0x2e, 0xa5, 0x30, 0x28, 0x0c, 0x8d, 0xed, 0x81, 0x4e, 0x6e, 0xe0, 0xef,
	0xcb, 0x6e, 0x97, 0xff, 0x23, 0xfb, 0x6a, 0xbe, 0xdc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x44,
	0x9b, 0x5d, 0x85, 0x02, 0x00, 0x00,
}
