
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(EnterFishingReq) },
		func() ProtoMessage { return new(EnterFishingRsp) },
		func() ProtoMessage { return new(StartFishingReq) },
		func() ProtoMessage { return new(StartFishingRsp) },
		func() ProtoMessage { return new(FishCastRodReq) },
		func() ProtoMessage { return new(FishCastRodRsp) },
		func() ProtoMessage { return new(FishChosenNotify) },
		func() ProtoMessage { return new(FishEscapeNotify) },
		func() ProtoMessage { return new(FishBiteReq) },
		func() ProtoMessage { return new(FishBiteRsp) },
		func() ProtoMessage { return new(FishBattleBeginReq) },
		func() ProtoMessage { return new(FishBattleBeginRsp) },
		func() ProtoMessage { return new(FishBattleEndReq) },
		func() ProtoMessage { return new(FishBattleEndRsp) },
		func() ProtoMessage { return new(ExitFishingReq) },
		func() ProtoMessage { return new(ExitFishingRsp) },
		func() ProtoMessage { return new(FishAttractNotify) },
		func() ProtoMessage { return new(FishBaitGoneNotify) },
		func() ProtoMessage { return new(PlayerFishingDataNotify) },
		func() ProtoMessage { return new(FishPoolDataNotify) },
	)
}

const (
	ProtoCommandEnterFishingReq         ProtoCommand = 5826
	ProtoCommandEnterFishingRsp         ProtoCommand = 5818
	ProtoCommandStartFishingReq         ProtoCommand = 5825
	ProtoCommandStartFishingRsp         ProtoCommand = 5807
	ProtoCommandFishCastRodReq          ProtoCommand = 5802
	ProtoCommandFishCastRodRsp          ProtoCommand = 5831
	ProtoCommandFishChosenNotify        ProtoCommand = 5829
	ProtoCommandFishEscapeNotify        ProtoCommand = 5822
	ProtoCommandFishBiteReq             ProtoCommand = 5844
	ProtoCommandFishBiteRsp             ProtoCommand = 5849
	ProtoCommandFishBattleBeginReq      ProtoCommand = 5820
	ProtoCommandFishBattleBeginRsp      ProtoCommand = 5845
	ProtoCommandFishBattleEndReq        ProtoCommand = 5841
	ProtoCommandFishBattleEndRsp        ProtoCommand = 5842
	ProtoCommandExitFishingReq          ProtoCommand = 5814
	ProtoCommandExitFishingRsp          ProtoCommand = 5847
	ProtoCommandFishAttractNotify       ProtoCommand = 5837
	ProtoCommandFishBaitGoneNotify      ProtoCommand = 5823
	ProtoCommandPlayerFishingDataNotify ProtoCommand = 5835
	ProtoCommandFishPoolDataNotify      ProtoCommand = 5848
)

func (*EnterFishingReq) ProtoCommand() ProtoCommand         { return ProtoCommandEnterFishingReq }
func (*EnterFishingReq) ProtoMessageType() ProtoMessageType { return "EnterFishingReq" }

func (*EnterFishingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEnterFishingRsp }
func (*EnterFishingRsp) ProtoMessageType() ProtoMessageType { return "EnterFishingRsp" }

func (*StartFishingReq) ProtoCommand() ProtoCommand         { return ProtoCommandStartFishingReq }
func (*StartFishingReq) ProtoMessageType() ProtoMessageType { return "StartFishingReq" }

func (*StartFishingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandStartFishingRsp }
func (*StartFishingRsp) ProtoMessageType() ProtoMessageType { return "StartFishingRsp" }

func (*FishCastRodReq) ProtoCommand() ProtoCommand         { return ProtoCommandFishCastRodReq }
func (*FishCastRodReq) ProtoMessageType() ProtoMessageType { return "FishCastRodReq" }

func (*FishCastRodRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFishCastRodRsp }
func (*FishCastRodRsp) ProtoMessageType() ProtoMessageType { return "FishCastRodRsp" }

func (*FishChosenNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFishChosenNotify }
func (*FishChosenNotify) ProtoMessageType() ProtoMessageType { return "FishChosenNotify" }

func (*FishEscapeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFishEscapeNotify }
func (*FishEscapeNotify) ProtoMessageType() ProtoMessageType { return "FishEscapeNotify" }

func (*FishBiteReq) ProtoCommand() ProtoCommand         { return ProtoCommandFishBiteReq }
func (*FishBiteReq) ProtoMessageType() ProtoMessageType { return "FishBiteReq" }

func (*FishBiteRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFishBiteRsp }
func (*FishBiteRsp) ProtoMessageType() ProtoMessageType { return "FishBiteRsp" }

func (*FishBattleBeginReq) ProtoCommand() ProtoCommand         { return ProtoCommandFishBattleBeginReq }
func (*FishBattleBeginReq) ProtoMessageType() ProtoMessageType { return "FishBattleBeginReq" }

func (*FishBattleBeginRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFishBattleBeginRsp }
func (*FishBattleBeginRsp) ProtoMessageType() ProtoMessageType { return "FishBattleBeginRsp" }

func (*FishBattleEndReq) ProtoCommand() ProtoCommand         { return ProtoCommandFishBattleEndReq }
func (*FishBattleEndReq) ProtoMessageType() ProtoMessageType { return "FishBattleEndReq" }

func (*FishBattleEndRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFishBattleEndRsp }
func (*FishBattleEndRsp) ProtoMessageType() ProtoMessageType { return "FishBattleEndRsp" }

func (*ExitFishingReq) ProtoCommand() ProtoCommand         { return ProtoCommandExitFishingReq }
func (*ExitFishingReq) ProtoMessageType() ProtoMessageType { return "ExitFishingReq" }

func (*ExitFishingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandExitFishingRsp }
func (*ExitFishingRsp) ProtoMessageType() ProtoMessageType { return "ExitFishingRsp" }

func (*FishAttractNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFishAttractNotify }
func (*FishAttractNotify) ProtoMessageType() ProtoMessageType { return "FishAttractNotify" }

func (*FishBaitGoneNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFishBaitGoneNotify }
func (*FishBaitGoneNotify) ProtoMessageType() ProtoMessageType { return "FishBaitGoneNotify" }

func (*PlayerFishingDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerFishingDataNotify
}
func (*PlayerFishingDataNotify) ProtoMessageType() ProtoMessageType { return "PlayerFishingDataNotify" }

func (*FishPoolDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFishPoolDataNotify }
func (*FishPoolDataNotify) ProtoMessageType() ProtoMessageType { return "FishPoolDataNotify" }
