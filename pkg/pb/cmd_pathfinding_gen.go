
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(QueryPathReq) },
		func() ProtoMessage { return new(QueryPathRsp) },
		func() ProtoMessage { return new(ObstacleModifyNotify) },
		func() ProtoMessage { return new(PathfindingPingNotify) },
		func() ProtoMessage { return new(PathfindingEnterSceneReq) },
		func() ProtoMessage { return new(PathfindingEnterSceneRsp) },
		func() ProtoMessage { return new(GMShowObstacleReq) },
		func() ProtoMessage { return new(GMShowObstacleRsp) },
		func() ProtoMessage { return new(GMShowNavMeshReq) },
		func() ProtoMessage { return new(GMShowNavMeshRsp) },
		func() ProtoMessage { return new(NavMeshStatsNotify) },
	)
}

const (
	ProtoCommandQueryPathReq             ProtoCommand = 2372
	ProtoCommandQueryPathRsp             ProtoCommand = 2398
	ProtoCommandObstacleModifyNotify     ProtoCommand = 2312
	ProtoCommandPathfindingPingNotify    ProtoCommand = 2335
	ProtoCommandPathfindingEnterSceneReq ProtoCommand = 2307
	ProtoCommandPathfindingEnterSceneRsp ProtoCommand = 2321
	ProtoCommandGMShowObstacleReq        ProtoCommand = 2361
	ProtoCommandGMShowObstacleRsp        ProtoCommand = 2329
	ProtoCommandGMShowNavMeshReq         ProtoCommand = 2357
	ProtoCommandGMShowNavMeshRsp         ProtoCommand = 2400
	ProtoCommandNavMeshStatsNotify       ProtoCommand = 2316
)

func (*QueryPathReq) ProtoCommand() ProtoCommand         { return ProtoCommandQueryPathReq }
func (*QueryPathReq) ProtoMessageType() ProtoMessageType { return "QueryPathReq" }

func (*QueryPathRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQueryPathRsp }
func (*QueryPathRsp) ProtoMessageType() ProtoMessageType { return "QueryPathRsp" }

func (*ObstacleModifyNotify) ProtoCommand() ProtoCommand         { return ProtoCommandObstacleModifyNotify }
func (*ObstacleModifyNotify) ProtoMessageType() ProtoMessageType { return "ObstacleModifyNotify" }

func (*PathfindingPingNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPathfindingPingNotify }
func (*PathfindingPingNotify) ProtoMessageType() ProtoMessageType { return "PathfindingPingNotify" }

func (*PathfindingEnterSceneReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPathfindingEnterSceneReq
}
func (*PathfindingEnterSceneReq) ProtoMessageType() ProtoMessageType {
	return "PathfindingEnterSceneReq"
}

func (*PathfindingEnterSceneRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPathfindingEnterSceneRsp
}
func (*PathfindingEnterSceneRsp) ProtoMessageType() ProtoMessageType {
	return "PathfindingEnterSceneRsp"
}

func (*GMShowObstacleReq) ProtoCommand() ProtoCommand         { return ProtoCommandGMShowObstacleReq }
func (*GMShowObstacleReq) ProtoMessageType() ProtoMessageType { return "GMShowObstacleReq" }

func (*GMShowObstacleRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGMShowObstacleRsp }
func (*GMShowObstacleRsp) ProtoMessageType() ProtoMessageType { return "GMShowObstacleRsp" }

func (*GMShowNavMeshReq) ProtoCommand() ProtoCommand         { return ProtoCommandGMShowNavMeshReq }
func (*GMShowNavMeshReq) ProtoMessageType() ProtoMessageType { return "GMShowNavMeshReq" }

func (*GMShowNavMeshRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGMShowNavMeshRsp }
func (*GMShowNavMeshRsp) ProtoMessageType() ProtoMessageType { return "GMShowNavMeshRsp" }

func (*NavMeshStatsNotify) ProtoCommand() ProtoCommand         { return ProtoCommandNavMeshStatsNotify }
func (*NavMeshStatsNotify) ProtoMessageType() ProtoMessageType { return "NavMeshStatsNotify" }
