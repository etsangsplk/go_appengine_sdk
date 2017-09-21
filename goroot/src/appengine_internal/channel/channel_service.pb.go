// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/channel

It has these top-level messages:
	ChannelServiceError
	CreateChannelRequest
	CreateChannelResponse
	SendMessageRequest
*/
package appengine

import proto "appengine_internal/github.com/golang/protobuf/proto"
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

type ChannelServiceError_ErrorCode int32

const (
	ChannelServiceError_OK                             ChannelServiceError_ErrorCode = 0
	ChannelServiceError_INTERNAL_ERROR                 ChannelServiceError_ErrorCode = 1
	ChannelServiceError_INVALID_CHANNEL_KEY            ChannelServiceError_ErrorCode = 2
	ChannelServiceError_BAD_MESSAGE                    ChannelServiceError_ErrorCode = 3
	ChannelServiceError_INVALID_CHANNEL_TOKEN_DURATION ChannelServiceError_ErrorCode = 4
	ChannelServiceError_APPID_ALIAS_REQUIRED           ChannelServiceError_ErrorCode = 5
)

var ChannelServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "INVALID_CHANNEL_KEY",
	3: "BAD_MESSAGE",
	4: "INVALID_CHANNEL_TOKEN_DURATION",
	5: "APPID_ALIAS_REQUIRED",
}
var ChannelServiceError_ErrorCode_value = map[string]int32{
	"OK":                             0,
	"INTERNAL_ERROR":                 1,
	"INVALID_CHANNEL_KEY":            2,
	"BAD_MESSAGE":                    3,
	"INVALID_CHANNEL_TOKEN_DURATION": 4,
	"APPID_ALIAS_REQUIRED":           5,
}

func (x ChannelServiceError_ErrorCode) Enum() *ChannelServiceError_ErrorCode {
	p := new(ChannelServiceError_ErrorCode)
	*p = x
	return p
}
func (x ChannelServiceError_ErrorCode) String() string {
	return proto.EnumName(ChannelServiceError_ErrorCode_name, int32(x))
}
func (x *ChannelServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChannelServiceError_ErrorCode_value, data, "ChannelServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ChannelServiceError_ErrorCode(value)
	return nil
}

type ChannelServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChannelServiceError) Reset()         { *m = ChannelServiceError{} }
func (m *ChannelServiceError) String() string { return proto.CompactTextString(m) }
func (*ChannelServiceError) ProtoMessage()    {}

type CreateChannelRequest struct {
	ApplicationKey   *string `protobuf:"bytes,1,req,name=application_key,json=applicationKey" json:"application_key,omitempty"`
	DurationMinutes  *int32  `protobuf:"varint,2,opt,name=duration_minutes,json=durationMinutes" json:"duration_minutes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateChannelRequest) Reset()         { *m = CreateChannelRequest{} }
func (m *CreateChannelRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChannelRequest) ProtoMessage()    {}

func (m *CreateChannelRequest) GetApplicationKey() string {
	if m != nil && m.ApplicationKey != nil {
		return *m.ApplicationKey
	}
	return ""
}

func (m *CreateChannelRequest) GetDurationMinutes() int32 {
	if m != nil && m.DurationMinutes != nil {
		return *m.DurationMinutes
	}
	return 0
}

type CreateChannelResponse struct {
	Token            *string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	DurationMinutes  *int32  `protobuf:"varint,3,opt,name=duration_minutes,json=durationMinutes" json:"duration_minutes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateChannelResponse) Reset()         { *m = CreateChannelResponse{} }
func (m *CreateChannelResponse) String() string { return proto.CompactTextString(m) }
func (*CreateChannelResponse) ProtoMessage()    {}

func (m *CreateChannelResponse) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *CreateChannelResponse) GetDurationMinutes() int32 {
	if m != nil && m.DurationMinutes != nil {
		return *m.DurationMinutes
	}
	return 0
}

type SendMessageRequest struct {
	ApplicationKey   *string `protobuf:"bytes,1,req,name=application_key,json=applicationKey" json:"application_key,omitempty"`
	Message          *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}

func (m *SendMessageRequest) GetApplicationKey() string {
	if m != nil && m.ApplicationKey != nil {
		return *m.ApplicationKey
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ChannelServiceError)(nil), "appengine.ChannelServiceError")
	proto.RegisterType((*CreateChannelRequest)(nil), "appengine.CreateChannelRequest")
	proto.RegisterType((*CreateChannelResponse)(nil), "appengine.CreateChannelResponse")
	proto.RegisterType((*SendMessageRequest)(nil), "appengine.SendMessageRequest")
	proto.RegisterEnum("appengine.ChannelServiceError_ErrorCode", ChannelServiceError_ErrorCode_name, ChannelServiceError_ErrorCode_value)
}

func init() {
}
