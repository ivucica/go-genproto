// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1alpha1/cel_service.proto

package expr

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Severities of issues.
type IssueDetails_Severity int32

const (
	// An unspecified severity.
	IssueDetails_SEVERITY_UNSPECIFIED IssueDetails_Severity = 0
	// Deprecation issue for statements and method that may no longer be
	// supported or maintained.
	IssueDetails_DEPRECATION IssueDetails_Severity = 1
	// Warnings such as: unused variables.
	IssueDetails_WARNING IssueDetails_Severity = 2
	// Errors such as: unmatched curly braces or variable redefinition.
	IssueDetails_ERROR IssueDetails_Severity = 3
)

var IssueDetails_Severity_name = map[int32]string{
	0: "SEVERITY_UNSPECIFIED",
	1: "DEPRECATION",
	2: "WARNING",
	3: "ERROR",
}

var IssueDetails_Severity_value = map[string]int32{
	"SEVERITY_UNSPECIFIED": 0,
	"DEPRECATION":          1,
	"WARNING":              2,
	"ERROR":                3,
}

func (x IssueDetails_Severity) String() string {
	return proto.EnumName(IssueDetails_Severity_name, int32(x))
}

func (IssueDetails_Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{6, 0}
}

// Request message for the Parse method.
type ParseRequest struct {
	// Required. Source text in CEL syntax.
	CelSource string `protobuf:"bytes,1,opt,name=cel_source,json=celSource,proto3" json:"cel_source,omitempty"`
	// Tag for version of CEL syntax, for future use.
	SyntaxVersion string `protobuf:"bytes,2,opt,name=syntax_version,json=syntaxVersion,proto3" json:"syntax_version,omitempty"`
	// File or resource for source text, used in [SourceInfo][google.api.expr.v1alpha1.SourceInfo].
	SourceLocation string `protobuf:"bytes,3,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	// Prevent macro expansion.  See "Macros" in Language Defiinition.
	DisableMacros        bool     `protobuf:"varint,4,opt,name=disable_macros,json=disableMacros,proto3" json:"disable_macros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseRequest) Reset()         { *m = ParseRequest{} }
func (m *ParseRequest) String() string { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()    {}
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{0}
}

func (m *ParseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseRequest.Unmarshal(m, b)
}
func (m *ParseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseRequest.Marshal(b, m, deterministic)
}
func (m *ParseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseRequest.Merge(m, src)
}
func (m *ParseRequest) XXX_Size() int {
	return xxx_messageInfo_ParseRequest.Size(m)
}
func (m *ParseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseRequest proto.InternalMessageInfo

func (m *ParseRequest) GetCelSource() string {
	if m != nil {
		return m.CelSource
	}
	return ""
}

func (m *ParseRequest) GetSyntaxVersion() string {
	if m != nil {
		return m.SyntaxVersion
	}
	return ""
}

func (m *ParseRequest) GetSourceLocation() string {
	if m != nil {
		return m.SourceLocation
	}
	return ""
}

func (m *ParseRequest) GetDisableMacros() bool {
	if m != nil {
		return m.DisableMacros
	}
	return false
}

// Response message for the Parse method.
type ParseResponse struct {
	// The parsed representation, or unset if parsing failed.
	ParsedExpr *ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
	// Any number of issues with [StatusDetails][] as the details.
	Issues               []*status.Status `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ParseResponse) Reset()         { *m = ParseResponse{} }
func (m *ParseResponse) String() string { return proto.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()    {}
func (*ParseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{1}
}

func (m *ParseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResponse.Unmarshal(m, b)
}
func (m *ParseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResponse.Marshal(b, m, deterministic)
}
func (m *ParseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResponse.Merge(m, src)
}
func (m *ParseResponse) XXX_Size() int {
	return xxx_messageInfo_ParseResponse.Size(m)
}
func (m *ParseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResponse proto.InternalMessageInfo

func (m *ParseResponse) GetParsedExpr() *ParsedExpr {
	if m != nil {
		return m.ParsedExpr
	}
	return nil
}

func (m *ParseResponse) GetIssues() []*status.Status {
	if m != nil {
		return m.Issues
	}
	return nil
}

// Request message for the Check method.
type CheckRequest struct {
	// Required. The parsed representation of the CEL program.
	ParsedExpr *ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
	// Declarations of types for external variables and functions.
	// Required if program uses external variables or functions
	// not in the default environment.
	TypeEnv []*Decl `protobuf:"bytes,2,rep,name=type_env,json=typeEnv,proto3" json:"type_env,omitempty"`
	// The protocol buffer context.  See "Name Resolution" in the
	// Language Definition.
	Container string `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
	// If true, use only the declarations in [type_env][google.api.expr.v1alpha1.CheckRequest.type_env].  If false (default),
	// add declarations for the standard definitions to the type environment.  See
	// "Standard Definitions" in the Language Definition.
	NoStdEnv             bool     `protobuf:"varint,4,opt,name=no_std_env,json=noStdEnv,proto3" json:"no_std_env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{2}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetParsedExpr() *ParsedExpr {
	if m != nil {
		return m.ParsedExpr
	}
	return nil
}

func (m *CheckRequest) GetTypeEnv() []*Decl {
	if m != nil {
		return m.TypeEnv
	}
	return nil
}

func (m *CheckRequest) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *CheckRequest) GetNoStdEnv() bool {
	if m != nil {
		return m.NoStdEnv
	}
	return false
}

// Response message for the Check method.
type CheckResponse struct {
	// The annotated representation, or unset if checking failed.
	CheckedExpr *CheckedExpr `protobuf:"bytes,1,opt,name=checked_expr,json=checkedExpr,proto3" json:"checked_expr,omitempty"`
	// Any number of issues with [StatusDetails][] as the details.
	Issues               []*status.Status `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{3}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetCheckedExpr() *CheckedExpr {
	if m != nil {
		return m.CheckedExpr
	}
	return nil
}

func (m *CheckResponse) GetIssues() []*status.Status {
	if m != nil {
		return m.Issues
	}
	return nil
}

// Request message for the Eval method.
type EvalRequest struct {
	// Required. Either the parsed or annotated representation of the CEL program.
	//
	// Types that are valid to be assigned to ExprKind:
	//	*EvalRequest_ParsedExpr
	//	*EvalRequest_CheckedExpr
	ExprKind isEvalRequest_ExprKind `protobuf_oneof:"expr_kind"`
	// Bindings for the external variables.  The types SHOULD be compatible
	// with the type environment in [CheckRequest][google.api.expr.v1alpha1.CheckRequest], if checked.
	Bindings map[string]*ExprValue `protobuf:"bytes,3,rep,name=bindings,proto3" json:"bindings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// SHOULD be the same container as used in [CheckRequest][google.api.expr.v1alpha1.CheckRequest], if checked.
	Container            string   `protobuf:"bytes,4,opt,name=container,proto3" json:"container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvalRequest) Reset()         { *m = EvalRequest{} }
func (m *EvalRequest) String() string { return proto.CompactTextString(m) }
func (*EvalRequest) ProtoMessage()    {}
func (*EvalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{4}
}

func (m *EvalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalRequest.Unmarshal(m, b)
}
func (m *EvalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalRequest.Marshal(b, m, deterministic)
}
func (m *EvalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalRequest.Merge(m, src)
}
func (m *EvalRequest) XXX_Size() int {
	return xxx_messageInfo_EvalRequest.Size(m)
}
func (m *EvalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvalRequest proto.InternalMessageInfo

type isEvalRequest_ExprKind interface {
	isEvalRequest_ExprKind()
}

type EvalRequest_ParsedExpr struct {
	ParsedExpr *ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3,oneof"`
}

type EvalRequest_CheckedExpr struct {
	CheckedExpr *CheckedExpr `protobuf:"bytes,2,opt,name=checked_expr,json=checkedExpr,proto3,oneof"`
}

func (*EvalRequest_ParsedExpr) isEvalRequest_ExprKind() {}

func (*EvalRequest_CheckedExpr) isEvalRequest_ExprKind() {}

func (m *EvalRequest) GetExprKind() isEvalRequest_ExprKind {
	if m != nil {
		return m.ExprKind
	}
	return nil
}

func (m *EvalRequest) GetParsedExpr() *ParsedExpr {
	if x, ok := m.GetExprKind().(*EvalRequest_ParsedExpr); ok {
		return x.ParsedExpr
	}
	return nil
}

func (m *EvalRequest) GetCheckedExpr() *CheckedExpr {
	if x, ok := m.GetExprKind().(*EvalRequest_CheckedExpr); ok {
		return x.CheckedExpr
	}
	return nil
}

func (m *EvalRequest) GetBindings() map[string]*ExprValue {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *EvalRequest) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EvalRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EvalRequest_OneofMarshaler, _EvalRequest_OneofUnmarshaler, _EvalRequest_OneofSizer, []interface{}{
		(*EvalRequest_ParsedExpr)(nil),
		(*EvalRequest_CheckedExpr)(nil),
	}
}

func _EvalRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EvalRequest)
	// expr_kind
	switch x := m.ExprKind.(type) {
	case *EvalRequest_ParsedExpr:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ParsedExpr); err != nil {
			return err
		}
	case *EvalRequest_CheckedExpr:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CheckedExpr); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EvalRequest.ExprKind has unexpected type %T", x)
	}
	return nil
}

func _EvalRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EvalRequest)
	switch tag {
	case 1: // expr_kind.parsed_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ParsedExpr)
		err := b.DecodeMessage(msg)
		m.ExprKind = &EvalRequest_ParsedExpr{msg}
		return true, err
	case 2: // expr_kind.checked_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CheckedExpr)
		err := b.DecodeMessage(msg)
		m.ExprKind = &EvalRequest_CheckedExpr{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EvalRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EvalRequest)
	// expr_kind
	switch x := m.ExprKind.(type) {
	case *EvalRequest_ParsedExpr:
		s := proto.Size(x.ParsedExpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EvalRequest_CheckedExpr:
		s := proto.Size(x.CheckedExpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for the Eval method.
type EvalResponse struct {
	// The execution result, or unset if execution couldn't start.
	Result *ExprValue `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// Any number of issues with [StatusDetails][] as the details.
	// Note that CEL execution errors are reified into [ExprValue][google.api.expr.v1alpha1.ExprValue].
	// Nevertheless, we'll allow out-of-band issues to be raised,
	// which also makes the replies more regular.
	Issues               []*status.Status `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EvalResponse) Reset()         { *m = EvalResponse{} }
func (m *EvalResponse) String() string { return proto.CompactTextString(m) }
func (*EvalResponse) ProtoMessage()    {}
func (*EvalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{5}
}

func (m *EvalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalResponse.Unmarshal(m, b)
}
func (m *EvalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalResponse.Marshal(b, m, deterministic)
}
func (m *EvalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalResponse.Merge(m, src)
}
func (m *EvalResponse) XXX_Size() int {
	return xxx_messageInfo_EvalResponse.Size(m)
}
func (m *EvalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvalResponse proto.InternalMessageInfo

func (m *EvalResponse) GetResult() *ExprValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EvalResponse) GetIssues() []*status.Status {
	if m != nil {
		return m.Issues
	}
	return nil
}

// Warnings or errors in service execution are represented by
// [google.rpc.Status][google.rpc.Status] messages, with the following message
// in the details field.
type IssueDetails struct {
	// The severity of the issue.
	Severity IssueDetails_Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=google.api.expr.v1alpha1.IssueDetails_Severity" json:"severity,omitempty"`
	// Position in the source, if known.
	Position *SourcePosition `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	// Expression ID from [Expr][google.api.expr.v1alpha1.Expr], 0 if unknown.
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueDetails) Reset()         { *m = IssueDetails{} }
func (m *IssueDetails) String() string { return proto.CompactTextString(m) }
func (*IssueDetails) ProtoMessage()    {}
func (*IssueDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35b2125e64b6d66, []int{6}
}

func (m *IssueDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueDetails.Unmarshal(m, b)
}
func (m *IssueDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueDetails.Marshal(b, m, deterministic)
}
func (m *IssueDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueDetails.Merge(m, src)
}
func (m *IssueDetails) XXX_Size() int {
	return xxx_messageInfo_IssueDetails.Size(m)
}
func (m *IssueDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueDetails.DiscardUnknown(m)
}

var xxx_messageInfo_IssueDetails proto.InternalMessageInfo

func (m *IssueDetails) GetSeverity() IssueDetails_Severity {
	if m != nil {
		return m.Severity
	}
	return IssueDetails_SEVERITY_UNSPECIFIED
}

func (m *IssueDetails) GetPosition() *SourcePosition {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *IssueDetails) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.api.expr.v1alpha1.IssueDetails_Severity", IssueDetails_Severity_name, IssueDetails_Severity_value)
	proto.RegisterType((*ParseRequest)(nil), "google.api.expr.v1alpha1.ParseRequest")
	proto.RegisterType((*ParseResponse)(nil), "google.api.expr.v1alpha1.ParseResponse")
	proto.RegisterType((*CheckRequest)(nil), "google.api.expr.v1alpha1.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "google.api.expr.v1alpha1.CheckResponse")
	proto.RegisterType((*EvalRequest)(nil), "google.api.expr.v1alpha1.EvalRequest")
	proto.RegisterMapType((map[string]*ExprValue)(nil), "google.api.expr.v1alpha1.EvalRequest.BindingsEntry")
	proto.RegisterType((*EvalResponse)(nil), "google.api.expr.v1alpha1.EvalResponse")
	proto.RegisterType((*IssueDetails)(nil), "google.api.expr.v1alpha1.IssueDetails")
}

func init() {
	proto.RegisterFile("google/api/expr/v1alpha1/cel_service.proto", fileDescriptor_f35b2125e64b6d66)
}

var fileDescriptor_f35b2125e64b6d66 = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x41, 0x6f, 0xdb, 0x36,
	0x14, 0xc7, 0x2d, 0x39, 0x49, 0xed, 0x27, 0x3b, 0x31, 0x88, 0x01, 0x33, 0x8c, 0x6c, 0x08, 0xd4,
	0x25, 0x0d, 0x7a, 0x90, 0x50, 0xf7, 0xb2, 0xae, 0xbb, 0x34, 0xb6, 0xd6, 0x78, 0x5b, 0x13, 0x83,
	0xee, 0x52, 0xac, 0x17, 0x8d, 0x91, 0x08, 0x97, 0x88, 0x4a, 0x6a, 0xa4, 0xac, 0x25, 0xb7, 0x61,
	0xc0, 0x3e, 0xc6, 0x6e, 0xfb, 0x38, 0xfb, 0x40, 0x3b, 0x0e, 0x22, 0xe9, 0xc4, 0xee, 0xa0, 0xa4,
	0x19, 0x76, 0x93, 0x9e, 0x7e, 0xef, 0xaf, 0xf7, 0xfe, 0x7c, 0x24, 0xe1, 0xf1, 0x5c, 0x88, 0x79,
	0x46, 0x43, 0x92, 0xb3, 0x90, 0x5e, 0xe6, 0x32, 0x2c, 0x9f, 0x90, 0x2c, 0x7f, 0x47, 0x9e, 0x84,
	0x09, 0xcd, 0x62, 0x45, 0x65, 0xc9, 0x12, 0x1a, 0xe4, 0x52, 0x14, 0x02, 0xf5, 0x0d, 0x1b, 0x90,
	0x9c, 0x05, 0x15, 0x1b, 0x2c, 0xd9, 0xc1, 0x41, 0xbd, 0xca, 0x3b, 0x9a, 0x5c, 0xd0, 0xd4, 0x28,
	0x0c, 0x1e, 0xd6, 0x72, 0xb4, 0x24, 0x99, 0x85, 0xf6, 0x6b, 0x21, 0x75, 0xc5, 0x0b, 0x72, 0x69,
	0xb1, 0x4f, 0x2d, 0x26, 0xf3, 0x24, 0x54, 0x05, 0x29, 0x16, 0xca, 0x7c, 0xf0, 0xff, 0x74, 0xa0,
	0x33, 0x25, 0x52, 0x51, 0x4c, 0x7f, 0x5e, 0x50, 0x55, 0xa0, 0xcf, 0x00, 0x74, 0x33, 0x62, 0x21,
	0x13, 0xda, 0x77, 0xf6, 0x9c, 0xc3, 0x36, 0x6e, 0x27, 0x34, 0x9b, 0xe9, 0x00, 0xda, 0x87, 0x6d,
	0x23, 0x1c, 0x97, 0x54, 0x2a, 0x26, 0x78, 0xdf, 0xd5, 0x48, 0xd7, 0x44, 0xcf, 0x4c, 0x10, 0x3d,
	0x82, 0x1d, 0xa3, 0x10, 0x67, 0x22, 0x21, 0x45, 0xc5, 0x35, 0x35, 0xb7, 0x6d, 0xc2, 0xdf, 0xdb,
	0x68, 0xa5, 0x97, 0x32, 0x45, 0xce, 0x33, 0x1a, 0xbf, 0x27, 0x89, 0x14, 0xaa, 0xbf, 0xb1, 0xe7,
	0x1c, 0xb6, 0x70, 0xd7, 0x46, 0x5f, 0xe9, 0xa0, 0xff, 0x9b, 0x03, 0x5d, 0x5b, 0xa6, 0xca, 0x05,
	0x57, 0x14, 0x45, 0xe0, 0xe5, 0x55, 0x20, 0x8d, 0xab, 0xb6, 0x75, 0xa1, 0xde, 0xf0, 0x8b, 0xa0,
	0xce, 0xf5, 0x40, 0x67, 0xa7, 0xd1, 0x65, 0x2e, 0x31, 0xe4, 0xd7, 0xcf, 0xe8, 0x31, 0x6c, 0x31,
	0xa5, 0x16, 0x54, 0xf5, 0xdd, 0xbd, 0xe6, 0xa1, 0x37, 0x44, 0x4b, 0x05, 0x99, 0x27, 0xc1, 0x4c,
	0x3b, 0x85, 0x2d, 0xe1, 0xff, 0xe5, 0x40, 0x67, 0x54, 0x2d, 0xd1, 0xd2, 0xab, 0xff, 0xa9, 0x86,
	0x67, 0xd0, 0x2a, 0xae, 0x72, 0x1a, 0x53, 0x5e, 0xda, 0x2a, 0x3e, 0xaf, 0xd7, 0x18, 0xd3, 0x24,
	0xc3, 0x0f, 0x2a, 0x3e, 0xe2, 0x25, 0xda, 0x85, 0x76, 0x22, 0x78, 0x41, 0x18, 0xa7, 0xd2, 0x3a,
	0x7c, 0x13, 0x40, 0xbb, 0x00, 0x5c, 0xc4, 0xaa, 0x48, 0xb5, 0xb4, 0x31, 0xb6, 0xc5, 0xc5, 0xac,
	0x48, 0x23, 0x5e, 0xfa, 0xbf, 0x3b, 0xd0, 0xb5, 0xed, 0x58, 0x4f, 0x8f, 0xa1, 0x63, 0x47, 0x70,
	0xb5, 0xa1, 0xfd, 0xfa, 0x62, 0x46, 0x86, 0xd6, 0x1d, 0x79, 0xc9, 0xcd, 0xcb, 0xbd, 0x6c, 0xfd,
	0xb5, 0x09, 0x5e, 0x54, 0x92, 0x6c, 0xe9, 0xea, 0xcb, 0xff, 0xec, 0xea, 0x71, 0x63, 0xcd, 0xd7,
	0x6f, 0x3f, 0x68, 0xc7, 0xbd, 0x47, 0x3b, 0xc7, 0x8d, 0xf5, 0x86, 0x4e, 0xa1, 0x75, 0xce, 0x78,
	0xca, 0xf8, 0x5c, 0xf5, 0x9b, 0xba, 0xa5, 0xa7, 0xf5, 0x3a, 0x2b, 0xdd, 0x04, 0x47, 0x36, 0x2b,
	0xe2, 0x85, 0xbc, 0xc2, 0xd7, 0x22, 0xeb, 0x2b, 0xb7, 0xf1, 0xc1, 0xca, 0x0d, 0x7e, 0x82, 0xee,
	0x5a, 0x22, 0xea, 0x41, 0xf3, 0x82, 0x5e, 0xd9, 0xfd, 0x58, 0x3d, 0xa2, 0x67, 0xb0, 0x59, 0x92,
	0x6c, 0x41, 0x6d, 0x5b, 0x0f, 0x6f, 0x29, 0xe7, 0x32, 0x97, 0x67, 0x15, 0x8a, 0x4d, 0xc6, 0x57,
	0xee, 0x97, 0xce, 0x91, 0x07, 0xed, 0x8a, 0x8a, 0x2f, 0x18, 0x4f, 0xfd, 0x5f, 0xa0, 0x63, 0x6a,
	0xb6, 0x83, 0xf0, 0x1c, 0xb6, 0x24, 0x55, 0x8b, 0xac, 0xb0, 0xee, 0x7f, 0x94, 0xb8, 0x4d, 0xb9,
	0xdf, 0xda, 0xbb, 0xd0, 0x99, 0x54, 0x8f, 0x63, 0x5a, 0x10, 0x96, 0x29, 0xf4, 0x1d, 0xb4, 0x14,
	0x2d, 0xa9, 0x64, 0x85, 0x69, 0x76, 0x7b, 0x18, 0xd6, 0xff, 0x7b, 0x35, 0x33, 0x98, 0xd9, 0x34,
	0x7c, 0x2d, 0x80, 0xc6, 0xd0, 0xca, 0x85, 0x62, 0xc5, 0xf2, 0x98, 0xf2, 0x86, 0x87, 0xf5, 0x62,
	0xe6, 0x80, 0x9b, 0x5a, 0x1e, 0x5f, 0x67, 0xa2, 0x6d, 0x70, 0x59, 0xaa, 0x37, 0x57, 0x13, 0xbb,
	0x2c, 0xf5, 0x5f, 0x41, 0x6b, 0xf9, 0x2f, 0xd4, 0x87, 0x4f, 0x66, 0xd1, 0x59, 0x84, 0x27, 0xaf,
	0x7f, 0x8c, 0x7f, 0x38, 0x99, 0x4d, 0xa3, 0xd1, 0xe4, 0x9b, 0x49, 0x34, 0xee, 0x35, 0xd0, 0x0e,
	0x78, 0xe3, 0x68, 0x8a, 0xa3, 0xd1, 0x8b, 0xd7, 0x93, 0xd3, 0x93, 0x9e, 0x83, 0x3c, 0x78, 0xf0,
	0xe6, 0x05, 0x3e, 0x99, 0x9c, 0xbc, 0xec, 0xb9, 0xa8, 0x0d, 0x9b, 0x11, 0xc6, 0xa7, 0xb8, 0xd7,
	0x1c, 0xfe, 0xe1, 0x02, 0x8c, 0x68, 0x36, 0x33, 0xb7, 0x07, 0x7a, 0x0b, 0x9b, 0x7a, 0xa0, 0xd1,
	0xc1, 0x1d, 0x13, 0x6f, 0x07, 0x6c, 0xf0, 0xe8, 0x4e, 0xce, 0x2c, 0xaa, 0xdf, 0xa8, 0xb4, 0xf5,
	0x88, 0xdf, 0xa6, 0xbd, 0x7a, 0xc0, 0xdd, 0xa6, 0xbd, 0x76, 0x72, 0xf8, 0x0d, 0xf4, 0x06, 0x36,
	0xaa, 0x11, 0x42, 0xfb, 0x1f, 0xb5, 0x2d, 0x06, 0x07, 0x77, 0x61, 0x4b, 0xe1, 0x23, 0x09, 0xbb,
	0x89, 0x78, 0x5f, 0x8b, 0x1f, 0xed, 0xdc, 0x98, 0x37, 0xad, 0xae, 0xb4, 0xa9, 0xf3, 0xf6, 0x6b,
	0x0b, 0xcf, 0x45, 0x46, 0xf8, 0x3c, 0x10, 0x72, 0x1e, 0xce, 0x29, 0xd7, 0x17, 0x5e, 0x68, 0x3e,
	0x91, 0x9c, 0xa9, 0x7f, 0xdf, 0x99, 0xcf, 0xab, 0xb7, 0xbf, 0x1d, 0xe7, 0x7c, 0x4b, 0xb3, 0x4f,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x76, 0x96, 0x35, 0xf1, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CelServiceClient is the client API for CelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CelServiceClient interface {
	// Transforms CEL source text into a parsed representation.
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	// Runs static checks on a parsed CEL representation and return
	// an annotated representation, or a set of issues.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Evaluates a parsed or annotation CEL representation given
	// values of external bindings.
	Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error)
}

type celServiceClient struct {
	cc *grpc.ClientConn
}

func NewCelServiceClient(cc *grpc.ClientConn) CelServiceClient {
	return &celServiceClient{cc}
}

func (c *celServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error) {
	out := new(EvalResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Eval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CelServiceServer is the server API for CelService service.
type CelServiceServer interface {
	// Transforms CEL source text into a parsed representation.
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	// Runs static checks on a parsed CEL representation and return
	// an annotated representation, or a set of issues.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Evaluates a parsed or annotation CEL representation given
	// values of external bindings.
	Eval(context.Context, *EvalRequest) (*EvalResponse, error)
}

func RegisterCelServiceServer(s *grpc.Server, srv CelServiceServer) {
	s.RegisterService(&_CelService_serviceDesc, srv)
}

func _CelService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Eval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Eval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Eval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Eval(ctx, req.(*EvalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.api.expr.v1alpha1.CelService",
	HandlerType: (*CelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _CelService_Parse_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CelService_Check_Handler,
		},
		{
			MethodName: "Eval",
			Handler:    _CelService_Eval_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/api/expr/v1alpha1/cel_service.proto",
}
