
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AchievementAllDataNotify) },
		func() ProtoMessage { return new(AchievementUpdateNotify) },
		func() ProtoMessage { return new(TakeAchievementRewardReq) },
		func() ProtoMessage { return new(TakeAchievementRewardRsp) },
		func() ProtoMessage { return new(TakeAchievementGoalRewardReq) },
		func() ProtoMessage { return new(TakeAchievementGoalRewardRsp) },
	)
}

const (
	ProtoCommandAchievementAllDataNotify     ProtoCommand = 2676
	ProtoCommandAchievementUpdateNotify      ProtoCommand = 2668
	ProtoCommandTakeAchievementRewardReq     ProtoCommand = 2675
	ProtoCommandTakeAchievementRewardRsp     ProtoCommand = 2657
	ProtoCommandTakeAchievementGoalRewardReq ProtoCommand = 2652
	ProtoCommandTakeAchievementGoalRewardRsp ProtoCommand = 2681
)

func (*AchievementAllDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAchievementAllDataNotify
}
func (*AchievementAllDataNotify) ProtoMessageType() ProtoMessageType {
	return "AchievementAllDataNotify"
}

func (*AchievementUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAchievementUpdateNotify
}
func (*AchievementUpdateNotify) ProtoMessageType() ProtoMessageType { return "AchievementUpdateNotify" }

func (*TakeAchievementRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeAchievementRewardReq
}
func (*TakeAchievementRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeAchievementRewardReq"
}

func (*TakeAchievementRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeAchievementRewardRsp
}
func (*TakeAchievementRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeAchievementRewardRsp"
}

func (*TakeAchievementGoalRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeAchievementGoalRewardReq
}
func (*TakeAchievementGoalRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeAchievementGoalRewardReq"
}

func (*TakeAchievementGoalRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeAchievementGoalRewardRsp
}
func (*TakeAchievementGoalRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeAchievementGoalRewardRsp"
}
