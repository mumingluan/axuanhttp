// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: server_only/cmd_recharge.server.proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid            uint32 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	ProductId      string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName    string `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductNum     uint32 `protobuf:"varint,5,opt,name=product_num,json=productNum,proto3" json:"product_num,omitempty"`
	CoinNum        uint32 `protobuf:"varint,6,opt,name=coin_num,json=coinNum,proto3" json:"coin_num,omitempty"`
	TotalFee       string `protobuf:"bytes,7,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	Currency       string `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	PriceTier      string `protobuf:"bytes,9,opt,name=price_tier,json=priceTier,proto3" json:"price_tier,omitempty"`
	TradeNo        string `protobuf:"bytes,10,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
	TradeTime      uint32 `protobuf:"varint,11,opt,name=trade_time,json=tradeTime,proto3" json:"trade_time,omitempty"`
	ChannelId      uint32 `protobuf:"varint,12,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChannelOrderNo string `protobuf:"bytes,13,opt,name=channel_order_no,json=channelOrderNo,proto3" json:"channel_order_no,omitempty"`
	PayPlat        string `protobuf:"bytes,14,opt,name=pay_plat,json=payPlat,proto3" json:"pay_plat,omitempty"`
	Extend         string `protobuf:"bytes,15,opt,name=extend,proto3" json:"extend,omitempty"`
	CreateTime     uint32 `protobuf:"varint,16,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Bonus          string `protobuf:"bytes,17,opt,name=bonus,proto3" json:"bonus,omitempty"`
	BonusNum       uint32 `protobuf:"varint,18,opt,name=bonus_num,json=bonusNum,proto3" json:"bonus_num,omitempty"`
	VipPointNum    uint32 `protobuf:"varint,19,opt,name=vip_point_num,json=vipPointNum,proto3" json:"vip_point_num,omitempty"`
	PayType        string `protobuf:"bytes,20,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	PayVendor      string `protobuf:"bytes,21,opt,name=pay_vendor,json=payVendor,proto3" json:"pay_vendor,omitempty"`
	ClientType     string `protobuf:"bytes,22,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	Device         string `protobuf:"bytes,23,opt,name=device,proto3" json:"device,omitempty"`
	ClientIp       string `protobuf:"bytes,24,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Order) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Order) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Order) GetProductNum() uint32 {
	if x != nil {
		return x.ProductNum
	}
	return 0
}

func (x *Order) GetCoinNum() uint32 {
	if x != nil {
		return x.CoinNum
	}
	return 0
}

func (x *Order) GetTotalFee() string {
	if x != nil {
		return x.TotalFee
	}
	return ""
}

func (x *Order) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Order) GetPriceTier() string {
	if x != nil {
		return x.PriceTier
	}
	return ""
}

func (x *Order) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *Order) GetTradeTime() uint32 {
	if x != nil {
		return x.TradeTime
	}
	return 0
}

func (x *Order) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *Order) GetChannelOrderNo() string {
	if x != nil {
		return x.ChannelOrderNo
	}
	return ""
}

func (x *Order) GetPayPlat() string {
	if x != nil {
		return x.PayPlat
	}
	return ""
}

func (x *Order) GetExtend() string {
	if x != nil {
		return x.Extend
	}
	return ""
}

func (x *Order) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Order) GetBonus() string {
	if x != nil {
		return x.Bonus
	}
	return ""
}

func (x *Order) GetBonusNum() uint32 {
	if x != nil {
		return x.BonusNum
	}
	return 0
}

func (x *Order) GetVipPointNum() uint32 {
	if x != nil {
		return x.VipPointNum
	}
	return 0
}

func (x *Order) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

func (x *Order) GetPayVendor() string {
	if x != nil {
		return x.PayVendor
	}
	return ""
}

func (x *Order) GetClientType() string {
	if x != nil {
		return x.ClientType
	}
	return ""
}

func (x *Order) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Order) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type RechargeOrderNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *RechargeOrderNotify) Reset() {
	*x = RechargeOrderNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RechargeOrderNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeOrderNotify) ProtoMessage() {}

func (x *RechargeOrderNotify) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeOrderNotify.ProtoReflect.Descriptor instead.
func (*RechargeOrderNotify) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{1}
}

func (x *RechargeOrderNotify) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type MarkOrderFinishedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    uint32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	FinishTime uint32 `protobuf:"varint,2,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	IsRetry    bool   `protobuf:"varint,3,opt,name=is_retry,json=isRetry,proto3" json:"is_retry,omitempty"`
}

func (x *MarkOrderFinishedReq) Reset() {
	*x = MarkOrderFinishedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkOrderFinishedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderFinishedReq) ProtoMessage() {}

func (x *MarkOrderFinishedReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderFinishedReq.ProtoReflect.Descriptor instead.
func (*MarkOrderFinishedReq) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{2}
}

func (x *MarkOrderFinishedReq) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *MarkOrderFinishedReq) GetFinishTime() uint32 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

func (x *MarkOrderFinishedReq) GetIsRetry() bool {
	if x != nil {
		return x.IsRetry
	}
	return false
}

type MarkOrderFinishedRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OrderId uint32 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *MarkOrderFinishedRsp) Reset() {
	*x = MarkOrderFinishedRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkOrderFinishedRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderFinishedRsp) ProtoMessage() {}

func (x *MarkOrderFinishedRsp) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderFinishedRsp.ProtoReflect.Descriptor instead.
func (*MarkOrderFinishedRsp) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{3}
}

func (x *MarkOrderFinishedRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MarkOrderFinishedRsp) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetUnfinishedOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUnfinishedOrderReq) Reset() {
	*x = GetUnfinishedOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnfinishedOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnfinishedOrderReq) ProtoMessage() {}

func (x *GetUnfinishedOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnfinishedOrderReq.ProtoReflect.Descriptor instead.
func (*GetUnfinishedOrderReq) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{4}
}

type GetUnfinishedOrderRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   int32    `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OrderList []*Order `protobuf:"bytes,2,rep,name=order_list,json=orderList,proto3" json:"order_list,omitempty"`
}

func (x *GetUnfinishedOrderRsp) Reset() {
	*x = GetUnfinishedOrderRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnfinishedOrderRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnfinishedOrderRsp) ProtoMessage() {}

func (x *GetUnfinishedOrderRsp) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnfinishedOrderRsp.ProtoReflect.Descriptor instead.
func (*GetUnfinishedOrderRsp) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetUnfinishedOrderRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetUnfinishedOrderRsp) GetOrderList() []*Order {
	if x != nil {
		return x.OrderList
	}
	return nil
}

type GetSpecificUnfinishedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetSpecificUnfinishedReq) Reset() {
	*x = GetSpecificUnfinishedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpecificUnfinishedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpecificUnfinishedReq) ProtoMessage() {}

func (x *GetSpecificUnfinishedReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpecificUnfinishedReq.ProtoReflect.Descriptor instead.
func (*GetSpecificUnfinishedReq) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{6}
}

func (x *GetSpecificUnfinishedReq) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetSpecificUnfinishedRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Order   *Order `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetSpecificUnfinishedRsp) Reset() {
	*x = GetSpecificUnfinishedRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_only_cmd_recharge_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpecificUnfinishedRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpecificUnfinishedRsp) ProtoMessage() {}

func (x *GetSpecificUnfinishedRsp) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_recharge_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpecificUnfinishedRsp.ProtoReflect.Descriptor instead.
func (*GetSpecificUnfinishedRsp) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_recharge_server_proto_rawDescGZIP(), []int{7}
}

func (x *GetSpecificUnfinishedRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetSpecificUnfinishedRsp) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_server_only_cmd_recharge_server_proto protoreflect.FileDescriptor

var file_server_only_cmd_recharge_server_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2f, 0x63, 0x6d,
	0x64, 0x5f, 0x72, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd,
	0x05, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x69, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x50, 0x6c, 0x61, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x76, 0x69,
	0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x76, 0x69, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79,
	0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x79, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x22, 0x39,
	0x0a, 0x13, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x6d, 0x0a, 0x14, 0x4d, 0x61, 0x72,
	0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x52, 0x65, 0x74, 0x72, 0x79, 0x22, 0x4b, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x5e,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x35,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x55, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x55, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_server_only_cmd_recharge_server_proto_rawDescOnce sync.Once
	file_server_only_cmd_recharge_server_proto_rawDescData = file_server_only_cmd_recharge_server_proto_rawDesc
)

func file_server_only_cmd_recharge_server_proto_rawDescGZIP() []byte {
	file_server_only_cmd_recharge_server_proto_rawDescOnce.Do(func() {
		file_server_only_cmd_recharge_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_only_cmd_recharge_server_proto_rawDescData)
	})
	return file_server_only_cmd_recharge_server_proto_rawDescData
}

var file_server_only_cmd_recharge_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_only_cmd_recharge_server_proto_goTypes = []interface{}{
	(*Order)(nil),                    // 0: proto.Order
	(*RechargeOrderNotify)(nil),      // 1: proto.RechargeOrderNotify
	(*MarkOrderFinishedReq)(nil),     // 2: proto.MarkOrderFinishedReq
	(*MarkOrderFinishedRsp)(nil),     // 3: proto.MarkOrderFinishedRsp
	(*GetUnfinishedOrderReq)(nil),    // 4: proto.GetUnfinishedOrderReq
	(*GetUnfinishedOrderRsp)(nil),    // 5: proto.GetUnfinishedOrderRsp
	(*GetSpecificUnfinishedReq)(nil), // 6: proto.GetSpecificUnfinishedReq
	(*GetSpecificUnfinishedRsp)(nil), // 7: proto.GetSpecificUnfinishedRsp
}
var file_server_only_cmd_recharge_server_proto_depIdxs = []int32{
	0, // 0: proto.RechargeOrderNotify.order:type_name -> proto.Order
	0, // 1: proto.GetUnfinishedOrderRsp.order_list:type_name -> proto.Order
	0, // 2: proto.GetSpecificUnfinishedRsp.order:type_name -> proto.Order
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_only_cmd_recharge_server_proto_init() }
func file_server_only_cmd_recharge_server_proto_init() {
	if File_server_only_cmd_recharge_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_only_cmd_recharge_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RechargeOrderNotify); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkOrderFinishedReq); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkOrderFinishedRsp); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnfinishedOrderReq); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnfinishedOrderRsp); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpecificUnfinishedReq); i {
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
		file_server_only_cmd_recharge_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpecificUnfinishedRsp); i {
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
			RawDescriptor: file_server_only_cmd_recharge_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_only_cmd_recharge_server_proto_goTypes,
		DependencyIndexes: file_server_only_cmd_recharge_server_proto_depIdxs,
		MessageInfos:      file_server_only_cmd_recharge_server_proto_msgTypes,
	}.Build()
	File_server_only_cmd_recharge_server_proto = out.File
	file_server_only_cmd_recharge_server_proto_rawDesc = nil
	file_server_only_cmd_recharge_server_proto_goTypes = nil
	file_server_only_cmd_recharge_server_proto_depIdxs = nil
}
