// Code generated by protoc_gen_go.
// DO NOT EDIT!

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/memcache

It has these top-level messages:
	MemcacheServiceError
	AppOverride
	MemcacheGetRequest
	MemcacheGetResponse
	MemcacheSetRequest
	MemcacheSetResponse
	MemcacheDeleteRequest
	MemcacheDeleteResponse
	MemcacheIncrementRequest
	MemcacheIncrementResponse
	MemcacheBatchIncrementRequest
	MemcacheBatchIncrementResponse
	MemcacheFlushRequest
	MemcacheFlushResponse
	MemcacheStatsRequest
	MergedNamespaceStats
	MemcacheHotKey
	MemcacheStatsResponse
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

type MemcacheServiceError_ErrorCode int32

const (
	MemcacheServiceError_OK                MemcacheServiceError_ErrorCode = 0
	MemcacheServiceError_UNSPECIFIED_ERROR MemcacheServiceError_ErrorCode = 1
	MemcacheServiceError_NAMESPACE_NOT_SET MemcacheServiceError_ErrorCode = 2
	MemcacheServiceError_PERMISSION_DENIED MemcacheServiceError_ErrorCode = 3
	MemcacheServiceError_INVALID_VALUE     MemcacheServiceError_ErrorCode = 6
	MemcacheServiceError_UNAVAILABLE       MemcacheServiceError_ErrorCode = 9
)

var MemcacheServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "UNSPECIFIED_ERROR",
	2: "NAMESPACE_NOT_SET",
	3: "PERMISSION_DENIED",
	6: "INVALID_VALUE",
	9: "UNAVAILABLE",
}
var MemcacheServiceError_ErrorCode_value = map[string]int32{
	"OK":                0,
	"UNSPECIFIED_ERROR": 1,
	"NAMESPACE_NOT_SET": 2,
	"PERMISSION_DENIED": 3,
	"INVALID_VALUE":     6,
	"UNAVAILABLE":       9,
}

func (x MemcacheServiceError_ErrorCode) Enum() *MemcacheServiceError_ErrorCode {
	p := new(MemcacheServiceError_ErrorCode)
	*p = x
	return p
}
func (x MemcacheServiceError_ErrorCode) String() string {
	return proto.EnumName(MemcacheServiceError_ErrorCode_name, int32(x))
}
func (x *MemcacheServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheServiceError_ErrorCode_value, data, "MemcacheServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = MemcacheServiceError_ErrorCode(value)
	return nil
}

type MemcacheGetResponse_GetStatusCode int32

const (
	MemcacheGetResponse_HIT               MemcacheGetResponse_GetStatusCode = 1
	MemcacheGetResponse_MISS              MemcacheGetResponse_GetStatusCode = 2
	MemcacheGetResponse_TRUNCATED         MemcacheGetResponse_GetStatusCode = 3
	MemcacheGetResponse_DEADLINE_EXCEEDED MemcacheGetResponse_GetStatusCode = 4
	MemcacheGetResponse_UNREACHABLE       MemcacheGetResponse_GetStatusCode = 5
	MemcacheGetResponse_OTHER_ERROR       MemcacheGetResponse_GetStatusCode = 6
)

var MemcacheGetResponse_GetStatusCode_name = map[int32]string{
	1: "HIT",
	2: "MISS",
	3: "TRUNCATED",
	4: "DEADLINE_EXCEEDED",
	5: "UNREACHABLE",
	6: "OTHER_ERROR",
}
var MemcacheGetResponse_GetStatusCode_value = map[string]int32{
	"HIT":               1,
	"MISS":              2,
	"TRUNCATED":         3,
	"DEADLINE_EXCEEDED": 4,
	"UNREACHABLE":       5,
	"OTHER_ERROR":       6,
}

func (x MemcacheGetResponse_GetStatusCode) Enum() *MemcacheGetResponse_GetStatusCode {
	p := new(MemcacheGetResponse_GetStatusCode)
	*p = x
	return p
}
func (x MemcacheGetResponse_GetStatusCode) String() string {
	return proto.EnumName(MemcacheGetResponse_GetStatusCode_name, int32(x))
}
func (x *MemcacheGetResponse_GetStatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheGetResponse_GetStatusCode_value, data, "MemcacheGetResponse_GetStatusCode")
	if err != nil {
		return err
	}
	*x = MemcacheGetResponse_GetStatusCode(value)
	return nil
}

type MemcacheSetRequest_SetPolicy int32

const (
	MemcacheSetRequest_SET     MemcacheSetRequest_SetPolicy = 1
	MemcacheSetRequest_ADD     MemcacheSetRequest_SetPolicy = 2
	MemcacheSetRequest_REPLACE MemcacheSetRequest_SetPolicy = 3
	MemcacheSetRequest_CAS     MemcacheSetRequest_SetPolicy = 4
)

var MemcacheSetRequest_SetPolicy_name = map[int32]string{
	1: "SET",
	2: "ADD",
	3: "REPLACE",
	4: "CAS",
}
var MemcacheSetRequest_SetPolicy_value = map[string]int32{
	"SET":     1,
	"ADD":     2,
	"REPLACE": 3,
	"CAS":     4,
}

func (x MemcacheSetRequest_SetPolicy) Enum() *MemcacheSetRequest_SetPolicy {
	p := new(MemcacheSetRequest_SetPolicy)
	*p = x
	return p
}
func (x MemcacheSetRequest_SetPolicy) String() string {
	return proto.EnumName(MemcacheSetRequest_SetPolicy_name, int32(x))
}
func (x *MemcacheSetRequest_SetPolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheSetRequest_SetPolicy_value, data, "MemcacheSetRequest_SetPolicy")
	if err != nil {
		return err
	}
	*x = MemcacheSetRequest_SetPolicy(value)
	return nil
}

type MemcacheSetResponse_SetStatusCode int32

const (
	MemcacheSetResponse_STORED            MemcacheSetResponse_SetStatusCode = 1
	MemcacheSetResponse_NOT_STORED        MemcacheSetResponse_SetStatusCode = 2
	MemcacheSetResponse_ERROR             MemcacheSetResponse_SetStatusCode = 3
	MemcacheSetResponse_EXISTS            MemcacheSetResponse_SetStatusCode = 4
	MemcacheSetResponse_DEADLINE_EXCEEDED MemcacheSetResponse_SetStatusCode = 5
	MemcacheSetResponse_UNREACHABLE       MemcacheSetResponse_SetStatusCode = 6
	MemcacheSetResponse_OTHER_ERROR       MemcacheSetResponse_SetStatusCode = 7
)

var MemcacheSetResponse_SetStatusCode_name = map[int32]string{
	1: "STORED",
	2: "NOT_STORED",
	3: "ERROR",
	4: "EXISTS",
	5: "DEADLINE_EXCEEDED",
	6: "UNREACHABLE",
	7: "OTHER_ERROR",
}
var MemcacheSetResponse_SetStatusCode_value = map[string]int32{
	"STORED":            1,
	"NOT_STORED":        2,
	"ERROR":             3,
	"EXISTS":            4,
	"DEADLINE_EXCEEDED": 5,
	"UNREACHABLE":       6,
	"OTHER_ERROR":       7,
}

func (x MemcacheSetResponse_SetStatusCode) Enum() *MemcacheSetResponse_SetStatusCode {
	p := new(MemcacheSetResponse_SetStatusCode)
	*p = x
	return p
}
func (x MemcacheSetResponse_SetStatusCode) String() string {
	return proto.EnumName(MemcacheSetResponse_SetStatusCode_name, int32(x))
}
func (x *MemcacheSetResponse_SetStatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheSetResponse_SetStatusCode_value, data, "MemcacheSetResponse_SetStatusCode")
	if err != nil {
		return err
	}
	*x = MemcacheSetResponse_SetStatusCode(value)
	return nil
}

type MemcacheDeleteResponse_DeleteStatusCode int32

const (
	MemcacheDeleteResponse_DELETED           MemcacheDeleteResponse_DeleteStatusCode = 1
	MemcacheDeleteResponse_NOT_FOUND         MemcacheDeleteResponse_DeleteStatusCode = 2
	MemcacheDeleteResponse_DEADLINE_EXCEEDED MemcacheDeleteResponse_DeleteStatusCode = 3
	MemcacheDeleteResponse_UNREACHABLE       MemcacheDeleteResponse_DeleteStatusCode = 4
	MemcacheDeleteResponse_OTHER_ERROR       MemcacheDeleteResponse_DeleteStatusCode = 5
)

var MemcacheDeleteResponse_DeleteStatusCode_name = map[int32]string{
	1: "DELETED",
	2: "NOT_FOUND",
	3: "DEADLINE_EXCEEDED",
	4: "UNREACHABLE",
	5: "OTHER_ERROR",
}
var MemcacheDeleteResponse_DeleteStatusCode_value = map[string]int32{
	"DELETED":           1,
	"NOT_FOUND":         2,
	"DEADLINE_EXCEEDED": 3,
	"UNREACHABLE":       4,
	"OTHER_ERROR":       5,
}

func (x MemcacheDeleteResponse_DeleteStatusCode) Enum() *MemcacheDeleteResponse_DeleteStatusCode {
	p := new(MemcacheDeleteResponse_DeleteStatusCode)
	*p = x
	return p
}
func (x MemcacheDeleteResponse_DeleteStatusCode) String() string {
	return proto.EnumName(MemcacheDeleteResponse_DeleteStatusCode_name, int32(x))
}
func (x *MemcacheDeleteResponse_DeleteStatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheDeleteResponse_DeleteStatusCode_value, data, "MemcacheDeleteResponse_DeleteStatusCode")
	if err != nil {
		return err
	}
	*x = MemcacheDeleteResponse_DeleteStatusCode(value)
	return nil
}

type MemcacheIncrementRequest_Direction int32

const (
	MemcacheIncrementRequest_INCREMENT MemcacheIncrementRequest_Direction = 1
	MemcacheIncrementRequest_DECREMENT MemcacheIncrementRequest_Direction = 2
)

var MemcacheIncrementRequest_Direction_name = map[int32]string{
	1: "INCREMENT",
	2: "DECREMENT",
}
var MemcacheIncrementRequest_Direction_value = map[string]int32{
	"INCREMENT": 1,
	"DECREMENT": 2,
}

func (x MemcacheIncrementRequest_Direction) Enum() *MemcacheIncrementRequest_Direction {
	p := new(MemcacheIncrementRequest_Direction)
	*p = x
	return p
}
func (x MemcacheIncrementRequest_Direction) String() string {
	return proto.EnumName(MemcacheIncrementRequest_Direction_name, int32(x))
}
func (x *MemcacheIncrementRequest_Direction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheIncrementRequest_Direction_value, data, "MemcacheIncrementRequest_Direction")
	if err != nil {
		return err
	}
	*x = MemcacheIncrementRequest_Direction(value)
	return nil
}

type MemcacheIncrementResponse_IncrementStatusCode int32

const (
	MemcacheIncrementResponse_OK                MemcacheIncrementResponse_IncrementStatusCode = 1
	MemcacheIncrementResponse_NOT_CHANGED       MemcacheIncrementResponse_IncrementStatusCode = 2
	MemcacheIncrementResponse_ERROR             MemcacheIncrementResponse_IncrementStatusCode = 3
	MemcacheIncrementResponse_DEADLINE_EXCEEDED MemcacheIncrementResponse_IncrementStatusCode = 4
	MemcacheIncrementResponse_UNREACHABLE       MemcacheIncrementResponse_IncrementStatusCode = 5
	MemcacheIncrementResponse_OTHER_ERROR       MemcacheIncrementResponse_IncrementStatusCode = 6
)

var MemcacheIncrementResponse_IncrementStatusCode_name = map[int32]string{
	1: "OK",
	2: "NOT_CHANGED",
	3: "ERROR",
	4: "DEADLINE_EXCEEDED",
	5: "UNREACHABLE",
	6: "OTHER_ERROR",
}
var MemcacheIncrementResponse_IncrementStatusCode_value = map[string]int32{
	"OK":                1,
	"NOT_CHANGED":       2,
	"ERROR":             3,
	"DEADLINE_EXCEEDED": 4,
	"UNREACHABLE":       5,
	"OTHER_ERROR":       6,
}

func (x MemcacheIncrementResponse_IncrementStatusCode) Enum() *MemcacheIncrementResponse_IncrementStatusCode {
	p := new(MemcacheIncrementResponse_IncrementStatusCode)
	*p = x
	return p
}
func (x MemcacheIncrementResponse_IncrementStatusCode) String() string {
	return proto.EnumName(MemcacheIncrementResponse_IncrementStatusCode_name, int32(x))
}
func (x *MemcacheIncrementResponse_IncrementStatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MemcacheIncrementResponse_IncrementStatusCode_value, data, "MemcacheIncrementResponse_IncrementStatusCode")
	if err != nil {
		return err
	}
	*x = MemcacheIncrementResponse_IncrementStatusCode(value)
	return nil
}

type MemcacheServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MemcacheServiceError) Reset()         { *m = MemcacheServiceError{} }
func (m *MemcacheServiceError) String() string { return proto.CompactTextString(m) }
func (*MemcacheServiceError) ProtoMessage()    {}

type AppOverride struct {
	AppId            *string `protobuf:"bytes,1,req,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AppOverride) Reset()         { *m = AppOverride{} }
func (m *AppOverride) String() string { return proto.CompactTextString(m) }
func (*AppOverride) ProtoMessage()    {}

func (m *AppOverride) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

type MemcacheGetRequest struct {
	Key              [][]byte     `protobuf:"bytes,1,rep,name=key" json:"key,omitempty"`
	NameSpace        *string      `protobuf:"bytes,2,opt,name=name_space,json=nameSpace,def=" json:"name_space,omitempty"`
	ForCas           *bool        `protobuf:"varint,4,opt,name=for_cas,json=forCas" json:"for_cas,omitempty"`
	Override         *AppOverride `protobuf:"bytes,5,opt,name=override" json:"override,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MemcacheGetRequest) Reset()         { *m = MemcacheGetRequest{} }
func (m *MemcacheGetRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheGetRequest) ProtoMessage()    {}

func (m *MemcacheGetRequest) GetKey() [][]byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemcacheGetRequest) GetNameSpace() string {
	if m != nil && m.NameSpace != nil {
		return *m.NameSpace
	}
	return ""
}

func (m *MemcacheGetRequest) GetForCas() bool {
	if m != nil && m.ForCas != nil {
		return *m.ForCas
	}
	return false
}

func (m *MemcacheGetRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

type MemcacheGetResponse struct {
	Item             []*MemcacheGetResponse_Item         `protobuf:"group,1,rep,name=Item,json=item" json:"item,omitempty"`
	GetStatus        []MemcacheGetResponse_GetStatusCode `protobuf:"varint,7,rep,name=get_status,json=getStatus,enum=appengine.MemcacheGetResponse_GetStatusCode" json:"get_status,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *MemcacheGetResponse) Reset()         { *m = MemcacheGetResponse{} }
func (m *MemcacheGetResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheGetResponse) ProtoMessage()    {}

func (m *MemcacheGetResponse) GetItem() []*MemcacheGetResponse_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *MemcacheGetResponse) GetGetStatus() []MemcacheGetResponse_GetStatusCode {
	if m != nil {
		return m.GetStatus
	}
	return nil
}

type MemcacheGetResponse_Item struct {
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Value            []byte  `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	Flags            *uint32 `protobuf:"fixed32,4,opt,name=flags" json:"flags,omitempty"`
	CasId            *uint64 `protobuf:"fixed64,5,opt,name=cas_id,json=casId" json:"cas_id,omitempty"`
	ExpiresInSeconds *int32  `protobuf:"varint,6,opt,name=expires_in_seconds,json=expiresInSeconds" json:"expires_in_seconds,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MemcacheGetResponse_Item) Reset()         { *m = MemcacheGetResponse_Item{} }
func (m *MemcacheGetResponse_Item) String() string { return proto.CompactTextString(m) }
func (*MemcacheGetResponse_Item) ProtoMessage()    {}

func (m *MemcacheGetResponse_Item) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemcacheGetResponse_Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MemcacheGetResponse_Item) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *MemcacheGetResponse_Item) GetCasId() uint64 {
	if m != nil && m.CasId != nil {
		return *m.CasId
	}
	return 0
}

func (m *MemcacheGetResponse_Item) GetExpiresInSeconds() int32 {
	if m != nil && m.ExpiresInSeconds != nil {
		return *m.ExpiresInSeconds
	}
	return 0
}

type MemcacheSetRequest struct {
	Item             []*MemcacheSetRequest_Item `protobuf:"group,1,rep,name=Item,json=item" json:"item,omitempty"`
	NameSpace        *string                    `protobuf:"bytes,7,opt,name=name_space,json=nameSpace,def=" json:"name_space,omitempty"`
	Override         *AppOverride               `protobuf:"bytes,10,opt,name=override" json:"override,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *MemcacheSetRequest) Reset()         { *m = MemcacheSetRequest{} }
func (m *MemcacheSetRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheSetRequest) ProtoMessage()    {}

func (m *MemcacheSetRequest) GetItem() []*MemcacheSetRequest_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *MemcacheSetRequest) GetNameSpace() string {
	if m != nil && m.NameSpace != nil {
		return *m.NameSpace
	}
	return ""
}

func (m *MemcacheSetRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

type MemcacheSetRequest_Item struct {
	Key              []byte                        `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Value            []byte                        `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	Flags            *uint32                       `protobuf:"fixed32,4,opt,name=flags" json:"flags,omitempty"`
	SetPolicy        *MemcacheSetRequest_SetPolicy `protobuf:"varint,5,opt,name=set_policy,json=setPolicy,enum=appengine.MemcacheSetRequest_SetPolicy,def=1" json:"set_policy,omitempty"`
	ExpirationTime   *uint32                       `protobuf:"fixed32,6,opt,name=expiration_time,json=expirationTime,def=0" json:"expiration_time,omitempty"`
	CasId            *uint64                       `protobuf:"fixed64,8,opt,name=cas_id,json=casId" json:"cas_id,omitempty"`
	ForCas           *bool                         `protobuf:"varint,9,opt,name=for_cas,json=forCas" json:"for_cas,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *MemcacheSetRequest_Item) Reset()         { *m = MemcacheSetRequest_Item{} }
func (m *MemcacheSetRequest_Item) String() string { return proto.CompactTextString(m) }
func (*MemcacheSetRequest_Item) ProtoMessage()    {}

const Default_MemcacheSetRequest_Item_SetPolicy MemcacheSetRequest_SetPolicy = MemcacheSetRequest_SET
const Default_MemcacheSetRequest_Item_ExpirationTime uint32 = 0

func (m *MemcacheSetRequest_Item) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemcacheSetRequest_Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MemcacheSetRequest_Item) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *MemcacheSetRequest_Item) GetSetPolicy() MemcacheSetRequest_SetPolicy {
	if m != nil && m.SetPolicy != nil {
		return *m.SetPolicy
	}
	return Default_MemcacheSetRequest_Item_SetPolicy
}

func (m *MemcacheSetRequest_Item) GetExpirationTime() uint32 {
	if m != nil && m.ExpirationTime != nil {
		return *m.ExpirationTime
	}
	return Default_MemcacheSetRequest_Item_ExpirationTime
}

func (m *MemcacheSetRequest_Item) GetCasId() uint64 {
	if m != nil && m.CasId != nil {
		return *m.CasId
	}
	return 0
}

func (m *MemcacheSetRequest_Item) GetForCas() bool {
	if m != nil && m.ForCas != nil {
		return *m.ForCas
	}
	return false
}

type MemcacheSetResponse struct {
	SetStatus        []MemcacheSetResponse_SetStatusCode `protobuf:"varint,1,rep,name=set_status,json=setStatus,enum=appengine.MemcacheSetResponse_SetStatusCode" json:"set_status,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *MemcacheSetResponse) Reset()         { *m = MemcacheSetResponse{} }
func (m *MemcacheSetResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheSetResponse) ProtoMessage()    {}

func (m *MemcacheSetResponse) GetSetStatus() []MemcacheSetResponse_SetStatusCode {
	if m != nil {
		return m.SetStatus
	}
	return nil
}

type MemcacheDeleteRequest struct {
	Item             []*MemcacheDeleteRequest_Item `protobuf:"group,1,rep,name=Item,json=item" json:"item,omitempty"`
	NameSpace        *string                       `protobuf:"bytes,4,opt,name=name_space,json=nameSpace,def=" json:"name_space,omitempty"`
	Override         *AppOverride                  `protobuf:"bytes,5,opt,name=override" json:"override,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *MemcacheDeleteRequest) Reset()         { *m = MemcacheDeleteRequest{} }
func (m *MemcacheDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheDeleteRequest) ProtoMessage()    {}

func (m *MemcacheDeleteRequest) GetItem() []*MemcacheDeleteRequest_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *MemcacheDeleteRequest) GetNameSpace() string {
	if m != nil && m.NameSpace != nil {
		return *m.NameSpace
	}
	return ""
}

func (m *MemcacheDeleteRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

type MemcacheDeleteRequest_Item struct {
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	DeleteTime       *uint32 `protobuf:"fixed32,3,opt,name=delete_time,json=deleteTime,def=0" json:"delete_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MemcacheDeleteRequest_Item) Reset()         { *m = MemcacheDeleteRequest_Item{} }
func (m *MemcacheDeleteRequest_Item) String() string { return proto.CompactTextString(m) }
func (*MemcacheDeleteRequest_Item) ProtoMessage()    {}

const Default_MemcacheDeleteRequest_Item_DeleteTime uint32 = 0

func (m *MemcacheDeleteRequest_Item) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemcacheDeleteRequest_Item) GetDeleteTime() uint32 {
	if m != nil && m.DeleteTime != nil {
		return *m.DeleteTime
	}
	return Default_MemcacheDeleteRequest_Item_DeleteTime
}

type MemcacheDeleteResponse struct {
	DeleteStatus     []MemcacheDeleteResponse_DeleteStatusCode `protobuf:"varint,1,rep,name=delete_status,json=deleteStatus,enum=appengine.MemcacheDeleteResponse_DeleteStatusCode" json:"delete_status,omitempty"`
	XXX_unrecognized []byte                                    `json:"-"`
}

func (m *MemcacheDeleteResponse) Reset()         { *m = MemcacheDeleteResponse{} }
func (m *MemcacheDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheDeleteResponse) ProtoMessage()    {}

func (m *MemcacheDeleteResponse) GetDeleteStatus() []MemcacheDeleteResponse_DeleteStatusCode {
	if m != nil {
		return m.DeleteStatus
	}
	return nil
}

type MemcacheIncrementRequest struct {
	Key              []byte                              `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	NameSpace        *string                             `protobuf:"bytes,4,opt,name=name_space,json=nameSpace,def=" json:"name_space,omitempty"`
	Delta            *uint64                             `protobuf:"varint,2,opt,name=delta,def=1" json:"delta,omitempty"`
	Direction        *MemcacheIncrementRequest_Direction `protobuf:"varint,3,opt,name=direction,enum=appengine.MemcacheIncrementRequest_Direction,def=1" json:"direction,omitempty"`
	InitialValue     *uint64                             `protobuf:"varint,5,opt,name=initial_value,json=initialValue" json:"initial_value,omitempty"`
	InitialFlags     *uint32                             `protobuf:"fixed32,6,opt,name=initial_flags,json=initialFlags" json:"initial_flags,omitempty"`
	Override         *AppOverride                        `protobuf:"bytes,7,opt,name=override" json:"override,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *MemcacheIncrementRequest) Reset()         { *m = MemcacheIncrementRequest{} }
func (m *MemcacheIncrementRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheIncrementRequest) ProtoMessage()    {}

const Default_MemcacheIncrementRequest_Delta uint64 = 1
const Default_MemcacheIncrementRequest_Direction MemcacheIncrementRequest_Direction = MemcacheIncrementRequest_INCREMENT

func (m *MemcacheIncrementRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemcacheIncrementRequest) GetNameSpace() string {
	if m != nil && m.NameSpace != nil {
		return *m.NameSpace
	}
	return ""
}

func (m *MemcacheIncrementRequest) GetDelta() uint64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return Default_MemcacheIncrementRequest_Delta
}

func (m *MemcacheIncrementRequest) GetDirection() MemcacheIncrementRequest_Direction {
	if m != nil && m.Direction != nil {
		return *m.Direction
	}
	return Default_MemcacheIncrementRequest_Direction
}

func (m *MemcacheIncrementRequest) GetInitialValue() uint64 {
	if m != nil && m.InitialValue != nil {
		return *m.InitialValue
	}
	return 0
}

func (m *MemcacheIncrementRequest) GetInitialFlags() uint32 {
	if m != nil && m.InitialFlags != nil {
		return *m.InitialFlags
	}
	return 0
}

func (m *MemcacheIncrementRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

type MemcacheIncrementResponse struct {
	NewValue         *uint64                                        `protobuf:"varint,1,opt,name=new_value,json=newValue" json:"new_value,omitempty"`
	IncrementStatus  *MemcacheIncrementResponse_IncrementStatusCode `protobuf:"varint,2,opt,name=increment_status,json=incrementStatus,enum=appengine.MemcacheIncrementResponse_IncrementStatusCode" json:"increment_status,omitempty"`
	XXX_unrecognized []byte                                         `json:"-"`
}

func (m *MemcacheIncrementResponse) Reset()         { *m = MemcacheIncrementResponse{} }
func (m *MemcacheIncrementResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheIncrementResponse) ProtoMessage()    {}

func (m *MemcacheIncrementResponse) GetNewValue() uint64 {
	if m != nil && m.NewValue != nil {
		return *m.NewValue
	}
	return 0
}

func (m *MemcacheIncrementResponse) GetIncrementStatus() MemcacheIncrementResponse_IncrementStatusCode {
	if m != nil && m.IncrementStatus != nil {
		return *m.IncrementStatus
	}
	return MemcacheIncrementResponse_OK
}

type MemcacheBatchIncrementRequest struct {
	NameSpace        *string                     `protobuf:"bytes,1,opt,name=name_space,json=nameSpace,def=" json:"name_space,omitempty"`
	Item             []*MemcacheIncrementRequest `protobuf:"bytes,2,rep,name=item" json:"item,omitempty"`
	Override         *AppOverride                `protobuf:"bytes,3,opt,name=override" json:"override,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *MemcacheBatchIncrementRequest) Reset()         { *m = MemcacheBatchIncrementRequest{} }
func (m *MemcacheBatchIncrementRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheBatchIncrementRequest) ProtoMessage()    {}

func (m *MemcacheBatchIncrementRequest) GetNameSpace() string {
	if m != nil && m.NameSpace != nil {
		return *m.NameSpace
	}
	return ""
}

func (m *MemcacheBatchIncrementRequest) GetItem() []*MemcacheIncrementRequest {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *MemcacheBatchIncrementRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

type MemcacheBatchIncrementResponse struct {
	Item             []*MemcacheIncrementResponse `protobuf:"bytes,1,rep,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *MemcacheBatchIncrementResponse) Reset()         { *m = MemcacheBatchIncrementResponse{} }
func (m *MemcacheBatchIncrementResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheBatchIncrementResponse) ProtoMessage()    {}

func (m *MemcacheBatchIncrementResponse) GetItem() []*MemcacheIncrementResponse {
	if m != nil {
		return m.Item
	}
	return nil
}

type MemcacheFlushRequest struct {
	Override         *AppOverride `protobuf:"bytes,1,opt,name=override" json:"override,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MemcacheFlushRequest) Reset()         { *m = MemcacheFlushRequest{} }
func (m *MemcacheFlushRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheFlushRequest) ProtoMessage()    {}

func (m *MemcacheFlushRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

type MemcacheFlushResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MemcacheFlushResponse) Reset()         { *m = MemcacheFlushResponse{} }
func (m *MemcacheFlushResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheFlushResponse) ProtoMessage()    {}

type MemcacheStatsRequest struct {
	Override         *AppOverride `protobuf:"bytes,1,opt,name=override" json:"override,omitempty"`
	MaxHotkeyCount   *int32       `protobuf:"varint,2,opt,name=max_hotkey_count,json=maxHotkeyCount,def=0" json:"max_hotkey_count,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MemcacheStatsRequest) Reset()         { *m = MemcacheStatsRequest{} }
func (m *MemcacheStatsRequest) String() string { return proto.CompactTextString(m) }
func (*MemcacheStatsRequest) ProtoMessage()    {}

const Default_MemcacheStatsRequest_MaxHotkeyCount int32 = 0

func (m *MemcacheStatsRequest) GetOverride() *AppOverride {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *MemcacheStatsRequest) GetMaxHotkeyCount() int32 {
	if m != nil && m.MaxHotkeyCount != nil {
		return *m.MaxHotkeyCount
	}
	return Default_MemcacheStatsRequest_MaxHotkeyCount
}

type MergedNamespaceStats struct {
	Hits             *uint64           `protobuf:"varint,1,req,name=hits" json:"hits,omitempty"`
	Misses           *uint64           `protobuf:"varint,2,req,name=misses" json:"misses,omitempty"`
	ByteHits         *uint64           `protobuf:"varint,3,req,name=byte_hits,json=byteHits" json:"byte_hits,omitempty"`
	Items            *uint64           `protobuf:"varint,4,req,name=items" json:"items,omitempty"`
	Bytes            *uint64           `protobuf:"varint,5,req,name=bytes" json:"bytes,omitempty"`
	OldestItemAge    *uint32           `protobuf:"fixed32,6,req,name=oldest_item_age,json=oldestItemAge" json:"oldest_item_age,omitempty"`
	Hotkeys          []*MemcacheHotKey `protobuf:"bytes,7,rep,name=hotkeys" json:"hotkeys,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MergedNamespaceStats) Reset()         { *m = MergedNamespaceStats{} }
func (m *MergedNamespaceStats) String() string { return proto.CompactTextString(m) }
func (*MergedNamespaceStats) ProtoMessage()    {}

func (m *MergedNamespaceStats) GetHits() uint64 {
	if m != nil && m.Hits != nil {
		return *m.Hits
	}
	return 0
}

func (m *MergedNamespaceStats) GetMisses() uint64 {
	if m != nil && m.Misses != nil {
		return *m.Misses
	}
	return 0
}

func (m *MergedNamespaceStats) GetByteHits() uint64 {
	if m != nil && m.ByteHits != nil {
		return *m.ByteHits
	}
	return 0
}

func (m *MergedNamespaceStats) GetItems() uint64 {
	if m != nil && m.Items != nil {
		return *m.Items
	}
	return 0
}

func (m *MergedNamespaceStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *MergedNamespaceStats) GetOldestItemAge() uint32 {
	if m != nil && m.OldestItemAge != nil {
		return *m.OldestItemAge
	}
	return 0
}

func (m *MergedNamespaceStats) GetHotkeys() []*MemcacheHotKey {
	if m != nil {
		return m.Hotkeys
	}
	return nil
}

type MemcacheHotKey struct {
	Key              []byte   `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Qps              *float64 `protobuf:"fixed64,2,req,name=qps" json:"qps,omitempty"`
	NameSpace        *string  `protobuf:"bytes,3,opt,name=name_space,json=nameSpace" json:"name_space,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MemcacheHotKey) Reset()         { *m = MemcacheHotKey{} }
func (m *MemcacheHotKey) String() string { return proto.CompactTextString(m) }
func (*MemcacheHotKey) ProtoMessage()    {}

func (m *MemcacheHotKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemcacheHotKey) GetQps() float64 {
	if m != nil && m.Qps != nil {
		return *m.Qps
	}
	return 0
}

func (m *MemcacheHotKey) GetNameSpace() string {
	if m != nil && m.NameSpace != nil {
		return *m.NameSpace
	}
	return ""
}

type MemcacheStatsResponse struct {
	Stats            *MergedNamespaceStats `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *MemcacheStatsResponse) Reset()         { *m = MemcacheStatsResponse{} }
func (m *MemcacheStatsResponse) String() string { return proto.CompactTextString(m) }
func (*MemcacheStatsResponse) ProtoMessage()    {}

func (m *MemcacheStatsResponse) GetStats() *MergedNamespaceStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*MemcacheServiceError)(nil), "appengine.MemcacheServiceError")
	proto.RegisterType((*AppOverride)(nil), "appengine.AppOverride")
	proto.RegisterType((*MemcacheGetRequest)(nil), "appengine.MemcacheGetRequest")
	proto.RegisterType((*MemcacheGetResponse)(nil), "appengine.MemcacheGetResponse")
	proto.RegisterType((*MemcacheGetResponse_Item)(nil), "appengine.MemcacheGetResponse.Item")
	proto.RegisterType((*MemcacheSetRequest)(nil), "appengine.MemcacheSetRequest")
	proto.RegisterType((*MemcacheSetRequest_Item)(nil), "appengine.MemcacheSetRequest.Item")
	proto.RegisterType((*MemcacheSetResponse)(nil), "appengine.MemcacheSetResponse")
	proto.RegisterType((*MemcacheDeleteRequest)(nil), "appengine.MemcacheDeleteRequest")
	proto.RegisterType((*MemcacheDeleteRequest_Item)(nil), "appengine.MemcacheDeleteRequest.Item")
	proto.RegisterType((*MemcacheDeleteResponse)(nil), "appengine.MemcacheDeleteResponse")
	proto.RegisterType((*MemcacheIncrementRequest)(nil), "appengine.MemcacheIncrementRequest")
	proto.RegisterType((*MemcacheIncrementResponse)(nil), "appengine.MemcacheIncrementResponse")
	proto.RegisterType((*MemcacheBatchIncrementRequest)(nil), "appengine.MemcacheBatchIncrementRequest")
	proto.RegisterType((*MemcacheBatchIncrementResponse)(nil), "appengine.MemcacheBatchIncrementResponse")
	proto.RegisterType((*MemcacheFlushRequest)(nil), "appengine.MemcacheFlushRequest")
	proto.RegisterType((*MemcacheFlushResponse)(nil), "appengine.MemcacheFlushResponse")
	proto.RegisterType((*MemcacheStatsRequest)(nil), "appengine.MemcacheStatsRequest")
	proto.RegisterType((*MergedNamespaceStats)(nil), "appengine.MergedNamespaceStats")
	proto.RegisterType((*MemcacheHotKey)(nil), "appengine.MemcacheHotKey")
	proto.RegisterType((*MemcacheStatsResponse)(nil), "appengine.MemcacheStatsResponse")
	proto.RegisterEnum("appengine.MemcacheServiceError_ErrorCode", MemcacheServiceError_ErrorCode_name, MemcacheServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.MemcacheGetResponse_GetStatusCode", MemcacheGetResponse_GetStatusCode_name, MemcacheGetResponse_GetStatusCode_value)
	proto.RegisterEnum("appengine.MemcacheSetRequest_SetPolicy", MemcacheSetRequest_SetPolicy_name, MemcacheSetRequest_SetPolicy_value)
	proto.RegisterEnum("appengine.MemcacheSetResponse_SetStatusCode", MemcacheSetResponse_SetStatusCode_name, MemcacheSetResponse_SetStatusCode_value)
	proto.RegisterEnum("appengine.MemcacheDeleteResponse_DeleteStatusCode", MemcacheDeleteResponse_DeleteStatusCode_name, MemcacheDeleteResponse_DeleteStatusCode_value)
	proto.RegisterEnum("appengine.MemcacheIncrementRequest_Direction", MemcacheIncrementRequest_Direction_name, MemcacheIncrementRequest_Direction_value)
	proto.RegisterEnum("appengine.MemcacheIncrementResponse_IncrementStatusCode", MemcacheIncrementResponse_IncrementStatusCode_name, MemcacheIncrementResponse_IncrementStatusCode_value)
}

func init() {
}
