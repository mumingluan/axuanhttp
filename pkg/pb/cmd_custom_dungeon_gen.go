
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(EnterCustomDungeonReq) },
		func() ProtoMessage { return new(EnterCustomDungeonRsp) },
		func() ProtoMessage { return new(SaveCustomDungeonRoomReq) },
		func() ProtoMessage { return new(SaveCustomDungeonRoomRsp) },
		func() ProtoMessage { return new(ChangeCustomDungeonRoomReq) },
		func() ProtoMessage { return new(ChangeCustomDungeonRoomRsp) },
		func() ProtoMessage { return new(RemoveCustomDungeonReq) },
		func() ProtoMessage { return new(RemoveCustomDungeonRsp) },
		func() ProtoMessage { return new(TryCustomDungeonReq) },
		func() ProtoMessage { return new(TryCustomDungeonRsp) },
		func() ProtoMessage { return new(PublishCustomDungeonReq) },
		func() ProtoMessage { return new(PublishCustomDungeonRsp) },
		func() ProtoMessage { return new(ExitCustomDungeonTryReq) },
		func() ProtoMessage { return new(ExitCustomDungeonTryRsp) },
		func() ProtoMessage { return new(CustomDungeonUpdateNotify) },
		func() ProtoMessage { return new(GetRecommendCustomDungeonReq) },
		func() ProtoMessage { return new(GetRecommendCustomDungeonRsp) },
		func() ProtoMessage { return new(GetStoreCustomDungeonReq) },
		func() ProtoMessage { return new(GetStoreCustomDungeonRsp) },
		func() ProtoMessage { return new(SearchCustomDungeonReq) },
		func() ProtoMessage { return new(SearchCustomDungeonRsp) },
		func() ProtoMessage { return new(StoreCustomDungeonReq) },
		func() ProtoMessage { return new(StoreCustomDungeonRsp) },
		func() ProtoMessage { return new(LikeCustomDungeonReq) },
		func() ProtoMessage { return new(LikeCustomDungeonRsp) },
		func() ProtoMessage { return new(GetCustomDungeonReq) },
		func() ProtoMessage { return new(GetCustomDungeonRsp) },
		func() ProtoMessage { return new(CustomDungeonRecoverNotify) },
		func() ProtoMessage { return new(BackPlayCustomDungeonOfficialReq) },
		func() ProtoMessage { return new(BackPlayCustomDungeonOfficialRsp) },
		func() ProtoMessage { return new(CustomDungeonOfficialNotify) },
		func() ProtoMessage { return new(ReplayCustomDungeonReq) },
		func() ProtoMessage { return new(ReplayCustomDungeonRsp) },
		func() ProtoMessage { return new(CustomDungeonBattleRecordNotify) },
		func() ProtoMessage { return new(OutStuckCustomDungeonReq) },
		func() ProtoMessage { return new(OutStuckCustomDungeonRsp) },
	)
}

const (
	ProtoCommandEnterCustomDungeonReq            ProtoCommand = 6226
	ProtoCommandEnterCustomDungeonRsp            ProtoCommand = 6218
	ProtoCommandSaveCustomDungeonRoomReq         ProtoCommand = 6225
	ProtoCommandSaveCustomDungeonRoomRsp         ProtoCommand = 6207
	ProtoCommandChangeCustomDungeonRoomReq       ProtoCommand = 6222
	ProtoCommandChangeCustomDungeonRoomRsp       ProtoCommand = 6244
	ProtoCommandRemoveCustomDungeonReq           ProtoCommand = 6249
	ProtoCommandRemoveCustomDungeonRsp           ProtoCommand = 6220
	ProtoCommandTryCustomDungeonReq              ProtoCommand = 6245
	ProtoCommandTryCustomDungeonRsp              ProtoCommand = 6241
	ProtoCommandPublishCustomDungeonReq          ProtoCommand = 6242
	ProtoCommandPublishCustomDungeonRsp          ProtoCommand = 6214
	ProtoCommandExitCustomDungeonTryReq          ProtoCommand = 6247
	ProtoCommandExitCustomDungeonTryRsp          ProtoCommand = 6237
	ProtoCommandCustomDungeonUpdateNotify        ProtoCommand = 6223
	ProtoCommandGetRecommendCustomDungeonReq     ProtoCommand = 6235
	ProtoCommandGetRecommendCustomDungeonRsp     ProtoCommand = 6248
	ProtoCommandGetStoreCustomDungeonReq         ProtoCommand = 6250
	ProtoCommandGetStoreCustomDungeonRsp         ProtoCommand = 6212
	ProtoCommandSearchCustomDungeonReq           ProtoCommand = 6233
	ProtoCommandSearchCustomDungeonRsp           ProtoCommand = 6215
	ProtoCommandStoreCustomDungeonReq            ProtoCommand = 6213
	ProtoCommandStoreCustomDungeonRsp            ProtoCommand = 6201
	ProtoCommandLikeCustomDungeonReq             ProtoCommand = 6210
	ProtoCommandLikeCustomDungeonRsp             ProtoCommand = 6219
	ProtoCommandGetCustomDungeonReq              ProtoCommand = 6209
	ProtoCommandGetCustomDungeonRsp              ProtoCommand = 6227
	ProtoCommandCustomDungeonRecoverNotify       ProtoCommand = 6217
	ProtoCommandBackPlayCustomDungeonOfficialReq ProtoCommand = 6203
	ProtoCommandBackPlayCustomDungeonOfficialRsp ProtoCommand = 6204
	ProtoCommandCustomDungeonOfficialNotify      ProtoCommand = 6221
	ProtoCommandReplayCustomDungeonReq           ProtoCommand = 6243
	ProtoCommandReplayCustomDungeonRsp           ProtoCommand = 6240
	ProtoCommandCustomDungeonBattleRecordNotify  ProtoCommand = 6236
	ProtoCommandOutStuckCustomDungeonReq         ProtoCommand = 6211
	ProtoCommandOutStuckCustomDungeonRsp         ProtoCommand = 6234
)

func (*EnterCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandEnterCustomDungeonReq }
func (*EnterCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "EnterCustomDungeonReq" }

func (*EnterCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEnterCustomDungeonRsp }
func (*EnterCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "EnterCustomDungeonRsp" }

func (*SaveCustomDungeonRoomReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSaveCustomDungeonRoomReq
}
func (*SaveCustomDungeonRoomReq) ProtoMessageType() ProtoMessageType {
	return "SaveCustomDungeonRoomReq"
}

func (*SaveCustomDungeonRoomRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSaveCustomDungeonRoomRsp
}
func (*SaveCustomDungeonRoomRsp) ProtoMessageType() ProtoMessageType {
	return "SaveCustomDungeonRoomRsp"
}

func (*ChangeCustomDungeonRoomReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeCustomDungeonRoomReq
}
func (*ChangeCustomDungeonRoomReq) ProtoMessageType() ProtoMessageType {
	return "ChangeCustomDungeonRoomReq"
}

func (*ChangeCustomDungeonRoomRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeCustomDungeonRoomRsp
}
func (*ChangeCustomDungeonRoomRsp) ProtoMessageType() ProtoMessageType {
	return "ChangeCustomDungeonRoomRsp"
}

func (*RemoveCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveCustomDungeonReq }
func (*RemoveCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "RemoveCustomDungeonReq" }

func (*RemoveCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveCustomDungeonRsp }
func (*RemoveCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "RemoveCustomDungeonRsp" }

func (*TryCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandTryCustomDungeonReq }
func (*TryCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "TryCustomDungeonReq" }

func (*TryCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTryCustomDungeonRsp }
func (*TryCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "TryCustomDungeonRsp" }

func (*PublishCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPublishCustomDungeonReq
}
func (*PublishCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "PublishCustomDungeonReq" }

func (*PublishCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPublishCustomDungeonRsp
}
func (*PublishCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "PublishCustomDungeonRsp" }

func (*ExitCustomDungeonTryReq) ProtoCommand() ProtoCommand {
	return ProtoCommandExitCustomDungeonTryReq
}
func (*ExitCustomDungeonTryReq) ProtoMessageType() ProtoMessageType { return "ExitCustomDungeonTryReq" }

func (*ExitCustomDungeonTryRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandExitCustomDungeonTryRsp
}
func (*ExitCustomDungeonTryRsp) ProtoMessageType() ProtoMessageType { return "ExitCustomDungeonTryRsp" }

func (*CustomDungeonUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCustomDungeonUpdateNotify
}
func (*CustomDungeonUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "CustomDungeonUpdateNotify"
}

func (*GetRecommendCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRecommendCustomDungeonReq
}
func (*GetRecommendCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "GetRecommendCustomDungeonReq"
}

func (*GetRecommendCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRecommendCustomDungeonRsp
}
func (*GetRecommendCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "GetRecommendCustomDungeonRsp"
}

func (*GetStoreCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetStoreCustomDungeonReq
}
func (*GetStoreCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "GetStoreCustomDungeonReq"
}

func (*GetStoreCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetStoreCustomDungeonRsp
}
func (*GetStoreCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "GetStoreCustomDungeonRsp"
}

func (*SearchCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandSearchCustomDungeonReq }
func (*SearchCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "SearchCustomDungeonReq" }

func (*SearchCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSearchCustomDungeonRsp }
func (*SearchCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "SearchCustomDungeonRsp" }

func (*StoreCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandStoreCustomDungeonReq }
func (*StoreCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "StoreCustomDungeonReq" }

func (*StoreCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandStoreCustomDungeonRsp }
func (*StoreCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "StoreCustomDungeonRsp" }

func (*LikeCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandLikeCustomDungeonReq }
func (*LikeCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "LikeCustomDungeonReq" }

func (*LikeCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLikeCustomDungeonRsp }
func (*LikeCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "LikeCustomDungeonRsp" }

func (*GetCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetCustomDungeonReq }
func (*GetCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "GetCustomDungeonReq" }

func (*GetCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetCustomDungeonRsp }
func (*GetCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "GetCustomDungeonRsp" }

func (*CustomDungeonRecoverNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCustomDungeonRecoverNotify
}
func (*CustomDungeonRecoverNotify) ProtoMessageType() ProtoMessageType {
	return "CustomDungeonRecoverNotify"
}

func (*BackPlayCustomDungeonOfficialReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBackPlayCustomDungeonOfficialReq
}
func (*BackPlayCustomDungeonOfficialReq) ProtoMessageType() ProtoMessageType {
	return "BackPlayCustomDungeonOfficialReq"
}

func (*BackPlayCustomDungeonOfficialRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBackPlayCustomDungeonOfficialRsp
}
func (*BackPlayCustomDungeonOfficialRsp) ProtoMessageType() ProtoMessageType {
	return "BackPlayCustomDungeonOfficialRsp"
}

func (*CustomDungeonOfficialNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCustomDungeonOfficialNotify
}
func (*CustomDungeonOfficialNotify) ProtoMessageType() ProtoMessageType {
	return "CustomDungeonOfficialNotify"
}

func (*ReplayCustomDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandReplayCustomDungeonReq }
func (*ReplayCustomDungeonReq) ProtoMessageType() ProtoMessageType { return "ReplayCustomDungeonReq" }

func (*ReplayCustomDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReplayCustomDungeonRsp }
func (*ReplayCustomDungeonRsp) ProtoMessageType() ProtoMessageType { return "ReplayCustomDungeonRsp" }

func (*CustomDungeonBattleRecordNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCustomDungeonBattleRecordNotify
}
func (*CustomDungeonBattleRecordNotify) ProtoMessageType() ProtoMessageType {
	return "CustomDungeonBattleRecordNotify"
}

func (*OutStuckCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandOutStuckCustomDungeonReq
}
func (*OutStuckCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "OutStuckCustomDungeonReq"
}

func (*OutStuckCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandOutStuckCustomDungeonRsp
}
func (*OutStuckCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "OutStuckCustomDungeonRsp"
}
