// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ums.proto

package ums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 添加｜｜更新会员
type SaveOrUpdateMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Nickname string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Status   string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Avatar   string `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender   string `protobuf:"bytes,10,opt,name=gender,proto3" json:"gender,omitempty"`
	Email    string `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SaveOrUpdateMemberReq) Reset() {
	*x = SaveOrUpdateMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveOrUpdateMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveOrUpdateMemberReq) ProtoMessage() {}

func (x *SaveOrUpdateMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveOrUpdateMemberReq.ProtoReflect.Descriptor instead.
func (*SaveOrUpdateMemberReq) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{0}
}

func (x *SaveOrUpdateMemberReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveOrUpdateMemberReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SaveOrUpdateMemberReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SaveOrUpdateMemberReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SaveOrUpdateMemberReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SaveOrUpdateMemberReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SaveOrUpdateMemberReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *SaveOrUpdateMemberReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SaveOrUpdateMemberResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *SaveOrUpdateMemberResp) Reset() {
	*x = SaveOrUpdateMemberResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveOrUpdateMemberResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveOrUpdateMemberResp) ProtoMessage() {}

func (x *SaveOrUpdateMemberResp) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveOrUpdateMemberResp.ProtoReflect.Descriptor instead.
func (*SaveOrUpdateMemberResp) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{1}
}

func (x *SaveOrUpdateMemberResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// 会员删除
type MemberDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MemberDeleteReq) Reset() {
	*x = MemberDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberDeleteReq) ProtoMessage() {}

func (x *MemberDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberDeleteReq.ProtoReflect.Descriptor instead.
func (*MemberDeleteReq) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{2}
}

func (x *MemberDeleteReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MemberDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *MemberDeleteResp) Reset() {
	*x = MemberDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberDeleteResp) ProtoMessage() {}

func (x *MemberDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberDeleteResp.ProtoReflect.Descriptor instead.
func (*MemberDeleteResp) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{3}
}

func (x *MemberDeleteResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// 会员列表
type MemberListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64  `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Status   string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Nickname string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *MemberListReq) Reset() {
	*x = MemberListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListReq) ProtoMessage() {}

func (x *MemberListReq) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListReq.ProtoReflect.Descriptor instead.
func (*MemberListReq) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{4}
}

func (x *MemberListReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *MemberListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MemberListReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *MemberListReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MemberListReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type MemberListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname  string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Status    string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Avatar    string `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender    string `protobuf:"bytes,10,opt,name=gender,proto3" json:"gender,omitempty"`
	Email     string `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt string `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *MemberListData) Reset() {
	*x = MemberListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListData) ProtoMessage() {}

func (x *MemberListData) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListData.ProtoReflect.Descriptor instead.
func (*MemberListData) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{5}
}

func (x *MemberListData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberListData) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *MemberListData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *MemberListData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MemberListData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *MemberListData) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *MemberListData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberListData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type MemberListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*MemberListData `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MemberListResp) Reset() {
	*x = MemberListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ums_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListResp) ProtoMessage() {}

func (x *MemberListResp) ProtoReflect() protoreflect.Message {
	mi := &file_ums_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListResp.ProtoReflect.Descriptor instead.
func (*MemberListResp) Descriptor() ([]byte, []int) {
	return file_ums_proto_rawDescGZIP(), []int{6}
}

func (x *MemberListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemberListResp) GetList() []*MemberListData {
	if x != nil {
		return x.List
	}
	return nil
}

var File_ums_proto protoreflect.FileDescriptor

var file_ums_proto_rawDesc = []byte{
	0x0a, 0x09, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x75, 0x6d, 0x73,
	0x22, 0xd3, 0x01, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2c, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x6e, 0x67, 0x22, 0x23, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4f, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xc8, 0x01, 0x0a, 0x03, 0x55, 0x6d, 0x73,
	0x12, 0x4d, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3b, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x0a,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x75, 0x6d, 0x73,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ums_proto_rawDescOnce sync.Once
	file_ums_proto_rawDescData = file_ums_proto_rawDesc
)

func file_ums_proto_rawDescGZIP() []byte {
	file_ums_proto_rawDescOnce.Do(func() {
		file_ums_proto_rawDescData = protoimpl.X.CompressGZIP(file_ums_proto_rawDescData)
	})
	return file_ums_proto_rawDescData
}

var file_ums_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ums_proto_goTypes = []interface{}{
	(*SaveOrUpdateMemberReq)(nil),  // 0: ums.SaveOrUpdateMemberReq
	(*SaveOrUpdateMemberResp)(nil), // 1: ums.SaveOrUpdateMemberResp
	(*MemberDeleteReq)(nil),        // 2: ums.MemberDeleteReq
	(*MemberDeleteResp)(nil),       // 3: ums.MemberDeleteResp
	(*MemberListReq)(nil),          // 4: ums.MemberListReq
	(*MemberListData)(nil),         // 5: ums.MemberListData
	(*MemberListResp)(nil),         // 6: ums.MemberListResp
}
var file_ums_proto_depIdxs = []int32{
	5, // 0: ums.MemberListResp.list:type_name -> ums.MemberListData
	0, // 1: ums.Ums.SaveOrUpdateMember:input_type -> ums.SaveOrUpdateMemberReq
	2, // 2: ums.Ums.MemberDelete:input_type -> ums.MemberDeleteReq
	4, // 3: ums.Ums.MemberList:input_type -> ums.MemberListReq
	1, // 4: ums.Ums.SaveOrUpdateMember:output_type -> ums.SaveOrUpdateMemberResp
	3, // 5: ums.Ums.MemberDelete:output_type -> ums.MemberDeleteResp
	6, // 6: ums.Ums.MemberList:output_type -> ums.MemberListResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ums_proto_init() }
func file_ums_proto_init() {
	if File_ums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveOrUpdateMemberReq); i {
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
		file_ums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveOrUpdateMemberResp); i {
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
		file_ums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberDeleteReq); i {
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
		file_ums_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberDeleteResp); i {
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
		file_ums_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberListReq); i {
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
		file_ums_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberListData); i {
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
		file_ums_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberListResp); i {
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
			RawDescriptor: file_ums_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ums_proto_goTypes,
		DependencyIndexes: file_ums_proto_depIdxs,
		MessageInfos:      file_ums_proto_msgTypes,
	}.Build()
	File_ums_proto = out.File
	file_ums_proto_rawDesc = nil
	file_ums_proto_goTypes = nil
	file_ums_proto_depIdxs = nil
}