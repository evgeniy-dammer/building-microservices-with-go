// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Currencies is an enum which represents the allowed currencies for the API
type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_BGN Currencies = 3
	Currencies_CZK Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_GBP Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_PLN Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_CHF Currencies = 11
	Currencies_ISK Currencies = 12
	Currencies_NOK Currencies = 13
	Currencies_HRK Currencies = 14
	Currencies_RUB Currencies = 15
	Currencies_TRY Currencies = 16
	Currencies_AUD Currencies = 17
	Currencies_BRL Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_CNY Currencies = 20
	Currencies_HKD Currencies = 21
	Currencies_IDR Currencies = 22
	Currencies_ILS Currencies = 23
	Currencies_INR Currencies = 24
	Currencies_KRW Currencies = 25
	Currencies_MXN Currencies = 26
	Currencies_MYR Currencies = 27
	Currencies_NZD Currencies = 28
	Currencies_PHP Currencies = 29
	Currencies_SGD Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_ZAR Currencies = 32
)

var Currencies_name = map[int32]string{
	0:  "EUR",
	1:  "USD",
	2:  "JPY",
	3:  "BGN",
	4:  "CZK",
	5:  "DKK",
	6:  "GBP",
	7:  "HUF",
	8:  "PLN",
	9:  "RON",
	10: "SEK",
	11: "CHF",
	12: "ISK",
	13: "NOK",
	14: "HRK",
	15: "RUB",
	16: "TRY",
	17: "AUD",
	18: "BRL",
	19: "CAD",
	20: "CNY",
	21: "HKD",
	22: "IDR",
	23: "ILS",
	24: "INR",
	25: "KRW",
	26: "MXN",
	27: "MYR",
	28: "NZD",
	29: "PHP",
	30: "SGD",
	31: "THB",
	32: "ZAR",
}

var Currencies_value = map[string]int32{
	"EUR": 0,
	"USD": 1,
	"JPY": 2,
	"BGN": 3,
	"CZK": 4,
	"DKK": 5,
	"GBP": 6,
	"HUF": 7,
	"PLN": 8,
	"RON": 9,
	"SEK": 10,
	"CHF": 11,
	"ISK": 12,
	"NOK": 13,
	"HRK": 14,
	"RUB": 15,
	"TRY": 16,
	"AUD": 17,
	"BRL": 18,
	"CAD": 19,
	"CNY": 20,
	"HKD": 21,
	"IDR": 22,
	"ILS": 23,
	"INR": 24,
	"KRW": 25,
	"MXN": 26,
	"MYR": 27,
	"NZD": 28,
	"PHP": 29,
	"SGD": 30,
	"THB": 31,
	"ZAR": 32,
}

func (x Currencies) String() string {
	return proto.EnumName(Currencies_name, int32(x))
}

func (Currencies) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

// RateRequest defines the request for a GetRate call
type RateRequest struct {
	// Base is the base currency code for the rate
	Base Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=protos.Currencies" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination          Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=protos.Currencies" json:"Destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *RateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateRequest.Merge(m, src)
}
func (m *RateRequest) XXX_Size() int {
	return xxx_messageInfo_RateRequest.Size(m)
}
func (m *RateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateRequest proto.InternalMessageInfo

func (m *RateRequest) GetBase() Currencies {
	if m != nil {
		return m.Base
	}
	return Currencies_EUR
}

func (m *RateRequest) GetDestination() Currencies {
	if m != nil {
		return m.Destination
	}
	return Currencies_EUR
}

// RateResponse is the response from a GetRate call
type RateResponse struct {
	// Base is the base currency code for the rate
	Base Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=protos.Currencies" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=protos.Currencies" json:"Destination,omitempty"`
	// Rate number and can be used to convert between the two currencies specified in the request.
	Rate                 float64  `protobuf:"fixed64,3,opt,name=Rate,proto3" json:"Rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{1}
}

func (m *RateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateResponse.Unmarshal(m, b)
}
func (m *RateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateResponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateResponse.Merge(m, src)
}
func (m *RateResponse) XXX_Size() int {
	return xxx_messageInfo_RateResponse.Size(m)
}
func (m *RateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateResponse proto.InternalMessageInfo

func (m *RateResponse) GetBase() Currencies {
	if m != nil {
		return m.Base
	}
	return Currencies_EUR
}

func (m *RateResponse) GetDestination() Currencies {
	if m != nil {
		return m.Destination
	}
	return Currencies_EUR
}

func (m *RateResponse) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

type StreamingRateResponse struct {
	// Types that are valid to be assigned to Message:
	//	*StreamingRateResponse_RateResponse
	//	*StreamingRateResponse_Error
	Message              isStreamingRateResponse_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StreamingRateResponse) Reset()         { *m = StreamingRateResponse{} }
func (m *StreamingRateResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingRateResponse) ProtoMessage()    {}
func (*StreamingRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{2}
}

func (m *StreamingRateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRateResponse.Unmarshal(m, b)
}
func (m *StreamingRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRateResponse.Marshal(b, m, deterministic)
}
func (m *StreamingRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRateResponse.Merge(m, src)
}
func (m *StreamingRateResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingRateResponse.Size(m)
}
func (m *StreamingRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRateResponse proto.InternalMessageInfo

type isStreamingRateResponse_Message interface {
	isStreamingRateResponse_Message()
}

type StreamingRateResponse_RateResponse struct {
	RateResponse *RateResponse `protobuf:"bytes,1,opt,name=rate_response,json=rateResponse,proto3,oneof"`
}

type StreamingRateResponse_Error struct {
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*StreamingRateResponse_RateResponse) isStreamingRateResponse_Message() {}

func (*StreamingRateResponse_Error) isStreamingRateResponse_Message() {}

func (m *StreamingRateResponse) GetMessage() isStreamingRateResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *StreamingRateResponse) GetRateResponse() *RateResponse {
	if x, ok := m.GetMessage().(*StreamingRateResponse_RateResponse); ok {
		return x.RateResponse
	}
	return nil
}

func (m *StreamingRateResponse) GetError() *status.Status {
	if x, ok := m.GetMessage().(*StreamingRateResponse_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamingRateResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamingRateResponse_RateResponse)(nil),
		(*StreamingRateResponse_Error)(nil),
	}
}

func init() {
	proto.RegisterEnum("protos.Currencies", Currencies_name, Currencies_value)
	proto.RegisterType((*RateRequest)(nil), "protos.RateRequest")
	proto.RegisterType((*RateResponse)(nil), "protos.RateResponse")
	proto.RegisterType((*StreamingRateResponse)(nil), "protos.StreamingRateResponse")
}

func init() {
	proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea)
}

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6f, 0x12, 0x41,
	0x14, 0xc6, 0x3b, 0x40, 0x0b, 0x1d, 0x28, 0x3e, 0xa7, 0xad, 0x45, 0xb4, 0x4a, 0x38, 0x18, 0xd2,
	0xc3, 0x62, 0xb0, 0x37, 0x4f, 0x6c, 0xb7, 0x65, 0x75, 0x71, 0x4b, 0x66, 0x25, 0x0a, 0x17, 0xb3,
	0x6c, 0x26, 0x64, 0xa3, 0xdd, 0xc5, 0x99, 0xe1, 0xd0, 0x9b, 0x67, 0xfd, 0x7b, 0xfc, 0xff, 0xcc,
	0x7b, 0xdb, 0x46, 0x9a, 0xe0, 0xd1, 0xd3, 0xfe, 0x32, 0xf3, 0x7d, 0xef, 0x7d, 0x79, 0x6f, 0x96,
	0x37, 0x93, 0xb5, 0xd6, 0x2a, 0x4b, 0x6e, 0x9d, 0x95, 0xce, 0x6d, 0x2e, 0xf6, 0xe8, 0x63, 0xda,
	0x27, 0xcb, 0x3c, 0x5f, 0x7e, 0x53, 0x7d, 0xbd, 0x4a, 0xfa, 0xc6, 0xc6, 0x76, 0x6d, 0x0a, 0x41,
	0xf7, 0x2b, 0xaf, 0xcb, 0xd8, 0x2a, 0xa9, 0xbe, 0xaf, 0x95, 0xb1, 0xe2, 0x15, 0xaf, 0xb8, 0xb1,
	0x51, 0x2d, 0xd6, 0x61, 0xbd, 0xe6, 0x40, 0x14, 0x22, 0xe3, 0x5c, 0x14, 0x55, 0x53, 0x65, 0x24,
	0xdd, 0x8b, 0x73, 0x5e, 0xf7, 0x94, 0xb1, 0x69, 0x16, 0xdb, 0x34, 0xcf, 0x5a, 0xa5, 0x7f, 0xca,
	0x37, 0x65, 0xdd, 0x1f, 0x8c, 0x37, 0x8a, 0x6e, 0x66, 0x95, 0x67, 0x46, 0xfd, 0xdf, 0x76, 0x42,
	0xf0, 0x0a, 0x76, 0x6b, 0x95, 0x3b, 0xac, 0xc7, 0x24, 0x71, 0xf7, 0x17, 0xe3, 0xc7, 0x91, 0xd5,
	0x2a, 0xbe, 0x49, 0xb3, 0xe5, 0x83, 0x2c, 0x6f, 0xf9, 0x81, 0x8e, 0xad, 0xfa, 0xa2, 0xef, 0x0e,
	0x28, 0x54, 0x7d, 0x70, 0x74, 0xdf, 0x65, 0x53, 0xec, 0xef, 0xc8, 0x86, 0xde, 0x34, 0x9f, 0xf1,
	0x5d, 0xa5, 0x75, 0xae, 0x29, 0x5a, 0x7d, 0x20, 0x9c, 0x62, 0xde, 0x8e, 0x5e, 0x25, 0x4e, 0x44,
	0xf3, 0xf6, 0x77, 0x64, 0x21, 0x71, 0xf7, 0x79, 0xf5, 0x46, 0x19, 0x13, 0x2f, 0xd5, 0xd9, 0xef,
	0x12, 0xe7, 0x7f, 0xd3, 0x8b, 0x2a, 0x2f, 0x5f, 0x4e, 0x25, 0xec, 0x20, 0x4c, 0x23, 0x0f, 0x18,
	0xc2, 0xfb, 0xc9, 0x0c, 0x4a, 0x08, 0xee, 0x28, 0x84, 0x32, 0xc2, 0xc5, 0x3c, 0x80, 0x0a, 0x82,
	0x17, 0x04, 0xb0, 0x8b, 0x30, 0x72, 0x27, 0xb0, 0x87, 0xe0, 0x4f, 0xaf, 0xa0, 0x8a, 0x30, 0x19,
	0x87, 0x50, 0x43, 0x90, 0xd7, 0x21, 0xec, 0x23, 0x44, 0x97, 0x01, 0x70, 0xb2, 0xfb, 0x57, 0x50,
	0x47, 0x78, 0x17, 0x05, 0xd0, 0x40, 0x08, 0xaf, 0x03, 0x38, 0x20, 0xbb, 0x0c, 0xa0, 0x49, 0xae,
	0xa9, 0x0b, 0x8f, 0x10, 0x3e, 0xca, 0x19, 0x00, 0xc2, 0x70, 0xea, 0xc1, 0x63, 0x8a, 0x21, 0xc7,
	0x20, 0xa8, 0xce, 0xd0, 0x83, 0x43, 0x82, 0x70, 0x06, 0x47, 0x64, 0x0f, 0x3c, 0x38, 0xa6, 0xca,
	0x9e, 0x84, 0x27, 0x04, 0xe3, 0x08, 0x4e, 0x08, 0x42, 0x09, 0x2d, 0x84, 0x40, 0x7e, 0x82, 0xa7,
	0x08, 0x1f, 0x3e, 0x87, 0xd0, 0x26, 0x98, 0x49, 0x78, 0x46, 0x31, 0xe6, 0x1e, 0x3c, 0xa7, 0xf0,
	0xfe, 0x04, 0x4e, 0x29, 0xf3, 0xc8, 0x83, 0x17, 0x14, 0xc3, 0x77, 0xe1, 0x25, 0xc2, 0x7c, 0x28,
	0xa1, 0x33, 0xf8, 0xc9, 0x78, 0xed, 0x6e, 0x6e, 0xb7, 0xe2, 0x9c, 0x57, 0x47, 0xca, 0xe2, 0x7a,
	0xc4, 0xe1, 0xc3, 0x65, 0xd1, 0x9b, 0x6e, 0x6f, 0xdd, 0xa0, 0xf0, 0x79, 0x33, 0x5a, 0x2f, 0x4c,
	0xa2, 0xd3, 0x85, 0xc2, 0x0b, 0xb3, 0xdd, 0x7c, 0x7a, 0x7f, 0xb8, 0xf5, 0xd1, 0xf4, 0xd8, 0x6b,
	0xe6, 0xf2, 0x79, 0xcd, 0xe9, 0x17, 0xaa, 0x45, 0xf1, 0xbf, 0xbd, 0xf9, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x31, 0x31, 0x62, 0x2f, 0x88, 0x03, 0x00, 0x00,
}
