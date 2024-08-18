
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(SeaLampPlayerContributionNotify) },
		func() ProtoMessage { return new(SeaLampProgressNotify) },
		func() ProtoMessage { return new(SeaLampBroadcastNotify) },
		func() ProtoMessage { return new(SeaLampSetProgressByMuipReq) },
		func() ProtoMessage { return new(SeaLampSetProgressByMuipRsp) },
		func() ProtoMessage { return new(SeaLampProgressImplementNotify) },
		func() ProtoMessage { return new(SeaLampClearProgressByGmNotify) },
		func() ProtoMessage { return new(SeaLampAddProgressByMuipReq) },
		func() ProtoMessage { return new(SeaLampAddProgressByMuipRsp) },
		func() ProtoMessage { return new(GetActivityDataByMuipReq) },
		func() ProtoMessage { return new(GetActivityDataByMuipRsp) },
	)
}

const (
	ProtoCommandSeaLampPlayerContributionNotify ProtoCommand = 10672
	ProtoCommandSeaLampProgressNotify           ProtoCommand = 10698
	ProtoCommandSeaLampBroadcastNotify          ProtoCommand = 10612
	ProtoCommandSeaLampSetProgressByMuipReq     ProtoCommand = 10635
	ProtoCommandSeaLampSetProgressByMuipRsp     ProtoCommand = 10607
	ProtoCommandSeaLampProgressImplementNotify  ProtoCommand = 10621
	ProtoCommandSeaLampClearProgressByGmNotify  ProtoCommand = 10603
	ProtoCommandSeaLampAddProgressByMuipReq     ProtoCommand = 10690
	ProtoCommandSeaLampAddProgressByMuipRsp     ProtoCommand = 10673
	ProtoCommandGetActivityDataByMuipReq        ProtoCommand = 10699
	ProtoCommandGetActivityDataByMuipRsp        ProtoCommand = 10631
)

func (*SeaLampPlayerContributionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampPlayerContributionNotify
}
func (*SeaLampPlayerContributionNotify) ProtoMessageType() ProtoMessageType {
	return "SeaLampPlayerContributionNotify"
}

func (*SeaLampProgressNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSeaLampProgressNotify }
func (*SeaLampProgressNotify) ProtoMessageType() ProtoMessageType { return "SeaLampProgressNotify" }

func (*SeaLampBroadcastNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSeaLampBroadcastNotify }
func (*SeaLampBroadcastNotify) ProtoMessageType() ProtoMessageType { return "SeaLampBroadcastNotify" }

func (*SeaLampSetProgressByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampSetProgressByMuipReq
}
func (*SeaLampSetProgressByMuipReq) ProtoMessageType() ProtoMessageType {
	return "SeaLampSetProgressByMuipReq"
}

func (*SeaLampSetProgressByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampSetProgressByMuipRsp
}
func (*SeaLampSetProgressByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "SeaLampSetProgressByMuipRsp"
}

func (*SeaLampProgressImplementNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampProgressImplementNotify
}
func (*SeaLampProgressImplementNotify) ProtoMessageType() ProtoMessageType {
	return "SeaLampProgressImplementNotify"
}

func (*SeaLampClearProgressByGmNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampClearProgressByGmNotify
}
func (*SeaLampClearProgressByGmNotify) ProtoMessageType() ProtoMessageType {
	return "SeaLampClearProgressByGmNotify"
}

func (*SeaLampAddProgressByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampAddProgressByMuipReq
}
func (*SeaLampAddProgressByMuipReq) ProtoMessageType() ProtoMessageType {
	return "SeaLampAddProgressByMuipReq"
}

func (*SeaLampAddProgressByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampAddProgressByMuipRsp
}
func (*SeaLampAddProgressByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "SeaLampAddProgressByMuipRsp"
}

func (*GetActivityDataByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetActivityDataByMuipReq
}
func (*GetActivityDataByMuipReq) ProtoMessageType() ProtoMessageType {
	return "GetActivityDataByMuipReq"
}

func (*GetActivityDataByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetActivityDataByMuipRsp
}
func (*GetActivityDataByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "GetActivityDataByMuipRsp"
}
