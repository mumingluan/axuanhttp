
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(UpdateMpStatusNotify) },
		func() ProtoMessage { return new(DelMpStatusNotify) },
		func() ProtoMessage { return new(GetPlayerMpStatusListReq) },
		func() ProtoMessage { return new(GetPlayerMpStatusListRsp) },
		func() ProtoMessage { return new(GetPlayerMpStatusInfoReq) },
		func() ProtoMessage { return new(GetPlayerMpStatusInfoRsp) },
	)
}

const (
	ProtoCommandUpdateMpStatusNotify     ProtoCommand = 10472
	ProtoCommandDelMpStatusNotify        ProtoCommand = 10498
	ProtoCommandGetPlayerMpStatusListReq ProtoCommand = 10412
	ProtoCommandGetPlayerMpStatusListRsp ProtoCommand = 10435
	ProtoCommandGetPlayerMpStatusInfoReq ProtoCommand = 10407
	ProtoCommandGetPlayerMpStatusInfoRsp ProtoCommand = 10421
)

func (*UpdateMpStatusNotify) ProtoCommand() ProtoCommand         { return ProtoCommandUpdateMpStatusNotify }
func (*UpdateMpStatusNotify) ProtoMessageType() ProtoMessageType { return "UpdateMpStatusNotify" }

func (*DelMpStatusNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDelMpStatusNotify }
func (*DelMpStatusNotify) ProtoMessageType() ProtoMessageType { return "DelMpStatusNotify" }

func (*GetPlayerMpStatusListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerMpStatusListReq
}
func (*GetPlayerMpStatusListReq) ProtoMessageType() ProtoMessageType {
	return "GetPlayerMpStatusListReq"
}

func (*GetPlayerMpStatusListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerMpStatusListRsp
}
func (*GetPlayerMpStatusListRsp) ProtoMessageType() ProtoMessageType {
	return "GetPlayerMpStatusListRsp"
}

func (*GetPlayerMpStatusInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerMpStatusInfoReq
}
func (*GetPlayerMpStatusInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetPlayerMpStatusInfoReq"
}

func (*GetPlayerMpStatusInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerMpStatusInfoRsp
}
func (*GetPlayerMpStatusInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetPlayerMpStatusInfoRsp"
}
