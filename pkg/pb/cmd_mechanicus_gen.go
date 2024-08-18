
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetMechanicusInfoReq) },
		func() ProtoMessage { return new(GetMechanicusInfoRsp) },
		func() ProtoMessage { return new(MechanicusSequenceOpenNotify) },
		func() ProtoMessage { return new(MechanicusCoinNotify) },
		func() ProtoMessage { return new(MechanicusOpenNotify) },
		func() ProtoMessage { return new(MechanicusCloseNotify) },
		func() ProtoMessage { return new(MechanicusUnlockGearReq) },
		func() ProtoMessage { return new(MechanicusUnlockGearRsp) },
		func() ProtoMessage { return new(MechanicusLevelupGearReq) },
		func() ProtoMessage { return new(MechanicusLevelupGearRsp) },
		func() ProtoMessage { return new(EnterMechanicusDungeonReq) },
		func() ProtoMessage { return new(EnterMechanicusDungeonRsp) },
		func() ProtoMessage { return new(MechanicusCandidateTeamCreateReq) },
		func() ProtoMessage { return new(MechanicusCandidateTeamCreateRsp) },
	)
}

const (
	ProtoCommandGetMechanicusInfoReq             ProtoCommand = 3972
	ProtoCommandGetMechanicusInfoRsp             ProtoCommand = 3998
	ProtoCommandMechanicusSequenceOpenNotify     ProtoCommand = 3912
	ProtoCommandMechanicusCoinNotify             ProtoCommand = 3935
	ProtoCommandMechanicusOpenNotify             ProtoCommand = 3907
	ProtoCommandMechanicusCloseNotify            ProtoCommand = 3921
	ProtoCommandMechanicusUnlockGearReq          ProtoCommand = 3903
	ProtoCommandMechanicusUnlockGearRsp          ProtoCommand = 3990
	ProtoCommandMechanicusLevelupGearReq         ProtoCommand = 3973
	ProtoCommandMechanicusLevelupGearRsp         ProtoCommand = 3999
	ProtoCommandEnterMechanicusDungeonReq        ProtoCommand = 3931
	ProtoCommandEnterMechanicusDungeonRsp        ProtoCommand = 3975
	ProtoCommandMechanicusCandidateTeamCreateReq ProtoCommand = 3981
	ProtoCommandMechanicusCandidateTeamCreateRsp ProtoCommand = 3905
)

func (*GetMechanicusInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetMechanicusInfoReq }
func (*GetMechanicusInfoReq) ProtoMessageType() ProtoMessageType { return "GetMechanicusInfoReq" }

func (*GetMechanicusInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetMechanicusInfoRsp }
func (*GetMechanicusInfoRsp) ProtoMessageType() ProtoMessageType { return "GetMechanicusInfoRsp" }

func (*MechanicusSequenceOpenNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusSequenceOpenNotify
}
func (*MechanicusSequenceOpenNotify) ProtoMessageType() ProtoMessageType {
	return "MechanicusSequenceOpenNotify"
}

func (*MechanicusCoinNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMechanicusCoinNotify }
func (*MechanicusCoinNotify) ProtoMessageType() ProtoMessageType { return "MechanicusCoinNotify" }

func (*MechanicusOpenNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMechanicusOpenNotify }
func (*MechanicusOpenNotify) ProtoMessageType() ProtoMessageType { return "MechanicusOpenNotify" }

func (*MechanicusCloseNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMechanicusCloseNotify }
func (*MechanicusCloseNotify) ProtoMessageType() ProtoMessageType { return "MechanicusCloseNotify" }

func (*MechanicusUnlockGearReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusUnlockGearReq
}
func (*MechanicusUnlockGearReq) ProtoMessageType() ProtoMessageType { return "MechanicusUnlockGearReq" }

func (*MechanicusUnlockGearRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusUnlockGearRsp
}
func (*MechanicusUnlockGearRsp) ProtoMessageType() ProtoMessageType { return "MechanicusUnlockGearRsp" }

func (*MechanicusLevelupGearReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusLevelupGearReq
}
func (*MechanicusLevelupGearReq) ProtoMessageType() ProtoMessageType {
	return "MechanicusLevelupGearReq"
}

func (*MechanicusLevelupGearRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusLevelupGearRsp
}
func (*MechanicusLevelupGearRsp) ProtoMessageType() ProtoMessageType {
	return "MechanicusLevelupGearRsp"
}

func (*EnterMechanicusDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterMechanicusDungeonReq
}
func (*EnterMechanicusDungeonReq) ProtoMessageType() ProtoMessageType {
	return "EnterMechanicusDungeonReq"
}

func (*EnterMechanicusDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandEnterMechanicusDungeonRsp
}
func (*EnterMechanicusDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "EnterMechanicusDungeonRsp"
}

func (*MechanicusCandidateTeamCreateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusCandidateTeamCreateReq
}
func (*MechanicusCandidateTeamCreateReq) ProtoMessageType() ProtoMessageType {
	return "MechanicusCandidateTeamCreateReq"
}

func (*MechanicusCandidateTeamCreateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMechanicusCandidateTeamCreateRsp
}
func (*MechanicusCandidateTeamCreateRsp) ProtoMessageType() ProtoMessageType {
	return "MechanicusCandidateTeamCreateRsp"
}
