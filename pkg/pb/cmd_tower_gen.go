
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(TowerBriefDataNotify) },
		func() ProtoMessage { return new(TowerFloorRecordChangeNotify) },
		func() ProtoMessage { return new(TowerCurLevelRecordChangeNotify) },
		func() ProtoMessage { return new(TowerDailyRewardProgressChangeNotify) },
		func() ProtoMessage { return new(TowerTeamSelectReq) },
		func() ProtoMessage { return new(TowerTeamSelectRsp) },
		func() ProtoMessage { return new(TowerAllDataReq) },
		func() ProtoMessage { return new(TowerAllDataRsp) },
		func() ProtoMessage { return new(TowerEnterLevelReq) },
		func() ProtoMessage { return new(TowerEnterLevelRsp) },
		func() ProtoMessage { return new(TowerBuffSelectReq) },
		func() ProtoMessage { return new(TowerBuffSelectRsp) },
		func() ProtoMessage { return new(TowerSurrenderReq) },
		func() ProtoMessage { return new(TowerSurrenderRsp) },
		func() ProtoMessage { return new(TowerGetFloorStarRewardReq) },
		func() ProtoMessage { return new(TowerGetFloorStarRewardRsp) },
		func() ProtoMessage { return new(TowerLevelEndNotify) },
		func() ProtoMessage { return new(TowerLevelStarCondNotify) },
		func() ProtoMessage { return new(TowerMiddleLevelChangeTeamNotify) },
		func() ProtoMessage { return new(TowerRecordHandbookReq) },
		func() ProtoMessage { return new(TowerRecordHandbookRsp) },
	)
}

const (
	ProtoCommandTowerBriefDataNotify                 ProtoCommand = 2472
	ProtoCommandTowerFloorRecordChangeNotify         ProtoCommand = 2498
	ProtoCommandTowerCurLevelRecordChangeNotify      ProtoCommand = 2412
	ProtoCommandTowerDailyRewardProgressChangeNotify ProtoCommand = 2435
	ProtoCommandTowerTeamSelectReq                   ProtoCommand = 2421
	ProtoCommandTowerTeamSelectRsp                   ProtoCommand = 2403
	ProtoCommandTowerAllDataReq                      ProtoCommand = 2490
	ProtoCommandTowerAllDataRsp                      ProtoCommand = 2473
	ProtoCommandTowerEnterLevelReq                   ProtoCommand = 2431
	ProtoCommandTowerEnterLevelRsp                   ProtoCommand = 2475
	ProtoCommandTowerBuffSelectReq                   ProtoCommand = 2448
	ProtoCommandTowerBuffSelectRsp                   ProtoCommand = 2497
	ProtoCommandTowerSurrenderReq                    ProtoCommand = 2422
	ProtoCommandTowerSurrenderRsp                    ProtoCommand = 2465
	ProtoCommandTowerGetFloorStarRewardReq           ProtoCommand = 2404
	ProtoCommandTowerGetFloorStarRewardRsp           ProtoCommand = 2493
	ProtoCommandTowerLevelEndNotify                  ProtoCommand = 2495
	ProtoCommandTowerLevelStarCondNotify             ProtoCommand = 2406
	ProtoCommandTowerMiddleLevelChangeTeamNotify     ProtoCommand = 2434
	ProtoCommandTowerRecordHandbookReq               ProtoCommand = 2450
	ProtoCommandTowerRecordHandbookRsp               ProtoCommand = 2443
)

func (*TowerBriefDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandTowerBriefDataNotify }
func (*TowerBriefDataNotify) ProtoMessageType() ProtoMessageType { return "TowerBriefDataNotify" }

func (*TowerFloorRecordChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerFloorRecordChangeNotify
}
func (*TowerFloorRecordChangeNotify) ProtoMessageType() ProtoMessageType {
	return "TowerFloorRecordChangeNotify"
}

func (*TowerCurLevelRecordChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerCurLevelRecordChangeNotify
}
func (*TowerCurLevelRecordChangeNotify) ProtoMessageType() ProtoMessageType {
	return "TowerCurLevelRecordChangeNotify"
}

func (*TowerDailyRewardProgressChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerDailyRewardProgressChangeNotify
}
func (*TowerDailyRewardProgressChangeNotify) ProtoMessageType() ProtoMessageType {
	return "TowerDailyRewardProgressChangeNotify"
}

func (*TowerTeamSelectReq) ProtoCommand() ProtoCommand         { return ProtoCommandTowerTeamSelectReq }
func (*TowerTeamSelectReq) ProtoMessageType() ProtoMessageType { return "TowerTeamSelectReq" }

func (*TowerTeamSelectRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTowerTeamSelectRsp }
func (*TowerTeamSelectRsp) ProtoMessageType() ProtoMessageType { return "TowerTeamSelectRsp" }

func (*TowerAllDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandTowerAllDataReq }
func (*TowerAllDataReq) ProtoMessageType() ProtoMessageType { return "TowerAllDataReq" }

func (*TowerAllDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTowerAllDataRsp }
func (*TowerAllDataRsp) ProtoMessageType() ProtoMessageType { return "TowerAllDataRsp" }

func (*TowerEnterLevelReq) ProtoCommand() ProtoCommand         { return ProtoCommandTowerEnterLevelReq }
func (*TowerEnterLevelReq) ProtoMessageType() ProtoMessageType { return "TowerEnterLevelReq" }

func (*TowerEnterLevelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTowerEnterLevelRsp }
func (*TowerEnterLevelRsp) ProtoMessageType() ProtoMessageType { return "TowerEnterLevelRsp" }

func (*TowerBuffSelectReq) ProtoCommand() ProtoCommand         { return ProtoCommandTowerBuffSelectReq }
func (*TowerBuffSelectReq) ProtoMessageType() ProtoMessageType { return "TowerBuffSelectReq" }

func (*TowerBuffSelectRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTowerBuffSelectRsp }
func (*TowerBuffSelectRsp) ProtoMessageType() ProtoMessageType { return "TowerBuffSelectRsp" }

func (*TowerSurrenderReq) ProtoCommand() ProtoCommand         { return ProtoCommandTowerSurrenderReq }
func (*TowerSurrenderReq) ProtoMessageType() ProtoMessageType { return "TowerSurrenderReq" }

func (*TowerSurrenderRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTowerSurrenderRsp }
func (*TowerSurrenderRsp) ProtoMessageType() ProtoMessageType { return "TowerSurrenderRsp" }

func (*TowerGetFloorStarRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerGetFloorStarRewardReq
}
func (*TowerGetFloorStarRewardReq) ProtoMessageType() ProtoMessageType {
	return "TowerGetFloorStarRewardReq"
}

func (*TowerGetFloorStarRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerGetFloorStarRewardRsp
}
func (*TowerGetFloorStarRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TowerGetFloorStarRewardRsp"
}

func (*TowerLevelEndNotify) ProtoCommand() ProtoCommand         { return ProtoCommandTowerLevelEndNotify }
func (*TowerLevelEndNotify) ProtoMessageType() ProtoMessageType { return "TowerLevelEndNotify" }

func (*TowerLevelStarCondNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerLevelStarCondNotify
}
func (*TowerLevelStarCondNotify) ProtoMessageType() ProtoMessageType {
	return "TowerLevelStarCondNotify"
}

func (*TowerMiddleLevelChangeTeamNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTowerMiddleLevelChangeTeamNotify
}
func (*TowerMiddleLevelChangeTeamNotify) ProtoMessageType() ProtoMessageType {
	return "TowerMiddleLevelChangeTeamNotify"
}

func (*TowerRecordHandbookReq) ProtoCommand() ProtoCommand         { return ProtoCommandTowerRecordHandbookReq }
func (*TowerRecordHandbookReq) ProtoMessageType() ProtoMessageType { return "TowerRecordHandbookReq" }

func (*TowerRecordHandbookRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTowerRecordHandbookRsp }
func (*TowerRecordHandbookRsp) ProtoMessageType() ProtoMessageType { return "TowerRecordHandbookRsp" }
