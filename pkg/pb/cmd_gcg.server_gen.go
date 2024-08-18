
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(ServerGCGOperationReq) },
		func() ProtoMessage { return new(ServerGCGOperationRsp) },
		func() ProtoMessage { return new(ServerGCGAskDuelReq) },
		func() ProtoMessage { return new(ServerGCGNewGameInfoNotify) },
		func() ProtoMessage { return new(ServerGCGGMCommandReq) },
		func() ProtoMessage { return new(ServerGCGGMCommandRsp) },
		func() ProtoMessage { return new(ServerGCGInitFinishReq) },
		func() ProtoMessage { return new(ServerGCGSyncReq) },
		func() ProtoMessage { return new(ServerGCGSyncRsp) },
		func() ProtoMessage { return new(ServerGCGCreateSingleGameReq) },
		func() ProtoMessage { return new(ServerGCGCreateSingleGameRsp) },
		func() ProtoMessage { return new(ServerGCGGiveUpGameReq) },
		func() ProtoMessage { return new(ServerGCGGiveUpGameRsp) },
		func() ProtoMessage { return new(ServerGCGSkillPreviewAskReq) },
		func() ProtoMessage { return new(ServerGCGSkillPreviewAskRsp) },
		func() ProtoMessage { return new(ServerGCGCreateMultiGameReq) },
		func() ProtoMessage { return new(ServerGCGCreateMultiGameRsp) },
		func() ProtoMessage { return new(ServerGCGAddBothAITaskReq) },
		func() ProtoMessage { return new(ServerGCGAddBothAITaskRsp) },
		func() ProtoMessage { return new(ServerGCGCreateMatchGameNotify) },
		func() ProtoMessage { return new(ServerGCGCreateMatchAIGameNotify) },
		func() ProtoMessage { return new(ServerGCGMatchFinishNotify) },
		func() ProtoMessage { return new(ServerGCGBackToDuelReq) },
		func() ProtoMessage { return new(ServerGCGBackToDuelRsp) },
		func() ProtoMessage { return new(ServerGCGMessagePackNotify) },
		func() ProtoMessage { return new(ServerGCGHeartBeatNotify) },
		func() ProtoMessage { return new(ServerGCGSkillPreviewNotify) },
		func() ProtoMessage { return new(ServerGCGChallengeUpdateNotify) },
	)
}

const (
	ProtoCommandServerGCGOperationReq            ProtoCommand = 12436
	ProtoCommandServerGCGOperationRsp            ProtoCommand = 12407
	ProtoCommandServerGCGAskDuelReq              ProtoCommand = 12457
	ProtoCommandServerGCGNewGameInfoNotify       ProtoCommand = 12360
	ProtoCommandServerGCGGMCommandReq            ProtoCommand = 12373
	ProtoCommandServerGCGGMCommandRsp            ProtoCommand = 12338
	ProtoCommandServerGCGInitFinishReq           ProtoCommand = 12334
	ProtoCommandServerGCGSyncReq                 ProtoCommand = 12456
	ProtoCommandServerGCGSyncRsp                 ProtoCommand = 12328
	ProtoCommandServerGCGCreateSingleGameReq     ProtoCommand = 12489
	ProtoCommandServerGCGCreateSingleGameRsp     ProtoCommand = 12308
	ProtoCommandServerGCGGiveUpGameReq           ProtoCommand = 12499
	ProtoCommandServerGCGGiveUpGameRsp           ProtoCommand = 12492
	ProtoCommandServerGCGSkillPreviewAskReq      ProtoCommand = 12319
	ProtoCommandServerGCGSkillPreviewAskRsp      ProtoCommand = 12477
	ProtoCommandServerGCGCreateMultiGameReq      ProtoCommand = 12476
	ProtoCommandServerGCGCreateMultiGameRsp      ProtoCommand = 12490
	ProtoCommandServerGCGAddBothAITaskReq        ProtoCommand = 12423
	ProtoCommandServerGCGAddBothAITaskRsp        ProtoCommand = 12439
	ProtoCommandServerGCGCreateMatchGameNotify   ProtoCommand = 12497
	ProtoCommandServerGCGCreateMatchAIGameNotify ProtoCommand = 12392
	ProtoCommandServerGCGMatchFinishNotify       ProtoCommand = 12329
	ProtoCommandServerGCGBackToDuelReq           ProtoCommand = 12455
	ProtoCommandServerGCGBackToDuelRsp           ProtoCommand = 12309
	ProtoCommandServerGCGMessagePackNotify       ProtoCommand = 12463
	ProtoCommandServerGCGHeartBeatNotify         ProtoCommand = 12405
	ProtoCommandServerGCGSkillPreviewNotify      ProtoCommand = 12414
	ProtoCommandServerGCGChallengeUpdateNotify   ProtoCommand = 12332
)

func (*ServerGCGOperationReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGOperationReq }
func (*ServerGCGOperationReq) ProtoMessageType() ProtoMessageType { return "ServerGCGOperationReq" }

func (*ServerGCGOperationRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGOperationRsp }
func (*ServerGCGOperationRsp) ProtoMessageType() ProtoMessageType { return "ServerGCGOperationRsp" }

func (*ServerGCGAskDuelReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGAskDuelReq }
func (*ServerGCGAskDuelReq) ProtoMessageType() ProtoMessageType { return "ServerGCGAskDuelReq" }

func (*ServerGCGNewGameInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGNewGameInfoNotify
}
func (*ServerGCGNewGameInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGNewGameInfoNotify"
}

func (*ServerGCGGMCommandReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGGMCommandReq }
func (*ServerGCGGMCommandReq) ProtoMessageType() ProtoMessageType { return "ServerGCGGMCommandReq" }

func (*ServerGCGGMCommandRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGGMCommandRsp }
func (*ServerGCGGMCommandRsp) ProtoMessageType() ProtoMessageType { return "ServerGCGGMCommandRsp" }

func (*ServerGCGInitFinishReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGInitFinishReq }
func (*ServerGCGInitFinishReq) ProtoMessageType() ProtoMessageType { return "ServerGCGInitFinishReq" }

func (*ServerGCGSyncReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGSyncReq }
func (*ServerGCGSyncReq) ProtoMessageType() ProtoMessageType { return "ServerGCGSyncReq" }

func (*ServerGCGSyncRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGSyncRsp }
func (*ServerGCGSyncRsp) ProtoMessageType() ProtoMessageType { return "ServerGCGSyncRsp" }

func (*ServerGCGCreateSingleGameReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGCreateSingleGameReq
}
func (*ServerGCGCreateSingleGameReq) ProtoMessageType() ProtoMessageType {
	return "ServerGCGCreateSingleGameReq"
}

func (*ServerGCGCreateSingleGameRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGCreateSingleGameRsp
}
func (*ServerGCGCreateSingleGameRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGCGCreateSingleGameRsp"
}

func (*ServerGCGGiveUpGameReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGGiveUpGameReq }
func (*ServerGCGGiveUpGameReq) ProtoMessageType() ProtoMessageType { return "ServerGCGGiveUpGameReq" }

func (*ServerGCGGiveUpGameRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGGiveUpGameRsp }
func (*ServerGCGGiveUpGameRsp) ProtoMessageType() ProtoMessageType { return "ServerGCGGiveUpGameRsp" }

func (*ServerGCGSkillPreviewAskReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGSkillPreviewAskReq
}
func (*ServerGCGSkillPreviewAskReq) ProtoMessageType() ProtoMessageType {
	return "ServerGCGSkillPreviewAskReq"
}

func (*ServerGCGSkillPreviewAskRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGSkillPreviewAskRsp
}
func (*ServerGCGSkillPreviewAskRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGCGSkillPreviewAskRsp"
}

func (*ServerGCGCreateMultiGameReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGCreateMultiGameReq
}
func (*ServerGCGCreateMultiGameReq) ProtoMessageType() ProtoMessageType {
	return "ServerGCGCreateMultiGameReq"
}

func (*ServerGCGCreateMultiGameRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGCreateMultiGameRsp
}
func (*ServerGCGCreateMultiGameRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGCGCreateMultiGameRsp"
}

func (*ServerGCGAddBothAITaskReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGAddBothAITaskReq
}
func (*ServerGCGAddBothAITaskReq) ProtoMessageType() ProtoMessageType {
	return "ServerGCGAddBothAITaskReq"
}

func (*ServerGCGAddBothAITaskRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGAddBothAITaskRsp
}
func (*ServerGCGAddBothAITaskRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGCGAddBothAITaskRsp"
}

func (*ServerGCGCreateMatchGameNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGCreateMatchGameNotify
}
func (*ServerGCGCreateMatchGameNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGCreateMatchGameNotify"
}

func (*ServerGCGCreateMatchAIGameNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGCreateMatchAIGameNotify
}
func (*ServerGCGCreateMatchAIGameNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGCreateMatchAIGameNotify"
}

func (*ServerGCGMatchFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGMatchFinishNotify
}
func (*ServerGCGMatchFinishNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGMatchFinishNotify"
}

func (*ServerGCGBackToDuelReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGBackToDuelReq }
func (*ServerGCGBackToDuelReq) ProtoMessageType() ProtoMessageType { return "ServerGCGBackToDuelReq" }

func (*ServerGCGBackToDuelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGCGBackToDuelRsp }
func (*ServerGCGBackToDuelRsp) ProtoMessageType() ProtoMessageType { return "ServerGCGBackToDuelRsp" }

func (*ServerGCGMessagePackNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGMessagePackNotify
}
func (*ServerGCGMessagePackNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGMessagePackNotify"
}

func (*ServerGCGHeartBeatNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGHeartBeatNotify
}
func (*ServerGCGHeartBeatNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGHeartBeatNotify"
}

func (*ServerGCGSkillPreviewNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGSkillPreviewNotify
}
func (*ServerGCGSkillPreviewNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGSkillPreviewNotify"
}

func (*ServerGCGChallengeUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGCGChallengeUpdateNotify
}
func (*ServerGCGChallengeUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGCGChallengeUpdateNotify"
}
