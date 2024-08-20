
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(OpActivityStateNotify) },
		func() ProtoMessage { return new(SignInInfoReq) },
		func() ProtoMessage { return new(SignInInfoRsp) },
		func() ProtoMessage { return new(GetSignInRewardReq) },
		func() ProtoMessage { return new(GetSignInRewardRsp) },
		func() ProtoMessage { return new(BonusActivityUpdateNotify) },
		func() ProtoMessage { return new(BonusActivityInfoReq) },
		func() ProtoMessage { return new(BonusActivityInfoRsp) },
		func() ProtoMessage { return new(GetBonusActivityRewardReq) },
		func() ProtoMessage { return new(GetBonusActivityRewardRsp) },
	)
}

const (
	ProtoCommandOpActivityStateNotify     ProtoCommand = 2572
	ProtoCommandSignInInfoReq             ProtoCommand = 2512
	ProtoCommandSignInInfoRsp             ProtoCommand = 2535
	ProtoCommandGetSignInRewardReq        ProtoCommand = 2507
	ProtoCommandGetSignInRewardRsp        ProtoCommand = 2521
	ProtoCommandBonusActivityUpdateNotify ProtoCommand = 2575
	ProtoCommandBonusActivityInfoReq      ProtoCommand = 2548
	ProtoCommandBonusActivityInfoRsp      ProtoCommand = 2597
	ProtoCommandGetBonusActivityRewardReq ProtoCommand = 2581
	ProtoCommandGetBonusActivityRewardRsp ProtoCommand = 2505
)

func (*OpActivityStateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOpActivityStateNotify }
func (*OpActivityStateNotify) ProtoMessageType() ProtoMessageType { return "OpActivityStateNotify" }

func (*SignInInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandSignInInfoReq }
func (*SignInInfoReq) ProtoMessageType() ProtoMessageType { return "SignInInfoReq" }

func (*SignInInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSignInInfoRsp }
func (*SignInInfoRsp) ProtoMessageType() ProtoMessageType { return "SignInInfoRsp" }

func (*GetSignInRewardReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetSignInRewardReq }
func (*GetSignInRewardReq) ProtoMessageType() ProtoMessageType { return "GetSignInRewardReq" }

func (*GetSignInRewardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetSignInRewardRsp }
func (*GetSignInRewardRsp) ProtoMessageType() ProtoMessageType { return "GetSignInRewardRsp" }

func (*BonusActivityUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBonusActivityUpdateNotify
}
func (*BonusActivityUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "BonusActivityUpdateNotify"
}

func (*BonusActivityInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandBonusActivityInfoReq }
func (*BonusActivityInfoReq) ProtoMessageType() ProtoMessageType { return "BonusActivityInfoReq" }

func (*BonusActivityInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBonusActivityInfoRsp }
func (*BonusActivityInfoRsp) ProtoMessageType() ProtoMessageType { return "BonusActivityInfoRsp" }

func (*GetBonusActivityRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetBonusActivityRewardReq
}
func (*GetBonusActivityRewardReq) ProtoMessageType() ProtoMessageType {
	return "GetBonusActivityRewardReq"
}

func (*GetBonusActivityRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetBonusActivityRewardRsp
}
func (*GetBonusActivityRewardRsp) ProtoMessageType() ProtoMessageType {
	return "GetBonusActivityRewardRsp"
}