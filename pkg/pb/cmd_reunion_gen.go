
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(ReunionBriefInfoReq) },
		func() ProtoMessage { return new(ReunionBriefInfoRsp) },
		func() ProtoMessage { return new(TakeReunionFirstGiftRewardReq) },
		func() ProtoMessage { return new(TakeReunionFirstGiftRewardRsp) },
		func() ProtoMessage { return new(GetReunionSignInInfoReq) },
		func() ProtoMessage { return new(GetReunionSignInInfoRsp) },
		func() ProtoMessage { return new(TakeReunionSignInRewardReq) },
		func() ProtoMessage { return new(TakeReunionSignInRewardRsp) },
		func() ProtoMessage { return new(GetReunionMissionInfoReq) },
		func() ProtoMessage { return new(GetReunionMissionInfoRsp) },
		func() ProtoMessage { return new(TakeReunionWatcherRewardReq) },
		func() ProtoMessage { return new(TakeReunionWatcherRewardRsp) },
		func() ProtoMessage { return new(UpdateReunionWatcherNotify) },
		func() ProtoMessage { return new(TakeReunionMissionRewardReq) },
		func() ProtoMessage { return new(TakeReunionMissionRewardRsp) },
		func() ProtoMessage { return new(GetReunionPrivilegeInfoReq) },
		func() ProtoMessage { return new(GetReunionPrivilegeInfoRsp) },
		func() ProtoMessage { return new(ReunionSettleNotify) },
		func() ProtoMessage { return new(ReunionActivateNotify) },
		func() ProtoMessage { return new(ReunionPrivilegeChangeNotify) },
		func() ProtoMessage { return new(ReunionDailyRefreshNotify) },
	)
}

const (
	ProtoCommandReunionBriefInfoReq           ProtoCommand = 5076
	ProtoCommandReunionBriefInfoRsp           ProtoCommand = 5068
	ProtoCommandTakeReunionFirstGiftRewardReq ProtoCommand = 5075
	ProtoCommandTakeReunionFirstGiftRewardRsp ProtoCommand = 5057
	ProtoCommandGetReunionSignInInfoReq       ProtoCommand = 5052
	ProtoCommandGetReunionSignInInfoRsp       ProtoCommand = 5081
	ProtoCommandTakeReunionSignInRewardReq    ProtoCommand = 5079
	ProtoCommandTakeReunionSignInRewardRsp    ProtoCommand = 5072
	ProtoCommandGetReunionMissionInfoReq      ProtoCommand = 5094
	ProtoCommandGetReunionMissionInfoRsp      ProtoCommand = 5099
	ProtoCommandTakeReunionWatcherRewardReq   ProtoCommand = 5070
	ProtoCommandTakeReunionWatcherRewardRsp   ProtoCommand = 5095
	ProtoCommandUpdateReunionWatcherNotify    ProtoCommand = 5091
	ProtoCommandTakeReunionMissionRewardReq   ProtoCommand = 5092
	ProtoCommandTakeReunionMissionRewardRsp   ProtoCommand = 5064
	ProtoCommandGetReunionPrivilegeInfoReq    ProtoCommand = 5097
	ProtoCommandGetReunionPrivilegeInfoRsp    ProtoCommand = 5087
	ProtoCommandReunionSettleNotify           ProtoCommand = 5073
	ProtoCommandReunionActivateNotify         ProtoCommand = 5085
	ProtoCommandReunionPrivilegeChangeNotify  ProtoCommand = 5098
	ProtoCommandReunionDailyRefreshNotify     ProtoCommand = 5100
)

func (*ReunionBriefInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandReunionBriefInfoReq }
func (*ReunionBriefInfoReq) ProtoMessageType() ProtoMessageType { return "ReunionBriefInfoReq" }

func (*ReunionBriefInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReunionBriefInfoRsp }
func (*ReunionBriefInfoRsp) ProtoMessageType() ProtoMessageType { return "ReunionBriefInfoRsp" }

func (*TakeReunionFirstGiftRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionFirstGiftRewardReq
}
func (*TakeReunionFirstGiftRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeReunionFirstGiftRewardReq"
}

func (*TakeReunionFirstGiftRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionFirstGiftRewardRsp
}
func (*TakeReunionFirstGiftRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeReunionFirstGiftRewardRsp"
}

func (*GetReunionSignInInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetReunionSignInInfoReq
}
func (*GetReunionSignInInfoReq) ProtoMessageType() ProtoMessageType { return "GetReunionSignInInfoReq" }

func (*GetReunionSignInInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetReunionSignInInfoRsp
}
func (*GetReunionSignInInfoRsp) ProtoMessageType() ProtoMessageType { return "GetReunionSignInInfoRsp" }

func (*TakeReunionSignInRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionSignInRewardReq
}
func (*TakeReunionSignInRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeReunionSignInRewardReq"
}

func (*TakeReunionSignInRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionSignInRewardRsp
}
func (*TakeReunionSignInRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeReunionSignInRewardRsp"
}

func (*GetReunionMissionInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetReunionMissionInfoReq
}
func (*GetReunionMissionInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetReunionMissionInfoReq"
}

func (*GetReunionMissionInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetReunionMissionInfoRsp
}
func (*GetReunionMissionInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetReunionMissionInfoRsp"
}

func (*TakeReunionWatcherRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionWatcherRewardReq
}
func (*TakeReunionWatcherRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeReunionWatcherRewardReq"
}

func (*TakeReunionWatcherRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionWatcherRewardRsp
}
func (*TakeReunionWatcherRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeReunionWatcherRewardRsp"
}

func (*UpdateReunionWatcherNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdateReunionWatcherNotify
}
func (*UpdateReunionWatcherNotify) ProtoMessageType() ProtoMessageType {
	return "UpdateReunionWatcherNotify"
}

func (*TakeReunionMissionRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionMissionRewardReq
}
func (*TakeReunionMissionRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeReunionMissionRewardReq"
}

func (*TakeReunionMissionRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeReunionMissionRewardRsp
}
func (*TakeReunionMissionRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeReunionMissionRewardRsp"
}

func (*GetReunionPrivilegeInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetReunionPrivilegeInfoReq
}
func (*GetReunionPrivilegeInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetReunionPrivilegeInfoReq"
}

func (*GetReunionPrivilegeInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetReunionPrivilegeInfoRsp
}
func (*GetReunionPrivilegeInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetReunionPrivilegeInfoRsp"
}

func (*ReunionSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandReunionSettleNotify }
func (*ReunionSettleNotify) ProtoMessageType() ProtoMessageType { return "ReunionSettleNotify" }

func (*ReunionActivateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandReunionActivateNotify }
func (*ReunionActivateNotify) ProtoMessageType() ProtoMessageType { return "ReunionActivateNotify" }

func (*ReunionPrivilegeChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandReunionPrivilegeChangeNotify
}
func (*ReunionPrivilegeChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ReunionPrivilegeChangeNotify"
}

func (*ReunionDailyRefreshNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandReunionDailyRefreshNotify
}
func (*ReunionDailyRefreshNotify) ProtoMessageType() ProtoMessageType {
	return "ReunionDailyRefreshNotify"
}
