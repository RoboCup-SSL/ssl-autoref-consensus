// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rcon.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	rcon.proto
	referee.proto

It has these top-level messages:
	SSL_RefereeRemoteControlRequest
	SSL_RefereeRemoteControlReply
	SSL_Referee
*/
package main

import proto "github.com/golang/protobuf/proto"
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

// Which type of card to issue.
type SSL_RefereeRemoteControlRequest_CardInfo_CardType int32

const (
	SSL_RefereeRemoteControlRequest_CardInfo_CARD_YELLOW SSL_RefereeRemoteControlRequest_CardInfo_CardType = 0
	SSL_RefereeRemoteControlRequest_CardInfo_CARD_RED    SSL_RefereeRemoteControlRequest_CardInfo_CardType = 1
)

var SSL_RefereeRemoteControlRequest_CardInfo_CardType_name = map[int32]string{
	0: "CARD_YELLOW",
	1: "CARD_RED",
}
var SSL_RefereeRemoteControlRequest_CardInfo_CardType_value = map[string]int32{
	"CARD_YELLOW": 0,
	"CARD_RED":    1,
}

func (x SSL_RefereeRemoteControlRequest_CardInfo_CardType) Enum() *SSL_RefereeRemoteControlRequest_CardInfo_CardType {
	p := new(SSL_RefereeRemoteControlRequest_CardInfo_CardType)
	*p = x
	return p
}
func (x SSL_RefereeRemoteControlRequest_CardInfo_CardType) String() string {
	return proto.EnumName(SSL_RefereeRemoteControlRequest_CardInfo_CardType_name, int32(x))
}
func (x *SSL_RefereeRemoteControlRequest_CardInfo_CardType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_RefereeRemoteControlRequest_CardInfo_CardType_value, data, "SSL_RefereeRemoteControlRequest_CardInfo_CardType")
	if err != nil {
		return err
	}
	*x = SSL_RefereeRemoteControlRequest_CardInfo_CardType(value)
	return nil
}
func (SSL_RefereeRemoteControlRequest_CardInfo_CardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

// Which team to issue the card to.
type SSL_RefereeRemoteControlRequest_CardInfo_CardTeam int32

const (
	SSL_RefereeRemoteControlRequest_CardInfo_TEAM_YELLOW SSL_RefereeRemoteControlRequest_CardInfo_CardTeam = 0
	SSL_RefereeRemoteControlRequest_CardInfo_TEAM_BLUE   SSL_RefereeRemoteControlRequest_CardInfo_CardTeam = 1
)

var SSL_RefereeRemoteControlRequest_CardInfo_CardTeam_name = map[int32]string{
	0: "TEAM_YELLOW",
	1: "TEAM_BLUE",
}
var SSL_RefereeRemoteControlRequest_CardInfo_CardTeam_value = map[string]int32{
	"TEAM_YELLOW": 0,
	"TEAM_BLUE":   1,
}

func (x SSL_RefereeRemoteControlRequest_CardInfo_CardTeam) Enum() *SSL_RefereeRemoteControlRequest_CardInfo_CardTeam {
	p := new(SSL_RefereeRemoteControlRequest_CardInfo_CardTeam)
	*p = x
	return p
}
func (x SSL_RefereeRemoteControlRequest_CardInfo_CardTeam) String() string {
	return proto.EnumName(SSL_RefereeRemoteControlRequest_CardInfo_CardTeam_name, int32(x))
}
func (x *SSL_RefereeRemoteControlRequest_CardInfo_CardTeam) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_RefereeRemoteControlRequest_CardInfo_CardTeam_value, data, "SSL_RefereeRemoteControlRequest_CardInfo_CardTeam")
	if err != nil {
		return err
	}
	*x = SSL_RefereeRemoteControlRequest_CardInfo_CardTeam(value)
	return nil
}
func (SSL_RefereeRemoteControlRequest_CardInfo_CardTeam) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1}
}

// The outcome of the request.
type SSL_RefereeRemoteControlReply_Outcome int32

const (
	// The request was accepted.
	SSL_RefereeRemoteControlReply_OK SSL_RefereeRemoteControlReply_Outcome = 0
	// The request was rejected because it contained more than one action.
	SSL_RefereeRemoteControlReply_MULTIPLE_ACTIONS SSL_RefereeRemoteControlReply_Outcome = 1
	// The request was rejected because the requested stage does not exist,
	// is not accessible from the current game state, or is a game half
	// (which must be entered by means of a NORMAL_START command, not a
	// stage change).
	SSL_RefereeRemoteControlReply_BAD_STAGE SSL_RefereeRemoteControlReply_Outcome = 2
	// The request was rejected because the requested command does not
	// exist or is not accessible from the current game state.
	SSL_RefereeRemoteControlReply_BAD_COMMAND SSL_RefereeRemoteControlReply_Outcome = 3
	// The request was rejected because a designated position was provided
	// with a command that does not need one or with the command field
	// absent, or the command field was set to a ball placement command but
	// a designated position was not provided.
	SSL_RefereeRemoteControlReply_BAD_DESIGNATED_POSITION SSL_RefereeRemoteControlReply_Outcome = 4
	// The request was rejected because the command_counter field does not
	// match.
	SSL_RefereeRemoteControlReply_BAD_COMMAND_COUNTER SSL_RefereeRemoteControlReply_Outcome = 5
	// The request was rejected because a card cannot be issued at this
	// time.
	SSL_RefereeRemoteControlReply_BAD_CARD    SSL_RefereeRemoteControlReply_Outcome = 6
	SSL_RefereeRemoteControlReply_NO_MAJORITY SSL_RefereeRemoteControlReply_Outcome = 7
)

var SSL_RefereeRemoteControlReply_Outcome_name = map[int32]string{
	0: "OK",
	1: "MULTIPLE_ACTIONS",
	2: "BAD_STAGE",
	3: "BAD_COMMAND",
	4: "BAD_DESIGNATED_POSITION",
	5: "BAD_COMMAND_COUNTER",
	6: "BAD_CARD",
	7: "NO_MAJORITY",
}
var SSL_RefereeRemoteControlReply_Outcome_value = map[string]int32{
	"OK":                      0,
	"MULTIPLE_ACTIONS":        1,
	"BAD_STAGE":               2,
	"BAD_COMMAND":             3,
	"BAD_DESIGNATED_POSITION": 4,
	"BAD_COMMAND_COUNTER":     5,
	"BAD_CARD":                6,
	"NO_MAJORITY":             7,
}

func (x SSL_RefereeRemoteControlReply_Outcome) Enum() *SSL_RefereeRemoteControlReply_Outcome {
	p := new(SSL_RefereeRemoteControlReply_Outcome)
	*p = x
	return p
}
func (x SSL_RefereeRemoteControlReply_Outcome) String() string {
	return proto.EnumName(SSL_RefereeRemoteControlReply_Outcome_name, int32(x))
}
func (x *SSL_RefereeRemoteControlReply_Outcome) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_RefereeRemoteControlReply_Outcome_value, data, "SSL_RefereeRemoteControlReply_Outcome")
	if err != nil {
		return err
	}
	*x = SSL_RefereeRemoteControlReply_Outcome(value)
	return nil
}
func (SSL_RefereeRemoteControlReply_Outcome) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// The TCP half-connection from controller to referee box carries a sequence of
// these messages, sent on demand, each preceded by its length in bytes as a
// 4-byte big-endian integer.
//
// There is no rule on how often requests can be, or need to be, sent. A remote
// control client connected to the referee box over a reliable link will
// typically send messages only when a change of game state is needed. However,
// a remote control client connected over an unreliable link (e.g. wifi) may
// wish to send a no-op request (a message with only message_id filled in) at a
// fixed frequency to detect and report connection failures. Such no-op
// requests cause no change to game state but solicit a reply from the referee
// box with the OK outcome.
//
// Each request may contain at most one action. An action is a stage, a
// command, or a card. Setting more than one action in a single request results
// in the request being rejected. This simplifies understanding the order in
// which actions occur.
type SSL_RefereeRemoteControlRequest struct {
	// The message ID. This number may be selected arbitrarily by the client.
	// It is never interpreted by the referee box. It is returned unmodified in
	// the SSL_RefereeRemoteControlReply message to allow correlating replies
	// to requests.
	MessageId *uint32 `protobuf:"varint,1,req,name=message_id,json=messageId" json:"message_id,omitempty"`
	// The stage of the game to move to, which should be omitted if the stage
	// need not change.
	//
	// Do not use this field to enter the NORMAL_FIRST_HALF,
	// NORMAL_SECOND_HALF, EXTRA_FIRST_HALF, or EXTRA_SECOND_HALF stages.
	// Instead, prepare a kickoff and then issue the NORMAL_START command via
	// the command field.
	Stage *SSL_Referee_Stage `protobuf:"varint,2,opt,name=stage,enum=SSL_Referee_Stage" json:"stage,omitempty"`
	// The command to be issued, which should be omitted if no command should
	// be issued at this time (i.e. the command currently in force may continue
	// to be used, or a stage change is occurring which comes with its own new
	// command). Sending a request with this field set always increments the
	// command counter in the broadcast referee box packet, which implies a
	// “new activity”.
	//
	// The TIMEOUT_YELLOW and TIMEOUT_BLUE commands must be used to report
	// entering timeouts. The timeout counter and clock are updated when those
	// commands are issued. STOP is used to end the timeout.
	//
	// HALT can be issued during a timeout to stop the timeout clock. The
	// appropriate choice of TIMEOUT_YELLOW or TIMEOUT_BLUE can be issued to
	// resume the timeout clock, or STOP can be issued to end the timeout.
	//
	// The GOAL_YELLOW and GOAL_BLUE commands must be used to report goals
	// scored. The goal counters in the TeamInfo messages are incremented when
	// those commands are issued.
	Command *SSL_Referee_Command `protobuf:"varint,3,opt,name=command,enum=SSL_Referee_Command" json:"command,omitempty"`
	// The coordinates of the Designated Position (whose purpose depends on the
	// current command). This is measured in millimetres and correspond to
	// SSL-Vision coordinates. this field must be present if and only if the
	// command field is present and is set to a ball placement command. If the
	// command field is absent, the designated position does not change. If the
	// command field is present and the designated position field is absent, no
	// designated position is included in the new command.
	DesignatedPosition *SSL_Referee_Point                        `protobuf:"bytes,4,opt,name=designated_position,json=designatedPosition" json:"designated_position,omitempty"`
	Card               *SSL_RefereeRemoteControlRequest_CardInfo `protobuf:"bytes,5,opt,name=card" json:"card,omitempty"`
	// The command_counter of the most recent multicast referee packet observed
	// by the remote control. If this does not match the command_counter value
	// of the current referee packet, the request is rejected.
	//
	// The purpose of this field is to avoid race conditions between input from
	// a human and input from an autonomous software referee using the remote
	// control protocol. For example, consider the case where the game is in
	// play (Normal Start) and the human operator presses the Halt button while
	// the autonomous referee simultaneously sends the Stop command. Without
	// this field:
	// 1. The Halt button is pressed, resulting in a new command being started.
	// 2. The autonomous referee sends the Stop command to the remote control
	//    port.
	// 3. The Stop command is accepted and takes precedence due to arriving
	//    later.
	// In the worst case, this could result in the operator having to press
	// Halt multiple times before the field is made safe.
	//
	// By including the command_counter field in all requests on the remote
	// control port, and autonomous referee avoids this problem. Instead, the
	// situation would develop as follows:
	// 1. The Halt button is pressed, resulting in a new command being started.
	// 2. The autonomous referee sends the Stop command, but with the Normal
	//    Start command’s command_counter.
	// 3. The referee box rejects the Stop command and remains halted, due to
	//    the mismatch in command_counter value.
	// 4. The autonomous referee waits for the next multicast packet from the
	//    referee box, to get the new command_counter value.
	// 5. The autonomous referee box observes that the game is halted and
	//    decides it should not send the Stop command.
	//
	// If this field is omitted, the check for a matching counter value is
	// bypassed and requests are never rejected for this reason. This is
	// appropriate for remote control software managed by a human operator,
	// where no race condition is possible.
	LastCommandCounter *uint32 `protobuf:"varint,6,opt,name=last_command_counter,json=lastCommandCounter" json:"last_command_counter,omitempty"`
	ImplId             *string `protobuf:"bytes,7,opt,name=impl_id,json=implId" json:"impl_id,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *SSL_RefereeRemoteControlRequest) Reset()                    { *m = SSL_RefereeRemoteControlRequest{} }
func (m *SSL_RefereeRemoteControlRequest) String() string            { return proto.CompactTextString(m) }
func (*SSL_RefereeRemoteControlRequest) ProtoMessage()               {}
func (*SSL_RefereeRemoteControlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SSL_RefereeRemoteControlRequest) GetMessageId() uint32 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *SSL_RefereeRemoteControlRequest) GetStage() SSL_Referee_Stage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return SSL_Referee_NORMAL_FIRST_HALF_PRE
}

func (m *SSL_RefereeRemoteControlRequest) GetCommand() SSL_Referee_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return SSL_Referee_HALT
}

func (m *SSL_RefereeRemoteControlRequest) GetDesignatedPosition() *SSL_Referee_Point {
	if m != nil {
		return m.DesignatedPosition
	}
	return nil
}

func (m *SSL_RefereeRemoteControlRequest) GetCard() *SSL_RefereeRemoteControlRequest_CardInfo {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *SSL_RefereeRemoteControlRequest) GetLastCommandCounter() uint32 {
	if m != nil && m.LastCommandCounter != nil {
		return *m.LastCommandCounter
	}
	return 0
}

func (m *SSL_RefereeRemoteControlRequest) GetImplId() string {
	if m != nil && m.ImplId != nil {
		return *m.ImplId
	}
	return ""
}

// The card to issue.
type SSL_RefereeRemoteControlRequest_CardInfo struct {
	Type             *SSL_RefereeRemoteControlRequest_CardInfo_CardType `protobuf:"varint,1,req,name=type,enum=SSL_RefereeRemoteControlRequest_CardInfo_CardType" json:"type,omitempty"`
	Team             *SSL_RefereeRemoteControlRequest_CardInfo_CardTeam `protobuf:"varint,2,req,name=team,enum=SSL_RefereeRemoteControlRequest_CardInfo_CardTeam" json:"team,omitempty"`
	XXX_unrecognized []byte                                             `json:"-"`
}

func (m *SSL_RefereeRemoteControlRequest_CardInfo) Reset() {
	*m = SSL_RefereeRemoteControlRequest_CardInfo{}
}
func (m *SSL_RefereeRemoteControlRequest_CardInfo) String() string { return proto.CompactTextString(m) }
func (*SSL_RefereeRemoteControlRequest_CardInfo) ProtoMessage()    {}
func (*SSL_RefereeRemoteControlRequest_CardInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *SSL_RefereeRemoteControlRequest_CardInfo) GetType() SSL_RefereeRemoteControlRequest_CardInfo_CardType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SSL_RefereeRemoteControlRequest_CardInfo_CARD_YELLOW
}

func (m *SSL_RefereeRemoteControlRequest_CardInfo) GetTeam() SSL_RefereeRemoteControlRequest_CardInfo_CardTeam {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return SSL_RefereeRemoteControlRequest_CardInfo_TEAM_YELLOW
}

// The TCP half-connection from referee box to controller carries a sequence of
// these messages, sent precisely once per received
// SSL_RefereeRemoteControlRequest message, each preceded by its length in
// bytes as as 4-byte big-endian integer.
type SSL_RefereeRemoteControlReply struct {
	// The message ID of the request message to which this reply corresponds.
	MessageId        *uint32                                `protobuf:"varint,1,req,name=message_id,json=messageId" json:"message_id,omitempty"`
	Outcome          *SSL_RefereeRemoteControlReply_Outcome `protobuf:"varint,2,req,name=outcome,enum=SSL_RefereeRemoteControlReply_Outcome" json:"outcome,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *SSL_RefereeRemoteControlReply) Reset()                    { *m = SSL_RefereeRemoteControlReply{} }
func (m *SSL_RefereeRemoteControlReply) String() string            { return proto.CompactTextString(m) }
func (*SSL_RefereeRemoteControlReply) ProtoMessage()               {}
func (*SSL_RefereeRemoteControlReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SSL_RefereeRemoteControlReply) GetMessageId() uint32 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *SSL_RefereeRemoteControlReply) GetOutcome() SSL_RefereeRemoteControlReply_Outcome {
	if m != nil && m.Outcome != nil {
		return *m.Outcome
	}
	return SSL_RefereeRemoteControlReply_OK
}

func init() {
	proto.RegisterType((*SSL_RefereeRemoteControlRequest)(nil), "SSL_RefereeRemoteControlRequest")
	proto.RegisterType((*SSL_RefereeRemoteControlRequest_CardInfo)(nil), "SSL_RefereeRemoteControlRequest.CardInfo")
	proto.RegisterType((*SSL_RefereeRemoteControlReply)(nil), "SSL_RefereeRemoteControlReply")
	proto.RegisterEnum("SSL_RefereeRemoteControlRequest_CardInfo_CardType", SSL_RefereeRemoteControlRequest_CardInfo_CardType_name, SSL_RefereeRemoteControlRequest_CardInfo_CardType_value)
	proto.RegisterEnum("SSL_RefereeRemoteControlRequest_CardInfo_CardTeam", SSL_RefereeRemoteControlRequest_CardInfo_CardTeam_name, SSL_RefereeRemoteControlRequest_CardInfo_CardTeam_value)
	proto.RegisterEnum("SSL_RefereeRemoteControlReply_Outcome", SSL_RefereeRemoteControlReply_Outcome_name, SSL_RefereeRemoteControlReply_Outcome_value)
}

func init() { proto.RegisterFile("rcon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6a, 0xdb, 0x4c,
	0x14, 0x85, 0x23, 0xc5, 0xb6, 0xe2, 0x9b, 0xdf, 0xf9, 0x87, 0x89, 0xc1, 0x22, 0x25, 0xd4, 0x78,
	0x51, 0x94, 0x2e, 0x44, 0xf1, 0xbe, 0x50, 0x45, 0x52, 0x83, 0x5a, 0xd9, 0x32, 0x23, 0x99, 0x92,
	0xd5, 0x20, 0xac, 0x89, 0x11, 0x58, 0x1a, 0x55, 0x1a, 0x2f, 0xfc, 0x2e, 0x85, 0x3e, 0x59, 0x1f,
	0xa2, 0x6f, 0x50, 0x66, 0x24, 0x93, 0xb4, 0xd0, 0xa6, 0x74, 0xe7, 0x39, 0xf7, 0x3b, 0xc7, 0xe7,
	0x5e, 0x04, 0x50, 0x6f, 0x78, 0x69, 0x57, 0x35, 0x17, 0xfc, 0x6a, 0x54, 0xb3, 0x07, 0x56, 0x33,
	0xd6, 0x3e, 0x67, 0xdf, 0x7a, 0xf0, 0x32, 0x8e, 0x43, 0x4a, 0x5a, 0x95, 0xb0, 0x82, 0x0b, 0xe6,
	0xf2, 0x52, 0xd4, 0x7c, 0x47, 0xd8, 0xe7, 0x3d, 0x6b, 0x04, 0xbe, 0x06, 0x28, 0x58, 0xd3, 0xa4,
	0x5b, 0x46, 0xf3, 0xcc, 0xd4, 0xa6, 0xba, 0x35, 0x22, 0xc3, 0x4e, 0x09, 0x32, 0x6c, 0x41, 0xbf,
	0x11, 0xe9, 0x96, 0x99, 0xfa, 0x54, 0xb3, 0x2e, 0xe6, 0xd8, 0x7e, 0x92, 0x67, 0xc7, 0x72, 0x42,
	0x5a, 0x00, 0xdb, 0x60, 0x6c, 0x78, 0x51, 0xa4, 0x65, 0x66, 0x9e, 0x2a, 0x76, 0xfc, 0x13, 0xeb,
	0xb6, 0x33, 0x72, 0x84, 0xb0, 0x0b, 0x97, 0x19, 0x6b, 0xf2, 0x6d, 0x99, 0x0a, 0x96, 0xd1, 0x8a,
	0x37, 0xb9, 0xc8, 0x79, 0x69, 0xf6, 0xa6, 0x9a, 0x75, 0xfe, 0xcb, 0xff, 0xac, 0x78, 0x5e, 0x0a,
	0x82, 0x1f, 0xf1, 0x55, 0x47, 0xe3, 0xb7, 0xd0, 0xdb, 0xa4, 0x75, 0x66, 0xf6, 0x95, 0xeb, 0xc6,
	0x7e, 0x66, 0x5b, 0xdb, 0x4d, 0xeb, 0x2c, 0x28, 0x1f, 0x38, 0x51, 0x36, 0xfc, 0x06, 0xc6, 0xbb,
	0xb4, 0x11, 0xb4, 0xeb, 0x44, 0x37, 0x7c, 0x5f, 0x0a, 0x56, 0x9b, 0x83, 0xa9, 0x66, 0x8d, 0x08,
	0x96, 0xb3, 0xae, 0xb7, 0xdb, 0x4e, 0xf0, 0x04, 0x8c, 0xbc, 0xa8, 0x76, 0xf2, 0x56, 0xc6, 0x54,
	0xb3, 0x86, 0x64, 0x20, 0x9f, 0x41, 0x76, 0xf5, 0x5d, 0x83, 0xb3, 0x63, 0x3a, 0x7e, 0x0f, 0x3d,
	0x71, 0xa8, 0x98, 0x3a, 0xe7, 0xc5, 0x7c, 0xfe, 0xd7, 0xb5, 0xd4, 0x8f, 0xe4, 0x50, 0x31, 0xa2,
	0xfc, 0x2a, 0x87, 0xa5, 0x85, 0xa9, 0xff, 0x53, 0x0e, 0x4b, 0x0b, 0xa2, 0xfc, 0xb3, 0x9b, 0xb6,
	0x9b, 0x4c, 0xc6, 0xff, 0xc3, 0xb9, 0xeb, 0x10, 0x8f, 0xde, 0xfb, 0x61, 0x18, 0x7d, 0x42, 0x27,
	0xf8, 0x3f, 0x38, 0x53, 0x02, 0xf1, 0x3d, 0xa4, 0xcd, 0x5e, 0x77, 0x28, 0x4b, 0x0b, 0x89, 0x26,
	0xbe, 0xb3, 0x78, 0x44, 0x47, 0x30, 0x54, 0xc2, 0x6d, 0xb8, 0xf6, 0x91, 0x36, 0xfb, 0xaa, 0xc3,
	0xf5, 0xef, 0x2b, 0x55, 0xbb, 0xc3, 0x73, 0x5f, 0xd7, 0x3b, 0x30, 0xf8, 0x5e, 0x6c, 0x78, 0xc1,
	0xba, 0x15, 0x5f, 0xd9, 0x7f, 0xcc, 0xb3, 0xa3, 0x96, 0x26, 0x47, 0xdb, 0xec, 0x8b, 0x06, 0x46,
	0x27, 0xe2, 0x01, 0xe8, 0xd1, 0x47, 0x74, 0x82, 0xc7, 0x80, 0x16, 0xeb, 0x30, 0x09, 0x56, 0xa1,
	0x4f, 0x1d, 0x37, 0x09, 0xa2, 0x65, 0x8c, 0x34, 0xd9, 0xfd, 0xd6, 0xf1, 0x68, 0x9c, 0x38, 0x77,
	0x3e, 0xd2, 0xe5, 0x6e, 0xf2, 0xe9, 0x46, 0x8b, 0x85, 0xb3, 0xf4, 0xd0, 0x29, 0x7e, 0x01, 0x13,
	0x29, 0x78, 0x7e, 0x1c, 0xdc, 0x2d, 0x9d, 0xc4, 0xf7, 0xe8, 0x2a, 0x8a, 0x03, 0xe9, 0x46, 0x3d,
	0x3c, 0x81, 0xcb, 0x27, 0x34, 0x75, 0xa3, 0xf5, 0x32, 0xf1, 0x09, 0xea, 0xcb, 0xe3, 0xa9, 0x81,
	0x43, 0x3c, 0x34, 0x90, 0xa1, 0xcb, 0x88, 0x2e, 0x9c, 0x0f, 0x11, 0x09, 0x92, 0x7b, 0x64, 0xfc,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x42, 0x89, 0x48, 0xc2, 0x9d, 0x03, 0x00, 0x00,
}