
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerInvestigationAllInfoNotify) },
		func() ProtoMessage { return new(TakeInvestigationRewardReq) },
		func() ProtoMessage { return new(TakeInvestigationRewardRsp) },
		func() ProtoMessage { return new(TakeInvestigationTargetRewardReq) },
		func() ProtoMessage { return new(TakeInvestigationTargetRewardRsp) },
		func() ProtoMessage { return new(GetInvestigationMonsterReq) },
		func() ProtoMessage { return new(GetInvestigationMonsterRsp) },
		func() ProtoMessage { return new(PlayerInvestigationNotify) },
		func() ProtoMessage { return new(PlayerInvestigationTargetNotify) },
		func() ProtoMessage { return new(MarkTargetInvestigationMonsterNotify) },
		func() ProtoMessage { return new(InvestigationMonsterUpdateNotify) },
		func() ProtoMessage { return new(InvestigationQuestDailyNotify) },
		func() ProtoMessage { return new(InvestigationReadQuestDailyNotify) },
	)
}

const (
	ProtoCommandPlayerInvestigationAllInfoNotify     ProtoCommand = 1928
	ProtoCommandTakeInvestigationRewardReq           ProtoCommand = 1912
	ProtoCommandTakeInvestigationRewardRsp           ProtoCommand = 1922
	ProtoCommandTakeInvestigationTargetRewardReq     ProtoCommand = 1918
	ProtoCommandTakeInvestigationTargetRewardRsp     ProtoCommand = 1916
	ProtoCommandGetInvestigationMonsterReq           ProtoCommand = 1901
	ProtoCommandGetInvestigationMonsterRsp           ProtoCommand = 1910
	ProtoCommandPlayerInvestigationNotify            ProtoCommand = 1911
	ProtoCommandPlayerInvestigationTargetNotify      ProtoCommand = 1929
	ProtoCommandMarkTargetInvestigationMonsterNotify ProtoCommand = 1915
	ProtoCommandInvestigationMonsterUpdateNotify     ProtoCommand = 1906
	ProtoCommandInvestigationQuestDailyNotify        ProtoCommand = 1921
	ProtoCommandInvestigationReadQuestDailyNotify    ProtoCommand = 1902
)

func (*PlayerInvestigationAllInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerInvestigationAllInfoNotify
}
func (*PlayerInvestigationAllInfoNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerInvestigationAllInfoNotify"
}

func (*TakeInvestigationRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeInvestigationRewardReq
}
func (*TakeInvestigationRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeInvestigationRewardReq"
}

func (*TakeInvestigationRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeInvestigationRewardRsp
}
func (*TakeInvestigationRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeInvestigationRewardRsp"
}

func (*TakeInvestigationTargetRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeInvestigationTargetRewardReq
}
func (*TakeInvestigationTargetRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeInvestigationTargetRewardReq"
}

func (*TakeInvestigationTargetRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeInvestigationTargetRewardRsp
}
func (*TakeInvestigationTargetRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeInvestigationTargetRewardRsp"
}

func (*GetInvestigationMonsterReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetInvestigationMonsterReq
}
func (*GetInvestigationMonsterReq) ProtoMessageType() ProtoMessageType {
	return "GetInvestigationMonsterReq"
}

func (*GetInvestigationMonsterRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetInvestigationMonsterRsp
}
func (*GetInvestigationMonsterRsp) ProtoMessageType() ProtoMessageType {
	return "GetInvestigationMonsterRsp"
}

func (*PlayerInvestigationNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerInvestigationNotify
}
func (*PlayerInvestigationNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerInvestigationNotify"
}

func (*PlayerInvestigationTargetNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerInvestigationTargetNotify
}
func (*PlayerInvestigationTargetNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerInvestigationTargetNotify"
}

func (*MarkTargetInvestigationMonsterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMarkTargetInvestigationMonsterNotify
}
func (*MarkTargetInvestigationMonsterNotify) ProtoMessageType() ProtoMessageType {
	return "MarkTargetInvestigationMonsterNotify"
}

func (*InvestigationMonsterUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInvestigationMonsterUpdateNotify
}
func (*InvestigationMonsterUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "InvestigationMonsterUpdateNotify"
}

func (*InvestigationQuestDailyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInvestigationQuestDailyNotify
}
func (*InvestigationQuestDailyNotify) ProtoMessageType() ProtoMessageType {
	return "InvestigationQuestDailyNotify"
}

func (*InvestigationReadQuestDailyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInvestigationReadQuestDailyNotify
}
func (*InvestigationReadQuestDailyNotify) ProtoMessageType() ProtoMessageType {
	return "InvestigationReadQuestDailyNotify"
}
