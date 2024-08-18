// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: cmd/cmd_codex.proto

package pb

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

type CodexTypeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodexIdList              []uint32          `protobuf:"varint,14,rep,packed,name=codex_id_list,json=codexIdList,proto3" json:"codex_id_list,omitempty"`
	WeaponMaxPromoteLevelMap map[uint32]uint32 `protobuf:"bytes,4,rep,name=weapon_max_promote_level_map,json=weaponMaxPromoteLevelMap,proto3" json:"weapon_max_promote_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Type                     CodexType         `protobuf:"varint,13,opt,name=type,proto3,enum=proto.CodexType" json:"type,omitempty"`
	HaveViewedList           []bool            `protobuf:"varint,5,rep,packed,name=have_viewed_list,json=haveViewedList,proto3" json:"have_viewed_list,omitempty"`
}

func (x *CodexTypeData) Reset() {
	*x = CodexTypeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodexTypeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodexTypeData) ProtoMessage() {}

func (x *CodexTypeData) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodexTypeData.ProtoReflect.Descriptor instead.
func (*CodexTypeData) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{0}
}

func (x *CodexTypeData) GetCodexIdList() []uint32 {
	if x != nil {
		return x.CodexIdList
	}
	return nil
}

func (x *CodexTypeData) GetWeaponMaxPromoteLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.WeaponMaxPromoteLevelMap
	}
	return nil
}

func (x *CodexTypeData) GetType() CodexType {
	if x != nil {
		return x.Type
	}
	return CodexType_CODEX_NONE
}

func (x *CodexTypeData) GetHaveViewedList() []bool {
	if x != nil {
		return x.HaveViewedList
	}
	return nil
}

type CodexDataFullNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastReadPushtipsCodexId  uint32           `protobuf:"varint,4,opt,name=last_read_pushtips_codex_id,json=lastReadPushtipsCodexId,proto3" json:"last_read_pushtips_codex_id,omitempty"`
	RecentViewedPushtipsList []uint32         `protobuf:"varint,2,rep,packed,name=recent_viewed_pushtips_list,json=recentViewedPushtipsList,proto3" json:"recent_viewed_pushtips_list,omitempty"`
	LastReadPushtipsTypeId   uint32           `protobuf:"varint,3,opt,name=last_read_pushtips_type_id,json=lastReadPushtipsTypeId,proto3" json:"last_read_pushtips_type_id,omitempty"`
	TypeDataList             []*CodexTypeData `protobuf:"bytes,6,rep,name=type_data_list,json=typeDataList,proto3" json:"type_data_list,omitempty"`
}

func (x *CodexDataFullNotify) Reset() {
	*x = CodexDataFullNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodexDataFullNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodexDataFullNotify) ProtoMessage() {}

func (x *CodexDataFullNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodexDataFullNotify.ProtoReflect.Descriptor instead.
func (*CodexDataFullNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{1}
}

func (x *CodexDataFullNotify) GetLastReadPushtipsCodexId() uint32 {
	if x != nil {
		return x.LastReadPushtipsCodexId
	}
	return 0
}

func (x *CodexDataFullNotify) GetRecentViewedPushtipsList() []uint32 {
	if x != nil {
		return x.RecentViewedPushtipsList
	}
	return nil
}

func (x *CodexDataFullNotify) GetLastReadPushtipsTypeId() uint32 {
	if x != nil {
		return x.LastReadPushtipsTypeId
	}
	return 0
}

func (x *CodexDataFullNotify) GetTypeDataList() []*CodexTypeData {
	if x != nil {
		return x.TypeDataList
	}
	return nil
}

type CodexDataUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    uint32    `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	WeaponMaxPromoteLevel uint32    `protobuf:"varint,15,opt,name=weapon_max_promote_level,json=weaponMaxPromoteLevel,proto3" json:"weapon_max_promote_level,omitempty"`
	Type                  CodexType `protobuf:"varint,11,opt,name=type,proto3,enum=proto.CodexType" json:"type,omitempty"`
}

func (x *CodexDataUpdateNotify) Reset() {
	*x = CodexDataUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodexDataUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodexDataUpdateNotify) ProtoMessage() {}

func (x *CodexDataUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodexDataUpdateNotify.ProtoReflect.Descriptor instead.
func (*CodexDataUpdateNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{2}
}

func (x *CodexDataUpdateNotify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CodexDataUpdateNotify) GetWeaponMaxPromoteLevel() uint32 {
	if x != nil {
		return x.WeaponMaxPromoteLevel
	}
	return 0
}

func (x *CodexDataUpdateNotify) GetType() CodexType {
	if x != nil {
		return x.Type
	}
	return CodexType_CODEX_NONE
}

type QueryCodexMonsterBeKilledNumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodexIdList []uint32 `protobuf:"varint,14,rep,packed,name=codex_id_list,json=codexIdList,proto3" json:"codex_id_list,omitempty"`
}

func (x *QueryCodexMonsterBeKilledNumReq) Reset() {
	*x = QueryCodexMonsterBeKilledNumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCodexMonsterBeKilledNumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCodexMonsterBeKilledNumReq) ProtoMessage() {}

func (x *QueryCodexMonsterBeKilledNumReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCodexMonsterBeKilledNumReq.ProtoReflect.Descriptor instead.
func (*QueryCodexMonsterBeKilledNumReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{3}
}

func (x *QueryCodexMonsterBeKilledNumReq) GetCodexIdList() []uint32 {
	if x != nil {
		return x.CodexIdList
	}
	return nil
}

type QueryCodexMonsterBeKilledNumRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodexIdList       []uint32 `protobuf:"varint,4,rep,packed,name=codex_id_list,json=codexIdList,proto3" json:"codex_id_list,omitempty"`
	BeCapturedNumList []uint32 `protobuf:"varint,6,rep,packed,name=be_captured_num_list,json=beCapturedNumList,proto3" json:"be_captured_num_list,omitempty"`
	BeKilledNumList   []uint32 `protobuf:"varint,12,rep,packed,name=be_killed_num_list,json=beKilledNumList,proto3" json:"be_killed_num_list,omitempty"`
	Retcode           int32    `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *QueryCodexMonsterBeKilledNumRsp) Reset() {
	*x = QueryCodexMonsterBeKilledNumRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCodexMonsterBeKilledNumRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCodexMonsterBeKilledNumRsp) ProtoMessage() {}

func (x *QueryCodexMonsterBeKilledNumRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCodexMonsterBeKilledNumRsp.ProtoReflect.Descriptor instead.
func (*QueryCodexMonsterBeKilledNumRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{4}
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetCodexIdList() []uint32 {
	if x != nil {
		return x.CodexIdList
	}
	return nil
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetBeCapturedNumList() []uint32 {
	if x != nil {
		return x.BeCapturedNumList
	}
	return nil
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetBeKilledNumList() []uint32 {
	if x != nil {
		return x.BeKilledNumList
	}
	return nil
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type ViewCodexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeDataList []*CodexTypeData `protobuf:"bytes,10,rep,name=type_data_list,json=typeDataList,proto3" json:"type_data_list,omitempty"`
}

func (x *ViewCodexReq) Reset() {
	*x = ViewCodexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCodexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCodexReq) ProtoMessage() {}

func (x *ViewCodexReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCodexReq.ProtoReflect.Descriptor instead.
func (*ViewCodexReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{5}
}

func (x *ViewCodexReq) GetTypeDataList() []*CodexTypeData {
	if x != nil {
		return x.TypeDataList
	}
	return nil
}

type ViewCodexRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                  int32            `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PushTipsRewardList       []uint32         `protobuf:"varint,10,rep,packed,name=push_tips_reward_list,json=pushTipsRewardList,proto3" json:"push_tips_reward_list,omitempty"`
	RecentViewedPushtipsList []uint32         `protobuf:"varint,3,rep,packed,name=recent_viewed_pushtips_list,json=recentViewedPushtipsList,proto3" json:"recent_viewed_pushtips_list,omitempty"`
	TypeDataList             []*CodexTypeData `protobuf:"bytes,9,rep,name=type_data_list,json=typeDataList,proto3" json:"type_data_list,omitempty"`
	PushTipsReadList         []uint32         `protobuf:"varint,15,rep,packed,name=push_tips_read_list,json=pushTipsReadList,proto3" json:"push_tips_read_list,omitempty"`
}

func (x *ViewCodexRsp) Reset() {
	*x = ViewCodexRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCodexRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCodexRsp) ProtoMessage() {}

func (x *ViewCodexRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCodexRsp.ProtoReflect.Descriptor instead.
func (*ViewCodexRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{6}
}

func (x *ViewCodexRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ViewCodexRsp) GetPushTipsRewardList() []uint32 {
	if x != nil {
		return x.PushTipsRewardList
	}
	return nil
}

func (x *ViewCodexRsp) GetRecentViewedPushtipsList() []uint32 {
	if x != nil {
		return x.RecentViewedPushtipsList
	}
	return nil
}

func (x *ViewCodexRsp) GetTypeDataList() []*CodexTypeData {
	if x != nil {
		return x.TypeDataList
	}
	return nil
}

func (x *ViewCodexRsp) GetPushTipsReadList() []uint32 {
	if x != nil {
		return x.PushTipsReadList
	}
	return nil
}

type SetCodexPushtipsReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId  uint32 `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	CodexId uint32 `protobuf:"varint,14,opt,name=codex_id,json=codexId,proto3" json:"codex_id,omitempty"`
}

func (x *SetCodexPushtipsReadReq) Reset() {
	*x = SetCodexPushtipsReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCodexPushtipsReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCodexPushtipsReadReq) ProtoMessage() {}

func (x *SetCodexPushtipsReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCodexPushtipsReadReq.ProtoReflect.Descriptor instead.
func (*SetCodexPushtipsReadReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{7}
}

func (x *SetCodexPushtipsReadReq) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *SetCodexPushtipsReadReq) GetCodexId() uint32 {
	if x != nil {
		return x.CodexId
	}
	return 0
}

type SetCodexPushtipsReadRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TypeId  uint32 `protobuf:"varint,5,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	CodexId uint32 `protobuf:"varint,14,opt,name=codex_id,json=codexId,proto3" json:"codex_id,omitempty"`
}

func (x *SetCodexPushtipsReadRsp) Reset() {
	*x = SetCodexPushtipsReadRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_codex_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCodexPushtipsReadRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCodexPushtipsReadRsp) ProtoMessage() {}

func (x *SetCodexPushtipsReadRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_codex_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCodexPushtipsReadRsp.ProtoReflect.Descriptor instead.
func (*SetCodexPushtipsReadRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_codex_proto_rawDescGZIP(), []int{8}
}

func (x *SetCodexPushtipsReadRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SetCodexPushtipsReadRsp) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *SetCodexPushtipsReadRsp) GetCodexId() uint32 {
	if x != nil {
		return x.CodexId
	}
	return 0
}

var File_cmd_cmd_codex_proto protoreflect.FileDescriptor

var file_cmd_cmd_codex_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x0d, 0x43,
	0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d,
	0x63, 0x6f, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x72, 0x0a, 0x1c, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x57, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x77, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x4d, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x78,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x61,
	0x76, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x4b, 0x0a, 0x1d, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x4d, 0x61,
	0x78, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x8a, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x46,
	0x75, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3c, 0x0a, 0x1b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x74, 0x69, 0x70, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x73, 0x68, 0x74, 0x69, 0x70, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x78, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x74, 0x69, 0x70,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x18, 0x72, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x50, 0x75, 0x73, 0x68, 0x74, 0x69,
	0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x1a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x6c, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x73, 0x68, 0x74, 0x69, 0x70, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x86,
	0x01, 0x0a, 0x15, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x77, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x77, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x78, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x65, 0x4b, 0x69,
	0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f,
	0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xbd,
	0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x42, 0x65, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x52,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x78,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x62, 0x65, 0x5f, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x62, 0x65, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64,
	0x4e, 0x75, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x62, 0x65, 0x5f, 0x6b, 0x69,
	0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x65, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4a,
	0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x12, 0x3a,
	0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x74, 0x79,
	0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x85, 0x02, 0x0a, 0x0c, 0x56,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x69,
	0x70, 0x73, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x74, 0x69,
	0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x18, 0x72,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x50, 0x75, 0x73, 0x68, 0x74,
	0x69, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x10, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x4d, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x50, 0x75,
	0x73, 0x68, 0x74, 0x69, 0x70, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x5f,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x49,
	0x64, 0x22, 0x67, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x50, 0x75, 0x73,
	0x68, 0x74, 0x69, 0x70, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_codex_proto_rawDescOnce sync.Once
	file_cmd_cmd_codex_proto_rawDescData = file_cmd_cmd_codex_proto_rawDesc
)

func file_cmd_cmd_codex_proto_rawDescGZIP() []byte {
	file_cmd_cmd_codex_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_codex_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_codex_proto_rawDescData)
	})
	return file_cmd_cmd_codex_proto_rawDescData
}

var file_cmd_cmd_codex_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cmd_cmd_codex_proto_goTypes = []interface{}{
	(*CodexTypeData)(nil),                   // 0: proto.CodexTypeData
	(*CodexDataFullNotify)(nil),             // 1: proto.CodexDataFullNotify
	(*CodexDataUpdateNotify)(nil),           // 2: proto.CodexDataUpdateNotify
	(*QueryCodexMonsterBeKilledNumReq)(nil), // 3: proto.QueryCodexMonsterBeKilledNumReq
	(*QueryCodexMonsterBeKilledNumRsp)(nil), // 4: proto.QueryCodexMonsterBeKilledNumRsp
	(*ViewCodexReq)(nil),                    // 5: proto.ViewCodexReq
	(*ViewCodexRsp)(nil),                    // 6: proto.ViewCodexRsp
	(*SetCodexPushtipsReadReq)(nil),         // 7: proto.SetCodexPushtipsReadReq
	(*SetCodexPushtipsReadRsp)(nil),         // 8: proto.SetCodexPushtipsReadRsp
	nil,                                     // 9: proto.CodexTypeData.WeaponMaxPromoteLevelMapEntry
	(CodexType)(0),                          // 10: proto.CodexType
}
var file_cmd_cmd_codex_proto_depIdxs = []int32{
	9,  // 0: proto.CodexTypeData.weapon_max_promote_level_map:type_name -> proto.CodexTypeData.WeaponMaxPromoteLevelMapEntry
	10, // 1: proto.CodexTypeData.type:type_name -> proto.CodexType
	0,  // 2: proto.CodexDataFullNotify.type_data_list:type_name -> proto.CodexTypeData
	10, // 3: proto.CodexDataUpdateNotify.type:type_name -> proto.CodexType
	0,  // 4: proto.ViewCodexReq.type_data_list:type_name -> proto.CodexTypeData
	0,  // 5: proto.ViewCodexRsp.type_data_list:type_name -> proto.CodexTypeData
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_cmd_codex_proto_init() }
func file_cmd_cmd_codex_proto_init() {
	if File_cmd_cmd_codex_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_cmd_codex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodexTypeData); i {
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
		file_cmd_cmd_codex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodexDataFullNotify); i {
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
		file_cmd_cmd_codex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodexDataUpdateNotify); i {
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
		file_cmd_cmd_codex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCodexMonsterBeKilledNumReq); i {
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
		file_cmd_cmd_codex_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCodexMonsterBeKilledNumRsp); i {
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
		file_cmd_cmd_codex_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCodexReq); i {
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
		file_cmd_cmd_codex_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCodexRsp); i {
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
		file_cmd_cmd_codex_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCodexPushtipsReadReq); i {
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
		file_cmd_cmd_codex_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCodexPushtipsReadRsp); i {
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
			RawDescriptor: file_cmd_cmd_codex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_codex_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_codex_proto_depIdxs,
		MessageInfos:      file_cmd_cmd_codex_proto_msgTypes,
	}.Build()
	File_cmd_cmd_codex_proto = out.File
	file_cmd_cmd_codex_proto_rawDesc = nil
	file_cmd_cmd_codex_proto_goTypes = nil
	file_cmd_cmd_codex_proto_depIdxs = nil
}
