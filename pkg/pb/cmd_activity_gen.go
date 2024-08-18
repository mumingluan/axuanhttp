
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetActivityScheduleReq) },
		func() ProtoMessage { return new(GetActivityScheduleRsp) },
		func() ProtoMessage { return new(GetActivityInfoReq) },
		func() ProtoMessage { return new(GetActivityInfoRsp) },
		func() ProtoMessage { return new(ActivityPlayOpenAnimNotify) },
		func() ProtoMessage { return new(ActivityInfoNotify) },
		func() ProtoMessage { return new(ActivityScheduleInfoNotify) },
		func() ProtoMessage { return new(ActivityTakeWatcherRewardReq) },
		func() ProtoMessage { return new(ActivityTakeWatcherRewardRsp) },
		func() ProtoMessage { return new(ActivityUpdateWatcherNotify) },
		func() ProtoMessage { return new(ActivitySelectAvatarCardReq) },
		func() ProtoMessage { return new(ActivitySelectAvatarCardRsp) },
		func() ProtoMessage { return new(ActivityCoinInfoNotify) },
		func() ProtoMessage { return new(SeaLampFlyLampReq) },
		func() ProtoMessage { return new(SeaLampFlyLampRsp) },
		func() ProtoMessage { return new(SeaLampTakeContributionRewardReq) },
		func() ProtoMessage { return new(SeaLampTakeContributionRewardRsp) },
		func() ProtoMessage { return new(SeaLampTakePhaseRewardReq) },
		func() ProtoMessage { return new(SeaLampTakePhaseRewardRsp) },
		func() ProtoMessage { return new(SeaLampContributeItemReq) },
		func() ProtoMessage { return new(SeaLampContributeItemRsp) },
		func() ProtoMessage { return new(SeaLampFlyLampNotify) },
		func() ProtoMessage { return new(SeaLampCoinNotify) },
		func() ProtoMessage { return new(SeaLampPopularityNotify) },
		func() ProtoMessage { return new(LoadActivityTerrainNotify) },
		func() ProtoMessage { return new(ServerAnnounceNotify) },
		func() ProtoMessage { return new(ServerAnnounceRevokeNotify) },
		func() ProtoMessage { return new(ActivityBannerNotify) },
		func() ProtoMessage { return new(ActivityBannerClearReq) },
		func() ProtoMessage { return new(ActivityBannerClearRsp) },
		func() ProtoMessage { return new(SalesmanDeliverItemReq) },
		func() ProtoMessage { return new(SalesmanDeliverItemRsp) },
		func() ProtoMessage { return new(SalesmanTakeRewardReq) },
		func() ProtoMessage { return new(SalesmanTakeRewardRsp) },
		func() ProtoMessage { return new(ActivityCondStateChangeNotify) },
		func() ProtoMessage { return new(SalesmanTakeSpecialRewardReq) },
		func() ProtoMessage { return new(SalesmanTakeSpecialRewardRsp) },
		func() ProtoMessage { return new(GetAuthSalesmanInfoReq) },
		func() ProtoMessage { return new(GetAuthSalesmanInfoRsp) },
		func() ProtoMessage { return new(EnterTrialAvatarActivityDungeonReq) },
		func() ProtoMessage { return new(EnterTrialAvatarActivityDungeonRsp) },
		func() ProtoMessage { return new(ReceivedTrialAvatarActivityRewardReq) },
		func() ProtoMessage { return new(ReceivedTrialAvatarActivityRewardRsp) },
		func() ProtoMessage { return new(TrialAvatarFirstPassDungeonNotify) },
		func() ProtoMessage { return new(TrialAvatarInDungeonIndexNotify) },
		func() ProtoMessage { return new(TakeDeliveryDailyRewardReq) },
		func() ProtoMessage { return new(TakeDeliveryDailyRewardRsp) },
		func() ProtoMessage { return new(FinishDeliveryNotify) },
		func() ProtoMessage { return new(SelectAsterMidDifficultyReq) },
		func() ProtoMessage { return new(SelectAsterMidDifficultyRsp) },
		func() ProtoMessage { return new(AsterProgressInfoNotify) },
		func() ProtoMessage { return new(AsterLittleInfoNotify) },
		func() ProtoMessage { return new(AsterMidInfoNotify) },
		func() ProtoMessage { return new(AsterMiscInfoNotify) },
		func() ProtoMessage { return new(TakeAsterSpecialRewardReq) },
		func() ProtoMessage { return new(TakeAsterSpecialRewardRsp) },
		func() ProtoMessage { return new(AsterLargeInfoNotify) },
		func() ProtoMessage { return new(FlightActivitySettleNotify) },
		func() ProtoMessage { return new(FlightActivityRestartReq) },
		func() ProtoMessage { return new(FlightActivityRestartRsp) },
		func() ProtoMessage { return new(AsterMidCampInfoNotify) },
		func() ProtoMessage { return new(DragonSpineChapterOpenNotify) },
		func() ProtoMessage { return new(DragonSpineChapterProgressChangeNotify) },
		func() ProtoMessage { return new(DragonSpineChapterFinishNotify) },
		func() ProtoMessage { return new(DragonSpineCoinChangeNotify) },
		func() ProtoMessage { return new(ActivitySaleChangeNotify) },
		func() ProtoMessage { return new(StartEffigyChallengeReq) },
		func() ProtoMessage { return new(StartEffigyChallengeRsp) },
		func() ProtoMessage { return new(EffigyChallengeInfoNotify) },
		func() ProtoMessage { return new(EffigyChallengeResultNotify) },
		func() ProtoMessage { return new(TakeEffigyFirstPassRewardReq) },
		func() ProtoMessage { return new(TakeEffigyFirstPassRewardRsp) },
		func() ProtoMessage { return new(TakeEffigyRewardReq) },
		func() ProtoMessage { return new(TakeEffigyRewardRsp) },
		func() ProtoMessage { return new(SelectEffigyChallengeConditionReq) },
		func() ProtoMessage { return new(SelectEffigyChallengeConditionRsp) },
		func() ProtoMessage { return new(RestartEffigyChallengeReq) },
		func() ProtoMessage { return new(RestartEffigyChallengeRsp) },
		func() ProtoMessage { return new(TreasureMapRegionInfoNotify) },
		func() ProtoMessage { return new(TreasureMapCurrencyNotify) },
		func() ProtoMessage { return new(TreasureMapRegionActiveNotify) },
		func() ProtoMessage { return new(TreasureMapMpChallengeNotify) },
		func() ProtoMessage { return new(TreasureMapBonusChallengeNotify) },
		func() ProtoMessage { return new(TreasureMapGuideTaskDoneNotify) },
		func() ProtoMessage { return new(TreasureMapPreTaskDoneNotify) },
		func() ProtoMessage { return new(BlessingScanReq) },
		func() ProtoMessage { return new(BlessingScanRsp) },
		func() ProtoMessage { return new(BlessingRedeemRewardReq) },
		func() ProtoMessage { return new(BlessingRedeemRewardRsp) },
		func() ProtoMessage { return new(BlessingGetFriendPicListReq) },
		func() ProtoMessage { return new(BlessingGetFriendPicListRsp) },
		func() ProtoMessage { return new(BlessingGiveFriendPicReq) },
		func() ProtoMessage { return new(BlessingGiveFriendPicRsp) },
		func() ProtoMessage { return new(BlessingAcceptGivePicReq) },
		func() ProtoMessage { return new(BlessingAcceptGivePicRsp) },
		func() ProtoMessage { return new(BlessingGetAllRecvPicRecordListReq) },
		func() ProtoMessage { return new(BlessingGetAllRecvPicRecordListRsp) },
		func() ProtoMessage { return new(BlessingRecvFriendPicNotify) },
		func() ProtoMessage { return new(BlessingAcceptAllGivePicReq) },
		func() ProtoMessage { return new(BlessingAcceptAllGivePicRsp) },
		func() ProtoMessage { return new(ExpeditionStartReq) },
		func() ProtoMessage { return new(ExpeditionStartRsp) },
		func() ProtoMessage { return new(ExpeditionRecallReq) },
		func() ProtoMessage { return new(ExpeditionRecallRsp) },
		func() ProtoMessage { return new(ExpeditionTakeRewardReq) },
		func() ProtoMessage { return new(ExpeditionTakeRewardRsp) },
		func() ProtoMessage { return new(GetExpeditionAssistInfoListReq) },
		func() ProtoMessage { return new(GetExpeditionAssistInfoListRsp) },
		func() ProtoMessage { return new(SetCurExpeditionChallengeIdReq) },
		func() ProtoMessage { return new(SetCurExpeditionChallengeIdRsp) },
		func() ProtoMessage { return new(ExpeditionChallengeEnterRegionNotify) },
		func() ProtoMessage { return new(ExpeditionChallengeFinishedNotify) },
		func() ProtoMessage { return new(FleurFairBalloonSettleNotify) },
		func() ProtoMessage { return new(FleurFairFallSettleNotify) },
		func() ProtoMessage { return new(FleurFairMusicGameSettleReq) },
		func() ProtoMessage { return new(FleurFairMusicGameSettleRsp) },
		func() ProtoMessage { return new(FleurFairMusicGameStartReq) },
		func() ProtoMessage { return new(FleurFairMusicGameStartRsp) },
		func() ProtoMessage { return new(FleurFairReplayMiniGameReq) },
		func() ProtoMessage { return new(FleurFairReplayMiniGameRsp) },
		func() ProtoMessage { return new(StartArenaChallengeLevelReq) },
		func() ProtoMessage { return new(StartArenaChallengeLevelRsp) },
		func() ProtoMessage { return new(ArenaChallengeFinishNotify) },
		func() ProtoMessage { return new(WaterSpritePhaseFinishNotify) },
		func() ProtoMessage { return new(ActivityTakeWatcherRewardBatchReq) },
		func() ProtoMessage { return new(ActivityTakeWatcherRewardBatchRsp) },
		func() ProtoMessage { return new(ChannelerSlabStageActiveChallengeIndexNotify) },
		func() ProtoMessage { return new(ChannelerSlabStageOneoffDungeonNotify) },
		func() ProtoMessage { return new(ChannellerSlabWearBuffReq) },
		func() ProtoMessage { return new(ChannellerSlabWearBuffRsp) },
		func() ProtoMessage { return new(ChannellerSlabTakeoffBuffReq) },
		func() ProtoMessage { return new(ChannellerSlabTakeoffBuffRsp) },
		func() ProtoMessage { return new(ChannellerSlabEnterLoopDungeonReq) },
		func() ProtoMessage { return new(ChannellerSlabEnterLoopDungeonRsp) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonTakeFirstPassRewardReq) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonTakeFirstPassRewardRsp) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonTakeScoreRewardReq) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonTakeScoreRewardRsp) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonChallengeInfoNotify) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonSelectConditionReq) },
		func() ProtoMessage { return new(ChannellerSlabLoopDungeonSelectConditionRsp) },
		func() ProtoMessage { return new(ChannellerSlabOneOffDungeonInfoReq) },
		func() ProtoMessage { return new(ChannellerSlabOneOffDungeonInfoRsp) },
		func() ProtoMessage { return new(ChannellerSlabOneOffDungeonInfoNotify) },
		func() ProtoMessage { return new(ChannellerSlabSaveAssistInfoReq) },
		func() ProtoMessage { return new(ChannellerSlabSaveAssistInfoRsp) },
		func() ProtoMessage { return new(MistTrialSelectAvatarAndEnterDungeonReq) },
		func() ProtoMessage { return new(MistTrialSelectAvatarAndEnterDungeonRsp) },
		func() ProtoMessage { return new(MistTrialGetChallengeMissionReq) },
		func() ProtoMessage { return new(MistTrialGetChallengeMissionRsp) },
		func() ProtoMessage { return new(MistTrialDunegonFailNotify) },
		func() ProtoMessage { return new(ChannellerSlabCheckEnterLoopDungeonReq) },
		func() ProtoMessage { return new(ChannellerSlabCheckEnterLoopDungeonRsp) },
		func() ProtoMessage { return new(HideAndSeekSelectSkillReq) },
		func() ProtoMessage { return new(HideAndSeekSelectSkillRsp) },
		func() ProtoMessage { return new(ActivityTakeScoreRewardReq) },
		func() ProtoMessage { return new(ActivityTakeScoreRewardRsp) },
		func() ProtoMessage { return new(ActivityTakeAllScoreRewardReq) },
		func() ProtoMessage { return new(ActivityTakeAllScoreRewardRsp) },
		func() ProtoMessage { return new(HideAndSeekChooseMapReq) },
		func() ProtoMessage { return new(HideAndSeekChooseMapRsp) },
		func() ProtoMessage { return new(CommonPlayerTipsNotify) },
		func() ProtoMessage { return new(FindHilichurlFinishSecondQuestNotify) },
		func() ProtoMessage { return new(FindHilichurlAcceptQuestNotify) },
		func() ProtoMessage { return new(SummerTimeFloatSignalPositionNotify) },
		func() ProtoMessage { return new(SummerTimeFloatSignalUpdateNotify) },
		func() ProtoMessage { return new(SummerTimeSprintBoatSettleNotify) },
		func() ProtoMessage { return new(SummerTimeSprintBoatRestartReq) },
		func() ProtoMessage { return new(SummerTimeSprintBoatRestartRsp) },
		func() ProtoMessage { return new(StartBuoyantCombatGalleryReq) },
		func() ProtoMessage { return new(StartBuoyantCombatGalleryRsp) },
		func() ProtoMessage { return new(BuoyantCombatSettleNotify) },
		func() ProtoMessage { return new(SetLimitOptimizationNotify) },
		func() ProtoMessage { return new(EchoShellUpdateNotify) },
		func() ProtoMessage { return new(EchoShellTakeRewardReq) },
		func() ProtoMessage { return new(EchoShellTakeRewardRsp) },
		func() ProtoMessage { return new(BounceConjuringSettleNotify) },
		func() ProtoMessage { return new(BlitzRushParkourRestartReq) },
		func() ProtoMessage { return new(BlitzRushParkourRestartRsp) },
		func() ProtoMessage { return new(EnterChessDungeonReq) },
		func() ProtoMessage { return new(EnterChessDungeonRsp) },
		func() ProtoMessage { return new(TreasureMapHostInfoNotify) },
		func() ProtoMessage { return new(SumoSaveTeamReq) },
		func() ProtoMessage { return new(SumoSaveTeamRsp) },
		func() ProtoMessage { return new(SumoSelectTeamAndEnterDungeonReq) },
		func() ProtoMessage { return new(SumoSelectTeamAndEnterDungeonRsp) },
		func() ProtoMessage { return new(SumoDungeonSettleNotify) },
		func() ProtoMessage { return new(SumoEnterDungeonNotify) },
		func() ProtoMessage { return new(SumoSwitchTeamReq) },
		func() ProtoMessage { return new(SumoSwitchTeamRsp) },
		func() ProtoMessage { return new(SumoLeaveDungeonNotify) },
		func() ProtoMessage { return new(SumoRestartDungeonReq) },
		func() ProtoMessage { return new(SumoRestartDungeonRsp) },
		func() ProtoMessage { return new(ActivityDisableTransferPointInteractionNotify) },
		func() ProtoMessage { return new(SumoSetNoSwitchPunishTimeNotify) },
		func() ProtoMessage { return new(FishingGallerySettleNotify) },
		func() ProtoMessage { return new(LunaRiteSacrificeReq) },
		func() ProtoMessage { return new(LunaRiteSacrificeRsp) },
		func() ProtoMessage { return new(LunaRiteTakeSacrificeRewardReq) },
		func() ProtoMessage { return new(LunaRiteTakeSacrificeRewardRsp) },
		func() ProtoMessage { return new(LunaRiteHintPointReq) },
		func() ProtoMessage { return new(LunaRiteHintPointRsp) },
		func() ProtoMessage { return new(LunaRiteHintPointRemoveNotify) },
		func() ProtoMessage { return new(LunaRiteGroupBundleRegisterNotify) },
		func() ProtoMessage { return new(LunaRiteAreaFinishNotify) },
		func() ProtoMessage { return new(PlantFlowerGetSeedInfoReq) },
		func() ProtoMessage { return new(PlantFlowerGetSeedInfoRsp) },
		func() ProtoMessage { return new(PlantFlowerTakeSeedRewardReq) },
		func() ProtoMessage { return new(PlantFlowerTakeSeedRewardRsp) },
		func() ProtoMessage { return new(PlantFlowerSetFlowerWishReq) },
		func() ProtoMessage { return new(PlantFlowerSetFlowerWishRsp) },
		func() ProtoMessage { return new(PlantFlowerGetFriendFlowerWishListReq) },
		func() ProtoMessage { return new(PlantFlowerGetFriendFlowerWishListRsp) },
		func() ProtoMessage { return new(PlantFlowerGiveFriendFlowerReq) },
		func() ProtoMessage { return new(PlantFlowerGiveFriendFlowerRsp) },
		func() ProtoMessage { return new(PlantFlowerGetRecvFlowerListReq) },
		func() ProtoMessage { return new(PlantFlowerGetRecvFlowerListRsp) },
		func() ProtoMessage { return new(PlantFlowerHaveRecvFlowerNotify) },
		func() ProtoMessage { return new(PlantFlowerAcceptGiveFlowerReq) },
		func() ProtoMessage { return new(PlantFlowerAcceptGiveFlowerRsp) },
		func() ProtoMessage { return new(PlantFlowerAcceptAllGiveFlowerReq) },
		func() ProtoMessage { return new(PlantFlowerAcceptAllGiveFlowerRsp) },
		func() ProtoMessage { return new(PlantFlowerGetCanGiveFriendFlowerReq) },
		func() ProtoMessage { return new(PlantFlowerGetCanGiveFriendFlowerRsp) },
		func() ProtoMessage { return new(PlantFlowerEditFlowerCombinationReq) },
		func() ProtoMessage { return new(PlantFlowerEditFlowerCombinationRsp) },
		func() ProtoMessage { return new(MusicGameSettleReq) },
		func() ProtoMessage { return new(MusicGameSettleRsp) },
		func() ProtoMessage { return new(MusicGameStartReq) },
		func() ProtoMessage { return new(MusicGameStartRsp) },
		func() ProtoMessage { return new(DoRoguelikeDungeonCardGachaReq) },
		func() ProtoMessage { return new(DoRoguelikeDungeonCardGachaRsp) },
		func() ProtoMessage { return new(RefreshRoguelikeDungeonCardReq) },
		func() ProtoMessage { return new(RefreshRoguelikeDungeonCardRsp) },
		func() ProtoMessage { return new(SelectRoguelikeDungeonCardReq) },
		func() ProtoMessage { return new(SelectRoguelikeDungeonCardRsp) },
		func() ProtoMessage { return new(EquipRoguelikeRuneReq) },
		func() ProtoMessage { return new(EquipRoguelikeRuneRsp) },
		func() ProtoMessage { return new(TriggerRoguelikeRuneReq) },
		func() ProtoMessage { return new(TriggerRoguelikeRuneRsp) },
		func() ProtoMessage { return new(TriggerRoguelikeCurseNotify) },
		func() ProtoMessage { return new(UpgradeRoguelikeShikigamiReq) },
		func() ProtoMessage { return new(UpgradeRoguelikeShikigamiRsp) },
		func() ProtoMessage { return new(RoguelikeSelectAvatarAndEnterDungeonReq) },
		func() ProtoMessage { return new(RoguelikeSelectAvatarAndEnterDungeonRsp) },
		func() ProtoMessage { return new(RoguelikeGiveUpReq) },
		func() ProtoMessage { return new(RoguelikeGiveUpRsp) },
		func() ProtoMessage { return new(RoguelikeTakeStageFirstPassRewardReq) },
		func() ProtoMessage { return new(RoguelikeTakeStageFirstPassRewardRsp) },
		func() ProtoMessage { return new(GiveUpRoguelikeDungeonCardReq) },
		func() ProtoMessage { return new(GiveUpRoguelikeDungeonCardRsp) },
		func() ProtoMessage { return new(EnterRoguelikeDungeonNotify) },
		func() ProtoMessage { return new(StartRogueEliteCellChallengeReq) },
		func() ProtoMessage { return new(StartRogueEliteCellChallengeRsp) },
		func() ProtoMessage { return new(StartRogueNormalCellChallengeReq) },
		func() ProtoMessage { return new(StartRogueNormalCellChallengeRsp) },
		func() ProtoMessage { return new(RogueCellUpdateNotify) },
		func() ProtoMessage { return new(RogueDungeonPlayerCellChangeNotify) },
		func() ProtoMessage { return new(RogueHealAvatarsReq) },
		func() ProtoMessage { return new(RogueHealAvatarsRsp) },
		func() ProtoMessage { return new(RogueResumeDungeonReq) },
		func() ProtoMessage { return new(RogueResumeDungeonRsp) },
		func() ProtoMessage { return new(ClearRoguelikeCurseNotify) },
		func() ProtoMessage { return new(RoguelikeCardGachaNotify) },
		func() ProtoMessage { return new(RogueSwitchAvatarReq) },
		func() ProtoMessage { return new(RogueSwitchAvatarRsp) },
		func() ProtoMessage { return new(DisableRoguelikeTrapNotify) },
		func() ProtoMessage { return new(RoguelikeRuneRecordUpdateNotify) },
		func() ProtoMessage { return new(RoguelikeMistClearNotify) },
		func() ProtoMessage { return new(RoguelikeEffectDataNotify) },
		func() ProtoMessage { return new(RoguelikeEffectViewReq) },
		func() ProtoMessage { return new(RoguelikeEffectViewRsp) },
		func() ProtoMessage { return new(RoguelikeResourceBonusPropUpdateNotify) },
		func() ProtoMessage { return new(RoguelikeRefreshCardCostUpdateNotify) },
		func() ProtoMessage { return new(DigActivityMarkPointChangeNotify) },
		func() ProtoMessage { return new(DigActivityChangeGadgetStateReq) },
		func() ProtoMessage { return new(DigActivityChangeGadgetStateRsp) },
		func() ProtoMessage { return new(WinterCampStageInfoChangeNotify) },
		func() ProtoMessage { return new(WinterCampRaceScoreNotify) },
		func() ProtoMessage { return new(WinterCampGiveFriendItemReq) },
		func() ProtoMessage { return new(WinterCampGiveFriendItemRsp) },
		func() ProtoMessage { return new(WinterCampSetWishListReq) },
		func() ProtoMessage { return new(WinterCampSetWishListRsp) },
		func() ProtoMessage { return new(WinterCampGetFriendWishListReq) },
		func() ProtoMessage { return new(WinterCampGetFriendWishListRsp) },
		func() ProtoMessage { return new(WinterCampRecvItemNotify) },
		func() ProtoMessage { return new(WinterCampAcceptGiveItemReq) },
		func() ProtoMessage { return new(WinterCampAcceptGiveItemRsp) },
		func() ProtoMessage { return new(WinterCampAcceptAllGiveItemReq) },
		func() ProtoMessage { return new(WinterCampAcceptAllGiveItemRsp) },
		func() ProtoMessage { return new(WinterCampGetCanGiveFriendItemReq) },
		func() ProtoMessage { return new(WinterCampGetCanGiveFriendItemRsp) },
		func() ProtoMessage { return new(WinterCampGetRecvItemListReq) },
		func() ProtoMessage { return new(WinterCampGetRecvItemListRsp) },
		func() ProtoMessage { return new(WinterCampEditSnowmanCombinationReq) },
		func() ProtoMessage { return new(WinterCampEditSnowmanCombinationRsp) },
		func() ProtoMessage { return new(WinterCampTriathlonSettleNotify) },
		func() ProtoMessage { return new(WinterCampTakeExploreRewardReq) },
		func() ProtoMessage { return new(WinterCampTakeExploreRewardRsp) },
		func() ProtoMessage { return new(WinterCampTakeBattleRewardReq) },
		func() ProtoMessage { return new(WinterCampTakeBattleRewardRsp) },
		func() ProtoMessage { return new(WinterCampTriathlonRestartReq) },
		func() ProtoMessage { return new(WinterCampTriathlonRestartRsp) },
		func() ProtoMessage { return new(MistTrialSettleNotify) },
		func() ProtoMessage { return new(MistTrialGetDungeonExhibitionDataReq) },
		func() ProtoMessage { return new(MistTrialGetDungeonExhibitionDataRsp) },
		func() ProtoMessage { return new(PotionResetChallengeReq) },
		func() ProtoMessage { return new(PotionResetChallengeRsp) },
		func() ProtoMessage { return new(PotionEnterDungeonReq) },
		func() ProtoMessage { return new(PotionEnterDungeonRsp) },
		func() ProtoMessage { return new(PotionEnterDungeonNotify) },
		func() ProtoMessage { return new(PotionSaveDungeonResultReq) },
		func() ProtoMessage { return new(PotionSaveDungeonResultRsp) },
		func() ProtoMessage { return new(PotionRestartDungeonReq) },
		func() ProtoMessage { return new(PotionRestartDungeonRsp) },
		func() ProtoMessage { return new(TanukiTravelFinishGuideQuestNotify) },
		func() ProtoMessage { return new(FinishLanternProjectionReq) },
		func() ProtoMessage { return new(FinishLanternProjectionRsp) },
		func() ProtoMessage { return new(ViewLanternProjectionTipsReq) },
		func() ProtoMessage { return new(ViewLanternProjectionTipsRsp) },
		func() ProtoMessage { return new(ViewLanternProjectionLevelTipsReq) },
		func() ProtoMessage { return new(ViewLanternProjectionLevelTipsRsp) },
		func() ProtoMessage { return new(SalvagePreventSettleNotify) },
		func() ProtoMessage { return new(SalvageEscortSettleNotify) },
		func() ProtoMessage { return new(LanternRiteTakeSkinRewardReq) },
		func() ProtoMessage { return new(LanternRiteTakeSkinRewardRsp) },
		func() ProtoMessage { return new(SalvagePreventRestartReq) },
		func() ProtoMessage { return new(SalvagePreventRestartRsp) },
		func() ProtoMessage { return new(SalvageEscortRestartReq) },
		func() ProtoMessage { return new(SalvageEscortRestartRsp) },
		func() ProtoMessage { return new(LanternRiteStartFireworksReformReq) },
		func() ProtoMessage { return new(LanternRiteStartFireworksReformRsp) },
		func() ProtoMessage { return new(LanternRiteDoFireworksReformReq) },
		func() ProtoMessage { return new(LanternRiteDoFireworksReformRsp) },
		func() ProtoMessage { return new(LanternRiteEndFireworksReformReq) },
		func() ProtoMessage { return new(LanternRiteEndFireworksReformRsp) },
		func() ProtoMessage { return new(UpdateSalvageBundleMarkReq) },
		func() ProtoMessage { return new(UpdateSalvageBundleMarkRsp) },
		func() ProtoMessage { return new(MichiaeMatsuriDarkPressureLevelUpdateNotify) },
		func() ProtoMessage { return new(MichiaeMatsuriInteractStatueReq) },
		func() ProtoMessage { return new(MichiaeMatsuriInteractStatueRsp) },
		func() ProtoMessage { return new(MichiaeMatsuriUnlockCrystalSkillReq) },
		func() ProtoMessage { return new(MichiaeMatsuriUnlockCrystalSkillRsp) },
		func() ProtoMessage { return new(MichiaeMatsuriStartBossChallengeReq) },
		func() ProtoMessage { return new(MichiaeMatsuriStartBossChallengeRsp) },
		func() ProtoMessage { return new(MichiaeMatsuriStartDarkChallengeReq) },
		func() ProtoMessage { return new(MichiaeMatsuriStartDarkChallengeRsp) },
		func() ProtoMessage { return new(MichiaeMatsuriRemoveChestMarkNotify) },
		func() ProtoMessage { return new(MichiaeMatsuriRemoveChallengeMarkNotify) },
		func() ProtoMessage { return new(MichiaeMatsuriGainCrystalExpUpdateNotify) },
		func() ProtoMessage { return new(BartenderCompleteOrderReq) },
		func() ProtoMessage { return new(BartenderCompleteOrderRsp) },
		func() ProtoMessage { return new(BartenderCancelOrderReq) },
		func() ProtoMessage { return new(BartenderCancelOrderRsp) },
		func() ProtoMessage { return new(BartenderGetFormulaReq) },
		func() ProtoMessage { return new(BartenderGetFormulaRsp) },
		func() ProtoMessage { return new(BartenderStartLevelReq) },
		func() ProtoMessage { return new(BartenderStartLevelRsp) },
		func() ProtoMessage { return new(BartenderCancelLevelReq) },
		func() ProtoMessage { return new(BartenderCancelLevelRsp) },
		func() ProtoMessage { return new(BartenderLevelProgressNotify) },
		func() ProtoMessage { return new(BartenderFinishLevelReq) },
		func() ProtoMessage { return new(BartenderFinishLevelRsp) },
		func() ProtoMessage { return new(CrystalLinkEnterDungeonReq) },
		func() ProtoMessage { return new(CrystalLinkEnterDungeonRsp) },
		func() ProtoMessage { return new(CrystalLinkDungeonInfoNotify) },
		func() ProtoMessage { return new(CrystalLinkRestartDungeonReq) },
		func() ProtoMessage { return new(CrystalLinkRestartDungeonRsp) },
		func() ProtoMessage { return new(QuickOpenActivityReq) },
		func() ProtoMessage { return new(QuickOpenActivityRsp) },
		func() ProtoMessage { return new(IrodoriEditFlowerCombinationReq) },
		func() ProtoMessage { return new(IrodoriEditFlowerCombinationRsp) },
		func() ProtoMessage { return new(IrodoriScanEntityReq) },
		func() ProtoMessage { return new(IrodoriScanEntityRsp) },
		func() ProtoMessage { return new(IrodoriFillPoetryReq) },
		func() ProtoMessage { return new(IrodoriFillPoetryRsp) },
		func() ProtoMessage { return new(IrodoriChessEquipCardReq) },
		func() ProtoMessage { return new(IrodoriChessEquipCardRsp) },
		func() ProtoMessage { return new(IrodoriChessUnequipCardReq) },
		func() ProtoMessage { return new(IrodoriChessUnequipCardRsp) },
		func() ProtoMessage { return new(EnterIrodoriChessDungeonReq) },
		func() ProtoMessage { return new(EnterIrodoriChessDungeonRsp) },
		func() ProtoMessage { return new(IrodoriMasterStartGalleryReq) },
		func() ProtoMessage { return new(IrodoriMasterStartGalleryRsp) },
		func() ProtoMessage { return new(IrodoriMasterGalleryCgEndNotify) },
		func() ProtoMessage { return new(IrodoriMasterGallerySettleNotify) },
		func() ProtoMessage { return new(PhotoActivityFinishReq) },
		func() ProtoMessage { return new(PhotoActivityFinishRsp) },
		func() ProtoMessage { return new(PhotoActivityClientViewReq) },
		func() ProtoMessage { return new(PhotoActivityClientViewRsp) },
		func() ProtoMessage { return new(SpiceActivityFinishMakeSpiceReq) },
		func() ProtoMessage { return new(SpiceActivityFinishMakeSpiceRsp) },
		func() ProtoMessage { return new(SpiceActivityProcessFoodReq) },
		func() ProtoMessage { return new(SpiceActivityProcessFoodRsp) },
		func() ProtoMessage { return new(SpiceActivityGivingRecordNotify) },
		func() ProtoMessage { return new(GachaActivityPercentNotify) },
		func() ProtoMessage { return new(GachaActivityUpdateElemNotify) },
		func() ProtoMessage { return new(GachaActivityCreateRobotReq) },
		func() ProtoMessage { return new(GachaActivityCreateRobotRsp) },
		func() ProtoMessage { return new(GachaActivityTakeRewardReq) },
		func() ProtoMessage { return new(GachaActivityTakeRewardRsp) },
		func() ProtoMessage { return new(GachaActivityResetReq) },
		func() ProtoMessage { return new(GachaActivityResetRsp) },
		func() ProtoMessage { return new(GachaActivityNextStageReq) },
		func() ProtoMessage { return new(GachaActivityNextStageRsp) },
		func() ProtoMessage { return new(ActivityGiveFriendGiftReq) },
		func() ProtoMessage { return new(ActivityGiveFriendGiftRsp) },
		func() ProtoMessage { return new(ActivityGetRecvGiftListReq) },
		func() ProtoMessage { return new(ActivityGetRecvGiftListRsp) },
		func() ProtoMessage { return new(ActivityHaveRecvGiftNotify) },
		func() ProtoMessage { return new(ActivityAcceptGiveGiftReq) },
		func() ProtoMessage { return new(ActivityAcceptGiveGiftRsp) },
		func() ProtoMessage { return new(ActivityAcceptAllGiveGiftReq) },
		func() ProtoMessage { return new(ActivityAcceptAllGiveGiftRsp) },
		func() ProtoMessage { return new(ActivityGetCanGiveFriendGiftReq) },
		func() ProtoMessage { return new(ActivityGetCanGiveFriendGiftRsp) },
		func() ProtoMessage { return new(ActivitySetGiftWishReq) },
		func() ProtoMessage { return new(ActivitySetGiftWishRsp) },
		func() ProtoMessage { return new(ActivityGetFriendGiftWishListReq) },
		func() ProtoMessage { return new(ActivityGetFriendGiftWishListRsp) },
		func() ProtoMessage { return new(LuminanceStoneChallengeSettleNotify) },
		func() ProtoMessage { return new(StartRogueDiaryPlayReq) },
		func() ProtoMessage { return new(StartRogueDiaryPlayRsp) },
		func() ProtoMessage { return new(ResetRogueDiaryPlayReq) },
		func() ProtoMessage { return new(ResetRogueDiaryPlayRsp) },
		func() ProtoMessage { return new(EnterRogueDiaryDungeonReq) },
		func() ProtoMessage { return new(EnterRogueDiaryDungeonRsp) },
		func() ProtoMessage { return new(ResumeRogueDiaryDungeonReq) },
		func() ProtoMessage { return new(ResumeRogueDiaryDungeonRsp) },
		func() ProtoMessage { return new(RogueDiaryDungeonInfoNotify) },
		func() ProtoMessage { return new(RogueDiaryDungeonSettleNotify) },
		func() ProtoMessage { return new(StartRogueDiaryRoomReq) },
		func() ProtoMessage { return new(StartRogueDiaryRoomRsp) },
		func() ProtoMessage { return new(RogueDiaryTiredAvatarNotify) },
		func() ProtoMessage { return new(ReserveRogueDiaryAvatarReq) },
		func() ProtoMessage { return new(ReserveRogueDiaryAvatarRsp) },
		func() ProtoMessage { return new(GetRogueDairyRepairInfoReq) },
		func() ProtoMessage { return new(GetRogueDairyRepairInfoRsp) },
		func() ProtoMessage { return new(RefreshRogueDiaryCardReq) },
		func() ProtoMessage { return new(RefreshRogueDiaryCardRsp) },
		func() ProtoMessage { return new(RogueFinishRepairReq) },
		func() ProtoMessage { return new(RogueFinishRepairRsp) },
		func() ProtoMessage { return new(TryInterruptRogueDiaryDungeonReq) },
		func() ProtoMessage { return new(TryInterruptRogueDiaryDungeonRsp) },
		func() ProtoMessage { return new(RogueDiaryRepairInfoNotify) },
		func() ProtoMessage { return new(RetryCurRogueDiaryDungeonReq) },
		func() ProtoMessage { return new(RetryCurRogueDiaryDungeonRsp) },
		func() ProtoMessage { return new(RogueDiaryReviveAvatarReq) },
		func() ProtoMessage { return new(RogueDiaryReviveAvatarRsp) },
		func() ProtoMessage { return new(TryEnterNextRogueDiaryDungeonReq) },
		func() ProtoMessage { return new(TryEnterNextRogueDiaryDungeonRsp) },
		func() ProtoMessage { return new(RogueDiaryCoinAddNotify) },
		func() ProtoMessage { return new(SummerTimeV2BoatSettleNotify) },
		func() ProtoMessage { return new(ActivityPushTipsInfoNotify) },
		func() ProtoMessage { return new(ActivityReadPushTipsReq) },
		func() ProtoMessage { return new(ActivityReadPushTipsRsp) },
		func() ProtoMessage { return new(SummerTimeV2RestartBoatGalleryReq) },
		func() ProtoMessage { return new(SummerTimeV2RestartBoatGalleryRsp) },
		func() ProtoMessage { return new(SummerTimeV2RestartDungeonReq) },
		func() ProtoMessage { return new(SummerTimeV2RestartDungeonRsp) },
		func() ProtoMessage { return new(IslandPartySettleNotify) },
		func() ProtoMessage { return new(GearActivityStartPlayGearReq) },
		func() ProtoMessage { return new(GearActivityStartPlayGearRsp) },
		func() ProtoMessage { return new(GearActivityFinishPlayGearReq) },
		func() ProtoMessage { return new(GearActivityFinishPlayGearRsp) },
		func() ProtoMessage { return new(GearActivityStartPlayPictureReq) },
		func() ProtoMessage { return new(GearActivityStartPlayPictureRsp) },
		func() ProtoMessage { return new(GearActivityFinishPlayPictureReq) },
		func() ProtoMessage { return new(GearActivityFinishPlayPictureRsp) },
		func() ProtoMessage { return new(GravenInnocenceRaceSettleNotify) },
		func() ProtoMessage { return new(GravenInnocenceRaceRestartReq) },
		func() ProtoMessage { return new(GravenInnocenceRaceRestartRsp) },
		func() ProtoMessage { return new(GravenInnocenceEditCarveCombinationReq) },
		func() ProtoMessage { return new(GravenInnocenceEditCarveCombinationRsp) },
		func() ProtoMessage { return new(GravenInnocencePhotoFinishReq) },
		func() ProtoMessage { return new(GravenInnocencePhotoFinishRsp) },
		func() ProtoMessage { return new(GravenInnocencePhotoReminderNotify) },
		func() ProtoMessage { return new(InstableSprayEnterDungeonReq) },
		func() ProtoMessage { return new(InstableSprayEnterDungeonRsp) },
		func() ProtoMessage { return new(InstableSpraySwitchTeamReq) },
		func() ProtoMessage { return new(InstableSpraySwitchTeamRsp) },
		func() ProtoMessage { return new(InstableSprayLevelFinishNotify) },
		func() ProtoMessage { return new(InstableSprayRestartDungeonReq) },
		func() ProtoMessage { return new(InstableSprayRestartDungeonRsp) },
		func() ProtoMessage { return new(MuqadasPotionActivityEnterDungeonReq) },
		func() ProtoMessage { return new(MuqadasPotionActivityEnterDungeonRsp) },
		func() ProtoMessage { return new(MuqadasPotionDungeonSettleNotify) },
		func() ProtoMessage { return new(MuqadasPotionRestartDungeonReq) },
		func() ProtoMessage { return new(MuqadasPotionRestartDungeonRsp) },
		func() ProtoMessage { return new(MuqadasPotionCaptureWeaknessReq) },
		func() ProtoMessage { return new(MuqadasPotionCaptureWeaknessRsp) },
		func() ProtoMessage { return new(TreasureSeelieCollectOrbsNotify) },
		func() ProtoMessage { return new(VintageMarketDeliverItemReq) },
		func() ProtoMessage { return new(VintageMarketDeliverItemRsp) },
		func() ProtoMessage { return new(SceneGalleryVintageHuntingSettleNotify) },
		func() ProtoMessage { return new(VintagePresentFinishNoify) },
		func() ProtoMessage { return new(VintageDecorateBoothReq) },
		func() ProtoMessage { return new(VintageDecorateBoothRsp) },
		func() ProtoMessage { return new(VintageHuntingStartGalleryReq) },
		func() ProtoMessage { return new(VintageHuntingStartGalleryRsp) },
		func() ProtoMessage { return new(VintageCampGroupBundleRegisterNotify) },
		func() ProtoMessage { return new(VintageCampStageFinishNotify) },
		func() ProtoMessage { return new(VintageMarketStartStorePlayReq) },
		func() ProtoMessage { return new(VintageMarketStartStorePlayRsp) },
		func() ProtoMessage { return new(VintageMarketFinishStorePlayReq) },
		func() ProtoMessage { return new(VintageMarketFinishStorePlayRsp) },
		func() ProtoMessage { return new(VintagePresentFinishNotify) },
		func() ProtoMessage { return new(VintageMarketStoreUnlockSlotReq) },
		func() ProtoMessage { return new(VintageMarketStoreUnlockSlotRsp) },
		func() ProtoMessage { return new(VintageMarketStoreChooseStrategyReq) },
		func() ProtoMessage { return new(VintageMarketStoreChooseStrategyRsp) },
		func() ProtoMessage { return new(VintageMarketStoreViewStrategyReq) },
		func() ProtoMessage { return new(VintageMarketStoreViewStrategyRsp) },
		func() ProtoMessage { return new(VintageMarketDividendFinishNotify) },
		func() ProtoMessage { return new(VintageMarketNpcEventFinishNotify) },
		func() ProtoMessage { return new(WindFieldRestartDungeonReq) },
		func() ProtoMessage { return new(WindFieldRestartDungeonRsp) },
		func() ProtoMessage { return new(EnterFungusFighterPlotDungeonReq) },
		func() ProtoMessage { return new(EnterFungusFighterPlotDungeonRsp) },
		func() ProtoMessage { return new(FungusFighterPlotInfoNotify) },
		func() ProtoMessage { return new(FungusCultivateReq) },
		func() ProtoMessage { return new(FungusCultivateRsp) },
		func() ProtoMessage { return new(FungusRenameReq) },
		func() ProtoMessage { return new(FungusRenameRsp) },
		func() ProtoMessage { return new(EnterFungusFighterTrainingDungeonReq) },
		func() ProtoMessage { return new(EnterFungusFighterTrainingDungeonRsp) },
		func() ProtoMessage { return new(FungusFighterRuntimeDataNotify) },
		func() ProtoMessage { return new(FungusFighterTrainingSelectFungusReq) },
		func() ProtoMessage { return new(FungusFighterTrainingSelectFungusRsp) },
		func() ProtoMessage { return new(FungusFighterTrainingGallerySettleNotify) },
		func() ProtoMessage { return new(FungusFighterClearTrainingRuntimeDataReq) },
		func() ProtoMessage { return new(FungusFighterClearTrainingRuntimeDataRsp) },
		func() ProtoMessage { return new(FungusFighterUseBackupFungusReq) },
		func() ProtoMessage { return new(FungusFighterUseBackupFungusRsp) },
		func() ProtoMessage { return new(FungusFighterRestartTraningDungeonReq) },
		func() ProtoMessage { return new(FungusFighterRestartTraningDungeonRsp) },
		func() ProtoMessage { return new(CharAmusementSettleNotify) },
		func() ProtoMessage { return new(EffigyChallengeV2EnterDungeonReq) },
		func() ProtoMessage { return new(EffigyChallengeV2EnterDungeonRsp) },
		func() ProtoMessage { return new(EffigyChallengeV2RestartDungeonReq) },
		func() ProtoMessage { return new(EffigyChallengeV2RestartDungeonRsp) },
		func() ProtoMessage { return new(EffigyChallengeV2ChooseSkillReq) },
		func() ProtoMessage { return new(EffigyChallengeV2ChooseSkillRsp) },
		func() ProtoMessage { return new(EffigyChallengeV2DungeonInfoNotify) },
		func() ProtoMessage { return new(CoinCollectChooseSkillReq) },
		func() ProtoMessage { return new(CoinCollectChooseSkillRsp) },
		func() ProtoMessage { return new(RestartCoinCollectPlaySingleModeReq) },
		func() ProtoMessage { return new(RestartCoinCollectPlaySingleModeRsp) },
		func() ProtoMessage { return new(EndCoinCollectPlaySingleModeReq) },
		func() ProtoMessage { return new(EndCoinCollectPlaySingleModeRsp) },
		func() ProtoMessage { return new(CoinCollectPrepareReq) },
		func() ProtoMessage { return new(CoinCollectPrepareRsp) },
		func() ProtoMessage { return new(CoinCollectInterruptPlayReq) },
		func() ProtoMessage { return new(CoinCollectInterruptPlayRsp) },
		func() ProtoMessage { return new(CoinCollectCheckDoubleStartPlayReq) },
		func() ProtoMessage { return new(CoinCollectCheckDoubleStartPlayRsp) },
		func() ProtoMessage { return new(SingleStartBrickBreakerReq) },
		func() ProtoMessage { return new(SingleStartBrickBreakerRsp) },
		func() ProtoMessage { return new(SingleRestartBrickBreakerReq) },
		func() ProtoMessage { return new(SingleRestartBrickBreakerRsp) },
		func() ProtoMessage { return new(BrickBreakerSettleNotify) },
		func() ProtoMessage { return new(BrickBreakerTwiceStartReq) },
		func() ProtoMessage { return new(BrickBreakerTwiceStartRsp) },
		func() ProtoMessage { return new(BrickBreakerQuitReq) },
		func() ProtoMessage { return new(BrickBreakerQuitRsp) },
		func() ProtoMessage { return new(LanV3BoatGameStartSingleReq) },
		func() ProtoMessage { return new(LanV3BoatGameStartSingleRsp) },
		func() ProtoMessage { return new(LanV3BoatGameTransferFinishNotify) },
		func() ProtoMessage { return new(LanV3RaceSettleNotify) },
		func() ProtoMessage { return new(LanV3RaceRestartReq) },
		func() ProtoMessage { return new(LanV3RaceRestartRsp) },
		func() ProtoMessage { return new(LanV3BoatInterruptSettleStageReq) },
		func() ProtoMessage { return new(LanV3BoatInterruptSettleStageRsp) },
		func() ProtoMessage { return new(LanV3ShadowFinishLevelReq) },
		func() ProtoMessage { return new(LanV3ShadowFinishLevelRsp) },
		func() ProtoMessage { return new(DuelHeartEnterDungeonReq) },
		func() ProtoMessage { return new(DuelHeartEnterDungeonRsp) },
		func() ProtoMessage { return new(DuelHeartRestartDungeonReq) },
		func() ProtoMessage { return new(DuelHeartRestartDungeonRsp) },
		func() ProtoMessage { return new(DuelHeartSelectDifficultyReq) },
		func() ProtoMessage { return new(DuelHeartSelectDifficultyRsp) },
		func() ProtoMessage { return new(DuelHeartSettleNotify) },
		func() ProtoMessage { return new(DuelHeartCgEndNotify) },
		func() ProtoMessage { return new(TeamChainEnterDungeonReq) },
		func() ProtoMessage { return new(TeamChainEnterDungeonRsp) },
		func() ProtoMessage { return new(TeamChainRestartDungeonReq) },
		func() ProtoMessage { return new(TeamChainRestartDungeonRsp) },
		func() ProtoMessage { return new(TeamChainDungeonInfoNotify) },
		func() ProtoMessage { return new(TeamChainTakeCostumeRewardReq) },
		func() ProtoMessage { return new(TeamChainTakeCostumeRewardRsp) },
		func() ProtoMessage { return new(ElectroherculesBattleSelectDifficultyReq) },
		func() ProtoMessage { return new(ElectroherculesBattleSelectDifficultyRsp) },
		func() ProtoMessage { return new(ElectroherculesBattleSettleNotify) },
	)
}

const (
	ProtoCommandGetActivityScheduleReq                          ProtoCommand = 2136
	ProtoCommandGetActivityScheduleRsp                          ProtoCommand = 2107
	ProtoCommandGetActivityInfoReq                              ProtoCommand = 2095
	ProtoCommandGetActivityInfoRsp                              ProtoCommand = 2041
	ProtoCommandActivityPlayOpenAnimNotify                      ProtoCommand = 2157
	ProtoCommandActivityInfoNotify                              ProtoCommand = 2060
	ProtoCommandActivityScheduleInfoNotify                      ProtoCommand = 2073
	ProtoCommandActivityTakeWatcherRewardReq                    ProtoCommand = 2038
	ProtoCommandActivityTakeWatcherRewardRsp                    ProtoCommand = 2034
	ProtoCommandActivityUpdateWatcherNotify                     ProtoCommand = 2156
	ProtoCommandActivitySelectAvatarCardReq                     ProtoCommand = 2028
	ProtoCommandActivitySelectAvatarCardRsp                     ProtoCommand = 2189
	ProtoCommandActivityCoinInfoNotify                          ProtoCommand = 2008
	ProtoCommandSeaLampFlyLampReq                               ProtoCommand = 2199
	ProtoCommandSeaLampFlyLampRsp                               ProtoCommand = 2192
	ProtoCommandSeaLampTakeContributionRewardReq                ProtoCommand = 2019
	ProtoCommandSeaLampTakeContributionRewardRsp                ProtoCommand = 2177
	ProtoCommandSeaLampTakePhaseRewardReq                       ProtoCommand = 2176
	ProtoCommandSeaLampTakePhaseRewardRsp                       ProtoCommand = 2190
	ProtoCommandSeaLampContributeItemReq                        ProtoCommand = 2123
	ProtoCommandSeaLampContributeItemRsp                        ProtoCommand = 2139
	ProtoCommandSeaLampFlyLampNotify                            ProtoCommand = 2105
	ProtoCommandSeaLampCoinNotify                               ProtoCommand = 2114
	ProtoCommandSeaLampPopularityNotify                         ProtoCommand = 2032
	ProtoCommandLoadActivityTerrainNotify                       ProtoCommand = 2029
	ProtoCommandServerAnnounceNotify                            ProtoCommand = 2197
	ProtoCommandServerAnnounceRevokeNotify                      ProtoCommand = 2092
	ProtoCommandActivityBannerNotify                            ProtoCommand = 2155
	ProtoCommandActivityBannerClearReq                          ProtoCommand = 2009
	ProtoCommandActivityBannerClearRsp                          ProtoCommand = 2163
	ProtoCommandSalesmanDeliverItemReq                          ProtoCommand = 2138
	ProtoCommandSalesmanDeliverItemRsp                          ProtoCommand = 2104
	ProtoCommandSalesmanTakeRewardReq                           ProtoCommand = 2191
	ProtoCommandSalesmanTakeRewardRsp                           ProtoCommand = 2110
	ProtoCommandActivityCondStateChangeNotify                   ProtoCommand = 2140
	ProtoCommandSalesmanTakeSpecialRewardReq                    ProtoCommand = 2145
	ProtoCommandSalesmanTakeSpecialRewardRsp                    ProtoCommand = 2124
	ProtoCommandGetAuthSalesmanInfoReq                          ProtoCommand = 2070
	ProtoCommandGetAuthSalesmanInfoRsp                          ProtoCommand = 2004
	ProtoCommandEnterTrialAvatarActivityDungeonReq              ProtoCommand = 2118
	ProtoCommandEnterTrialAvatarActivityDungeonRsp              ProtoCommand = 2183
	ProtoCommandReceivedTrialAvatarActivityRewardReq            ProtoCommand = 2130
	ProtoCommandReceivedTrialAvatarActivityRewardRsp            ProtoCommand = 2076
	ProtoCommandTrialAvatarFirstPassDungeonNotify               ProtoCommand = 2013
	ProtoCommandTrialAvatarInDungeonIndexNotify                 ProtoCommand = 2186
	ProtoCommandTakeDeliveryDailyRewardReq                      ProtoCommand = 2121
	ProtoCommandTakeDeliveryDailyRewardRsp                      ProtoCommand = 2162
	ProtoCommandFinishDeliveryNotify                            ProtoCommand = 2089
	ProtoCommandSelectAsterMidDifficultyReq                     ProtoCommand = 2134
	ProtoCommandSelectAsterMidDifficultyRsp                     ProtoCommand = 2180
	ProtoCommandAsterProgressInfoNotify                         ProtoCommand = 2016
	ProtoCommandAsterLittleInfoNotify                           ProtoCommand = 2068
	ProtoCommandAsterMidInfoNotify                              ProtoCommand = 2031
	ProtoCommandAsterMiscInfoNotify                             ProtoCommand = 2036
	ProtoCommandTakeAsterSpecialRewardReq                       ProtoCommand = 2097
	ProtoCommandTakeAsterSpecialRewardRsp                       ProtoCommand = 2193
	ProtoCommandAsterLargeInfoNotify                            ProtoCommand = 2146
	ProtoCommandFlightActivitySettleNotify                      ProtoCommand = 2195
	ProtoCommandFlightActivityRestartReq                        ProtoCommand = 2037
	ProtoCommandFlightActivityRestartRsp                        ProtoCommand = 2165
	ProtoCommandAsterMidCampInfoNotify                          ProtoCommand = 2133
	ProtoCommandDragonSpineChapterOpenNotify                    ProtoCommand = 2022
	ProtoCommandDragonSpineChapterProgressChangeNotify          ProtoCommand = 2065
	ProtoCommandDragonSpineChapterFinishNotify                  ProtoCommand = 2069
	ProtoCommandDragonSpineCoinChangeNotify                     ProtoCommand = 2088
	ProtoCommandActivitySaleChangeNotify                        ProtoCommand = 2071
	ProtoCommandStartEffigyChallengeReq                         ProtoCommand = 2169
	ProtoCommandStartEffigyChallengeRsp                         ProtoCommand = 2173
	ProtoCommandEffigyChallengeInfoNotify                       ProtoCommand = 2090
	ProtoCommandEffigyChallengeResultNotify                     ProtoCommand = 2046
	ProtoCommandTakeEffigyFirstPassRewardReq                    ProtoCommand = 2196
	ProtoCommandTakeEffigyFirstPassRewardRsp                    ProtoCommand = 2061
	ProtoCommandTakeEffigyRewardReq                             ProtoCommand = 2040
	ProtoCommandTakeEffigyRewardRsp                             ProtoCommand = 2007
	ProtoCommandSelectEffigyChallengeConditionReq               ProtoCommand = 2064
	ProtoCommandSelectEffigyChallengeConditionRsp               ProtoCommand = 2039
	ProtoCommandRestartEffigyChallengeReq                       ProtoCommand = 2148
	ProtoCommandRestartEffigyChallengeRsp                       ProtoCommand = 2042
	ProtoCommandTreasureMapRegionInfoNotify                     ProtoCommand = 2185
	ProtoCommandTreasureMapCurrencyNotify                       ProtoCommand = 2171
	ProtoCommandTreasureMapRegionActiveNotify                   ProtoCommand = 2122
	ProtoCommandTreasureMapMpChallengeNotify                    ProtoCommand = 2048
	ProtoCommandTreasureMapBonusChallengeNotify                 ProtoCommand = 2115
	ProtoCommandTreasureMapGuideTaskDoneNotify                  ProtoCommand = 2119
	ProtoCommandTreasureMapPreTaskDoneNotify                    ProtoCommand = 2152
	ProtoCommandBlessingScanReq                                 ProtoCommand = 2081
	ProtoCommandBlessingScanRsp                                 ProtoCommand = 2093
	ProtoCommandBlessingRedeemRewardReq                         ProtoCommand = 2137
	ProtoCommandBlessingRedeemRewardRsp                         ProtoCommand = 2098
	ProtoCommandBlessingGetFriendPicListReq                     ProtoCommand = 2043
	ProtoCommandBlessingGetFriendPicListRsp                     ProtoCommand = 2056
	ProtoCommandBlessingGiveFriendPicReq                        ProtoCommand = 2062
	ProtoCommandBlessingGiveFriendPicRsp                        ProtoCommand = 2053
	ProtoCommandBlessingAcceptGivePicReq                        ProtoCommand = 2006
	ProtoCommandBlessingAcceptGivePicRsp                        ProtoCommand = 2055
	ProtoCommandBlessingGetAllRecvPicRecordListReq              ProtoCommand = 2096
	ProtoCommandBlessingGetAllRecvPicRecordListRsp              ProtoCommand = 2083
	ProtoCommandBlessingRecvFriendPicNotify                     ProtoCommand = 2178
	ProtoCommandBlessingAcceptAllGivePicReq                     ProtoCommand = 2045
	ProtoCommandBlessingAcceptAllGivePicRsp                     ProtoCommand = 2044
	ProtoCommandExpeditionStartReq                              ProtoCommand = 2087
	ProtoCommandExpeditionStartRsp                              ProtoCommand = 2135
	ProtoCommandExpeditionRecallReq                             ProtoCommand = 2131
	ProtoCommandExpeditionRecallRsp                             ProtoCommand = 2129
	ProtoCommandExpeditionTakeRewardReq                         ProtoCommand = 2149
	ProtoCommandExpeditionTakeRewardRsp                         ProtoCommand = 2080
	ProtoCommandGetExpeditionAssistInfoListReq                  ProtoCommand = 2150
	ProtoCommandGetExpeditionAssistInfoListRsp                  ProtoCommand = 2035
	ProtoCommandSetCurExpeditionChallengeIdReq                  ProtoCommand = 2021
	ProtoCommandSetCurExpeditionChallengeIdRsp                  ProtoCommand = 2049
	ProtoCommandExpeditionChallengeEnterRegionNotify            ProtoCommand = 2154
	ProtoCommandExpeditionChallengeFinishedNotify               ProtoCommand = 2091
	ProtoCommandFleurFairBalloonSettleNotify                    ProtoCommand = 2099
	ProtoCommandFleurFairFallSettleNotify                       ProtoCommand = 2017
	ProtoCommandFleurFairMusicGameSettleReq                     ProtoCommand = 2194
	ProtoCommandFleurFairMusicGameSettleRsp                     ProtoCommand = 2113
	ProtoCommandFleurFairMusicGameStartReq                      ProtoCommand = 2167
	ProtoCommandFleurFairMusicGameStartRsp                      ProtoCommand = 2079
	ProtoCommandFleurFairReplayMiniGameReq                      ProtoCommand = 2181
	ProtoCommandFleurFairReplayMiniGameRsp                      ProtoCommand = 2052
	ProtoCommandStartArenaChallengeLevelReq                     ProtoCommand = 2127
	ProtoCommandStartArenaChallengeLevelRsp                     ProtoCommand = 2125
	ProtoCommandArenaChallengeFinishNotify                      ProtoCommand = 2030
	ProtoCommandWaterSpritePhaseFinishNotify                    ProtoCommand = 2025
	ProtoCommandActivityTakeWatcherRewardBatchReq               ProtoCommand = 2159
	ProtoCommandActivityTakeWatcherRewardBatchRsp               ProtoCommand = 2109
	ProtoCommandChannelerSlabStageActiveChallengeIndexNotify    ProtoCommand = 8734
	ProtoCommandChannelerSlabStageOneoffDungeonNotify           ProtoCommand = 8203
	ProtoCommandChannellerSlabWearBuffReq                       ProtoCommand = 8107
	ProtoCommandChannellerSlabWearBuffRsp                       ProtoCommand = 8600
	ProtoCommandChannellerSlabTakeoffBuffReq                    ProtoCommand = 8516
	ProtoCommandChannellerSlabTakeoffBuffRsp                    ProtoCommand = 8237
	ProtoCommandChannellerSlabEnterLoopDungeonReq               ProtoCommand = 8869
	ProtoCommandChannellerSlabEnterLoopDungeonRsp               ProtoCommand = 8081
	ProtoCommandChannellerSlabLoopDungeonTakeFirstPassRewardReq ProtoCommand = 8589
	ProtoCommandChannellerSlabLoopDungeonTakeFirstPassRewardRsp ProtoCommand = 8539
	ProtoCommandChannellerSlabLoopDungeonTakeScoreRewardReq     ProtoCommand = 8684
	ProtoCommandChannellerSlabLoopDungeonTakeScoreRewardRsp     ProtoCommand = 8433
	ProtoCommandChannellerSlabLoopDungeonChallengeInfoNotify    ProtoCommand = 8224
	ProtoCommandChannellerSlabLoopDungeonSelectConditionReq     ProtoCommand = 8503
	ProtoCommandChannellerSlabLoopDungeonSelectConditionRsp     ProtoCommand = 8509
	ProtoCommandChannellerSlabOneOffDungeonInfoReq              ProtoCommand = 8409
	ProtoCommandChannellerSlabOneOffDungeonInfoRsp              ProtoCommand = 8268
	ProtoCommandChannellerSlabOneOffDungeonInfoNotify           ProtoCommand = 8729
	ProtoCommandChannellerSlabSaveAssistInfoReq                 ProtoCommand = 8416
	ProtoCommandChannellerSlabSaveAssistInfoRsp                 ProtoCommand = 8932
	ProtoCommandMistTrialSelectAvatarAndEnterDungeonReq         ProtoCommand = 8666
	ProtoCommandMistTrialSelectAvatarAndEnterDungeonRsp         ProtoCommand = 8239
	ProtoCommandMistTrialGetChallengeMissionReq                 ProtoCommand = 8893
	ProtoCommandMistTrialGetChallengeMissionRsp                 ProtoCommand = 8508
	ProtoCommandMistTrialDunegonFailNotify                      ProtoCommand = 8135
	ProtoCommandChannellerSlabCheckEnterLoopDungeonReq          ProtoCommand = 8745
	ProtoCommandChannellerSlabCheckEnterLoopDungeonRsp          ProtoCommand = 8452
	ProtoCommandHideAndSeekSelectSkillReq                       ProtoCommand = 8183
	ProtoCommandHideAndSeekSelectSkillRsp                       ProtoCommand = 8088
	ProtoCommandActivityTakeScoreRewardReq                      ProtoCommand = 8971
	ProtoCommandActivityTakeScoreRewardRsp                      ProtoCommand = 8583
	ProtoCommandActivityTakeAllScoreRewardReq                   ProtoCommand = 8372
	ProtoCommandActivityTakeAllScoreRewardRsp                   ProtoCommand = 8043
	ProtoCommandHideAndSeekChooseMapReq                         ProtoCommand = 8512
	ProtoCommandHideAndSeekChooseMapRsp                         ProtoCommand = 8533
	ProtoCommandCommonPlayerTipsNotify                          ProtoCommand = 8466
	ProtoCommandFindHilichurlFinishSecondQuestNotify            ProtoCommand = 8901
	ProtoCommandFindHilichurlAcceptQuestNotify                  ProtoCommand = 8659
	ProtoCommandSummerTimeFloatSignalPositionNotify             ProtoCommand = 8077
	ProtoCommandSummerTimeFloatSignalUpdateNotify               ProtoCommand = 8781
	ProtoCommandSummerTimeSprintBoatSettleNotify                ProtoCommand = 8651
	ProtoCommandSummerTimeSprintBoatRestartReq                  ProtoCommand = 8410
	ProtoCommandSummerTimeSprintBoatRestartRsp                  ProtoCommand = 8356
	ProtoCommandStartBuoyantCombatGalleryReq                    ProtoCommand = 8732
	ProtoCommandStartBuoyantCombatGalleryRsp                    ProtoCommand = 8680
	ProtoCommandBuoyantCombatSettleNotify                       ProtoCommand = 8305
	ProtoCommandSetLimitOptimizationNotify                      ProtoCommand = 8851
	ProtoCommandEchoShellUpdateNotify                           ProtoCommand = 8150
	ProtoCommandEchoShellTakeRewardReq                          ProtoCommand = 8114
	ProtoCommandEchoShellTakeRewardRsp                          ProtoCommand = 8797
	ProtoCommandBounceConjuringSettleNotify                     ProtoCommand = 8084
	ProtoCommandBlitzRushParkourRestartReq                      ProtoCommand = 8653
	ProtoCommandBlitzRushParkourRestartRsp                      ProtoCommand = 8944
	ProtoCommandEnterChessDungeonReq                            ProtoCommand = 8191
	ProtoCommandEnterChessDungeonRsp                            ProtoCommand = 8592
	ProtoCommandTreasureMapHostInfoNotify                       ProtoCommand = 8681
	ProtoCommandSumoSaveTeamReq                                 ProtoCommand = 8313
	ProtoCommandSumoSaveTeamRsp                                 ProtoCommand = 8319
	ProtoCommandSumoSelectTeamAndEnterDungeonReq                ProtoCommand = 8215
	ProtoCommandSumoSelectTeamAndEnterDungeonRsp                ProtoCommand = 8193
	ProtoCommandSumoDungeonSettleNotify                         ProtoCommand = 8291
	ProtoCommandSumoEnterDungeonNotify                          ProtoCommand = 8013
	ProtoCommandSumoSwitchTeamReq                               ProtoCommand = 8351
	ProtoCommandSumoSwitchTeamRsp                               ProtoCommand = 8525
	ProtoCommandSumoLeaveDungeonNotify                          ProtoCommand = 8640
	ProtoCommandSumoRestartDungeonReq                           ProtoCommand = 8612
	ProtoCommandSumoRestartDungeonRsp                           ProtoCommand = 8214
	ProtoCommandActivityDisableTransferPointInteractionNotify   ProtoCommand = 8982
	ProtoCommandSumoSetNoSwitchPunishTimeNotify                 ProtoCommand = 8935
	ProtoCommandFishingGallerySettleNotify                      ProtoCommand = 8780
	ProtoCommandLunaRiteSacrificeReq                            ProtoCommand = 8805
	ProtoCommandLunaRiteSacrificeRsp                            ProtoCommand = 8080
	ProtoCommandLunaRiteTakeSacrificeRewardReq                  ProtoCommand = 8045
	ProtoCommandLunaRiteTakeSacrificeRewardRsp                  ProtoCommand = 8397
	ProtoCommandLunaRiteHintPointReq                            ProtoCommand = 8195
	ProtoCommandLunaRiteHintPointRsp                            ProtoCommand = 8765
	ProtoCommandLunaRiteHintPointRemoveNotify                   ProtoCommand = 8787
	ProtoCommandLunaRiteGroupBundleRegisterNotify               ProtoCommand = 8465
	ProtoCommandLunaRiteAreaFinishNotify                        ProtoCommand = 8213
	ProtoCommandPlantFlowerGetSeedInfoReq                       ProtoCommand = 8560
	ProtoCommandPlantFlowerGetSeedInfoRsp                       ProtoCommand = 8764
	ProtoCommandPlantFlowerTakeSeedRewardReq                    ProtoCommand = 8968
	ProtoCommandPlantFlowerTakeSeedRewardRsp                    ProtoCommand = 8860
	ProtoCommandPlantFlowerSetFlowerWishReq                     ProtoCommand = 8547
	ProtoCommandPlantFlowerSetFlowerWishRsp                     ProtoCommand = 8910
	ProtoCommandPlantFlowerGetFriendFlowerWishListReq           ProtoCommand = 8126
	ProtoCommandPlantFlowerGetFriendFlowerWishListRsp           ProtoCommand = 8511
	ProtoCommandPlantFlowerGiveFriendFlowerReq                  ProtoCommand = 8846
	ProtoCommandPlantFlowerGiveFriendFlowerRsp                  ProtoCommand = 8386
	ProtoCommandPlantFlowerGetRecvFlowerListReq                 ProtoCommand = 8270
	ProtoCommandPlantFlowerGetRecvFlowerListRsp                 ProtoCommand = 8374
	ProtoCommandPlantFlowerHaveRecvFlowerNotify                 ProtoCommand = 8078
	ProtoCommandPlantFlowerAcceptGiveFlowerReq                  ProtoCommand = 8383
	ProtoCommandPlantFlowerAcceptGiveFlowerRsp                  ProtoCommand = 8567
	ProtoCommandPlantFlowerAcceptAllGiveFlowerReq               ProtoCommand = 8808
	ProtoCommandPlantFlowerAcceptAllGiveFlowerRsp               ProtoCommand = 8888
	ProtoCommandPlantFlowerGetCanGiveFriendFlowerReq            ProtoCommand = 8716
	ProtoCommandPlantFlowerGetCanGiveFriendFlowerRsp            ProtoCommand = 8766
	ProtoCommandPlantFlowerEditFlowerCombinationReq             ProtoCommand = 8843
	ProtoCommandPlantFlowerEditFlowerCombinationRsp             ProtoCommand = 8788
	ProtoCommandMusicGameSettleReq                              ProtoCommand = 8892
	ProtoCommandMusicGameSettleRsp                              ProtoCommand = 8673
	ProtoCommandMusicGameStartReq                               ProtoCommand = 8406
	ProtoCommandMusicGameStartRsp                               ProtoCommand = 8326
	ProtoCommandDoRoguelikeDungeonCardGachaReq                  ProtoCommand = 8148
	ProtoCommandDoRoguelikeDungeonCardGachaRsp                  ProtoCommand = 8472
	ProtoCommandRefreshRoguelikeDungeonCardReq                  ProtoCommand = 8279
	ProtoCommandRefreshRoguelikeDungeonCardRsp                  ProtoCommand = 8349
	ProtoCommandSelectRoguelikeDungeonCardReq                   ProtoCommand = 8085
	ProtoCommandSelectRoguelikeDungeonCardRsp                   ProtoCommand = 8138
	ProtoCommandEquipRoguelikeRuneReq                           ProtoCommand = 8306
	ProtoCommandEquipRoguelikeRuneRsp                           ProtoCommand = 8705
	ProtoCommandTriggerRoguelikeRuneReq                         ProtoCommand = 8463
	ProtoCommandTriggerRoguelikeRuneRsp                         ProtoCommand = 8065
	ProtoCommandTriggerRoguelikeCurseNotify                     ProtoCommand = 8412
	ProtoCommandUpgradeRoguelikeShikigamiReq                    ProtoCommand = 8151
	ProtoCommandUpgradeRoguelikeShikigamiRsp                    ProtoCommand = 8966
	ProtoCommandRoguelikeSelectAvatarAndEnterDungeonReq         ProtoCommand = 8457
	ProtoCommandRoguelikeSelectAvatarAndEnterDungeonRsp         ProtoCommand = 8538
	ProtoCommandRoguelikeGiveUpReq                              ProtoCommand = 8660
	ProtoCommandRoguelikeGiveUpRsp                              ProtoCommand = 8139
	ProtoCommandRoguelikeTakeStageFirstPassRewardReq            ProtoCommand = 8421
	ProtoCommandRoguelikeTakeStageFirstPassRewardRsp            ProtoCommand = 8552
	ProtoCommandGiveUpRoguelikeDungeonCardReq                   ProtoCommand = 8353
	ProtoCommandGiveUpRoguelikeDungeonCardRsp                   ProtoCommand = 8497
	ProtoCommandEnterRoguelikeDungeonNotify                     ProtoCommand = 8652
	ProtoCommandStartRogueEliteCellChallengeReq                 ProtoCommand = 8242
	ProtoCommandStartRogueEliteCellChallengeRsp                 ProtoCommand = 8958
	ProtoCommandStartRogueNormalCellChallengeReq                ProtoCommand = 8205
	ProtoCommandStartRogueNormalCellChallengeRsp                ProtoCommand = 8036
	ProtoCommandRogueCellUpdateNotify                           ProtoCommand = 8642
	ProtoCommandRogueDungeonPlayerCellChangeNotify              ProtoCommand = 8347
	ProtoCommandRogueHealAvatarsReq                             ProtoCommand = 8947
	ProtoCommandRogueHealAvatarsRsp                             ProtoCommand = 8949
	ProtoCommandRogueResumeDungeonReq                           ProtoCommand = 8795
	ProtoCommandRogueResumeDungeonRsp                           ProtoCommand = 8647
	ProtoCommandClearRoguelikeCurseNotify                       ProtoCommand = 8207
	ProtoCommandRoguelikeCardGachaNotify                        ProtoCommand = 8925
	ProtoCommandRogueSwitchAvatarReq                            ProtoCommand = 8201
	ProtoCommandRogueSwitchAvatarRsp                            ProtoCommand = 8915
	ProtoCommandDisableRoguelikeTrapNotify                      ProtoCommand = 8259
	ProtoCommandRoguelikeRuneRecordUpdateNotify                 ProtoCommand = 8973
	ProtoCommandRoguelikeMistClearNotify                        ProtoCommand = 8324
	ProtoCommandRoguelikeEffectDataNotify                       ProtoCommand = 8222
	ProtoCommandRoguelikeEffectViewReq                          ProtoCommand = 8528
	ProtoCommandRoguelikeEffectViewRsp                          ProtoCommand = 8639
	ProtoCommandRoguelikeResourceBonusPropUpdateNotify          ProtoCommand = 8555
	ProtoCommandRoguelikeRefreshCardCostUpdateNotify            ProtoCommand = 8927
	ProtoCommandDigActivityMarkPointChangeNotify                ProtoCommand = 8109
	ProtoCommandDigActivityChangeGadgetStateReq                 ProtoCommand = 8464
	ProtoCommandDigActivityChangeGadgetStateRsp                 ProtoCommand = 8430
	ProtoCommandWinterCampStageInfoChangeNotify                 ProtoCommand = 8154
	ProtoCommandWinterCampRaceScoreNotify                       ProtoCommand = 8149
	ProtoCommandWinterCampGiveFriendItemReq                     ProtoCommand = 8572
	ProtoCommandWinterCampGiveFriendItemRsp                     ProtoCommand = 8264
	ProtoCommandWinterCampSetWishListReq                        ProtoCommand = 8753
	ProtoCommandWinterCampSetWishListRsp                        ProtoCommand = 8281
	ProtoCommandWinterCampGetFriendWishListReq                  ProtoCommand = 8946
	ProtoCommandWinterCampGetFriendWishListRsp                  ProtoCommand = 8937
	ProtoCommandWinterCampRecvItemNotify                        ProtoCommand = 8580
	ProtoCommandWinterCampAcceptGiveItemReq                     ProtoCommand = 8387
	ProtoCommandWinterCampAcceptGiveItemRsp                     ProtoCommand = 8185
	ProtoCommandWinterCampAcceptAllGiveItemReq                  ProtoCommand = 9000
	ProtoCommandWinterCampAcceptAllGiveItemRsp                  ProtoCommand = 8626
	ProtoCommandWinterCampGetCanGiveFriendItemReq               ProtoCommand = 8964
	ProtoCommandWinterCampGetCanGiveFriendItemRsp               ProtoCommand = 8357
	ProtoCommandWinterCampGetRecvItemListReq                    ProtoCommand = 8143
	ProtoCommandWinterCampGetRecvItemListRsp                    ProtoCommand = 8423
	ProtoCommandWinterCampEditSnowmanCombinationReq             ProtoCommand = 8144
	ProtoCommandWinterCampEditSnowmanCombinationRsp             ProtoCommand = 8142
	ProtoCommandWinterCampTriathlonSettleNotify                 ProtoCommand = 8342
	ProtoCommandWinterCampTakeExploreRewardReq                  ProtoCommand = 8607
	ProtoCommandWinterCampTakeExploreRewardRsp                  ProtoCommand = 8978
	ProtoCommandWinterCampTakeBattleRewardReq                   ProtoCommand = 8401
	ProtoCommandWinterCampTakeBattleRewardRsp                   ProtoCommand = 8153
	ProtoCommandWinterCampTriathlonRestartReq                   ProtoCommand = 8844
	ProtoCommandWinterCampTriathlonRestartRsp                   ProtoCommand = 8569
	ProtoCommandMistTrialSettleNotify                           ProtoCommand = 8373
	ProtoCommandMistTrialGetDungeonExhibitionDataReq            ProtoCommand = 8740
	ProtoCommandMistTrialGetDungeonExhibitionDataRsp            ProtoCommand = 8066
	ProtoCommandPotionResetChallengeReq                         ProtoCommand = 8377
	ProtoCommandPotionResetChallengeRsp                         ProtoCommand = 8067
	ProtoCommandPotionEnterDungeonReq                           ProtoCommand = 8261
	ProtoCommandPotionEnterDungeonRsp                           ProtoCommand = 8482
	ProtoCommandPotionEnterDungeonNotify                        ProtoCommand = 8531
	ProtoCommandPotionSaveDungeonResultReq                      ProtoCommand = 8192
	ProtoCommandPotionSaveDungeonResultRsp                      ProtoCommand = 8688
	ProtoCommandPotionRestartDungeonReq                         ProtoCommand = 8273
	ProtoCommandPotionRestartDungeonRsp                         ProtoCommand = 8062
	ProtoCommandTanukiTravelFinishGuideQuestNotify              ProtoCommand = 8924
	ProtoCommandFinishLanternProjectionReq                      ProtoCommand = 8704
	ProtoCommandFinishLanternProjectionRsp                      ProtoCommand = 8713
	ProtoCommandViewLanternProjectionTipsReq                    ProtoCommand = 8218
	ProtoCommandViewLanternProjectionTipsRsp                    ProtoCommand = 8590
	ProtoCommandViewLanternProjectionLevelTipsReq               ProtoCommand = 8758
	ProtoCommandViewLanternProjectionLevelTipsRsp               ProtoCommand = 8411
	ProtoCommandSalvagePreventSettleNotify                      ProtoCommand = 8231
	ProtoCommandSalvageEscortSettleNotify                       ProtoCommand = 8499
	ProtoCommandLanternRiteTakeSkinRewardReq                    ProtoCommand = 8826
	ProtoCommandLanternRiteTakeSkinRewardRsp                    ProtoCommand = 8777
	ProtoCommandSalvagePreventRestartReq                        ProtoCommand = 8367
	ProtoCommandSalvagePreventRestartRsp                        ProtoCommand = 8938
	ProtoCommandSalvageEscortRestartReq                         ProtoCommand = 8396
	ProtoCommandSalvageEscortRestartRsp                         ProtoCommand = 8959
	ProtoCommandLanternRiteStartFireworksReformReq              ProtoCommand = 8518
	ProtoCommandLanternRiteStartFireworksReformRsp              ProtoCommand = 8862
	ProtoCommandLanternRiteDoFireworksReformReq                 ProtoCommand = 8226
	ProtoCommandLanternRiteDoFireworksReformRsp                 ProtoCommand = 8657
	ProtoCommandLanternRiteEndFireworksReformReq                ProtoCommand = 8277
	ProtoCommandLanternRiteEndFireworksReformRsp                ProtoCommand = 8933
	ProtoCommandUpdateSalvageBundleMarkReq                      ProtoCommand = 8967
	ProtoCommandUpdateSalvageBundleMarkRsp                      ProtoCommand = 8459
	ProtoCommandMichiaeMatsuriDarkPressureLevelUpdateNotify     ProtoCommand = 8825
	ProtoCommandMichiaeMatsuriInteractStatueReq                 ProtoCommand = 8718
	ProtoCommandMichiaeMatsuriInteractStatueRsp                 ProtoCommand = 8449
	ProtoCommandMichiaeMatsuriUnlockCrystalSkillReq             ProtoCommand = 8345
	ProtoCommandMichiaeMatsuriUnlockCrystalSkillRsp             ProtoCommand = 8588
	ProtoCommandMichiaeMatsuriStartBossChallengeReq             ProtoCommand = 8703
	ProtoCommandMichiaeMatsuriStartBossChallengeRsp             ProtoCommand = 8426
	ProtoCommandMichiaeMatsuriStartDarkChallengeReq             ProtoCommand = 8054
	ProtoCommandMichiaeMatsuriStartDarkChallengeRsp             ProtoCommand = 8791
	ProtoCommandMichiaeMatsuriRemoveChestMarkNotify             ProtoCommand = 8726
	ProtoCommandMichiaeMatsuriRemoveChallengeMarkNotify         ProtoCommand = 8072
	ProtoCommandMichiaeMatsuriGainCrystalExpUpdateNotify        ProtoCommand = 8523
	ProtoCommandBartenderCompleteOrderReq                       ProtoCommand = 8414
	ProtoCommandBartenderCompleteOrderRsp                       ProtoCommand = 8125
	ProtoCommandBartenderCancelOrderReq                         ProtoCommand = 8442
	ProtoCommandBartenderCancelOrderRsp                         ProtoCommand = 8837
	ProtoCommandBartenderGetFormulaReq                          ProtoCommand = 8462
	ProtoCommandBartenderGetFormulaRsp                          ProtoCommand = 8842
	ProtoCommandBartenderStartLevelReq                          ProtoCommand = 8507
	ProtoCommandBartenderStartLevelRsp                          ProtoCommand = 8402
	ProtoCommandBartenderCancelLevelReq                         ProtoCommand = 8771
	ProtoCommandBartenderCancelLevelRsp                         ProtoCommand = 8686
	ProtoCommandBartenderLevelProgressNotify                    ProtoCommand = 8756
	ProtoCommandBartenderFinishLevelReq                         ProtoCommand = 8227
	ProtoCommandBartenderFinishLevelRsp                         ProtoCommand = 8093
	ProtoCommandCrystalLinkEnterDungeonReq                      ProtoCommand = 8325
	ProtoCommandCrystalLinkEnterDungeonRsp                      ProtoCommand = 8147
	ProtoCommandCrystalLinkDungeonInfoNotify                    ProtoCommand = 8858
	ProtoCommandCrystalLinkRestartDungeonReq                    ProtoCommand = 8022
	ProtoCommandCrystalLinkRestartDungeonRsp                    ProtoCommand = 8119
	ProtoCommandQuickOpenActivityReq                            ProtoCommand = 8178
	ProtoCommandQuickOpenActivityRsp                            ProtoCommand = 8882
	ProtoCommandIrodoriEditFlowerCombinationReq                 ProtoCommand = 8608
	ProtoCommandIrodoriEditFlowerCombinationRsp                 ProtoCommand = 8833
	ProtoCommandIrodoriScanEntityReq                            ProtoCommand = 8767
	ProtoCommandIrodoriScanEntityRsp                            ProtoCommand = 8026
	ProtoCommandIrodoriFillPoetryReq                            ProtoCommand = 8129
	ProtoCommandIrodoriFillPoetryRsp                            ProtoCommand = 8880
	ProtoCommandIrodoriChessEquipCardReq                        ProtoCommand = 8561
	ProtoCommandIrodoriChessEquipCardRsp                        ProtoCommand = 8308
	ProtoCommandIrodoriChessUnequipCardReq                      ProtoCommand = 8057
	ProtoCommandIrodoriChessUnequipCardRsp                      ProtoCommand = 8817
	ProtoCommandEnterIrodoriChessDungeonReq                     ProtoCommand = 8717
	ProtoCommandEnterIrodoriChessDungeonRsp                     ProtoCommand = 8546
	ProtoCommandIrodoriMasterStartGalleryReq                    ProtoCommand = 8165
	ProtoCommandIrodoriMasterStartGalleryRsp                    ProtoCommand = 8381
	ProtoCommandIrodoriMasterGalleryCgEndNotify                 ProtoCommand = 8061
	ProtoCommandIrodoriMasterGallerySettleNotify                ProtoCommand = 8340
	ProtoCommandPhotoActivityFinishReq                          ProtoCommand = 8921
	ProtoCommandPhotoActivityFinishRsp                          ProtoCommand = 8854
	ProtoCommandPhotoActivityClientViewReq                      ProtoCommand = 8709
	ProtoCommandPhotoActivityClientViewRsp                      ProtoCommand = 8983
	ProtoCommandSpiceActivityFinishMakeSpiceReq                 ProtoCommand = 8096
	ProtoCommandSpiceActivityFinishMakeSpiceRsp                 ProtoCommand = 8481
	ProtoCommandSpiceActivityProcessFoodReq                     ProtoCommand = 8216
	ProtoCommandSpiceActivityProcessFoodRsp                     ProtoCommand = 8772
	ProtoCommandSpiceActivityGivingRecordNotify                 ProtoCommand = 8407
	ProtoCommandGachaActivityPercentNotify                      ProtoCommand = 8450
	ProtoCommandGachaActivityUpdateElemNotify                   ProtoCommand = 8919
	ProtoCommandGachaActivityCreateRobotReq                     ProtoCommand = 8614
	ProtoCommandGachaActivityCreateRobotRsp                     ProtoCommand = 8610
	ProtoCommandGachaActivityTakeRewardReq                      ProtoCommand = 8930
	ProtoCommandGachaActivityTakeRewardRsp                      ProtoCommand = 8768
	ProtoCommandGachaActivityResetReq                           ProtoCommand = 8163
	ProtoCommandGachaActivityResetRsp                           ProtoCommand = 8240
	ProtoCommandGachaActivityNextStageReq                       ProtoCommand = 8257
	ProtoCommandGachaActivityNextStageRsp                       ProtoCommand = 8918
	ProtoCommandActivityGiveFriendGiftReq                       ProtoCommand = 8233
	ProtoCommandActivityGiveFriendGiftRsp                       ProtoCommand = 8696
	ProtoCommandActivityGetRecvGiftListReq                      ProtoCommand = 8725
	ProtoCommandActivityGetRecvGiftListRsp                      ProtoCommand = 8120
	ProtoCommandActivityHaveRecvGiftNotify                      ProtoCommand = 8733
	ProtoCommandActivityAcceptGiveGiftReq                       ProtoCommand = 8095
	ProtoCommandActivityAcceptGiveGiftRsp                       ProtoCommand = 8502
	ProtoCommandActivityAcceptAllGiveGiftReq                    ProtoCommand = 8113
	ProtoCommandActivityAcceptAllGiveGiftRsp                    ProtoCommand = 8132
	ProtoCommandActivityGetCanGiveFriendGiftReq                 ProtoCommand = 8559
	ProtoCommandActivityGetCanGiveFriendGiftRsp                 ProtoCommand = 8848
	ProtoCommandActivitySetGiftWishReq                          ProtoCommand = 8017
	ProtoCommandActivitySetGiftWishRsp                          ProtoCommand = 8554
	ProtoCommandActivityGetFriendGiftWishListReq                ProtoCommand = 8806
	ProtoCommandActivityGetFriendGiftWishListRsp                ProtoCommand = 8253
	ProtoCommandLuminanceStoneChallengeSettleNotify             ProtoCommand = 8186
	ProtoCommandStartRogueDiaryPlayReq                          ProtoCommand = 8419
	ProtoCommandStartRogueDiaryPlayRsp                          ProtoCommand = 8385
	ProtoCommandResetRogueDiaryPlayReq                          ProtoCommand = 8127
	ProtoCommandResetRogueDiaryPlayRsp                          ProtoCommand = 8948
	ProtoCommandEnterRogueDiaryDungeonReq                       ProtoCommand = 8943
	ProtoCommandEnterRogueDiaryDungeonRsp                       ProtoCommand = 8352
	ProtoCommandResumeRogueDiaryDungeonReq                      ProtoCommand = 8838
	ProtoCommandResumeRogueDiaryDungeonRsp                      ProtoCommand = 8989
	ProtoCommandRogueDiaryDungeonInfoNotify                     ProtoCommand = 8597
	ProtoCommandRogueDiaryDungeonSettleNotify                   ProtoCommand = 8895
	ProtoCommandStartRogueDiaryRoomReq                          ProtoCommand = 8159
	ProtoCommandStartRogueDiaryRoomRsp                          ProtoCommand = 8793
	ProtoCommandRogueDiaryTiredAvatarNotify                     ProtoCommand = 8514
	ProtoCommandReserveRogueDiaryAvatarReq                      ProtoCommand = 8748
	ProtoCommandReserveRogueDiaryAvatarRsp                      ProtoCommand = 8799
	ProtoCommandGetRogueDairyRepairInfoReq                      ProtoCommand = 8014
	ProtoCommandGetRogueDairyRepairInfoRsp                      ProtoCommand = 8447
	ProtoCommandRefreshRogueDiaryCardReq                        ProtoCommand = 8991
	ProtoCommandRefreshRogueDiaryCardRsp                        ProtoCommand = 8028
	ProtoCommandRogueFinishRepairReq                            ProtoCommand = 8363
	ProtoCommandRogueFinishRepairRsp                            ProtoCommand = 8535
	ProtoCommandTryInterruptRogueDiaryDungeonReq                ProtoCommand = 8617
	ProtoCommandTryInterruptRogueDiaryDungeonRsp                ProtoCommand = 8903
	ProtoCommandRogueDiaryRepairInfoNotify                      ProtoCommand = 8641
	ProtoCommandRetryCurRogueDiaryDungeonReq                    ProtoCommand = 8398
	ProtoCommandRetryCurRogueDiaryDungeonRsp                    ProtoCommand = 8334
	ProtoCommandRogueDiaryReviveAvatarReq                       ProtoCommand = 8038
	ProtoCommandRogueDiaryReviveAvatarRsp                       ProtoCommand = 8343
	ProtoCommandTryEnterNextRogueDiaryDungeonReq                ProtoCommand = 8280
	ProtoCommandTryEnterNextRogueDiaryDungeonRsp                ProtoCommand = 8362
	ProtoCommandRogueDiaryCoinAddNotify                         ProtoCommand = 8602
	ProtoCommandSummerTimeV2BoatSettleNotify                    ProtoCommand = 8870
	ProtoCommandActivityPushTipsInfoNotify                      ProtoCommand = 8513
	ProtoCommandActivityReadPushTipsReq                         ProtoCommand = 8145
	ProtoCommandActivityReadPushTipsRsp                         ProtoCommand = 8574
	ProtoCommandSummerTimeV2RestartBoatGalleryReq               ProtoCommand = 8476
	ProtoCommandSummerTimeV2RestartBoatGalleryRsp               ProtoCommand = 8004
	ProtoCommandSummerTimeV2RestartDungeonReq                   ProtoCommand = 8346
	ProtoCommandSummerTimeV2RestartDungeonRsp                   ProtoCommand = 8996
	ProtoCommandIslandPartySettleNotify                         ProtoCommand = 24601
	ProtoCommandGearActivityStartPlayGearReq                    ProtoCommand = 23467
	ProtoCommandGearActivityStartPlayGearRsp                    ProtoCommand = 21025
	ProtoCommandGearActivityFinishPlayGearReq                   ProtoCommand = 21834
	ProtoCommandGearActivityFinishPlayGearRsp                   ProtoCommand = 21800
	ProtoCommandGearActivityStartPlayPictureReq                 ProtoCommand = 24550
	ProtoCommandGearActivityStartPlayPictureRsp                 ProtoCommand = 23388
	ProtoCommandGearActivityFinishPlayPictureReq                ProtoCommand = 21054
	ProtoCommandGearActivityFinishPlayPictureRsp                ProtoCommand = 21851
	ProtoCommandGravenInnocenceRaceSettleNotify                 ProtoCommand = 20681
	ProtoCommandGravenInnocenceRaceRestartReq                   ProtoCommand = 22882
	ProtoCommandGravenInnocenceRaceRestartRsp                   ProtoCommand = 21880
	ProtoCommandGravenInnocenceEditCarveCombinationReq          ProtoCommand = 23107
	ProtoCommandGravenInnocenceEditCarveCombinationRsp          ProtoCommand = 20702
	ProtoCommandGravenInnocencePhotoFinishReq                   ProtoCommand = 21750
	ProtoCommandGravenInnocencePhotoFinishRsp                   ProtoCommand = 23948
	ProtoCommandGravenInnocencePhotoReminderNotify              ProtoCommand = 23864
	ProtoCommandInstableSprayEnterDungeonReq                    ProtoCommand = 24312
	ProtoCommandInstableSprayEnterDungeonRsp                    ProtoCommand = 23381
	ProtoCommandInstableSpraySwitchTeamReq                      ProtoCommand = 24857
	ProtoCommandInstableSpraySwitchTeamRsp                      ProtoCommand = 24152
	ProtoCommandInstableSprayLevelFinishNotify                  ProtoCommand = 21961
	ProtoCommandInstableSprayRestartDungeonReq                  ProtoCommand = 23678
	ProtoCommandInstableSprayRestartDungeonRsp                  ProtoCommand = 24923
	ProtoCommandMuqadasPotionActivityEnterDungeonReq            ProtoCommand = 24602
	ProtoCommandMuqadasPotionActivityEnterDungeonRsp            ProtoCommand = 21804
	ProtoCommandMuqadasPotionDungeonSettleNotify                ProtoCommand = 20005
	ProtoCommandMuqadasPotionRestartDungeonReq                  ProtoCommand = 22391
	ProtoCommandMuqadasPotionRestartDungeonRsp                  ProtoCommand = 21208
	ProtoCommandMuqadasPotionCaptureWeaknessReq                 ProtoCommand = 20011
	ProtoCommandMuqadasPotionCaptureWeaknessRsp                 ProtoCommand = 24081
	ProtoCommandTreasureSeelieCollectOrbsNotify                 ProtoCommand = 20754
	ProtoCommandVintageMarketDeliverItemReq                     ProtoCommand = 23141
	ProtoCommandVintageMarketDeliverItemRsp                     ProtoCommand = 22181
	ProtoCommandSceneGalleryVintageHuntingSettleNotify          ProtoCommand = 20324
	ProtoCommandVintagePresentFinishNoify                       ProtoCommand = 24142
	ProtoCommandVintageDecorateBoothReq                         ProtoCommand = 20846
	ProtoCommandVintageDecorateBoothRsp                         ProtoCommand = 20993
	ProtoCommandVintageHuntingStartGalleryReq                   ProtoCommand = 21780
	ProtoCommandVintageHuntingStartGalleryRsp                   ProtoCommand = 21951
	ProtoCommandVintageCampGroupBundleRegisterNotify            ProtoCommand = 24244
	ProtoCommandVintageCampStageFinishNotify                    ProtoCommand = 22830
	ProtoCommandVintageMarketStartStorePlayReq                  ProtoCommand = 22864
	ProtoCommandVintageMarketStartStorePlayRsp                  ProtoCommand = 22130
	ProtoCommandVintageMarketFinishStorePlayReq                 ProtoCommand = 20676
	ProtoCommandVintageMarketFinishStorePlayRsp                 ProtoCommand = 23462
	ProtoCommandVintagePresentFinishNotify                      ProtoCommand = 20086
	ProtoCommandVintageMarketStoreUnlockSlotReq                 ProtoCommand = 20626
	ProtoCommandVintageMarketStoreUnlockSlotRsp                 ProtoCommand = 20733
	ProtoCommandVintageMarketStoreChooseStrategyReq             ProtoCommand = 21248
	ProtoCommandVintageMarketStoreChooseStrategyRsp             ProtoCommand = 24860
	ProtoCommandVintageMarketStoreViewStrategyReq               ProtoCommand = 21700
	ProtoCommandVintageMarketStoreViewStrategyRsp               ProtoCommand = 21814
	ProtoCommandVintageMarketDividendFinishNotify               ProtoCommand = 23147
	ProtoCommandVintageMarketNpcEventFinishNotify               ProtoCommand = 24201
	ProtoCommandWindFieldRestartDungeonReq                      ProtoCommand = 20731
	ProtoCommandWindFieldRestartDungeonRsp                      ProtoCommand = 24712
	ProtoCommandEnterFungusFighterPlotDungeonReq                ProtoCommand = 23053
	ProtoCommandEnterFungusFighterPlotDungeonRsp                ProtoCommand = 21008
	ProtoCommandFungusFighterPlotInfoNotify                     ProtoCommand = 22174
	ProtoCommandFungusCultivateReq                              ProtoCommand = 21749
	ProtoCommandFungusCultivateRsp                              ProtoCommand = 23532
	ProtoCommandFungusRenameReq                                 ProtoCommand = 22006
	ProtoCommandFungusRenameRsp                                 ProtoCommand = 20066
	ProtoCommandEnterFungusFighterTrainingDungeonReq            ProtoCommand = 23860
	ProtoCommandEnterFungusFighterTrainingDungeonRsp            ProtoCommand = 21593
	ProtoCommandFungusFighterRuntimeDataNotify                  ProtoCommand = 24674
	ProtoCommandFungusFighterTrainingSelectFungusReq            ProtoCommand = 23903
	ProtoCommandFungusFighterTrainingSelectFungusRsp            ProtoCommand = 21570
	ProtoCommandFungusFighterTrainingGallerySettleNotify        ProtoCommand = 23931
	ProtoCommandFungusFighterClearTrainingRuntimeDataReq        ProtoCommand = 24137
	ProtoCommandFungusFighterClearTrainingRuntimeDataRsp        ProtoCommand = 22991
	ProtoCommandFungusFighterUseBackupFungusReq                 ProtoCommand = 21266
	ProtoCommandFungusFighterUseBackupFungusRsp                 ProtoCommand = 23428
	ProtoCommandFungusFighterRestartTraningDungeonReq           ProtoCommand = 23980
	ProtoCommandFungusFighterRestartTraningDungeonRsp           ProtoCommand = 22890
	ProtoCommandCharAmusementSettleNotify                       ProtoCommand = 23133
	ProtoCommandEffigyChallengeV2EnterDungeonReq                ProtoCommand = 23489
	ProtoCommandEffigyChallengeV2EnterDungeonRsp                ProtoCommand = 24917
	ProtoCommandEffigyChallengeV2RestartDungeonReq              ProtoCommand = 24522
	ProtoCommandEffigyChallengeV2RestartDungeonRsp              ProtoCommand = 23167
	ProtoCommandEffigyChallengeV2ChooseSkillReq                 ProtoCommand = 21269
	ProtoCommandEffigyChallengeV2ChooseSkillRsp                 ProtoCommand = 22448
	ProtoCommandEffigyChallengeV2DungeonInfoNotify              ProtoCommand = 22835
	ProtoCommandCoinCollectChooseSkillReq                       ProtoCommand = 21214
	ProtoCommandCoinCollectChooseSkillRsp                       ProtoCommand = 24700
	ProtoCommandRestartCoinCollectPlaySingleModeReq             ProtoCommand = 22367
	ProtoCommandRestartCoinCollectPlaySingleModeRsp             ProtoCommand = 24828
	ProtoCommandEndCoinCollectPlaySingleModeReq                 ProtoCommand = 23216
	ProtoCommandEndCoinCollectPlaySingleModeRsp                 ProtoCommand = 20302
	ProtoCommandCoinCollectPrepareReq                           ProtoCommand = 21718
	ProtoCommandCoinCollectPrepareRsp                           ProtoCommand = 20930
	ProtoCommandCoinCollectInterruptPlayReq                     ProtoCommand = 22934
	ProtoCommandCoinCollectInterruptPlayRsp                     ProtoCommand = 22509
	ProtoCommandCoinCollectCheckDoubleStartPlayReq              ProtoCommand = 21294
	ProtoCommandCoinCollectCheckDoubleStartPlayRsp              ProtoCommand = 21244
	ProtoCommandSingleStartBrickBreakerReq                      ProtoCommand = 23113
	ProtoCommandSingleStartBrickBreakerRsp                      ProtoCommand = 20622
	ProtoCommandSingleRestartBrickBreakerReq                    ProtoCommand = 22704
	ProtoCommandSingleRestartBrickBreakerRsp                    ProtoCommand = 22112
	ProtoCommandBrickBreakerSettleNotify                        ProtoCommand = 24616
	ProtoCommandBrickBreakerTwiceStartReq                       ProtoCommand = 24028
	ProtoCommandBrickBreakerTwiceStartRsp                       ProtoCommand = 22591
	ProtoCommandBrickBreakerQuitReq                             ProtoCommand = 20137
	ProtoCommandBrickBreakerQuitRsp                             ProtoCommand = 20667
	ProtoCommandLanV3BoatGameStartSingleReq                     ProtoCommand = 23637
	ProtoCommandLanV3BoatGameStartSingleRsp                     ProtoCommand = 22069
	ProtoCommandLanV3BoatGameTransferFinishNotify               ProtoCommand = 20683
	ProtoCommandLanV3RaceSettleNotify                           ProtoCommand = 24629
	ProtoCommandLanV3RaceRestartReq                             ProtoCommand = 20331
	ProtoCommandLanV3RaceRestartRsp                             ProtoCommand = 23477
	ProtoCommandLanV3BoatInterruptSettleStageReq                ProtoCommand = 20951
	ProtoCommandLanV3BoatInterruptSettleStageRsp                ProtoCommand = 24759
	ProtoCommandLanV3ShadowFinishLevelReq                       ProtoCommand = 20227
	ProtoCommandLanV3ShadowFinishLevelRsp                       ProtoCommand = 20480
	ProtoCommandDuelHeartEnterDungeonReq                        ProtoCommand = 20076
	ProtoCommandDuelHeartEnterDungeonRsp                        ProtoCommand = 24080
	ProtoCommandDuelHeartRestartDungeonReq                      ProtoCommand = 23780
	ProtoCommandDuelHeartRestartDungeonRsp                      ProtoCommand = 24816
	ProtoCommandDuelHeartSelectDifficultyReq                    ProtoCommand = 22297
	ProtoCommandDuelHeartSelectDifficultyRsp                    ProtoCommand = 22312
	ProtoCommandDuelHeartSettleNotify                           ProtoCommand = 22250
	ProtoCommandDuelHeartCgEndNotify                            ProtoCommand = 20093
	ProtoCommandTeamChainEnterDungeonReq                        ProtoCommand = 20636
	ProtoCommandTeamChainEnterDungeonRsp                        ProtoCommand = 21776
	ProtoCommandTeamChainRestartDungeonReq                      ProtoCommand = 22269
	ProtoCommandTeamChainRestartDungeonRsp                      ProtoCommand = 21313
	ProtoCommandTeamChainDungeonInfoNotify                      ProtoCommand = 24946
	ProtoCommandTeamChainTakeCostumeRewardReq                   ProtoCommand = 21207
	ProtoCommandTeamChainTakeCostumeRewardRsp                   ProtoCommand = 24258
	ProtoCommandElectroherculesBattleSelectDifficultyReq        ProtoCommand = 24212
	ProtoCommandElectroherculesBattleSelectDifficultyRsp        ProtoCommand = 23684
	ProtoCommandElectroherculesBattleSettleNotify               ProtoCommand = 21083
)

func (*GetActivityScheduleReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetActivityScheduleReq }
func (*GetActivityScheduleReq) ProtoMessageType() ProtoMessageType { return "GetActivityScheduleReq" }

func (*GetActivityScheduleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetActivityScheduleRsp }
func (*GetActivityScheduleRsp) ProtoMessageType() ProtoMessageType { return "GetActivityScheduleRsp" }

func (*GetActivityInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetActivityInfoReq }
func (*GetActivityInfoReq) ProtoMessageType() ProtoMessageType { return "GetActivityInfoReq" }

func (*GetActivityInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetActivityInfoRsp }
func (*GetActivityInfoRsp) ProtoMessageType() ProtoMessageType { return "GetActivityInfoRsp" }

func (*ActivityPlayOpenAnimNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityPlayOpenAnimNotify
}
func (*ActivityPlayOpenAnimNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityPlayOpenAnimNotify"
}

func (*ActivityInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandActivityInfoNotify }
func (*ActivityInfoNotify) ProtoMessageType() ProtoMessageType { return "ActivityInfoNotify" }

func (*ActivityScheduleInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityScheduleInfoNotify
}
func (*ActivityScheduleInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityScheduleInfoNotify"
}

func (*ActivityTakeWatcherRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeWatcherRewardReq
}
func (*ActivityTakeWatcherRewardReq) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeWatcherRewardReq"
}

func (*ActivityTakeWatcherRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeWatcherRewardRsp
}
func (*ActivityTakeWatcherRewardRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeWatcherRewardRsp"
}

func (*ActivityUpdateWatcherNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityUpdateWatcherNotify
}
func (*ActivityUpdateWatcherNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityUpdateWatcherNotify"
}

func (*ActivitySelectAvatarCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivitySelectAvatarCardReq
}
func (*ActivitySelectAvatarCardReq) ProtoMessageType() ProtoMessageType {
	return "ActivitySelectAvatarCardReq"
}

func (*ActivitySelectAvatarCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivitySelectAvatarCardRsp
}
func (*ActivitySelectAvatarCardRsp) ProtoMessageType() ProtoMessageType {
	return "ActivitySelectAvatarCardRsp"
}

func (*ActivityCoinInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandActivityCoinInfoNotify }
func (*ActivityCoinInfoNotify) ProtoMessageType() ProtoMessageType { return "ActivityCoinInfoNotify" }

func (*SeaLampFlyLampReq) ProtoCommand() ProtoCommand         { return ProtoCommandSeaLampFlyLampReq }
func (*SeaLampFlyLampReq) ProtoMessageType() ProtoMessageType { return "SeaLampFlyLampReq" }

func (*SeaLampFlyLampRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSeaLampFlyLampRsp }
func (*SeaLampFlyLampRsp) ProtoMessageType() ProtoMessageType { return "SeaLampFlyLampRsp" }

func (*SeaLampTakeContributionRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampTakeContributionRewardReq
}
func (*SeaLampTakeContributionRewardReq) ProtoMessageType() ProtoMessageType {
	return "SeaLampTakeContributionRewardReq"
}

func (*SeaLampTakeContributionRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampTakeContributionRewardRsp
}
func (*SeaLampTakeContributionRewardRsp) ProtoMessageType() ProtoMessageType {
	return "SeaLampTakeContributionRewardRsp"
}

func (*SeaLampTakePhaseRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampTakePhaseRewardReq
}
func (*SeaLampTakePhaseRewardReq) ProtoMessageType() ProtoMessageType {
	return "SeaLampTakePhaseRewardReq"
}

func (*SeaLampTakePhaseRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampTakePhaseRewardRsp
}
func (*SeaLampTakePhaseRewardRsp) ProtoMessageType() ProtoMessageType {
	return "SeaLampTakePhaseRewardRsp"
}

func (*SeaLampContributeItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampContributeItemReq
}
func (*SeaLampContributeItemReq) ProtoMessageType() ProtoMessageType {
	return "SeaLampContributeItemReq"
}

func (*SeaLampContributeItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampContributeItemRsp
}
func (*SeaLampContributeItemRsp) ProtoMessageType() ProtoMessageType {
	return "SeaLampContributeItemRsp"
}

func (*SeaLampFlyLampNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSeaLampFlyLampNotify }
func (*SeaLampFlyLampNotify) ProtoMessageType() ProtoMessageType { return "SeaLampFlyLampNotify" }

func (*SeaLampCoinNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSeaLampCoinNotify }
func (*SeaLampCoinNotify) ProtoMessageType() ProtoMessageType { return "SeaLampCoinNotify" }

func (*SeaLampPopularityNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSeaLampPopularityNotify
}
func (*SeaLampPopularityNotify) ProtoMessageType() ProtoMessageType { return "SeaLampPopularityNotify" }

func (*LoadActivityTerrainNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLoadActivityTerrainNotify
}
func (*LoadActivityTerrainNotify) ProtoMessageType() ProtoMessageType {
	return "LoadActivityTerrainNotify"
}

func (*ServerAnnounceNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerAnnounceNotify }
func (*ServerAnnounceNotify) ProtoMessageType() ProtoMessageType { return "ServerAnnounceNotify" }

func (*ServerAnnounceRevokeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerAnnounceRevokeNotify
}
func (*ServerAnnounceRevokeNotify) ProtoMessageType() ProtoMessageType {
	return "ServerAnnounceRevokeNotify"
}

func (*ActivityBannerNotify) ProtoCommand() ProtoCommand         { return ProtoCommandActivityBannerNotify }
func (*ActivityBannerNotify) ProtoMessageType() ProtoMessageType { return "ActivityBannerNotify" }

func (*ActivityBannerClearReq) ProtoCommand() ProtoCommand         { return ProtoCommandActivityBannerClearReq }
func (*ActivityBannerClearReq) ProtoMessageType() ProtoMessageType { return "ActivityBannerClearReq" }

func (*ActivityBannerClearRsp) ProtoCommand() ProtoCommand         { return ProtoCommandActivityBannerClearRsp }
func (*ActivityBannerClearRsp) ProtoMessageType() ProtoMessageType { return "ActivityBannerClearRsp" }

func (*SalesmanDeliverItemReq) ProtoCommand() ProtoCommand         { return ProtoCommandSalesmanDeliverItemReq }
func (*SalesmanDeliverItemReq) ProtoMessageType() ProtoMessageType { return "SalesmanDeliverItemReq" }

func (*SalesmanDeliverItemRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSalesmanDeliverItemRsp }
func (*SalesmanDeliverItemRsp) ProtoMessageType() ProtoMessageType { return "SalesmanDeliverItemRsp" }

func (*SalesmanTakeRewardReq) ProtoCommand() ProtoCommand         { return ProtoCommandSalesmanTakeRewardReq }
func (*SalesmanTakeRewardReq) ProtoMessageType() ProtoMessageType { return "SalesmanTakeRewardReq" }

func (*SalesmanTakeRewardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSalesmanTakeRewardRsp }
func (*SalesmanTakeRewardRsp) ProtoMessageType() ProtoMessageType { return "SalesmanTakeRewardRsp" }

func (*ActivityCondStateChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityCondStateChangeNotify
}
func (*ActivityCondStateChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityCondStateChangeNotify"
}

func (*SalesmanTakeSpecialRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSalesmanTakeSpecialRewardReq
}
func (*SalesmanTakeSpecialRewardReq) ProtoMessageType() ProtoMessageType {
	return "SalesmanTakeSpecialRewardReq"
}

func (*SalesmanTakeSpecialRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSalesmanTakeSpecialRewardRsp
}
func (*SalesmanTakeSpecialRewardRsp) ProtoMessageType() ProtoMessageType {
	return "SalesmanTakeSpecialRewardRsp"
}

func (*GetAuthSalesmanInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetAuthSalesmanInfoReq }
func (*GetAuthSalesmanInfoReq) ProtoMessageType() ProtoMessageType { return "GetAuthSalesmanInfoReq" }

func (*GetAuthSalesmanInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetAuthSalesmanInfoRsp }
func (*GetAuthSalesmanInfoRsp) ProtoMessageType() ProtoMessageType { return "GetAuthSalesmanInfoRsp" }

func (*EnterTrialAvatarActivityDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterTrialAvatarActivityDungeonReq
}
func (*EnterTrialAvatarActivityDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EnterTrialAvatarActivityDungeonReq"
}

func (*EnterTrialAvatarActivityDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterTrialAvatarActivityDungeonRsp
}
func (*EnterTrialAvatarActivityDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EnterTrialAvatarActivityDungeonRsp"
}

func (*ReceivedTrialAvatarActivityRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandReceivedTrialAvatarActivityRewardReq
}
func (*ReceivedTrialAvatarActivityRewardReq) ProtoMessageType() ProtoMessageType {
	return "ReceivedTrialAvatarActivityRewardReq"
}

func (*ReceivedTrialAvatarActivityRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandReceivedTrialAvatarActivityRewardRsp
}
func (*ReceivedTrialAvatarActivityRewardRsp) ProtoMessageType() ProtoMessageType {
	return "ReceivedTrialAvatarActivityRewardRsp"
}

func (*TrialAvatarFirstPassDungeonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTrialAvatarFirstPassDungeonNotify
}
func (*TrialAvatarFirstPassDungeonNotify) ProtoMessageType() ProtoMessageType {
	return "TrialAvatarFirstPassDungeonNotify"
}

func (*TrialAvatarInDungeonIndexNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTrialAvatarInDungeonIndexNotify
}
func (*TrialAvatarInDungeonIndexNotify) ProtoMessageType() ProtoMessageType {
	return "TrialAvatarInDungeonIndexNotify"
}

func (*TakeDeliveryDailyRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeDeliveryDailyRewardReq
}
func (*TakeDeliveryDailyRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeDeliveryDailyRewardReq"
}

func (*TakeDeliveryDailyRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeDeliveryDailyRewardRsp
}
func (*TakeDeliveryDailyRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeDeliveryDailyRewardRsp"
}

func (*FinishDeliveryNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFinishDeliveryNotify }
func (*FinishDeliveryNotify) ProtoMessageType() ProtoMessageType { return "FinishDeliveryNotify" }

func (*SelectAsterMidDifficultyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSelectAsterMidDifficultyReq
}
func (*SelectAsterMidDifficultyReq) ProtoMessageType() ProtoMessageType {
	return "SelectAsterMidDifficultyReq"
}

func (*SelectAsterMidDifficultyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSelectAsterMidDifficultyRsp
}
func (*SelectAsterMidDifficultyRsp) ProtoMessageType() ProtoMessageType {
	return "SelectAsterMidDifficultyRsp"
}

func (*AsterProgressInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAsterProgressInfoNotify
}
func (*AsterProgressInfoNotify) ProtoMessageType() ProtoMessageType { return "AsterProgressInfoNotify" }

func (*AsterLittleInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAsterLittleInfoNotify }
func (*AsterLittleInfoNotify) ProtoMessageType() ProtoMessageType { return "AsterLittleInfoNotify" }

func (*AsterMidInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAsterMidInfoNotify }
func (*AsterMidInfoNotify) ProtoMessageType() ProtoMessageType { return "AsterMidInfoNotify" }

func (*AsterMiscInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAsterMiscInfoNotify }
func (*AsterMiscInfoNotify) ProtoMessageType() ProtoMessageType { return "AsterMiscInfoNotify" }

func (*TakeAsterSpecialRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeAsterSpecialRewardReq
}
func (*TakeAsterSpecialRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeAsterSpecialRewardReq"
}

func (*TakeAsterSpecialRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeAsterSpecialRewardRsp
}
func (*TakeAsterSpecialRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeAsterSpecialRewardRsp"
}

func (*AsterLargeInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAsterLargeInfoNotify }
func (*AsterLargeInfoNotify) ProtoMessageType() ProtoMessageType { return "AsterLargeInfoNotify" }

func (*FlightActivitySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFlightActivitySettleNotify
}
func (*FlightActivitySettleNotify) ProtoMessageType() ProtoMessageType {
	return "FlightActivitySettleNotify"
}

func (*FlightActivityRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFlightActivityRestartReq
}
func (*FlightActivityRestartReq) ProtoMessageType() ProtoMessageType {
	return "FlightActivityRestartReq"
}

func (*FlightActivityRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFlightActivityRestartRsp
}
func (*FlightActivityRestartRsp) ProtoMessageType() ProtoMessageType {
	return "FlightActivityRestartRsp"
}

func (*AsterMidCampInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAsterMidCampInfoNotify }
func (*AsterMidCampInfoNotify) ProtoMessageType() ProtoMessageType { return "AsterMidCampInfoNotify" }

func (*DragonSpineChapterOpenNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDragonSpineChapterOpenNotify
}
func (*DragonSpineChapterOpenNotify) ProtoMessageType() ProtoMessageType {
	return "DragonSpineChapterOpenNotify"
}

func (*DragonSpineChapterProgressChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDragonSpineChapterProgressChangeNotify
}
func (*DragonSpineChapterProgressChangeNotify) ProtoMessageType() ProtoMessageType {
	return "DragonSpineChapterProgressChangeNotify"
}

func (*DragonSpineChapterFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDragonSpineChapterFinishNotify
}
func (*DragonSpineChapterFinishNotify) ProtoMessageType() ProtoMessageType {
	return "DragonSpineChapterFinishNotify"
}

func (*DragonSpineCoinChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDragonSpineCoinChangeNotify
}
func (*DragonSpineCoinChangeNotify) ProtoMessageType() ProtoMessageType {
	return "DragonSpineCoinChangeNotify"
}

func (*ActivitySaleChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivitySaleChangeNotify
}
func (*ActivitySaleChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ActivitySaleChangeNotify"
}

func (*StartEffigyChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandStartEffigyChallengeReq
}
func (*StartEffigyChallengeReq) ProtoMessageType() ProtoMessageType { return "StartEffigyChallengeReq" }

func (*StartEffigyChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandStartEffigyChallengeRsp
}
func (*StartEffigyChallengeRsp) ProtoMessageType() ProtoMessageType { return "StartEffigyChallengeRsp" }

func (*EffigyChallengeInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeInfoNotify
}
func (*EffigyChallengeInfoNotify) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeInfoNotify"
}

func (*EffigyChallengeResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeResultNotify
}
func (*EffigyChallengeResultNotify) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeResultNotify"
}

func (*TakeEffigyFirstPassRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeEffigyFirstPassRewardReq
}
func (*TakeEffigyFirstPassRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeEffigyFirstPassRewardReq"
}

func (*TakeEffigyFirstPassRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeEffigyFirstPassRewardRsp
}
func (*TakeEffigyFirstPassRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeEffigyFirstPassRewardRsp"
}

func (*TakeEffigyRewardReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeEffigyRewardReq }
func (*TakeEffigyRewardReq) ProtoMessageType() ProtoMessageType { return "TakeEffigyRewardReq" }

func (*TakeEffigyRewardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeEffigyRewardRsp }
func (*TakeEffigyRewardRsp) ProtoMessageType() ProtoMessageType { return "TakeEffigyRewardRsp" }

func (*SelectEffigyChallengeConditionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSelectEffigyChallengeConditionReq
}
func (*SelectEffigyChallengeConditionReq) ProtoMessageType() ProtoMessageType {
	return "SelectEffigyChallengeConditionReq"
}

func (*SelectEffigyChallengeConditionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSelectEffigyChallengeConditionRsp
}
func (*SelectEffigyChallengeConditionRsp) ProtoMessageType() ProtoMessageType {
	return "SelectEffigyChallengeConditionRsp"
}

func (*RestartEffigyChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRestartEffigyChallengeReq
}
func (*RestartEffigyChallengeReq) ProtoMessageType() ProtoMessageType {
	return "RestartEffigyChallengeReq"
}

func (*RestartEffigyChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRestartEffigyChallengeRsp
}
func (*RestartEffigyChallengeRsp) ProtoMessageType() ProtoMessageType {
	return "RestartEffigyChallengeRsp"
}

func (*TreasureMapRegionInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapRegionInfoNotify
}
func (*TreasureMapRegionInfoNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapRegionInfoNotify"
}

func (*TreasureMapCurrencyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapCurrencyNotify
}
func (*TreasureMapCurrencyNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapCurrencyNotify"
}

func (*TreasureMapRegionActiveNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapRegionActiveNotify
}
func (*TreasureMapRegionActiveNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapRegionActiveNotify"
}

func (*TreasureMapMpChallengeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapMpChallengeNotify
}
func (*TreasureMapMpChallengeNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapMpChallengeNotify"
}

func (*TreasureMapBonusChallengeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapBonusChallengeNotify
}
func (*TreasureMapBonusChallengeNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapBonusChallengeNotify"
}

func (*TreasureMapGuideTaskDoneNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapGuideTaskDoneNotify
}
func (*TreasureMapGuideTaskDoneNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapGuideTaskDoneNotify"
}

func (*TreasureMapPreTaskDoneNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapPreTaskDoneNotify
}
func (*TreasureMapPreTaskDoneNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapPreTaskDoneNotify"
}

func (*BlessingScanReq) ProtoCommand() ProtoCommand         { return ProtoCommandBlessingScanReq }
func (*BlessingScanReq) ProtoMessageType() ProtoMessageType { return "BlessingScanReq" }

func (*BlessingScanRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBlessingScanRsp }
func (*BlessingScanRsp) ProtoMessageType() ProtoMessageType { return "BlessingScanRsp" }

func (*BlessingRedeemRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingRedeemRewardReq
}
func (*BlessingRedeemRewardReq) ProtoMessageType() ProtoMessageType { return "BlessingRedeemRewardReq" }

func (*BlessingRedeemRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingRedeemRewardRsp
}
func (*BlessingRedeemRewardRsp) ProtoMessageType() ProtoMessageType { return "BlessingRedeemRewardRsp" }

func (*BlessingGetFriendPicListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingGetFriendPicListReq
}
func (*BlessingGetFriendPicListReq) ProtoMessageType() ProtoMessageType {
	return "BlessingGetFriendPicListReq"
}

func (*BlessingGetFriendPicListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingGetFriendPicListRsp
}
func (*BlessingGetFriendPicListRsp) ProtoMessageType() ProtoMessageType {
	return "BlessingGetFriendPicListRsp"
}

func (*BlessingGiveFriendPicReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingGiveFriendPicReq
}
func (*BlessingGiveFriendPicReq) ProtoMessageType() ProtoMessageType {
	return "BlessingGiveFriendPicReq"
}

func (*BlessingGiveFriendPicRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingGiveFriendPicRsp
}
func (*BlessingGiveFriendPicRsp) ProtoMessageType() ProtoMessageType {
	return "BlessingGiveFriendPicRsp"
}

func (*BlessingAcceptGivePicReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingAcceptGivePicReq
}
func (*BlessingAcceptGivePicReq) ProtoMessageType() ProtoMessageType {
	return "BlessingAcceptGivePicReq"
}

func (*BlessingAcceptGivePicRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingAcceptGivePicRsp
}
func (*BlessingAcceptGivePicRsp) ProtoMessageType() ProtoMessageType {
	return "BlessingAcceptGivePicRsp"
}

func (*BlessingGetAllRecvPicRecordListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingGetAllRecvPicRecordListReq
}
func (*BlessingGetAllRecvPicRecordListReq) ProtoMessageType() ProtoMessageType {
	return "BlessingGetAllRecvPicRecordListReq"
}

func (*BlessingGetAllRecvPicRecordListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingGetAllRecvPicRecordListRsp
}
func (*BlessingGetAllRecvPicRecordListRsp) ProtoMessageType() ProtoMessageType {
	return "BlessingGetAllRecvPicRecordListRsp"
}

func (*BlessingRecvFriendPicNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingRecvFriendPicNotify
}
func (*BlessingRecvFriendPicNotify) ProtoMessageType() ProtoMessageType {
	return "BlessingRecvFriendPicNotify"
}

func (*BlessingAcceptAllGivePicReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingAcceptAllGivePicReq
}
func (*BlessingAcceptAllGivePicReq) ProtoMessageType() ProtoMessageType {
	return "BlessingAcceptAllGivePicReq"
}

func (*BlessingAcceptAllGivePicRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlessingAcceptAllGivePicRsp
}
func (*BlessingAcceptAllGivePicRsp) ProtoMessageType() ProtoMessageType {
	return "BlessingAcceptAllGivePicRsp"
}

func (*ExpeditionStartReq) ProtoCommand() ProtoCommand         { return ProtoCommandExpeditionStartReq }
func (*ExpeditionStartReq) ProtoMessageType() ProtoMessageType { return "ExpeditionStartReq" }

func (*ExpeditionStartRsp) ProtoCommand() ProtoCommand         { return ProtoCommandExpeditionStartRsp }
func (*ExpeditionStartRsp) ProtoMessageType() ProtoMessageType { return "ExpeditionStartRsp" }

func (*ExpeditionRecallReq) ProtoCommand() ProtoCommand         { return ProtoCommandExpeditionRecallReq }
func (*ExpeditionRecallReq) ProtoMessageType() ProtoMessageType { return "ExpeditionRecallReq" }

func (*ExpeditionRecallRsp) ProtoCommand() ProtoCommand         { return ProtoCommandExpeditionRecallRsp }
func (*ExpeditionRecallRsp) ProtoMessageType() ProtoMessageType { return "ExpeditionRecallRsp" }

func (*ExpeditionTakeRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandExpeditionTakeRewardReq
}
func (*ExpeditionTakeRewardReq) ProtoMessageType() ProtoMessageType { return "ExpeditionTakeRewardReq" }

func (*ExpeditionTakeRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandExpeditionTakeRewardRsp
}
func (*ExpeditionTakeRewardRsp) ProtoMessageType() ProtoMessageType { return "ExpeditionTakeRewardRsp" }

func (*GetExpeditionAssistInfoListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetExpeditionAssistInfoListReq
}
func (*GetExpeditionAssistInfoListReq) ProtoMessageType() ProtoMessageType {
	return "GetExpeditionAssistInfoListReq"
}

func (*GetExpeditionAssistInfoListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetExpeditionAssistInfoListRsp
}
func (*GetExpeditionAssistInfoListRsp) ProtoMessageType() ProtoMessageType {
	return "GetExpeditionAssistInfoListRsp"
}

func (*SetCurExpeditionChallengeIdReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetCurExpeditionChallengeIdReq
}
func (*SetCurExpeditionChallengeIdReq) ProtoMessageType() ProtoMessageType {
	return "SetCurExpeditionChallengeIdReq"
}

func (*SetCurExpeditionChallengeIdRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetCurExpeditionChallengeIdRsp
}
func (*SetCurExpeditionChallengeIdRsp) ProtoMessageType() ProtoMessageType {
	return "SetCurExpeditionChallengeIdRsp"
}

func (*ExpeditionChallengeEnterRegionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandExpeditionChallengeEnterRegionNotify
}
func (*ExpeditionChallengeEnterRegionNotify) ProtoMessageType() ProtoMessageType {
	return "ExpeditionChallengeEnterRegionNotify"
}

func (*ExpeditionChallengeFinishedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandExpeditionChallengeFinishedNotify
}
func (*ExpeditionChallengeFinishedNotify) ProtoMessageType() ProtoMessageType {
	return "ExpeditionChallengeFinishedNotify"
}

func (*FleurFairBalloonSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairBalloonSettleNotify
}
func (*FleurFairBalloonSettleNotify) ProtoMessageType() ProtoMessageType {
	return "FleurFairBalloonSettleNotify"
}

func (*FleurFairFallSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairFallSettleNotify
}
func (*FleurFairFallSettleNotify) ProtoMessageType() ProtoMessageType {
	return "FleurFairFallSettleNotify"
}

func (*FleurFairMusicGameSettleReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairMusicGameSettleReq
}
func (*FleurFairMusicGameSettleReq) ProtoMessageType() ProtoMessageType {
	return "FleurFairMusicGameSettleReq"
}

func (*FleurFairMusicGameSettleRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairMusicGameSettleRsp
}
func (*FleurFairMusicGameSettleRsp) ProtoMessageType() ProtoMessageType {
	return "FleurFairMusicGameSettleRsp"
}

func (*FleurFairMusicGameStartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairMusicGameStartReq
}
func (*FleurFairMusicGameStartReq) ProtoMessageType() ProtoMessageType {
	return "FleurFairMusicGameStartReq"
}

func (*FleurFairMusicGameStartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairMusicGameStartRsp
}
func (*FleurFairMusicGameStartRsp) ProtoMessageType() ProtoMessageType {
	return "FleurFairMusicGameStartRsp"
}

func (*FleurFairReplayMiniGameReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairReplayMiniGameReq
}
func (*FleurFairReplayMiniGameReq) ProtoMessageType() ProtoMessageType {
	return "FleurFairReplayMiniGameReq"
}

func (*FleurFairReplayMiniGameRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFleurFairReplayMiniGameRsp
}
func (*FleurFairReplayMiniGameRsp) ProtoMessageType() ProtoMessageType {
	return "FleurFairReplayMiniGameRsp"
}

func (*StartArenaChallengeLevelReq) ProtoCommand() ProtoCommand {
	return ProtoCommandStartArenaChallengeLevelReq
}
func (*StartArenaChallengeLevelReq) ProtoMessageType() ProtoMessageType {
	return "StartArenaChallengeLevelReq"
}

func (*StartArenaChallengeLevelRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandStartArenaChallengeLevelRsp
}
func (*StartArenaChallengeLevelRsp) ProtoMessageType() ProtoMessageType {
	return "StartArenaChallengeLevelRsp"
}

func (*ArenaChallengeFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandArenaChallengeFinishNotify
}
func (*ArenaChallengeFinishNotify) ProtoMessageType() ProtoMessageType {
	return "ArenaChallengeFinishNotify"
}

func (*WaterSpritePhaseFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWaterSpritePhaseFinishNotify
}
func (*WaterSpritePhaseFinishNotify) ProtoMessageType() ProtoMessageType {
	return "WaterSpritePhaseFinishNotify"
}

func (*ActivityTakeWatcherRewardBatchReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeWatcherRewardBatchReq
}
func (*ActivityTakeWatcherRewardBatchReq) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeWatcherRewardBatchReq"
}

func (*ActivityTakeWatcherRewardBatchRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeWatcherRewardBatchRsp
}
func (*ActivityTakeWatcherRewardBatchRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeWatcherRewardBatchRsp"
}

func (*ChannelerSlabStageActiveChallengeIndexNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChannelerSlabStageActiveChallengeIndexNotify
}
func (*ChannelerSlabStageActiveChallengeIndexNotify) ProtoMessageType() ProtoMessageType {
	return "ChannelerSlabStageActiveChallengeIndexNotify"
}

func (*ChannelerSlabStageOneoffDungeonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChannelerSlabStageOneoffDungeonNotify
}
func (*ChannelerSlabStageOneoffDungeonNotify) ProtoMessageType() ProtoMessageType {
	return "ChannelerSlabStageOneoffDungeonNotify"
}

func (*ChannellerSlabWearBuffReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabWearBuffReq
}
func (*ChannellerSlabWearBuffReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabWearBuffReq"
}

func (*ChannellerSlabWearBuffRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabWearBuffRsp
}
func (*ChannellerSlabWearBuffRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabWearBuffRsp"
}

func (*ChannellerSlabTakeoffBuffReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabTakeoffBuffReq
}
func (*ChannellerSlabTakeoffBuffReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabTakeoffBuffReq"
}

func (*ChannellerSlabTakeoffBuffRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabTakeoffBuffRsp
}
func (*ChannellerSlabTakeoffBuffRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabTakeoffBuffRsp"
}

func (*ChannellerSlabEnterLoopDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabEnterLoopDungeonReq
}
func (*ChannellerSlabEnterLoopDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabEnterLoopDungeonReq"
}

func (*ChannellerSlabEnterLoopDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabEnterLoopDungeonRsp
}
func (*ChannellerSlabEnterLoopDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabEnterLoopDungeonRsp"
}

func (*ChannellerSlabLoopDungeonTakeFirstPassRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonTakeFirstPassRewardReq
}
func (*ChannellerSlabLoopDungeonTakeFirstPassRewardReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonTakeFirstPassRewardReq"
}

func (*ChannellerSlabLoopDungeonTakeFirstPassRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonTakeFirstPassRewardRsp
}
func (*ChannellerSlabLoopDungeonTakeFirstPassRewardRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonTakeFirstPassRewardRsp"
}

func (*ChannellerSlabLoopDungeonTakeScoreRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonTakeScoreRewardReq
}
func (*ChannellerSlabLoopDungeonTakeScoreRewardReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonTakeScoreRewardReq"
}

func (*ChannellerSlabLoopDungeonTakeScoreRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonTakeScoreRewardRsp
}
func (*ChannellerSlabLoopDungeonTakeScoreRewardRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonTakeScoreRewardRsp"
}

func (*ChannellerSlabLoopDungeonChallengeInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonChallengeInfoNotify
}
func (*ChannellerSlabLoopDungeonChallengeInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonChallengeInfoNotify"
}

func (*ChannellerSlabLoopDungeonSelectConditionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonSelectConditionReq
}
func (*ChannellerSlabLoopDungeonSelectConditionReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonSelectConditionReq"
}

func (*ChannellerSlabLoopDungeonSelectConditionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabLoopDungeonSelectConditionRsp
}
func (*ChannellerSlabLoopDungeonSelectConditionRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabLoopDungeonSelectConditionRsp"
}

func (*ChannellerSlabOneOffDungeonInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabOneOffDungeonInfoReq
}
func (*ChannellerSlabOneOffDungeonInfoReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabOneOffDungeonInfoReq"
}

func (*ChannellerSlabOneOffDungeonInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabOneOffDungeonInfoRsp
}
func (*ChannellerSlabOneOffDungeonInfoRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabOneOffDungeonInfoRsp"
}

func (*ChannellerSlabOneOffDungeonInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabOneOffDungeonInfoNotify
}
func (*ChannellerSlabOneOffDungeonInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabOneOffDungeonInfoNotify"
}

func (*ChannellerSlabSaveAssistInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabSaveAssistInfoReq
}
func (*ChannellerSlabSaveAssistInfoReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabSaveAssistInfoReq"
}

func (*ChannellerSlabSaveAssistInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabSaveAssistInfoRsp
}
func (*ChannellerSlabSaveAssistInfoRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabSaveAssistInfoRsp"
}

func (*MistTrialSelectAvatarAndEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialSelectAvatarAndEnterDungeonReq
}
func (*MistTrialSelectAvatarAndEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "MistTrialSelectAvatarAndEnterDungeonReq"
}

func (*MistTrialSelectAvatarAndEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialSelectAvatarAndEnterDungeonRsp
}
func (*MistTrialSelectAvatarAndEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "MistTrialSelectAvatarAndEnterDungeonRsp"
}

func (*MistTrialGetChallengeMissionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialGetChallengeMissionReq
}
func (*MistTrialGetChallengeMissionReq) ProtoMessageType() ProtoMessageType {
	return "MistTrialGetChallengeMissionReq"
}

func (*MistTrialGetChallengeMissionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialGetChallengeMissionRsp
}
func (*MistTrialGetChallengeMissionRsp) ProtoMessageType() ProtoMessageType {
	return "MistTrialGetChallengeMissionRsp"
}

func (*MistTrialDunegonFailNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialDunegonFailNotify
}
func (*MistTrialDunegonFailNotify) ProtoMessageType() ProtoMessageType {
	return "MistTrialDunegonFailNotify"
}

func (*ChannellerSlabCheckEnterLoopDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabCheckEnterLoopDungeonReq
}
func (*ChannellerSlabCheckEnterLoopDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabCheckEnterLoopDungeonReq"
}

func (*ChannellerSlabCheckEnterLoopDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChannellerSlabCheckEnterLoopDungeonRsp
}
func (*ChannellerSlabCheckEnterLoopDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ChannellerSlabCheckEnterLoopDungeonRsp"
}

func (*HideAndSeekSelectSkillReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekSelectSkillReq
}
func (*HideAndSeekSelectSkillReq) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekSelectSkillReq"
}

func (*HideAndSeekSelectSkillRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekSelectSkillRsp
}
func (*HideAndSeekSelectSkillRsp) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekSelectSkillRsp"
}

func (*ActivityTakeScoreRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeScoreRewardReq
}
func (*ActivityTakeScoreRewardReq) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeScoreRewardReq"
}

func (*ActivityTakeScoreRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeScoreRewardRsp
}
func (*ActivityTakeScoreRewardRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeScoreRewardRsp"
}

func (*ActivityTakeAllScoreRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeAllScoreRewardReq
}
func (*ActivityTakeAllScoreRewardReq) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeAllScoreRewardReq"
}

func (*ActivityTakeAllScoreRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityTakeAllScoreRewardRsp
}
func (*ActivityTakeAllScoreRewardRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityTakeAllScoreRewardRsp"
}

func (*HideAndSeekChooseMapReq) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekChooseMapReq
}
func (*HideAndSeekChooseMapReq) ProtoMessageType() ProtoMessageType { return "HideAndSeekChooseMapReq" }

func (*HideAndSeekChooseMapRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekChooseMapRsp
}
func (*HideAndSeekChooseMapRsp) ProtoMessageType() ProtoMessageType { return "HideAndSeekChooseMapRsp" }

func (*CommonPlayerTipsNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCommonPlayerTipsNotify }
func (*CommonPlayerTipsNotify) ProtoMessageType() ProtoMessageType { return "CommonPlayerTipsNotify" }

func (*FindHilichurlFinishSecondQuestNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFindHilichurlFinishSecondQuestNotify
}
func (*FindHilichurlFinishSecondQuestNotify) ProtoMessageType() ProtoMessageType {
	return "FindHilichurlFinishSecondQuestNotify"
}

func (*FindHilichurlAcceptQuestNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFindHilichurlAcceptQuestNotify
}
func (*FindHilichurlAcceptQuestNotify) ProtoMessageType() ProtoMessageType {
	return "FindHilichurlAcceptQuestNotify"
}

func (*SummerTimeFloatSignalPositionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeFloatSignalPositionNotify
}
func (*SummerTimeFloatSignalPositionNotify) ProtoMessageType() ProtoMessageType {
	return "SummerTimeFloatSignalPositionNotify"
}

func (*SummerTimeFloatSignalUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeFloatSignalUpdateNotify
}
func (*SummerTimeFloatSignalUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "SummerTimeFloatSignalUpdateNotify"
}

func (*SummerTimeSprintBoatSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeSprintBoatSettleNotify
}
func (*SummerTimeSprintBoatSettleNotify) ProtoMessageType() ProtoMessageType {
	return "SummerTimeSprintBoatSettleNotify"
}

func (*SummerTimeSprintBoatRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeSprintBoatRestartReq
}
func (*SummerTimeSprintBoatRestartReq) ProtoMessageType() ProtoMessageType {
	return "SummerTimeSprintBoatRestartReq"
}

func (*SummerTimeSprintBoatRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeSprintBoatRestartRsp
}
func (*SummerTimeSprintBoatRestartRsp) ProtoMessageType() ProtoMessageType {
	return "SummerTimeSprintBoatRestartRsp"
}

func (*StartBuoyantCombatGalleryReq) ProtoCommand() ProtoCommand {
	return ProtoCommandStartBuoyantCombatGalleryReq
}
func (*StartBuoyantCombatGalleryReq) ProtoMessageType() ProtoMessageType {
	return "StartBuoyantCombatGalleryReq"
}

func (*StartBuoyantCombatGalleryRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandStartBuoyantCombatGalleryRsp
}
func (*StartBuoyantCombatGalleryRsp) ProtoMessageType() ProtoMessageType {
	return "StartBuoyantCombatGalleryRsp"
}

func (*BuoyantCombatSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBuoyantCombatSettleNotify
}
func (*BuoyantCombatSettleNotify) ProtoMessageType() ProtoMessageType {
	return "BuoyantCombatSettleNotify"
}

func (*SetLimitOptimizationNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSetLimitOptimizationNotify
}
func (*SetLimitOptimizationNotify) ProtoMessageType() ProtoMessageType {
	return "SetLimitOptimizationNotify"
}

func (*EchoShellUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEchoShellUpdateNotify }
func (*EchoShellUpdateNotify) ProtoMessageType() ProtoMessageType { return "EchoShellUpdateNotify" }

func (*EchoShellTakeRewardReq) ProtoCommand() ProtoCommand         { return ProtoCommandEchoShellTakeRewardReq }
func (*EchoShellTakeRewardReq) ProtoMessageType() ProtoMessageType { return "EchoShellTakeRewardReq" }

func (*EchoShellTakeRewardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEchoShellTakeRewardRsp }
func (*EchoShellTakeRewardRsp) ProtoMessageType() ProtoMessageType { return "EchoShellTakeRewardRsp" }

func (*BounceConjuringSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBounceConjuringSettleNotify
}
func (*BounceConjuringSettleNotify) ProtoMessageType() ProtoMessageType {
	return "BounceConjuringSettleNotify"
}

func (*BlitzRushParkourRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBlitzRushParkourRestartReq
}
func (*BlitzRushParkourRestartReq) ProtoMessageType() ProtoMessageType {
	return "BlitzRushParkourRestartReq"
}

func (*BlitzRushParkourRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBlitzRushParkourRestartRsp
}
func (*BlitzRushParkourRestartRsp) ProtoMessageType() ProtoMessageType {
	return "BlitzRushParkourRestartRsp"
}

func (*EnterChessDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandEnterChessDungeonReq }
func (*EnterChessDungeonReq) ProtoMessageType() ProtoMessageType { return "EnterChessDungeonReq" }

func (*EnterChessDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEnterChessDungeonRsp }
func (*EnterChessDungeonRsp) ProtoMessageType() ProtoMessageType { return "EnterChessDungeonRsp" }

func (*TreasureMapHostInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapHostInfoNotify
}
func (*TreasureMapHostInfoNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapHostInfoNotify"
}

func (*SumoSaveTeamReq) ProtoCommand() ProtoCommand         { return ProtoCommandSumoSaveTeamReq }
func (*SumoSaveTeamReq) ProtoMessageType() ProtoMessageType { return "SumoSaveTeamReq" }

func (*SumoSaveTeamRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSumoSaveTeamRsp }
func (*SumoSaveTeamRsp) ProtoMessageType() ProtoMessageType { return "SumoSaveTeamRsp" }

func (*SumoSelectTeamAndEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSumoSelectTeamAndEnterDungeonReq
}
func (*SumoSelectTeamAndEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "SumoSelectTeamAndEnterDungeonReq"
}

func (*SumoSelectTeamAndEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSumoSelectTeamAndEnterDungeonRsp
}
func (*SumoSelectTeamAndEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "SumoSelectTeamAndEnterDungeonRsp"
}

func (*SumoDungeonSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSumoDungeonSettleNotify
}
func (*SumoDungeonSettleNotify) ProtoMessageType() ProtoMessageType { return "SumoDungeonSettleNotify" }

func (*SumoEnterDungeonNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSumoEnterDungeonNotify }
func (*SumoEnterDungeonNotify) ProtoMessageType() ProtoMessageType { return "SumoEnterDungeonNotify" }

func (*SumoSwitchTeamReq) ProtoCommand() ProtoCommand         { return ProtoCommandSumoSwitchTeamReq }
func (*SumoSwitchTeamReq) ProtoMessageType() ProtoMessageType { return "SumoSwitchTeamReq" }

func (*SumoSwitchTeamRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSumoSwitchTeamRsp }
func (*SumoSwitchTeamRsp) ProtoMessageType() ProtoMessageType { return "SumoSwitchTeamRsp" }

func (*SumoLeaveDungeonNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSumoLeaveDungeonNotify }
func (*SumoLeaveDungeonNotify) ProtoMessageType() ProtoMessageType { return "SumoLeaveDungeonNotify" }

func (*SumoRestartDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandSumoRestartDungeonReq }
func (*SumoRestartDungeonReq) ProtoMessageType() ProtoMessageType { return "SumoRestartDungeonReq" }

func (*SumoRestartDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSumoRestartDungeonRsp }
func (*SumoRestartDungeonRsp) ProtoMessageType() ProtoMessageType { return "SumoRestartDungeonRsp" }

func (*ActivityDisableTransferPointInteractionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityDisableTransferPointInteractionNotify
}
func (*ActivityDisableTransferPointInteractionNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityDisableTransferPointInteractionNotify"
}

func (*SumoSetNoSwitchPunishTimeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSumoSetNoSwitchPunishTimeNotify
}
func (*SumoSetNoSwitchPunishTimeNotify) ProtoMessageType() ProtoMessageType {
	return "SumoSetNoSwitchPunishTimeNotify"
}

func (*FishingGallerySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFishingGallerySettleNotify
}
func (*FishingGallerySettleNotify) ProtoMessageType() ProtoMessageType {
	return "FishingGallerySettleNotify"
}

func (*LunaRiteSacrificeReq) ProtoCommand() ProtoCommand         { return ProtoCommandLunaRiteSacrificeReq }
func (*LunaRiteSacrificeReq) ProtoMessageType() ProtoMessageType { return "LunaRiteSacrificeReq" }

func (*LunaRiteSacrificeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLunaRiteSacrificeRsp }
func (*LunaRiteSacrificeRsp) ProtoMessageType() ProtoMessageType { return "LunaRiteSacrificeRsp" }

func (*LunaRiteTakeSacrificeRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLunaRiteTakeSacrificeRewardReq
}
func (*LunaRiteTakeSacrificeRewardReq) ProtoMessageType() ProtoMessageType {
	return "LunaRiteTakeSacrificeRewardReq"
}

func (*LunaRiteTakeSacrificeRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLunaRiteTakeSacrificeRewardRsp
}
func (*LunaRiteTakeSacrificeRewardRsp) ProtoMessageType() ProtoMessageType {
	return "LunaRiteTakeSacrificeRewardRsp"
}

func (*LunaRiteHintPointReq) ProtoCommand() ProtoCommand         { return ProtoCommandLunaRiteHintPointReq }
func (*LunaRiteHintPointReq) ProtoMessageType() ProtoMessageType { return "LunaRiteHintPointReq" }

func (*LunaRiteHintPointRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLunaRiteHintPointRsp }
func (*LunaRiteHintPointRsp) ProtoMessageType() ProtoMessageType { return "LunaRiteHintPointRsp" }

func (*LunaRiteHintPointRemoveNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLunaRiteHintPointRemoveNotify
}
func (*LunaRiteHintPointRemoveNotify) ProtoMessageType() ProtoMessageType {
	return "LunaRiteHintPointRemoveNotify"
}

func (*LunaRiteGroupBundleRegisterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLunaRiteGroupBundleRegisterNotify
}
func (*LunaRiteGroupBundleRegisterNotify) ProtoMessageType() ProtoMessageType {
	return "LunaRiteGroupBundleRegisterNotify"
}

func (*LunaRiteAreaFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLunaRiteAreaFinishNotify
}
func (*LunaRiteAreaFinishNotify) ProtoMessageType() ProtoMessageType {
	return "LunaRiteAreaFinishNotify"
}

func (*PlantFlowerGetSeedInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetSeedInfoReq
}
func (*PlantFlowerGetSeedInfoReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetSeedInfoReq"
}

func (*PlantFlowerGetSeedInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetSeedInfoRsp
}
func (*PlantFlowerGetSeedInfoRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetSeedInfoRsp"
}

func (*PlantFlowerTakeSeedRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerTakeSeedRewardReq
}
func (*PlantFlowerTakeSeedRewardReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerTakeSeedRewardReq"
}

func (*PlantFlowerTakeSeedRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerTakeSeedRewardRsp
}
func (*PlantFlowerTakeSeedRewardRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerTakeSeedRewardRsp"
}

func (*PlantFlowerSetFlowerWishReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerSetFlowerWishReq
}
func (*PlantFlowerSetFlowerWishReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerSetFlowerWishReq"
}

func (*PlantFlowerSetFlowerWishRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerSetFlowerWishRsp
}
func (*PlantFlowerSetFlowerWishRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerSetFlowerWishRsp"
}

func (*PlantFlowerGetFriendFlowerWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetFriendFlowerWishListReq
}
func (*PlantFlowerGetFriendFlowerWishListReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetFriendFlowerWishListReq"
}

func (*PlantFlowerGetFriendFlowerWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetFriendFlowerWishListRsp
}
func (*PlantFlowerGetFriendFlowerWishListRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetFriendFlowerWishListRsp"
}

func (*PlantFlowerGiveFriendFlowerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGiveFriendFlowerReq
}
func (*PlantFlowerGiveFriendFlowerReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGiveFriendFlowerReq"
}

func (*PlantFlowerGiveFriendFlowerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGiveFriendFlowerRsp
}
func (*PlantFlowerGiveFriendFlowerRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGiveFriendFlowerRsp"
}

func (*PlantFlowerGetRecvFlowerListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetRecvFlowerListReq
}
func (*PlantFlowerGetRecvFlowerListReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetRecvFlowerListReq"
}

func (*PlantFlowerGetRecvFlowerListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetRecvFlowerListRsp
}
func (*PlantFlowerGetRecvFlowerListRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetRecvFlowerListRsp"
}

func (*PlantFlowerHaveRecvFlowerNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerHaveRecvFlowerNotify
}
func (*PlantFlowerHaveRecvFlowerNotify) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerHaveRecvFlowerNotify"
}

func (*PlantFlowerAcceptGiveFlowerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerAcceptGiveFlowerReq
}
func (*PlantFlowerAcceptGiveFlowerReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerAcceptGiveFlowerReq"
}

func (*PlantFlowerAcceptGiveFlowerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerAcceptGiveFlowerRsp
}
func (*PlantFlowerAcceptGiveFlowerRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerAcceptGiveFlowerRsp"
}

func (*PlantFlowerAcceptAllGiveFlowerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerAcceptAllGiveFlowerReq
}
func (*PlantFlowerAcceptAllGiveFlowerReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerAcceptAllGiveFlowerReq"
}

func (*PlantFlowerAcceptAllGiveFlowerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerAcceptAllGiveFlowerRsp
}
func (*PlantFlowerAcceptAllGiveFlowerRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerAcceptAllGiveFlowerRsp"
}

func (*PlantFlowerGetCanGiveFriendFlowerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetCanGiveFriendFlowerReq
}
func (*PlantFlowerGetCanGiveFriendFlowerReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetCanGiveFriendFlowerReq"
}

func (*PlantFlowerGetCanGiveFriendFlowerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerGetCanGiveFriendFlowerRsp
}
func (*PlantFlowerGetCanGiveFriendFlowerRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerGetCanGiveFriendFlowerRsp"
}

func (*PlantFlowerEditFlowerCombinationReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerEditFlowerCombinationReq
}
func (*PlantFlowerEditFlowerCombinationReq) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerEditFlowerCombinationReq"
}

func (*PlantFlowerEditFlowerCombinationRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlantFlowerEditFlowerCombinationRsp
}
func (*PlantFlowerEditFlowerCombinationRsp) ProtoMessageType() ProtoMessageType {
	return "PlantFlowerEditFlowerCombinationRsp"
}

func (*MusicGameSettleReq) ProtoCommand() ProtoCommand         { return ProtoCommandMusicGameSettleReq }
func (*MusicGameSettleReq) ProtoMessageType() ProtoMessageType { return "MusicGameSettleReq" }

func (*MusicGameSettleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMusicGameSettleRsp }
func (*MusicGameSettleRsp) ProtoMessageType() ProtoMessageType { return "MusicGameSettleRsp" }

func (*MusicGameStartReq) ProtoCommand() ProtoCommand         { return ProtoCommandMusicGameStartReq }
func (*MusicGameStartReq) ProtoMessageType() ProtoMessageType { return "MusicGameStartReq" }

func (*MusicGameStartRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMusicGameStartRsp }
func (*MusicGameStartRsp) ProtoMessageType() ProtoMessageType { return "MusicGameStartRsp" }

func (*DoRoguelikeDungeonCardGachaReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDoRoguelikeDungeonCardGachaReq
}
func (*DoRoguelikeDungeonCardGachaReq) ProtoMessageType() ProtoMessageType {
	return "DoRoguelikeDungeonCardGachaReq"
}

func (*DoRoguelikeDungeonCardGachaRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDoRoguelikeDungeonCardGachaRsp
}
func (*DoRoguelikeDungeonCardGachaRsp) ProtoMessageType() ProtoMessageType {
	return "DoRoguelikeDungeonCardGachaRsp"
}

func (*RefreshRoguelikeDungeonCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshRoguelikeDungeonCardReq
}
func (*RefreshRoguelikeDungeonCardReq) ProtoMessageType() ProtoMessageType {
	return "RefreshRoguelikeDungeonCardReq"
}

func (*RefreshRoguelikeDungeonCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshRoguelikeDungeonCardRsp
}
func (*RefreshRoguelikeDungeonCardRsp) ProtoMessageType() ProtoMessageType {
	return "RefreshRoguelikeDungeonCardRsp"
}

func (*SelectRoguelikeDungeonCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSelectRoguelikeDungeonCardReq
}
func (*SelectRoguelikeDungeonCardReq) ProtoMessageType() ProtoMessageType {
	return "SelectRoguelikeDungeonCardReq"
}

func (*SelectRoguelikeDungeonCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSelectRoguelikeDungeonCardRsp
}
func (*SelectRoguelikeDungeonCardRsp) ProtoMessageType() ProtoMessageType {
	return "SelectRoguelikeDungeonCardRsp"
}

func (*EquipRoguelikeRuneReq) ProtoCommand() ProtoCommand         { return ProtoCommandEquipRoguelikeRuneReq }
func (*EquipRoguelikeRuneReq) ProtoMessageType() ProtoMessageType { return "EquipRoguelikeRuneReq" }

func (*EquipRoguelikeRuneRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEquipRoguelikeRuneRsp }
func (*EquipRoguelikeRuneRsp) ProtoMessageType() ProtoMessageType { return "EquipRoguelikeRuneRsp" }

func (*TriggerRoguelikeRuneReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTriggerRoguelikeRuneReq
}
func (*TriggerRoguelikeRuneReq) ProtoMessageType() ProtoMessageType { return "TriggerRoguelikeRuneReq" }

func (*TriggerRoguelikeRuneRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTriggerRoguelikeRuneRsp
}
func (*TriggerRoguelikeRuneRsp) ProtoMessageType() ProtoMessageType { return "TriggerRoguelikeRuneRsp" }

func (*TriggerRoguelikeCurseNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTriggerRoguelikeCurseNotify
}
func (*TriggerRoguelikeCurseNotify) ProtoMessageType() ProtoMessageType {
	return "TriggerRoguelikeCurseNotify"
}

func (*UpgradeRoguelikeShikigamiReq) ProtoCommand() ProtoCommand {
	return ProtoCommandUpgradeRoguelikeShikigamiReq
}
func (*UpgradeRoguelikeShikigamiReq) ProtoMessageType() ProtoMessageType {
	return "UpgradeRoguelikeShikigamiReq"
}

func (*UpgradeRoguelikeShikigamiRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandUpgradeRoguelikeShikigamiRsp
}
func (*UpgradeRoguelikeShikigamiRsp) ProtoMessageType() ProtoMessageType {
	return "UpgradeRoguelikeShikigamiRsp"
}

func (*RoguelikeSelectAvatarAndEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeSelectAvatarAndEnterDungeonReq
}
func (*RoguelikeSelectAvatarAndEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "RoguelikeSelectAvatarAndEnterDungeonReq"
}

func (*RoguelikeSelectAvatarAndEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeSelectAvatarAndEnterDungeonRsp
}
func (*RoguelikeSelectAvatarAndEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "RoguelikeSelectAvatarAndEnterDungeonRsp"
}

func (*RoguelikeGiveUpReq) ProtoCommand() ProtoCommand         { return ProtoCommandRoguelikeGiveUpReq }
func (*RoguelikeGiveUpReq) ProtoMessageType() ProtoMessageType { return "RoguelikeGiveUpReq" }

func (*RoguelikeGiveUpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRoguelikeGiveUpRsp }
func (*RoguelikeGiveUpRsp) ProtoMessageType() ProtoMessageType { return "RoguelikeGiveUpRsp" }

func (*RoguelikeTakeStageFirstPassRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeTakeStageFirstPassRewardReq
}
func (*RoguelikeTakeStageFirstPassRewardReq) ProtoMessageType() ProtoMessageType {
	return "RoguelikeTakeStageFirstPassRewardReq"
}

func (*RoguelikeTakeStageFirstPassRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeTakeStageFirstPassRewardRsp
}
func (*RoguelikeTakeStageFirstPassRewardRsp) ProtoMessageType() ProtoMessageType {
	return "RoguelikeTakeStageFirstPassRewardRsp"
}

func (*GiveUpRoguelikeDungeonCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGiveUpRoguelikeDungeonCardReq
}
func (*GiveUpRoguelikeDungeonCardReq) ProtoMessageType() ProtoMessageType {
	return "GiveUpRoguelikeDungeonCardReq"
}

func (*GiveUpRoguelikeDungeonCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGiveUpRoguelikeDungeonCardRsp
}
func (*GiveUpRoguelikeDungeonCardRsp) ProtoMessageType() ProtoMessageType {
	return "GiveUpRoguelikeDungeonCardRsp"
}

func (*EnterRoguelikeDungeonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterRoguelikeDungeonNotify
}
func (*EnterRoguelikeDungeonNotify) ProtoMessageType() ProtoMessageType {
	return "EnterRoguelikeDungeonNotify"
}

func (*StartRogueEliteCellChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandStartRogueEliteCellChallengeReq
}
func (*StartRogueEliteCellChallengeReq) ProtoMessageType() ProtoMessageType {
	return "StartRogueEliteCellChallengeReq"
}

func (*StartRogueEliteCellChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandStartRogueEliteCellChallengeRsp
}
func (*StartRogueEliteCellChallengeRsp) ProtoMessageType() ProtoMessageType {
	return "StartRogueEliteCellChallengeRsp"
}

func (*StartRogueNormalCellChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandStartRogueNormalCellChallengeReq
}
func (*StartRogueNormalCellChallengeReq) ProtoMessageType() ProtoMessageType {
	return "StartRogueNormalCellChallengeReq"
}

func (*StartRogueNormalCellChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandStartRogueNormalCellChallengeRsp
}
func (*StartRogueNormalCellChallengeRsp) ProtoMessageType() ProtoMessageType {
	return "StartRogueNormalCellChallengeRsp"
}

func (*RogueCellUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRogueCellUpdateNotify }
func (*RogueCellUpdateNotify) ProtoMessageType() ProtoMessageType { return "RogueCellUpdateNotify" }

func (*RogueDungeonPlayerCellChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDungeonPlayerCellChangeNotify
}
func (*RogueDungeonPlayerCellChangeNotify) ProtoMessageType() ProtoMessageType {
	return "RogueDungeonPlayerCellChangeNotify"
}

func (*RogueHealAvatarsReq) ProtoCommand() ProtoCommand         { return ProtoCommandRogueHealAvatarsReq }
func (*RogueHealAvatarsReq) ProtoMessageType() ProtoMessageType { return "RogueHealAvatarsReq" }

func (*RogueHealAvatarsRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRogueHealAvatarsRsp }
func (*RogueHealAvatarsRsp) ProtoMessageType() ProtoMessageType { return "RogueHealAvatarsRsp" }

func (*RogueResumeDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandRogueResumeDungeonReq }
func (*RogueResumeDungeonReq) ProtoMessageType() ProtoMessageType { return "RogueResumeDungeonReq" }

func (*RogueResumeDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRogueResumeDungeonRsp }
func (*RogueResumeDungeonRsp) ProtoMessageType() ProtoMessageType { return "RogueResumeDungeonRsp" }

func (*ClearRoguelikeCurseNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClearRoguelikeCurseNotify
}
func (*ClearRoguelikeCurseNotify) ProtoMessageType() ProtoMessageType {
	return "ClearRoguelikeCurseNotify"
}

func (*RoguelikeCardGachaNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeCardGachaNotify
}
func (*RoguelikeCardGachaNotify) ProtoMessageType() ProtoMessageType {
	return "RoguelikeCardGachaNotify"
}

func (*RogueSwitchAvatarReq) ProtoCommand() ProtoCommand         { return ProtoCommandRogueSwitchAvatarReq }
func (*RogueSwitchAvatarReq) ProtoMessageType() ProtoMessageType { return "RogueSwitchAvatarReq" }

func (*RogueSwitchAvatarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRogueSwitchAvatarRsp }
func (*RogueSwitchAvatarRsp) ProtoMessageType() ProtoMessageType { return "RogueSwitchAvatarRsp" }

func (*DisableRoguelikeTrapNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDisableRoguelikeTrapNotify
}
func (*DisableRoguelikeTrapNotify) ProtoMessageType() ProtoMessageType {
	return "DisableRoguelikeTrapNotify"
}

func (*RoguelikeRuneRecordUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeRuneRecordUpdateNotify
}
func (*RoguelikeRuneRecordUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "RoguelikeRuneRecordUpdateNotify"
}

func (*RoguelikeMistClearNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeMistClearNotify
}
func (*RoguelikeMistClearNotify) ProtoMessageType() ProtoMessageType {
	return "RoguelikeMistClearNotify"
}

func (*RoguelikeEffectDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeEffectDataNotify
}
func (*RoguelikeEffectDataNotify) ProtoMessageType() ProtoMessageType {
	return "RoguelikeEffectDataNotify"
}

func (*RoguelikeEffectViewReq) ProtoCommand() ProtoCommand         { return ProtoCommandRoguelikeEffectViewReq }
func (*RoguelikeEffectViewReq) ProtoMessageType() ProtoMessageType { return "RoguelikeEffectViewReq" }

func (*RoguelikeEffectViewRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRoguelikeEffectViewRsp }
func (*RoguelikeEffectViewRsp) ProtoMessageType() ProtoMessageType { return "RoguelikeEffectViewRsp" }

func (*RoguelikeResourceBonusPropUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeResourceBonusPropUpdateNotify
}
func (*RoguelikeResourceBonusPropUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "RoguelikeResourceBonusPropUpdateNotify"
}

func (*RoguelikeRefreshCardCostUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRoguelikeRefreshCardCostUpdateNotify
}
func (*RoguelikeRefreshCardCostUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "RoguelikeRefreshCardCostUpdateNotify"
}

func (*DigActivityMarkPointChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDigActivityMarkPointChangeNotify
}
func (*DigActivityMarkPointChangeNotify) ProtoMessageType() ProtoMessageType {
	return "DigActivityMarkPointChangeNotify"
}

func (*DigActivityChangeGadgetStateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDigActivityChangeGadgetStateReq
}
func (*DigActivityChangeGadgetStateReq) ProtoMessageType() ProtoMessageType {
	return "DigActivityChangeGadgetStateReq"
}

func (*DigActivityChangeGadgetStateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDigActivityChangeGadgetStateRsp
}
func (*DigActivityChangeGadgetStateRsp) ProtoMessageType() ProtoMessageType {
	return "DigActivityChangeGadgetStateRsp"
}

func (*WinterCampStageInfoChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampStageInfoChangeNotify
}
func (*WinterCampStageInfoChangeNotify) ProtoMessageType() ProtoMessageType {
	return "WinterCampStageInfoChangeNotify"
}

func (*WinterCampRaceScoreNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampRaceScoreNotify
}
func (*WinterCampRaceScoreNotify) ProtoMessageType() ProtoMessageType {
	return "WinterCampRaceScoreNotify"
}

func (*WinterCampGiveFriendItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGiveFriendItemReq
}
func (*WinterCampGiveFriendItemReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampGiveFriendItemReq"
}

func (*WinterCampGiveFriendItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGiveFriendItemRsp
}
func (*WinterCampGiveFriendItemRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampGiveFriendItemRsp"
}

func (*WinterCampSetWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampSetWishListReq
}
func (*WinterCampSetWishListReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampSetWishListReq"
}

func (*WinterCampSetWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampSetWishListRsp
}
func (*WinterCampSetWishListRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampSetWishListRsp"
}

func (*WinterCampGetFriendWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGetFriendWishListReq
}
func (*WinterCampGetFriendWishListReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampGetFriendWishListReq"
}

func (*WinterCampGetFriendWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGetFriendWishListRsp
}
func (*WinterCampGetFriendWishListRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampGetFriendWishListRsp"
}

func (*WinterCampRecvItemNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampRecvItemNotify
}
func (*WinterCampRecvItemNotify) ProtoMessageType() ProtoMessageType {
	return "WinterCampRecvItemNotify"
}

func (*WinterCampAcceptGiveItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampAcceptGiveItemReq
}
func (*WinterCampAcceptGiveItemReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampAcceptGiveItemReq"
}

func (*WinterCampAcceptGiveItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampAcceptGiveItemRsp
}
func (*WinterCampAcceptGiveItemRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampAcceptGiveItemRsp"
}

func (*WinterCampAcceptAllGiveItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampAcceptAllGiveItemReq
}
func (*WinterCampAcceptAllGiveItemReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampAcceptAllGiveItemReq"
}

func (*WinterCampAcceptAllGiveItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampAcceptAllGiveItemRsp
}
func (*WinterCampAcceptAllGiveItemRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampAcceptAllGiveItemRsp"
}

func (*WinterCampGetCanGiveFriendItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGetCanGiveFriendItemReq
}
func (*WinterCampGetCanGiveFriendItemReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampGetCanGiveFriendItemReq"
}

func (*WinterCampGetCanGiveFriendItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGetCanGiveFriendItemRsp
}
func (*WinterCampGetCanGiveFriendItemRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampGetCanGiveFriendItemRsp"
}

func (*WinterCampGetRecvItemListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGetRecvItemListReq
}
func (*WinterCampGetRecvItemListReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampGetRecvItemListReq"
}

func (*WinterCampGetRecvItemListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampGetRecvItemListRsp
}
func (*WinterCampGetRecvItemListRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampGetRecvItemListRsp"
}

func (*WinterCampEditSnowmanCombinationReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampEditSnowmanCombinationReq
}
func (*WinterCampEditSnowmanCombinationReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampEditSnowmanCombinationReq"
}

func (*WinterCampEditSnowmanCombinationRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampEditSnowmanCombinationRsp
}
func (*WinterCampEditSnowmanCombinationRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampEditSnowmanCombinationRsp"
}

func (*WinterCampTriathlonSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTriathlonSettleNotify
}
func (*WinterCampTriathlonSettleNotify) ProtoMessageType() ProtoMessageType {
	return "WinterCampTriathlonSettleNotify"
}

func (*WinterCampTakeExploreRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTakeExploreRewardReq
}
func (*WinterCampTakeExploreRewardReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampTakeExploreRewardReq"
}

func (*WinterCampTakeExploreRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTakeExploreRewardRsp
}
func (*WinterCampTakeExploreRewardRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampTakeExploreRewardRsp"
}

func (*WinterCampTakeBattleRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTakeBattleRewardReq
}
func (*WinterCampTakeBattleRewardReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampTakeBattleRewardReq"
}

func (*WinterCampTakeBattleRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTakeBattleRewardRsp
}
func (*WinterCampTakeBattleRewardRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampTakeBattleRewardRsp"
}

func (*WinterCampTriathlonRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTriathlonRestartReq
}
func (*WinterCampTriathlonRestartReq) ProtoMessageType() ProtoMessageType {
	return "WinterCampTriathlonRestartReq"
}

func (*WinterCampTriathlonRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWinterCampTriathlonRestartRsp
}
func (*WinterCampTriathlonRestartRsp) ProtoMessageType() ProtoMessageType {
	return "WinterCampTriathlonRestartRsp"
}

func (*MistTrialSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMistTrialSettleNotify }
func (*MistTrialSettleNotify) ProtoMessageType() ProtoMessageType { return "MistTrialSettleNotify" }

func (*MistTrialGetDungeonExhibitionDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialGetDungeonExhibitionDataReq
}
func (*MistTrialGetDungeonExhibitionDataReq) ProtoMessageType() ProtoMessageType {
	return "MistTrialGetDungeonExhibitionDataReq"
}

func (*MistTrialGetDungeonExhibitionDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMistTrialGetDungeonExhibitionDataRsp
}
func (*MistTrialGetDungeonExhibitionDataRsp) ProtoMessageType() ProtoMessageType {
	return "MistTrialGetDungeonExhibitionDataRsp"
}

func (*PotionResetChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionResetChallengeReq
}
func (*PotionResetChallengeReq) ProtoMessageType() ProtoMessageType { return "PotionResetChallengeReq" }

func (*PotionResetChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionResetChallengeRsp
}
func (*PotionResetChallengeRsp) ProtoMessageType() ProtoMessageType { return "PotionResetChallengeRsp" }

func (*PotionEnterDungeonReq) ProtoCommand() ProtoCommand         { return ProtoCommandPotionEnterDungeonReq }
func (*PotionEnterDungeonReq) ProtoMessageType() ProtoMessageType { return "PotionEnterDungeonReq" }

func (*PotionEnterDungeonRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPotionEnterDungeonRsp }
func (*PotionEnterDungeonRsp) ProtoMessageType() ProtoMessageType { return "PotionEnterDungeonRsp" }

func (*PotionEnterDungeonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionEnterDungeonNotify
}
func (*PotionEnterDungeonNotify) ProtoMessageType() ProtoMessageType {
	return "PotionEnterDungeonNotify"
}

func (*PotionSaveDungeonResultReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionSaveDungeonResultReq
}
func (*PotionSaveDungeonResultReq) ProtoMessageType() ProtoMessageType {
	return "PotionSaveDungeonResultReq"
}

func (*PotionSaveDungeonResultRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionSaveDungeonResultRsp
}
func (*PotionSaveDungeonResultRsp) ProtoMessageType() ProtoMessageType {
	return "PotionSaveDungeonResultRsp"
}

func (*PotionRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionRestartDungeonReq
}
func (*PotionRestartDungeonReq) ProtoMessageType() ProtoMessageType { return "PotionRestartDungeonReq" }

func (*PotionRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPotionRestartDungeonRsp
}
func (*PotionRestartDungeonRsp) ProtoMessageType() ProtoMessageType { return "PotionRestartDungeonRsp" }

func (*TanukiTravelFinishGuideQuestNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTanukiTravelFinishGuideQuestNotify
}
func (*TanukiTravelFinishGuideQuestNotify) ProtoMessageType() ProtoMessageType {
	return "TanukiTravelFinishGuideQuestNotify"
}

func (*FinishLanternProjectionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFinishLanternProjectionReq
}
func (*FinishLanternProjectionReq) ProtoMessageType() ProtoMessageType {
	return "FinishLanternProjectionReq"
}

func (*FinishLanternProjectionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFinishLanternProjectionRsp
}
func (*FinishLanternProjectionRsp) ProtoMessageType() ProtoMessageType {
	return "FinishLanternProjectionRsp"
}

func (*ViewLanternProjectionTipsReq) ProtoCommand() ProtoCommand {
	return ProtoCommandViewLanternProjectionTipsReq
}
func (*ViewLanternProjectionTipsReq) ProtoMessageType() ProtoMessageType {
	return "ViewLanternProjectionTipsReq"
}

func (*ViewLanternProjectionTipsRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandViewLanternProjectionTipsRsp
}
func (*ViewLanternProjectionTipsRsp) ProtoMessageType() ProtoMessageType {
	return "ViewLanternProjectionTipsRsp"
}

func (*ViewLanternProjectionLevelTipsReq) ProtoCommand() ProtoCommand {
	return ProtoCommandViewLanternProjectionLevelTipsReq
}
func (*ViewLanternProjectionLevelTipsReq) ProtoMessageType() ProtoMessageType {
	return "ViewLanternProjectionLevelTipsReq"
}

func (*ViewLanternProjectionLevelTipsRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandViewLanternProjectionLevelTipsRsp
}
func (*ViewLanternProjectionLevelTipsRsp) ProtoMessageType() ProtoMessageType {
	return "ViewLanternProjectionLevelTipsRsp"
}

func (*SalvagePreventSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSalvagePreventSettleNotify
}
func (*SalvagePreventSettleNotify) ProtoMessageType() ProtoMessageType {
	return "SalvagePreventSettleNotify"
}

func (*SalvageEscortSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSalvageEscortSettleNotify
}
func (*SalvageEscortSettleNotify) ProtoMessageType() ProtoMessageType {
	return "SalvageEscortSettleNotify"
}

func (*LanternRiteTakeSkinRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteTakeSkinRewardReq
}
func (*LanternRiteTakeSkinRewardReq) ProtoMessageType() ProtoMessageType {
	return "LanternRiteTakeSkinRewardReq"
}

func (*LanternRiteTakeSkinRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteTakeSkinRewardRsp
}
func (*LanternRiteTakeSkinRewardRsp) ProtoMessageType() ProtoMessageType {
	return "LanternRiteTakeSkinRewardRsp"
}

func (*SalvagePreventRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSalvagePreventRestartReq
}
func (*SalvagePreventRestartReq) ProtoMessageType() ProtoMessageType {
	return "SalvagePreventRestartReq"
}

func (*SalvagePreventRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSalvagePreventRestartRsp
}
func (*SalvagePreventRestartRsp) ProtoMessageType() ProtoMessageType {
	return "SalvagePreventRestartRsp"
}

func (*SalvageEscortRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSalvageEscortRestartReq
}
func (*SalvageEscortRestartReq) ProtoMessageType() ProtoMessageType { return "SalvageEscortRestartReq" }

func (*SalvageEscortRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSalvageEscortRestartRsp
}
func (*SalvageEscortRestartRsp) ProtoMessageType() ProtoMessageType { return "SalvageEscortRestartRsp" }

func (*LanternRiteStartFireworksReformReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteStartFireworksReformReq
}
func (*LanternRiteStartFireworksReformReq) ProtoMessageType() ProtoMessageType {
	return "LanternRiteStartFireworksReformReq"
}

func (*LanternRiteStartFireworksReformRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteStartFireworksReformRsp
}
func (*LanternRiteStartFireworksReformRsp) ProtoMessageType() ProtoMessageType {
	return "LanternRiteStartFireworksReformRsp"
}

func (*LanternRiteDoFireworksReformReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteDoFireworksReformReq
}
func (*LanternRiteDoFireworksReformReq) ProtoMessageType() ProtoMessageType {
	return "LanternRiteDoFireworksReformReq"
}

func (*LanternRiteDoFireworksReformRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteDoFireworksReformRsp
}
func (*LanternRiteDoFireworksReformRsp) ProtoMessageType() ProtoMessageType {
	return "LanternRiteDoFireworksReformRsp"
}

func (*LanternRiteEndFireworksReformReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteEndFireworksReformReq
}
func (*LanternRiteEndFireworksReformReq) ProtoMessageType() ProtoMessageType {
	return "LanternRiteEndFireworksReformReq"
}

func (*LanternRiteEndFireworksReformRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanternRiteEndFireworksReformRsp
}
func (*LanternRiteEndFireworksReformRsp) ProtoMessageType() ProtoMessageType {
	return "LanternRiteEndFireworksReformRsp"
}

func (*UpdateSalvageBundleMarkReq) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdateSalvageBundleMarkReq
}
func (*UpdateSalvageBundleMarkReq) ProtoMessageType() ProtoMessageType {
	return "UpdateSalvageBundleMarkReq"
}

func (*UpdateSalvageBundleMarkRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdateSalvageBundleMarkRsp
}
func (*UpdateSalvageBundleMarkRsp) ProtoMessageType() ProtoMessageType {
	return "UpdateSalvageBundleMarkRsp"
}

func (*MichiaeMatsuriDarkPressureLevelUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriDarkPressureLevelUpdateNotify
}
func (*MichiaeMatsuriDarkPressureLevelUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriDarkPressureLevelUpdateNotify"
}

func (*MichiaeMatsuriInteractStatueReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriInteractStatueReq
}
func (*MichiaeMatsuriInteractStatueReq) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriInteractStatueReq"
}

func (*MichiaeMatsuriInteractStatueRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriInteractStatueRsp
}
func (*MichiaeMatsuriInteractStatueRsp) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriInteractStatueRsp"
}

func (*MichiaeMatsuriUnlockCrystalSkillReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriUnlockCrystalSkillReq
}
func (*MichiaeMatsuriUnlockCrystalSkillReq) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriUnlockCrystalSkillReq"
}

func (*MichiaeMatsuriUnlockCrystalSkillRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriUnlockCrystalSkillRsp
}
func (*MichiaeMatsuriUnlockCrystalSkillRsp) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriUnlockCrystalSkillRsp"
}

func (*MichiaeMatsuriStartBossChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriStartBossChallengeReq
}
func (*MichiaeMatsuriStartBossChallengeReq) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriStartBossChallengeReq"
}

func (*MichiaeMatsuriStartBossChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriStartBossChallengeRsp
}
func (*MichiaeMatsuriStartBossChallengeRsp) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriStartBossChallengeRsp"
}

func (*MichiaeMatsuriStartDarkChallengeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriStartDarkChallengeReq
}
func (*MichiaeMatsuriStartDarkChallengeReq) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriStartDarkChallengeReq"
}

func (*MichiaeMatsuriStartDarkChallengeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriStartDarkChallengeRsp
}
func (*MichiaeMatsuriStartDarkChallengeRsp) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriStartDarkChallengeRsp"
}

func (*MichiaeMatsuriRemoveChestMarkNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriRemoveChestMarkNotify
}
func (*MichiaeMatsuriRemoveChestMarkNotify) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriRemoveChestMarkNotify"
}

func (*MichiaeMatsuriRemoveChallengeMarkNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriRemoveChallengeMarkNotify
}
func (*MichiaeMatsuriRemoveChallengeMarkNotify) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriRemoveChallengeMarkNotify"
}

func (*MichiaeMatsuriGainCrystalExpUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMichiaeMatsuriGainCrystalExpUpdateNotify
}
func (*MichiaeMatsuriGainCrystalExpUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "MichiaeMatsuriGainCrystalExpUpdateNotify"
}

func (*BartenderCompleteOrderReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderCompleteOrderReq
}
func (*BartenderCompleteOrderReq) ProtoMessageType() ProtoMessageType {
	return "BartenderCompleteOrderReq"
}

func (*BartenderCompleteOrderRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderCompleteOrderRsp
}
func (*BartenderCompleteOrderRsp) ProtoMessageType() ProtoMessageType {
	return "BartenderCompleteOrderRsp"
}

func (*BartenderCancelOrderReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderCancelOrderReq
}
func (*BartenderCancelOrderReq) ProtoMessageType() ProtoMessageType { return "BartenderCancelOrderReq" }

func (*BartenderCancelOrderRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderCancelOrderRsp
}
func (*BartenderCancelOrderRsp) ProtoMessageType() ProtoMessageType { return "BartenderCancelOrderRsp" }

func (*BartenderGetFormulaReq) ProtoCommand() ProtoCommand         { return ProtoCommandBartenderGetFormulaReq }
func (*BartenderGetFormulaReq) ProtoMessageType() ProtoMessageType { return "BartenderGetFormulaReq" }

func (*BartenderGetFormulaRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBartenderGetFormulaRsp }
func (*BartenderGetFormulaRsp) ProtoMessageType() ProtoMessageType { return "BartenderGetFormulaRsp" }

func (*BartenderStartLevelReq) ProtoCommand() ProtoCommand         { return ProtoCommandBartenderStartLevelReq }
func (*BartenderStartLevelReq) ProtoMessageType() ProtoMessageType { return "BartenderStartLevelReq" }

func (*BartenderStartLevelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBartenderStartLevelRsp }
func (*BartenderStartLevelRsp) ProtoMessageType() ProtoMessageType { return "BartenderStartLevelRsp" }

func (*BartenderCancelLevelReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderCancelLevelReq
}
func (*BartenderCancelLevelReq) ProtoMessageType() ProtoMessageType { return "BartenderCancelLevelReq" }

func (*BartenderCancelLevelRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderCancelLevelRsp
}
func (*BartenderCancelLevelRsp) ProtoMessageType() ProtoMessageType { return "BartenderCancelLevelRsp" }

func (*BartenderLevelProgressNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderLevelProgressNotify
}
func (*BartenderLevelProgressNotify) ProtoMessageType() ProtoMessageType {
	return "BartenderLevelProgressNotify"
}

func (*BartenderFinishLevelReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderFinishLevelReq
}
func (*BartenderFinishLevelReq) ProtoMessageType() ProtoMessageType { return "BartenderFinishLevelReq" }

func (*BartenderFinishLevelRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBartenderFinishLevelRsp
}
func (*BartenderFinishLevelRsp) ProtoMessageType() ProtoMessageType { return "BartenderFinishLevelRsp" }

func (*CrystalLinkEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCrystalLinkEnterDungeonReq
}
func (*CrystalLinkEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "CrystalLinkEnterDungeonReq"
}

func (*CrystalLinkEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCrystalLinkEnterDungeonRsp
}
func (*CrystalLinkEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "CrystalLinkEnterDungeonRsp"
}

func (*CrystalLinkDungeonInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCrystalLinkDungeonInfoNotify
}
func (*CrystalLinkDungeonInfoNotify) ProtoMessageType() ProtoMessageType {
	return "CrystalLinkDungeonInfoNotify"
}

func (*CrystalLinkRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCrystalLinkRestartDungeonReq
}
func (*CrystalLinkRestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "CrystalLinkRestartDungeonReq"
}

func (*CrystalLinkRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCrystalLinkRestartDungeonRsp
}
func (*CrystalLinkRestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "CrystalLinkRestartDungeonRsp"
}

func (*QuickOpenActivityReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuickOpenActivityReq }
func (*QuickOpenActivityReq) ProtoMessageType() ProtoMessageType { return "QuickOpenActivityReq" }

func (*QuickOpenActivityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuickOpenActivityRsp }
func (*QuickOpenActivityRsp) ProtoMessageType() ProtoMessageType { return "QuickOpenActivityRsp" }

func (*IrodoriEditFlowerCombinationReq) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriEditFlowerCombinationReq
}
func (*IrodoriEditFlowerCombinationReq) ProtoMessageType() ProtoMessageType {
	return "IrodoriEditFlowerCombinationReq"
}

func (*IrodoriEditFlowerCombinationRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriEditFlowerCombinationRsp
}
func (*IrodoriEditFlowerCombinationRsp) ProtoMessageType() ProtoMessageType {
	return "IrodoriEditFlowerCombinationRsp"
}

func (*IrodoriScanEntityReq) ProtoCommand() ProtoCommand         { return ProtoCommandIrodoriScanEntityReq }
func (*IrodoriScanEntityReq) ProtoMessageType() ProtoMessageType { return "IrodoriScanEntityReq" }

func (*IrodoriScanEntityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandIrodoriScanEntityRsp }
func (*IrodoriScanEntityRsp) ProtoMessageType() ProtoMessageType { return "IrodoriScanEntityRsp" }

func (*IrodoriFillPoetryReq) ProtoCommand() ProtoCommand         { return ProtoCommandIrodoriFillPoetryReq }
func (*IrodoriFillPoetryReq) ProtoMessageType() ProtoMessageType { return "IrodoriFillPoetryReq" }

func (*IrodoriFillPoetryRsp) ProtoCommand() ProtoCommand         { return ProtoCommandIrodoriFillPoetryRsp }
func (*IrodoriFillPoetryRsp) ProtoMessageType() ProtoMessageType { return "IrodoriFillPoetryRsp" }

func (*IrodoriChessEquipCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriChessEquipCardReq
}
func (*IrodoriChessEquipCardReq) ProtoMessageType() ProtoMessageType {
	return "IrodoriChessEquipCardReq"
}

func (*IrodoriChessEquipCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriChessEquipCardRsp
}
func (*IrodoriChessEquipCardRsp) ProtoMessageType() ProtoMessageType {
	return "IrodoriChessEquipCardRsp"
}

func (*IrodoriChessUnequipCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriChessUnequipCardReq
}
func (*IrodoriChessUnequipCardReq) ProtoMessageType() ProtoMessageType {
	return "IrodoriChessUnequipCardReq"
}

func (*IrodoriChessUnequipCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriChessUnequipCardRsp
}
func (*IrodoriChessUnequipCardRsp) ProtoMessageType() ProtoMessageType {
	return "IrodoriChessUnequipCardRsp"
}

func (*EnterIrodoriChessDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterIrodoriChessDungeonReq
}
func (*EnterIrodoriChessDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EnterIrodoriChessDungeonReq"
}

func (*EnterIrodoriChessDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterIrodoriChessDungeonRsp
}
func (*EnterIrodoriChessDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EnterIrodoriChessDungeonRsp"
}

func (*IrodoriMasterStartGalleryReq) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriMasterStartGalleryReq
}
func (*IrodoriMasterStartGalleryReq) ProtoMessageType() ProtoMessageType {
	return "IrodoriMasterStartGalleryReq"
}

func (*IrodoriMasterStartGalleryRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriMasterStartGalleryRsp
}
func (*IrodoriMasterStartGalleryRsp) ProtoMessageType() ProtoMessageType {
	return "IrodoriMasterStartGalleryRsp"
}

func (*IrodoriMasterGalleryCgEndNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriMasterGalleryCgEndNotify
}
func (*IrodoriMasterGalleryCgEndNotify) ProtoMessageType() ProtoMessageType {
	return "IrodoriMasterGalleryCgEndNotify"
}

func (*IrodoriMasterGallerySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIrodoriMasterGallerySettleNotify
}
func (*IrodoriMasterGallerySettleNotify) ProtoMessageType() ProtoMessageType {
	return "IrodoriMasterGallerySettleNotify"
}

func (*PhotoActivityFinishReq) ProtoCommand() ProtoCommand         { return ProtoCommandPhotoActivityFinishReq }
func (*PhotoActivityFinishReq) ProtoMessageType() ProtoMessageType { return "PhotoActivityFinishReq" }

func (*PhotoActivityFinishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPhotoActivityFinishRsp }
func (*PhotoActivityFinishRsp) ProtoMessageType() ProtoMessageType { return "PhotoActivityFinishRsp" }

func (*PhotoActivityClientViewReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPhotoActivityClientViewReq
}
func (*PhotoActivityClientViewReq) ProtoMessageType() ProtoMessageType {
	return "PhotoActivityClientViewReq"
}

func (*PhotoActivityClientViewRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPhotoActivityClientViewRsp
}
func (*PhotoActivityClientViewRsp) ProtoMessageType() ProtoMessageType {
	return "PhotoActivityClientViewRsp"
}

func (*SpiceActivityFinishMakeSpiceReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSpiceActivityFinishMakeSpiceReq
}
func (*SpiceActivityFinishMakeSpiceReq) ProtoMessageType() ProtoMessageType {
	return "SpiceActivityFinishMakeSpiceReq"
}

func (*SpiceActivityFinishMakeSpiceRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSpiceActivityFinishMakeSpiceRsp
}
func (*SpiceActivityFinishMakeSpiceRsp) ProtoMessageType() ProtoMessageType {
	return "SpiceActivityFinishMakeSpiceRsp"
}

func (*SpiceActivityProcessFoodReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSpiceActivityProcessFoodReq
}
func (*SpiceActivityProcessFoodReq) ProtoMessageType() ProtoMessageType {
	return "SpiceActivityProcessFoodReq"
}

func (*SpiceActivityProcessFoodRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSpiceActivityProcessFoodRsp
}
func (*SpiceActivityProcessFoodRsp) ProtoMessageType() ProtoMessageType {
	return "SpiceActivityProcessFoodRsp"
}

func (*SpiceActivityGivingRecordNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSpiceActivityGivingRecordNotify
}
func (*SpiceActivityGivingRecordNotify) ProtoMessageType() ProtoMessageType {
	return "SpiceActivityGivingRecordNotify"
}

func (*GachaActivityPercentNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityPercentNotify
}
func (*GachaActivityPercentNotify) ProtoMessageType() ProtoMessageType {
	return "GachaActivityPercentNotify"
}

func (*GachaActivityUpdateElemNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityUpdateElemNotify
}
func (*GachaActivityUpdateElemNotify) ProtoMessageType() ProtoMessageType {
	return "GachaActivityUpdateElemNotify"
}

func (*GachaActivityCreateRobotReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityCreateRobotReq
}
func (*GachaActivityCreateRobotReq) ProtoMessageType() ProtoMessageType {
	return "GachaActivityCreateRobotReq"
}

func (*GachaActivityCreateRobotRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityCreateRobotRsp
}
func (*GachaActivityCreateRobotRsp) ProtoMessageType() ProtoMessageType {
	return "GachaActivityCreateRobotRsp"
}

func (*GachaActivityTakeRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityTakeRewardReq
}
func (*GachaActivityTakeRewardReq) ProtoMessageType() ProtoMessageType {
	return "GachaActivityTakeRewardReq"
}

func (*GachaActivityTakeRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityTakeRewardRsp
}
func (*GachaActivityTakeRewardRsp) ProtoMessageType() ProtoMessageType {
	return "GachaActivityTakeRewardRsp"
}

func (*GachaActivityResetReq) ProtoCommand() ProtoCommand         { return ProtoCommandGachaActivityResetReq }
func (*GachaActivityResetReq) ProtoMessageType() ProtoMessageType { return "GachaActivityResetReq" }

func (*GachaActivityResetRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGachaActivityResetRsp }
func (*GachaActivityResetRsp) ProtoMessageType() ProtoMessageType { return "GachaActivityResetRsp" }

func (*GachaActivityNextStageReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityNextStageReq
}
func (*GachaActivityNextStageReq) ProtoMessageType() ProtoMessageType {
	return "GachaActivityNextStageReq"
}

func (*GachaActivityNextStageRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGachaActivityNextStageRsp
}
func (*GachaActivityNextStageRsp) ProtoMessageType() ProtoMessageType {
	return "GachaActivityNextStageRsp"
}

func (*ActivityGiveFriendGiftReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGiveFriendGiftReq
}
func (*ActivityGiveFriendGiftReq) ProtoMessageType() ProtoMessageType {
	return "ActivityGiveFriendGiftReq"
}

func (*ActivityGiveFriendGiftRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGiveFriendGiftRsp
}
func (*ActivityGiveFriendGiftRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityGiveFriendGiftRsp"
}

func (*ActivityGetRecvGiftListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGetRecvGiftListReq
}
func (*ActivityGetRecvGiftListReq) ProtoMessageType() ProtoMessageType {
	return "ActivityGetRecvGiftListReq"
}

func (*ActivityGetRecvGiftListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGetRecvGiftListRsp
}
func (*ActivityGetRecvGiftListRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityGetRecvGiftListRsp"
}

func (*ActivityHaveRecvGiftNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityHaveRecvGiftNotify
}
func (*ActivityHaveRecvGiftNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityHaveRecvGiftNotify"
}

func (*ActivityAcceptGiveGiftReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityAcceptGiveGiftReq
}
func (*ActivityAcceptGiveGiftReq) ProtoMessageType() ProtoMessageType {
	return "ActivityAcceptGiveGiftReq"
}

func (*ActivityAcceptGiveGiftRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityAcceptGiveGiftRsp
}
func (*ActivityAcceptGiveGiftRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityAcceptGiveGiftRsp"
}

func (*ActivityAcceptAllGiveGiftReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityAcceptAllGiveGiftReq
}
func (*ActivityAcceptAllGiveGiftReq) ProtoMessageType() ProtoMessageType {
	return "ActivityAcceptAllGiveGiftReq"
}

func (*ActivityAcceptAllGiveGiftRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityAcceptAllGiveGiftRsp
}
func (*ActivityAcceptAllGiveGiftRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityAcceptAllGiveGiftRsp"
}

func (*ActivityGetCanGiveFriendGiftReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGetCanGiveFriendGiftReq
}
func (*ActivityGetCanGiveFriendGiftReq) ProtoMessageType() ProtoMessageType {
	return "ActivityGetCanGiveFriendGiftReq"
}

func (*ActivityGetCanGiveFriendGiftRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGetCanGiveFriendGiftRsp
}
func (*ActivityGetCanGiveFriendGiftRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityGetCanGiveFriendGiftRsp"
}

func (*ActivitySetGiftWishReq) ProtoCommand() ProtoCommand         { return ProtoCommandActivitySetGiftWishReq }
func (*ActivitySetGiftWishReq) ProtoMessageType() ProtoMessageType { return "ActivitySetGiftWishReq" }

func (*ActivitySetGiftWishRsp) ProtoCommand() ProtoCommand         { return ProtoCommandActivitySetGiftWishRsp }
func (*ActivitySetGiftWishRsp) ProtoMessageType() ProtoMessageType { return "ActivitySetGiftWishRsp" }

func (*ActivityGetFriendGiftWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGetFriendGiftWishListReq
}
func (*ActivityGetFriendGiftWishListReq) ProtoMessageType() ProtoMessageType {
	return "ActivityGetFriendGiftWishListReq"
}

func (*ActivityGetFriendGiftWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityGetFriendGiftWishListRsp
}
func (*ActivityGetFriendGiftWishListRsp) ProtoMessageType() ProtoMessageType {
	return "ActivityGetFriendGiftWishListRsp"
}

func (*LuminanceStoneChallengeSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLuminanceStoneChallengeSettleNotify
}
func (*LuminanceStoneChallengeSettleNotify) ProtoMessageType() ProtoMessageType {
	return "LuminanceStoneChallengeSettleNotify"
}

func (*StartRogueDiaryPlayReq) ProtoCommand() ProtoCommand         { return ProtoCommandStartRogueDiaryPlayReq }
func (*StartRogueDiaryPlayReq) ProtoMessageType() ProtoMessageType { return "StartRogueDiaryPlayReq" }

func (*StartRogueDiaryPlayRsp) ProtoCommand() ProtoCommand         { return ProtoCommandStartRogueDiaryPlayRsp }
func (*StartRogueDiaryPlayRsp) ProtoMessageType() ProtoMessageType { return "StartRogueDiaryPlayRsp" }

func (*ResetRogueDiaryPlayReq) ProtoCommand() ProtoCommand         { return ProtoCommandResetRogueDiaryPlayReq }
func (*ResetRogueDiaryPlayReq) ProtoMessageType() ProtoMessageType { return "ResetRogueDiaryPlayReq" }

func (*ResetRogueDiaryPlayRsp) ProtoCommand() ProtoCommand         { return ProtoCommandResetRogueDiaryPlayRsp }
func (*ResetRogueDiaryPlayRsp) ProtoMessageType() ProtoMessageType { return "ResetRogueDiaryPlayRsp" }

func (*EnterRogueDiaryDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterRogueDiaryDungeonReq
}
func (*EnterRogueDiaryDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EnterRogueDiaryDungeonReq"
}

func (*EnterRogueDiaryDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterRogueDiaryDungeonRsp
}
func (*EnterRogueDiaryDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EnterRogueDiaryDungeonRsp"
}

func (*ResumeRogueDiaryDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandResumeRogueDiaryDungeonReq
}
func (*ResumeRogueDiaryDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ResumeRogueDiaryDungeonReq"
}

func (*ResumeRogueDiaryDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandResumeRogueDiaryDungeonRsp
}
func (*ResumeRogueDiaryDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ResumeRogueDiaryDungeonRsp"
}

func (*RogueDiaryDungeonInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryDungeonInfoNotify
}
func (*RogueDiaryDungeonInfoNotify) ProtoMessageType() ProtoMessageType {
	return "RogueDiaryDungeonInfoNotify"
}

func (*RogueDiaryDungeonSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryDungeonSettleNotify
}
func (*RogueDiaryDungeonSettleNotify) ProtoMessageType() ProtoMessageType {
	return "RogueDiaryDungeonSettleNotify"
}

func (*StartRogueDiaryRoomReq) ProtoCommand() ProtoCommand         { return ProtoCommandStartRogueDiaryRoomReq }
func (*StartRogueDiaryRoomReq) ProtoMessageType() ProtoMessageType { return "StartRogueDiaryRoomReq" }

func (*StartRogueDiaryRoomRsp) ProtoCommand() ProtoCommand         { return ProtoCommandStartRogueDiaryRoomRsp }
func (*StartRogueDiaryRoomRsp) ProtoMessageType() ProtoMessageType { return "StartRogueDiaryRoomRsp" }

func (*RogueDiaryTiredAvatarNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryTiredAvatarNotify
}
func (*RogueDiaryTiredAvatarNotify) ProtoMessageType() ProtoMessageType {
	return "RogueDiaryTiredAvatarNotify"
}

func (*ReserveRogueDiaryAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandReserveRogueDiaryAvatarReq
}
func (*ReserveRogueDiaryAvatarReq) ProtoMessageType() ProtoMessageType {
	return "ReserveRogueDiaryAvatarReq"
}

func (*ReserveRogueDiaryAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandReserveRogueDiaryAvatarRsp
}
func (*ReserveRogueDiaryAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "ReserveRogueDiaryAvatarRsp"
}

func (*GetRogueDairyRepairInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRogueDairyRepairInfoReq
}
func (*GetRogueDairyRepairInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetRogueDairyRepairInfoReq"
}

func (*GetRogueDairyRepairInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRogueDairyRepairInfoRsp
}
func (*GetRogueDairyRepairInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetRogueDairyRepairInfoRsp"
}

func (*RefreshRogueDiaryCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshRogueDiaryCardReq
}
func (*RefreshRogueDiaryCardReq) ProtoMessageType() ProtoMessageType {
	return "RefreshRogueDiaryCardReq"
}

func (*RefreshRogueDiaryCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshRogueDiaryCardRsp
}
func (*RefreshRogueDiaryCardRsp) ProtoMessageType() ProtoMessageType {
	return "RefreshRogueDiaryCardRsp"
}

func (*RogueFinishRepairReq) ProtoCommand() ProtoCommand         { return ProtoCommandRogueFinishRepairReq }
func (*RogueFinishRepairReq) ProtoMessageType() ProtoMessageType { return "RogueFinishRepairReq" }

func (*RogueFinishRepairRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRogueFinishRepairRsp }
func (*RogueFinishRepairRsp) ProtoMessageType() ProtoMessageType { return "RogueFinishRepairRsp" }

func (*TryInterruptRogueDiaryDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTryInterruptRogueDiaryDungeonReq
}
func (*TryInterruptRogueDiaryDungeonReq) ProtoMessageType() ProtoMessageType {
	return "TryInterruptRogueDiaryDungeonReq"
}

func (*TryInterruptRogueDiaryDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTryInterruptRogueDiaryDungeonRsp
}
func (*TryInterruptRogueDiaryDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "TryInterruptRogueDiaryDungeonRsp"
}

func (*RogueDiaryRepairInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryRepairInfoNotify
}
func (*RogueDiaryRepairInfoNotify) ProtoMessageType() ProtoMessageType {
	return "RogueDiaryRepairInfoNotify"
}

func (*RetryCurRogueDiaryDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRetryCurRogueDiaryDungeonReq
}
func (*RetryCurRogueDiaryDungeonReq) ProtoMessageType() ProtoMessageType {
	return "RetryCurRogueDiaryDungeonReq"
}

func (*RetryCurRogueDiaryDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRetryCurRogueDiaryDungeonRsp
}
func (*RetryCurRogueDiaryDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "RetryCurRogueDiaryDungeonRsp"
}

func (*RogueDiaryReviveAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryReviveAvatarReq
}
func (*RogueDiaryReviveAvatarReq) ProtoMessageType() ProtoMessageType {
	return "RogueDiaryReviveAvatarReq"
}

func (*RogueDiaryReviveAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryReviveAvatarRsp
}
func (*RogueDiaryReviveAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "RogueDiaryReviveAvatarRsp"
}

func (*TryEnterNextRogueDiaryDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTryEnterNextRogueDiaryDungeonReq
}
func (*TryEnterNextRogueDiaryDungeonReq) ProtoMessageType() ProtoMessageType {
	return "TryEnterNextRogueDiaryDungeonReq"
}

func (*TryEnterNextRogueDiaryDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTryEnterNextRogueDiaryDungeonRsp
}
func (*TryEnterNextRogueDiaryDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "TryEnterNextRogueDiaryDungeonRsp"
}

func (*RogueDiaryCoinAddNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRogueDiaryCoinAddNotify
}
func (*RogueDiaryCoinAddNotify) ProtoMessageType() ProtoMessageType { return "RogueDiaryCoinAddNotify" }

func (*SummerTimeV2BoatSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeV2BoatSettleNotify
}
func (*SummerTimeV2BoatSettleNotify) ProtoMessageType() ProtoMessageType {
	return "SummerTimeV2BoatSettleNotify"
}

func (*ActivityPushTipsInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityPushTipsInfoNotify
}
func (*ActivityPushTipsInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ActivityPushTipsInfoNotify"
}

func (*ActivityReadPushTipsReq) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityReadPushTipsReq
}
func (*ActivityReadPushTipsReq) ProtoMessageType() ProtoMessageType { return "ActivityReadPushTipsReq" }

func (*ActivityReadPushTipsRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandActivityReadPushTipsRsp
}
func (*ActivityReadPushTipsRsp) ProtoMessageType() ProtoMessageType { return "ActivityReadPushTipsRsp" }

func (*SummerTimeV2RestartBoatGalleryReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeV2RestartBoatGalleryReq
}
func (*SummerTimeV2RestartBoatGalleryReq) ProtoMessageType() ProtoMessageType {
	return "SummerTimeV2RestartBoatGalleryReq"
}

func (*SummerTimeV2RestartBoatGalleryRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeV2RestartBoatGalleryRsp
}
func (*SummerTimeV2RestartBoatGalleryRsp) ProtoMessageType() ProtoMessageType {
	return "SummerTimeV2RestartBoatGalleryRsp"
}

func (*SummerTimeV2RestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeV2RestartDungeonReq
}
func (*SummerTimeV2RestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "SummerTimeV2RestartDungeonReq"
}

func (*SummerTimeV2RestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSummerTimeV2RestartDungeonRsp
}
func (*SummerTimeV2RestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "SummerTimeV2RestartDungeonRsp"
}

func (*IslandPartySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIslandPartySettleNotify
}
func (*IslandPartySettleNotify) ProtoMessageType() ProtoMessageType { return "IslandPartySettleNotify" }

func (*GearActivityStartPlayGearReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityStartPlayGearReq
}
func (*GearActivityStartPlayGearReq) ProtoMessageType() ProtoMessageType {
	return "GearActivityStartPlayGearReq"
}

func (*GearActivityStartPlayGearRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityStartPlayGearRsp
}
func (*GearActivityStartPlayGearRsp) ProtoMessageType() ProtoMessageType {
	return "GearActivityStartPlayGearRsp"
}

func (*GearActivityFinishPlayGearReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityFinishPlayGearReq
}
func (*GearActivityFinishPlayGearReq) ProtoMessageType() ProtoMessageType {
	return "GearActivityFinishPlayGearReq"
}

func (*GearActivityFinishPlayGearRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityFinishPlayGearRsp
}
func (*GearActivityFinishPlayGearRsp) ProtoMessageType() ProtoMessageType {
	return "GearActivityFinishPlayGearRsp"
}

func (*GearActivityStartPlayPictureReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityStartPlayPictureReq
}
func (*GearActivityStartPlayPictureReq) ProtoMessageType() ProtoMessageType {
	return "GearActivityStartPlayPictureReq"
}

func (*GearActivityStartPlayPictureRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityStartPlayPictureRsp
}
func (*GearActivityStartPlayPictureRsp) ProtoMessageType() ProtoMessageType {
	return "GearActivityStartPlayPictureRsp"
}

func (*GearActivityFinishPlayPictureReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityFinishPlayPictureReq
}
func (*GearActivityFinishPlayPictureReq) ProtoMessageType() ProtoMessageType {
	return "GearActivityFinishPlayPictureReq"
}

func (*GearActivityFinishPlayPictureRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGearActivityFinishPlayPictureRsp
}
func (*GearActivityFinishPlayPictureRsp) ProtoMessageType() ProtoMessageType {
	return "GearActivityFinishPlayPictureRsp"
}

func (*GravenInnocenceRaceSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocenceRaceSettleNotify
}
func (*GravenInnocenceRaceSettleNotify) ProtoMessageType() ProtoMessageType {
	return "GravenInnocenceRaceSettleNotify"
}

func (*GravenInnocenceRaceRestartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocenceRaceRestartReq
}
func (*GravenInnocenceRaceRestartReq) ProtoMessageType() ProtoMessageType {
	return "GravenInnocenceRaceRestartReq"
}

func (*GravenInnocenceRaceRestartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocenceRaceRestartRsp
}
func (*GravenInnocenceRaceRestartRsp) ProtoMessageType() ProtoMessageType {
	return "GravenInnocenceRaceRestartRsp"
}

func (*GravenInnocenceEditCarveCombinationReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocenceEditCarveCombinationReq
}
func (*GravenInnocenceEditCarveCombinationReq) ProtoMessageType() ProtoMessageType {
	return "GravenInnocenceEditCarveCombinationReq"
}

func (*GravenInnocenceEditCarveCombinationRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocenceEditCarveCombinationRsp
}
func (*GravenInnocenceEditCarveCombinationRsp) ProtoMessageType() ProtoMessageType {
	return "GravenInnocenceEditCarveCombinationRsp"
}

func (*GravenInnocencePhotoFinishReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocencePhotoFinishReq
}
func (*GravenInnocencePhotoFinishReq) ProtoMessageType() ProtoMessageType {
	return "GravenInnocencePhotoFinishReq"
}

func (*GravenInnocencePhotoFinishRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocencePhotoFinishRsp
}
func (*GravenInnocencePhotoFinishRsp) ProtoMessageType() ProtoMessageType {
	return "GravenInnocencePhotoFinishRsp"
}

func (*GravenInnocencePhotoReminderNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGravenInnocencePhotoReminderNotify
}
func (*GravenInnocencePhotoReminderNotify) ProtoMessageType() ProtoMessageType {
	return "GravenInnocencePhotoReminderNotify"
}

func (*InstableSprayEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSprayEnterDungeonReq
}
func (*InstableSprayEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "InstableSprayEnterDungeonReq"
}

func (*InstableSprayEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSprayEnterDungeonRsp
}
func (*InstableSprayEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "InstableSprayEnterDungeonRsp"
}

func (*InstableSpraySwitchTeamReq) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSpraySwitchTeamReq
}
func (*InstableSpraySwitchTeamReq) ProtoMessageType() ProtoMessageType {
	return "InstableSpraySwitchTeamReq"
}

func (*InstableSpraySwitchTeamRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSpraySwitchTeamRsp
}
func (*InstableSpraySwitchTeamRsp) ProtoMessageType() ProtoMessageType {
	return "InstableSpraySwitchTeamRsp"
}

func (*InstableSprayLevelFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSprayLevelFinishNotify
}
func (*InstableSprayLevelFinishNotify) ProtoMessageType() ProtoMessageType {
	return "InstableSprayLevelFinishNotify"
}

func (*InstableSprayRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSprayRestartDungeonReq
}
func (*InstableSprayRestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "InstableSprayRestartDungeonReq"
}

func (*InstableSprayRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSprayRestartDungeonRsp
}
func (*InstableSprayRestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "InstableSprayRestartDungeonRsp"
}

func (*MuqadasPotionActivityEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionActivityEnterDungeonReq
}
func (*MuqadasPotionActivityEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionActivityEnterDungeonReq"
}

func (*MuqadasPotionActivityEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionActivityEnterDungeonRsp
}
func (*MuqadasPotionActivityEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionActivityEnterDungeonRsp"
}

func (*MuqadasPotionDungeonSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionDungeonSettleNotify
}
func (*MuqadasPotionDungeonSettleNotify) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionDungeonSettleNotify"
}

func (*MuqadasPotionRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionRestartDungeonReq
}
func (*MuqadasPotionRestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionRestartDungeonReq"
}

func (*MuqadasPotionRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionRestartDungeonRsp
}
func (*MuqadasPotionRestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionRestartDungeonRsp"
}

func (*MuqadasPotionCaptureWeaknessReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionCaptureWeaknessReq
}
func (*MuqadasPotionCaptureWeaknessReq) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionCaptureWeaknessReq"
}

func (*MuqadasPotionCaptureWeaknessRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMuqadasPotionCaptureWeaknessRsp
}
func (*MuqadasPotionCaptureWeaknessRsp) ProtoMessageType() ProtoMessageType {
	return "MuqadasPotionCaptureWeaknessRsp"
}

func (*TreasureSeelieCollectOrbsNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureSeelieCollectOrbsNotify
}
func (*TreasureSeelieCollectOrbsNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureSeelieCollectOrbsNotify"
}

func (*VintageMarketDeliverItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketDeliverItemReq
}
func (*VintageMarketDeliverItemReq) ProtoMessageType() ProtoMessageType {
	return "VintageMarketDeliverItemReq"
}

func (*VintageMarketDeliverItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketDeliverItemRsp
}
func (*VintageMarketDeliverItemRsp) ProtoMessageType() ProtoMessageType {
	return "VintageMarketDeliverItemRsp"
}

func (*SceneGalleryVintageHuntingSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSceneGalleryVintageHuntingSettleNotify
}
func (*SceneGalleryVintageHuntingSettleNotify) ProtoMessageType() ProtoMessageType {
	return "SceneGalleryVintageHuntingSettleNotify"
}

func (*VintagePresentFinishNoify) ProtoCommand() ProtoCommand {
	return ProtoCommandVintagePresentFinishNoify
}
func (*VintagePresentFinishNoify) ProtoMessageType() ProtoMessageType {
	return "VintagePresentFinishNoify"
}

func (*VintageDecorateBoothReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageDecorateBoothReq
}
func (*VintageDecorateBoothReq) ProtoMessageType() ProtoMessageType { return "VintageDecorateBoothReq" }

func (*VintageDecorateBoothRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageDecorateBoothRsp
}
func (*VintageDecorateBoothRsp) ProtoMessageType() ProtoMessageType { return "VintageDecorateBoothRsp" }

func (*VintageHuntingStartGalleryReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageHuntingStartGalleryReq
}
func (*VintageHuntingStartGalleryReq) ProtoMessageType() ProtoMessageType {
	return "VintageHuntingStartGalleryReq"
}

func (*VintageHuntingStartGalleryRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageHuntingStartGalleryRsp
}
func (*VintageHuntingStartGalleryRsp) ProtoMessageType() ProtoMessageType {
	return "VintageHuntingStartGalleryRsp"
}

func (*VintageCampGroupBundleRegisterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageCampGroupBundleRegisterNotify
}
func (*VintageCampGroupBundleRegisterNotify) ProtoMessageType() ProtoMessageType {
	return "VintageCampGroupBundleRegisterNotify"
}

func (*VintageCampStageFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageCampStageFinishNotify
}
func (*VintageCampStageFinishNotify) ProtoMessageType() ProtoMessageType {
	return "VintageCampStageFinishNotify"
}

func (*VintageMarketStartStorePlayReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStartStorePlayReq
}
func (*VintageMarketStartStorePlayReq) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStartStorePlayReq"
}

func (*VintageMarketStartStorePlayRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStartStorePlayRsp
}
func (*VintageMarketStartStorePlayRsp) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStartStorePlayRsp"
}

func (*VintageMarketFinishStorePlayReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketFinishStorePlayReq
}
func (*VintageMarketFinishStorePlayReq) ProtoMessageType() ProtoMessageType {
	return "VintageMarketFinishStorePlayReq"
}

func (*VintageMarketFinishStorePlayRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketFinishStorePlayRsp
}
func (*VintageMarketFinishStorePlayRsp) ProtoMessageType() ProtoMessageType {
	return "VintageMarketFinishStorePlayRsp"
}

func (*VintagePresentFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandVintagePresentFinishNotify
}
func (*VintagePresentFinishNotify) ProtoMessageType() ProtoMessageType {
	return "VintagePresentFinishNotify"
}

func (*VintageMarketStoreUnlockSlotReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStoreUnlockSlotReq
}
func (*VintageMarketStoreUnlockSlotReq) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStoreUnlockSlotReq"
}

func (*VintageMarketStoreUnlockSlotRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStoreUnlockSlotRsp
}
func (*VintageMarketStoreUnlockSlotRsp) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStoreUnlockSlotRsp"
}

func (*VintageMarketStoreChooseStrategyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStoreChooseStrategyReq
}
func (*VintageMarketStoreChooseStrategyReq) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStoreChooseStrategyReq"
}

func (*VintageMarketStoreChooseStrategyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStoreChooseStrategyRsp
}
func (*VintageMarketStoreChooseStrategyRsp) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStoreChooseStrategyRsp"
}

func (*VintageMarketStoreViewStrategyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStoreViewStrategyReq
}
func (*VintageMarketStoreViewStrategyReq) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStoreViewStrategyReq"
}

func (*VintageMarketStoreViewStrategyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketStoreViewStrategyRsp
}
func (*VintageMarketStoreViewStrategyRsp) ProtoMessageType() ProtoMessageType {
	return "VintageMarketStoreViewStrategyRsp"
}

func (*VintageMarketDividendFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketDividendFinishNotify
}
func (*VintageMarketDividendFinishNotify) ProtoMessageType() ProtoMessageType {
	return "VintageMarketDividendFinishNotify"
}

func (*VintageMarketNpcEventFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandVintageMarketNpcEventFinishNotify
}
func (*VintageMarketNpcEventFinishNotify) ProtoMessageType() ProtoMessageType {
	return "VintageMarketNpcEventFinishNotify"
}

func (*WindFieldRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandWindFieldRestartDungeonReq
}
func (*WindFieldRestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "WindFieldRestartDungeonReq"
}

func (*WindFieldRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandWindFieldRestartDungeonRsp
}
func (*WindFieldRestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "WindFieldRestartDungeonRsp"
}

func (*EnterFungusFighterPlotDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterFungusFighterPlotDungeonReq
}
func (*EnterFungusFighterPlotDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EnterFungusFighterPlotDungeonReq"
}

func (*EnterFungusFighterPlotDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterFungusFighterPlotDungeonRsp
}
func (*EnterFungusFighterPlotDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EnterFungusFighterPlotDungeonRsp"
}

func (*FungusFighterPlotInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterPlotInfoNotify
}
func (*FungusFighterPlotInfoNotify) ProtoMessageType() ProtoMessageType {
	return "FungusFighterPlotInfoNotify"
}

func (*FungusCultivateReq) ProtoCommand() ProtoCommand         { return ProtoCommandFungusCultivateReq }
func (*FungusCultivateReq) ProtoMessageType() ProtoMessageType { return "FungusCultivateReq" }

func (*FungusCultivateRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFungusCultivateRsp }
func (*FungusCultivateRsp) ProtoMessageType() ProtoMessageType { return "FungusCultivateRsp" }

func (*FungusRenameReq) ProtoCommand() ProtoCommand         { return ProtoCommandFungusRenameReq }
func (*FungusRenameReq) ProtoMessageType() ProtoMessageType { return "FungusRenameReq" }

func (*FungusRenameRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFungusRenameRsp }
func (*FungusRenameRsp) ProtoMessageType() ProtoMessageType { return "FungusRenameRsp" }

func (*EnterFungusFighterTrainingDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterFungusFighterTrainingDungeonReq
}
func (*EnterFungusFighterTrainingDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EnterFungusFighterTrainingDungeonReq"
}

func (*EnterFungusFighterTrainingDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterFungusFighterTrainingDungeonRsp
}
func (*EnterFungusFighterTrainingDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EnterFungusFighterTrainingDungeonRsp"
}

func (*FungusFighterRuntimeDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterRuntimeDataNotify
}
func (*FungusFighterRuntimeDataNotify) ProtoMessageType() ProtoMessageType {
	return "FungusFighterRuntimeDataNotify"
}

func (*FungusFighterTrainingSelectFungusReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterTrainingSelectFungusReq
}
func (*FungusFighterTrainingSelectFungusReq) ProtoMessageType() ProtoMessageType {
	return "FungusFighterTrainingSelectFungusReq"
}

func (*FungusFighterTrainingSelectFungusRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterTrainingSelectFungusRsp
}
func (*FungusFighterTrainingSelectFungusRsp) ProtoMessageType() ProtoMessageType {
	return "FungusFighterTrainingSelectFungusRsp"
}

func (*FungusFighterTrainingGallerySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterTrainingGallerySettleNotify
}
func (*FungusFighterTrainingGallerySettleNotify) ProtoMessageType() ProtoMessageType {
	return "FungusFighterTrainingGallerySettleNotify"
}

func (*FungusFighterClearTrainingRuntimeDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterClearTrainingRuntimeDataReq
}
func (*FungusFighterClearTrainingRuntimeDataReq) ProtoMessageType() ProtoMessageType {
	return "FungusFighterClearTrainingRuntimeDataReq"
}

func (*FungusFighterClearTrainingRuntimeDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterClearTrainingRuntimeDataRsp
}
func (*FungusFighterClearTrainingRuntimeDataRsp) ProtoMessageType() ProtoMessageType {
	return "FungusFighterClearTrainingRuntimeDataRsp"
}

func (*FungusFighterUseBackupFungusReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterUseBackupFungusReq
}
func (*FungusFighterUseBackupFungusReq) ProtoMessageType() ProtoMessageType {
	return "FungusFighterUseBackupFungusReq"
}

func (*FungusFighterUseBackupFungusRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterUseBackupFungusRsp
}
func (*FungusFighterUseBackupFungusRsp) ProtoMessageType() ProtoMessageType {
	return "FungusFighterUseBackupFungusRsp"
}

func (*FungusFighterRestartTraningDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterRestartTraningDungeonReq
}
func (*FungusFighterRestartTraningDungeonReq) ProtoMessageType() ProtoMessageType {
	return "FungusFighterRestartTraningDungeonReq"
}

func (*FungusFighterRestartTraningDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterRestartTraningDungeonRsp
}
func (*FungusFighterRestartTraningDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "FungusFighterRestartTraningDungeonRsp"
}

func (*CharAmusementSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCharAmusementSettleNotify
}
func (*CharAmusementSettleNotify) ProtoMessageType() ProtoMessageType {
	return "CharAmusementSettleNotify"
}

func (*EffigyChallengeV2EnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2EnterDungeonReq
}
func (*EffigyChallengeV2EnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2EnterDungeonReq"
}

func (*EffigyChallengeV2EnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2EnterDungeonRsp
}
func (*EffigyChallengeV2EnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2EnterDungeonRsp"
}

func (*EffigyChallengeV2RestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2RestartDungeonReq
}
func (*EffigyChallengeV2RestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2RestartDungeonReq"
}

func (*EffigyChallengeV2RestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2RestartDungeonRsp
}
func (*EffigyChallengeV2RestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2RestartDungeonRsp"
}

func (*EffigyChallengeV2ChooseSkillReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2ChooseSkillReq
}
func (*EffigyChallengeV2ChooseSkillReq) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2ChooseSkillReq"
}

func (*EffigyChallengeV2ChooseSkillRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2ChooseSkillRsp
}
func (*EffigyChallengeV2ChooseSkillRsp) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2ChooseSkillRsp"
}

func (*EffigyChallengeV2DungeonInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEffigyChallengeV2DungeonInfoNotify
}
func (*EffigyChallengeV2DungeonInfoNotify) ProtoMessageType() ProtoMessageType {
	return "EffigyChallengeV2DungeonInfoNotify"
}

func (*CoinCollectChooseSkillReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectChooseSkillReq
}
func (*CoinCollectChooseSkillReq) ProtoMessageType() ProtoMessageType {
	return "CoinCollectChooseSkillReq"
}

func (*CoinCollectChooseSkillRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectChooseSkillRsp
}
func (*CoinCollectChooseSkillRsp) ProtoMessageType() ProtoMessageType {
	return "CoinCollectChooseSkillRsp"
}

func (*RestartCoinCollectPlaySingleModeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRestartCoinCollectPlaySingleModeReq
}
func (*RestartCoinCollectPlaySingleModeReq) ProtoMessageType() ProtoMessageType {
	return "RestartCoinCollectPlaySingleModeReq"
}

func (*RestartCoinCollectPlaySingleModeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRestartCoinCollectPlaySingleModeRsp
}
func (*RestartCoinCollectPlaySingleModeRsp) ProtoMessageType() ProtoMessageType {
	return "RestartCoinCollectPlaySingleModeRsp"
}

func (*EndCoinCollectPlaySingleModeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEndCoinCollectPlaySingleModeReq
}
func (*EndCoinCollectPlaySingleModeReq) ProtoMessageType() ProtoMessageType {
	return "EndCoinCollectPlaySingleModeReq"
}

func (*EndCoinCollectPlaySingleModeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEndCoinCollectPlaySingleModeRsp
}
func (*EndCoinCollectPlaySingleModeRsp) ProtoMessageType() ProtoMessageType {
	return "EndCoinCollectPlaySingleModeRsp"
}

func (*CoinCollectPrepareReq) ProtoCommand() ProtoCommand         { return ProtoCommandCoinCollectPrepareReq }
func (*CoinCollectPrepareReq) ProtoMessageType() ProtoMessageType { return "CoinCollectPrepareReq" }

func (*CoinCollectPrepareRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCoinCollectPrepareRsp }
func (*CoinCollectPrepareRsp) ProtoMessageType() ProtoMessageType { return "CoinCollectPrepareRsp" }

func (*CoinCollectInterruptPlayReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectInterruptPlayReq
}
func (*CoinCollectInterruptPlayReq) ProtoMessageType() ProtoMessageType {
	return "CoinCollectInterruptPlayReq"
}

func (*CoinCollectInterruptPlayRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectInterruptPlayRsp
}
func (*CoinCollectInterruptPlayRsp) ProtoMessageType() ProtoMessageType {
	return "CoinCollectInterruptPlayRsp"
}

func (*CoinCollectCheckDoubleStartPlayReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectCheckDoubleStartPlayReq
}
func (*CoinCollectCheckDoubleStartPlayReq) ProtoMessageType() ProtoMessageType {
	return "CoinCollectCheckDoubleStartPlayReq"
}

func (*CoinCollectCheckDoubleStartPlayRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectCheckDoubleStartPlayRsp
}
func (*CoinCollectCheckDoubleStartPlayRsp) ProtoMessageType() ProtoMessageType {
	return "CoinCollectCheckDoubleStartPlayRsp"
}

func (*SingleStartBrickBreakerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSingleStartBrickBreakerReq
}
func (*SingleStartBrickBreakerReq) ProtoMessageType() ProtoMessageType {
	return "SingleStartBrickBreakerReq"
}

func (*SingleStartBrickBreakerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSingleStartBrickBreakerRsp
}
func (*SingleStartBrickBreakerRsp) ProtoMessageType() ProtoMessageType {
	return "SingleStartBrickBreakerRsp"
}

func (*SingleRestartBrickBreakerReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSingleRestartBrickBreakerReq
}
func (*SingleRestartBrickBreakerReq) ProtoMessageType() ProtoMessageType {
	return "SingleRestartBrickBreakerReq"
}

func (*SingleRestartBrickBreakerRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSingleRestartBrickBreakerRsp
}
func (*SingleRestartBrickBreakerRsp) ProtoMessageType() ProtoMessageType {
	return "SingleRestartBrickBreakerRsp"
}

func (*BrickBreakerSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerSettleNotify
}
func (*BrickBreakerSettleNotify) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerSettleNotify"
}

func (*BrickBreakerTwiceStartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerTwiceStartReq
}
func (*BrickBreakerTwiceStartReq) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerTwiceStartReq"
}

func (*BrickBreakerTwiceStartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBrickBreakerTwiceStartRsp
}
func (*BrickBreakerTwiceStartRsp) ProtoMessageType() ProtoMessageType {
	return "BrickBreakerTwiceStartRsp"
}

func (*BrickBreakerQuitReq) ProtoCommand() ProtoCommand         { return ProtoCommandBrickBreakerQuitReq }
func (*BrickBreakerQuitReq) ProtoMessageType() ProtoMessageType { return "BrickBreakerQuitReq" }

func (*BrickBreakerQuitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBrickBreakerQuitRsp }
func (*BrickBreakerQuitRsp) ProtoMessageType() ProtoMessageType { return "BrickBreakerQuitRsp" }

func (*LanV3BoatGameStartSingleReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3BoatGameStartSingleReq
}
func (*LanV3BoatGameStartSingleReq) ProtoMessageType() ProtoMessageType {
	return "LanV3BoatGameStartSingleReq"
}

func (*LanV3BoatGameStartSingleRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3BoatGameStartSingleRsp
}
func (*LanV3BoatGameStartSingleRsp) ProtoMessageType() ProtoMessageType {
	return "LanV3BoatGameStartSingleRsp"
}

func (*LanV3BoatGameTransferFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3BoatGameTransferFinishNotify
}
func (*LanV3BoatGameTransferFinishNotify) ProtoMessageType() ProtoMessageType {
	return "LanV3BoatGameTransferFinishNotify"
}

func (*LanV3RaceSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLanV3RaceSettleNotify }
func (*LanV3RaceSettleNotify) ProtoMessageType() ProtoMessageType { return "LanV3RaceSettleNotify" }

func (*LanV3RaceRestartReq) ProtoCommand() ProtoCommand         { return ProtoCommandLanV3RaceRestartReq }
func (*LanV3RaceRestartReq) ProtoMessageType() ProtoMessageType { return "LanV3RaceRestartReq" }

func (*LanV3RaceRestartRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLanV3RaceRestartRsp }
func (*LanV3RaceRestartRsp) ProtoMessageType() ProtoMessageType { return "LanV3RaceRestartRsp" }

func (*LanV3BoatInterruptSettleStageReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3BoatInterruptSettleStageReq
}
func (*LanV3BoatInterruptSettleStageReq) ProtoMessageType() ProtoMessageType {
	return "LanV3BoatInterruptSettleStageReq"
}

func (*LanV3BoatInterruptSettleStageRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3BoatInterruptSettleStageRsp
}
func (*LanV3BoatInterruptSettleStageRsp) ProtoMessageType() ProtoMessageType {
	return "LanV3BoatInterruptSettleStageRsp"
}

func (*LanV3ShadowFinishLevelReq) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3ShadowFinishLevelReq
}
func (*LanV3ShadowFinishLevelReq) ProtoMessageType() ProtoMessageType {
	return "LanV3ShadowFinishLevelReq"
}

func (*LanV3ShadowFinishLevelRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandLanV3ShadowFinishLevelRsp
}
func (*LanV3ShadowFinishLevelRsp) ProtoMessageType() ProtoMessageType {
	return "LanV3ShadowFinishLevelRsp"
}

func (*DuelHeartEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDuelHeartEnterDungeonReq
}
func (*DuelHeartEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "DuelHeartEnterDungeonReq"
}

func (*DuelHeartEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDuelHeartEnterDungeonRsp
}
func (*DuelHeartEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "DuelHeartEnterDungeonRsp"
}

func (*DuelHeartRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDuelHeartRestartDungeonReq
}
func (*DuelHeartRestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "DuelHeartRestartDungeonReq"
}

func (*DuelHeartRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDuelHeartRestartDungeonRsp
}
func (*DuelHeartRestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "DuelHeartRestartDungeonRsp"
}

func (*DuelHeartSelectDifficultyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDuelHeartSelectDifficultyReq
}
func (*DuelHeartSelectDifficultyReq) ProtoMessageType() ProtoMessageType {
	return "DuelHeartSelectDifficultyReq"
}

func (*DuelHeartSelectDifficultyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDuelHeartSelectDifficultyRsp
}
func (*DuelHeartSelectDifficultyRsp) ProtoMessageType() ProtoMessageType {
	return "DuelHeartSelectDifficultyRsp"
}

func (*DuelHeartSettleNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDuelHeartSettleNotify }
func (*DuelHeartSettleNotify) ProtoMessageType() ProtoMessageType { return "DuelHeartSettleNotify" }

func (*DuelHeartCgEndNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDuelHeartCgEndNotify }
func (*DuelHeartCgEndNotify) ProtoMessageType() ProtoMessageType { return "DuelHeartCgEndNotify" }

func (*TeamChainEnterDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainEnterDungeonReq
}
func (*TeamChainEnterDungeonReq) ProtoMessageType() ProtoMessageType {
	return "TeamChainEnterDungeonReq"
}

func (*TeamChainEnterDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainEnterDungeonRsp
}
func (*TeamChainEnterDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "TeamChainEnterDungeonRsp"
}

func (*TeamChainRestartDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainRestartDungeonReq
}
func (*TeamChainRestartDungeonReq) ProtoMessageType() ProtoMessageType {
	return "TeamChainRestartDungeonReq"
}

func (*TeamChainRestartDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainRestartDungeonRsp
}
func (*TeamChainRestartDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "TeamChainRestartDungeonRsp"
}

func (*TeamChainDungeonInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainDungeonInfoNotify
}
func (*TeamChainDungeonInfoNotify) ProtoMessageType() ProtoMessageType {
	return "TeamChainDungeonInfoNotify"
}

func (*TeamChainTakeCostumeRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainTakeCostumeRewardReq
}
func (*TeamChainTakeCostumeRewardReq) ProtoMessageType() ProtoMessageType {
	return "TeamChainTakeCostumeRewardReq"
}

func (*TeamChainTakeCostumeRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamChainTakeCostumeRewardRsp
}
func (*TeamChainTakeCostumeRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TeamChainTakeCostumeRewardRsp"
}

func (*ElectroherculesBattleSelectDifficultyReq) ProtoCommand() ProtoCommand {
	return ProtoCommandElectroherculesBattleSelectDifficultyReq
}
func (*ElectroherculesBattleSelectDifficultyReq) ProtoMessageType() ProtoMessageType {
	return "ElectroherculesBattleSelectDifficultyReq"
}

func (*ElectroherculesBattleSelectDifficultyRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandElectroherculesBattleSelectDifficultyRsp
}
func (*ElectroherculesBattleSelectDifficultyRsp) ProtoMessageType() ProtoMessageType {
	return "ElectroherculesBattleSelectDifficultyRsp"
}

func (*ElectroherculesBattleSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandElectroherculesBattleSettleNotify
}
func (*ElectroherculesBattleSettleNotify) ProtoMessageType() ProtoMessageType {
	return "ElectroherculesBattleSettleNotify"
}
