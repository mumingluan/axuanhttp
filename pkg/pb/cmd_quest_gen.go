
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(QuestListNotify) },
		func() ProtoMessage { return new(QuestListUpdateNotify) },
		func() ProtoMessage { return new(QuestDelNotify) },
		func() ProtoMessage { return new(FinishedParentQuestNotify) },
		func() ProtoMessage { return new(FinishedParentQuestUpdateNotify) },
		func() ProtoMessage { return new(AddQuestContentProgressReq) },
		func() ProtoMessage { return new(AddQuestContentProgressRsp) },
		func() ProtoMessage { return new(GetQuestTalkHistoryReq) },
		func() ProtoMessage { return new(GetQuestTalkHistoryRsp) },
		func() ProtoMessage { return new(QuestCreateEntityReq) },
		func() ProtoMessage { return new(QuestCreateEntityRsp) },
		func() ProtoMessage { return new(QuestDestroyEntityReq) },
		func() ProtoMessage { return new(QuestDestroyEntityRsp) },
		func() ProtoMessage { return new(ChapterStateNotify) },
		func() ProtoMessage { return new(QuestProgressUpdateNotify) },
		func() ProtoMessage { return new(QuestUpdateQuestVarReq) },
		func() ProtoMessage { return new(QuestUpdateQuestVarRsp) },
		func() ProtoMessage { return new(QuestUpdateQuestVarNotify) },
		func() ProtoMessage { return new(QuestDestroyNpcReq) },
		func() ProtoMessage { return new(QuestDestroyNpcRsp) },
		func() ProtoMessage { return new(BargainStartNotify) },
		func() ProtoMessage { return new(BargainOfferPriceReq) },
		func() ProtoMessage { return new(BargainOfferPriceRsp) },
		func() ProtoMessage { return new(BargainTerminateNotify) },
		func() ProtoMessage { return new(GetBargainDataReq) },
		func() ProtoMessage { return new(GetBargainDataRsp) },
		func() ProtoMessage { return new(GetAllActivatedBargainDataReq) },
		func() ProtoMessage { return new(GetAllActivatedBargainDataRsp) },
		func() ProtoMessage { return new(ServerCondMeetQuestListUpdateNotify) },
		func() ProtoMessage { return new(QuestGlobalVarNotify) },
		func() ProtoMessage { return new(QuestTransmitReq) },
		func() ProtoMessage { return new(QuestTransmitRsp) },
		func() ProtoMessage { return new(PersonalLineAllDataReq) },
		func() ProtoMessage { return new(PersonalLineAllDataRsp) },
		func() ProtoMessage { return new(RedeemLegendaryKeyReq) },
		func() ProtoMessage { return new(RedeemLegendaryKeyRsp) },
		func() ProtoMessage { return new(UnlockPersonalLineReq) },
		func() ProtoMessage { return new(UnlockPersonalLineRsp) },
		func() ProtoMessage { return new(CancelFinishParentQuestNotify) },
		func() ProtoMessage { return new(QuestUpdateQuestTimeVarNotify) },
		func() ProtoMessage { return new(PersonalLineNewUnlockNotify) },
		func() ProtoMessage { return new(NpcTalkStateNotify) },
		func() ProtoMessage { return new(GetQuestLackingResourceReq) },
		func() ProtoMessage { return new(GetQuestLackingResourceRsp) },
		func() ProtoMessage { return new(GetParentQuestVideoKeyReq) },
		func() ProtoMessage { return new(GetParentQuestVideoKeyRsp) },
		func() ProtoMessage { return new(ParentQuestInferenceDataNotify) },
		func() ProtoMessage { return new(InterpretInferenceWordReq) },
		func() ProtoMessage { return new(InterpretInferenceWordRsp) },
		func() ProtoMessage { return new(AssociateInferenceWordReq) },
		func() ProtoMessage { return new(AssociateInferenceWordRsp) },
		func() ProtoMessage { return new(SubmitInferenceWordReq) },
		func() ProtoMessage { return new(SubmitInferenceWordRsp) },
		func() ProtoMessage { return new(QuestRenameAvatarReq) },
		func() ProtoMessage { return new(QuestRenameAvatarRsp) },
	)
}

const (
	ProtoCommandQuestListNotify                     ProtoCommand = 472
	ProtoCommandQuestListUpdateNotify               ProtoCommand = 498
	ProtoCommandQuestDelNotify                      ProtoCommand = 412
	ProtoCommandFinishedParentQuestNotify           ProtoCommand = 435
	ProtoCommandFinishedParentQuestUpdateNotify     ProtoCommand = 407
	ProtoCommandAddQuestContentProgressReq          ProtoCommand = 421
	ProtoCommandAddQuestContentProgressRsp          ProtoCommand = 403
	ProtoCommandGetQuestTalkHistoryReq              ProtoCommand = 490
	ProtoCommandGetQuestTalkHistoryRsp              ProtoCommand = 473
	ProtoCommandQuestCreateEntityReq                ProtoCommand = 499
	ProtoCommandQuestCreateEntityRsp                ProtoCommand = 431
	ProtoCommandQuestDestroyEntityReq               ProtoCommand = 475
	ProtoCommandQuestDestroyEntityRsp               ProtoCommand = 448
	ProtoCommandChapterStateNotify                  ProtoCommand = 405
	ProtoCommandQuestProgressUpdateNotify           ProtoCommand = 482
	ProtoCommandQuestUpdateQuestVarReq              ProtoCommand = 447
	ProtoCommandQuestUpdateQuestVarRsp              ProtoCommand = 439
	ProtoCommandQuestUpdateQuestVarNotify           ProtoCommand = 453
	ProtoCommandQuestDestroyNpcReq                  ProtoCommand = 422
	ProtoCommandQuestDestroyNpcRsp                  ProtoCommand = 465
	ProtoCommandBargainStartNotify                  ProtoCommand = 404
	ProtoCommandBargainOfferPriceReq                ProtoCommand = 493
	ProtoCommandBargainOfferPriceRsp                ProtoCommand = 427
	ProtoCommandBargainTerminateNotify              ProtoCommand = 494
	ProtoCommandGetBargainDataReq                   ProtoCommand = 488
	ProtoCommandGetBargainDataRsp                   ProtoCommand = 426
	ProtoCommandGetAllActivatedBargainDataReq       ProtoCommand = 463
	ProtoCommandGetAllActivatedBargainDataRsp       ProtoCommand = 495
	ProtoCommandServerCondMeetQuestListUpdateNotify ProtoCommand = 406
	ProtoCommandQuestGlobalVarNotify                ProtoCommand = 434
	ProtoCommandQuestTransmitReq                    ProtoCommand = 450
	ProtoCommandQuestTransmitRsp                    ProtoCommand = 443
	ProtoCommandPersonalLineAllDataReq              ProtoCommand = 474
	ProtoCommandPersonalLineAllDataRsp              ProtoCommand = 476
	ProtoCommandRedeemLegendaryKeyReq               ProtoCommand = 446
	ProtoCommandRedeemLegendaryKeyRsp               ProtoCommand = 441
	ProtoCommandUnlockPersonalLineReq               ProtoCommand = 449
	ProtoCommandUnlockPersonalLineRsp               ProtoCommand = 491
	ProtoCommandCancelFinishParentQuestNotify       ProtoCommand = 424
	ProtoCommandQuestUpdateQuestTimeVarNotify       ProtoCommand = 456
	ProtoCommandPersonalLineNewUnlockNotify         ProtoCommand = 442
	ProtoCommandNpcTalkStateNotify                  ProtoCommand = 430
	ProtoCommandGetQuestLackingResourceReq          ProtoCommand = 467
	ProtoCommandGetQuestLackingResourceRsp          ProtoCommand = 458
	ProtoCommandGetParentQuestVideoKeyReq           ProtoCommand = 470
	ProtoCommandGetParentQuestVideoKeyRsp           ProtoCommand = 417
	ProtoCommandParentQuestInferenceDataNotify      ProtoCommand = 402
	ProtoCommandInterpretInferenceWordReq           ProtoCommand = 419
	ProtoCommandInterpretInferenceWordRsp           ProtoCommand = 461
	ProtoCommandAssociateInferenceWordReq           ProtoCommand = 429
	ProtoCommandAssociateInferenceWordRsp           ProtoCommand = 457
	ProtoCommandSubmitInferenceWordReq              ProtoCommand = 500
	ProtoCommandSubmitInferenceWordRsp              ProtoCommand = 416
	ProtoCommandQuestRenameAvatarReq                ProtoCommand = 487
	ProtoCommandQuestRenameAvatarRsp                ProtoCommand = 440
)

func (*QuestListNotify) ProtoCommand() ProtoCommand         { return ProtoCommandQuestListNotify }
func (*QuestListNotify) ProtoMessageType() ProtoMessageType { return "QuestListNotify" }

func (*QuestListUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandQuestListUpdateNotify }
func (*QuestListUpdateNotify) ProtoMessageType() ProtoMessageType { return "QuestListUpdateNotify" }

func (*QuestDelNotify) ProtoCommand() ProtoCommand         { return ProtoCommandQuestDelNotify }
func (*QuestDelNotify) ProtoMessageType() ProtoMessageType { return "QuestDelNotify" }

func (*FinishedParentQuestNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFinishedParentQuestNotify
}
func (*FinishedParentQuestNotify) ProtoMessageType() ProtoMessageType {
	return "FinishedParentQuestNotify"
}

func (*FinishedParentQuestUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFinishedParentQuestUpdateNotify
}
func (*FinishedParentQuestUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "FinishedParentQuestUpdateNotify"
}

func (*AddQuestContentProgressReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAddQuestContentProgressReq
}
func (*AddQuestContentProgressReq) ProtoMessageType() ProtoMessageType {
	return "AddQuestContentProgressReq"
}

func (*AddQuestContentProgressRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAddQuestContentProgressRsp
}
func (*AddQuestContentProgressRsp) ProtoMessageType() ProtoMessageType {
	return "AddQuestContentProgressRsp"
}

func (*GetQuestTalkHistoryReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetQuestTalkHistoryReq }
func (*GetQuestTalkHistoryReq) ProtoMessageType() ProtoMessageType { return "GetQuestTalkHistoryReq" }

func (*GetQuestTalkHistoryRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetQuestTalkHistoryRsp }
func (*GetQuestTalkHistoryRsp) ProtoMessageType() ProtoMessageType { return "GetQuestTalkHistoryRsp" }

func (*QuestCreateEntityReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuestCreateEntityReq }
func (*QuestCreateEntityReq) ProtoMessageType() ProtoMessageType { return "QuestCreateEntityReq" }

func (*QuestCreateEntityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuestCreateEntityRsp }
func (*QuestCreateEntityRsp) ProtoMessageType() ProtoMessageType { return "QuestCreateEntityRsp" }

func (*QuestDestroyEntityReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuestDestroyEntityReq }
func (*QuestDestroyEntityReq) ProtoMessageType() ProtoMessageType { return "QuestDestroyEntityReq" }

func (*QuestDestroyEntityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuestDestroyEntityRsp }
func (*QuestDestroyEntityRsp) ProtoMessageType() ProtoMessageType { return "QuestDestroyEntityRsp" }

func (*ChapterStateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChapterStateNotify }
func (*ChapterStateNotify) ProtoMessageType() ProtoMessageType { return "ChapterStateNotify" }

func (*QuestProgressUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandQuestProgressUpdateNotify
}
func (*QuestProgressUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "QuestProgressUpdateNotify"
}

func (*QuestUpdateQuestVarReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuestUpdateQuestVarReq }
func (*QuestUpdateQuestVarReq) ProtoMessageType() ProtoMessageType { return "QuestUpdateQuestVarReq" }

func (*QuestUpdateQuestVarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuestUpdateQuestVarRsp }
func (*QuestUpdateQuestVarRsp) ProtoMessageType() ProtoMessageType { return "QuestUpdateQuestVarRsp" }

func (*QuestUpdateQuestVarNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandQuestUpdateQuestVarNotify
}
func (*QuestUpdateQuestVarNotify) ProtoMessageType() ProtoMessageType {
	return "QuestUpdateQuestVarNotify"
}

func (*QuestDestroyNpcReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuestDestroyNpcReq }
func (*QuestDestroyNpcReq) ProtoMessageType() ProtoMessageType { return "QuestDestroyNpcReq" }

func (*QuestDestroyNpcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuestDestroyNpcRsp }
func (*QuestDestroyNpcRsp) ProtoMessageType() ProtoMessageType { return "QuestDestroyNpcRsp" }

func (*BargainStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandBargainStartNotify }
func (*BargainStartNotify) ProtoMessageType() ProtoMessageType { return "BargainStartNotify" }

func (*BargainOfferPriceReq) ProtoCommand() ProtoCommand         { return ProtoCommandBargainOfferPriceReq }
func (*BargainOfferPriceReq) ProtoMessageType() ProtoMessageType { return "BargainOfferPriceReq" }

func (*BargainOfferPriceRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBargainOfferPriceRsp }
func (*BargainOfferPriceRsp) ProtoMessageType() ProtoMessageType { return "BargainOfferPriceRsp" }

func (*BargainTerminateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandBargainTerminateNotify }
func (*BargainTerminateNotify) ProtoMessageType() ProtoMessageType { return "BargainTerminateNotify" }

func (*GetBargainDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetBargainDataReq }
func (*GetBargainDataReq) ProtoMessageType() ProtoMessageType { return "GetBargainDataReq" }

func (*GetBargainDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetBargainDataRsp }
func (*GetBargainDataRsp) ProtoMessageType() ProtoMessageType { return "GetBargainDataRsp" }

func (*GetAllActivatedBargainDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllActivatedBargainDataReq
}
func (*GetAllActivatedBargainDataReq) ProtoMessageType() ProtoMessageType {
	return "GetAllActivatedBargainDataReq"
}

func (*GetAllActivatedBargainDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllActivatedBargainDataRsp
}
func (*GetAllActivatedBargainDataRsp) ProtoMessageType() ProtoMessageType {
	return "GetAllActivatedBargainDataRsp"
}

func (*ServerCondMeetQuestListUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCondMeetQuestListUpdateNotify
}
func (*ServerCondMeetQuestListUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "ServerCondMeetQuestListUpdateNotify"
}

func (*QuestGlobalVarNotify) ProtoCommand() ProtoCommand         { return ProtoCommandQuestGlobalVarNotify }
func (*QuestGlobalVarNotify) ProtoMessageType() ProtoMessageType { return "QuestGlobalVarNotify" }

func (*QuestTransmitReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuestTransmitReq }
func (*QuestTransmitReq) ProtoMessageType() ProtoMessageType { return "QuestTransmitReq" }

func (*QuestTransmitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuestTransmitRsp }
func (*QuestTransmitRsp) ProtoMessageType() ProtoMessageType { return "QuestTransmitRsp" }

func (*PersonalLineAllDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandPersonalLineAllDataReq }
func (*PersonalLineAllDataReq) ProtoMessageType() ProtoMessageType { return "PersonalLineAllDataReq" }

func (*PersonalLineAllDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPersonalLineAllDataRsp }
func (*PersonalLineAllDataRsp) ProtoMessageType() ProtoMessageType { return "PersonalLineAllDataRsp" }

func (*RedeemLegendaryKeyReq) ProtoCommand() ProtoCommand         { return ProtoCommandRedeemLegendaryKeyReq }
func (*RedeemLegendaryKeyReq) ProtoMessageType() ProtoMessageType { return "RedeemLegendaryKeyReq" }

func (*RedeemLegendaryKeyRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRedeemLegendaryKeyRsp }
func (*RedeemLegendaryKeyRsp) ProtoMessageType() ProtoMessageType { return "RedeemLegendaryKeyRsp" }

func (*UnlockPersonalLineReq) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockPersonalLineReq }
func (*UnlockPersonalLineReq) ProtoMessageType() ProtoMessageType { return "UnlockPersonalLineReq" }

func (*UnlockPersonalLineRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockPersonalLineRsp }
func (*UnlockPersonalLineRsp) ProtoMessageType() ProtoMessageType { return "UnlockPersonalLineRsp" }

func (*CancelFinishParentQuestNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCancelFinishParentQuestNotify
}
func (*CancelFinishParentQuestNotify) ProtoMessageType() ProtoMessageType {
	return "CancelFinishParentQuestNotify"
}

func (*QuestUpdateQuestTimeVarNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandQuestUpdateQuestTimeVarNotify
}
func (*QuestUpdateQuestTimeVarNotify) ProtoMessageType() ProtoMessageType {
	return "QuestUpdateQuestTimeVarNotify"
}

func (*PersonalLineNewUnlockNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPersonalLineNewUnlockNotify
}
func (*PersonalLineNewUnlockNotify) ProtoMessageType() ProtoMessageType {
	return "PersonalLineNewUnlockNotify"
}

func (*NpcTalkStateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandNpcTalkStateNotify }
func (*NpcTalkStateNotify) ProtoMessageType() ProtoMessageType { return "NpcTalkStateNotify" }

func (*GetQuestLackingResourceReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetQuestLackingResourceReq
}
func (*GetQuestLackingResourceReq) ProtoMessageType() ProtoMessageType {
	return "GetQuestLackingResourceReq"
}

func (*GetQuestLackingResourceRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetQuestLackingResourceRsp
}
func (*GetQuestLackingResourceRsp) ProtoMessageType() ProtoMessageType {
	return "GetQuestLackingResourceRsp"
}

func (*GetParentQuestVideoKeyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetParentQuestVideoKeyReq
}
func (*GetParentQuestVideoKeyReq) ProtoMessageType() ProtoMessageType {
	return "GetParentQuestVideoKeyReq"
}

func (*GetParentQuestVideoKeyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetParentQuestVideoKeyRsp
}
func (*GetParentQuestVideoKeyRsp) ProtoMessageType() ProtoMessageType {
	return "GetParentQuestVideoKeyRsp"
}

func (*ParentQuestInferenceDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandParentQuestInferenceDataNotify
}
func (*ParentQuestInferenceDataNotify) ProtoMessageType() ProtoMessageType {
	return "ParentQuestInferenceDataNotify"
}

func (*InterpretInferenceWordReq) ProtoCommand() ProtoCommand {
	return ProtoCommandInterpretInferenceWordReq
}
func (*InterpretInferenceWordReq) ProtoMessageType() ProtoMessageType {
	return "InterpretInferenceWordReq"
}

func (*InterpretInferenceWordRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandInterpretInferenceWordRsp
}
func (*InterpretInferenceWordRsp) ProtoMessageType() ProtoMessageType {
	return "InterpretInferenceWordRsp"
}

func (*AssociateInferenceWordReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAssociateInferenceWordReq
}
func (*AssociateInferenceWordReq) ProtoMessageType() ProtoMessageType {
	return "AssociateInferenceWordReq"
}

func (*AssociateInferenceWordRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAssociateInferenceWordRsp
}
func (*AssociateInferenceWordRsp) ProtoMessageType() ProtoMessageType {
	return "AssociateInferenceWordRsp"
}

func (*SubmitInferenceWordReq) ProtoCommand() ProtoCommand         { return ProtoCommandSubmitInferenceWordReq }
func (*SubmitInferenceWordReq) ProtoMessageType() ProtoMessageType { return "SubmitInferenceWordReq" }

func (*SubmitInferenceWordRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSubmitInferenceWordRsp }
func (*SubmitInferenceWordRsp) ProtoMessageType() ProtoMessageType { return "SubmitInferenceWordRsp" }

func (*QuestRenameAvatarReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuestRenameAvatarReq }
func (*QuestRenameAvatarReq) ProtoMessageType() ProtoMessageType { return "QuestRenameAvatarReq" }

func (*QuestRenameAvatarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuestRenameAvatarRsp }
func (*QuestRenameAvatarRsp) ProtoMessageType() ProtoMessageType { return "QuestRenameAvatarRsp" }
