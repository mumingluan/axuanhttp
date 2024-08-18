
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AllCoopInfoNotify) },
		func() ProtoMessage { return new(MainCoopUpdateNotify) },
		func() ProtoMessage { return new(SaveMainCoopReq) },
		func() ProtoMessage { return new(SaveMainCoopRsp) },
		func() ProtoMessage { return new(FinishMainCoopReq) },
		func() ProtoMessage { return new(FinishMainCoopRsp) },
		func() ProtoMessage { return new(CoopDataNotify) },
		func() ProtoMessage { return new(CoopChapterUpdateNotify) },
		func() ProtoMessage { return new(CoopCgUpdateNotify) },
		func() ProtoMessage { return new(CoopRewardUpdateNotify) },
		func() ProtoMessage { return new(UnlockCoopChapterReq) },
		func() ProtoMessage { return new(UnlockCoopChapterRsp) },
		func() ProtoMessage { return new(CoopPointUpdateNotify) },
		func() ProtoMessage { return new(StartCoopPointReq) },
		func() ProtoMessage { return new(StartCoopPointRsp) },
		func() ProtoMessage { return new(CancelCoopTaskReq) },
		func() ProtoMessage { return new(CancelCoopTaskRsp) },
		func() ProtoMessage { return new(TakeCoopRewardReq) },
		func() ProtoMessage { return new(TakeCoopRewardRsp) },
		func() ProtoMessage { return new(CoopProgressUpdateNotify) },
		func() ProtoMessage { return new(SaveCoopDialogReq) },
		func() ProtoMessage { return new(SaveCoopDialogRsp) },
		func() ProtoMessage { return new(CoopCgShowNotify) },
		func() ProtoMessage { return new(SetCoopChapterViewedReq) },
		func() ProtoMessage { return new(SetCoopChapterViewedRsp) },
		func() ProtoMessage { return new(MainCoopFailNotify) },
	)
}

const (
	ProtoCommandAllCoopInfoNotify        ProtoCommand = 1976
	ProtoCommandMainCoopUpdateNotify     ProtoCommand = 1968
	ProtoCommandSaveMainCoopReq          ProtoCommand = 1975
	ProtoCommandSaveMainCoopRsp          ProtoCommand = 1957
	ProtoCommandFinishMainCoopReq        ProtoCommand = 1952
	ProtoCommandFinishMainCoopRsp        ProtoCommand = 1981
	ProtoCommandCoopDataNotify           ProtoCommand = 1979
	ProtoCommandCoopChapterUpdateNotify  ProtoCommand = 1972
	ProtoCommandCoopCgUpdateNotify       ProtoCommand = 1994
	ProtoCommandCoopRewardUpdateNotify   ProtoCommand = 1999
	ProtoCommandUnlockCoopChapterReq     ProtoCommand = 1970
	ProtoCommandUnlockCoopChapterRsp     ProtoCommand = 1995
	ProtoCommandCoopPointUpdateNotify    ProtoCommand = 1991
	ProtoCommandStartCoopPointReq        ProtoCommand = 1992
	ProtoCommandStartCoopPointRsp        ProtoCommand = 1964
	ProtoCommandCancelCoopTaskReq        ProtoCommand = 1997
	ProtoCommandCancelCoopTaskRsp        ProtoCommand = 1987
	ProtoCommandTakeCoopRewardReq        ProtoCommand = 1973
	ProtoCommandTakeCoopRewardRsp        ProtoCommand = 1985
	ProtoCommandCoopProgressUpdateNotify ProtoCommand = 1998
	ProtoCommandSaveCoopDialogReq        ProtoCommand = 2000
	ProtoCommandSaveCoopDialogRsp        ProtoCommand = 1962
	ProtoCommandCoopCgShowNotify         ProtoCommand = 1983
	ProtoCommandSetCoopChapterViewedReq  ProtoCommand = 1965
	ProtoCommandSetCoopChapterViewedRsp  ProtoCommand = 1963
	ProtoCommandMainCoopFailNotify       ProtoCommand = 1951
)

func (*AllCoopInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAllCoopInfoNotify }
func (*AllCoopInfoNotify) ProtoMessageType() ProtoMessageType { return "AllCoopInfoNotify" }

func (*MainCoopUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMainCoopUpdateNotify }
func (*MainCoopUpdateNotify) ProtoMessageType() ProtoMessageType { return "MainCoopUpdateNotify" }

func (*SaveMainCoopReq) ProtoCommand() ProtoCommand         { return ProtoCommandSaveMainCoopReq }
func (*SaveMainCoopReq) ProtoMessageType() ProtoMessageType { return "SaveMainCoopReq" }

func (*SaveMainCoopRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSaveMainCoopRsp }
func (*SaveMainCoopRsp) ProtoMessageType() ProtoMessageType { return "SaveMainCoopRsp" }

func (*FinishMainCoopReq) ProtoCommand() ProtoCommand         { return ProtoCommandFinishMainCoopReq }
func (*FinishMainCoopReq) ProtoMessageType() ProtoMessageType { return "FinishMainCoopReq" }

func (*FinishMainCoopRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFinishMainCoopRsp }
func (*FinishMainCoopRsp) ProtoMessageType() ProtoMessageType { return "FinishMainCoopRsp" }

func (*CoopDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCoopDataNotify }
func (*CoopDataNotify) ProtoMessageType() ProtoMessageType { return "CoopDataNotify" }

func (*CoopChapterUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCoopChapterUpdateNotify
}
func (*CoopChapterUpdateNotify) ProtoMessageType() ProtoMessageType { return "CoopChapterUpdateNotify" }

func (*CoopCgUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCoopCgUpdateNotify }
func (*CoopCgUpdateNotify) ProtoMessageType() ProtoMessageType { return "CoopCgUpdateNotify" }

func (*CoopRewardUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCoopRewardUpdateNotify }
func (*CoopRewardUpdateNotify) ProtoMessageType() ProtoMessageType { return "CoopRewardUpdateNotify" }

func (*UnlockCoopChapterReq) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockCoopChapterReq }
func (*UnlockCoopChapterReq) ProtoMessageType() ProtoMessageType { return "UnlockCoopChapterReq" }

func (*UnlockCoopChapterRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockCoopChapterRsp }
func (*UnlockCoopChapterRsp) ProtoMessageType() ProtoMessageType { return "UnlockCoopChapterRsp" }

func (*CoopPointUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCoopPointUpdateNotify }
func (*CoopPointUpdateNotify) ProtoMessageType() ProtoMessageType { return "CoopPointUpdateNotify" }

func (*StartCoopPointReq) ProtoCommand() ProtoCommand         { return ProtoCommandStartCoopPointReq }
func (*StartCoopPointReq) ProtoMessageType() ProtoMessageType { return "StartCoopPointReq" }

func (*StartCoopPointRsp) ProtoCommand() ProtoCommand         { return ProtoCommandStartCoopPointRsp }
func (*StartCoopPointRsp) ProtoMessageType() ProtoMessageType { return "StartCoopPointRsp" }

func (*CancelCoopTaskReq) ProtoCommand() ProtoCommand         { return ProtoCommandCancelCoopTaskReq }
func (*CancelCoopTaskReq) ProtoMessageType() ProtoMessageType { return "CancelCoopTaskReq" }

func (*CancelCoopTaskRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCancelCoopTaskRsp }
func (*CancelCoopTaskRsp) ProtoMessageType() ProtoMessageType { return "CancelCoopTaskRsp" }

func (*TakeCoopRewardReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeCoopRewardReq }
func (*TakeCoopRewardReq) ProtoMessageType() ProtoMessageType { return "TakeCoopRewardReq" }

func (*TakeCoopRewardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeCoopRewardRsp }
func (*TakeCoopRewardRsp) ProtoMessageType() ProtoMessageType { return "TakeCoopRewardRsp" }

func (*CoopProgressUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCoopProgressUpdateNotify
}
func (*CoopProgressUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "CoopProgressUpdateNotify"
}

func (*SaveCoopDialogReq) ProtoCommand() ProtoCommand         { return ProtoCommandSaveCoopDialogReq }
func (*SaveCoopDialogReq) ProtoMessageType() ProtoMessageType { return "SaveCoopDialogReq" }

func (*SaveCoopDialogRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSaveCoopDialogRsp }
func (*SaveCoopDialogRsp) ProtoMessageType() ProtoMessageType { return "SaveCoopDialogRsp" }

func (*CoopCgShowNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCoopCgShowNotify }
func (*CoopCgShowNotify) ProtoMessageType() ProtoMessageType { return "CoopCgShowNotify" }

func (*SetCoopChapterViewedReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetCoopChapterViewedReq
}
func (*SetCoopChapterViewedReq) ProtoMessageType() ProtoMessageType { return "SetCoopChapterViewedReq" }

func (*SetCoopChapterViewedRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetCoopChapterViewedRsp
}
func (*SetCoopChapterViewedRsp) ProtoMessageType() ProtoMessageType { return "SetCoopChapterViewedRsp" }

func (*MainCoopFailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMainCoopFailNotify }
func (*MainCoopFailNotify) ProtoMessageType() ProtoMessageType { return "MainCoopFailNotify" }
