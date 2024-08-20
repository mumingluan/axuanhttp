
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetGachaInfoReq) },
		func() ProtoMessage { return new(GetGachaInfoRsp) },
		func() ProtoMessage { return new(DoGachaReq) },
		func() ProtoMessage { return new(DoGachaRsp) },
		func() ProtoMessage { return new(GachaWishReq) },
		func() ProtoMessage { return new(GachaWishRsp) },
		func() ProtoMessage { return new(GachaOpenWishNotify) },
		func() ProtoMessage { return new(GachaSimpleInfoNotify) },
	)
}

const (
	ProtoCommandGetGachaInfoReq       ProtoCommand = 1572
	ProtoCommandGetGachaInfoRsp       ProtoCommand = 1598
	ProtoCommandDoGachaReq            ProtoCommand = 1512
	ProtoCommandDoGachaRsp            ProtoCommand = 1535
	ProtoCommandGachaWishReq          ProtoCommand = 1507
	ProtoCommandGachaWishRsp          ProtoCommand = 1521
	ProtoCommandGachaOpenWishNotify   ProtoCommand = 1503
	ProtoCommandGachaSimpleInfoNotify ProtoCommand = 1590
)

func (*GetGachaInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetGachaInfoReq }
func (*GetGachaInfoReq) ProtoMessageType() ProtoMessageType { return "GetGachaInfoReq" }

func (*GetGachaInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetGachaInfoRsp }
func (*GetGachaInfoRsp) ProtoMessageType() ProtoMessageType { return "GetGachaInfoRsp" }

func (*DoGachaReq) ProtoCommand() ProtoCommand         { return ProtoCommandDoGachaReq }
func (*DoGachaReq) ProtoMessageType() ProtoMessageType { return "DoGachaReq" }

func (*DoGachaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDoGachaRsp }
func (*DoGachaRsp) ProtoMessageType() ProtoMessageType { return "DoGachaRsp" }

func (*GachaWishReq) ProtoCommand() ProtoCommand         { return ProtoCommandGachaWishReq }
func (*GachaWishReq) ProtoMessageType() ProtoMessageType { return "GachaWishReq" }

func (*GachaWishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGachaWishRsp }
func (*GachaWishRsp) ProtoMessageType() ProtoMessageType { return "GachaWishRsp" }

func (*GachaOpenWishNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGachaOpenWishNotify }
func (*GachaOpenWishNotify) ProtoMessageType() ProtoMessageType { return "GachaOpenWishNotify" }

func (*GachaSimpleInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGachaSimpleInfoNotify }
func (*GachaSimpleInfoNotify) ProtoMessageType() ProtoMessageType { return "GachaSimpleInfoNotify" }