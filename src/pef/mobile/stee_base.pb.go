// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stee_base.proto

package mobile

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

type STEEDataType int32

const (
	STEEDataType_DATATYPE_GET_REPORT_CONFIG STEEDataType = 1001
	STEEDataType_DATATYPE_DO_REPORT_DATA    STEEDataType = 1002
	STEEDataType_DATATYPE_CHECK_REQ         STEEDataType = 1003
	STEEDataType_DATATYPE_EVENT_REPORT      STEEDataType = 1004
)

var STEEDataType_name = map[int32]string{
	1001: "DATATYPE_GET_REPORT_CONFIG",
	1002: "DATATYPE_DO_REPORT_DATA",
	1003: "DATATYPE_CHECK_REQ",
	1004: "DATATYPE_EVENT_REPORT",
}

var STEEDataType_value = map[string]int32{
	"DATATYPE_GET_REPORT_CONFIG": 1001,
	"DATATYPE_DO_REPORT_DATA":    1002,
	"DATATYPE_CHECK_REQ":         1003,
	"DATATYPE_EVENT_REPORT":      1004,
}

func (x STEEDataType) Enum() *STEEDataType {
	p := new(STEEDataType)
	*p = x
	return p
}

func (x STEEDataType) String() string {
	return proto.EnumName(STEEDataType_name, int32(x))
}

func (x *STEEDataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(STEEDataType_value, data, "STEEDataType")
	if err != nil {
		return err
	}
	*x = STEEDataType(value)
	return nil
}

func (STEEDataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{0}
}

// 通讯错误类型，服务器侧示情况添加
type STEEErrorCode int32

const (
	STEEErrorCode_ERR_APP_KEY_NOT_FOUND_OR_NOT_EXISTS STEEErrorCode = -10001
	STEEErrorCode_ERR_CHECK_SIGN_FAILED               STEEErrorCode = -10002
	STEEErrorCode_ERR_UNKONW_BUSINESS_TYPE            STEEErrorCode = -10003
	STEEErrorCode_ERR_DESERIALIZATION_FAILED          STEEErrorCode = -10004
	STEEErrorCode_ERR_INTERNAL_SERVER_ERROR           STEEErrorCode = -10005
	STEEErrorCode_ERR_RESOURE_LIMITE                  STEEErrorCode = -10006
	STEEErrorCode_ERR_NOERROR                         STEEErrorCode = 0
)

var STEEErrorCode_name = map[int32]string{
	-10001: "ERR_APP_KEY_NOT_FOUND_OR_NOT_EXISTS",
	-10002: "ERR_CHECK_SIGN_FAILED",
	-10003: "ERR_UNKONW_BUSINESS_TYPE",
	-10004: "ERR_DESERIALIZATION_FAILED",
	-10005: "ERR_INTERNAL_SERVER_ERROR",
	-10006: "ERR_RESOURE_LIMITE",
	0:      "ERR_NOERROR",
}

var STEEErrorCode_value = map[string]int32{
	"ERR_APP_KEY_NOT_FOUND_OR_NOT_EXISTS": -10001,
	"ERR_CHECK_SIGN_FAILED":               -10002,
	"ERR_UNKONW_BUSINESS_TYPE":            -10003,
	"ERR_DESERIALIZATION_FAILED":          -10004,
	"ERR_INTERNAL_SERVER_ERROR":           -10005,
	"ERR_RESOURE_LIMITE":                  -10006,
	"ERR_NOERROR":                         0,
}

func (x STEEErrorCode) Enum() *STEEErrorCode {
	p := new(STEEErrorCode)
	*p = x
	return p
}

func (x STEEErrorCode) String() string {
	return proto.EnumName(STEEErrorCode_name, int32(x))
}

func (x *STEEErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(STEEErrorCode_value, data, "STEEErrorCode")
	if err != nil {
		return err
	}
	*x = STEEErrorCode(value)
	return nil
}

func (STEEErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{1}
}

type STEERequestHeader_OSType int32

const (
	STEERequestHeader_Android STEERequestHeader_OSType = 0
	STEERequestHeader_iOS     STEERequestHeader_OSType = 1
)

var STEERequestHeader_OSType_name = map[int32]string{
	0: "Android",
	1: "iOS",
}

var STEERequestHeader_OSType_value = map[string]int32{
	"Android": 0,
	"iOS":     1,
}

func (x STEERequestHeader_OSType) Enum() *STEERequestHeader_OSType {
	p := new(STEERequestHeader_OSType)
	*p = x
	return p
}

func (x STEERequestHeader_OSType) String() string {
	return proto.EnumName(STEERequestHeader_OSType_name, int32(x))
}

func (x *STEERequestHeader_OSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(STEERequestHeader_OSType_value, data, "STEERequestHeader_OSType")
	if err != nil {
		return err
	}
	*x = STEERequestHeader_OSType(value)
	return nil
}

func (STEERequestHeader_OSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{0, 0}
}

type STEERequestHeader_OSArch int32

const (
	STEERequestHeader_ARM    STEERequestHeader_OSArch = 0
	STEERequestHeader_X86    STEERequestHeader_OSArch = 1
	STEERequestHeader_X86_64 STEERequestHeader_OSArch = 2
	STEERequestHeader_ARM64  STEERequestHeader_OSArch = 3
)

var STEERequestHeader_OSArch_name = map[int32]string{
	0: "ARM",
	1: "X86",
	2: "X86_64",
	3: "ARM64",
}

var STEERequestHeader_OSArch_value = map[string]int32{
	"ARM":    0,
	"X86":    1,
	"X86_64": 2,
	"ARM64":  3,
}

func (x STEERequestHeader_OSArch) Enum() *STEERequestHeader_OSArch {
	p := new(STEERequestHeader_OSArch)
	*p = x
	return p
}

func (x STEERequestHeader_OSArch) String() string {
	return proto.EnumName(STEERequestHeader_OSArch_name, int32(x))
}

func (x *STEERequestHeader_OSArch) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(STEERequestHeader_OSArch_value, data, "STEERequestHeader_OSArch")
	if err != nil {
		return err
	}
	*x = STEERequestHeader_OSArch(value)
	return nil
}

func (STEERequestHeader_OSArch) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{0, 1}
}

type STEERequestHeader struct {
	SdkVer               *string                   `protobuf:"bytes,1,opt,name=sdk_ver,json=sdkVer" json:"sdk_ver,omitempty"`
	AppVer               *string                   `protobuf:"bytes,2,opt,name=app_ver,json=appVer" json:"app_ver,omitempty"`
	AppVerCode           *int32                    `protobuf:"varint,3,opt,name=app_ver_code,json=appVerCode" json:"app_ver_code,omitempty"`
	AppKey               *string                   `protobuf:"bytes,4,opt,name=app_key,json=appKey" json:"app_key,omitempty"`
	OsType               *STEERequestHeader_OSType `protobuf:"varint,5,opt,name=os_type,json=osType,enum=com.dingxiang.mobile.STEERequestHeader_OSType" json:"os_type,omitempty"`
	OsVer                *string                   `protobuf:"bytes,6,opt,name=os_ver,json=osVer" json:"os_ver,omitempty"`
	ConstId              *string                   `protobuf:"bytes,7,opt,name=const_id,json=constId" json:"const_id,omitempty"`
	ProtoVersion         *int32                    `protobuf:"varint,8,opt,name=proto_version,json=protoVersion" json:"proto_version,omitempty"`
	AppCode              *string                   `protobuf:"bytes,9,opt,name=app_code,json=appCode" json:"app_code,omitempty"`
	OsArch               *STEERequestHeader_OSArch `protobuf:"varint,10,opt,name=os_arch,json=osArch,enum=com.dingxiang.mobile.STEERequestHeader_OSArch" json:"os_arch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *STEERequestHeader) Reset()         { *m = STEERequestHeader{} }
func (m *STEERequestHeader) String() string { return proto.CompactTextString(m) }
func (*STEERequestHeader) ProtoMessage()    {}
func (*STEERequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{0}
}

func (m *STEERequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STEERequestHeader.Unmarshal(m, b)
}
func (m *STEERequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STEERequestHeader.Marshal(b, m, deterministic)
}
func (m *STEERequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STEERequestHeader.Merge(m, src)
}
func (m *STEERequestHeader) XXX_Size() int {
	return xxx_messageInfo_STEERequestHeader.Size(m)
}
func (m *STEERequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_STEERequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_STEERequestHeader proto.InternalMessageInfo

func (m *STEERequestHeader) GetSdkVer() string {
	if m != nil && m.SdkVer != nil {
		return *m.SdkVer
	}
	return ""
}

func (m *STEERequestHeader) GetAppVer() string {
	if m != nil && m.AppVer != nil {
		return *m.AppVer
	}
	return ""
}

func (m *STEERequestHeader) GetAppVerCode() int32 {
	if m != nil && m.AppVerCode != nil {
		return *m.AppVerCode
	}
	return 0
}

func (m *STEERequestHeader) GetAppKey() string {
	if m != nil && m.AppKey != nil {
		return *m.AppKey
	}
	return ""
}

func (m *STEERequestHeader) GetOsType() STEERequestHeader_OSType {
	if m != nil && m.OsType != nil {
		return *m.OsType
	}
	return STEERequestHeader_Android
}

func (m *STEERequestHeader) GetOsVer() string {
	if m != nil && m.OsVer != nil {
		return *m.OsVer
	}
	return ""
}

func (m *STEERequestHeader) GetConstId() string {
	if m != nil && m.ConstId != nil {
		return *m.ConstId
	}
	return ""
}

func (m *STEERequestHeader) GetProtoVersion() int32 {
	if m != nil && m.ProtoVersion != nil {
		return *m.ProtoVersion
	}
	return 0
}

func (m *STEERequestHeader) GetAppCode() string {
	if m != nil && m.AppCode != nil {
		return *m.AppCode
	}
	return ""
}

func (m *STEERequestHeader) GetOsArch() STEERequestHeader_OSArch {
	if m != nil && m.OsArch != nil {
		return *m.OsArch
	}
	return STEERequestHeader_ARM
}

type STEERequest struct {
	Header               *STEERequestHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Type                 *STEEDataType      `protobuf:"varint,2,req,name=type,enum=com.dingxiang.mobile.STEEDataType" json:"type,omitempty"`
	Data                 []byte             `protobuf:"bytes,3,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *STEERequest) Reset()         { *m = STEERequest{} }
func (m *STEERequest) String() string { return proto.CompactTextString(m) }
func (*STEERequest) ProtoMessage()    {}
func (*STEERequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{1}
}

func (m *STEERequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STEERequest.Unmarshal(m, b)
}
func (m *STEERequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STEERequest.Marshal(b, m, deterministic)
}
func (m *STEERequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STEERequest.Merge(m, src)
}
func (m *STEERequest) XXX_Size() int {
	return xxx_messageInfo_STEERequest.Size(m)
}
func (m *STEERequest) XXX_DiscardUnknown() {
	xxx_messageInfo_STEERequest.DiscardUnknown(m)
}

var xxx_messageInfo_STEERequest proto.InternalMessageInfo

func (m *STEERequest) GetHeader() *STEERequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *STEERequest) GetType() STEEDataType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return STEEDataType_DATATYPE_GET_REPORT_CONFIG
}

func (m *STEERequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type STEEResponse struct {
	Type                 *STEEDataType  `protobuf:"varint,1,req,name=type,enum=com.dingxiang.mobile.STEEDataType" json:"type,omitempty"`
	Data                 []byte         `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	ErrCode              *STEEErrorCode `protobuf:"varint,3,req,name=err_code,json=errCode,enum=com.dingxiang.mobile.STEEErrorCode" json:"err_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *STEEResponse) Reset()         { *m = STEEResponse{} }
func (m *STEEResponse) String() string { return proto.CompactTextString(m) }
func (*STEEResponse) ProtoMessage()    {}
func (*STEEResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{2}
}

func (m *STEEResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STEEResponse.Unmarshal(m, b)
}
func (m *STEEResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STEEResponse.Marshal(b, m, deterministic)
}
func (m *STEEResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STEEResponse.Merge(m, src)
}
func (m *STEEResponse) XXX_Size() int {
	return xxx_messageInfo_STEEResponse.Size(m)
}
func (m *STEEResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_STEEResponse.DiscardUnknown(m)
}

var xxx_messageInfo_STEEResponse proto.InternalMessageInfo

func (m *STEEResponse) GetType() STEEDataType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return STEEDataType_DATATYPE_GET_REPORT_CONFIG
}

func (m *STEEResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *STEEResponse) GetErrCode() STEEErrorCode {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return STEEErrorCode_ERR_APP_KEY_NOT_FOUND_OR_NOT_EXISTS
}

type Timestamp struct {
	Seconds              *int64   `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos                *int64   `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cce331aad09138, []int{3}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int64 {
	if m != nil && m.Nanos != nil {
		return *m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterEnum("com.dingxiang.mobile.STEEDataType", STEEDataType_name, STEEDataType_value)
	proto.RegisterEnum("com.dingxiang.mobile.STEEErrorCode", STEEErrorCode_name, STEEErrorCode_value)
	proto.RegisterEnum("com.dingxiang.mobile.STEERequestHeader_OSType", STEERequestHeader_OSType_name, STEERequestHeader_OSType_value)
	proto.RegisterEnum("com.dingxiang.mobile.STEERequestHeader_OSArch", STEERequestHeader_OSArch_name, STEERequestHeader_OSArch_value)
	proto.RegisterType((*STEERequestHeader)(nil), "com.dingxiang.mobile.STEERequestHeader")
	proto.RegisterType((*STEERequest)(nil), "com.dingxiang.mobile.STEERequest")
	proto.RegisterType((*STEEResponse)(nil), "com.dingxiang.mobile.STEEResponse")
	proto.RegisterType((*Timestamp)(nil), "com.dingxiang.mobile.Timestamp")
}

func init() { proto.RegisterFile("stee_base.proto", fileDescriptor_52cce331aad09138) }

var fileDescriptor_52cce331aad09138 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x6e, 0xf3, 0x34,
	0x14, 0xfe, 0x93, 0xac, 0xcd, 0x7a, 0xda, 0xff, 0xff, 0x83, 0xb5, 0x69, 0xd9, 0x84, 0x58, 0xd5,
	0x09, 0x56, 0xed, 0xa2, 0x82, 0x69, 0xaa, 0x26, 0x21, 0x81, 0xb2, 0xd6, 0xeb, 0xa2, 0x76, 0x49,
	0x71, 0xd2, 0xb2, 0x71, 0x63, 0x65, 0x8d, 0xb5, 0x45, 0xa3, 0x71, 0x88, 0x03, 0xa2, 0xb7, 0xbc,
	0x00, 0xf7, 0xbc, 0x05, 0x4f, 0xc3, 0x33, 0xb0, 0x0d, 0x78, 0x04, 0x90, 0x9d, 0xb5, 0x05, 0xc1,
	0x24, 0x20, 0x37, 0xf6, 0x39, 0xdf, 0xf7, 0x9d, 0xf3, 0x9d, 0xd8, 0x86, 0xb7, 0xa2, 0x60, 0x8c,
	0xde, 0x44, 0x82, 0x75, 0xb2, 0x9c, 0x17, 0x1c, 0x6d, 0xcd, 0xf8, 0xbc, 0x13, 0x27, 0xe9, 0xed,
	0xb7, 0x49, 0x94, 0xde, 0x76, 0xe6, 0xfc, 0x26, 0xf9, 0x92, 0xb5, 0x7e, 0x32, 0xe0, 0x9d, 0x20,
	0xc4, 0x98, 0xb0, 0xaf, 0xbe, 0x66, 0xa2, 0xb8, 0x60, 0x51, 0xcc, 0x72, 0xb4, 0x03, 0xa6, 0x88,
	0xef, 0xe9, 0x37, 0x2c, 0xb7, 0xb5, 0xa6, 0xd6, 0xae, 0x91, 0xaa, 0x88, 0xef, 0xa7, 0x25, 0x10,
	0x65, 0x99, 0x02, 0xf4, 0x12, 0x88, 0xb2, 0x4c, 0x02, 0x4d, 0x68, 0x3c, 0x03, 0x74, 0xc6, 0x63,
	0x66, 0x1b, 0x4d, 0xad, 0x5d, 0x21, 0x50, 0xa2, 0x3d, 0x1e, 0xb3, 0xa5, 0xf4, 0x9e, 0x2d, 0xec,
	0x8d, 0x95, 0x74, 0xc8, 0x16, 0x68, 0x00, 0x26, 0x17, 0xb4, 0x58, 0x64, 0xcc, 0xae, 0x34, 0xb5,
	0xf6, 0x9b, 0xe3, 0x4e, 0xe7, 0x9f, 0xac, 0x76, 0xfe, 0x66, 0xb3, 0xe3, 0x07, 0xe1, 0x22, 0x63,
	0xa4, 0xca, 0x85, 0x5c, 0xd1, 0x36, 0x54, 0xb9, 0x50, 0xde, 0xaa, 0xaa, 0x41, 0x85, 0x0b, 0x69,
	0x6d, 0x17, 0x36, 0x67, 0x3c, 0x15, 0x05, 0x4d, 0x62, 0xdb, 0x54, 0x80, 0xa9, 0x62, 0x37, 0x46,
	0x07, 0xf0, 0x5a, 0xfd, 0x1c, 0x29, 0x12, 0x09, 0x4f, 0xed, 0x4d, 0x65, 0xbb, 0xa1, 0x92, 0xd3,
	0x32, 0x27, 0xf5, 0xd2, 0xb8, 0x1a, 0xab, 0x56, 0xea, 0xa3, 0x2c, 0x53, 0x33, 0x95, 0xd6, 0xa3,
	0x7c, 0x76, 0x67, 0xc3, 0x7f, 0xb5, 0xee, 0xe4, 0xb3, 0x3b, 0x69, 0x5d, 0xae, 0xad, 0xf7, 0xa0,
	0x5a, 0x0e, 0x83, 0xea, 0x60, 0x3a, 0x69, 0x9c, 0xf3, 0x24, 0xb6, 0x5e, 0x21, 0x13, 0x8c, 0xc4,
	0x0f, 0x2c, 0xad, 0xf5, 0x91, 0xc4, 0x25, 0x53, 0xa6, 0x1c, 0x72, 0x59, 0x62, 0x57, 0xa7, 0x5d,
	0x4b, 0x43, 0x00, 0xd5, 0xab, 0xd3, 0x2e, 0xed, 0x9e, 0x58, 0x3a, 0xaa, 0x41, 0xc5, 0x21, 0x97,
	0xdd, 0x13, 0xcb, 0x68, 0xfd, 0xa0, 0x41, 0xfd, 0x4f, 0x7d, 0xd1, 0xa7, 0x50, 0xbd, 0x53, 0xbd,
	0x6d, 0xad, 0xa9, 0xb7, 0xeb, 0xc7, 0x87, 0xff, 0xd2, 0x2a, 0x79, 0x96, 0xa1, 0x2e, 0x6c, 0xa8,
	0x43, 0xd2, 0x9b, 0x7a, 0xfb, 0xcd, 0x71, 0xeb, 0x65, 0x79, 0x3f, 0x2a, 0x22, 0x75, 0x30, 0x8a,
	0x8f, 0x10, 0x6c, 0xc4, 0x51, 0x11, 0xd9, 0x46, 0x53, 0x6f, 0x37, 0x88, 0xda, 0x4b, 0x73, 0x8d,
	0xb2, 0x93, 0xc8, 0x78, 0x2a, 0xd8, 0xaa, 0xb8, 0xf6, 0x3f, 0x8b, 0xeb, 0xeb, 0xe2, 0xe8, 0x13,
	0xd8, 0x64, 0xf9, 0xea, 0x1e, 0xca, 0x7a, 0x07, 0x2f, 0xd7, 0xc3, 0x79, 0xce, 0xd5, 0x05, 0x25,
	0x26, 0xcb, 0xd5, 0xa6, 0xf5, 0x31, 0xd4, 0xc2, 0x64, 0xce, 0x44, 0x11, 0xcd, 0x33, 0x64, 0x83,
	0x29, 0xd8, 0x8c, 0xa7, 0xb1, 0x50, 0x4f, 0xc1, 0x20, 0xcb, 0x10, 0x6d, 0x41, 0x25, 0x8d, 0x52,
	0x2e, 0xd4, 0x4b, 0x30, 0x48, 0x19, 0x1c, 0x7d, 0xf7, 0x3c, 0xd9, 0xd2, 0x27, 0xda, 0x87, 0xbd,
	0xbe, 0x13, 0x3a, 0xe1, 0xf5, 0x18, 0xd3, 0x01, 0x0e, 0x29, 0xc1, 0x63, 0x9f, 0x84, 0xb4, 0xe7,
	0x7b, 0xe7, 0xee, 0xc0, 0xfa, 0xd9, 0x44, 0xef, 0xc2, 0xce, 0x8a, 0xd0, 0xf7, 0x97, 0xb8, 0x4c,
	0x59, 0x0f, 0x26, 0xda, 0x01, 0xb4, 0x42, 0x7b, 0x17, 0xb8, 0x37, 0xa4, 0x04, 0x7f, 0x66, 0x3d,
	0x9a, 0x68, 0x0f, 0xb6, 0x57, 0x00, 0x9e, 0x62, 0x6f, 0x59, 0xd9, 0x7a, 0x32, 0x8f, 0xbe, 0xd7,
	0xe1, 0xf5, 0x5f, 0x86, 0x43, 0x1f, 0xc2, 0x01, 0x26, 0x84, 0x3a, 0xe3, 0x31, 0x1d, 0xe2, 0x6b,
	0xea, 0xf9, 0x21, 0x3d, 0xf7, 0x27, 0x5e, 0x9f, 0xfa, 0x44, 0x05, 0xf8, 0xca, 0x0d, 0xc2, 0xc0,
	0xfa, 0xed, 0xc7, 0xdf, 0xcb, 0x4f, 0x43, 0x2d, 0xd8, 0x96, 0x8a, 0xb2, 0x67, 0xe0, 0x0e, 0x3c,
	0x7a, 0xee, 0xb8, 0x23, 0xdc, 0xb7, 0x7e, 0x5d, 0x73, 0xde, 0x07, 0x5b, 0x72, 0x26, 0xde, 0xd0,
	0xf7, 0x3e, 0xa7, 0x67, 0x93, 0xc0, 0xf5, 0x70, 0x10, 0x50, 0xe9, 0xc9, 0xfa, 0x65, 0x4d, 0x3b,
	0x84, 0x3d, 0x49, 0xeb, 0xe3, 0x00, 0x13, 0xd7, 0x19, 0xb9, 0x5f, 0x38, 0xa1, 0xeb, 0xaf, 0xea,
	0x3d, 0xad, 0x89, 0x1f, 0xc0, 0xae, 0x24, 0xba, 0x5e, 0x88, 0x89, 0xe7, 0x8c, 0x68, 0x80, 0xc9,
	0x14, 0x13, 0x8a, 0x09, 0xf1, 0x89, 0xf5, 0xb8, 0xe6, 0xed, 0x03, 0x92, 0x3c, 0x82, 0x03, 0x7f,
	0x42, 0x30, 0x1d, 0xb9, 0x97, 0x6e, 0x88, 0xad, 0x87, 0x35, 0xe1, 0x2d, 0xd4, 0x25, 0xc1, 0xf3,
	0x4b, 0xe9, 0xab, 0x33, 0xfd, 0xc2, 0xf8, 0x23, 0x00, 0x00, 0xff, 0xff, 0x36, 0x6b, 0xd7, 0x50,
	0x13, 0x05, 0x00, 0x00,
}