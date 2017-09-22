// Code generated by protoc_gen_go. DO NOT EDIT.

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/capability

It has these top-level messages:
	CapabilityConfigList
	CapabilityConfig
	IsEnabledRequest
	IsEnabledResponse
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

type CapabilityConfig_Status int32

const (
	CapabilityConfig_DEFAULT   CapabilityConfig_Status = 0
	CapabilityConfig_ENABLED   CapabilityConfig_Status = 1
	CapabilityConfig_SCHEDULED CapabilityConfig_Status = 2
	CapabilityConfig_DISABLED  CapabilityConfig_Status = 3
	CapabilityConfig_UNKNOWN   CapabilityConfig_Status = 4
)

var CapabilityConfig_Status_name = map[int32]string{
	0: "DEFAULT",
	1: "ENABLED",
	2: "SCHEDULED",
	3: "DISABLED",
	4: "UNKNOWN",
}
var CapabilityConfig_Status_value = map[string]int32{
	"DEFAULT":   0,
	"ENABLED":   1,
	"SCHEDULED": 2,
	"DISABLED":  3,
	"UNKNOWN":   4,
}

func (x CapabilityConfig_Status) Enum() *CapabilityConfig_Status {
	p := new(CapabilityConfig_Status)
	*p = x
	return p
}
func (x CapabilityConfig_Status) String() string {
	return proto.EnumName(CapabilityConfig_Status_name, int32(x))
}
func (x *CapabilityConfig_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CapabilityConfig_Status_value, data, "CapabilityConfig_Status")
	if err != nil {
		return err
	}
	*x = CapabilityConfig_Status(value)
	return nil
}

type IsEnabledResponse_SummaryStatus int32

const (
	IsEnabledResponse_DEFAULT          IsEnabledResponse_SummaryStatus = 0
	IsEnabledResponse_ENABLED          IsEnabledResponse_SummaryStatus = 1
	IsEnabledResponse_SCHEDULED_FUTURE IsEnabledResponse_SummaryStatus = 2
	IsEnabledResponse_SCHEDULED_NOW    IsEnabledResponse_SummaryStatus = 3
	IsEnabledResponse_DISABLED         IsEnabledResponse_SummaryStatus = 4
	IsEnabledResponse_UNKNOWN          IsEnabledResponse_SummaryStatus = 5
)

var IsEnabledResponse_SummaryStatus_name = map[int32]string{
	0: "DEFAULT",
	1: "ENABLED",
	2: "SCHEDULED_FUTURE",
	3: "SCHEDULED_NOW",
	4: "DISABLED",
	5: "UNKNOWN",
}
var IsEnabledResponse_SummaryStatus_value = map[string]int32{
	"DEFAULT":          0,
	"ENABLED":          1,
	"SCHEDULED_FUTURE": 2,
	"SCHEDULED_NOW":    3,
	"DISABLED":         4,
	"UNKNOWN":          5,
}

func (x IsEnabledResponse_SummaryStatus) Enum() *IsEnabledResponse_SummaryStatus {
	p := new(IsEnabledResponse_SummaryStatus)
	*p = x
	return p
}
func (x IsEnabledResponse_SummaryStatus) String() string {
	return proto.EnumName(IsEnabledResponse_SummaryStatus_name, int32(x))
}
func (x *IsEnabledResponse_SummaryStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IsEnabledResponse_SummaryStatus_value, data, "IsEnabledResponse_SummaryStatus")
	if err != nil {
		return err
	}
	*x = IsEnabledResponse_SummaryStatus(value)
	return nil
}

type CapabilityConfigList struct {
	Config           []*CapabilityConfig `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
	DefaultConfig    *CapabilityConfig   `protobuf:"bytes,2,opt,name=default_config,json=defaultConfig" json:"default_config,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CapabilityConfigList) Reset()         { *m = CapabilityConfigList{} }
func (m *CapabilityConfigList) String() string { return proto.CompactTextString(m) }
func (*CapabilityConfigList) ProtoMessage()    {}

func (m *CapabilityConfigList) GetConfig() []*CapabilityConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *CapabilityConfigList) GetDefaultConfig() *CapabilityConfig {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

type CapabilityConfig struct {
	Package          *string                  `protobuf:"bytes,1,req,name=package" json:"package,omitempty"`
	Capability       *string                  `protobuf:"bytes,2,req,name=capability" json:"capability,omitempty"`
	Status           *CapabilityConfig_Status `protobuf:"varint,3,opt,name=status,enum=appengine.CapabilityConfig_Status,def=4" json:"status,omitempty"`
	ScheduledTime    *string                  `protobuf:"bytes,7,opt,name=scheduled_time,json=scheduledTime" json:"scheduled_time,omitempty"`
	InternalMessage  *string                  `protobuf:"bytes,4,opt,name=internal_message,json=internalMessage" json:"internal_message,omitempty"`
	AdminMessage     *string                  `protobuf:"bytes,5,opt,name=admin_message,json=adminMessage" json:"admin_message,omitempty"`
	ErrorMessage     *string                  `protobuf:"bytes,6,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *CapabilityConfig) Reset()         { *m = CapabilityConfig{} }
func (m *CapabilityConfig) String() string { return proto.CompactTextString(m) }
func (*CapabilityConfig) ProtoMessage()    {}

const Default_CapabilityConfig_Status CapabilityConfig_Status = CapabilityConfig_UNKNOWN

func (m *CapabilityConfig) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *CapabilityConfig) GetCapability() string {
	if m != nil && m.Capability != nil {
		return *m.Capability
	}
	return ""
}

func (m *CapabilityConfig) GetStatus() CapabilityConfig_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_CapabilityConfig_Status
}

func (m *CapabilityConfig) GetScheduledTime() string {
	if m != nil && m.ScheduledTime != nil {
		return *m.ScheduledTime
	}
	return ""
}

func (m *CapabilityConfig) GetInternalMessage() string {
	if m != nil && m.InternalMessage != nil {
		return *m.InternalMessage
	}
	return ""
}

func (m *CapabilityConfig) GetAdminMessage() string {
	if m != nil && m.AdminMessage != nil {
		return *m.AdminMessage
	}
	return ""
}

func (m *CapabilityConfig) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

type IsEnabledRequest struct {
	Package          *string  `protobuf:"bytes,1,req,name=package" json:"package,omitempty"`
	Capability       []string `protobuf:"bytes,2,rep,name=capability" json:"capability,omitempty"`
	Call             []string `protobuf:"bytes,3,rep,name=call" json:"call,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IsEnabledRequest) Reset()         { *m = IsEnabledRequest{} }
func (m *IsEnabledRequest) String() string { return proto.CompactTextString(m) }
func (*IsEnabledRequest) ProtoMessage()    {}

func (m *IsEnabledRequest) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *IsEnabledRequest) GetCapability() []string {
	if m != nil {
		return m.Capability
	}
	return nil
}

func (m *IsEnabledRequest) GetCall() []string {
	if m != nil {
		return m.Call
	}
	return nil
}

type IsEnabledResponse struct {
	SummaryStatus      *IsEnabledResponse_SummaryStatus `protobuf:"varint,1,opt,name=summary_status,json=summaryStatus,enum=appengine.IsEnabledResponse_SummaryStatus" json:"summary_status,omitempty"`
	TimeUntilScheduled *int64                           `protobuf:"varint,2,opt,name=time_until_scheduled,json=timeUntilScheduled" json:"time_until_scheduled,omitempty"`
	Config             []*CapabilityConfig              `protobuf:"bytes,3,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized   []byte                           `json:"-"`
}

func (m *IsEnabledResponse) Reset()         { *m = IsEnabledResponse{} }
func (m *IsEnabledResponse) String() string { return proto.CompactTextString(m) }
func (*IsEnabledResponse) ProtoMessage()    {}

func (m *IsEnabledResponse) GetSummaryStatus() IsEnabledResponse_SummaryStatus {
	if m != nil && m.SummaryStatus != nil {
		return *m.SummaryStatus
	}
	return IsEnabledResponse_DEFAULT
}

func (m *IsEnabledResponse) GetTimeUntilScheduled() int64 {
	if m != nil && m.TimeUntilScheduled != nil {
		return *m.TimeUntilScheduled
	}
	return 0
}

func (m *IsEnabledResponse) GetConfig() []*CapabilityConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*CapabilityConfigList)(nil), "appengine.CapabilityConfigList")
	proto.RegisterType((*CapabilityConfig)(nil), "appengine.CapabilityConfig")
	proto.RegisterType((*IsEnabledRequest)(nil), "appengine.IsEnabledRequest")
	proto.RegisterType((*IsEnabledResponse)(nil), "appengine.IsEnabledResponse")
	proto.RegisterEnum("appengine.CapabilityConfig_Status", CapabilityConfig_Status_name, CapabilityConfig_Status_value)
	proto.RegisterEnum("appengine.IsEnabledResponse_SummaryStatus", IsEnabledResponse_SummaryStatus_name, IsEnabledResponse_SummaryStatus_value)
}

func init() {
}
