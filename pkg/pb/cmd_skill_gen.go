
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(UnlockAvatarTalentReq) },
		func() ProtoMessage { return new(UnlockAvatarTalentRsp) },
		func() ProtoMessage { return new(AvatarUnlockTalentNotify) },
		func() ProtoMessage { return new(AvatarSkillDepotChangeNotify) },
		func() ProtoMessage { return new(BigTalentPointConvertReq) },
		func() ProtoMessage { return new(BigTalentPointConvertRsp) },
		func() ProtoMessage { return new(AvatarSkillMaxChargeCountNotify) },
		func() ProtoMessage { return new(AvatarSkillInfoNotify) },
		func() ProtoMessage { return new(ProudSkillUpgradeReq) },
		func() ProtoMessage { return new(ProudSkillUpgradeRsp) },
		func() ProtoMessage { return new(ProudSkillChangeNotify) },
		func() ProtoMessage { return new(AvatarSkillUpgradeReq) },
		func() ProtoMessage { return new(AvatarSkillUpgradeRsp) },
		func() ProtoMessage { return new(AvatarSkillChangeNotify) },
		func() ProtoMessage { return new(ProudSkillExtraLevelNotify) },
		func() ProtoMessage { return new(CanUseSkillNotify) },
		func() ProtoMessage { return new(TeamResonanceChangeNotify) },
	)
}

const (
	ProtoCommandUnlockAvatarTalentReq           ProtoCommand = 1072
	ProtoCommandUnlockAvatarTalentRsp           ProtoCommand = 1098
	ProtoCommandAvatarUnlockTalentNotify        ProtoCommand = 1012
	ProtoCommandAvatarSkillDepotChangeNotify    ProtoCommand = 1035
	ProtoCommandBigTalentPointConvertReq        ProtoCommand = 1007
	ProtoCommandBigTalentPointConvertRsp        ProtoCommand = 1021
	ProtoCommandAvatarSkillMaxChargeCountNotify ProtoCommand = 1003
	ProtoCommandAvatarSkillInfoNotify           ProtoCommand = 1090
	ProtoCommandProudSkillUpgradeReq            ProtoCommand = 1073
	ProtoCommandProudSkillUpgradeRsp            ProtoCommand = 1099
	ProtoCommandProudSkillChangeNotify          ProtoCommand = 1031
	ProtoCommandAvatarSkillUpgradeReq           ProtoCommand = 1075
	ProtoCommandAvatarSkillUpgradeRsp           ProtoCommand = 1048
	ProtoCommandAvatarSkillChangeNotify         ProtoCommand = 1097
	ProtoCommandProudSkillExtraLevelNotify      ProtoCommand = 1081
	ProtoCommandCanUseSkillNotify               ProtoCommand = 1005
	ProtoCommandTeamResonanceChangeNotify       ProtoCommand = 1082
)

func (*UnlockAvatarTalentReq) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockAvatarTalentReq }
func (*UnlockAvatarTalentReq) ProtoMessageType() ProtoMessageType { return "UnlockAvatarTalentReq" }

func (*UnlockAvatarTalentRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockAvatarTalentRsp }
func (*UnlockAvatarTalentRsp) ProtoMessageType() ProtoMessageType { return "UnlockAvatarTalentRsp" }

func (*AvatarUnlockTalentNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarUnlockTalentNotify
}
func (*AvatarUnlockTalentNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarUnlockTalentNotify"
}

func (*AvatarSkillDepotChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarSkillDepotChangeNotify
}
func (*AvatarSkillDepotChangeNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarSkillDepotChangeNotify"
}

func (*BigTalentPointConvertReq) ProtoCommand() ProtoCommand {
	return ProtoCommandBigTalentPointConvertReq
}
func (*BigTalentPointConvertReq) ProtoMessageType() ProtoMessageType {
	return "BigTalentPointConvertReq"
}

func (*BigTalentPointConvertRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandBigTalentPointConvertRsp
}
func (*BigTalentPointConvertRsp) ProtoMessageType() ProtoMessageType {
	return "BigTalentPointConvertRsp"
}

func (*AvatarSkillMaxChargeCountNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarSkillMaxChargeCountNotify
}
func (*AvatarSkillMaxChargeCountNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarSkillMaxChargeCountNotify"
}

func (*AvatarSkillInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarSkillInfoNotify }
func (*AvatarSkillInfoNotify) ProtoMessageType() ProtoMessageType { return "AvatarSkillInfoNotify" }

func (*ProudSkillUpgradeReq) ProtoCommand() ProtoCommand         { return ProtoCommandProudSkillUpgradeReq }
func (*ProudSkillUpgradeReq) ProtoMessageType() ProtoMessageType { return "ProudSkillUpgradeReq" }

func (*ProudSkillUpgradeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandProudSkillUpgradeRsp }
func (*ProudSkillUpgradeRsp) ProtoMessageType() ProtoMessageType { return "ProudSkillUpgradeRsp" }

func (*ProudSkillChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandProudSkillChangeNotify }
func (*ProudSkillChangeNotify) ProtoMessageType() ProtoMessageType { return "ProudSkillChangeNotify" }

func (*AvatarSkillUpgradeReq) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarSkillUpgradeReq }
func (*AvatarSkillUpgradeReq) ProtoMessageType() ProtoMessageType { return "AvatarSkillUpgradeReq" }

func (*AvatarSkillUpgradeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarSkillUpgradeRsp }
func (*AvatarSkillUpgradeRsp) ProtoMessageType() ProtoMessageType { return "AvatarSkillUpgradeRsp" }

func (*AvatarSkillChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarSkillChangeNotify
}
func (*AvatarSkillChangeNotify) ProtoMessageType() ProtoMessageType { return "AvatarSkillChangeNotify" }

func (*ProudSkillExtraLevelNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandProudSkillExtraLevelNotify
}
func (*ProudSkillExtraLevelNotify) ProtoMessageType() ProtoMessageType {
	return "ProudSkillExtraLevelNotify"
}

func (*CanUseSkillNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCanUseSkillNotify }
func (*CanUseSkillNotify) ProtoMessageType() ProtoMessageType { return "CanUseSkillNotify" }

func (*TeamResonanceChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTeamResonanceChangeNotify
}
func (*TeamResonanceChangeNotify) ProtoMessageType() ProtoMessageType {
	return "TeamResonanceChangeNotify"
}
