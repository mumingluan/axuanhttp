
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetShopReq) },
		func() ProtoMessage { return new(GetShopRsp) },
		func() ProtoMessage { return new(BuyGoodsReq) },
		func() ProtoMessage { return new(BuyGoodsRsp) },
		func() ProtoMessage { return new(GetShopmallDataReq) },
		func() ProtoMessage { return new(GetShopmallDataRsp) },
		func() ProtoMessage { return new(GetActivityShopSheetInfoReq) },
		func() ProtoMessage { return new(GetActivityShopSheetInfoRsp) },
		func() ProtoMessage { return new(BatchBuyGoodsReq) },
		func() ProtoMessage { return new(BatchBuyGoodsRsp) },
	)
}

const (
	ProtoCommandGetShopReq                  ProtoCommand = 772
	ProtoCommandGetShopRsp                  ProtoCommand = 798
	ProtoCommandBuyGoodsReq                 ProtoCommand = 712
	ProtoCommandBuyGoodsRsp                 ProtoCommand = 735
	ProtoCommandGetShopmallDataReq          ProtoCommand = 707
	ProtoCommandGetShopmallDataRsp          ProtoCommand = 721
	ProtoCommandGetActivityShopSheetInfoReq ProtoCommand = 703
	ProtoCommandGetActivityShopSheetInfoRsp ProtoCommand = 790
	ProtoCommandBatchBuyGoodsReq            ProtoCommand = 773
	ProtoCommandBatchBuyGoodsRsp            ProtoCommand = 799
)

func (*GetShopReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetShopReq }
func (*GetShopReq) ProtoMessageType() ProtoMessageType { return "GetShopReq" }

func (*GetShopRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetShopRsp }
func (*GetShopRsp) ProtoMessageType() ProtoMessageType { return "GetShopRsp" }

func (*BuyGoodsReq) ProtoCommand() ProtoCommand         { return ProtoCommandBuyGoodsReq }
func (*BuyGoodsReq) ProtoMessageType() ProtoMessageType { return "BuyGoodsReq" }

func (*BuyGoodsRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBuyGoodsRsp }
func (*BuyGoodsRsp) ProtoMessageType() ProtoMessageType { return "BuyGoodsRsp" }

func (*GetShopmallDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetShopmallDataReq }
func (*GetShopmallDataReq) ProtoMessageType() ProtoMessageType { return "GetShopmallDataReq" }

func (*GetShopmallDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetShopmallDataRsp }
func (*GetShopmallDataRsp) ProtoMessageType() ProtoMessageType { return "GetShopmallDataRsp" }

func (*GetActivityShopSheetInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetActivityShopSheetInfoReq
}
func (*GetActivityShopSheetInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetActivityShopSheetInfoReq"
}

func (*GetActivityShopSheetInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetActivityShopSheetInfoRsp
}
func (*GetActivityShopSheetInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetActivityShopSheetInfoRsp"
}

func (*BatchBuyGoodsReq) ProtoCommand() ProtoCommand         { return ProtoCommandBatchBuyGoodsReq }
func (*BatchBuyGoodsReq) ProtoMessageType() ProtoMessageType { return "BatchBuyGoodsReq" }

func (*BatchBuyGoodsRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBatchBuyGoodsRsp }
func (*BatchBuyGoodsRsp) ProtoMessageType() ProtoMessageType { return "BatchBuyGoodsRsp" }