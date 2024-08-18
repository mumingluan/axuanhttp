
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(RegionSearchNotify) },
		func() ProtoMessage { return new(RegionSearchChangeRegionNotify) },
		func() ProtoMessage { return new(TakeRegionSearchRewardReq) },
		func() ProtoMessage { return new(TakeRegionSearchRewardRsp) },
		func() ProtoMessage { return new(GetRegionSearchReq) },
	)
}

const (
	ProtoCommandRegionSearchNotify             ProtoCommand = 5626
	ProtoCommandRegionSearchChangeRegionNotify ProtoCommand = 5618
	ProtoCommandTakeRegionSearchRewardReq      ProtoCommand = 5625
	ProtoCommandTakeRegionSearchRewardRsp      ProtoCommand = 5607
	ProtoCommandGetRegionSearchReq             ProtoCommand = 5602
)

func (*RegionSearchNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRegionSearchNotify }
func (*RegionSearchNotify) ProtoMessageType() ProtoMessageType { return "RegionSearchNotify" }

func (*RegionSearchChangeRegionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRegionSearchChangeRegionNotify
}
func (*RegionSearchChangeRegionNotify) ProtoMessageType() ProtoMessageType {
	return "RegionSearchChangeRegionNotify"
}

func (*TakeRegionSearchRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeRegionSearchRewardReq
}
func (*TakeRegionSearchRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeRegionSearchRewardReq"
}

func (*TakeRegionSearchRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeRegionSearchRewardRsp
}
func (*TakeRegionSearchRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeRegionSearchRewardRsp"
}

func (*GetRegionSearchReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetRegionSearchReq }
func (*GetRegionSearchReq) ProtoMessageType() ProtoMessageType { return "GetRegionSearchReq" }
