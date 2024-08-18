
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(NpcTalkReq) },
		func() ProtoMessage { return new(NpcTalkRsp) },
		func() ProtoMessage { return new(GetSceneNpcPositionReq) },
		func() ProtoMessage { return new(GetSceneNpcPositionRsp) },
		func() ProtoMessage { return new(MetNpcIdListNotify) },
		func() ProtoMessage { return new(MeetNpcReq) },
		func() ProtoMessage { return new(MeetNpcRsp) },
		func() ProtoMessage { return new(FinishedTalkIdListNotify) },
	)
}

const (
	ProtoCommandNpcTalkReq               ProtoCommand = 572
	ProtoCommandNpcTalkRsp               ProtoCommand = 598
	ProtoCommandGetSceneNpcPositionReq   ProtoCommand = 535
	ProtoCommandGetSceneNpcPositionRsp   ProtoCommand = 507
	ProtoCommandMetNpcIdListNotify       ProtoCommand = 521
	ProtoCommandMeetNpcReq               ProtoCommand = 503
	ProtoCommandMeetNpcRsp               ProtoCommand = 590
	ProtoCommandFinishedTalkIdListNotify ProtoCommand = 573
)

func (*NpcTalkReq) ProtoCommand() ProtoCommand         { return ProtoCommandNpcTalkReq }
func (*NpcTalkReq) ProtoMessageType() ProtoMessageType { return "NpcTalkReq" }

func (*NpcTalkRsp) ProtoCommand() ProtoCommand         { return ProtoCommandNpcTalkRsp }
func (*NpcTalkRsp) ProtoMessageType() ProtoMessageType { return "NpcTalkRsp" }

func (*GetSceneNpcPositionReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetSceneNpcPositionReq }
func (*GetSceneNpcPositionReq) ProtoMessageType() ProtoMessageType { return "GetSceneNpcPositionReq" }

func (*GetSceneNpcPositionRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetSceneNpcPositionRsp }
func (*GetSceneNpcPositionRsp) ProtoMessageType() ProtoMessageType { return "GetSceneNpcPositionRsp" }

func (*MetNpcIdListNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMetNpcIdListNotify }
func (*MetNpcIdListNotify) ProtoMessageType() ProtoMessageType { return "MetNpcIdListNotify" }

func (*MeetNpcReq) ProtoCommand() ProtoCommand         { return ProtoCommandMeetNpcReq }
func (*MeetNpcReq) ProtoMessageType() ProtoMessageType { return "MeetNpcReq" }

func (*MeetNpcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMeetNpcRsp }
func (*MeetNpcRsp) ProtoMessageType() ProtoMessageType { return "MeetNpcRsp" }

func (*FinishedTalkIdListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFinishedTalkIdListNotify
}
func (*FinishedTalkIdListNotify) ProtoMessageType() ProtoMessageType {
	return "FinishedTalkIdListNotify"
}
