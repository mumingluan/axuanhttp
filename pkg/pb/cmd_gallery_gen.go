
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GalleryStartNotify) },
		func() ProtoMessage { return new(GalleryBalloonShootNotify) },
		func() ProtoMessage { return new(GalleryBalloonScoreNotify) },
		func() ProtoMessage { return new(GalleryStopNotify) },
		func() ProtoMessage { return new(GalleryFallCatchNotify) },
		func() ProtoMessage { return new(GalleryFallScoreNotify) },
		func() ProtoMessage { return new(GetAllSceneGalleryInfoReq) },
		func() ProtoMessage { return new(GetAllSceneGalleryInfoRsp) },
		func() ProtoMessage { return new(GalleryFlowerCatchNotify) },
		func() ProtoMessage { return new(GalleryPreStartNotify) },
		func() ProtoMessage { return new(GalleryBulletHitNotify) },
		func() ProtoMessage { return new(GalleryBrokenFloorFallNotify) },
		func() ProtoMessage { return new(InterruptGalleryReq) },
		func() ProtoMessage { return new(InterruptGalleryRsp) },
		func() ProtoMessage { return new(SceneGalleryInfoNotify) },
		func() ProtoMessage { return new(GalleryBounceConjuringHitNotify) },
		func() ProtoMessage { return new(GallerySumoKillMonsterNotify) },
		func() ProtoMessage { return new(GalleryCrystalLinkKillMonsterNotify) },
		func() ProtoMessage { return new(GalleryCrystalLinkBuffInfoNotify) },
		func() ProtoMessage { return new(HomeGalleryInPlayingNotify) },
		func() ProtoMessage { return new(GalleryIslandPartyDownHillInfoNotify) },
		func() ProtoMessage { return new(IslandPartyRaftInfoNotify) },
		func() ProtoMessage { return new(IslandPartySailInfoNotify) },
		func() ProtoMessage { return new(BackRebornGalleryReq) },
		func() ProtoMessage { return new(BackRebornGalleryRsp) },
		func() ProtoMessage { return new(GalleryWillStartCountdownNotify) },
		func() ProtoMessage { return new(InstableSprayGalleryInfoNotify) },
		func() ProtoMessage { return new(WindFieldGalleryInfoNotify) },
		func() ProtoMessage { return new(WindFieldGalleryChallengeInfoNotify) },
		func() ProtoMessage { return new(FungusFighterTrainingInfoNotify) },
		func() ProtoMessage { return new(FungusCaptureSettleNotify) },
		func() ProtoMessage { return new(HideAndSeekPlayerCapturedNotify) },
		func() ProtoMessage { return new(CoinCollectGallerySettleNotify) },
	)
}

const (
	ProtoCommandGalleryStartNotify                   ProtoCommand = 5572
	ProtoCommandGalleryBalloonShootNotify            ProtoCommand = 5598
	ProtoCommandGalleryBalloonScoreNotify            ProtoCommand = 5512
	ProtoCommandGalleryStopNotify                    ProtoCommand = 5535
	ProtoCommandGalleryFallCatchNotify               ProtoCommand = 5507
	ProtoCommandGalleryFallScoreNotify               ProtoCommand = 5521
	ProtoCommandGetAllSceneGalleryInfoReq            ProtoCommand = 5503
	ProtoCommandGetAllSceneGalleryInfoRsp            ProtoCommand = 5590
	ProtoCommandGalleryFlowerCatchNotify             ProtoCommand = 5573
	ProtoCommandGalleryPreStartNotify                ProtoCommand = 5599
	ProtoCommandGalleryBulletHitNotify               ProtoCommand = 5531
	ProtoCommandGalleryBrokenFloorFallNotify         ProtoCommand = 5575
	ProtoCommandInterruptGalleryReq                  ProtoCommand = 5548
	ProtoCommandInterruptGalleryRsp                  ProtoCommand = 5597
	ProtoCommandSceneGalleryInfoNotify               ProtoCommand = 5581
	ProtoCommandGalleryBounceConjuringHitNotify      ProtoCommand = 5505
	ProtoCommandGallerySumoKillMonsterNotify         ProtoCommand = 5582
	ProtoCommandGalleryCrystalLinkKillMonsterNotify  ProtoCommand = 5547
	ProtoCommandGalleryCrystalLinkBuffInfoNotify     ProtoCommand = 5539
	ProtoCommandHomeGalleryInPlayingNotify           ProtoCommand = 5553
	ProtoCommandGalleryIslandPartyDownHillInfoNotify ProtoCommand = 5522
	ProtoCommandIslandPartyRaftInfoNotify            ProtoCommand = 5565
	ProtoCommandIslandPartySailInfoNotify            ProtoCommand = 5504
	ProtoCommandBackRebornGalleryReq                 ProtoCommand = 5593
	ProtoCommandBackRebornGalleryRsp                 ProtoCommand = 5527
	ProtoCommandGalleryWillStartCountdownNotify      ProtoCommand = 5594
	ProtoCommandInstableSprayGalleryInfoNotify       ProtoCommand = 5588
	ProtoCommandWindFieldGalleryInfoNotify           ProtoCommand = 5526
	ProtoCommandWindFieldGalleryChallengeInfoNotify  ProtoCommand = 5563
	ProtoCommandFungusFighterTrainingInfoNotify      ProtoCommand = 5595
	ProtoCommandFungusCaptureSettleNotify            ProtoCommand = 5506
	ProtoCommandHideAndSeekPlayerCapturedNotify      ProtoCommand = 5534
	ProtoCommandCoinCollectGallerySettleNotify       ProtoCommand = 5550
)

func (*GalleryStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGalleryStartNotify }
func (*GalleryStartNotify) ProtoMessageType() ProtoMessageType { return "GalleryStartNotify" }

func (*GalleryBalloonShootNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryBalloonShootNotify
}
func (*GalleryBalloonShootNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryBalloonShootNotify"
}

func (*GalleryBalloonScoreNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryBalloonScoreNotify
}
func (*GalleryBalloonScoreNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryBalloonScoreNotify"
}

func (*GalleryStopNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGalleryStopNotify }
func (*GalleryStopNotify) ProtoMessageType() ProtoMessageType { return "GalleryStopNotify" }

func (*GalleryFallCatchNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGalleryFallCatchNotify }
func (*GalleryFallCatchNotify) ProtoMessageType() ProtoMessageType { return "GalleryFallCatchNotify" }

func (*GalleryFallScoreNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGalleryFallScoreNotify }
func (*GalleryFallScoreNotify) ProtoMessageType() ProtoMessageType { return "GalleryFallScoreNotify" }

func (*GetAllSceneGalleryInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllSceneGalleryInfoReq
}
func (*GetAllSceneGalleryInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetAllSceneGalleryInfoReq"
}

func (*GetAllSceneGalleryInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllSceneGalleryInfoRsp
}
func (*GetAllSceneGalleryInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetAllSceneGalleryInfoRsp"
}

func (*GalleryFlowerCatchNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryFlowerCatchNotify
}
func (*GalleryFlowerCatchNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryFlowerCatchNotify"
}

func (*GalleryPreStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGalleryPreStartNotify }
func (*GalleryPreStartNotify) ProtoMessageType() ProtoMessageType { return "GalleryPreStartNotify" }

func (*GalleryBulletHitNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGalleryBulletHitNotify }
func (*GalleryBulletHitNotify) ProtoMessageType() ProtoMessageType { return "GalleryBulletHitNotify" }

func (*GalleryBrokenFloorFallNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryBrokenFloorFallNotify
}
func (*GalleryBrokenFloorFallNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryBrokenFloorFallNotify"
}

func (*InterruptGalleryReq) ProtoCommand() ProtoCommand         { return ProtoCommandInterruptGalleryReq }
func (*InterruptGalleryReq) ProtoMessageType() ProtoMessageType { return "InterruptGalleryReq" }

func (*InterruptGalleryRsp) ProtoCommand() ProtoCommand         { return ProtoCommandInterruptGalleryRsp }
func (*InterruptGalleryRsp) ProtoMessageType() ProtoMessageType { return "InterruptGalleryRsp" }

func (*SceneGalleryInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneGalleryInfoNotify }
func (*SceneGalleryInfoNotify) ProtoMessageType() ProtoMessageType { return "SceneGalleryInfoNotify" }

func (*GalleryBounceConjuringHitNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryBounceConjuringHitNotify
}
func (*GalleryBounceConjuringHitNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryBounceConjuringHitNotify"
}

func (*GallerySumoKillMonsterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGallerySumoKillMonsterNotify
}
func (*GallerySumoKillMonsterNotify) ProtoMessageType() ProtoMessageType {
	return "GallerySumoKillMonsterNotify"
}

func (*GalleryCrystalLinkKillMonsterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryCrystalLinkKillMonsterNotify
}
func (*GalleryCrystalLinkKillMonsterNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryCrystalLinkKillMonsterNotify"
}

func (*GalleryCrystalLinkBuffInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryCrystalLinkBuffInfoNotify
}
func (*GalleryCrystalLinkBuffInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryCrystalLinkBuffInfoNotify"
}

func (*HomeGalleryInPlayingNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHomeGalleryInPlayingNotify
}
func (*HomeGalleryInPlayingNotify) ProtoMessageType() ProtoMessageType {
	return "HomeGalleryInPlayingNotify"
}

func (*GalleryIslandPartyDownHillInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryIslandPartyDownHillInfoNotify
}
func (*GalleryIslandPartyDownHillInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryIslandPartyDownHillInfoNotify"
}

func (*IslandPartyRaftInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIslandPartyRaftInfoNotify
}
func (*IslandPartyRaftInfoNotify) ProtoMessageType() ProtoMessageType {
	return "IslandPartyRaftInfoNotify"
}

func (*IslandPartySailInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandIslandPartySailInfoNotify
}
func (*IslandPartySailInfoNotify) ProtoMessageType() ProtoMessageType {
	return "IslandPartySailInfoNotify"
}

func (*BackRebornGalleryReq) ProtoCommand() ProtoCommand         { return ProtoCommandBackRebornGalleryReq }
func (*BackRebornGalleryReq) ProtoMessageType() ProtoMessageType { return "BackRebornGalleryReq" }

func (*BackRebornGalleryRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBackRebornGalleryRsp }
func (*BackRebornGalleryRsp) ProtoMessageType() ProtoMessageType { return "BackRebornGalleryRsp" }

func (*GalleryWillStartCountdownNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGalleryWillStartCountdownNotify
}
func (*GalleryWillStartCountdownNotify) ProtoMessageType() ProtoMessageType {
	return "GalleryWillStartCountdownNotify"
}

func (*InstableSprayGalleryInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandInstableSprayGalleryInfoNotify
}
func (*InstableSprayGalleryInfoNotify) ProtoMessageType() ProtoMessageType {
	return "InstableSprayGalleryInfoNotify"
}

func (*WindFieldGalleryInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWindFieldGalleryInfoNotify
}
func (*WindFieldGalleryInfoNotify) ProtoMessageType() ProtoMessageType {
	return "WindFieldGalleryInfoNotify"
}

func (*WindFieldGalleryChallengeInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWindFieldGalleryChallengeInfoNotify
}
func (*WindFieldGalleryChallengeInfoNotify) ProtoMessageType() ProtoMessageType {
	return "WindFieldGalleryChallengeInfoNotify"
}

func (*FungusFighterTrainingInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusFighterTrainingInfoNotify
}
func (*FungusFighterTrainingInfoNotify) ProtoMessageType() ProtoMessageType {
	return "FungusFighterTrainingInfoNotify"
}

func (*FungusCaptureSettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFungusCaptureSettleNotify
}
func (*FungusCaptureSettleNotify) ProtoMessageType() ProtoMessageType {
	return "FungusCaptureSettleNotify"
}

func (*HideAndSeekPlayerCapturedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHideAndSeekPlayerCapturedNotify
}
func (*HideAndSeekPlayerCapturedNotify) ProtoMessageType() ProtoMessageType {
	return "HideAndSeekPlayerCapturedNotify"
}

func (*CoinCollectGallerySettleNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCoinCollectGallerySettleNotify
}
func (*CoinCollectGallerySettleNotify) ProtoMessageType() ProtoMessageType {
	return "CoinCollectGallerySettleNotify"
}
