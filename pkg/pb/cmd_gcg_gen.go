
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GCGOperationReq) },
		func() ProtoMessage { return new(GCGOperationRsp) },
		func() ProtoMessage { return new(GCGMessagePackNotify) },
		func() ProtoMessage { return new(GCGAskDuelReq) },
		func() ProtoMessage { return new(GCGAskDuelRsp) },
		func() ProtoMessage { return new(GCGGameBriefDataNotify) },
		func() ProtoMessage { return new(GCGInitFinishReq) },
		func() ProtoMessage { return new(GCGInitFinishRsp) },
		func() ProtoMessage { return new(GCGHeartBeatNotify) },
		func() ProtoMessage { return new(GCGSkillPreviewNotify) },
		func() ProtoMessage { return new(GCGSkillPreviewAskReq) },
		func() ProtoMessage { return new(GCGSkillPreviewAskRsp) },
		func() ProtoMessage { return new(GCGChallengeUpdateNotify) },
		func() ProtoMessage { return new(GCGBackToDuelReq) },
		func() ProtoMessage { return new(GCGBackToDuelRsp) },
		func() ProtoMessage { return new(GCGDebugReplayNotify) },
		func() ProtoMessage { return new(GCGGameMaxNotify) },
		func() ProtoMessage { return new(GCGGameCreateFailReasonNotify) },
		func() ProtoMessage { return new(GCGDSDataNotify) },
		func() ProtoMessage { return new(GCGDSCardBackUnlockNotify) },
		func() ProtoMessage { return new(GCGDSFieldUnlockNotify) },
		func() ProtoMessage { return new(GCGDSCardFaceUnlockNotify) },
		func() ProtoMessage { return new(GCGDSCardNumChangeNotify) },
		func() ProtoMessage { return new(GCGDSChangeCardFaceReq) },
		func() ProtoMessage { return new(GCGDSChangeCardFaceRsp) },
		func() ProtoMessage { return new(GCGDSChangeCardBackReq) },
		func() ProtoMessage { return new(GCGDSChangeCardBackRsp) },
		func() ProtoMessage { return new(GCGDSChangeFieldReq) },
		func() ProtoMessage { return new(GCGDSChangeFieldRsp) },
		func() ProtoMessage { return new(GCGDSChangeDeckNameReq) },
		func() ProtoMessage { return new(GCGDSChangeDeckNameRsp) },
		func() ProtoMessage { return new(GCGDSDeckSaveReq) },
		func() ProtoMessage { return new(GCGDSDeckSaveRsp) },
		func() ProtoMessage { return new(GCGDSChangeCurDeckReq) },
		func() ProtoMessage { return new(GCGDSChangeCurDeckRsp) },
		func() ProtoMessage { return new(GCGDSCurDeckChangeNotify) },
		func() ProtoMessage { return new(GCGDSDeleteDeckReq) },
		func() ProtoMessage { return new(GCGDSDeleteDeckRsp) },
		func() ProtoMessage { return new(GCGDSDeckUnlockNotify) },
		func() ProtoMessage { return new(GCGDSCardProficiencyNotify) },
		func() ProtoMessage { return new(GCGDSDeckUpdateNotify) },
		func() ProtoMessage { return new(GCGDSCardFaceUpdateNotify) },
		func() ProtoMessage { return new(GCGDSTakeCardProficiencyRewardReq) },
		func() ProtoMessage { return new(GCGDSTakeCardProficiencyRewardRsp) },
		func() ProtoMessage { return new(GCGDSBanCardNotify) },
		func() ProtoMessage { return new(GCGTCTavernInfoNotify) },
		func() ProtoMessage { return new(GCGTCInviteReq) },
		func() ProtoMessage { return new(GCGTCInviteRsp) },
		func() ProtoMessage { return new(GCGTCTavernChallengeDataNotify) },
		func() ProtoMessage { return new(GCGTCTavernChallengeUpdateNotify) },
		func() ProtoMessage { return new(GCGGrowthLevelNotify) },
		func() ProtoMessage { return new(GCGGrowthLevelRewardNotify) },
		func() ProtoMessage { return new(GCGGrowthLevelTakeRewardReq) },
		func() ProtoMessage { return new(GCGGrowthLevelTakeRewardRsp) },
		func() ProtoMessage { return new(GCGInviteGuestBattleReq) },
		func() ProtoMessage { return new(GCGInviteGuestBattleRsp) },
		func() ProtoMessage { return new(GCGInviteBattleNotify) },
		func() ProtoMessage { return new(GCGApplyInviteBattleReq) },
		func() ProtoMessage { return new(GCGApplyInviteBattleRsp) },
		func() ProtoMessage { return new(GCGApplyInviteBattleNotify) },
		func() ProtoMessage { return new(GCGWorldPlayerGCGStateReq) },
		func() ProtoMessage { return new(GCGWorldPlayerGCGStateRsp) },
		func() ProtoMessage { return new(GCGSettleNotify) },
		func() ProtoMessage { return new(GCGResourceStateNotify) },
		func() ProtoMessage { return new(GCGClientSettleReq) },
		func() ProtoMessage { return new(GCGClientSettleRsp) },
		func() ProtoMessage { return new(GCGSettleOptionReq) },
		func() ProtoMessage { return new(GCGSettleOptionRsp) },
		func() ProtoMessage { return new(GCGBasicDataNotify) },
		func() ProtoMessage { return new(GCGTavernNpcInfoNotify) },
		func() ProtoMessage { return new(GCGStartChallengeReq) },
		func() ProtoMessage { return new(GCGStartChallengeRsp) },
		func() ProtoMessage { return new(GCGWeekChallengeInfoNotify) },
		func() ProtoMessage { return new(GCGStartChallengeByCheckRewardReq) },
		func() ProtoMessage { return new(GCGStartChallengeByCheckRewardRsp) },
		func() ProtoMessage { return new(GCGLevelChallengeNotify) },
		func() ProtoMessage { return new(GCGLevelChallengeFinishNotify) },
		func() ProtoMessage { return new(GCGWorldChallengeUnlockNotify) },
		func() ProtoMessage { return new(GCGBossChallengeUpdateNotify) },
		func() ProtoMessage { return new(GCGLevelChallengeDeleteNotify) },
	)
}

const (
	ProtoCommandGCGOperationReq                   ProtoCommand = 7107
	ProtoCommandGCGOperationRsp                   ProtoCommand = 7600
	ProtoCommandGCGMessagePackNotify              ProtoCommand = 7516
	ProtoCommandGCGAskDuelReq                     ProtoCommand = 7237
	ProtoCommandGCGAskDuelRsp                     ProtoCommand = 7869
	ProtoCommandGCGGameBriefDataNotify            ProtoCommand = 7539
	ProtoCommandGCGInitFinishReq                  ProtoCommand = 7684
	ProtoCommandGCGInitFinishRsp                  ProtoCommand = 7433
	ProtoCommandGCGHeartBeatNotify                ProtoCommand = 7224
	ProtoCommandGCGSkillPreviewNotify             ProtoCommand = 7503
	ProtoCommandGCGSkillPreviewAskReq             ProtoCommand = 7509
	ProtoCommandGCGSkillPreviewAskRsp             ProtoCommand = 7409
	ProtoCommandGCGChallengeUpdateNotify          ProtoCommand = 7268
	ProtoCommandGCGBackToDuelReq                  ProtoCommand = 7729
	ProtoCommandGCGBackToDuelRsp                  ProtoCommand = 7416
	ProtoCommandGCGDebugReplayNotify              ProtoCommand = 7932
	ProtoCommandGCGGameMaxNotify                  ProtoCommand = 7666
	ProtoCommandGCGGameCreateFailReasonNotify     ProtoCommand = 7239
	ProtoCommandGCGDSDataNotify                   ProtoCommand = 7122
	ProtoCommandGCGDSCardBackUnlockNotify         ProtoCommand = 7265
	ProtoCommandGCGDSFieldUnlockNotify            ProtoCommand = 7333
	ProtoCommandGCGDSCardFaceUnlockNotify         ProtoCommand = 7049
	ProtoCommandGCGDSCardNumChangeNotify          ProtoCommand = 7358
	ProtoCommandGCGDSChangeCardFaceReq            ProtoCommand = 7169
	ProtoCommandGCGDSChangeCardFaceRsp            ProtoCommand = 7331
	ProtoCommandGCGDSChangeCardBackReq            ProtoCommand = 7292
	ProtoCommandGCGDSChangeCardBackRsp            ProtoCommand = 7044
	ProtoCommandGCGDSChangeFieldReq               ProtoCommand = 7541
	ProtoCommandGCGDSChangeFieldRsp               ProtoCommand = 7444
	ProtoCommandGCGDSChangeDeckNameReq            ProtoCommand = 7432
	ProtoCommandGCGDSChangeDeckNameRsp            ProtoCommand = 7916
	ProtoCommandGCGDSDeckSaveReq                  ProtoCommand = 7104
	ProtoCommandGCGDSDeckSaveRsp                  ProtoCommand = 7269
	ProtoCommandGCGDSChangeCurDeckReq             ProtoCommand = 7131
	ProtoCommandGCGDSChangeCurDeckRsp             ProtoCommand = 7301
	ProtoCommandGCGDSCurDeckChangeNotify          ProtoCommand = 7796
	ProtoCommandGCGDSDeleteDeckReq                ProtoCommand = 7988
	ProtoCommandGCGDSDeleteDeckRsp                ProtoCommand = 7524
	ProtoCommandGCGDSDeckUnlockNotify             ProtoCommand = 7732
	ProtoCommandGCGDSCardProficiencyNotify        ProtoCommand = 7680
	ProtoCommandGCGDSDeckUpdateNotify             ProtoCommand = 7305
	ProtoCommandGCGDSCardFaceUpdateNotify         ProtoCommand = 7851
	ProtoCommandGCGDSTakeCardProficiencyRewardReq ProtoCommand = 7581
	ProtoCommandGCGDSTakeCardProficiencyRewardRsp ProtoCommand = 7889
	ProtoCommandGCGDSBanCardNotify                ProtoCommand = 7839
	ProtoCommandGCGTCTavernInfoNotify             ProtoCommand = 7011
	ProtoCommandGCGTCInviteReq                    ProtoCommand = 7922
	ProtoCommandGCGTCInviteRsp                    ProtoCommand = 7328
	ProtoCommandGCGTCTavernChallengeDataNotify    ProtoCommand = 7294
	ProtoCommandGCGTCTavernChallengeUpdateNotify  ProtoCommand = 7184
	ProtoCommandGCGGrowthLevelNotify              ProtoCommand = 7736
	ProtoCommandGCGGrowthLevelRewardNotify        ProtoCommand = 7477
	ProtoCommandGCGGrowthLevelTakeRewardReq       ProtoCommand = 7051
	ProtoCommandGCGGrowthLevelTakeRewardRsp       ProtoCommand = 7670
	ProtoCommandGCGInviteGuestBattleReq           ProtoCommand = 7783
	ProtoCommandGCGInviteGuestBattleRsp           ProtoCommand = 7251
	ProtoCommandGCGInviteBattleNotify             ProtoCommand = 7692
	ProtoCommandGCGApplyInviteBattleReq           ProtoCommand = 7730
	ProtoCommandGCGApplyInviteBattleRsp           ProtoCommand = 7304
	ProtoCommandGCGApplyInviteBattleNotify        ProtoCommand = 7820
	ProtoCommandGCGWorldPlayerGCGStateReq         ProtoCommand = 7206
	ProtoCommandGCGWorldPlayerGCGStateRsp         ProtoCommand = 7136
	ProtoCommandGCGSettleNotify                   ProtoCommand = 7769
	ProtoCommandGCGResourceStateNotify            ProtoCommand = 7876
	ProtoCommandGCGClientSettleReq                ProtoCommand = 7506
	ProtoCommandGCGClientSettleRsp                ProtoCommand = 7105
	ProtoCommandGCGSettleOptionReq                ProtoCommand = 7124
	ProtoCommandGCGSettleOptionRsp                ProtoCommand = 7735
	ProtoCommandGCGBasicDataNotify                ProtoCommand = 7319
	ProtoCommandGCGTavernNpcInfoNotify            ProtoCommand = 7290
	ProtoCommandGCGStartChallengeReq              ProtoCommand = 7595
	ProtoCommandGCGStartChallengeRsp              ProtoCommand = 7763
	ProtoCommandGCGWeekChallengeInfoNotify        ProtoCommand = 7615
	ProtoCommandGCGStartChallengeByCheckRewardReq ProtoCommand = 7778
	ProtoCommandGCGStartChallengeByCheckRewardRsp ProtoCommand = 7619
	ProtoCommandGCGLevelChallengeNotify           ProtoCommand = 7055
	ProtoCommandGCGLevelChallengeFinishNotify     ProtoCommand = 7629
	ProtoCommandGCGWorldChallengeUnlockNotify     ProtoCommand = 7204
	ProtoCommandGCGBossChallengeUpdateNotify      ProtoCommand = 7073
	ProtoCommandGCGLevelChallengeDeleteNotify     ProtoCommand = 7648
)

func (*GCGOperationReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGOperationReq }
func (*GCGOperationReq) ProtoMessageType() ProtoMessageType { return "GCGOperationReq" }

func (*GCGOperationRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGOperationRsp }
func (*GCGOperationRsp) ProtoMessageType() ProtoMessageType { return "GCGOperationRsp" }

func (*GCGMessagePackNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGMessagePackNotify }
func (*GCGMessagePackNotify) ProtoMessageType() ProtoMessageType { return "GCGMessagePackNotify" }

func (*GCGAskDuelReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGAskDuelReq }
func (*GCGAskDuelReq) ProtoMessageType() ProtoMessageType { return "GCGAskDuelReq" }

func (*GCGAskDuelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGAskDuelRsp }
func (*GCGAskDuelRsp) ProtoMessageType() ProtoMessageType { return "GCGAskDuelRsp" }

func (*GCGGameBriefDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGGameBriefDataNotify }
func (*GCGGameBriefDataNotify) ProtoMessageType() ProtoMessageType { return "GCGGameBriefDataNotify" }

func (*GCGInitFinishReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGInitFinishReq }
func (*GCGInitFinishReq) ProtoMessageType() ProtoMessageType { return "GCGInitFinishReq" }

func (*GCGInitFinishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGInitFinishRsp }
func (*GCGInitFinishRsp) ProtoMessageType() ProtoMessageType { return "GCGInitFinishRsp" }

func (*GCGHeartBeatNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGHeartBeatNotify }
func (*GCGHeartBeatNotify) ProtoMessageType() ProtoMessageType { return "GCGHeartBeatNotify" }

func (*GCGSkillPreviewNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGSkillPreviewNotify }
func (*GCGSkillPreviewNotify) ProtoMessageType() ProtoMessageType { return "GCGSkillPreviewNotify" }

func (*GCGSkillPreviewAskReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGSkillPreviewAskReq }
func (*GCGSkillPreviewAskReq) ProtoMessageType() ProtoMessageType { return "GCGSkillPreviewAskReq" }

func (*GCGSkillPreviewAskRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGSkillPreviewAskRsp }
func (*GCGSkillPreviewAskRsp) ProtoMessageType() ProtoMessageType { return "GCGSkillPreviewAskRsp" }

func (*GCGChallengeUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGChallengeUpdateNotify
}
func (*GCGChallengeUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "GCGChallengeUpdateNotify"
}

func (*GCGBackToDuelReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGBackToDuelReq }
func (*GCGBackToDuelReq) ProtoMessageType() ProtoMessageType { return "GCGBackToDuelReq" }

func (*GCGBackToDuelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGBackToDuelRsp }
func (*GCGBackToDuelRsp) ProtoMessageType() ProtoMessageType { return "GCGBackToDuelRsp" }

func (*GCGDebugReplayNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDebugReplayNotify }
func (*GCGDebugReplayNotify) ProtoMessageType() ProtoMessageType { return "GCGDebugReplayNotify" }

func (*GCGGameMaxNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGGameMaxNotify }
func (*GCGGameMaxNotify) ProtoMessageType() ProtoMessageType { return "GCGGameMaxNotify" }

func (*GCGGameCreateFailReasonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGGameCreateFailReasonNotify
}
func (*GCGGameCreateFailReasonNotify) ProtoMessageType() ProtoMessageType {
	return "GCGGameCreateFailReasonNotify"
}

func (*GCGDSDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDataNotify }
func (*GCGDSDataNotify) ProtoMessageType() ProtoMessageType { return "GCGDSDataNotify" }

func (*GCGDSCardBackUnlockNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSCardBackUnlockNotify
}
func (*GCGDSCardBackUnlockNotify) ProtoMessageType() ProtoMessageType {
	return "GCGDSCardBackUnlockNotify"
}

func (*GCGDSFieldUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSFieldUnlockNotify }
func (*GCGDSFieldUnlockNotify) ProtoMessageType() ProtoMessageType { return "GCGDSFieldUnlockNotify" }

func (*GCGDSCardFaceUnlockNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSCardFaceUnlockNotify
}
func (*GCGDSCardFaceUnlockNotify) ProtoMessageType() ProtoMessageType {
	return "GCGDSCardFaceUnlockNotify"
}

func (*GCGDSCardNumChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSCardNumChangeNotify
}
func (*GCGDSCardNumChangeNotify) ProtoMessageType() ProtoMessageType {
	return "GCGDSCardNumChangeNotify"
}

func (*GCGDSChangeCardFaceReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeCardFaceReq }
func (*GCGDSChangeCardFaceReq) ProtoMessageType() ProtoMessageType { return "GCGDSChangeCardFaceReq" }

func (*GCGDSChangeCardFaceRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeCardFaceRsp }
func (*GCGDSChangeCardFaceRsp) ProtoMessageType() ProtoMessageType { return "GCGDSChangeCardFaceRsp" }

func (*GCGDSChangeCardBackReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeCardBackReq }
func (*GCGDSChangeCardBackReq) ProtoMessageType() ProtoMessageType { return "GCGDSChangeCardBackReq" }

func (*GCGDSChangeCardBackRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeCardBackRsp }
func (*GCGDSChangeCardBackRsp) ProtoMessageType() ProtoMessageType { return "GCGDSChangeCardBackRsp" }

func (*GCGDSChangeFieldReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeFieldReq }
func (*GCGDSChangeFieldReq) ProtoMessageType() ProtoMessageType { return "GCGDSChangeFieldReq" }

func (*GCGDSChangeFieldRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeFieldRsp }
func (*GCGDSChangeFieldRsp) ProtoMessageType() ProtoMessageType { return "GCGDSChangeFieldRsp" }

func (*GCGDSChangeDeckNameReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeDeckNameReq }
func (*GCGDSChangeDeckNameReq) ProtoMessageType() ProtoMessageType { return "GCGDSChangeDeckNameReq" }

func (*GCGDSChangeDeckNameRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeDeckNameRsp }
func (*GCGDSChangeDeckNameRsp) ProtoMessageType() ProtoMessageType { return "GCGDSChangeDeckNameRsp" }

func (*GCGDSDeckSaveReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDeckSaveReq }
func (*GCGDSDeckSaveReq) ProtoMessageType() ProtoMessageType { return "GCGDSDeckSaveReq" }

func (*GCGDSDeckSaveRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDeckSaveRsp }
func (*GCGDSDeckSaveRsp) ProtoMessageType() ProtoMessageType { return "GCGDSDeckSaveRsp" }

func (*GCGDSChangeCurDeckReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeCurDeckReq }
func (*GCGDSChangeCurDeckReq) ProtoMessageType() ProtoMessageType { return "GCGDSChangeCurDeckReq" }

func (*GCGDSChangeCurDeckRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSChangeCurDeckRsp }
func (*GCGDSChangeCurDeckRsp) ProtoMessageType() ProtoMessageType { return "GCGDSChangeCurDeckRsp" }

func (*GCGDSCurDeckChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSCurDeckChangeNotify
}
func (*GCGDSCurDeckChangeNotify) ProtoMessageType() ProtoMessageType {
	return "GCGDSCurDeckChangeNotify"
}

func (*GCGDSDeleteDeckReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDeleteDeckReq }
func (*GCGDSDeleteDeckReq) ProtoMessageType() ProtoMessageType { return "GCGDSDeleteDeckReq" }

func (*GCGDSDeleteDeckRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDeleteDeckRsp }
func (*GCGDSDeleteDeckRsp) ProtoMessageType() ProtoMessageType { return "GCGDSDeleteDeckRsp" }

func (*GCGDSDeckUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDeckUnlockNotify }
func (*GCGDSDeckUnlockNotify) ProtoMessageType() ProtoMessageType { return "GCGDSDeckUnlockNotify" }

func (*GCGDSCardProficiencyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSCardProficiencyNotify
}
func (*GCGDSCardProficiencyNotify) ProtoMessageType() ProtoMessageType {
	return "GCGDSCardProficiencyNotify"
}

func (*GCGDSDeckUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSDeckUpdateNotify }
func (*GCGDSDeckUpdateNotify) ProtoMessageType() ProtoMessageType { return "GCGDSDeckUpdateNotify" }

func (*GCGDSCardFaceUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSCardFaceUpdateNotify
}
func (*GCGDSCardFaceUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "GCGDSCardFaceUpdateNotify"
}

func (*GCGDSTakeCardProficiencyRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSTakeCardProficiencyRewardReq
}
func (*GCGDSTakeCardProficiencyRewardReq) ProtoMessageType() ProtoMessageType {
	return "GCGDSTakeCardProficiencyRewardReq"
}

func (*GCGDSTakeCardProficiencyRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGDSTakeCardProficiencyRewardRsp
}
func (*GCGDSTakeCardProficiencyRewardRsp) ProtoMessageType() ProtoMessageType {
	return "GCGDSTakeCardProficiencyRewardRsp"
}

func (*GCGDSBanCardNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGDSBanCardNotify }
func (*GCGDSBanCardNotify) ProtoMessageType() ProtoMessageType { return "GCGDSBanCardNotify" }

func (*GCGTCTavernInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGTCTavernInfoNotify }
func (*GCGTCTavernInfoNotify) ProtoMessageType() ProtoMessageType { return "GCGTCTavernInfoNotify" }

func (*GCGTCInviteReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGTCInviteReq }
func (*GCGTCInviteReq) ProtoMessageType() ProtoMessageType { return "GCGTCInviteReq" }

func (*GCGTCInviteRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGTCInviteRsp }
func (*GCGTCInviteRsp) ProtoMessageType() ProtoMessageType { return "GCGTCInviteRsp" }

func (*GCGTCTavernChallengeDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGTCTavernChallengeDataNotify
}
func (*GCGTCTavernChallengeDataNotify) ProtoMessageType() ProtoMessageType {
	return "GCGTCTavernChallengeDataNotify"
}

func (*GCGTCTavernChallengeUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGTCTavernChallengeUpdateNotify
}
func (*GCGTCTavernChallengeUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "GCGTCTavernChallengeUpdateNotify"
}

func (*GCGGrowthLevelNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGGrowthLevelNotify }
func (*GCGGrowthLevelNotify) ProtoMessageType() ProtoMessageType { return "GCGGrowthLevelNotify" }

func (*GCGGrowthLevelRewardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGGrowthLevelRewardNotify
}
func (*GCGGrowthLevelRewardNotify) ProtoMessageType() ProtoMessageType {
	return "GCGGrowthLevelRewardNotify"
}

func (*GCGGrowthLevelTakeRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGGrowthLevelTakeRewardReq
}
func (*GCGGrowthLevelTakeRewardReq) ProtoMessageType() ProtoMessageType {
	return "GCGGrowthLevelTakeRewardReq"
}

func (*GCGGrowthLevelTakeRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGGrowthLevelTakeRewardRsp
}
func (*GCGGrowthLevelTakeRewardRsp) ProtoMessageType() ProtoMessageType {
	return "GCGGrowthLevelTakeRewardRsp"
}

func (*GCGInviteGuestBattleReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGInviteGuestBattleReq
}
func (*GCGInviteGuestBattleReq) ProtoMessageType() ProtoMessageType { return "GCGInviteGuestBattleReq" }

func (*GCGInviteGuestBattleRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGInviteGuestBattleRsp
}
func (*GCGInviteGuestBattleRsp) ProtoMessageType() ProtoMessageType { return "GCGInviteGuestBattleRsp" }

func (*GCGInviteBattleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGInviteBattleNotify }
func (*GCGInviteBattleNotify) ProtoMessageType() ProtoMessageType { return "GCGInviteBattleNotify" }

func (*GCGApplyInviteBattleReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGApplyInviteBattleReq
}
func (*GCGApplyInviteBattleReq) ProtoMessageType() ProtoMessageType { return "GCGApplyInviteBattleReq" }

func (*GCGApplyInviteBattleRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGApplyInviteBattleRsp
}
func (*GCGApplyInviteBattleRsp) ProtoMessageType() ProtoMessageType { return "GCGApplyInviteBattleRsp" }

func (*GCGApplyInviteBattleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGApplyInviteBattleNotify
}
func (*GCGApplyInviteBattleNotify) ProtoMessageType() ProtoMessageType {
	return "GCGApplyInviteBattleNotify"
}

func (*GCGWorldPlayerGCGStateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGWorldPlayerGCGStateReq
}
func (*GCGWorldPlayerGCGStateReq) ProtoMessageType() ProtoMessageType {
	return "GCGWorldPlayerGCGStateReq"
}

func (*GCGWorldPlayerGCGStateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGWorldPlayerGCGStateRsp
}
func (*GCGWorldPlayerGCGStateRsp) ProtoMessageType() ProtoMessageType {
	return "GCGWorldPlayerGCGStateRsp"
}

func (*GCGSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGSettleNotify }
func (*GCGSettleNotify) ProtoMessageType() ProtoMessageType { return "GCGSettleNotify" }

func (*GCGResourceStateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGResourceStateNotify }
func (*GCGResourceStateNotify) ProtoMessageType() ProtoMessageType { return "GCGResourceStateNotify" }

func (*GCGClientSettleReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGClientSettleReq }
func (*GCGClientSettleReq) ProtoMessageType() ProtoMessageType { return "GCGClientSettleReq" }

func (*GCGClientSettleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGClientSettleRsp }
func (*GCGClientSettleRsp) ProtoMessageType() ProtoMessageType { return "GCGClientSettleRsp" }

func (*GCGSettleOptionReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGSettleOptionReq }
func (*GCGSettleOptionReq) ProtoMessageType() ProtoMessageType { return "GCGSettleOptionReq" }

func (*GCGSettleOptionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGSettleOptionRsp }
func (*GCGSettleOptionRsp) ProtoMessageType() ProtoMessageType { return "GCGSettleOptionRsp" }

func (*GCGBasicDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGBasicDataNotify }
func (*GCGBasicDataNotify) ProtoMessageType() ProtoMessageType { return "GCGBasicDataNotify" }

func (*GCGTavernNpcInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGCGTavernNpcInfoNotify }
func (*GCGTavernNpcInfoNotify) ProtoMessageType() ProtoMessageType { return "GCGTavernNpcInfoNotify" }

func (*GCGStartChallengeReq) ProtoCommand() ProtoCommand         { return ProtoCommandGCGStartChallengeReq }
func (*GCGStartChallengeReq) ProtoMessageType() ProtoMessageType { return "GCGStartChallengeReq" }

func (*GCGStartChallengeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGCGStartChallengeRsp }
func (*GCGStartChallengeRsp) ProtoMessageType() ProtoMessageType { return "GCGStartChallengeRsp" }

func (*GCGWeekChallengeInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGWeekChallengeInfoNotify
}
func (*GCGWeekChallengeInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GCGWeekChallengeInfoNotify"
}

func (*GCGStartChallengeByCheckRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGStartChallengeByCheckRewardReq
}
func (*GCGStartChallengeByCheckRewardReq) ProtoMessageType() ProtoMessageType {
	return "GCGStartChallengeByCheckRewardReq"
}

func (*GCGStartChallengeByCheckRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGStartChallengeByCheckRewardRsp
}
func (*GCGStartChallengeByCheckRewardRsp) ProtoMessageType() ProtoMessageType {
	return "GCGStartChallengeByCheckRewardRsp"
}

func (*GCGLevelChallengeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGLevelChallengeNotify
}
func (*GCGLevelChallengeNotify) ProtoMessageType() ProtoMessageType { return "GCGLevelChallengeNotify" }

func (*GCGLevelChallengeFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGLevelChallengeFinishNotify
}
func (*GCGLevelChallengeFinishNotify) ProtoMessageType() ProtoMessageType {
	return "GCGLevelChallengeFinishNotify"
}

func (*GCGWorldChallengeUnlockNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGWorldChallengeUnlockNotify
}
func (*GCGWorldChallengeUnlockNotify) ProtoMessageType() ProtoMessageType {
	return "GCGWorldChallengeUnlockNotify"
}

func (*GCGBossChallengeUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGBossChallengeUpdateNotify
}
func (*GCGBossChallengeUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "GCGBossChallengeUpdateNotify"
}

func (*GCGLevelChallengeDeleteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGCGLevelChallengeDeleteNotify
}
func (*GCGLevelChallengeDeleteNotify) ProtoMessageType() ProtoMessageType {
	return "GCGLevelChallengeDeleteNotify"
}
