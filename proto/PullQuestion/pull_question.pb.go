// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pull_question.proto

package pull_question

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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

type AllMyQuestionRequest struct {
	Questioner           string   `protobuf:"bytes,1,opt,name=questioner,proto3" json:"questioner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllMyQuestionRequest) Reset()         { *m = AllMyQuestionRequest{} }
func (m *AllMyQuestionRequest) String() string { return proto.CompactTextString(m) }
func (*AllMyQuestionRequest) ProtoMessage()    {}
func (*AllMyQuestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{0}
}
func (m *AllMyQuestionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMyQuestionRequest.Unmarshal(m, b)
}
func (m *AllMyQuestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMyQuestionRequest.Marshal(b, m, deterministic)
}
func (dst *AllMyQuestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMyQuestionRequest.Merge(dst, src)
}
func (m *AllMyQuestionRequest) XXX_Size() int {
	return xxx_messageInfo_AllMyQuestionRequest.Size(m)
}
func (m *AllMyQuestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMyQuestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllMyQuestionRequest proto.InternalMessageInfo

func (m *AllMyQuestionRequest) GetQuestioner() string {
	if m != nil {
		return m.Questioner
	}
	return ""
}

type AllMyQuestionReply struct {
	Question             []*QuestionList `protobuf:"bytes,1,rep,name=question,proto3" json:"question,omitempty"`
	Result               string          `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Message              bool            `protobuf:"varint,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AllMyQuestionReply) Reset()         { *m = AllMyQuestionReply{} }
func (m *AllMyQuestionReply) String() string { return proto.CompactTextString(m) }
func (*AllMyQuestionReply) ProtoMessage()    {}
func (*AllMyQuestionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{1}
}
func (m *AllMyQuestionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMyQuestionReply.Unmarshal(m, b)
}
func (m *AllMyQuestionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMyQuestionReply.Marshal(b, m, deterministic)
}
func (dst *AllMyQuestionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMyQuestionReply.Merge(dst, src)
}
func (m *AllMyQuestionReply) XXX_Size() int {
	return xxx_messageInfo_AllMyQuestionReply.Size(m)
}
func (m *AllMyQuestionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMyQuestionReply.DiscardUnknown(m)
}

var xxx_messageInfo_AllMyQuestionReply proto.InternalMessageInfo

func (m *AllMyQuestionReply) GetQuestion() []*QuestionList {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *AllMyQuestionReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *AllMyQuestionReply) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

type QuestionList struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Question             string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Questioner           string   `protobuf:"bytes,3,opt,name=questioner,proto3" json:"questioner,omitempty"`
	AnswerCount          string   `protobuf:"bytes,4,opt,name=answerCount,proto3" json:"answerCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionList) Reset()         { *m = QuestionList{} }
func (m *QuestionList) String() string { return proto.CompactTextString(m) }
func (*QuestionList) ProtoMessage()    {}
func (*QuestionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{2}
}
func (m *QuestionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionList.Unmarshal(m, b)
}
func (m *QuestionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionList.Marshal(b, m, deterministic)
}
func (dst *QuestionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionList.Merge(dst, src)
}
func (m *QuestionList) XXX_Size() int {
	return xxx_messageInfo_QuestionList.Size(m)
}
func (m *QuestionList) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionList.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionList proto.InternalMessageInfo

func (m *QuestionList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QuestionList) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *QuestionList) GetQuestioner() string {
	if m != nil {
		return m.Questioner
	}
	return ""
}

func (m *QuestionList) GetAnswerCount() string {
	if m != nil {
		return m.AnswerCount
	}
	return ""
}

type AllMyAnswerRequest struct {
	Answerer             string   `protobuf:"bytes,1,opt,name=answerer,proto3" json:"answerer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllMyAnswerRequest) Reset()         { *m = AllMyAnswerRequest{} }
func (m *AllMyAnswerRequest) String() string { return proto.CompactTextString(m) }
func (*AllMyAnswerRequest) ProtoMessage()    {}
func (*AllMyAnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{3}
}
func (m *AllMyAnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMyAnswerRequest.Unmarshal(m, b)
}
func (m *AllMyAnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMyAnswerRequest.Marshal(b, m, deterministic)
}
func (dst *AllMyAnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMyAnswerRequest.Merge(dst, src)
}
func (m *AllMyAnswerRequest) XXX_Size() int {
	return xxx_messageInfo_AllMyAnswerRequest.Size(m)
}
func (m *AllMyAnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMyAnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllMyAnswerRequest proto.InternalMessageInfo

func (m *AllMyAnswerRequest) GetAnswerer() string {
	if m != nil {
		return m.Answerer
	}
	return ""
}

type AllMyAnswerReply struct {
	Answer               []*AnswerList `protobuf:"bytes,1,rep,name=answer,proto3" json:"answer,omitempty"`
	Result               string        `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Message              bool          `protobuf:"varint,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AllMyAnswerReply) Reset()         { *m = AllMyAnswerReply{} }
func (m *AllMyAnswerReply) String() string { return proto.CompactTextString(m) }
func (*AllMyAnswerReply) ProtoMessage()    {}
func (*AllMyAnswerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{4}
}
func (m *AllMyAnswerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllMyAnswerReply.Unmarshal(m, b)
}
func (m *AllMyAnswerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllMyAnswerReply.Marshal(b, m, deterministic)
}
func (dst *AllMyAnswerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllMyAnswerReply.Merge(dst, src)
}
func (m *AllMyAnswerReply) XXX_Size() int {
	return xxx_messageInfo_AllMyAnswerReply.Size(m)
}
func (m *AllMyAnswerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AllMyAnswerReply.DiscardUnknown(m)
}

var xxx_messageInfo_AllMyAnswerReply proto.InternalMessageInfo

func (m *AllMyAnswerReply) GetAnswer() []*AnswerList {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (m *AllMyAnswerReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *AllMyAnswerReply) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

type AnswerList struct {
	Num                  string   `protobuf:"bytes,1,opt,name=num,proto3" json:"num,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Answerer             string   `protobuf:"bytes,3,opt,name=answerer,proto3" json:"answerer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerList) Reset()         { *m = AnswerList{} }
func (m *AnswerList) String() string { return proto.CompactTextString(m) }
func (*AnswerList) ProtoMessage()    {}
func (*AnswerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{5}
}
func (m *AnswerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerList.Unmarshal(m, b)
}
func (m *AnswerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerList.Marshal(b, m, deterministic)
}
func (dst *AnswerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerList.Merge(dst, src)
}
func (m *AnswerList) XXX_Size() int {
	return xxx_messageInfo_AnswerList.Size(m)
}
func (m *AnswerList) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerList.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerList proto.InternalMessageInfo

func (m *AnswerList) GetNum() string {
	if m != nil {
		return m.Num
	}
	return ""
}

func (m *AnswerList) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *AnswerList) GetAnswerer() string {
	if m != nil {
		return m.Answerer
	}
	return ""
}

type HighestRankingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HighestRankingRequest) Reset()         { *m = HighestRankingRequest{} }
func (m *HighestRankingRequest) String() string { return proto.CompactTextString(m) }
func (*HighestRankingRequest) ProtoMessage()    {}
func (*HighestRankingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{6}
}
func (m *HighestRankingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HighestRankingRequest.Unmarshal(m, b)
}
func (m *HighestRankingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HighestRankingRequest.Marshal(b, m, deterministic)
}
func (dst *HighestRankingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HighestRankingRequest.Merge(dst, src)
}
func (m *HighestRankingRequest) XXX_Size() int {
	return xxx_messageInfo_HighestRankingRequest.Size(m)
}
func (m *HighestRankingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HighestRankingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HighestRankingRequest proto.InternalMessageInfo

type HighestRankingReply struct {
	Question             []*QuestionList `protobuf:"bytes,1,rep,name=question,proto3" json:"question,omitempty"`
	Result               string          `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Message              bool            `protobuf:"varint,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HighestRankingReply) Reset()         { *m = HighestRankingReply{} }
func (m *HighestRankingReply) String() string { return proto.CompactTextString(m) }
func (*HighestRankingReply) ProtoMessage()    {}
func (*HighestRankingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{7}
}
func (m *HighestRankingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HighestRankingReply.Unmarshal(m, b)
}
func (m *HighestRankingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HighestRankingReply.Marshal(b, m, deterministic)
}
func (dst *HighestRankingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HighestRankingReply.Merge(dst, src)
}
func (m *HighestRankingReply) XXX_Size() int {
	return xxx_messageInfo_HighestRankingReply.Size(m)
}
func (m *HighestRankingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HighestRankingReply.DiscardUnknown(m)
}

var xxx_messageInfo_HighestRankingReply proto.InternalMessageInfo

func (m *HighestRankingReply) GetQuestion() []*QuestionList {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *HighestRankingReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *HighestRankingReply) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

type RedisSortRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisSortRequest) Reset()         { *m = RedisSortRequest{} }
func (m *RedisSortRequest) String() string { return proto.CompactTextString(m) }
func (*RedisSortRequest) ProtoMessage()    {}
func (*RedisSortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{8}
}
func (m *RedisSortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisSortRequest.Unmarshal(m, b)
}
func (m *RedisSortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisSortRequest.Marshal(b, m, deterministic)
}
func (dst *RedisSortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisSortRequest.Merge(dst, src)
}
func (m *RedisSortRequest) XXX_Size() int {
	return xxx_messageInfo_RedisSortRequest.Size(m)
}
func (m *RedisSortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisSortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RedisSortRequest proto.InternalMessageInfo

type RedisSortReply struct {
	Question             []*QuestionList `protobuf:"bytes,1,rep,name=question,proto3" json:"question,omitempty"`
	Result               string          `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Message              bool            `protobuf:"varint,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RedisSortReply) Reset()         { *m = RedisSortReply{} }
func (m *RedisSortReply) String() string { return proto.CompactTextString(m) }
func (*RedisSortReply) ProtoMessage()    {}
func (*RedisSortReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pull_question_1944700b07cc89b4, []int{9}
}
func (m *RedisSortReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisSortReply.Unmarshal(m, b)
}
func (m *RedisSortReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisSortReply.Marshal(b, m, deterministic)
}
func (dst *RedisSortReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisSortReply.Merge(dst, src)
}
func (m *RedisSortReply) XXX_Size() int {
	return xxx_messageInfo_RedisSortReply.Size(m)
}
func (m *RedisSortReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisSortReply.DiscardUnknown(m)
}

var xxx_messageInfo_RedisSortReply proto.InternalMessageInfo

func (m *RedisSortReply) GetQuestion() []*QuestionList {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *RedisSortReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *RedisSortReply) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

func init() {
	proto.RegisterType((*AllMyQuestionRequest)(nil), "pull_question.AllMyQuestionRequest")
	proto.RegisterType((*AllMyQuestionReply)(nil), "pull_question.AllMyQuestionReply")
	proto.RegisterType((*QuestionList)(nil), "pull_question.QuestionList")
	proto.RegisterType((*AllMyAnswerRequest)(nil), "pull_question.AllMyAnswerRequest")
	proto.RegisterType((*AllMyAnswerReply)(nil), "pull_question.AllMyAnswerReply")
	proto.RegisterType((*AnswerList)(nil), "pull_question.AnswerList")
	proto.RegisterType((*HighestRankingRequest)(nil), "pull_question.HighestRankingRequest")
	proto.RegisterType((*HighestRankingReply)(nil), "pull_question.HighestRankingReply")
	proto.RegisterType((*RedisSortRequest)(nil), "pull_question.RedisSortRequest")
	proto.RegisterType((*RedisSortReply)(nil), "pull_question.RedisSortReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	AllMyQuestion(ctx context.Context, in *AllMyQuestionRequest, opts ...grpc.CallOption) (*AllMyQuestionReply, error)
	AllMyAnswer(ctx context.Context, in *AllMyAnswerRequest, opts ...grpc.CallOption) (*AllMyAnswerReply, error)
	HighestRanking(ctx context.Context, in *HighestRankingRequest, opts ...grpc.CallOption) (*HighestRankingReply, error)
	RedisSort(ctx context.Context, in *RedisSortRequest, opts ...grpc.CallOption) (*RedisSortReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) AllMyQuestion(ctx context.Context, in *AllMyQuestionRequest, opts ...grpc.CallOption) (*AllMyQuestionReply, error) {
	out := new(AllMyQuestionReply)
	err := c.cc.Invoke(ctx, "/pull_question.Greeter/AllMyQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) AllMyAnswer(ctx context.Context, in *AllMyAnswerRequest, opts ...grpc.CallOption) (*AllMyAnswerReply, error) {
	out := new(AllMyAnswerReply)
	err := c.cc.Invoke(ctx, "/pull_question.Greeter/AllMyAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) HighestRanking(ctx context.Context, in *HighestRankingRequest, opts ...grpc.CallOption) (*HighestRankingReply, error) {
	out := new(HighestRankingReply)
	err := c.cc.Invoke(ctx, "/pull_question.Greeter/HighestRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) RedisSort(ctx context.Context, in *RedisSortRequest, opts ...grpc.CallOption) (*RedisSortReply, error) {
	out := new(RedisSortReply)
	err := c.cc.Invoke(ctx, "/pull_question.Greeter/RedisSort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	AllMyQuestion(context.Context, *AllMyQuestionRequest) (*AllMyQuestionReply, error)
	AllMyAnswer(context.Context, *AllMyAnswerRequest) (*AllMyAnswerReply, error)
	HighestRanking(context.Context, *HighestRankingRequest) (*HighestRankingReply, error)
	RedisSort(context.Context, *RedisSortRequest) (*RedisSortReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_AllMyQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllMyQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AllMyQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pull_question.Greeter/AllMyQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AllMyQuestion(ctx, req.(*AllMyQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_AllMyAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllMyAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AllMyAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pull_question.Greeter/AllMyAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AllMyAnswer(ctx, req.(*AllMyAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_HighestRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HighestRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).HighestRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pull_question.Greeter/HighestRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).HighestRanking(ctx, req.(*HighestRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_RedisSort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedisSortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).RedisSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pull_question.Greeter/RedisSort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).RedisSort(ctx, req.(*RedisSortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pull_question.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllMyQuestion",
			Handler:    _Greeter_AllMyQuestion_Handler,
		},
		{
			MethodName: "AllMyAnswer",
			Handler:    _Greeter_AllMyAnswer_Handler,
		},
		{
			MethodName: "HighestRanking",
			Handler:    _Greeter_HighestRanking_Handler,
		},
		{
			MethodName: "RedisSort",
			Handler:    _Greeter_RedisSort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pull_question.proto",
}

func init() { proto.RegisterFile("pull_question.proto", fileDescriptor_pull_question_1944700b07cc89b4) }

var fileDescriptor_pull_question_1944700b07cc89b4 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0x09, 0xea, 0xc7, 0x94, 0x56, 0xd1, 0x94, 0x8f, 0x10, 0x04, 0x04, 0xc3, 0xa1, 0xa7,
	0x08, 0x8a, 0x04, 0xe7, 0xc2, 0x01, 0x0e, 0xf4, 0x40, 0x7a, 0x42, 0x42, 0x42, 0x81, 0x5a, 0xc1,
	0xc2, 0x4d, 0x82, 0xed, 0xa8, 0x54, 0x20, 0xc1, 0xaf, 0xe5, 0x77, 0xac, 0x9a, 0x38, 0x69, 0x62,
	0x75, 0xbb, 0xd2, 0x1e, 0x7a, 0x8b, 0xc7, 0x6f, 0xde, 0xbc, 0x79, 0x33, 0x0e, 0x4c, 0xf3, 0x82,
	0xf3, 0x2f, 0x3f, 0x0b, 0x2a, 0x15, 0xcb, 0xd2, 0x30, 0x17, 0x99, 0xca, 0x70, 0xdc, 0x09, 0x92,
	0x57, 0x70, 0x7b, 0xc1, 0xf9, 0x72, 0xf7, 0x51, 0x07, 0x22, 0x5a, 0x5e, 0xe1, 0x23, 0x80, 0x1a,
	0x43, 0x85, 0x67, 0x05, 0xd6, 0x6c, 0x18, 0xb5, 0x22, 0xe4, 0x2f, 0xa0, 0x91, 0x97, 0xf3, 0x1d,
	0xbe, 0x86, 0x41, 0x8d, 0xf1, 0xac, 0xc0, 0x99, 0x8d, 0xe6, 0x0f, 0xc2, 0xae, 0x88, 0x1a, 0xff,
	0x81, 0x49, 0x15, 0x35, 0x60, 0xbc, 0x0b, 0x3d, 0x41, 0x65, 0xc1, 0x95, 0x67, 0x97, 0xa5, 0xf4,
	0x09, 0x3d, 0xe8, 0x6f, 0xa8, 0x94, 0x71, 0x42, 0x3d, 0x27, 0xb0, 0x66, 0x83, 0xa8, 0x3e, 0x92,
	0x3f, 0x70, 0xab, 0xcd, 0x85, 0x13, 0xb0, 0xd9, 0x5a, 0x0b, 0xb5, 0xd9, 0x1a, 0xfd, 0x96, 0x94,
	0x8a, 0xf3, 0x50, 0xad, 0xdb, 0x9c, 0x63, 0x36, 0x87, 0x01, 0x8c, 0xe2, 0x54, 0x6e, 0xa9, 0x78,
	0x9b, 0x15, 0xa9, 0xf2, 0x6e, 0x96, 0x80, 0x76, 0x88, 0x3c, 0xd7, 0xed, 0x2f, 0xca, 0x58, 0x6d,
	0x9a, 0x0f, 0x83, 0x0a, 0xd4, 0x58, 0xd6, 0x9c, 0xc9, 0x16, 0xdc, 0x4e, 0xc6, 0xde, 0xae, 0x17,
	0xd0, 0xab, 0xee, 0xb5, 0x59, 0xf7, 0x0d, 0xb3, 0x2a, 0x6c, 0x69, 0x95, 0x06, 0x5e, 0xc3, 0xa8,
	0x08, 0xe0, 0xc0, 0x83, 0x2e, 0x38, 0x69, 0xb1, 0xd1, 0xea, 0xf6, 0x9f, 0x7b, 0x46, 0x2d, 0x42,
	0x33, 0xea, 0x4a, 0xed, 0x66, 0x1c, 0xa3, 0x99, 0x7b, 0x70, 0xe7, 0x3d, 0x4b, 0xbe, 0x53, 0xa9,
	0xa2, 0x38, 0xfd, 0xc1, 0xd2, 0x44, 0x3b, 0x40, 0xfe, 0x59, 0x30, 0x35, 0x6f, 0xce, 0xbc, 0x18,
	0x08, 0x6e, 0x44, 0xd7, 0x4c, 0xae, 0x32, 0xa1, 0x6a, 0x59, 0xbf, 0x61, 0xd2, 0x8a, 0x9d, 0x57,
	0xd0, 0xfc, 0xbf, 0x0d, 0xfd, 0x77, 0x82, 0x52, 0x45, 0x05, 0x7e, 0x82, 0x71, 0xe7, 0xd9, 0xe0,
	0x53, 0x73, 0xe4, 0x47, 0x1e, 0xa3, 0xff, 0xe4, 0x34, 0x28, 0xe7, 0x3b, 0x72, 0x03, 0x57, 0x30,
	0x6a, 0x2d, 0x18, 0x1e, 0xcd, 0xe9, 0xac, 0xab, 0xff, 0xf8, 0x14, 0xa4, 0x22, 0xfd, 0x0c, 0x93,
	0xee, 0x38, 0xf1, 0x99, 0x91, 0x74, 0x74, 0x0f, 0x7c, 0x72, 0x05, 0xaa, 0x62, 0x5f, 0xc2, 0xb0,
	0x19, 0x0b, 0x9a, 0x6a, 0xcc, 0x21, 0xfa, 0x0f, 0x2f, 0x07, 0x94, 0x74, 0x6f, 0x10, 0x5c, 0x96,
	0x85, 0x89, 0xc8, 0xbf, 0x85, 0xf4, 0x57, 0xbc, 0xc9, 0x39, 0x95, 0x5f, 0x7b, 0xe5, 0x5f, 0xef,
	0xe5, 0x45, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xaa, 0x1f, 0x9e, 0x0c, 0x05, 0x00, 0x00,
}