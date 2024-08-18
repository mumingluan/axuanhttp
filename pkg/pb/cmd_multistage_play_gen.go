
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(MultistagePlayInfoNotify) },
		func() ProtoMessage { return new(MultistagePlayFinishStageReq) },
		func() ProtoMessage { return new(InBattleMechanicusExcapeMonsterNotify) },
		func() ProtoMessage { return new(InBattleMechanicusLeftMonsterNotify) },
		func() ProtoMessage { return new(InBattleMechanicusBuildingPointsNotify) },
		func() ProtoMessage { return new(InBattleMechanicusPickCardReq) },
		func() ProtoMessage { return new(InBattleMechanicusPickCardRsp) },
		func() ProtoMessage { return new(InBattleMechanicusPickCardNotify) },
		func() ProtoMessage { return new(InBattleMechanicusConfirmCardReq) },
		func() ProtoMessage { return new(InBattleMechanicusConfirmCardRsp) },
		func() ProtoMessage { return new(InBattleMechanicusConfirmCardNotify) },
		func() ProtoMessage { return new(InBattleMechanicusCardResultNotify) },
		func() ProtoMessage { return new(MultistagePlayFinishStageRsp) },
		func() ProtoMessage { return new(InBattleMechanicusSettleNotify) },
		func() ProtoMessage { return new(MultistagePlaySettleNotify) },
		func() ProtoMessage { return new(MultistagePlayStageEndNotify) },
		func() ProtoMessage { return new(MultistagePlayEndNotify) },
		func() ProtoMessage { return new(FleurFairBuffEnergyNotify) },
		func() ProtoMessage { return new(FleurFairStageSettleNotify) },
		func() ProtoMessage { return new(FleurFairFinishGalleryStageNotify) },
		func() ProtoMessage { return new(HideAndSeekSelectAvatarReq) },
		func() ProtoMessage { return new(HideAndSeekSelectAvatarRsp) },
		func() ProtoMessage { return new(HideAndSeekSetReadyReq) },
		func() ProtoMessage { return new(HideAndSeekSetReadyRsp) },
		func() ProtoMessage { return new(HideAndSeekSettleNotify) },
		func() ProtoMessage { return new(HideAndSeekPlayerReadyNotify) },
		func() ProtoMessage { return new(HideAndSeekPlayerSetAvatarNotify) },
		func() ProtoMessage { return new(ChessPickCardReq) },
		func() ProtoMessage { return new(ChessPickCardRsp) },
		func() ProtoMessage { return new(ChessPickCardNotify) },
		func() ProtoMessage { return new(ChessManualRefreshCardsReq) },
		func() ProtoMessage { return new(ChessManualRefreshCardsRsp) },
		func() ProtoMessage { return new(ChessPlayerInfoNotify) },
		func() ProtoMessage { return new(ChessLeftMonstersNotify) },
		func() ProtoMessage { return new(ChessEscapedMonstersNotify) },
		func() ProtoMessage { return new(ChessSelectedCardsNotify) },
		func() ProtoMessage { return new(GlobalBuildingInfoNotify) },
		func() ProtoMessage { return new(IrodoriChessPlayerInfoNotify) },
		func() ProtoMessage { return new(IrodoriChessLeftMonsterNotify) },
		func() ProtoMessage { return new(BrickBreakerPlayerReadyNotify) },
		func() ProtoMessage { return new(BrickBreakerPlayerSetAvatarNotify) },
		func() ProtoMessage { return new(BrickBreakerPlayerSetSkillNotify) },
		func() ProtoMessage { return new(BrickBreakerSetReadyReq) },
		func() ProtoMessage { return new(BrickBreakerSetReadyRsp) },
		func() ProtoMessage { return new(BrickBreakerSelectAvatarReq) },
		func() ProtoMessage { return new(BrickBreakerSelectAvatarRsp) },
		func() ProtoMessage { return new(BrickBreakerSelectSkillReq) },
		func() ProtoMessage { return new(BrickBreakerSelectSkillRsp) },
		func() ProtoMessage { return new(BrickBreakerSetChangingReq) },
		func() ProtoMessage { return new(BrickBreakerSetChangingRsp) },
		func() ProtoMessage { return new(BrickBreakerPlayerSetChangingNotify) },
		func() ProtoMessage { return new(CoinCollectPrepareStageNotify) },
		func() ProtoMessage { return new(LanV3BoatSettleNotify) },
	)
}

const (
	ProtoCommandMultistagePlayInfoNotify               ProtoCommand = 5372
	ProtoCommandMultistagePlayFinishStageReq           ProtoCommand = 5398
	ProtoCommandInBattleMechanicusExcapeMonsterNotify  ProtoCommand = 5307
	ProtoCommandInBattleMechanicusLeftMonsterNotify    ProtoCommand = 5321
	ProtoCommandInBattleMechanicusBuildingPointsNotify ProtoCommand = 5303
	ProtoCommandInBattleMechanicusPickCardReq          ProtoCommand = 5390
	ProtoCommandInBattleMechanicusPickCardRsp          ProtoCommand = 5373
	ProtoCommandInBattleMechanicusPickCardNotify       ProtoCommand = 5399
	ProtoCommandInBattleMechanicusConfirmCardReq       ProtoCommand = 5331
	ProtoCommandInBattleMechanicusConfirmCardRsp       ProtoCommand = 5375
	ProtoCommandInBattleMechanicusConfirmCardNotify    ProtoCommand = 5348
	ProtoCommandInBattleMechanicusCardResultNotify     ProtoCommand = 5397
	ProtoCommandMultistagePlayFinishStageRsp           ProtoCommand = 5381
	ProtoCommandInBattleMechanicusSettleNotify         ProtoCommand = 5305
	ProtoCommandMultistagePlaySettleNotify             ProtoCommand = 5313
	ProtoCommandMultistagePlayStageEndNotify           ProtoCommand = 5379
	ProtoCommandMultistagePlayEndNotify                ProtoCommand = 5355
	ProtoCommandFleurFairBuffEnergyNotify              ProtoCommand = 5324
	ProtoCommandFleurFairStageSettleNotify             ProtoCommand = 5356
	ProtoCommandFleurFairFinishGalleryStageNotify      ProtoCommand = 5342
	ProtoCommandHideAndSeekSelectAvatarReq             ProtoCommand = 5330
	ProtoCommandHideAndSeekSelectAvatarRsp             ProtoCommand = 5367
	ProtoCommandHideAndSeekSetReadyReq                 ProtoCommand = 5358
	ProtoCommandHideAndSeekSetReadyRsp                 ProtoCommand = 5370
	ProtoCommandHideAndSeekSettleNotify                ProtoCommand = 5317
	ProtoCommandHideAndSeekPlayerReadyNotify           ProtoCommand = 5302
	ProtoCommandHideAndSeekPlayerSetAvatarNotify       ProtoCommand = 5319
	ProtoCommandChessPickCardReq                       ProtoCommand = 5333
	ProtoCommandChessPickCardRsp                       ProtoCommand = 5384
	ProtoCommandChessPickCardNotify                    ProtoCommand = 5380
	ProtoCommandChessManualRefreshCardsReq             ProtoCommand = 5389
	ProtoCommandChessManualRefreshCardsRsp             ProtoCommand = 5359
	ProtoCommandChessPlayerInfoNotify                  ProtoCommand = 5332
	ProtoCommandChessLeftMonstersNotify                ProtoCommand = 5360
	ProtoCommandChessEscapedMonstersNotify             ProtoCommand = 5314
	ProtoCommandChessSelectedCardsNotify               ProtoCommand = 5392
	ProtoCommandGlobalBuildingInfoNotify               ProtoCommand = 5320
	ProtoCommandIrodoriChessPlayerInfoNotify           ProtoCommand = 5364
	ProtoCommandIrodoriChessLeftMonsterNotify          ProtoCommand = 5338
	ProtoCommandBrickBreakerPlayerReadyNotify          ProtoCommand = 5345
	ProtoCommandBrickBreakerPlayerSetAvatarNotify      ProtoCommand = 5308
	ProtoCommandBrickBreakerPlayerSetSkillNotify       ProtoCommand = 5309
	ProtoCommandBrickBreakerSetReadyReq                ProtoCommand = 5352
	ProtoCommandBrickBreakerSetReadyRsp                ProtoCommand = 5377
	ProtoCommandBrickBreakerSelectAvatarReq            ProtoCommand = 5337
	ProtoCommandBrickBreakerSelectAvatarRsp            ProtoCommand = 5385
	ProtoCommandBrickBreakerSelectSkillReq             ProtoCommand = 5325
	ProtoCommandBrickBreakerSelectSkillRsp             ProtoCommand = 5378
	ProtoCommandBrickBreakerSetChangingReq             ProtoCommand = 5336
	ProtoCommandBrickBreakerSetChangingRsp             ProtoCommand = 5354
	ProtoCommandBrickBreakerPlayerSetChangingNotify    ProtoCommand = 5383
	ProtoCommandCoinCollectPrepareStageNotify          ProtoCommand = 6536
	ProtoCommandLanV3BoatSettleNotify                  ProtoCommand = 6539
)

func (*MultistagePlayInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMultistagePlayInfoNotify
}
func (*MultistagePlayInfoNotify) ProtoMessageType() ProtoMessageType {
	return "MultistagePlayInfoNotify"
}

func (*MultistagePlayFinishStageReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMultistagePlayFinishStageReq
}
func (*MultistagePlayFinishStageReq) ProtoMessageType() ProtoMessageType {
	return "MultistagePlayFinishStageReq"
}

func (*InBattleMechanicusExcapeMonsterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusExcapeMonsterNotify
}
func (*InBattleMechanicusExcapeMonsterNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusExcapeMonsterNotify"
}

func (*InBattleMechanicusLeftMonsterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusLeftMonsterNotify
}
func (*InBattleMechanicusLeftMonsterNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusLeftMonsterNotify"
}

func (*InBattleMechanicusBuildingPointsNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusBuildingPointsNotify
}
func (*InBattleMechanicusBuildingPointsNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusBuildingPointsNotify"
}

func (*InBattleMechanicusPickCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusPickCardReq
}
func (*InBattleMechanicusPickCardReq) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusPickCardReq"
}

func (*InBattleMechanicusPickCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusPickCardRsp
}
func (*InBattleMechanicusPickCardRsp) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusPickCardRsp"
}

func (*InBattleMechanicusPickCardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusPickCardNotify
}
func (*InBattleMechanicusPickCardNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusPickCardNotify"
}

func (*InBattleMechanicusConfirmCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusConfirmCardReq
}
func (*InBattleMechanicusConfirmCardReq) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusConfirmCardReq"
}

func (*InBattleMechanicusConfirmCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusConfirmCardRsp
}
func (*InBattleMechanicusConfirmCardRsp) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusConfirmCardRsp"
}

func (*InBattleMechanicusConfirmCardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusConfirmCardNotify
}
func (*InBattleMechanicusConfirmCardNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusConfirmCardNotify"
}

func (*InBattleMechanicusCardResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusCardResultNotify
}
func (*InBattleMechanicusCardResultNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusCardResultNotify"
}

func (*MultistagePlayFinishStageRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMultistagePlayFinishStageRsp
}
func (*MultistagePlayFinishStageRsp) ProtoMessageType() ProtoMessageType {
	return "MultistagePlayFinishStageRsp"
}

func (*InBattleMechanicusSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInBattleMechanicusSettleNotify
}
func (*InBattleMechanicusSettleNotify) ProtoMessageType() ProtoMessageType {
	return "InBattleMechanicusSettleNotify"
}

func (*MultistagePlaySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMultistagePlaySettleNotify
}
func (*MultistagePlaySettleNotify) ProtoMessageType() ProtoMessageType {
	return "MultistagePlaySettleNotify"
}

func (*MultistagePlayStageEndNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMultistagePlayStageEndNotify
}
func (*MultistagePlayStageEndNotify) ProtoMessageType() ProtoMessageType {
	return "MultistagePlayStageEndNotify"
}

func (*MultistagePlayEndNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMultistagePlayEndNotify
}
func (*MultistagePlayEndNotify) ProtoMessageType() ProtoMessageType { return "MultistagePlayEndNotify" }

func (*FleurFairBuffEnergyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairBuffEnergyNotify
}
func (*FleurFairBuffEnergyNotify) ProtoMessageType() ProtoMessageType {
	return "FleurFairBuffEnergyNotify"
}

func (*FleurFairStageSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairStageSettleNotify
}
func (*FleurFairStageSettleNotify) ProtoMessageType() ProtoMessageType {
	return "FleurFairStageSettleNotify"
}

func (*FleurFairFinishGalleryStageNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairFinishGalleryStageNotify
}
func (*FleurFairFinishGalleryStageNotify) ProtoMessageType() ProtoMessageType {
	return "FleurFairFinishGalleryStageNotify"
}

func (*HideAndSeekSelectAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekSelectAvatarReq
}
func (*HideAndSeekSelectAvatarReq) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekSelectAvatarReq"
}

func (*HideAndSeekSelectAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekSelectAvatarRsp
}
func (*HideAndSeekSelectAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekSelectAvatarRsp"
}

func (*HideAndSeekSetReadyReq) ProtoCommand() ProtoCommand         { return ProtoCommandHideAndSeekSetReadyReq }
func (*HideAndSeekSetReadyReq) ProtoMessageType() ProtoMessageType { return "HideAndSeekSetReadyReq" }

func (*HideAndSeekSetReadyRsp) ProtoCommand() ProtoCommand         { return ProtoCommandHideAndSeekSetReadyRsp }
func (*HideAndSeekSetReadyRsp) ProtoMessageType() ProtoMessageType { return "HideAndSeekSetReadyRsp" }

func (*HideAndSeekSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekSettleNotify
}
func (*HideAndSeekSettleNotify) ProtoMessageType() ProtoMessageType { return "HideAndSeekSettleNotify" }

func (*HideAndSeekPlayerReadyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekPlayerReadyNotify
}
func (*HideAndSeekPlayerReadyNotify) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekPlayerReadyNotify"
}

func (*HideAndSeekPlayerSetAvatarNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekPlayerSetAvatarNotify
}
func (*HideAndSeekPlayerSetAvatarNotify) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekPlayerSetAvatarNotify"
}

func (*ChessPickCardReq) ProtoCommand() ProtoCommand         { return ProtoCommandChessPickCardReq }
func (*ChessPickCardReq) ProtoMessageType() ProtoMessageType { return "ChessPickCardReq" }

func (*ChessPickCardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandChessPickCardRsp }
func (*ChessPickCardRsp) ProtoMessageType() ProtoMessageType { return "ChessPickCardRsp" }

func (*ChessPickCardNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChessPickCardNotify }
func (*ChessPickCardNotify) ProtoMessageType() ProtoMessageType { return "ChessPickCardNotify" }

func (*ChessManualRefreshCardsReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChessManualRefreshCardsReq
}
func (*ChessManualRefreshCardsReq) ProtoMessageType() ProtoMessageType {
	return "ChessManualRefreshCardsReq"
}

func (*ChessManualRefreshCardsRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChessManualRefreshCardsRsp
}
func (*ChessManualRefreshCardsRsp) ProtoMessageType() ProtoMessageType {
	return "ChessManualRefreshCardsRsp"
}

func (*ChessPlayerInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChessPlayerInfoNotify }
func (*ChessPlayerInfoNotify) ProtoMessageType() ProtoMessageType { return "ChessPlayerInfoNotify" }

func (*ChessLeftMonstersNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChessLeftMonstersNotify
}
func (*ChessLeftMonstersNotify) ProtoMessageType() ProtoMessageType { return "ChessLeftMonstersNotify" }

func (*ChessEscapedMonstersNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChessEscapedMonstersNotify
}
func (*ChessEscapedMonstersNotify) ProtoMessageType() ProtoMessageType {
	return "ChessEscapedMonstersNotify"
}

func (*ChessSelectedCardsNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChessSelectedCardsNotify
}
func (*ChessSelectedCardsNotify) ProtoMessageType() ProtoMessageType {
	return "ChessSelectedCardsNotify"
}

func (*GlobalBuildingInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGlobalBuildingInfoNotify
}
func (*GlobalBuildingInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GlobalBuildingInfoNotify"
}

func (*IrodoriChessPlayerInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriChessPlayerInfoNotify
}
func (*IrodoriChessPlayerInfoNotify) ProtoMessageType() ProtoMessageType {
	return "IrodoriChessPlayerInfoNotify"
}

func (*IrodoriChessLeftMonsterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriChessLeftMonsterNotify
}
func (*IrodoriChessLeftMonsterNotify) ProtoMessageType() ProtoMessageType {
	return "IrodoriChessLeftMonsterNotify"
}

func (*BrickBreakerPlayerReadyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerPlayerReadyNotify
}
func (*BrickBreakerPlayerReadyNotify) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerPlayerReadyNotify"
}

func (*BrickBreakerPlayerSetAvatarNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerPlayerSetAvatarNotify
}
func (*BrickBreakerPlayerSetAvatarNotify) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerPlayerSetAvatarNotify"
}

func (*BrickBreakerPlayerSetSkillNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerPlayerSetSkillNotify
}
func (*BrickBreakerPlayerSetSkillNotify) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerPlayerSetSkillNotify"
}

func (*BrickBreakerSetReadyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSetReadyReq
}
func (*BrickBreakerSetReadyReq) ProtoMessageType() ProtoMessageType { return "BrickBreakerSetReadyReq" }

func (*BrickBreakerSetReadyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSetReadyRsp
}
func (*BrickBreakerSetReadyRsp) ProtoMessageType() ProtoMessageType { return "BrickBreakerSetReadyRsp" }

func (*BrickBreakerSelectAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSelectAvatarReq
}
func (*BrickBreakerSelectAvatarReq) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSelectAvatarReq"
}

func (*BrickBreakerSelectAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSelectAvatarRsp
}
func (*BrickBreakerSelectAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSelectAvatarRsp"
}

func (*BrickBreakerSelectSkillReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSelectSkillReq
}
func (*BrickBreakerSelectSkillReq) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSelectSkillReq"
}

func (*BrickBreakerSelectSkillRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSelectSkillRsp
}
func (*BrickBreakerSelectSkillRsp) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSelectSkillRsp"
}

func (*BrickBreakerSetChangingReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSetChangingReq
}
func (*BrickBreakerSetChangingReq) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSetChangingReq"
}

func (*BrickBreakerSetChangingRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSetChangingRsp
}
func (*BrickBreakerSetChangingRsp) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSetChangingRsp"
}

func (*BrickBreakerPlayerSetChangingNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerPlayerSetChangingNotify
}
func (*BrickBreakerPlayerSetChangingNotify) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerPlayerSetChangingNotify"
}

func (*CoinCollectPrepareStageNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectPrepareStageNotify
}
func (*CoinCollectPrepareStageNotify) ProtoMessageType() ProtoMessageType {
	return "CoinCollectPrepareStageNotify"
}

func (*LanV3BoatSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLanV3BoatSettleNotify }
func (*LanV3BoatSettleNotify) ProtoMessageType() ProtoMessageType { return "LanV3BoatSettleNotify" }
