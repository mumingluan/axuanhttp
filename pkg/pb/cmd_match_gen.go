
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerStartMatchReq) },
		func() ProtoMessage { return new(PlayerStartMatchRsp) },
		func() ProtoMessage { return new(PlayerMatchInfoNotify) },
		func() ProtoMessage { return new(PlayerCancelMatchReq) },
		func() ProtoMessage { return new(PlayerCancelMatchRsp) },
		func() ProtoMessage { return new(PlayerMatchStopNotify) },
		func() ProtoMessage { return new(PlayerMatchSuccNotify) },
		func() ProtoMessage { return new(PlayerConfirmMatchReq) },
		func() ProtoMessage { return new(PlayerConfirmMatchRsp) },
		func() ProtoMessage { return new(PlayerAllowEnterMpAfterAgreeMatchNotify) },
		func() ProtoMessage { return new(PlayerMatchAgreedResultNotify) },
		func() ProtoMessage { return new(PlayerApplyEnterMpAfterMatchAgreedNotify) },
		func() ProtoMessage { return new(PlayerGeneralMatchDismissNotify) },
		func() ProtoMessage { return new(PlayerGeneralMatchConfirmNotify) },
		func() ProtoMessage { return new(PlayerGetForceQuitBanInfoReq) },
		func() ProtoMessage { return new(PlayerGetForceQuitBanInfoRsp) },
		func() ProtoMessage { return new(ServerTryCancelGeneralMatchNotify) },
		func() ProtoMessage { return new(PlayerGCGMatchDismissNotify) },
		func() ProtoMessage { return new(PlayerGCGMatchConfirmNotify) },
	)
}

const (
	ProtoCommandPlayerStartMatchReq                      ProtoCommand = 4176
	ProtoCommandPlayerStartMatchRsp                      ProtoCommand = 4168
	ProtoCommandPlayerMatchInfoNotify                    ProtoCommand = 4175
	ProtoCommandPlayerCancelMatchReq                     ProtoCommand = 4157
	ProtoCommandPlayerCancelMatchRsp                     ProtoCommand = 4152
	ProtoCommandPlayerMatchStopNotify                    ProtoCommand = 4181
	ProtoCommandPlayerMatchSuccNotify                    ProtoCommand = 4179
	ProtoCommandPlayerConfirmMatchReq                    ProtoCommand = 4172
	ProtoCommandPlayerConfirmMatchRsp                    ProtoCommand = 4194
	ProtoCommandPlayerAllowEnterMpAfterAgreeMatchNotify  ProtoCommand = 4199
	ProtoCommandPlayerMatchAgreedResultNotify            ProtoCommand = 4170
	ProtoCommandPlayerApplyEnterMpAfterMatchAgreedNotify ProtoCommand = 4195
	ProtoCommandPlayerGeneralMatchDismissNotify          ProtoCommand = 4191
	ProtoCommandPlayerGeneralMatchConfirmNotify          ProtoCommand = 4192
	ProtoCommandPlayerGetForceQuitBanInfoReq             ProtoCommand = 4164
	ProtoCommandPlayerGetForceQuitBanInfoRsp             ProtoCommand = 4197
	ProtoCommandServerTryCancelGeneralMatchNotify        ProtoCommand = 4187
	ProtoCommandPlayerGCGMatchDismissNotify              ProtoCommand = 4173
	ProtoCommandPlayerGCGMatchConfirmNotify              ProtoCommand = 4185
)

func (*PlayerStartMatchReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerStartMatchReq }
func (*PlayerStartMatchReq) ProtoMessageType() ProtoMessageType { return "PlayerStartMatchReq" }

func (*PlayerStartMatchRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerStartMatchRsp }
func (*PlayerStartMatchRsp) ProtoMessageType() ProtoMessageType { return "PlayerStartMatchRsp" }

func (*PlayerMatchInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerMatchInfoNotify }
func (*PlayerMatchInfoNotify) ProtoMessageType() ProtoMessageType { return "PlayerMatchInfoNotify" }

func (*PlayerCancelMatchReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCancelMatchReq }
func (*PlayerCancelMatchReq) ProtoMessageType() ProtoMessageType { return "PlayerCancelMatchReq" }

func (*PlayerCancelMatchRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerCancelMatchRsp }
func (*PlayerCancelMatchRsp) ProtoMessageType() ProtoMessageType { return "PlayerCancelMatchRsp" }

func (*PlayerMatchStopNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerMatchStopNotify }
func (*PlayerMatchStopNotify) ProtoMessageType() ProtoMessageType { return "PlayerMatchStopNotify" }

func (*PlayerMatchSuccNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerMatchSuccNotify }
func (*PlayerMatchSuccNotify) ProtoMessageType() ProtoMessageType { return "PlayerMatchSuccNotify" }

func (*PlayerConfirmMatchReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerConfirmMatchReq }
func (*PlayerConfirmMatchReq) ProtoMessageType() ProtoMessageType { return "PlayerConfirmMatchReq" }

func (*PlayerConfirmMatchRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerConfirmMatchRsp }
func (*PlayerConfirmMatchRsp) ProtoMessageType() ProtoMessageType { return "PlayerConfirmMatchRsp" }

func (*PlayerAllowEnterMpAfterAgreeMatchNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerAllowEnterMpAfterAgreeMatchNotify
}
func (*PlayerAllowEnterMpAfterAgreeMatchNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerAllowEnterMpAfterAgreeMatchNotify"
}

func (*PlayerMatchAgreedResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerMatchAgreedResultNotify
}
func (*PlayerMatchAgreedResultNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerMatchAgreedResultNotify"
}

func (*PlayerApplyEnterMpAfterMatchAgreedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerApplyEnterMpAfterMatchAgreedNotify
}
func (*PlayerApplyEnterMpAfterMatchAgreedNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerApplyEnterMpAfterMatchAgreedNotify"
}

func (*PlayerGeneralMatchDismissNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGeneralMatchDismissNotify
}
func (*PlayerGeneralMatchDismissNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerGeneralMatchDismissNotify"
}

func (*PlayerGeneralMatchConfirmNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGeneralMatchConfirmNotify
}
func (*PlayerGeneralMatchConfirmNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerGeneralMatchConfirmNotify"
}

func (*PlayerGetForceQuitBanInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGetForceQuitBanInfoReq
}
func (*PlayerGetForceQuitBanInfoReq) ProtoMessageType() ProtoMessageType {
	return "PlayerGetForceQuitBanInfoReq"
}

func (*PlayerGetForceQuitBanInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGetForceQuitBanInfoRsp
}
func (*PlayerGetForceQuitBanInfoRsp) ProtoMessageType() ProtoMessageType {
	return "PlayerGetForceQuitBanInfoRsp"
}

func (*ServerTryCancelGeneralMatchNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerTryCancelGeneralMatchNotify
}
func (*ServerTryCancelGeneralMatchNotify) ProtoMessageType() ProtoMessageType {
	return "ServerTryCancelGeneralMatchNotify"
}

func (*PlayerGCGMatchDismissNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGCGMatchDismissNotify
}
func (*PlayerGCGMatchDismissNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerGCGMatchDismissNotify"
}

func (*PlayerGCGMatchConfirmNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerGCGMatchConfirmNotify
}
func (*PlayerGCGMatchConfirmNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerGCGMatchConfirmNotify"
}
