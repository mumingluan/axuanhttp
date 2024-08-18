
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetUgcReq) },
		func() ProtoMessage { return new(GetUgcRsp) },
		func() ProtoMessage { return new(GetUgcBriefInfoReq) },
		func() ProtoMessage { return new(GetUgcBriefInfoRsp) },
		func() ProtoMessage { return new(SaveUgcReq) },
		func() ProtoMessage { return new(SaveUgcRsp) },
		func() ProtoMessage { return new(PublishUgcReq) },
		func() ProtoMessage { return new(PublishUgcRsp) },
		func() ProtoMessage { return new(CheckUgcUpdateReq) },
		func() ProtoMessage { return new(CheckUgcUpdateRsp) },
		func() ProtoMessage { return new(UgcNotify) },
		func() ProtoMessage { return new(CheckUgcStateReq) },
		func() ProtoMessage { return new(CheckUgcStateRsp) },
	)
}

const (
	ProtoCommandGetUgcReq          ProtoCommand = 6326
	ProtoCommandGetUgcRsp          ProtoCommand = 6318
	ProtoCommandGetUgcBriefInfoReq ProtoCommand = 6325
	ProtoCommandGetUgcBriefInfoRsp ProtoCommand = 6307
	ProtoCommandSaveUgcReq         ProtoCommand = 6329
	ProtoCommandSaveUgcRsp         ProtoCommand = 6322
	ProtoCommandPublishUgcReq      ProtoCommand = 6344
	ProtoCommandPublishUgcRsp      ProtoCommand = 6349
	ProtoCommandCheckUgcUpdateReq  ProtoCommand = 6320
	ProtoCommandCheckUgcUpdateRsp  ProtoCommand = 6345
	ProtoCommandUgcNotify          ProtoCommand = 6341
	ProtoCommandCheckUgcStateReq   ProtoCommand = 6342
	ProtoCommandCheckUgcStateRsp   ProtoCommand = 6314
)

func (*GetUgcReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetUgcReq }
func (*GetUgcReq) ProtoMessageType() ProtoMessageType { return "GetUgcReq" }

func (*GetUgcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetUgcRsp }
func (*GetUgcRsp) ProtoMessageType() ProtoMessageType { return "GetUgcRsp" }

func (*GetUgcBriefInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetUgcBriefInfoReq }
func (*GetUgcBriefInfoReq) ProtoMessageType() ProtoMessageType { return "GetUgcBriefInfoReq" }

func (*GetUgcBriefInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetUgcBriefInfoRsp }
func (*GetUgcBriefInfoRsp) ProtoMessageType() ProtoMessageType { return "GetUgcBriefInfoRsp" }

func (*SaveUgcReq) ProtoCommand() ProtoCommand         { return ProtoCommandSaveUgcReq }
func (*SaveUgcReq) ProtoMessageType() ProtoMessageType { return "SaveUgcReq" }

func (*SaveUgcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSaveUgcRsp }
func (*SaveUgcRsp) ProtoMessageType() ProtoMessageType { return "SaveUgcRsp" }

func (*PublishUgcReq) ProtoCommand() ProtoCommand         { return ProtoCommandPublishUgcReq }
func (*PublishUgcReq) ProtoMessageType() ProtoMessageType { return "PublishUgcReq" }

func (*PublishUgcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPublishUgcRsp }
func (*PublishUgcRsp) ProtoMessageType() ProtoMessageType { return "PublishUgcRsp" }

func (*CheckUgcUpdateReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckUgcUpdateReq }
func (*CheckUgcUpdateReq) ProtoMessageType() ProtoMessageType { return "CheckUgcUpdateReq" }

func (*CheckUgcUpdateRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCheckUgcUpdateRsp }
func (*CheckUgcUpdateRsp) ProtoMessageType() ProtoMessageType { return "CheckUgcUpdateRsp" }

func (*UgcNotify) ProtoCommand() ProtoCommand         { return ProtoCommandUgcNotify }
func (*UgcNotify) ProtoMessageType() ProtoMessageType { return "UgcNotify" }

func (*CheckUgcStateReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckUgcStateReq }
func (*CheckUgcStateReq) ProtoMessageType() ProtoMessageType { return "CheckUgcStateReq" }

func (*CheckUgcStateRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCheckUgcStateRsp }
func (*CheckUgcStateRsp) ProtoMessageType() ProtoMessageType { return "CheckUgcStateRsp" }
