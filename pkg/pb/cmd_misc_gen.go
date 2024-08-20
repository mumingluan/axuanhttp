
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(KeepAliveNotify) },
		func() ProtoMessage { return new(GmTalkReq) },
		func() ProtoMessage { return new(GmTalkRsp) },
		func() ProtoMessage { return new(ShowMessageNotify) },
		func() ProtoMessage { return new(PingReq) },
		func() ProtoMessage { return new(PingRsp) },
		func() ProtoMessage { return new(GetOnlinePlayerListReq) },
		func() ProtoMessage { return new(GetOnlinePlayerListRsp) },
		func() ProtoMessage { return new(ServerTimeNotify) },
		func() ProtoMessage { return new(ServerLogNotify) },
		func() ProtoMessage { return new(ClientReconnectNotify) },
		func() ProtoMessage { return new(RobotPushPlayerDataNotify) },
		func() ProtoMessage { return new(ClientReportNotify) },
		func() ProtoMessage { return new(UnionCmdNotify) },
		func() ProtoMessage { return new(GetOnlinePlayerInfoReq) },
		func() ProtoMessage { return new(GetOnlinePlayerInfoRsp) },
		func() ProtoMessage { return new(CheckSegmentCRCNotify) },
		func() ProtoMessage { return new(CheckSegmentCRCReq) },
		func() ProtoMessage { return new(WorldPlayerRTTNotify) },
		func() ProtoMessage { return new(EchoNotify) },
		func() ProtoMessage { return new(UpdateRedPointNotify) },
		func() ProtoMessage { return new(ClientBulletCreateNotify) },
		func() ProtoMessage { return new(ChangeServerGlobalValueNotify) },
		func() ProtoMessage { return new(GmTalkNotify) },
		func() ProtoMessage { return new(LastPacketPrintNotify) },
	)
}

const (
	ProtoCommandKeepAliveNotify               ProtoCommand = 72
	ProtoCommandGmTalkReq                     ProtoCommand = 98
	ProtoCommandGmTalkRsp                     ProtoCommand = 12
	ProtoCommandShowMessageNotify             ProtoCommand = 35
	ProtoCommandPingReq                       ProtoCommand = 7
	ProtoCommandPingRsp                       ProtoCommand = 21
	ProtoCommandGetOnlinePlayerListReq        ProtoCommand = 90
	ProtoCommandGetOnlinePlayerListRsp        ProtoCommand = 73
	ProtoCommandServerTimeNotify              ProtoCommand = 99
	ProtoCommandServerLogNotify               ProtoCommand = 31
	ProtoCommandClientReconnectNotify         ProtoCommand = 75
	ProtoCommandRobotPushPlayerDataNotify     ProtoCommand = 97
	ProtoCommandClientReportNotify            ProtoCommand = 81
	ProtoCommandUnionCmdNotify                ProtoCommand = 5
	ProtoCommandGetOnlinePlayerInfoReq        ProtoCommand = 82
	ProtoCommandGetOnlinePlayerInfoRsp        ProtoCommand = 47
	ProtoCommandCheckSegmentCRCNotify         ProtoCommand = 39
	ProtoCommandCheckSegmentCRCReq            ProtoCommand = 53
	ProtoCommandWorldPlayerRTTNotify          ProtoCommand = 22
	ProtoCommandEchoNotify                    ProtoCommand = 65
	ProtoCommandUpdateRedPointNotify          ProtoCommand = 93
	ProtoCommandClientBulletCreateNotify      ProtoCommand = 4
	ProtoCommandChangeServerGlobalValueNotify ProtoCommand = 27
	ProtoCommandGmTalkNotify                  ProtoCommand = 94
	ProtoCommandLastPacketPrintNotify         ProtoCommand = 88
)

func (*KeepAliveNotify) ProtoCommand() ProtoCommand         { return ProtoCommandKeepAliveNotify }
func (*KeepAliveNotify) ProtoMessageType() ProtoMessageType { return "KeepAliveNotify" }

func (*GmTalkReq) ProtoCommand() ProtoCommand         { return ProtoCommandGmTalkReq }
func (*GmTalkReq) ProtoMessageType() ProtoMessageType { return "GmTalkReq" }

func (*GmTalkRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGmTalkRsp }
func (*GmTalkRsp) ProtoMessageType() ProtoMessageType { return "GmTalkRsp" }

func (*ShowMessageNotify) ProtoCommand() ProtoCommand         { return ProtoCommandShowMessageNotify }
func (*ShowMessageNotify) ProtoMessageType() ProtoMessageType { return "ShowMessageNotify" }

func (*PingReq) ProtoCommand() ProtoCommand         { return ProtoCommandPingReq }
func (*PingReq) ProtoMessageType() ProtoMessageType { return "PingReq" }

func (*PingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPingRsp }
func (*PingRsp) ProtoMessageType() ProtoMessageType { return "PingRsp" }

func (*GetOnlinePlayerListReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetOnlinePlayerListReq }
func (*GetOnlinePlayerListReq) ProtoMessageType() ProtoMessageType { return "GetOnlinePlayerListReq" }

func (*GetOnlinePlayerListRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetOnlinePlayerListRsp }
func (*GetOnlinePlayerListRsp) ProtoMessageType() ProtoMessageType { return "GetOnlinePlayerListRsp" }

func (*ServerTimeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerTimeNotify }
func (*ServerTimeNotify) ProtoMessageType() ProtoMessageType { return "ServerTimeNotify" }

func (*ServerLogNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerLogNotify }
func (*ServerLogNotify) ProtoMessageType() ProtoMessageType { return "ServerLogNotify" }

func (*ClientReconnectNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClientReconnectNotify }
func (*ClientReconnectNotify) ProtoMessageType() ProtoMessageType { return "ClientReconnectNotify" }

func (*RobotPushPlayerDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRobotPushPlayerDataNotify
}
func (*RobotPushPlayerDataNotify) ProtoMessageType() ProtoMessageType {
	return "RobotPushPlayerDataNotify"
}

func (*ClientReportNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClientReportNotify }
func (*ClientReportNotify) ProtoMessageType() ProtoMessageType { return "ClientReportNotify" }

func (*UnionCmdNotify) ProtoCommand() ProtoCommand         { return ProtoCommandUnionCmdNotify }
func (*UnionCmdNotify) ProtoMessageType() ProtoMessageType { return "UnionCmdNotify" }

func (*GetOnlinePlayerInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetOnlinePlayerInfoReq }
func (*GetOnlinePlayerInfoReq) ProtoMessageType() ProtoMessageType { return "GetOnlinePlayerInfoReq" }

func (*GetOnlinePlayerInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetOnlinePlayerInfoRsp }
func (*GetOnlinePlayerInfoRsp) ProtoMessageType() ProtoMessageType { return "GetOnlinePlayerInfoRsp" }

func (*CheckSegmentCRCNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCheckSegmentCRCNotify }
func (*CheckSegmentCRCNotify) ProtoMessageType() ProtoMessageType { return "CheckSegmentCRCNotify" }

func (*CheckSegmentCRCReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckSegmentCRCReq }
func (*CheckSegmentCRCReq) ProtoMessageType() ProtoMessageType { return "CheckSegmentCRCReq" }

func (*WorldPlayerRTTNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWorldPlayerRTTNotify }
func (*WorldPlayerRTTNotify) ProtoMessageType() ProtoMessageType { return "WorldPlayerRTTNotify" }

func (*EchoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEchoNotify }
func (*EchoNotify) ProtoMessageType() ProtoMessageType { return "EchoNotify" }

func (*UpdateRedPointNotify) ProtoCommand() ProtoCommand         { return ProtoCommandUpdateRedPointNotify }
func (*UpdateRedPointNotify) ProtoMessageType() ProtoMessageType { return "UpdateRedPointNotify" }

func (*ClientBulletCreateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientBulletCreateNotify
}
func (*ClientBulletCreateNotify) ProtoMessageType() ProtoMessageType {
	return "ClientBulletCreateNotify"
}

func (*ChangeServerGlobalValueNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeServerGlobalValueNotify
}
func (*ChangeServerGlobalValueNotify) ProtoMessageType() ProtoMessageType {
	return "ChangeServerGlobalValueNotify"
}

func (*GmTalkNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGmTalkNotify }
func (*GmTalkNotify) ProtoMessageType() ProtoMessageType { return "GmTalkNotify" }

func (*LastPacketPrintNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLastPacketPrintNotify }
func (*LastPacketPrintNotify) ProtoMessageType() ProtoMessageType { return "LastPacketPrintNotify" }