// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: sys.proto

package sys

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

// 用户登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 添加｜｜更新角色
type SaveOrUpdateRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark   string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	CreateBy string `protobuf:"bytes,4,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	UpdateBy string `protobuf:"bytes,5,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
}

func (x *SaveOrUpdateRoleReq) Reset() {
	*x = SaveOrUpdateRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveOrUpdateRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveOrUpdateRoleReq) ProtoMessage() {}

func (x *SaveOrUpdateRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveOrUpdateRoleReq.ProtoReflect.Descriptor instead.
func (*SaveOrUpdateRoleReq) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{2}
}

func (x *SaveOrUpdateRoleReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveOrUpdateRoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveOrUpdateRoleReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SaveOrUpdateRoleReq) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *SaveOrUpdateRoleReq) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

type SaveOrUpdateRoleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *SaveOrUpdateRoleResp) Reset() {
	*x = SaveOrUpdateRoleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveOrUpdateRoleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveOrUpdateRoleResp) ProtoMessage() {}

func (x *SaveOrUpdateRoleResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveOrUpdateRoleResp.ProtoReflect.Descriptor instead.
func (*SaveOrUpdateRoleResp) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{3}
}

func (x *SaveOrUpdateRoleResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// 删除角色
type RoleDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *RoleDeleteReq) Reset() {
	*x = RoleDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteReq) ProtoMessage() {}

func (x *RoleDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteReq.ProtoReflect.Descriptor instead.
func (*RoleDeleteReq) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{4}
}

func (x *RoleDeleteReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type RoleDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *RoleDeleteResp) Reset() {
	*x = RoleDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteResp) ProtoMessage() {}

func (x *RoleDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteResp.ProtoReflect.Descriptor instead.
func (*RoleDeleteResp) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{5}
}

func (x *RoleDeleteResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// 角色列表
type RoleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64  `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RoleListReq) Reset() {
	*x = RoleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListReq) ProtoMessage() {}

func (x *RoleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListReq.ProtoReflect.Descriptor instead.
func (*RoleListReq) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{6}
}

func (x *RoleListReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *RoleListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RoleListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RoleListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark   string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	CreateBy string `protobuf:"bytes,4,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	CreateAt string `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateBy string `protobuf:"bytes,6,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	UpdateAt string `protobuf:"bytes,7,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *RoleListData) Reset() {
	*x = RoleListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListData) ProtoMessage() {}

func (x *RoleListData) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListData.ProtoReflect.Descriptor instead.
func (*RoleListData) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{7}
}

func (x *RoleListData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoleListData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleListData) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *RoleListData) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *RoleListData) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *RoleListData) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *RoleListData) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

type RoleListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*RoleListData `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *RoleListResp) Reset() {
	*x = RoleListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListResp) ProtoMessage() {}

func (x *RoleListResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListResp.ProtoReflect.Descriptor instead.
func (*RoleListResp) Descriptor() ([]byte, []int) {
	return file_sys_proto_rawDescGZIP(), []int{8}
}

func (x *RoleListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RoleListResp) GetList() []*RoleListData {
	if x != nil {
		return x.List
	}
	return nil
}

var File_sys_proto protoreflect.FileDescriptor

var file_sys_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x79, 0x73,
	0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x8b, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x22, 0x2a,
	0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x21, 0x0a, 0x0d, 0x52, 0x6f,
	0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x24, 0x0a,
	0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x6e, 0x67, 0x22, 0x59, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbe,
	0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22,
	0x4b, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xe2, 0x01, 0x0a,
	0x03, 0x53, 0x79, 0x73, 0x12, 0x2a, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x0e, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x47, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x0a, 0x52, 0x6f, 0x6c,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x73, 0x79,
	0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x2f, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x73,
	0x79, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x73, 0x79, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sys_proto_rawDescOnce sync.Once
	file_sys_proto_rawDescData = file_sys_proto_rawDesc
)

func file_sys_proto_rawDescGZIP() []byte {
	file_sys_proto_rawDescOnce.Do(func() {
		file_sys_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_proto_rawDescData)
	})
	return file_sys_proto_rawDescData
}

var file_sys_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sys_proto_goTypes = []interface{}{
	(*LoginReq)(nil),             // 0: sys.LoginReq
	(*LoginResp)(nil),            // 1: sys.LoginResp
	(*SaveOrUpdateRoleReq)(nil),  // 2: sys.SaveOrUpdateRoleReq
	(*SaveOrUpdateRoleResp)(nil), // 3: sys.SaveOrUpdateRoleResp
	(*RoleDeleteReq)(nil),        // 4: sys.RoleDeleteReq
	(*RoleDeleteResp)(nil),       // 5: sys.RoleDeleteResp
	(*RoleListReq)(nil),          // 6: sys.RoleListReq
	(*RoleListData)(nil),         // 7: sys.RoleListData
	(*RoleListResp)(nil),         // 8: sys.RoleListResp
}
var file_sys_proto_depIdxs = []int32{
	7, // 0: sys.RoleListResp.list:type_name -> sys.RoleListData
	0, // 1: sys.Sys.UserLogin:input_type -> sys.LoginReq
	2, // 2: sys.Sys.SaveOrUpdateRole:input_type -> sys.SaveOrUpdateRoleReq
	4, // 3: sys.Sys.RoleDelete:input_type -> sys.RoleDeleteReq
	6, // 4: sys.Sys.RoleList:input_type -> sys.RoleListReq
	1, // 5: sys.Sys.UserLogin:output_type -> sys.LoginResp
	3, // 6: sys.Sys.SaveOrUpdateRole:output_type -> sys.SaveOrUpdateRoleResp
	5, // 7: sys.Sys.RoleDelete:output_type -> sys.RoleDeleteResp
	8, // 8: sys.Sys.RoleList:output_type -> sys.RoleListResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sys_proto_init() }
func file_sys_proto_init() {
	if File_sys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_sys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_sys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveOrUpdateRoleReq); i {
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
		file_sys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveOrUpdateRoleResp); i {
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
		file_sys_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteReq); i {
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
		file_sys_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteResp); i {
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
		file_sys_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListReq); i {
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
		file_sys_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListData); i {
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
		file_sys_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListResp); i {
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
			RawDescriptor: file_sys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_proto_goTypes,
		DependencyIndexes: file_sys_proto_depIdxs,
		MessageInfos:      file_sys_proto_msgTypes,
	}.Build()
	File_sys_proto = out.File
	file_sys_proto_rawDesc = nil
	file_sys_proto_goTypes = nil
	file_sys_proto_depIdxs = nil
}