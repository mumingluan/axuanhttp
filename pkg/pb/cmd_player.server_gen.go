
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(SavePlayerDataReq) },
		func() ProtoMessage { return new(SavePlayerDataRsp) },
		func() ProtoMessage { return new(DisconnectClientNotify) },
		func() ProtoMessage { return new(SysCreateGroupReq) },
		func() ProtoMessage { return new(SaveBlockDataReq) },
		func() ProtoMessage { return new(SaveBlockDataRsp) },
		func() ProtoMessage { return new(SavePlayerExtraBinDataNotify) },
		func() ProtoMessage { return new(SysSavePlayerNotify) },
		func() ProtoMessage { return new(PlayerLoginBlockInfoNotify) },
	)
}

const (
	ProtoCommandSavePlayerDataReq            ProtoCommand = 10198
	ProtoCommandSavePlayerDataRsp            ProtoCommand = 10112
	ProtoCommandDisconnectClientNotify       ProtoCommand = 10173
	ProtoCommandSysCreateGroupReq            ProtoCommand = 10199
	ProtoCommandSaveBlockDataReq             ProtoCommand = 10131
	ProtoCommandSaveBlockDataRsp             ProtoCommand = 10175
	ProtoCommandSavePlayerExtraBinDataNotify ProtoCommand = 10148
	ProtoCommandSysSavePlayerNotify          ProtoCommand = 10197
	ProtoCommandPlayerLoginBlockInfoNotify   ProtoCommand = 10181
)

func (*SavePlayerDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandSavePlayerDataReq }
func (*SavePlayerDataReq) ProtoMessageType() ProtoMessageType { return "SavePlayerDataReq" }

func (*SavePlayerDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSavePlayerDataRsp }
func (*SavePlayerDataRsp) ProtoMessageType() ProtoMessageType { return "SavePlayerDataRsp" }

func (*DisconnectClientNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDisconnectClientNotify }
func (*DisconnectClientNotify) ProtoMessageType() ProtoMessageType { return "DisconnectClientNotify" }

func (*SysCreateGroupReq) ProtoCommand() ProtoCommand         { return ProtoCommandSysCreateGroupReq }
func (*SysCreateGroupReq) ProtoMessageType() ProtoMessageType { return "SysCreateGroupReq" }

func (*SaveBlockDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandSaveBlockDataReq }
func (*SaveBlockDataReq) ProtoMessageType() ProtoMessageType { return "SaveBlockDataReq" }

func (*SaveBlockDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSaveBlockDataRsp }
func (*SaveBlockDataRsp) ProtoMessageType() ProtoMessageType { return "SaveBlockDataRsp" }

func (*SavePlayerExtraBinDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSavePlayerExtraBinDataNotify
}
func (*SavePlayerExtraBinDataNotify) ProtoMessageType() ProtoMessageType {
	return "SavePlayerExtraBinDataNotify"
}

func (*SysSavePlayerNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSysSavePlayerNotify }
func (*SysSavePlayerNotify) ProtoMessageType() ProtoMessageType { return "SysSavePlayerNotify" }

func (*PlayerLoginBlockInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerLoginBlockInfoNotify
}
func (*PlayerLoginBlockInfoNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerLoginBlockInfoNotify"
}
