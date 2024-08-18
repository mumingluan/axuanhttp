
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GadgetInteractReq) },
		func() ProtoMessage { return new(GadgetInteractRsp) },
		func() ProtoMessage { return new(GadgetStateNotify) },
		func() ProtoMessage { return new(WorktopOptionNotify) },
		func() ProtoMessage { return new(SelectWorktopOptionReq) },
		func() ProtoMessage { return new(SelectWorktopOptionRsp) },
		func() ProtoMessage { return new(BossChestActivateNotify) },
		func() ProtoMessage { return new(BlossomChestInfoNotify) },
		func() ProtoMessage { return new(GadgetPlayStartNotify) },
		func() ProtoMessage { return new(GadgetPlayStopNotify) },
		func() ProtoMessage { return new(GadgetPlayDataNotify) },
		func() ProtoMessage { return new(GadgetPlayUidOpNotify) },
		func() ProtoMessage { return new(GadgetGeneralRewardInfoNotify) },
		func() ProtoMessage { return new(GadgetAutoPickDropInfoNotify) },
		func() ProtoMessage { return new(UpdateAbilityCreatedMovingPlatformNotify) },
		func() ProtoMessage { return new(FoundationReq) },
		func() ProtoMessage { return new(FoundationRsp) },
		func() ProtoMessage { return new(FoundationNotify) },
		func() ProtoMessage { return new(GadgetTalkChangeNotify) },
		func() ProtoMessage { return new(GadgetChainLevelUpdateNotify) },
		func() ProtoMessage { return new(GadgetChainLevelChangeNotify) },
		func() ProtoMessage { return new(VehicleInteractReq) },
		func() ProtoMessage { return new(VehicleInteractRsp) },
		func() ProtoMessage { return new(CreateVehicleReq) },
		func() ProtoMessage { return new(CreateVehicleRsp) },
		func() ProtoMessage { return new(RequestLiveInfoReq) },
		func() ProtoMessage { return new(RequestLiveInfoRsp) },
		func() ProtoMessage { return new(LiveStartNotify) },
		func() ProtoMessage { return new(ProjectorOptionReq) },
		func() ProtoMessage { return new(ProjectorOptionRsp) },
		func() ProtoMessage { return new(LiveEndNotify) },
		func() ProtoMessage { return new(VehicleStaminaNotify) },
		func() ProtoMessage { return new(GadgetCustomTreeInfoNotify) },
		func() ProtoMessage { return new(GadgetChangeLevelTagReq) },
		func() ProtoMessage { return new(GadgetChangeLevelTagRsp) },
		func() ProtoMessage { return new(NightCrowGadgetObservationMatchReq) },
		func() ProtoMessage { return new(NightCrowGadgetObservationMatchRsp) },
		func() ProtoMessage { return new(DeshretObeliskChestInfoNotify) },
	)
}

const (
	ProtoCommandGadgetInteractReq                        ProtoCommand = 872
	ProtoCommandGadgetInteractRsp                        ProtoCommand = 898
	ProtoCommandGadgetStateNotify                        ProtoCommand = 812
	ProtoCommandWorktopOptionNotify                      ProtoCommand = 835
	ProtoCommandSelectWorktopOptionReq                   ProtoCommand = 807
	ProtoCommandSelectWorktopOptionRsp                   ProtoCommand = 821
	ProtoCommandBossChestActivateNotify                  ProtoCommand = 803
	ProtoCommandBlossomChestInfoNotify                   ProtoCommand = 890
	ProtoCommandGadgetPlayStartNotify                    ProtoCommand = 873
	ProtoCommandGadgetPlayStopNotify                     ProtoCommand = 899
	ProtoCommandGadgetPlayDataNotify                     ProtoCommand = 831
	ProtoCommandGadgetPlayUidOpNotify                    ProtoCommand = 875
	ProtoCommandGadgetGeneralRewardInfoNotify            ProtoCommand = 848
	ProtoCommandGadgetAutoPickDropInfoNotify             ProtoCommand = 897
	ProtoCommandUpdateAbilityCreatedMovingPlatformNotify ProtoCommand = 881
	ProtoCommandFoundationReq                            ProtoCommand = 805
	ProtoCommandFoundationRsp                            ProtoCommand = 882
	ProtoCommandFoundationNotify                         ProtoCommand = 847
	ProtoCommandGadgetTalkChangeNotify                   ProtoCommand = 839
	ProtoCommandGadgetChainLevelUpdateNotify             ProtoCommand = 853
	ProtoCommandGadgetChainLevelChangeNotify             ProtoCommand = 822
	ProtoCommandVehicleInteractReq                       ProtoCommand = 865
	ProtoCommandVehicleInteractRsp                       ProtoCommand = 804
	ProtoCommandCreateVehicleReq                         ProtoCommand = 893
	ProtoCommandCreateVehicleRsp                         ProtoCommand = 827
	ProtoCommandRequestLiveInfoReq                       ProtoCommand = 894
	ProtoCommandRequestLiveInfoRsp                       ProtoCommand = 888
	ProtoCommandLiveStartNotify                          ProtoCommand = 826
	ProtoCommandProjectorOptionReq                       ProtoCommand = 863
	ProtoCommandProjectorOptionRsp                       ProtoCommand = 895
	ProtoCommandLiveEndNotify                            ProtoCommand = 806
	ProtoCommandVehicleStaminaNotify                     ProtoCommand = 834
	ProtoCommandGadgetCustomTreeInfoNotify               ProtoCommand = 850
	ProtoCommandGadgetChangeLevelTagReq                  ProtoCommand = 843
	ProtoCommandGadgetChangeLevelTagRsp                  ProtoCommand = 874
	ProtoCommandNightCrowGadgetObservationMatchReq       ProtoCommand = 876
	ProtoCommandNightCrowGadgetObservationMatchRsp       ProtoCommand = 846
	ProtoCommandDeshretObeliskChestInfoNotify            ProtoCommand = 841
)

func (*GadgetInteractReq) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetInteractReq }
func (*GadgetInteractReq) ProtoMessageType() ProtoMessageType { return "GadgetInteractReq" }

func (*GadgetInteractRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetInteractRsp }
func (*GadgetInteractRsp) ProtoMessageType() ProtoMessageType { return "GadgetInteractRsp" }

func (*GadgetStateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetStateNotify }
func (*GadgetStateNotify) ProtoMessageType() ProtoMessageType { return "GadgetStateNotify" }

func (*WorktopOptionNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWorktopOptionNotify }
func (*WorktopOptionNotify) ProtoMessageType() ProtoMessageType { return "WorktopOptionNotify" }

func (*SelectWorktopOptionReq) ProtoCommand() ProtoCommand         { return ProtoCommandSelectWorktopOptionReq }
func (*SelectWorktopOptionReq) ProtoMessageType() ProtoMessageType { return "SelectWorktopOptionReq" }

func (*SelectWorktopOptionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSelectWorktopOptionRsp }
func (*SelectWorktopOptionRsp) ProtoMessageType() ProtoMessageType { return "SelectWorktopOptionRsp" }

func (*BossChestActivateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBossChestActivateNotify
}
func (*BossChestActivateNotify) ProtoMessageType() ProtoMessageType { return "BossChestActivateNotify" }

func (*BlossomChestInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandBlossomChestInfoNotify }
func (*BlossomChestInfoNotify) ProtoMessageType() ProtoMessageType { return "BlossomChestInfoNotify" }

func (*GadgetPlayStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetPlayStartNotify }
func (*GadgetPlayStartNotify) ProtoMessageType() ProtoMessageType { return "GadgetPlayStartNotify" }

func (*GadgetPlayStopNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetPlayStopNotify }
func (*GadgetPlayStopNotify) ProtoMessageType() ProtoMessageType { return "GadgetPlayStopNotify" }

func (*GadgetPlayDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetPlayDataNotify }
func (*GadgetPlayDataNotify) ProtoMessageType() ProtoMessageType { return "GadgetPlayDataNotify" }

func (*GadgetPlayUidOpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetPlayUidOpNotify }
func (*GadgetPlayUidOpNotify) ProtoMessageType() ProtoMessageType { return "GadgetPlayUidOpNotify" }

func (*GadgetGeneralRewardInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetGeneralRewardInfoNotify
}
func (*GadgetGeneralRewardInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GadgetGeneralRewardInfoNotify"
}

func (*GadgetAutoPickDropInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetAutoPickDropInfoNotify
}
func (*GadgetAutoPickDropInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GadgetAutoPickDropInfoNotify"
}

func (*UpdateAbilityCreatedMovingPlatformNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdateAbilityCreatedMovingPlatformNotify
}
func (*UpdateAbilityCreatedMovingPlatformNotify) ProtoMessageType() ProtoMessageType {
	return "UpdateAbilityCreatedMovingPlatformNotify"
}

func (*FoundationReq) ProtoCommand() ProtoCommand         { return ProtoCommandFoundationReq }
func (*FoundationReq) ProtoMessageType() ProtoMessageType { return "FoundationReq" }

func (*FoundationRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFoundationRsp }
func (*FoundationRsp) ProtoMessageType() ProtoMessageType { return "FoundationRsp" }

func (*FoundationNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFoundationNotify }
func (*FoundationNotify) ProtoMessageType() ProtoMessageType { return "FoundationNotify" }

func (*GadgetTalkChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGadgetTalkChangeNotify }
func (*GadgetTalkChangeNotify) ProtoMessageType() ProtoMessageType { return "GadgetTalkChangeNotify" }

func (*GadgetChainLevelUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetChainLevelUpdateNotify
}
func (*GadgetChainLevelUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "GadgetChainLevelUpdateNotify"
}

func (*GadgetChainLevelChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetChainLevelChangeNotify
}
func (*GadgetChainLevelChangeNotify) ProtoMessageType() ProtoMessageType {
	return "GadgetChainLevelChangeNotify"
}

func (*VehicleInteractReq) ProtoCommand() ProtoCommand         { return ProtoCommandVehicleInteractReq }
func (*VehicleInteractReq) ProtoMessageType() ProtoMessageType { return "VehicleInteractReq" }

func (*VehicleInteractRsp) ProtoCommand() ProtoCommand         { return ProtoCommandVehicleInteractRsp }
func (*VehicleInteractRsp) ProtoMessageType() ProtoMessageType { return "VehicleInteractRsp" }

func (*CreateVehicleReq) ProtoCommand() ProtoCommand         { return ProtoCommandCreateVehicleReq }
func (*CreateVehicleReq) ProtoMessageType() ProtoMessageType { return "CreateVehicleReq" }

func (*CreateVehicleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCreateVehicleRsp }
func (*CreateVehicleRsp) ProtoMessageType() ProtoMessageType { return "CreateVehicleRsp" }

func (*RequestLiveInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandRequestLiveInfoReq }
func (*RequestLiveInfoReq) ProtoMessageType() ProtoMessageType { return "RequestLiveInfoReq" }

func (*RequestLiveInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRequestLiveInfoRsp }
func (*RequestLiveInfoRsp) ProtoMessageType() ProtoMessageType { return "RequestLiveInfoRsp" }

func (*LiveStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLiveStartNotify }
func (*LiveStartNotify) ProtoMessageType() ProtoMessageType { return "LiveStartNotify" }

func (*ProjectorOptionReq) ProtoCommand() ProtoCommand         { return ProtoCommandProjectorOptionReq }
func (*ProjectorOptionReq) ProtoMessageType() ProtoMessageType { return "ProjectorOptionReq" }

func (*ProjectorOptionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandProjectorOptionRsp }
func (*ProjectorOptionRsp) ProtoMessageType() ProtoMessageType { return "ProjectorOptionRsp" }

func (*LiveEndNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLiveEndNotify }
func (*LiveEndNotify) ProtoMessageType() ProtoMessageType { return "LiveEndNotify" }

func (*VehicleStaminaNotify) ProtoCommand() ProtoCommand         { return ProtoCommandVehicleStaminaNotify }
func (*VehicleStaminaNotify) ProtoMessageType() ProtoMessageType { return "VehicleStaminaNotify" }

func (*GadgetCustomTreeInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetCustomTreeInfoNotify
}
func (*GadgetCustomTreeInfoNotify) ProtoMessageType() ProtoMessageType {
	return "GadgetCustomTreeInfoNotify"
}

func (*GadgetChangeLevelTagReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetChangeLevelTagReq
}
func (*GadgetChangeLevelTagReq) ProtoMessageType() ProtoMessageType { return "GadgetChangeLevelTagReq" }

func (*GadgetChangeLevelTagRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGadgetChangeLevelTagRsp
}
func (*GadgetChangeLevelTagRsp) ProtoMessageType() ProtoMessageType { return "GadgetChangeLevelTagRsp" }

func (*NightCrowGadgetObservationMatchReq) ProtoCommand() ProtoCommand {
	return ProtoCommandNightCrowGadgetObservationMatchReq
}
func (*NightCrowGadgetObservationMatchReq) ProtoMessageType() ProtoMessageType {
	return "NightCrowGadgetObservationMatchReq"
}

func (*NightCrowGadgetObservationMatchRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandNightCrowGadgetObservationMatchRsp
}
func (*NightCrowGadgetObservationMatchRsp) ProtoMessageType() ProtoMessageType {
	return "NightCrowGadgetObservationMatchRsp"
}

func (*DeshretObeliskChestInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDeshretObeliskChestInfoNotify
}
func (*DeshretObeliskChestInfoNotify) ProtoMessageType() ProtoMessageType {
	return "DeshretObeliskChestInfoNotify"
}
