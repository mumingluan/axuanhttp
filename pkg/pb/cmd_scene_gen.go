
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerEnterSceneNotify) },
		func() ProtoMessage { return new(LeaveSceneReq) },
		func() ProtoMessage { return new(LeaveSceneRsp) },
		func() ProtoMessage { return new(SceneInitFinishReq) },
		func() ProtoMessage { return new(SceneInitFinishRsp) },
		func() ProtoMessage { return new(SceneEntityAppearNotify) },
		func() ProtoMessage { return new(SceneEntityDisappearNotify) },
		func() ProtoMessage { return new(SceneEntityMoveReq) },
		func() ProtoMessage { return new(SceneEntityMoveRsp) },
		func() ProtoMessage { return new(SceneAvatarStaminaStepReq) },
		func() ProtoMessage { return new(SceneAvatarStaminaStepRsp) },
		func() ProtoMessage { return new(SceneEntityMoveNotify) },
		func() ProtoMessage { return new(ScenePlayerLocationNotify) },
		func() ProtoMessage { return new(GetScenePointReq) },
		func() ProtoMessage { return new(GetScenePointRsp) },
		func() ProtoMessage { return new(EnterTransPointRegionNotify) },
		func() ProtoMessage { return new(ExitTransPointRegionNotify) },
		func() ProtoMessage { return new(ScenePointUnlockNotify) },
		func() ProtoMessage { return new(SceneTransToPointReq) },
		func() ProtoMessage { return new(SceneTransToPointRsp) },
		func() ProtoMessage { return new(EntityJumpNotify) },
		func() ProtoMessage { return new(GetSceneAreaReq) },
		func() ProtoMessage { return new(GetSceneAreaRsp) },
		func() ProtoMessage { return new(SceneAreaUnlockNotify) },
		func() ProtoMessage { return new(SceneEntityDrownReq) },
		func() ProtoMessage { return new(SceneEntityDrownRsp) },
		func() ProtoMessage { return new(SceneCreateEntityReq) },
		func() ProtoMessage { return new(SceneCreateEntityRsp) },
		func() ProtoMessage { return new(SceneDestroyEntityReq) },
		func() ProtoMessage { return new(SceneDestroyEntityRsp) },
		func() ProtoMessage { return new(SceneForceUnlockNotify) },
		func() ProtoMessage { return new(SceneForceLockNotify) },
		func() ProtoMessage { return new(EnterWorldAreaReq) },
		func() ProtoMessage { return new(EnterWorldAreaRsp) },
		func() ProtoMessage { return new(EntityForceSyncReq) },
		func() ProtoMessage { return new(EntityForceSyncRsp) },
		func() ProtoMessage { return new(GetAreaExplorePointReq) },
		func() ProtoMessage { return new(GetAreaExplorePointRsp) },
		func() ProtoMessage { return new(ClientTransmitReq) },
		func() ProtoMessage { return new(ClientTransmitRsp) },
		func() ProtoMessage { return new(EnterSceneWeatherAreaNotify) },
		func() ProtoMessage { return new(ExitSceneWeatherAreaNotify) },
		func() ProtoMessage { return new(SceneAreaWeatherNotify) },
		func() ProtoMessage { return new(ScenePlayerInfoNotify) },
		func() ProtoMessage { return new(WorldPlayerLocationNotify) },
		func() ProtoMessage { return new(BeginCameraSceneLookNotify) },
		func() ProtoMessage { return new(EndCameraSceneLookNotify) },
		func() ProtoMessage { return new(MarkEntityInMinMapNotify) },
		func() ProtoMessage { return new(UnmarkEntityInMinMapNotify) },
		func() ProtoMessage { return new(ExecuteGroupTriggerReq) },
		func() ProtoMessage { return new(ExecuteGroupTriggerRsp) },
		func() ProtoMessage { return new(LevelupCityReq) },
		func() ProtoMessage { return new(LevelupCityRsp) },
		func() ProtoMessage { return new(SceneRouteChangeNotify) },
		func() ProtoMessage { return new(PlatformStartRouteNotify) },
		func() ProtoMessage { return new(PlatformStopRouteNotify) },
		func() ProtoMessage { return new(PlatformChangeRouteNotify) },
		func() ProtoMessage { return new(ScenePlayerSoundNotify) },
		func() ProtoMessage { return new(PersonalSceneJumpReq) },
		func() ProtoMessage { return new(PersonalSceneJumpRsp) },
		func() ProtoMessage { return new(SealBattleBeginNotify) },
		func() ProtoMessage { return new(SealBattleEndNotify) },
		func() ProtoMessage { return new(SealBattleProgressNotify) },
		func() ProtoMessage { return new(ClientPauseNotify) },
		func() ProtoMessage { return new(PlayerEnterSceneInfoNotify) },
		func() ProtoMessage { return new(JoinPlayerSceneReq) },
		func() ProtoMessage { return new(JoinPlayerSceneRsp) },
		func() ProtoMessage { return new(SceneKickPlayerReq) },
		func() ProtoMessage { return new(SceneKickPlayerRsp) },
		func() ProtoMessage { return new(SceneKickPlayerNotify) },
		func() ProtoMessage { return new(HitClientTrivialNotify) },
		func() ProtoMessage { return new(BackMyWorldReq) },
		func() ProtoMessage { return new(BackMyWorldRsp) },
		func() ProtoMessage { return new(SeeMonsterReq) },
		func() ProtoMessage { return new(SeeMonsterRsp) },
		func() ProtoMessage { return new(AddSeenMonsterNotify) },
		func() ProtoMessage { return new(AllSeenMonsterNotify) },
		func() ProtoMessage { return new(SceneTimeNotify) },
		func() ProtoMessage { return new(EnterSceneReadyReq) },
		func() ProtoMessage { return new(EnterSceneReadyRsp) },
		func() ProtoMessage { return new(EnterScenePeerNotify) },
		func() ProtoMessage { return new(EnterSceneDoneReq) },
		func() ProtoMessage { return new(EnterSceneDoneRsp) },
		func() ProtoMessage { return new(WorldPlayerDieNotify) },
		func() ProtoMessage { return new(WorldPlayerReviveReq) },
		func() ProtoMessage { return new(WorldPlayerReviveRsp) },
		func() ProtoMessage { return new(JoinPlayerFailNotify) },
		func() ProtoMessage { return new(SetSceneWeatherAreaReq) },
		func() ProtoMessage { return new(SetSceneWeatherAreaRsp) },
		func() ProtoMessage { return new(ExecuteGadgetLuaReq) },
		func() ProtoMessage { return new(ExecuteGadgetLuaRsp) },
		func() ProtoMessage { return new(CutSceneBeginNotify) },
		func() ProtoMessage { return new(CutSceneFinishNotify) },
		func() ProtoMessage { return new(CutSceneEndNotify) },
		func() ProtoMessage { return new(ClientScriptEventNotify) },
		func() ProtoMessage { return new(SceneEntitiesMovesReq) },
		func() ProtoMessage { return new(SceneEntitiesMovesRsp) },
		func() ProtoMessage { return new(SceneEntitiesMoveCombineNotify) },
		func() ProtoMessage { return new(UnlockTransPointReq) },
		func() ProtoMessage { return new(UnlockTransPointRsp) },
		func() ProtoMessage { return new(SceneWeatherForcastReq) },
		func() ProtoMessage { return new(SceneWeatherForcastRsp) },
		func() ProtoMessage { return new(MarkMapReq) },
		func() ProtoMessage { return new(MarkMapRsp) },
		func() ProtoMessage { return new(AllMarkPointNotify) },
		func() ProtoMessage { return new(WorldDataNotify) },
		func() ProtoMessage { return new(EntityMoveRoomNotify) },
		func() ProtoMessage { return new(WorldPlayerInfoNotify) },
		func() ProtoMessage { return new(PostEnterSceneReq) },
		func() ProtoMessage { return new(PostEnterSceneRsp) },
		func() ProtoMessage { return new(PlayerChatReq) },
		func() ProtoMessage { return new(PlayerChatRsp) },
		func() ProtoMessage { return new(PlayerChatNotify) },
		func() ProtoMessage { return new(PlayerChatCDNotify) },
		func() ProtoMessage { return new(ChatHistoryNotify) },
		func() ProtoMessage { return new(SceneDataNotify) },
		func() ProtoMessage { return new(DungeonEntryToBeExploreNotify) },
		func() ProtoMessage { return new(GetDungeonEntryExploreConditionReq) },
		func() ProtoMessage { return new(GetDungeonEntryExploreConditionRsp) },
		func() ProtoMessage { return new(UnfreezeGroupLimitNotify) },
		func() ProtoMessage { return new(SetEntityClientDataNotify) },
		func() ProtoMessage { return new(GroupSuiteNotify) },
		func() ProtoMessage { return new(GroupUnloadNotify) },
		func() ProtoMessage { return new(MonsterAIConfigHashNotify) },
		func() ProtoMessage { return new(ShowTemplateReminderNotify) },
		func() ProtoMessage { return new(ShowCommonTipsNotify) },
		func() ProtoMessage { return new(CloseCommonTipsNotify) },
		func() ProtoMessage { return new(ChangeWorldToSingleModeNotify) },
		func() ProtoMessage { return new(SyncScenePlayTeamEntityNotify) },
		func() ProtoMessage { return new(DelScenePlayTeamEntityNotify) },
		func() ProtoMessage { return new(PlayerEyePointStateNotify) },
		func() ProtoMessage { return new(GetMapMarkTipsReq) },
		func() ProtoMessage { return new(GetMapMarkTipsRsp) },
		func() ProtoMessage { return new(ChangeWorldToSingleModeReq) },
		func() ProtoMessage { return new(ChangeWorldToSingleModeRsp) },
		func() ProtoMessage { return new(GetWorldMpInfoReq) },
		func() ProtoMessage { return new(GetWorldMpInfoRsp) },
		func() ProtoMessage { return new(EntityConfigHashNotify) },
		func() ProtoMessage { return new(ForceDragAvatarNotify) },
		func() ProtoMessage { return new(MonsterPointArrayRouteUpdateNotify) },
		func() ProtoMessage { return new(ForceDragBackTransferNotify) },
		func() ProtoMessage { return new(GetScenePerformanceReq) },
		func() ProtoMessage { return new(GetScenePerformanceRsp) },
		func() ProtoMessage { return new(SceneAudioNotify) },
		func() ProtoMessage { return new(HitTreeNotify) },
		func() ProtoMessage { return new(EntityTagChangeNotify) },
		func() ProtoMessage { return new(AvatarFollowRouteNotify) },
		func() ProtoMessage { return new(SceneEntityUpdateNotify) },
		func() ProtoMessage { return new(ClientHashDebugNotify) },
		func() ProtoMessage { return new(PlayerWorldSceneInfoListNotify) },
		func() ProtoMessage { return new(LuaEnvironmentEffectNotify) },
		func() ProtoMessage { return new(ClientLoadingCostumeVerificationNotify) },
		func() ProtoMessage { return new(ShowClientGuideNotify) },
		func() ProtoMessage { return new(ShowClientTutorialNotify) },
		func() ProtoMessage { return new(GetMapAreaReq) },
		func() ProtoMessage { return new(GetMapAreaRsp) },
		func() ProtoMessage { return new(MapAreaChangeNotify) },
		func() ProtoMessage { return new(LeaveWorldNotify) },
		func() ProtoMessage { return new(GuestBeginEnterSceneNotify) },
		func() ProtoMessage { return new(GuestPostEnterSceneNotify) },
		func() ProtoMessage { return new(LevelTagDataNotify) },
		func() ProtoMessage { return new(StopReminderNotify) },
		func() ProtoMessage { return new(AreaPlayInfoNotify) },
		func() ProtoMessage { return new(CheckGroupReplacedReq) },
		func() ProtoMessage { return new(CheckGroupReplacedRsp) },
		func() ProtoMessage { return new(DeathZoneObserveNotify) },
		func() ProtoMessage { return new(WorldChestOpenNotify) },
		func() ProtoMessage { return new(WidgetQuickHitTreeReq) },
		func() ProtoMessage { return new(WidgetQuickHitTreeRsp) },
		func() ProtoMessage { return new(BeginCameraSceneLookWithTemplateNotify) },
		func() ProtoMessage { return new(RefreshEntityAuthNotify) },
		func() ProtoMessage { return new(ScenePlayerBackgroundAvatarRefreshNotify) },
	)
}

const (
	ProtoCommandPlayerEnterSceneNotify                   ProtoCommand = 272
	ProtoCommandLeaveSceneReq                            ProtoCommand = 298
	ProtoCommandLeaveSceneRsp                            ProtoCommand = 212
	ProtoCommandSceneInitFinishReq                       ProtoCommand = 235
	ProtoCommandSceneInitFinishRsp                       ProtoCommand = 207
	ProtoCommandSceneEntityAppearNotify                  ProtoCommand = 221
	ProtoCommandSceneEntityDisappearNotify               ProtoCommand = 203
	ProtoCommandSceneEntityMoveReq                       ProtoCommand = 290
	ProtoCommandSceneEntityMoveRsp                       ProtoCommand = 273
	ProtoCommandSceneAvatarStaminaStepReq                ProtoCommand = 299
	ProtoCommandSceneAvatarStaminaStepRsp                ProtoCommand = 231
	ProtoCommandSceneEntityMoveNotify                    ProtoCommand = 275
	ProtoCommandScenePlayerLocationNotify                ProtoCommand = 248
	ProtoCommandGetScenePointReq                         ProtoCommand = 297
	ProtoCommandGetScenePointRsp                         ProtoCommand = 281
	ProtoCommandEnterTransPointRegionNotify              ProtoCommand = 205
	ProtoCommandExitTransPointRegionNotify               ProtoCommand = 282
	ProtoCommandScenePointUnlockNotify                   ProtoCommand = 247
	ProtoCommandSceneTransToPointReq                     ProtoCommand = 239
	ProtoCommandSceneTransToPointRsp                     ProtoCommand = 253
	ProtoCommandEntityJumpNotify                         ProtoCommand = 222
	ProtoCommandGetSceneAreaReq                          ProtoCommand = 265
	ProtoCommandGetSceneAreaRsp                          ProtoCommand = 204
	ProtoCommandSceneAreaUnlockNotify                    ProtoCommand = 293
	ProtoCommandSceneEntityDrownReq                      ProtoCommand = 227
	ProtoCommandSceneEntityDrownRsp                      ProtoCommand = 294
	ProtoCommandSceneCreateEntityReq                     ProtoCommand = 288
	ProtoCommandSceneCreateEntityRsp                     ProtoCommand = 226
	ProtoCommandSceneDestroyEntityReq                    ProtoCommand = 263
	ProtoCommandSceneDestroyEntityRsp                    ProtoCommand = 295
	ProtoCommandSceneForceUnlockNotify                   ProtoCommand = 206
	ProtoCommandSceneForceLockNotify                     ProtoCommand = 234
	ProtoCommandEnterWorldAreaReq                        ProtoCommand = 250
	ProtoCommandEnterWorldAreaRsp                        ProtoCommand = 243
	ProtoCommandEntityForceSyncReq                       ProtoCommand = 274
	ProtoCommandEntityForceSyncRsp                       ProtoCommand = 276
	ProtoCommandGetAreaExplorePointReq                   ProtoCommand = 241
	ProtoCommandGetAreaExplorePointRsp                   ProtoCommand = 249
	ProtoCommandClientTransmitReq                        ProtoCommand = 291
	ProtoCommandClientTransmitRsp                        ProtoCommand = 224
	ProtoCommandEnterSceneWeatherAreaNotify              ProtoCommand = 256
	ProtoCommandExitSceneWeatherAreaNotify               ProtoCommand = 242
	ProtoCommandSceneAreaWeatherNotify                   ProtoCommand = 230
	ProtoCommandScenePlayerInfoNotify                    ProtoCommand = 267
	ProtoCommandWorldPlayerLocationNotify                ProtoCommand = 258
	ProtoCommandBeginCameraSceneLookNotify               ProtoCommand = 270
	ProtoCommandEndCameraSceneLookNotify                 ProtoCommand = 217
	ProtoCommandMarkEntityInMinMapNotify                 ProtoCommand = 202
	ProtoCommandUnmarkEntityInMinMapNotify               ProtoCommand = 219
	ProtoCommandExecuteGroupTriggerReq                   ProtoCommand = 257
	ProtoCommandExecuteGroupTriggerRsp                   ProtoCommand = 300
	ProtoCommandLevelupCityReq                           ProtoCommand = 216
	ProtoCommandLevelupCityRsp                           ProtoCommand = 287
	ProtoCommandSceneRouteChangeNotify                   ProtoCommand = 240
	ProtoCommandPlatformStartRouteNotify                 ProtoCommand = 218
	ProtoCommandPlatformStopRouteNotify                  ProtoCommand = 266
	ProtoCommandPlatformChangeRouteNotify                ProtoCommand = 268
	ProtoCommandScenePlayerSoundNotify                   ProtoCommand = 233
	ProtoCommandPersonalSceneJumpReq                     ProtoCommand = 284
	ProtoCommandPersonalSceneJumpRsp                     ProtoCommand = 280
	ProtoCommandSealBattleBeginNotify                    ProtoCommand = 289
	ProtoCommandSealBattleEndNotify                      ProtoCommand = 259
	ProtoCommandSealBattleProgressNotify                 ProtoCommand = 232
	ProtoCommandClientPauseNotify                        ProtoCommand = 260
	ProtoCommandPlayerEnterSceneInfoNotify               ProtoCommand = 214
	ProtoCommandJoinPlayerSceneReq                       ProtoCommand = 292
	ProtoCommandJoinPlayerSceneRsp                       ProtoCommand = 220
	ProtoCommandSceneKickPlayerReq                       ProtoCommand = 264
	ProtoCommandSceneKickPlayerRsp                       ProtoCommand = 238
	ProtoCommandSceneKickPlayerNotify                    ProtoCommand = 211
	ProtoCommandHitClientTrivialNotify                   ProtoCommand = 244
	ProtoCommandBackMyWorldReq                           ProtoCommand = 286
	ProtoCommandBackMyWorldRsp                           ProtoCommand = 201
	ProtoCommandSeeMonsterReq                            ProtoCommand = 228
	ProtoCommandSeeMonsterRsp                            ProtoCommand = 251
	ProtoCommandAddSeenMonsterNotify                     ProtoCommand = 223
	ProtoCommandAllSeenMonsterNotify                     ProtoCommand = 271
	ProtoCommandSceneTimeNotify                          ProtoCommand = 245
	ProtoCommandEnterSceneReadyReq                       ProtoCommand = 208
	ProtoCommandEnterSceneReadyRsp                       ProtoCommand = 209
	ProtoCommandEnterScenePeerNotify                     ProtoCommand = 252
	ProtoCommandEnterSceneDoneReq                        ProtoCommand = 277
	ProtoCommandEnterSceneDoneRsp                        ProtoCommand = 237
	ProtoCommandWorldPlayerDieNotify                     ProtoCommand = 285
	ProtoCommandWorldPlayerReviveReq                     ProtoCommand = 225
	ProtoCommandWorldPlayerReviveRsp                     ProtoCommand = 278
	ProtoCommandJoinPlayerFailNotify                     ProtoCommand = 236
	ProtoCommandSetSceneWeatherAreaReq                   ProtoCommand = 254
	ProtoCommandSetSceneWeatherAreaRsp                   ProtoCommand = 283
	ProtoCommandExecuteGadgetLuaReq                      ProtoCommand = 269
	ProtoCommandExecuteGadgetLuaRsp                      ProtoCommand = 210
	ProtoCommandCutSceneBeginNotify                      ProtoCommand = 296
	ProtoCommandCutSceneFinishNotify                     ProtoCommand = 262
	ProtoCommandCutSceneEndNotify                        ProtoCommand = 215
	ProtoCommandClientScriptEventNotify                  ProtoCommand = 213
	ProtoCommandSceneEntitiesMovesReq                    ProtoCommand = 279
	ProtoCommandSceneEntitiesMovesRsp                    ProtoCommand = 255
	ProtoCommandSceneEntitiesMoveCombineNotify           ProtoCommand = 3387
	ProtoCommandUnlockTransPointReq                      ProtoCommand = 3035
	ProtoCommandUnlockTransPointRsp                      ProtoCommand = 3426
	ProtoCommandSceneWeatherForcastReq                   ProtoCommand = 3110
	ProtoCommandSceneWeatherForcastRsp                   ProtoCommand = 3012
	ProtoCommandMarkMapReq                               ProtoCommand = 3466
	ProtoCommandMarkMapRsp                               ProtoCommand = 3079
	ProtoCommandAllMarkPointNotify                       ProtoCommand = 3283
	ProtoCommandWorldDataNotify                          ProtoCommand = 3308
	ProtoCommandEntityMoveRoomNotify                     ProtoCommand = 3178
	ProtoCommandWorldPlayerInfoNotify                    ProtoCommand = 3116
	ProtoCommandPostEnterSceneReq                        ProtoCommand = 3312
	ProtoCommandPostEnterSceneRsp                        ProtoCommand = 3184
	ProtoCommandPlayerChatReq                            ProtoCommand = 3185
	ProtoCommandPlayerChatRsp                            ProtoCommand = 3228
	ProtoCommandPlayerChatNotify                         ProtoCommand = 3010
	ProtoCommandPlayerChatCDNotify                       ProtoCommand = 3367
	ProtoCommandChatHistoryNotify                        ProtoCommand = 3496
	ProtoCommandSceneDataNotify                          ProtoCommand = 3203
	ProtoCommandDungeonEntryToBeExploreNotify            ProtoCommand = 3147
	ProtoCommandGetDungeonEntryExploreConditionReq       ProtoCommand = 3165
	ProtoCommandGetDungeonEntryExploreConditionRsp       ProtoCommand = 3269
	ProtoCommandUnfreezeGroupLimitNotify                 ProtoCommand = 3220
	ProtoCommandSetEntityClientDataNotify                ProtoCommand = 3146
	ProtoCommandGroupSuiteNotify                         ProtoCommand = 3257
	ProtoCommandGroupUnloadNotify                        ProtoCommand = 3344
	ProtoCommandMonsterAIConfigHashNotify                ProtoCommand = 3039
	ProtoCommandShowTemplateReminderNotify               ProtoCommand = 3491
	ProtoCommandShowCommonTipsNotify                     ProtoCommand = 3352
	ProtoCommandCloseCommonTipsNotify                    ProtoCommand = 3194
	ProtoCommandChangeWorldToSingleModeNotify            ProtoCommand = 3006
	ProtoCommandSyncScenePlayTeamEntityNotify            ProtoCommand = 3333
	ProtoCommandDelScenePlayTeamEntityNotify             ProtoCommand = 3318
	ProtoCommandPlayerEyePointStateNotify                ProtoCommand = 3051
	ProtoCommandGetMapMarkTipsReq                        ProtoCommand = 3463
	ProtoCommandGetMapMarkTipsRsp                        ProtoCommand = 3327
	ProtoCommandChangeWorldToSingleModeReq               ProtoCommand = 3066
	ProtoCommandChangeWorldToSingleModeRsp               ProtoCommand = 3282
	ProtoCommandGetWorldMpInfoReq                        ProtoCommand = 3391
	ProtoCommandGetWorldMpInfoRsp                        ProtoCommand = 3320
	ProtoCommandEntityConfigHashNotify                   ProtoCommand = 3189
	ProtoCommandForceDragAvatarNotify                    ProtoCommand = 3235
	ProtoCommandMonsterPointArrayRouteUpdateNotify       ProtoCommand = 3410
	ProtoCommandForceDragBackTransferNotify              ProtoCommand = 3145
	ProtoCommandGetScenePerformanceReq                   ProtoCommand = 3419
	ProtoCommandGetScenePerformanceRsp                   ProtoCommand = 3137
	ProtoCommandSceneAudioNotify                         ProtoCommand = 3166
	ProtoCommandHitTreeNotify                            ProtoCommand = 3019
	ProtoCommandEntityTagChangeNotify                    ProtoCommand = 3316
	ProtoCommandAvatarFollowRouteNotify                  ProtoCommand = 3458
	ProtoCommandSceneEntityUpdateNotify                  ProtoCommand = 3412
	ProtoCommandClientHashDebugNotify                    ProtoCommand = 3086
	ProtoCommandPlayerWorldSceneInfoListNotify           ProtoCommand = 3129
	ProtoCommandLuaEnvironmentEffectNotify               ProtoCommand = 3408
	ProtoCommandClientLoadingCostumeVerificationNotify   ProtoCommand = 3487
	ProtoCommandShowClientGuideNotify                    ProtoCommand = 3005
	ProtoCommandShowClientTutorialNotify                 ProtoCommand = 3305
	ProtoCommandGetMapAreaReq                            ProtoCommand = 3108
	ProtoCommandGetMapAreaRsp                            ProtoCommand = 3328
	ProtoCommandMapAreaChangeNotify                      ProtoCommand = 3378
	ProtoCommandLeaveWorldNotify                         ProtoCommand = 3017
	ProtoCommandGuestBeginEnterSceneNotify               ProtoCommand = 3031
	ProtoCommandGuestPostEnterSceneNotify                ProtoCommand = 3144
	ProtoCommandLevelTagDataNotify                       ProtoCommand = 3314
	ProtoCommandStopReminderNotify                       ProtoCommand = 3004
	ProtoCommandAreaPlayInfoNotify                       ProtoCommand = 3323
	ProtoCommandCheckGroupReplacedReq                    ProtoCommand = 3113
	ProtoCommandCheckGroupReplacedRsp                    ProtoCommand = 3152
	ProtoCommandDeathZoneObserveNotify                   ProtoCommand = 3475
	ProtoCommandWorldChestOpenNotify                     ProtoCommand = 3295
	ProtoCommandWidgetQuickHitTreeReq                    ProtoCommand = 3345
	ProtoCommandWidgetQuickHitTreeRsp                    ProtoCommand = 3336
	ProtoCommandBeginCameraSceneLookWithTemplateNotify   ProtoCommand = 3160
	ProtoCommandRefreshEntityAuthNotify                  ProtoCommand = 3259
	ProtoCommandScenePlayerBackgroundAvatarRefreshNotify ProtoCommand = 3274
)

func (*PlayerEnterSceneNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerEnterSceneNotify }
func (*PlayerEnterSceneNotify) ProtoMessageType() ProtoMessageType { return "PlayerEnterSceneNotify" }

func (*LeaveSceneReq) ProtoCommand() ProtoCommand         { return ProtoCommandLeaveSceneReq }
func (*LeaveSceneReq) ProtoMessageType() ProtoMessageType { return "LeaveSceneReq" }

func (*LeaveSceneRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLeaveSceneRsp }
func (*LeaveSceneRsp) ProtoMessageType() ProtoMessageType { return "LeaveSceneRsp" }

func (*SceneInitFinishReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneInitFinishReq }
func (*SceneInitFinishReq) ProtoMessageType() ProtoMessageType { return "SceneInitFinishReq" }

func (*SceneInitFinishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneInitFinishRsp }
func (*SceneInitFinishRsp) ProtoMessageType() ProtoMessageType { return "SceneInitFinishRsp" }

func (*SceneEntityAppearNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneEntityAppearNotify
}
func (*SceneEntityAppearNotify) ProtoMessageType() ProtoMessageType { return "SceneEntityAppearNotify" }

func (*SceneEntityDisappearNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneEntityDisappearNotify
}
func (*SceneEntityDisappearNotify) ProtoMessageType() ProtoMessageType {
	return "SceneEntityDisappearNotify"
}

func (*SceneEntityMoveReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntityMoveReq }
func (*SceneEntityMoveReq) ProtoMessageType() ProtoMessageType { return "SceneEntityMoveReq" }

func (*SceneEntityMoveRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntityMoveRsp }
func (*SceneEntityMoveRsp) ProtoMessageType() ProtoMessageType { return "SceneEntityMoveRsp" }

func (*SceneAvatarStaminaStepReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneAvatarStaminaStepReq
}
func (*SceneAvatarStaminaStepReq) ProtoMessageType() ProtoMessageType {
	return "SceneAvatarStaminaStepReq"
}

func (*SceneAvatarStaminaStepRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneAvatarStaminaStepRsp
}
func (*SceneAvatarStaminaStepRsp) ProtoMessageType() ProtoMessageType {
	return "SceneAvatarStaminaStepRsp"
}

func (*SceneEntityMoveNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntityMoveNotify }
func (*SceneEntityMoveNotify) ProtoMessageType() ProtoMessageType { return "SceneEntityMoveNotify" }

func (*ScenePlayerLocationNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayerLocationNotify
}
func (*ScenePlayerLocationNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayerLocationNotify"
}

func (*GetScenePointReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetScenePointReq }
func (*GetScenePointReq) ProtoMessageType() ProtoMessageType { return "GetScenePointReq" }

func (*GetScenePointRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetScenePointRsp }
func (*GetScenePointRsp) ProtoMessageType() ProtoMessageType { return "GetScenePointRsp" }

func (*EnterTransPointRegionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterTransPointRegionNotify
}
func (*EnterTransPointRegionNotify) ProtoMessageType() ProtoMessageType {
	return "EnterTransPointRegionNotify"
}

func (*ExitTransPointRegionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandExitTransPointRegionNotify
}
func (*ExitTransPointRegionNotify) ProtoMessageType() ProtoMessageType {
	return "ExitTransPointRegionNotify"
}

func (*ScenePointUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandScenePointUnlockNotify }
func (*ScenePointUnlockNotify) ProtoMessageType() ProtoMessageType { return "ScenePointUnlockNotify" }

func (*SceneTransToPointReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneTransToPointReq }
func (*SceneTransToPointReq) ProtoMessageType() ProtoMessageType { return "SceneTransToPointReq" }

func (*SceneTransToPointRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneTransToPointRsp }
func (*SceneTransToPointRsp) ProtoMessageType() ProtoMessageType { return "SceneTransToPointRsp" }

func (*EntityJumpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityJumpNotify }
func (*EntityJumpNotify) ProtoMessageType() ProtoMessageType { return "EntityJumpNotify" }

func (*GetSceneAreaReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetSceneAreaReq }
func (*GetSceneAreaReq) ProtoMessageType() ProtoMessageType { return "GetSceneAreaReq" }

func (*GetSceneAreaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetSceneAreaRsp }
func (*GetSceneAreaRsp) ProtoMessageType() ProtoMessageType { return "GetSceneAreaRsp" }

func (*SceneAreaUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneAreaUnlockNotify }
func (*SceneAreaUnlockNotify) ProtoMessageType() ProtoMessageType { return "SceneAreaUnlockNotify" }

func (*SceneEntityDrownReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntityDrownReq }
func (*SceneEntityDrownReq) ProtoMessageType() ProtoMessageType { return "SceneEntityDrownReq" }

func (*SceneEntityDrownRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntityDrownRsp }
func (*SceneEntityDrownRsp) ProtoMessageType() ProtoMessageType { return "SceneEntityDrownRsp" }

func (*SceneCreateEntityReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneCreateEntityReq }
func (*SceneCreateEntityReq) ProtoMessageType() ProtoMessageType { return "SceneCreateEntityReq" }

func (*SceneCreateEntityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneCreateEntityRsp }
func (*SceneCreateEntityRsp) ProtoMessageType() ProtoMessageType { return "SceneCreateEntityRsp" }

func (*SceneDestroyEntityReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneDestroyEntityReq }
func (*SceneDestroyEntityReq) ProtoMessageType() ProtoMessageType { return "SceneDestroyEntityReq" }

func (*SceneDestroyEntityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneDestroyEntityRsp }
func (*SceneDestroyEntityRsp) ProtoMessageType() ProtoMessageType { return "SceneDestroyEntityRsp" }

func (*SceneForceUnlockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneForceUnlockNotify }
func (*SceneForceUnlockNotify) ProtoMessageType() ProtoMessageType { return "SceneForceUnlockNotify" }

func (*SceneForceLockNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneForceLockNotify }
func (*SceneForceLockNotify) ProtoMessageType() ProtoMessageType { return "SceneForceLockNotify" }

func (*EnterWorldAreaReq) ProtoCommand() ProtoCommand         { return ProtoCommandEnterWorldAreaReq }
func (*EnterWorldAreaReq) ProtoMessageType() ProtoMessageType { return "EnterWorldAreaReq" }

func (*EnterWorldAreaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEnterWorldAreaRsp }
func (*EnterWorldAreaRsp) ProtoMessageType() ProtoMessageType { return "EnterWorldAreaRsp" }

func (*EntityForceSyncReq) ProtoCommand() ProtoCommand         { return ProtoCommandEntityForceSyncReq }
func (*EntityForceSyncReq) ProtoMessageType() ProtoMessageType { return "EntityForceSyncReq" }

func (*EntityForceSyncRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEntityForceSyncRsp }
func (*EntityForceSyncRsp) ProtoMessageType() ProtoMessageType { return "EntityForceSyncRsp" }

func (*GetAreaExplorePointReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetAreaExplorePointReq }
func (*GetAreaExplorePointReq) ProtoMessageType() ProtoMessageType { return "GetAreaExplorePointReq" }

func (*GetAreaExplorePointRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetAreaExplorePointRsp }
func (*GetAreaExplorePointRsp) ProtoMessageType() ProtoMessageType { return "GetAreaExplorePointRsp" }

func (*ClientTransmitReq) ProtoCommand() ProtoCommand         { return ProtoCommandClientTransmitReq }
func (*ClientTransmitReq) ProtoMessageType() ProtoMessageType { return "ClientTransmitReq" }

func (*ClientTransmitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandClientTransmitRsp }
func (*ClientTransmitRsp) ProtoMessageType() ProtoMessageType { return "ClientTransmitRsp" }

func (*EnterSceneWeatherAreaNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterSceneWeatherAreaNotify
}
func (*EnterSceneWeatherAreaNotify) ProtoMessageType() ProtoMessageType {
	return "EnterSceneWeatherAreaNotify"
}

func (*ExitSceneWeatherAreaNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandExitSceneWeatherAreaNotify
}
func (*ExitSceneWeatherAreaNotify) ProtoMessageType() ProtoMessageType {
	return "ExitSceneWeatherAreaNotify"
}

func (*SceneAreaWeatherNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneAreaWeatherNotify }
func (*SceneAreaWeatherNotify) ProtoMessageType() ProtoMessageType { return "SceneAreaWeatherNotify" }

func (*ScenePlayerInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandScenePlayerInfoNotify }
func (*ScenePlayerInfoNotify) ProtoMessageType() ProtoMessageType { return "ScenePlayerInfoNotify" }

func (*WorldPlayerLocationNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldPlayerLocationNotify
}
func (*WorldPlayerLocationNotify) ProtoMessageType() ProtoMessageType {
	return "WorldPlayerLocationNotify"
}

func (*BeginCameraSceneLookNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBeginCameraSceneLookNotify
}
func (*BeginCameraSceneLookNotify) ProtoMessageType() ProtoMessageType {
	return "BeginCameraSceneLookNotify"
}

func (*EndCameraSceneLookNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEndCameraSceneLookNotify
}
func (*EndCameraSceneLookNotify) ProtoMessageType() ProtoMessageType {
	return "EndCameraSceneLookNotify"
}

func (*MarkEntityInMinMapNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMarkEntityInMinMapNotify
}
func (*MarkEntityInMinMapNotify) ProtoMessageType() ProtoMessageType {
	return "MarkEntityInMinMapNotify"
}

func (*UnmarkEntityInMinMapNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUnmarkEntityInMinMapNotify
}
func (*UnmarkEntityInMinMapNotify) ProtoMessageType() ProtoMessageType {
	return "UnmarkEntityInMinMapNotify"
}

func (*ExecuteGroupTriggerReq) ProtoCommand() ProtoCommand         { return ProtoCommandExecuteGroupTriggerReq }
func (*ExecuteGroupTriggerReq) ProtoMessageType() ProtoMessageType { return "ExecuteGroupTriggerReq" }

func (*ExecuteGroupTriggerRsp) ProtoCommand() ProtoCommand         { return ProtoCommandExecuteGroupTriggerRsp }
func (*ExecuteGroupTriggerRsp) ProtoMessageType() ProtoMessageType { return "ExecuteGroupTriggerRsp" }

func (*LevelupCityReq) ProtoCommand() ProtoCommand         { return ProtoCommandLevelupCityReq }
func (*LevelupCityReq) ProtoMessageType() ProtoMessageType { return "LevelupCityReq" }

func (*LevelupCityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLevelupCityRsp }
func (*LevelupCityRsp) ProtoMessageType() ProtoMessageType { return "LevelupCityRsp" }

func (*SceneRouteChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneRouteChangeNotify }
func (*SceneRouteChangeNotify) ProtoMessageType() ProtoMessageType { return "SceneRouteChangeNotify" }

func (*PlatformStartRouteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlatformStartRouteNotify
}
func (*PlatformStartRouteNotify) ProtoMessageType() ProtoMessageType {
	return "PlatformStartRouteNotify"
}

func (*PlatformStopRouteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlatformStopRouteNotify
}
func (*PlatformStopRouteNotify) ProtoMessageType() ProtoMessageType { return "PlatformStopRouteNotify" }

func (*PlatformChangeRouteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlatformChangeRouteNotify
}
func (*PlatformChangeRouteNotify) ProtoMessageType() ProtoMessageType {
	return "PlatformChangeRouteNotify"
}

func (*ScenePlayerSoundNotify) ProtoCommand() ProtoCommand         { return ProtoCommandScenePlayerSoundNotify }
func (*ScenePlayerSoundNotify) ProtoMessageType() ProtoMessageType { return "ScenePlayerSoundNotify" }

func (*PersonalSceneJumpReq) ProtoCommand() ProtoCommand         { return ProtoCommandPersonalSceneJumpReq }
func (*PersonalSceneJumpReq) ProtoMessageType() ProtoMessageType { return "PersonalSceneJumpReq" }

func (*PersonalSceneJumpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPersonalSceneJumpRsp }
func (*PersonalSceneJumpRsp) ProtoMessageType() ProtoMessageType { return "PersonalSceneJumpRsp" }

func (*SealBattleBeginNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSealBattleBeginNotify }
func (*SealBattleBeginNotify) ProtoMessageType() ProtoMessageType { return "SealBattleBeginNotify" }

func (*SealBattleEndNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSealBattleEndNotify }
func (*SealBattleEndNotify) ProtoMessageType() ProtoMessageType { return "SealBattleEndNotify" }

func (*SealBattleProgressNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSealBattleProgressNotify
}
func (*SealBattleProgressNotify) ProtoMessageType() ProtoMessageType {
	return "SealBattleProgressNotify"
}

func (*ClientPauseNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClientPauseNotify }
func (*ClientPauseNotify) ProtoMessageType() ProtoMessageType { return "ClientPauseNotify" }

func (*PlayerEnterSceneInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerEnterSceneInfoNotify
}
func (*PlayerEnterSceneInfoNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerEnterSceneInfoNotify"
}

func (*JoinPlayerSceneReq) ProtoCommand() ProtoCommand         { return ProtoCommandJoinPlayerSceneReq }
func (*JoinPlayerSceneReq) ProtoMessageType() ProtoMessageType { return "JoinPlayerSceneReq" }

func (*JoinPlayerSceneRsp) ProtoCommand() ProtoCommand         { return ProtoCommandJoinPlayerSceneRsp }
func (*JoinPlayerSceneRsp) ProtoMessageType() ProtoMessageType { return "JoinPlayerSceneRsp" }

func (*SceneKickPlayerReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneKickPlayerReq }
func (*SceneKickPlayerReq) ProtoMessageType() ProtoMessageType { return "SceneKickPlayerReq" }

func (*SceneKickPlayerRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneKickPlayerRsp }
func (*SceneKickPlayerRsp) ProtoMessageType() ProtoMessageType { return "SceneKickPlayerRsp" }

func (*SceneKickPlayerNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneKickPlayerNotify }
func (*SceneKickPlayerNotify) ProtoMessageType() ProtoMessageType { return "SceneKickPlayerNotify" }

func (*HitClientTrivialNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHitClientTrivialNotify }
func (*HitClientTrivialNotify) ProtoMessageType() ProtoMessageType { return "HitClientTrivialNotify" }

func (*BackMyWorldReq) ProtoCommand() ProtoCommand         { return ProtoCommandBackMyWorldReq }
func (*BackMyWorldReq) ProtoMessageType() ProtoMessageType { return "BackMyWorldReq" }

func (*BackMyWorldRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBackMyWorldRsp }
func (*BackMyWorldRsp) ProtoMessageType() ProtoMessageType { return "BackMyWorldRsp" }

func (*SeeMonsterReq) ProtoCommand() ProtoCommand         { return ProtoCommandSeeMonsterReq }
func (*SeeMonsterReq) ProtoMessageType() ProtoMessageType { return "SeeMonsterReq" }

func (*SeeMonsterRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSeeMonsterRsp }
func (*SeeMonsterRsp) ProtoMessageType() ProtoMessageType { return "SeeMonsterRsp" }

func (*AddSeenMonsterNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAddSeenMonsterNotify }
func (*AddSeenMonsterNotify) ProtoMessageType() ProtoMessageType { return "AddSeenMonsterNotify" }

func (*AllSeenMonsterNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAllSeenMonsterNotify }
func (*AllSeenMonsterNotify) ProtoMessageType() ProtoMessageType { return "AllSeenMonsterNotify" }

func (*SceneTimeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneTimeNotify }
func (*SceneTimeNotify) ProtoMessageType() ProtoMessageType { return "SceneTimeNotify" }

func (*EnterSceneReadyReq) ProtoCommand() ProtoCommand         { return ProtoCommandEnterSceneReadyReq }
func (*EnterSceneReadyReq) ProtoMessageType() ProtoMessageType { return "EnterSceneReadyReq" }

func (*EnterSceneReadyRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEnterSceneReadyRsp }
func (*EnterSceneReadyRsp) ProtoMessageType() ProtoMessageType { return "EnterSceneReadyRsp" }

func (*EnterScenePeerNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEnterScenePeerNotify }
func (*EnterScenePeerNotify) ProtoMessageType() ProtoMessageType { return "EnterScenePeerNotify" }

func (*EnterSceneDoneReq) ProtoCommand() ProtoCommand         { return ProtoCommandEnterSceneDoneReq }
func (*EnterSceneDoneReq) ProtoMessageType() ProtoMessageType { return "EnterSceneDoneReq" }

func (*EnterSceneDoneRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEnterSceneDoneRsp }
func (*EnterSceneDoneRsp) ProtoMessageType() ProtoMessageType { return "EnterSceneDoneRsp" }

func (*WorldPlayerDieNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWorldPlayerDieNotify }
func (*WorldPlayerDieNotify) ProtoMessageType() ProtoMessageType { return "WorldPlayerDieNotify" }

func (*WorldPlayerReviveReq) ProtoCommand() ProtoCommand         { return ProtoCommandWorldPlayerReviveReq }
func (*WorldPlayerReviveReq) ProtoMessageType() ProtoMessageType { return "WorldPlayerReviveReq" }

func (*WorldPlayerReviveRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWorldPlayerReviveRsp }
func (*WorldPlayerReviveRsp) ProtoMessageType() ProtoMessageType { return "WorldPlayerReviveRsp" }

func (*JoinPlayerFailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandJoinPlayerFailNotify }
func (*JoinPlayerFailNotify) ProtoMessageType() ProtoMessageType { return "JoinPlayerFailNotify" }

func (*SetSceneWeatherAreaReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetSceneWeatherAreaReq }
func (*SetSceneWeatherAreaReq) ProtoMessageType() ProtoMessageType { return "SetSceneWeatherAreaReq" }

func (*SetSceneWeatherAreaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetSceneWeatherAreaRsp }
func (*SetSceneWeatherAreaRsp) ProtoMessageType() ProtoMessageType { return "SetSceneWeatherAreaRsp" }

func (*ExecuteGadgetLuaReq) ProtoCommand() ProtoCommand         { return ProtoCommandExecuteGadgetLuaReq }
func (*ExecuteGadgetLuaReq) ProtoMessageType() ProtoMessageType { return "ExecuteGadgetLuaReq" }

func (*ExecuteGadgetLuaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandExecuteGadgetLuaRsp }
func (*ExecuteGadgetLuaRsp) ProtoMessageType() ProtoMessageType { return "ExecuteGadgetLuaRsp" }

func (*CutSceneBeginNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCutSceneBeginNotify }
func (*CutSceneBeginNotify) ProtoMessageType() ProtoMessageType { return "CutSceneBeginNotify" }

func (*CutSceneFinishNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCutSceneFinishNotify }
func (*CutSceneFinishNotify) ProtoMessageType() ProtoMessageType { return "CutSceneFinishNotify" }

func (*CutSceneEndNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCutSceneEndNotify }
func (*CutSceneEndNotify) ProtoMessageType() ProtoMessageType { return "CutSceneEndNotify" }

func (*ClientScriptEventNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientScriptEventNotify
}
func (*ClientScriptEventNotify) ProtoMessageType() ProtoMessageType { return "ClientScriptEventNotify" }

func (*SceneEntitiesMovesReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntitiesMovesReq }
func (*SceneEntitiesMovesReq) ProtoMessageType() ProtoMessageType { return "SceneEntitiesMovesReq" }

func (*SceneEntitiesMovesRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneEntitiesMovesRsp }
func (*SceneEntitiesMovesRsp) ProtoMessageType() ProtoMessageType { return "SceneEntitiesMovesRsp" }

func (*SceneEntitiesMoveCombineNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneEntitiesMoveCombineNotify
}
func (*SceneEntitiesMoveCombineNotify) ProtoMessageType() ProtoMessageType {
	return "SceneEntitiesMoveCombineNotify"
}

func (*UnlockTransPointReq) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockTransPointReq }
func (*UnlockTransPointReq) ProtoMessageType() ProtoMessageType { return "UnlockTransPointReq" }

func (*UnlockTransPointRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockTransPointRsp }
func (*UnlockTransPointRsp) ProtoMessageType() ProtoMessageType { return "UnlockTransPointRsp" }

func (*SceneWeatherForcastReq) ProtoCommand() ProtoCommand         { return ProtoCommandSceneWeatherForcastReq }
func (*SceneWeatherForcastReq) ProtoMessageType() ProtoMessageType { return "SceneWeatherForcastReq" }

func (*SceneWeatherForcastRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSceneWeatherForcastRsp }
func (*SceneWeatherForcastRsp) ProtoMessageType() ProtoMessageType { return "SceneWeatherForcastRsp" }

func (*MarkMapReq) ProtoCommand() ProtoCommand         { return ProtoCommandMarkMapReq }
func (*MarkMapReq) ProtoMessageType() ProtoMessageType { return "MarkMapReq" }

func (*MarkMapRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMarkMapRsp }
func (*MarkMapRsp) ProtoMessageType() ProtoMessageType { return "MarkMapRsp" }

func (*AllMarkPointNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAllMarkPointNotify }
func (*AllMarkPointNotify) ProtoMessageType() ProtoMessageType { return "AllMarkPointNotify" }

func (*WorldDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWorldDataNotify }
func (*WorldDataNotify) ProtoMessageType() ProtoMessageType { return "WorldDataNotify" }

func (*EntityMoveRoomNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityMoveRoomNotify }
func (*EntityMoveRoomNotify) ProtoMessageType() ProtoMessageType { return "EntityMoveRoomNotify" }

func (*WorldPlayerInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWorldPlayerInfoNotify }
func (*WorldPlayerInfoNotify) ProtoMessageType() ProtoMessageType { return "WorldPlayerInfoNotify" }

func (*PostEnterSceneReq) ProtoCommand() ProtoCommand         { return ProtoCommandPostEnterSceneReq }
func (*PostEnterSceneReq) ProtoMessageType() ProtoMessageType { return "PostEnterSceneReq" }

func (*PostEnterSceneRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPostEnterSceneRsp }
func (*PostEnterSceneRsp) ProtoMessageType() ProtoMessageType { return "PostEnterSceneRsp" }

func (*PlayerChatReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerChatReq }
func (*PlayerChatReq) ProtoMessageType() ProtoMessageType { return "PlayerChatReq" }

func (*PlayerChatRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerChatRsp }
func (*PlayerChatRsp) ProtoMessageType() ProtoMessageType { return "PlayerChatRsp" }

func (*PlayerChatNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerChatNotify }
func (*PlayerChatNotify) ProtoMessageType() ProtoMessageType { return "PlayerChatNotify" }

func (*PlayerChatCDNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerChatCDNotify }
func (*PlayerChatCDNotify) ProtoMessageType() ProtoMessageType { return "PlayerChatCDNotify" }

func (*ChatHistoryNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChatHistoryNotify }
func (*ChatHistoryNotify) ProtoMessageType() ProtoMessageType { return "ChatHistoryNotify" }

func (*SceneDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneDataNotify }
func (*SceneDataNotify) ProtoMessageType() ProtoMessageType { return "SceneDataNotify" }

func (*DungeonEntryToBeExploreNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDungeonEntryToBeExploreNotify
}
func (*DungeonEntryToBeExploreNotify) ProtoMessageType() ProtoMessageType {
	return "DungeonEntryToBeExploreNotify"
}

func (*GetDungeonEntryExploreConditionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetDungeonEntryExploreConditionReq
}
func (*GetDungeonEntryExploreConditionReq) ProtoMessageType() ProtoMessageType {
	return "GetDungeonEntryExploreConditionReq"
}

func (*GetDungeonEntryExploreConditionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetDungeonEntryExploreConditionRsp
}
func (*GetDungeonEntryExploreConditionRsp) ProtoMessageType() ProtoMessageType {
	return "GetDungeonEntryExploreConditionRsp"
}

func (*UnfreezeGroupLimitNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUnfreezeGroupLimitNotify
}
func (*UnfreezeGroupLimitNotify) ProtoMessageType() ProtoMessageType {
	return "UnfreezeGroupLimitNotify"
}

func (*SetEntityClientDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSetEntityClientDataNotify
}
func (*SetEntityClientDataNotify) ProtoMessageType() ProtoMessageType {
	return "SetEntityClientDataNotify"
}

func (*GroupSuiteNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGroupSuiteNotify }
func (*GroupSuiteNotify) ProtoMessageType() ProtoMessageType { return "GroupSuiteNotify" }

func (*GroupUnloadNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGroupUnloadNotify }
func (*GroupUnloadNotify) ProtoMessageType() ProtoMessageType { return "GroupUnloadNotify" }

func (*MonsterAIConfigHashNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMonsterAIConfigHashNotify
}
func (*MonsterAIConfigHashNotify) ProtoMessageType() ProtoMessageType {
	return "MonsterAIConfigHashNotify"
}

func (*ShowTemplateReminderNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandShowTemplateReminderNotify
}
func (*ShowTemplateReminderNotify) ProtoMessageType() ProtoMessageType {
	return "ShowTemplateReminderNotify"
}

func (*ShowCommonTipsNotify) ProtoCommand() ProtoCommand         { return ProtoCommandShowCommonTipsNotify }
func (*ShowCommonTipsNotify) ProtoMessageType() ProtoMessageType { return "ShowCommonTipsNotify" }

func (*CloseCommonTipsNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCloseCommonTipsNotify }
func (*CloseCommonTipsNotify) ProtoMessageType() ProtoMessageType { return "CloseCommonTipsNotify" }

func (*ChangeWorldToSingleModeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeWorldToSingleModeNotify
}
func (*ChangeWorldToSingleModeNotify) ProtoMessageType() ProtoMessageType {
	return "ChangeWorldToSingleModeNotify"
}

func (*SyncScenePlayTeamEntityNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSyncScenePlayTeamEntityNotify
}
func (*SyncScenePlayTeamEntityNotify) ProtoMessageType() ProtoMessageType {
	return "SyncScenePlayTeamEntityNotify"
}

func (*DelScenePlayTeamEntityNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDelScenePlayTeamEntityNotify
}
func (*DelScenePlayTeamEntityNotify) ProtoMessageType() ProtoMessageType {
	return "DelScenePlayTeamEntityNotify"
}

func (*PlayerEyePointStateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerEyePointStateNotify
}
func (*PlayerEyePointStateNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerEyePointStateNotify"
}

func (*GetMapMarkTipsReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetMapMarkTipsReq }
func (*GetMapMarkTipsReq) ProtoMessageType() ProtoMessageType { return "GetMapMarkTipsReq" }

func (*GetMapMarkTipsRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetMapMarkTipsRsp }
func (*GetMapMarkTipsRsp) ProtoMessageType() ProtoMessageType { return "GetMapMarkTipsRsp" }

func (*ChangeWorldToSingleModeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeWorldToSingleModeReq
}
func (*ChangeWorldToSingleModeReq) ProtoMessageType() ProtoMessageType {
	return "ChangeWorldToSingleModeReq"
}

func (*ChangeWorldToSingleModeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeWorldToSingleModeRsp
}
func (*ChangeWorldToSingleModeRsp) ProtoMessageType() ProtoMessageType {
	return "ChangeWorldToSingleModeRsp"
}

func (*GetWorldMpInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetWorldMpInfoReq }
func (*GetWorldMpInfoReq) ProtoMessageType() ProtoMessageType { return "GetWorldMpInfoReq" }

func (*GetWorldMpInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetWorldMpInfoRsp }
func (*GetWorldMpInfoRsp) ProtoMessageType() ProtoMessageType { return "GetWorldMpInfoRsp" }

func (*EntityConfigHashNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityConfigHashNotify }
func (*EntityConfigHashNotify) ProtoMessageType() ProtoMessageType { return "EntityConfigHashNotify" }

func (*ForceDragAvatarNotify) ProtoCommand() ProtoCommand         { return ProtoCommandForceDragAvatarNotify }
func (*ForceDragAvatarNotify) ProtoMessageType() ProtoMessageType { return "ForceDragAvatarNotify" }

func (*MonsterPointArrayRouteUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMonsterPointArrayRouteUpdateNotify
}
func (*MonsterPointArrayRouteUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "MonsterPointArrayRouteUpdateNotify"
}

func (*ForceDragBackTransferNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandForceDragBackTransferNotify
}
func (*ForceDragBackTransferNotify) ProtoMessageType() ProtoMessageType {
	return "ForceDragBackTransferNotify"
}

func (*GetScenePerformanceReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetScenePerformanceReq }
func (*GetScenePerformanceReq) ProtoMessageType() ProtoMessageType { return "GetScenePerformanceReq" }

func (*GetScenePerformanceRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetScenePerformanceRsp }
func (*GetScenePerformanceRsp) ProtoMessageType() ProtoMessageType { return "GetScenePerformanceRsp" }

func (*SceneAudioNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneAudioNotify }
func (*SceneAudioNotify) ProtoMessageType() ProtoMessageType { return "SceneAudioNotify" }

func (*HitTreeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHitTreeNotify }
func (*HitTreeNotify) ProtoMessageType() ProtoMessageType { return "HitTreeNotify" }

func (*EntityTagChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityTagChangeNotify }
func (*EntityTagChangeNotify) ProtoMessageType() ProtoMessageType { return "EntityTagChangeNotify" }

func (*AvatarFollowRouteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarFollowRouteNotify
}
func (*AvatarFollowRouteNotify) ProtoMessageType() ProtoMessageType { return "AvatarFollowRouteNotify" }

func (*SceneEntityUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneEntityUpdateNotify
}
func (*SceneEntityUpdateNotify) ProtoMessageType() ProtoMessageType { return "SceneEntityUpdateNotify" }

func (*ClientHashDebugNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClientHashDebugNotify }
func (*ClientHashDebugNotify) ProtoMessageType() ProtoMessageType { return "ClientHashDebugNotify" }

func (*PlayerWorldSceneInfoListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerWorldSceneInfoListNotify
}
func (*PlayerWorldSceneInfoListNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerWorldSceneInfoListNotify"
}

func (*LuaEnvironmentEffectNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLuaEnvironmentEffectNotify
}
func (*LuaEnvironmentEffectNotify) ProtoMessageType() ProtoMessageType {
	return "LuaEnvironmentEffectNotify"
}

func (*ClientLoadingCostumeVerificationNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientLoadingCostumeVerificationNotify
}
func (*ClientLoadingCostumeVerificationNotify) ProtoMessageType() ProtoMessageType {
	return "ClientLoadingCostumeVerificationNotify"
}

func (*ShowClientGuideNotify) ProtoCommand() ProtoCommand         { return ProtoCommandShowClientGuideNotify }
func (*ShowClientGuideNotify) ProtoMessageType() ProtoMessageType { return "ShowClientGuideNotify" }

func (*ShowClientTutorialNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandShowClientTutorialNotify
}
func (*ShowClientTutorialNotify) ProtoMessageType() ProtoMessageType {
	return "ShowClientTutorialNotify"
}

func (*GetMapAreaReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetMapAreaReq }
func (*GetMapAreaReq) ProtoMessageType() ProtoMessageType { return "GetMapAreaReq" }

func (*GetMapAreaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetMapAreaRsp }
func (*GetMapAreaRsp) ProtoMessageType() ProtoMessageType { return "GetMapAreaRsp" }

func (*MapAreaChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMapAreaChangeNotify }
func (*MapAreaChangeNotify) ProtoMessageType() ProtoMessageType { return "MapAreaChangeNotify" }

func (*LeaveWorldNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLeaveWorldNotify }
func (*LeaveWorldNotify) ProtoMessageType() ProtoMessageType { return "LeaveWorldNotify" }

func (*GuestBeginEnterSceneNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGuestBeginEnterSceneNotify
}
func (*GuestBeginEnterSceneNotify) ProtoMessageType() ProtoMessageType {
	return "GuestBeginEnterSceneNotify"
}

func (*GuestPostEnterSceneNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGuestPostEnterSceneNotify
}
func (*GuestPostEnterSceneNotify) ProtoMessageType() ProtoMessageType {
	return "GuestPostEnterSceneNotify"
}

func (*LevelTagDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLevelTagDataNotify }
func (*LevelTagDataNotify) ProtoMessageType() ProtoMessageType { return "LevelTagDataNotify" }

func (*StopReminderNotify) ProtoCommand() ProtoCommand         { return ProtoCommandStopReminderNotify }
func (*StopReminderNotify) ProtoMessageType() ProtoMessageType { return "StopReminderNotify" }

func (*AreaPlayInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAreaPlayInfoNotify }
func (*AreaPlayInfoNotify) ProtoMessageType() ProtoMessageType { return "AreaPlayInfoNotify" }

func (*CheckGroupReplacedReq) ProtoCommand() ProtoCommand         { return ProtoCommandCheckGroupReplacedReq }
func (*CheckGroupReplacedReq) ProtoMessageType() ProtoMessageType { return "CheckGroupReplacedReq" }

func (*CheckGroupReplacedRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCheckGroupReplacedRsp }
func (*CheckGroupReplacedRsp) ProtoMessageType() ProtoMessageType { return "CheckGroupReplacedRsp" }

func (*DeathZoneObserveNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDeathZoneObserveNotify }
func (*DeathZoneObserveNotify) ProtoMessageType() ProtoMessageType { return "DeathZoneObserveNotify" }

func (*WorldChestOpenNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWorldChestOpenNotify }
func (*WorldChestOpenNotify) ProtoMessageType() ProtoMessageType { return "WorldChestOpenNotify" }

func (*WidgetQuickHitTreeReq) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetQuickHitTreeReq }
func (*WidgetQuickHitTreeReq) ProtoMessageType() ProtoMessageType { return "WidgetQuickHitTreeReq" }

func (*WidgetQuickHitTreeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetQuickHitTreeRsp }
func (*WidgetQuickHitTreeRsp) ProtoMessageType() ProtoMessageType { return "WidgetQuickHitTreeRsp" }

func (*BeginCameraSceneLookWithTemplateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBeginCameraSceneLookWithTemplateNotify
}
func (*BeginCameraSceneLookWithTemplateNotify) ProtoMessageType() ProtoMessageType {
	return "BeginCameraSceneLookWithTemplateNotify"
}

func (*RefreshEntityAuthNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshEntityAuthNotify
}
func (*RefreshEntityAuthNotify) ProtoMessageType() ProtoMessageType { return "RefreshEntityAuthNotify" }

func (*ScenePlayerBackgroundAvatarRefreshNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayerBackgroundAvatarRefreshNotify
}
func (*ScenePlayerBackgroundAvatarRefreshNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayerBackgroundAvatarRefreshNotify"
}
