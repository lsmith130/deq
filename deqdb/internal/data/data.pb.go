// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data.proto

package data

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EventState int32

const (
	EventState_UNSPECIFIED_STATE   EventState = 0
	EventState_QUEUED              EventState = 1
	EventState_QUEUED_LINEAR       EventState = 7
	EventState_QUEUED_CONSTANT     EventState = 8
	EventState_OK                  EventState = 2
	EventState_INVALID             EventState = 4
	EventState_INTERNAL            EventState = 5
	EventState_SEND_LIMIT_EXCEEDED EventState = 6
	EventState_DEQUEUED_ERROR      EventState = 3
)

var EventState_name = map[int32]string{
	0: "UNSPECIFIED_STATE",
	1: "QUEUED",
	7: "QUEUED_LINEAR",
	8: "QUEUED_CONSTANT",
	2: "OK",
	4: "INVALID",
	5: "INTERNAL",
	6: "SEND_LIMIT_EXCEEDED",
	3: "DEQUEUED_ERROR",
}

var EventState_value = map[string]int32{
	"UNSPECIFIED_STATE":   0,
	"QUEUED":              1,
	"QUEUED_LINEAR":       7,
	"QUEUED_CONSTANT":     8,
	"OK":                  2,
	"INVALID":             4,
	"INTERNAL":            5,
	"SEND_LIMIT_EXCEEDED": 6,
	"DEQUEUED_ERROR":      3,
}

func (x EventState) String() string {
	return proto.EnumName(EventState_name, int32(x))
}

func (EventState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

type ChannelPayload struct {
	EventState EventState `protobuf:"varint,1,opt,name=event_state,json=eventState,proto3,enum=EventState" json:"event_state,omitempty"`
	// Deprecated as of db format v1.2.0. In db v1.2.0 and later, use SendCount message instead.
	DeprecatedSendCount int32 `protobuf:"varint,2,opt,name=deprecated_send_count,json=deprecatedSendCount,proto3" json:"deprecated_send_count,omitempty"`
}

func (m *ChannelPayload) Reset()         { *m = ChannelPayload{} }
func (m *ChannelPayload) String() string { return proto.CompactTextString(m) }
func (*ChannelPayload) ProtoMessage()    {}
func (*ChannelPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}
func (m *ChannelPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChannelPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChannelPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChannelPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelPayload.Merge(m, src)
}
func (m *ChannelPayload) XXX_Size() int {
	return m.Size()
}
func (m *ChannelPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelPayload.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelPayload proto.InternalMessageInfo

func (m *ChannelPayload) GetEventState() EventState {
	if m != nil {
		return m.EventState
	}
	return EventState_UNSPECIFIED_STATE
}

func (m *ChannelPayload) GetDeprecatedSendCount() int32 {
	if m != nil {
		return m.DeprecatedSendCount
	}
	return 0
}

type SendCount struct {
	SendCount int32 `protobuf:"varint,1,opt,name=send_count,json=sendCount,proto3" json:"send_count,omitempty"`
}

func (m *SendCount) Reset()         { *m = SendCount{} }
func (m *SendCount) String() string { return proto.CompactTextString(m) }
func (*SendCount) ProtoMessage()    {}
func (*SendCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}
func (m *SendCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCount.Merge(m, src)
}
func (m *SendCount) XXX_Size() int {
	return m.Size()
}
func (m *SendCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCount.DiscardUnknown(m)
}

var xxx_messageInfo_SendCount proto.InternalMessageInfo

func (m *SendCount) GetSendCount() int32 {
	if m != nil {
		return m.SendCount
	}
	return 0
}

type EventTimePayload struct {
	CreateTime int64 `protobuf:"fixed64,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (m *EventTimePayload) Reset()         { *m = EventTimePayload{} }
func (m *EventTimePayload) String() string { return proto.CompactTextString(m) }
func (*EventTimePayload) ProtoMessage()    {}
func (*EventTimePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}
func (m *EventTimePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTimePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTimePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTimePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTimePayload.Merge(m, src)
}
func (m *EventTimePayload) XXX_Size() int {
	return m.Size()
}
func (m *EventTimePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTimePayload.DiscardUnknown(m)
}

var xxx_messageInfo_EventTimePayload proto.InternalMessageInfo

func (m *EventTimePayload) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type IndexPayload struct {
	EventId    string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	CreateTime int64  `protobuf:"fixed64,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Version    int64  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *IndexPayload) Reset()         { *m = IndexPayload{} }
func (m *IndexPayload) String() string { return proto.CompactTextString(m) }
func (*IndexPayload) ProtoMessage()    {}
func (*IndexPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{3}
}
func (m *IndexPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexPayload.Merge(m, src)
}
func (m *IndexPayload) XXX_Size() int {
	return m.Size()
}
func (m *IndexPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexPayload.DiscardUnknown(m)
}

var xxx_messageInfo_IndexPayload proto.InternalMessageInfo

func (m *IndexPayload) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *IndexPayload) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *IndexPayload) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type EventPayload struct {
	Payload           []byte     `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	DefaultEventState EventState `protobuf:"varint,2,opt,name=default_event_state,json=defaultEventState,proto3,enum=EventState" json:"default_event_state,omitempty"`
	Indexes           []string   `protobuf:"bytes,3,rep,name=indexes,proto3" json:"indexes,omitempty"`
}

func (m *EventPayload) Reset()         { *m = EventPayload{} }
func (m *EventPayload) String() string { return proto.CompactTextString(m) }
func (*EventPayload) ProtoMessage()    {}
func (*EventPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{4}
}
func (m *EventPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPayload.Merge(m, src)
}
func (m *EventPayload) XXX_Size() int {
	return m.Size()
}
func (m *EventPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EventPayload proto.InternalMessageInfo

func (m *EventPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *EventPayload) GetDefaultEventState() EventState {
	if m != nil {
		return m.DefaultEventState
	}
	return EventState_UNSPECIFIED_STATE
}

func (m *EventPayload) GetIndexes() []string {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func init() {
	proto.RegisterEnum("EventState", EventState_name, EventState_value)
	proto.RegisterType((*ChannelPayload)(nil), "ChannelPayload")
	proto.RegisterType((*SendCount)(nil), "SendCount")
	proto.RegisterType((*EventTimePayload)(nil), "EventTimePayload")
	proto.RegisterType((*IndexPayload)(nil), "IndexPayload")
	proto.RegisterType((*EventPayload)(nil), "EventPayload")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xda, 0x34, 0x4e, 0x26, 0x21, 0x6c, 0x36, 0xaa, 0x30, 0x07, 0x4c, 0x94, 0x53, 0x54,
	0xa1, 0x1e, 0xda, 0x23, 0x27, 0x63, 0x2f, 0xd2, 0x8a, 0xe0, 0x94, 0xb5, 0x83, 0xb8, 0x59, 0x4b,
	0x76, 0x11, 0x96, 0x52, 0x3b, 0xb2, 0xb7, 0x15, 0x9c, 0xf8, 0x0b, 0xfc, 0x0b, 0xfe, 0x0a, 0xc7,
	0x1e, 0x39, 0xa2, 0xe4, 0x8f, 0x20, 0x7f, 0xa5, 0x51, 0x7b, 0x9b, 0xf7, 0xf6, 0xcd, 0x9b, 0x37,
	0xa3, 0x05, 0x90, 0x42, 0x8b, 0xf3, 0x6d, 0x9e, 0xe9, 0x6c, 0x96, 0xc3, 0xc8, 0xfb, 0x26, 0xd2,
	0x54, 0x6d, 0xae, 0xc4, 0x8f, 0x4d, 0x26, 0x24, 0x79, 0x0d, 0x03, 0x75, 0xab, 0x52, 0x1d, 0x17,
	0x5a, 0x68, 0x65, 0xa3, 0x29, 0x9a, 0x8f, 0x2e, 0x06, 0xe7, 0xb4, 0xe4, 0xc2, 0x92, 0xe2, 0xa0,
	0x0e, 0x35, 0xb9, 0x80, 0x53, 0xa9, 0xb6, 0xb9, 0x5a, 0x0b, 0xad, 0x64, 0x5c, 0xa8, 0x54, 0xc6,
	0xeb, 0xec, 0x26, 0xd5, 0xb6, 0x31, 0x45, 0xf3, 0x13, 0x3e, 0xb9, 0x7f, 0x0c, 0x55, 0x2a, 0xbd,
	0xf2, 0x69, 0x76, 0x06, 0xfd, 0x03, 0x20, 0x2f, 0x01, 0x8e, 0xba, 0x50, 0xd5, 0xd5, 0x2f, 0x0e,
	0xda, 0x4b, 0xc0, 0xd5, 0xe4, 0x28, 0xb9, 0x56, 0x6d, 0xc2, 0x57, 0x30, 0x58, 0xe7, 0x4a, 0x68,
	0x15, 0xeb, 0xe4, 0xba, 0x4e, 0x88, 0x39, 0xd4, 0x54, 0xa9, 0x9b, 0x49, 0x18, 0xb2, 0x54, 0xaa,
	0xef, 0x6d, 0xc3, 0x0b, 0xe8, 0xd5, 0x2b, 0x25, 0xb2, 0x52, 0xf7, 0xb9, 0x55, 0x61, 0xf6, 0xc8,
	0xcb, 0x78, 0xe8, 0x45, 0x6c, 0xb0, 0x6e, 0x55, 0x5e, 0x24, 0x59, 0x6a, 0x9b, 0x53, 0x34, 0x37,
	0x79, 0x0b, 0x67, 0x3f, 0x61, 0x58, 0x45, 0x6b, 0xa7, 0xd8, 0x60, 0x6d, 0xeb, 0xb2, 0x1a, 0x32,
	0xe4, 0x2d, 0x24, 0x6f, 0x60, 0x22, 0xd5, 0x57, 0x71, 0xb3, 0xd1, 0xf1, 0xf1, 0x69, 0x8d, 0xc7,
	0xa7, 0x1d, 0x37, 0xba, 0x7b, 0xaa, 0xb4, 0x4d, 0xca, 0x65, 0x54, 0x61, 0x9b, 0x53, 0xb3, 0xcc,
	0xde, 0xc0, 0xb3, 0xdf, 0x08, 0xe0, 0x48, 0x78, 0x0a, 0xe3, 0x55, 0x10, 0x5e, 0x51, 0x8f, 0xbd,
	0x63, 0xd4, 0x8f, 0xc3, 0xc8, 0x8d, 0x28, 0xee, 0x10, 0x80, 0xee, 0xc7, 0x15, 0x5d, 0x51, 0x1f,
	0x23, 0x32, 0x86, 0xa7, 0x75, 0x1d, 0x2f, 0x58, 0x40, 0x5d, 0x8e, 0x2d, 0x32, 0x81, 0x67, 0x0d,
	0xe5, 0x2d, 0x83, 0x30, 0x72, 0x83, 0x08, 0xf7, 0x48, 0x17, 0x8c, 0xe5, 0x7b, 0x6c, 0x90, 0x01,
	0x58, 0x2c, 0xf8, 0xe4, 0x2e, 0x98, 0x8f, 0x9f, 0x90, 0x21, 0xf4, 0x58, 0x10, 0x51, 0x1e, 0xb8,
	0x0b, 0x7c, 0x42, 0x9e, 0xc3, 0x24, 0xa4, 0x41, 0x69, 0xf4, 0x81, 0x45, 0x31, 0xfd, 0xec, 0x51,
	0xea, 0x53, 0x1f, 0x77, 0x09, 0x81, 0x91, 0x4f, 0x1b, 0x4b, 0xca, 0xf9, 0x92, 0x63, 0xf3, 0xad,
	0xfd, 0x67, 0xe7, 0xa0, 0xbb, 0x9d, 0x83, 0xfe, 0xed, 0x1c, 0xf4, 0x6b, 0xef, 0x74, 0xee, 0xf6,
	0x4e, 0xe7, 0xef, 0xde, 0xe9, 0x7c, 0xe9, 0x56, 0xdf, 0xf0, 0xf2, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0x2b, 0xa7, 0x9e, 0x94, 0x02, 0x00, 0x00,
}

func (m *ChannelPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EventState != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintData(dAtA, i, uint64(m.EventState))
	}
	if m.DeprecatedSendCount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintData(dAtA, i, uint64(m.DeprecatedSendCount))
	}
	return i, nil
}

func (m *SendCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendCount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SendCount != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintData(dAtA, i, uint64(m.SendCount))
	}
	return i, nil
}

func (m *EventTimePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTimePayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CreateTime != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.CreateTime))
		i += 8
	}
	return i, nil
}

func (m *IndexPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EventId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintData(dAtA, i, uint64(len(m.EventId)))
		i += copy(dAtA[i:], m.EventId)
	}
	if m.CreateTime != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.CreateTime))
		i += 8
	}
	if m.Version != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintData(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func (m *EventPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintData(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if m.DefaultEventState != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintData(dAtA, i, uint64(m.DefaultEventState))
	}
	if len(m.Indexes) > 0 {
		for _, s := range m.Indexes {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChannelPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventState != 0 {
		n += 1 + sovData(uint64(m.EventState))
	}
	if m.DeprecatedSendCount != 0 {
		n += 1 + sovData(uint64(m.DeprecatedSendCount))
	}
	return n
}

func (m *SendCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SendCount != 0 {
		n += 1 + sovData(uint64(m.SendCount))
	}
	return n
}

func (m *EventTimePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreateTime != 0 {
		n += 9
	}
	return n
}

func (m *IndexPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EventId)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.CreateTime != 0 {
		n += 9
	}
	if m.Version != 0 {
		n += 1 + sovData(uint64(m.Version))
	}
	return n
}

func (m *EventPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.DefaultEventState != 0 {
		n += 1 + sovData(uint64(m.DefaultEventState))
	}
	if len(m.Indexes) > 0 {
		for _, s := range m.Indexes {
			l = len(s)
			n += 1 + l + sovData(uint64(l))
		}
	}
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChannelPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChannelPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventState", wireType)
			}
			m.EventState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventState |= EventState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedSendCount", wireType)
			}
			m.DeprecatedSendCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeprecatedSendCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SendCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendCount", wireType)
			}
			m.SendCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventTimePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventTimePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTimePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			m.CreateTime = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.CreateTime = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IndexPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			m.CreateTime = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.CreateTime = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultEventState", wireType)
			}
			m.DefaultEventState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultEventState |= EventState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexes = append(m.Indexes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthData
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowData
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipData(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthData
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthData = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData   = fmt.Errorf("proto: integer overflow")
)
