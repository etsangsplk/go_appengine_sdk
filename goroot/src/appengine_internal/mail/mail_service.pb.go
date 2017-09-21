// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/mail

It has these top-level messages:
	MailServiceError
	MailAttachment
	MailHeader
	MailMessage
*/
package appengine

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MailServiceError_ErrorCode int32

const (
	MailServiceError_OK                      MailServiceError_ErrorCode = 0
	MailServiceError_INTERNAL_ERROR          MailServiceError_ErrorCode = 1
	MailServiceError_BAD_REQUEST             MailServiceError_ErrorCode = 2
	MailServiceError_UNAUTHORIZED_SENDER     MailServiceError_ErrorCode = 3
	MailServiceError_INVALID_ATTACHMENT_TYPE MailServiceError_ErrorCode = 4
	MailServiceError_INVALID_HEADER_NAME     MailServiceError_ErrorCode = 5
	MailServiceError_INVALID_CONTENT_ID      MailServiceError_ErrorCode = 6
)

var MailServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "BAD_REQUEST",
	3: "UNAUTHORIZED_SENDER",
	4: "INVALID_ATTACHMENT_TYPE",
	5: "INVALID_HEADER_NAME",
	6: "INVALID_CONTENT_ID",
}
var MailServiceError_ErrorCode_value = map[string]int32{
	"OK":                      0,
	"INTERNAL_ERROR":          1,
	"BAD_REQUEST":             2,
	"UNAUTHORIZED_SENDER":     3,
	"INVALID_ATTACHMENT_TYPE": 4,
	"INVALID_HEADER_NAME":     5,
	"INVALID_CONTENT_ID":      6,
}

func (x MailServiceError_ErrorCode) Enum() *MailServiceError_ErrorCode {
	p := new(MailServiceError_ErrorCode)
	*p = x
	return p
}
func (x MailServiceError_ErrorCode) String() string {
	return proto.EnumName(MailServiceError_ErrorCode_name, int32(x))
}
func (x *MailServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MailServiceError_ErrorCode_value, data, "MailServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = MailServiceError_ErrorCode(value)
	return nil
}

type MailServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MailServiceError) Reset()         { *m = MailServiceError{} }
func (m *MailServiceError) String() string { return proto.CompactTextString(m) }
func (*MailServiceError) ProtoMessage()    {}

type MailAttachment struct {
	FileName         *string `protobuf:"bytes,1,req" json:"FileName,omitempty"`
	Data             []byte  `protobuf:"bytes,2,req" json:"Data,omitempty"`
	ContentID        *string `protobuf:"bytes,3,opt" json:"ContentID,omitempty"`
	ContentIDSet     *bool   `protobuf:"varint,13,opt,name=ContentID_set" json:"ContentID_set,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MailAttachment) Reset()         { *m = MailAttachment{} }
func (m *MailAttachment) String() string { return proto.CompactTextString(m) }
func (*MailAttachment) ProtoMessage()    {}

func (m *MailAttachment) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *MailAttachment) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MailAttachment) GetContentID() string {
	if m != nil && m.ContentID != nil {
		return *m.ContentID
	}
	return ""
}

func (m *MailAttachment) GetContentIDSet() bool {
	if m != nil && m.ContentIDSet != nil {
		return *m.ContentIDSet
	}
	return false
}

type MailHeader struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MailHeader) Reset()         { *m = MailHeader{} }
func (m *MailHeader) String() string { return proto.CompactTextString(m) }
func (*MailHeader) ProtoMessage()    {}

func (m *MailHeader) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MailHeader) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type MailMessage struct {
	Sender           *string           `protobuf:"bytes,1,req" json:"Sender,omitempty"`
	ReplyTo          *string           `protobuf:"bytes,2,opt" json:"ReplyTo,omitempty"`
	To               []string          `protobuf:"bytes,3,rep" json:"To,omitempty"`
	Cc               []string          `protobuf:"bytes,4,rep" json:"Cc,omitempty"`
	Bcc              []string          `protobuf:"bytes,5,rep" json:"Bcc,omitempty"`
	Subject          *string           `protobuf:"bytes,6,req" json:"Subject,omitempty"`
	TextBody         *string           `protobuf:"bytes,7,opt" json:"TextBody,omitempty"`
	HtmlBody         *string           `protobuf:"bytes,8,opt" json:"HtmlBody,omitempty"`
	Attachment       []*MailAttachment `protobuf:"bytes,9,rep" json:"Attachment,omitempty"`
	Header           []*MailHeader     `protobuf:"bytes,10,rep" json:"Header,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MailMessage) Reset()         { *m = MailMessage{} }
func (m *MailMessage) String() string { return proto.CompactTextString(m) }
func (*MailMessage) ProtoMessage()    {}

func (m *MailMessage) GetSender() string {
	if m != nil && m.Sender != nil {
		return *m.Sender
	}
	return ""
}

func (m *MailMessage) GetReplyTo() string {
	if m != nil && m.ReplyTo != nil {
		return *m.ReplyTo
	}
	return ""
}

func (m *MailMessage) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *MailMessage) GetCc() []string {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *MailMessage) GetBcc() []string {
	if m != nil {
		return m.Bcc
	}
	return nil
}

func (m *MailMessage) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

func (m *MailMessage) GetTextBody() string {
	if m != nil && m.TextBody != nil {
		return *m.TextBody
	}
	return ""
}

func (m *MailMessage) GetHtmlBody() string {
	if m != nil && m.HtmlBody != nil {
		return *m.HtmlBody
	}
	return ""
}

func (m *MailMessage) GetAttachment() []*MailAttachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *MailMessage) GetHeader() []*MailHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterEnum("appengine.MailServiceError_ErrorCode", MailServiceError_ErrorCode_name, MailServiceError_ErrorCode_value)
}
