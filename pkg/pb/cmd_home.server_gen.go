
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetHomeDataReq) },
		func() ProtoMessage { return new(GetHomeDataRsp) },
		func() ProtoMessage { return new(SaveHomeDataReq) },
		func() ProtoMessage { return new(SaveHomeDataRsp) },
		func() ProtoMessage { return new(ServerTryEnterHomeReq) },
		func() ProtoMessage { return new(ServerTryEnterHomeRsp) },
		func() ProtoMessage { return new(ServerBlockHomeNotify) },
		func() ProtoMessage { return new(ServerGetFriendEnterHomeOptionReq) },
		func() ProtoMessage { return new(ServerGetFriendEnterHomeOptionRsp) },
		func() ProtoMessage { return new(SendHomeOfflineMsgReq) },
		func() ProtoMessage { return new(SendHomeOfflineMsgRsp) },
		func() ProtoMessage { return new(NewHomeOfflineMsgNotify) },
		func() ProtoMessage { return new(GetHomeOfflineMsgReq) },
		func() ProtoMessage { return new(GetHomeOfflineMsgRsp) },
		func() ProtoMessage { return new(RemoveHomeOfflineMsgNotify) },
		func() ProtoMessage { return new(ClearHomeOfflineMsgNotify) },
		func() ProtoMessage { return new(ServerHomeGetOnlineStatusReq) },
		func() ProtoMessage { return new(ServerHomeGetOnlineStatusRsp) },
		func() ProtoMessage { return new(ServerHomeKickPlayerReq) },
		func() ProtoMessage { return new(ServerHomeKickPlayerRsp) },
		func() ProtoMessage { return new(ServerHomeGetBlueprintBriefDataReq) },
		func() ProtoMessage { return new(ServerHomeGetBlueprintBriefDataRsp) },
		func() ProtoMessage { return new(ServerHomeUpdateBlueprintBriefDataReq) },
		func() ProtoMessage { return new(ServerHomeUpdateBlueprintBriefDataRsp) },
		func() ProtoMessage { return new(ServerHomeGetBlueprintDetailDataReq) },
		func() ProtoMessage { return new(ServerHomeGetBlueprintDetailDataRsp) },
		func() ProtoMessage { return new(ServerHomeSaveBlueprintDataReq) },
		func() ProtoMessage { return new(ServerHomeSaveBlueprintDataRsp) },
		func() ProtoMessage { return new(ServerHomeDeleteBlueprintDataReq) },
		func() ProtoMessage { return new(ServerHomeDeleteBlueprintDataRsp) },
		func() ProtoMessage { return new(ServerHomeBlueprintCacheInvalidNotify) },
	)
}

const (
	ProtoCommandGetHomeDataReq                        ProtoCommand = 12072
	ProtoCommandGetHomeDataRsp                        ProtoCommand = 12098
	ProtoCommandSaveHomeDataReq                       ProtoCommand = 12012
	ProtoCommandSaveHomeDataRsp                       ProtoCommand = 12035
	ProtoCommandServerTryEnterHomeReq                 ProtoCommand = 12007
	ProtoCommandServerTryEnterHomeRsp                 ProtoCommand = 12021
	ProtoCommandServerBlockHomeNotify                 ProtoCommand = 12003
	ProtoCommandServerGetFriendEnterHomeOptionReq     ProtoCommand = 12090
	ProtoCommandServerGetFriendEnterHomeOptionRsp     ProtoCommand = 12073
	ProtoCommandSendHomeOfflineMsgReq                 ProtoCommand = 12031
	ProtoCommandSendHomeOfflineMsgRsp                 ProtoCommand = 12075
	ProtoCommandNewHomeOfflineMsgNotify               ProtoCommand = 12048
	ProtoCommandGetHomeOfflineMsgReq                  ProtoCommand = 12097
	ProtoCommandGetHomeOfflineMsgRsp                  ProtoCommand = 12081
	ProtoCommandRemoveHomeOfflineMsgNotify            ProtoCommand = 12005
	ProtoCommandClearHomeOfflineMsgNotify             ProtoCommand = 12082
	ProtoCommandServerHomeGetOnlineStatusReq          ProtoCommand = 12047
	ProtoCommandServerHomeGetOnlineStatusRsp          ProtoCommand = 12039
	ProtoCommandServerHomeKickPlayerReq               ProtoCommand = 12053
	ProtoCommandServerHomeKickPlayerRsp               ProtoCommand = 12022
	ProtoCommandServerHomeGetBlueprintBriefDataReq    ProtoCommand = 12065
	ProtoCommandServerHomeGetBlueprintBriefDataRsp    ProtoCommand = 12004
	ProtoCommandServerHomeUpdateBlueprintBriefDataReq ProtoCommand = 12093
	ProtoCommandServerHomeUpdateBlueprintBriefDataRsp ProtoCommand = 12027
	ProtoCommandServerHomeGetBlueprintDetailDataReq   ProtoCommand = 12094
	ProtoCommandServerHomeGetBlueprintDetailDataRsp   ProtoCommand = 12088
	ProtoCommandServerHomeSaveBlueprintDataReq        ProtoCommand = 12026
	ProtoCommandServerHomeSaveBlueprintDataRsp        ProtoCommand = 12063
	ProtoCommandServerHomeDeleteBlueprintDataReq      ProtoCommand = 12095
	ProtoCommandServerHomeDeleteBlueprintDataRsp      ProtoCommand = 12006
	ProtoCommandServerHomeBlueprintCacheInvalidNotify ProtoCommand = 12034
)

func (*GetHomeDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetHomeDataReq }
func (*GetHomeDataReq) ProtoMessageType() ProtoMessageType { return "GetHomeDataReq" }

func (*GetHomeDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetHomeDataRsp }
func (*GetHomeDataRsp) ProtoMessageType() ProtoMessageType { return "GetHomeDataRsp" }

func (*SaveHomeDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandSaveHomeDataReq }
func (*SaveHomeDataReq) ProtoMessageType() ProtoMessageType { return "SaveHomeDataReq" }

func (*SaveHomeDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSaveHomeDataRsp }
func (*SaveHomeDataRsp) ProtoMessageType() ProtoMessageType { return "SaveHomeDataRsp" }

func (*ServerTryEnterHomeReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerTryEnterHomeReq }
func (*ServerTryEnterHomeReq) ProtoMessageType() ProtoMessageType { return "ServerTryEnterHomeReq" }

func (*ServerTryEnterHomeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerTryEnterHomeRsp }
func (*ServerTryEnterHomeRsp) ProtoMessageType() ProtoMessageType { return "ServerTryEnterHomeRsp" }

func (*ServerBlockHomeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerBlockHomeNotify }
func (*ServerBlockHomeNotify) ProtoMessageType() ProtoMessageType { return "ServerBlockHomeNotify" }

func (*ServerGetFriendEnterHomeOptionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendEnterHomeOptionReq
}
func (*ServerGetFriendEnterHomeOptionReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetFriendEnterHomeOptionReq"
}

func (*ServerGetFriendEnterHomeOptionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendEnterHomeOptionRsp
}
func (*ServerGetFriendEnterHomeOptionRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetFriendEnterHomeOptionRsp"
}

func (*SendHomeOfflineMsgReq) ProtoCommand() ProtoCommand         { return ProtoCommandSendHomeOfflineMsgReq }
func (*SendHomeOfflineMsgReq) ProtoMessageType() ProtoMessageType { return "SendHomeOfflineMsgReq" }

func (*SendHomeOfflineMsgRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSendHomeOfflineMsgRsp }
func (*SendHomeOfflineMsgRsp) ProtoMessageType() ProtoMessageType { return "SendHomeOfflineMsgRsp" }

func (*NewHomeOfflineMsgNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandNewHomeOfflineMsgNotify
}
func (*NewHomeOfflineMsgNotify) ProtoMessageType() ProtoMessageType { return "NewHomeOfflineMsgNotify" }

func (*GetHomeOfflineMsgReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetHomeOfflineMsgReq }
func (*GetHomeOfflineMsgReq) ProtoMessageType() ProtoMessageType { return "GetHomeOfflineMsgReq" }

func (*GetHomeOfflineMsgRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetHomeOfflineMsgRsp }
func (*GetHomeOfflineMsgRsp) ProtoMessageType() ProtoMessageType { return "GetHomeOfflineMsgRsp" }

func (*RemoveHomeOfflineMsgNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRemoveHomeOfflineMsgNotify
}
func (*RemoveHomeOfflineMsgNotify) ProtoMessageType() ProtoMessageType {
	return "RemoveHomeOfflineMsgNotify"
}

func (*ClearHomeOfflineMsgNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClearHomeOfflineMsgNotify
}
func (*ClearHomeOfflineMsgNotify) ProtoMessageType() ProtoMessageType {
	return "ClearHomeOfflineMsgNotify"
}

func (*ServerHomeGetOnlineStatusReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeGetOnlineStatusReq
}
func (*ServerHomeGetOnlineStatusReq) ProtoMessageType() ProtoMessageType {
	return "ServerHomeGetOnlineStatusReq"
}

func (*ServerHomeGetOnlineStatusRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeGetOnlineStatusRsp
}
func (*ServerHomeGetOnlineStatusRsp) ProtoMessageType() ProtoMessageType {
	return "ServerHomeGetOnlineStatusRsp"
}

func (*ServerHomeKickPlayerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeKickPlayerReq
}
func (*ServerHomeKickPlayerReq) ProtoMessageType() ProtoMessageType { return "ServerHomeKickPlayerReq" }

func (*ServerHomeKickPlayerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeKickPlayerRsp
}
func (*ServerHomeKickPlayerRsp) ProtoMessageType() ProtoMessageType { return "ServerHomeKickPlayerRsp" }

func (*ServerHomeGetBlueprintBriefDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeGetBlueprintBriefDataReq
}
func (*ServerHomeGetBlueprintBriefDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerHomeGetBlueprintBriefDataReq"
}

func (*ServerHomeGetBlueprintBriefDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeGetBlueprintBriefDataRsp
}
func (*ServerHomeGetBlueprintBriefDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerHomeGetBlueprintBriefDataRsp"
}

func (*ServerHomeUpdateBlueprintBriefDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeUpdateBlueprintBriefDataReq
}
func (*ServerHomeUpdateBlueprintBriefDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerHomeUpdateBlueprintBriefDataReq"
}

func (*ServerHomeUpdateBlueprintBriefDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeUpdateBlueprintBriefDataRsp
}
func (*ServerHomeUpdateBlueprintBriefDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerHomeUpdateBlueprintBriefDataRsp"
}

func (*ServerHomeGetBlueprintDetailDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeGetBlueprintDetailDataReq
}
func (*ServerHomeGetBlueprintDetailDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerHomeGetBlueprintDetailDataReq"
}

func (*ServerHomeGetBlueprintDetailDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeGetBlueprintDetailDataRsp
}
func (*ServerHomeGetBlueprintDetailDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerHomeGetBlueprintDetailDataRsp"
}

func (*ServerHomeSaveBlueprintDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeSaveBlueprintDataReq
}
func (*ServerHomeSaveBlueprintDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerHomeSaveBlueprintDataReq"
}

func (*ServerHomeSaveBlueprintDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeSaveBlueprintDataRsp
}
func (*ServerHomeSaveBlueprintDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerHomeSaveBlueprintDataRsp"
}

func (*ServerHomeDeleteBlueprintDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeDeleteBlueprintDataReq
}
func (*ServerHomeDeleteBlueprintDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerHomeDeleteBlueprintDataReq"
}

func (*ServerHomeDeleteBlueprintDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeDeleteBlueprintDataRsp
}
func (*ServerHomeDeleteBlueprintDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerHomeDeleteBlueprintDataRsp"
}

func (*ServerHomeBlueprintCacheInvalidNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerHomeBlueprintCacheInvalidNotify
}
func (*ServerHomeBlueprintCacheInvalidNotify) ProtoMessageType() ProtoMessageType {
	return "ServerHomeBlueprintCacheInvalidNotify"
}
