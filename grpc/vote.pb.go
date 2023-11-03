// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: grpc/vote.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UsersId     int32                  `protobuf:"varint,2,opt,name=users_id,json=usersId,proto3" json:"users_id,omitempty"`
	RatedUserId int32                  `protobuf:"varint,3,opt,name=rated_user_id,json=ratedUserId,proto3" json:"rated_user_id,omitempty"`
	Vote        int32                  `protobuf:"varint,4,opt,name=vote,proto3" json:"vote,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{0}
}

func (x *Vote) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vote) GetUsersId() int32 {
	if x != nil {
		return x.UsersId
	}
	return 0
}

func (x *Vote) GetRatedUserId() int32 {
	if x != nil {
		return x.RatedUserId
	}
	return 0
}

func (x *Vote) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

func (x *Vote) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Vote) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type AddVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RatedUserId int32 `protobuf:"varint,2,opt,name=rated_user_id,json=ratedUserId,proto3" json:"rated_user_id,omitempty"`
	Vote        int32 `protobuf:"varint,3,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *AddVoteRequest) Reset() {
	*x = AddVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVoteRequest) ProtoMessage() {}

func (x *AddVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVoteRequest.ProtoReflect.Descriptor instead.
func (*AddVoteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{1}
}

func (x *AddVoteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddVoteRequest) GetRatedUserId() int32 {
	if x != nil {
		return x.RatedUserId
	}
	return 0
}

func (x *AddVoteRequest) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

type ChangeVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RatedUserId int32 `protobuf:"varint,2,opt,name=rated_user_id,json=ratedUserId,proto3" json:"rated_user_id,omitempty"`
	Vote        int32 `protobuf:"varint,3,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *ChangeVoteRequest) Reset() {
	*x = ChangeVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeVoteRequest) ProtoMessage() {}

func (x *ChangeVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeVoteRequest.ProtoReflect.Descriptor instead.
func (*ChangeVoteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeVoteRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChangeVoteRequest) GetRatedUserId() int32 {
	if x != nil {
		return x.RatedUserId
	}
	return 0
}

func (x *ChangeVoteRequest) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

type GetAllVotesResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote []*Vote `protobuf:"bytes,1,rep,name=vote,proto3" json:"vote,omitempty"`
}

func (x *GetAllVotesResponce) Reset() {
	*x = GetAllVotesResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllVotesResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllVotesResponce) ProtoMessage() {}

func (x *GetAllVotesResponce) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllVotesResponce.ProtoReflect.Descriptor instead.
func (*GetAllVotesResponce) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllVotesResponce) GetVote() []*Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

type GetUsersRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUsersRateRequest) Reset() {
	*x = GetUsersRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRateRequest) ProtoMessage() {}

func (x *GetUsersRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRateRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRateRequest) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersRateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUsersRateResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote int32 `protobuf:"varint,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *GetUsersRateResponce) Reset() {
	*x = GetUsersRateResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRateResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRateResponce) ProtoMessage() {}

func (x *GetUsersRateResponce) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRateResponce.ProtoReflect.Descriptor instead.
func (*GetUsersRateResponce) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersRateResponce) GetVote() int32 {
	if x != nil {
		return x.Vote
	}
	return 0
}

type DeleteVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RatedUserId int32 `protobuf:"varint,2,opt,name=rated_user_id,json=ratedUserId,proto3" json:"rated_user_id,omitempty"`
}

func (x *DeleteVoteRequest) Reset() {
	*x = DeleteVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_vote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVoteRequest) ProtoMessage() {}

func (x *DeleteVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_vote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteVoteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_vote_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteVoteRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteVoteRequest) GetRatedUserId() int32 {
	if x != nil {
		return x.RatedUserId
	}
	return 0
}

var File_grpc_vote_proto protoreflect.FileDescriptor

var file_grpc_vote_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdf, 0x01, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x58, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x64, 0x0a, 0x11, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x74,
	0x65, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x76,
	0x6f, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x50, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xf2, 0x01, 0x0a, 0x0b, 0x56, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0e, 0x5a,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_vote_proto_rawDescOnce sync.Once
	file_grpc_vote_proto_rawDescData = file_grpc_vote_proto_rawDesc
)

func file_grpc_vote_proto_rawDescGZIP() []byte {
	file_grpc_vote_proto_rawDescOnce.Do(func() {
		file_grpc_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_vote_proto_rawDescData)
	})
	return file_grpc_vote_proto_rawDescData
}

var file_grpc_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_vote_proto_goTypes = []interface{}{
	(*Vote)(nil),                  // 0: Vote
	(*AddVoteRequest)(nil),        // 1: AddVoteRequest
	(*ChangeVoteRequest)(nil),     // 2: ChangeVoteRequest
	(*GetAllVotesResponce)(nil),   // 3: GetAllVotesResponce
	(*GetUsersRateRequest)(nil),   // 4: GetUsersRateRequest
	(*GetUsersRateResponce)(nil),  // 5: GetUsersRateResponce
	(*DeleteVoteRequest)(nil),     // 6: DeleteVoteRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_grpc_vote_proto_depIdxs = []int32{
	7, // 0: Vote.updated_at:type_name -> google.protobuf.Timestamp
	7, // 1: Vote.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 2: GetAllVotesResponce.vote:type_name -> Vote
	1, // 3: VoteService.AddVote:input_type -> AddVoteRequest
	2, // 4: VoteService.ChangeVote:input_type -> ChangeVoteRequest
	8, // 5: VoteService.GetAllVotes:input_type -> google.protobuf.Empty
	6, // 6: VoteService.DeleteVote:input_type -> DeleteVoteRequest
	8, // 7: VoteService.AddVote:output_type -> google.protobuf.Empty
	8, // 8: VoteService.ChangeVote:output_type -> google.protobuf.Empty
	3, // 9: VoteService.GetAllVotes:output_type -> GetAllVotesResponce
	8, // 10: VoteService.DeleteVote:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_vote_proto_init() }
func file_grpc_vote_proto_init() {
	if File_grpc_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddVoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeVoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllVotesResponce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_vote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRateResponce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_vote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_vote_proto_goTypes,
		DependencyIndexes: file_grpc_vote_proto_depIdxs,
		MessageInfos:      file_grpc_vote_proto_msgTypes,
	}.Build()
	File_grpc_vote_proto = out.File
	file_grpc_vote_proto_rawDesc = nil
	file_grpc_vote_proto_goTypes = nil
	file_grpc_vote_proto_depIdxs = nil
}
