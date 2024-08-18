
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(StopServerConfigNotify) },
		func() ProtoMessage { return new(NodeserverConnectedAndRegisteredNotify) },
		func() ProtoMessage { return new(MultiPlayerMsg) },
		func() ProtoMessage { return new(AddGateserverNotify) },
		func() ProtoMessage { return new(RegisterServiceNotify) },
		func() ProtoMessage { return new(PlayerTransferNotify) },
		func() ProtoMessage { return new(PacketFreqencyExceedNotify) },
		func() ProtoMessage { return new(SceneAsyncLoadGroupBatchNotify) },
		func() ProtoMessage { return new(ClientVersionSyncNotify) },
		func() ProtoMessage { return new(RegisterServiceSuccessNotify) },
		func() ProtoMessage { return new(ReloadConfigNotify) },
		func() ProtoMessage { return new(MultiserverServiceTypeNotify) },
		func() ProtoMessage { return new(MatchServiceStopNotify) },
		func() ProtoMessage { return new(MatchServiceStopImplementNotify) },
		func() ProtoMessage { return new(ServerBlockPlayerMpNotify) },
		func() ProtoMessage { return new(ServerBlockPlayerChatNotify) },
		func() ProtoMessage { return new(ServerCheckSegmentCrcNotify) },
	)
}

const (
	ProtoCommandStopServerConfigNotify                 ProtoCommand = 10072
	ProtoCommandNodeserverConnectedAndRegisteredNotify ProtoCommand = 10098
	ProtoCommandMultiPlayerMsg                         ProtoCommand = 10012
	ProtoCommandAddGateserverNotify                    ProtoCommand = 10035
	ProtoCommandRegisterServiceNotify                  ProtoCommand = 10007
	ProtoCommandPlayerTransferNotify                   ProtoCommand = 10021
	ProtoCommandPacketFreqencyExceedNotify             ProtoCommand = 10003
	ProtoCommandSceneAsyncLoadGroupBatchNotify         ProtoCommand = 10090
	ProtoCommandClientVersionSyncNotify                ProtoCommand = 10073
	ProtoCommandRegisterServiceSuccessNotify           ProtoCommand = 10099
	ProtoCommandReloadConfigNotify                     ProtoCommand = 10031
	ProtoCommandMultiserverServiceTypeNotify           ProtoCommand = 10075
	ProtoCommandMatchServiceStopNotify                 ProtoCommand = 10048
	ProtoCommandMatchServiceStopImplementNotify        ProtoCommand = 10097
	ProtoCommandServerBlockPlayerMpNotify              ProtoCommand = 10081
	ProtoCommandServerBlockPlayerChatNotify            ProtoCommand = 10005
	ProtoCommandServerCheckSegmentCrcNotify            ProtoCommand = 10082
)

func (*StopServerConfigNotify) ProtoCommand() ProtoCommand         { return ProtoCommandStopServerConfigNotify }
func (*StopServerConfigNotify) ProtoMessageType() ProtoMessageType { return "StopServerConfigNotify" }

func (*NodeserverConnectedAndRegisteredNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandNodeserverConnectedAndRegisteredNotify
}
func (*NodeserverConnectedAndRegisteredNotify) ProtoMessageType() ProtoMessageType {
	return "NodeserverConnectedAndRegisteredNotify"
}

func (*MultiPlayerMsg) ProtoCommand() ProtoCommand         { return ProtoCommandMultiPlayerMsg }
func (*MultiPlayerMsg) ProtoMessageType() ProtoMessageType { return "MultiPlayerMsg" }

func (*AddGateserverNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAddGateserverNotify }
func (*AddGateserverNotify) ProtoMessageType() ProtoMessageType { return "AddGateserverNotify" }

func (*RegisterServiceNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRegisterServiceNotify }
func (*RegisterServiceNotify) ProtoMessageType() ProtoMessageType { return "RegisterServiceNotify" }

func (*PlayerTransferNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerTransferNotify }
func (*PlayerTransferNotify) ProtoMessageType() ProtoMessageType { return "PlayerTransferNotify" }

func (*PacketFreqencyExceedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPacketFreqencyExceedNotify
}
func (*PacketFreqencyExceedNotify) ProtoMessageType() ProtoMessageType {
	return "PacketFreqencyExceedNotify"
}

func (*SceneAsyncLoadGroupBatchNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneAsyncLoadGroupBatchNotify
}
func (*SceneAsyncLoadGroupBatchNotify) ProtoMessageType() ProtoMessageType {
	return "SceneAsyncLoadGroupBatchNotify"
}

func (*ClientVersionSyncNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientVersionSyncNotify
}
func (*ClientVersionSyncNotify) ProtoMessageType() ProtoMessageType { return "ClientVersionSyncNotify" }

func (*RegisterServiceSuccessNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRegisterServiceSuccessNotify
}
func (*RegisterServiceSuccessNotify) ProtoMessageType() ProtoMessageType {
	return "RegisterServiceSuccessNotify"
}

func (*ReloadConfigNotify) ProtoCommand() ProtoCommand         { return ProtoCommandReloadConfigNotify }
func (*ReloadConfigNotify) ProtoMessageType() ProtoMessageType { return "ReloadConfigNotify" }

func (*MultiserverServiceTypeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMultiserverServiceTypeNotify
}
func (*MultiserverServiceTypeNotify) ProtoMessageType() ProtoMessageType {
	return "MultiserverServiceTypeNotify"
}

func (*MatchServiceStopNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMatchServiceStopNotify }
func (*MatchServiceStopNotify) ProtoMessageType() ProtoMessageType { return "MatchServiceStopNotify" }

func (*MatchServiceStopImplementNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMatchServiceStopImplementNotify
}
func (*MatchServiceStopImplementNotify) ProtoMessageType() ProtoMessageType {
	return "MatchServiceStopImplementNotify"
}

func (*ServerBlockPlayerMpNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerBlockPlayerMpNotify
}
func (*ServerBlockPlayerMpNotify) ProtoMessageType() ProtoMessageType {
	return "ServerBlockPlayerMpNotify"
}

func (*ServerBlockPlayerChatNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerBlockPlayerChatNotify
}
func (*ServerBlockPlayerChatNotify) ProtoMessageType() ProtoMessageType {
	return "ServerBlockPlayerChatNotify"
}

func (*ServerCheckSegmentCrcNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCheckSegmentCrcNotify
}
func (*ServerCheckSegmentCrcNotify) ProtoMessageType() ProtoMessageType {
	return "ServerCheckSegmentCrcNotify"
}
