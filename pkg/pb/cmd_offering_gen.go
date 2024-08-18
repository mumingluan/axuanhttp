
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerOfferingDataNotify) },
		func() ProtoMessage { return new(PlayerOfferingReq) },
		func() ProtoMessage { return new(PlayerOfferingRsp) },
		func() ProtoMessage { return new(TakeOfferingLevelRewardReq) },
		func() ProtoMessage { return new(TakeOfferingLevelRewardRsp) },
		func() ProtoMessage { return new(OfferingInteractReq) },
		func() ProtoMessage { return new(OfferingInteractRsp) },
	)
}

const (
	ProtoCommandPlayerOfferingDataNotify   ProtoCommand = 2923
	ProtoCommandPlayerOfferingReq          ProtoCommand = 2907
	ProtoCommandPlayerOfferingRsp          ProtoCommand = 2917
	ProtoCommandTakeOfferingLevelRewardReq ProtoCommand = 2919
	ProtoCommandTakeOfferingLevelRewardRsp ProtoCommand = 2911
	ProtoCommandOfferingInteractReq        ProtoCommand = 2918
	ProtoCommandOfferingInteractRsp        ProtoCommand = 2908
)

func (*PlayerOfferingDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerOfferingDataNotify
}
func (*PlayerOfferingDataNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerOfferingDataNotify"
}

func (*PlayerOfferingReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerOfferingReq }
func (*PlayerOfferingReq) ProtoMessageType() ProtoMessageType { return "PlayerOfferingReq" }

func (*PlayerOfferingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerOfferingRsp }
func (*PlayerOfferingRsp) ProtoMessageType() ProtoMessageType { return "PlayerOfferingRsp" }

func (*TakeOfferingLevelRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeOfferingLevelRewardReq
}
func (*TakeOfferingLevelRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeOfferingLevelRewardReq"
}

func (*TakeOfferingLevelRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeOfferingLevelRewardRsp
}
func (*TakeOfferingLevelRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeOfferingLevelRewardRsp"
}

func (*OfferingInteractReq) ProtoCommand() ProtoCommand         { return ProtoCommandOfferingInteractReq }
func (*OfferingInteractReq) ProtoMessageType() ProtoMessageType { return "OfferingInteractReq" }

func (*OfferingInteractRsp) ProtoCommand() ProtoCommand         { return ProtoCommandOfferingInteractRsp }
func (*OfferingInteractRsp) ProtoMessageType() ProtoMessageType { return "OfferingInteractRsp" }
