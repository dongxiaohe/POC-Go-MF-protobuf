// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maxwell_smarts/ticket_events.proto

/*
Package com_zendesk_maxwellsmarts_zendesk_ticketevents is a generated protocol buffer package.

It is generated from these files:
	maxwell_smarts/ticket_events.proto

It has these top-level messages:
	Comment
	TicketCreation
	Assignment
	AssignmentChange
	StatusChange
	TicketEvent
	TicketEvents
*/
package com_zendesk_maxwellsmarts_zendesk_ticketevents

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_zendesk_protobuf "../zendesk"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TicketStatus int32

const (
	TicketStatus_UNKNOWN_TICKET_STATUS TicketStatus = 0
	TicketStatus_NEW                   TicketStatus = 1
	TicketStatus_OPEN                  TicketStatus = 2
	TicketStatus_PENDING               TicketStatus = 3
	TicketStatus_HOLD                  TicketStatus = 4
	TicketStatus_SOLVED                TicketStatus = 5
	TicketStatus_CLOSED                TicketStatus = 6
	TicketStatus_DELETED               TicketStatus = 7
	TicketStatus_ARCHIVED              TicketStatus = 8
)

var TicketStatus_name = map[int32]string{
	0: "UNKNOWN_TICKET_STATUS",
	1: "NEW",
	2: "OPEN",
	3: "PENDING",
	4: "HOLD",
	5: "SOLVED",
	6: "CLOSED",
	7: "DELETED",
	8: "ARCHIVED",
}
var TicketStatus_value = map[string]int32{
	"UNKNOWN_TICKET_STATUS": 0,
	"NEW":      1,
	"OPEN":     2,
	"PENDING":  3,
	"HOLD":     4,
	"SOLVED":   5,
	"CLOSED":   6,
	"DELETED":  7,
	"ARCHIVED": 8,
}

func (x TicketStatus) String() string {
	return proto.EnumName(TicketStatus_name, int32(x))
}
func (TicketStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EventType int32

const (
	EventType_UNKNOWN_EVENT_TYPE EventType = 0
	EventType_TICKET_CREATION    EventType = 1
	EventType_COMMENT_CREATION   EventType = 2
	EventType_STATUS_CHANGE      EventType = 3
	EventType_ASSIGNMENT_CHANGE  EventType = 4
)

var EventType_name = map[int32]string{
	0: "UNKNOWN_EVENT_TYPE",
	1: "TICKET_CREATION",
	2: "COMMENT_CREATION",
	3: "STATUS_CHANGE",
	4: "ASSIGNMENT_CHANGE",
}
var EventType_value = map[string]int32{
	"UNKNOWN_EVENT_TYPE": 0,
	"TICKET_CREATION":    1,
	"COMMENT_CREATION":   2,
	"STATUS_CHANGE":      3,
	"ASSIGNMENT_CHANGE":  4,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Comment struct {
	Text   string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Public bool   `protobuf:"varint,2,opt,name=public" json:"public,omitempty"`
}

func (m *Comment) Reset()                    { *m = Comment{} }
func (m *Comment) String() string            { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()               {}
func (*Comment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Comment) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Comment) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

type TicketCreation struct {
	RequesterId  int64        `protobuf:"varint,1,opt,name=requester_id,json=requesterId" json:"requester_id,omitempty"`
	Status       TicketStatus `protobuf:"varint,2,opt,name=status,enum=com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketStatus" json:"status,omitempty"`
	Comment      *Comment     `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	Assignment   *Assignment  `protobuf:"bytes,4,opt,name=assignment" json:"assignment,omitempty"`
	CommentPart1 string       `protobuf:"bytes,5,opt,name=commentPart1" json:"commentPart1,omitempty"`
	CommentPart2 string       `protobuf:"bytes,6,opt,name=commentPart2" json:"commentPart2,omitempty"`
}

func (m *TicketCreation) Reset()                    { *m = TicketCreation{} }
func (m *TicketCreation) String() string            { return proto.CompactTextString(m) }
func (*TicketCreation) ProtoMessage()               {}
func (*TicketCreation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TicketCreation) GetRequesterId() int64 {
	if m != nil {
		return m.RequesterId
	}
	return 0
}

func (m *TicketCreation) GetStatus() TicketStatus {
	if m != nil {
		return m.Status
	}
	return TicketStatus_UNKNOWN_TICKET_STATUS
}

func (m *TicketCreation) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *TicketCreation) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func (m *TicketCreation) GetCommentPart1() string {
	if m != nil {
		return m.CommentPart1
	}
	return ""
}

func (m *TicketCreation) GetCommentPart2() string {
	if m != nil {
		return m.CommentPart2
	}
	return ""
}

type Assignment struct {
	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
}

func (m *Assignment) Reset()                    { *m = Assignment{} }
func (m *Assignment) String() string            { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()               {}
func (*Assignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Assignment) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Assignment) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

type AssignmentChange struct {
	Previous *Assignment `protobuf:"bytes,1,opt,name=previous" json:"previous,omitempty"`
	Current  *Assignment `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *AssignmentChange) Reset()                    { *m = AssignmentChange{} }
func (m *AssignmentChange) String() string            { return proto.CompactTextString(m) }
func (*AssignmentChange) ProtoMessage()               {}
func (*AssignmentChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AssignmentChange) GetPrevious() *Assignment {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *AssignmentChange) GetCurrent() *Assignment {
	if m != nil {
		return m.Current
	}
	return nil
}

type StatusChange struct {
	Previous TicketStatus `protobuf:"varint,1,opt,name=previous,enum=com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketStatus" json:"previous,omitempty"`
	Current  TicketStatus `protobuf:"varint,2,opt,name=current,enum=com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketStatus" json:"current,omitempty"`
}

func (m *StatusChange) Reset()                    { *m = StatusChange{} }
func (m *StatusChange) String() string            { return proto.CompactTextString(m) }
func (*StatusChange) ProtoMessage()               {}
func (*StatusChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StatusChange) GetPrevious() TicketStatus {
	if m != nil {
		return m.Previous
	}
	return TicketStatus_UNKNOWN_TICKET_STATUS
}

func (m *StatusChange) GetCurrent() TicketStatus {
	if m != nil {
		return m.Current
	}
	return TicketStatus_UNKNOWN_TICKET_STATUS
}

type TicketEvent struct {
	EventType EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,enum=com.zendesk.maxwellsmarts.zendesk.ticketevents.EventType" json:"event_type,omitempty"`
	// Types that are valid to be assigned to EventData:
	//	*TicketEvent_TicketCreation
	//	*TicketEvent_CommentCreation
	//	*TicketEvent_StatusChange
	//	*TicketEvent_AssignmentChange
	EventData isTicketEvent_EventData `protobuf_oneof:"event_data"`
}

func (m *TicketEvent) Reset()                    { *m = TicketEvent{} }
func (m *TicketEvent) String() string            { return proto.CompactTextString(m) }
func (*TicketEvent) ProtoMessage()               {}
func (*TicketEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isTicketEvent_EventData interface {
	isTicketEvent_EventData()
}

type TicketEvent_TicketCreation struct {
	TicketCreation *TicketCreation `protobuf:"bytes,2,opt,name=ticket_creation,json=ticketCreation,oneof"`
}
type TicketEvent_CommentCreation struct {
	CommentCreation *Comment `protobuf:"bytes,3,opt,name=comment_creation,json=commentCreation,oneof"`
}
type TicketEvent_StatusChange struct {
	StatusChange *StatusChange `protobuf:"bytes,4,opt,name=status_change,json=statusChange,oneof"`
}
type TicketEvent_AssignmentChange struct {
	AssignmentChange *AssignmentChange `protobuf:"bytes,5,opt,name=assignment_change,json=assignmentChange,oneof"`
}

func (*TicketEvent_TicketCreation) isTicketEvent_EventData()   {}
func (*TicketEvent_CommentCreation) isTicketEvent_EventData()  {}
func (*TicketEvent_StatusChange) isTicketEvent_EventData()     {}
func (*TicketEvent_AssignmentChange) isTicketEvent_EventData() {}

func (m *TicketEvent) GetEventData() isTicketEvent_EventData {
	if m != nil {
		return m.EventData
	}
	return nil
}

func (m *TicketEvent) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_UNKNOWN_EVENT_TYPE
}

func (m *TicketEvent) GetTicketCreation() *TicketCreation {
	if x, ok := m.GetEventData().(*TicketEvent_TicketCreation); ok {
		return x.TicketCreation
	}
	return nil
}

func (m *TicketEvent) GetCommentCreation() *Comment {
	if x, ok := m.GetEventData().(*TicketEvent_CommentCreation); ok {
		return x.CommentCreation
	}
	return nil
}

func (m *TicketEvent) GetStatusChange() *StatusChange {
	if x, ok := m.GetEventData().(*TicketEvent_StatusChange); ok {
		return x.StatusChange
	}
	return nil
}

func (m *TicketEvent) GetAssignmentChange() *AssignmentChange {
	if x, ok := m.GetEventData().(*TicketEvent_AssignmentChange); ok {
		return x.AssignmentChange
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TicketEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TicketEvent_OneofMarshaler, _TicketEvent_OneofUnmarshaler, _TicketEvent_OneofSizer, []interface{}{
		(*TicketEvent_TicketCreation)(nil),
		(*TicketEvent_CommentCreation)(nil),
		(*TicketEvent_StatusChange)(nil),
		(*TicketEvent_AssignmentChange)(nil),
	}
}

func _TicketEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TicketEvent)
	// event_data
	switch x := m.EventData.(type) {
	case *TicketEvent_TicketCreation:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TicketCreation); err != nil {
			return err
		}
	case *TicketEvent_CommentCreation:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CommentCreation); err != nil {
			return err
		}
	case *TicketEvent_StatusChange:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StatusChange); err != nil {
			return err
		}
	case *TicketEvent_AssignmentChange:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AssignmentChange); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TicketEvent.EventData has unexpected type %T", x)
	}
	return nil
}

func _TicketEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TicketEvent)
	switch tag {
	case 2: // event_data.ticket_creation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketCreation)
		err := b.DecodeMessage(msg)
		m.EventData = &TicketEvent_TicketCreation{msg}
		return true, err
	case 3: // event_data.comment_creation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Comment)
		err := b.DecodeMessage(msg)
		m.EventData = &TicketEvent_CommentCreation{msg}
		return true, err
	case 4: // event_data.status_change
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StatusChange)
		err := b.DecodeMessage(msg)
		m.EventData = &TicketEvent_StatusChange{msg}
		return true, err
	case 5: // event_data.assignment_change
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AssignmentChange)
		err := b.DecodeMessage(msg)
		m.EventData = &TicketEvent_AssignmentChange{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TicketEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TicketEvent)
	// event_data
	switch x := m.EventData.(type) {
	case *TicketEvent_TicketCreation:
		s := proto.Size(x.TicketCreation)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketEvent_CommentCreation:
		s := proto.Size(x.CommentCreation)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketEvent_StatusChange:
		s := proto.Size(x.StatusChange)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketEvent_AssignmentChange:
		s := proto.Size(x.AssignmentChange)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TicketEvents struct {
	Header    *com_zendesk_protobuf.ProtobufHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AccountId int32                                `protobuf:"varint,5,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ActorId   int64                                `protobuf:"varint,6,opt,name=actor_id,json=actorId" json:"actor_id,omitempty"`
	TicketId  int64                                `protobuf:"varint,7,opt,name=ticket_id,json=ticketId" json:"ticket_id,omitempty"`
	Events    []*TicketEvent                       `protobuf:"bytes,8,rep,name=events" json:"events,omitempty"`
}

func (m *TicketEvents) Reset()                    { *m = TicketEvents{} }
func (m *TicketEvents) String() string            { return proto.CompactTextString(m) }
func (*TicketEvents) ProtoMessage()               {}
func (*TicketEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TicketEvents) GetHeader() *com_zendesk_protobuf.ProtobufHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TicketEvents) GetAccountId() int32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *TicketEvents) GetActorId() int64 {
	if m != nil {
		return m.ActorId
	}
	return 0
}

func (m *TicketEvents) GetTicketId() int64 {
	if m != nil {
		return m.TicketId
	}
	return 0
}

func (m *TicketEvents) GetEvents() []*TicketEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Comment)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.Comment")
	proto.RegisterType((*TicketCreation)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketCreation")
	proto.RegisterType((*Assignment)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.Assignment")
	proto.RegisterType((*AssignmentChange)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.AssignmentChange")
	proto.RegisterType((*StatusChange)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.StatusChange")
	proto.RegisterType((*TicketEvent)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketEvent")
	proto.RegisterType((*TicketEvents)(nil), "com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketEvents")
	proto.RegisterEnum("com.zendesk.maxwellsmarts.zendesk.ticketevents.TicketStatus", TicketStatus_name, TicketStatus_value)
	proto.RegisterEnum("com.zendesk.maxwellsmarts.zendesk.ticketevents.EventType", EventType_name, EventType_value)
}

func init() { proto.RegisterFile("maxwell_smarts/ticket_events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xe3, 0x7c, 0xd8, 0xce, 0x49, 0xb6, 0x9d, 0x0e, 0x6c, 0xc9, 0x82, 0x90, 0x42, 0xc4,
	0x45, 0xb4, 0x17, 0x59, 0x11, 0x84, 0x10, 0xb0, 0x42, 0x9b, 0x75, 0xac, 0xc6, 0x6c, 0xd6, 0x09,
	0x63, 0x6f, 0xb6, 0x70, 0x63, 0xb9, 0xf6, 0xd0, 0x5a, 0x6d, 0xe2, 0x60, 0x8f, 0x4b, 0xcb, 0x03,
	0x20, 0xf1, 0x0e, 0x3c, 0x0c, 0x4f, 0x86, 0x90, 0x67, 0xc6, 0x4e, 0xd2, 0xbd, 0x4a, 0x73, 0x37,
	0x73, 0xe6, 0xcc, 0xef, 0xcc, 0x7f, 0xce, 0x07, 0xf4, 0x96, 0xfe, 0xdd, 0x1f, 0xf4, 0xe6, 0xc6,
	0x4b, 0x97, 0x7e, 0xc2, 0xd2, 0x17, 0x2c, 0x0a, 0xae, 0x29, 0xf3, 0xe8, 0x2d, 0x5d, 0xb1, 0x74,
	0xb0, 0x4e, 0x62, 0x16, 0xe3, 0x41, 0x10, 0x2f, 0x07, 0x7f, 0xd2, 0x55, 0x48, 0xd3, 0xeb, 0x81,
	0xf4, 0x17, 0xee, 0xa5, 0x55, 0x5c, 0x13, 0xb7, 0x3e, 0x3d, 0x95, 0xd6, 0x17, 0xfc, 0xfa, 0x45,
	0xf6, 0x9b, 0xe0, 0xf4, 0xbe, 0x01, 0xcd, 0x88, 0x97, 0x4b, 0xba, 0x62, 0x18, 0x43, 0x9d, 0xd1,
	0x3b, 0xd6, 0x51, 0xba, 0x4a, 0xbf, 0x49, 0xf8, 0x1a, 0x9f, 0x82, 0xba, 0xce, 0x2e, 0x6e, 0xa2,
	0xa0, 0x53, 0xed, 0x2a, 0x7d, 0x9d, 0xc8, 0x5d, 0xef, 0xef, 0x1a, 0x1c, 0xb9, 0x9c, 0x6f, 0x24,
	0xd4, 0x67, 0x51, 0xbc, 0xc2, 0x5f, 0x40, 0x3b, 0xa1, 0xbf, 0x67, 0x34, 0x65, 0x34, 0xf1, 0xa2,
	0x90, 0x63, 0x6a, 0xa4, 0x55, 0xda, 0xac, 0x10, 0xbb, 0xa0, 0xa6, 0xcc, 0x67, 0x59, 0xca, 0x69,
	0x47, 0xc3, 0x97, 0x7b, 0xaa, 0x18, 0x88, 0x90, 0x0e, 0x67, 0x10, 0xc9, 0xc2, 0x3f, 0x83, 0x16,
	0x08, 0x09, 0x9d, 0x5a, 0x57, 0xe9, 0xb7, 0x86, 0xdf, 0xee, 0x8b, 0x95, 0x3f, 0x40, 0x0a, 0x0e,
	0xfe, 0x15, 0xc0, 0x4f, 0xd3, 0xe8, 0x72, 0xc5, 0xa9, 0x75, 0x4e, 0xfd, 0x7e, 0x5f, 0xea, 0xa8,
	0x24, 0x90, 0x2d, 0x1a, 0xee, 0x41, 0x5b, 0x86, 0x99, 0xfb, 0x09, 0xfb, 0xaa, 0xd3, 0xe0, 0xdf,
	0xbd, 0x63, 0x7b, 0xe0, 0x33, 0xec, 0xa8, 0x1f, 0xf8, 0x0c, 0x49, 0x91, 0xb7, 0xde, 0x2b, 0x80,
	0x4d, 0x28, 0xfc, 0x09, 0x68, 0x59, 0xba, 0x9d, 0x01, 0x35, 0xdf, 0x5a, 0x21, 0x7e, 0x06, 0xfa,
	0x65, 0x12, 0x67, 0xeb, 0xfc, 0xa4, 0xca, 0x4f, 0x34, 0xbe, 0xb7, 0xc2, 0xde, 0xbf, 0x0a, 0xa0,
	0x0d, 0xc2, 0xb8, 0xf2, 0x57, 0x97, 0x14, 0x2f, 0x40, 0x5f, 0x27, 0xf4, 0x36, 0x8a, 0xb3, 0x94,
	0x93, 0x0e, 0xfb, 0x81, 0x92, 0x85, 0x5d, 0xd0, 0x82, 0x2c, 0x49, 0xf2, 0x8f, 0xad, 0x1e, 0x8c,
	0x2d, 0x50, 0xb9, 0x84, 0xb6, 0xa8, 0x0b, 0xf9, 0xfc, 0xf3, 0x07, 0xcf, 0x3f, 0xb4, 0xda, 0x36,
	0x02, 0x16, 0xbb, 0x02, 0x0e, 0x05, 0x97, 0x12, 0xfe, 0xa9, 0x43, 0x4b, 0x9c, 0x98, 0xb9, 0x1b,
	0x3e, 0x07, 0xe0, 0xfe, 0x1e, 0xbb, 0x5f, 0x53, 0xa9, 0xe1, 0xbb, 0x7d, 0x43, 0x71, 0x94, 0x7b,
	0xbf, 0xa6, 0xa4, 0x49, 0x8b, 0x25, 0x8e, 0xe0, 0x58, 0xce, 0x94, 0x40, 0x76, 0xaf, 0x4c, 0xc5,
	0x8f, 0x8f, 0x53, 0x52, 0xcc, 0x80, 0x49, 0x85, 0x1c, 0xb1, 0xdd, 0xa9, 0x10, 0x02, 0x92, 0x55,
	0xbb, 0x89, 0x75, 0x58, 0x97, 0x4e, 0x2a, 0xe4, 0x58, 0x22, 0xcb, 0x28, 0x01, 0x3c, 0x11, 0xc3,
	0xc0, 0x0b, 0x78, 0xf6, 0x65, 0xcb, 0xee, 0x9d, 0x98, 0xed, 0x0a, 0x9a, 0x54, 0x48, 0x3b, 0xdd,
	0xae, 0xa8, 0x18, 0x4e, 0x36, 0x6d, 0x5c, 0x04, 0x6a, 0xf0, 0x40, 0xaf, 0x1e, 0x5f, 0xc2, 0x65,
	0x30, 0xe4, 0x3f, 0xb0, 0xbd, 0x6e, 0x17, 0x05, 0x10, 0xfa, 0xcc, 0xef, 0xfd, 0xa7, 0x40, 0x7b,
	0xab, 0x3c, 0x52, 0xfc, 0x12, 0xd4, 0x2b, 0xea, 0x87, 0x34, 0x91, 0xed, 0xf9, 0xe5, 0xce, 0x23,
	0xca, 0x39, 0x3f, 0x97, 0x8b, 0x09, 0xf7, 0x25, 0xf2, 0x0e, 0xfe, 0x1c, 0xc0, 0x0f, 0x82, 0x38,
	0x5b, 0xb1, 0x7c, 0x20, 0xe4, 0x32, 0x1a, 0xa4, 0x29, 0x2d, 0x62, 0x5a, 0xf8, 0x01, 0x8b, 0xf9,
	0x1c, 0x51, 0xc5, 0xb4, 0xe0, 0x7b, 0x2b, 0xc4, 0x9f, 0x41, 0x53, 0x56, 0x4f, 0x14, 0x76, 0x34,
	0x7e, 0xa6, 0x0b, 0x83, 0x15, 0x62, 0x07, 0x54, 0x21, 0xb1, 0xa3, 0x77, 0x6b, 0xfd, 0xd6, 0xf0,
	0x87, 0xc7, 0x55, 0x14, 0x97, 0x48, 0x24, 0xea, 0xa7, 0xba, 0x5e, 0x45, 0x8d, 0xe7, 0x7f, 0x95,
	0x1f, 0x20, 0xd2, 0x84, 0x9f, 0xc1, 0xd3, 0x77, 0xf6, 0x1b, 0x7b, 0xf6, 0xde, 0xf6, 0x5c, 0xcb,
	0x78, 0x63, 0xba, 0x9e, 0xe3, 0x8e, 0xdc, 0x77, 0x0e, 0xaa, 0x60, 0x0d, 0x6a, 0xb6, 0xf9, 0x1e,
	0x29, 0x58, 0x87, 0xfa, 0x6c, 0x6e, 0xda, 0xa8, 0x8a, 0x5b, 0xa0, 0xcd, 0x4d, 0x7b, 0x6c, 0xd9,
	0x67, 0xa8, 0x96, 0x9b, 0x27, 0xb3, 0xe9, 0x18, 0xd5, 0x31, 0x80, 0xea, 0xcc, 0xa6, 0x0b, 0x73,
	0x8c, 0x1a, 0xf9, 0xda, 0x98, 0xce, 0x1c, 0x73, 0x8c, 0xd4, 0xdc, 0x7d, 0x6c, 0x4e, 0x4d, 0xd7,
	0x1c, 0x23, 0x0d, 0xb7, 0x41, 0x1f, 0x11, 0x63, 0x62, 0xe5, 0x6e, 0xfa, 0xf3, 0x3b, 0x68, 0x96,
	0x6d, 0x85, 0x4f, 0x01, 0x17, 0x8f, 0x30, 0x17, 0xa6, 0xed, 0x7a, 0xee, 0x2f, 0x73, 0x13, 0x55,
	0xf0, 0x47, 0x70, 0x2c, 0x1f, 0x65, 0x10, 0x73, 0xe4, 0x5a, 0x33, 0x1b, 0x29, 0xf8, 0x63, 0x40,
	0xc6, 0xec, 0xed, 0xdb, 0xdc, 0xad, 0xb4, 0x56, 0xf1, 0x09, 0x3c, 0x11, 0x0f, 0xf7, 0x8c, 0xc9,
	0xc8, 0x3e, 0x33, 0x51, 0x0d, 0x3f, 0x85, 0x93, 0x91, 0xe3, 0x58, 0x67, 0xb6, 0xf0, 0x15, 0xe6,
	0xfa, 0xeb, 0xea, 0x5c, 0xb9, 0x50, 0x79, 0x66, 0xbf, 0xfe, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x1e,
	0x73, 0x0c, 0x37, 0x26, 0x08, 0x00, 0x00,
}