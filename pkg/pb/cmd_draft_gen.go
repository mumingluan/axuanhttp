
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(DraftOwnerStartInviteReq) },
		func() ProtoMessage { return new(DraftOwnerStartInviteRsp) },
		func() ProtoMessage { return new(DraftOwnerInviteNotify) },
		func() ProtoMessage { return new(DraftGuestReplyInviteReq) },
		func() ProtoMessage { return new(DraftGuestReplyInviteRsp) },
		func() ProtoMessage { return new(DraftGuestReplyInviteNotify) },
		func() ProtoMessage { return new(DraftInviteResultNotify) },
		func() ProtoMessage { return new(DraftOwnerTwiceConfirmNotify) },
		func() ProtoMessage { return new(DraftGuestReplyTwiceConfirmReq) },
		func() ProtoMessage { return new(DraftGuestReplyTwiceConfirmRsp) },
		func() ProtoMessage { return new(DraftTwiceConfirmResultNotify) },
		func() ProtoMessage { return new(DraftGuestReplyTwiceConfirmNotify) },
	)
}

const (
	ProtoCommandDraftOwnerStartInviteReq          ProtoCommand = 5412
	ProtoCommandDraftOwnerStartInviteRsp          ProtoCommand = 5435
	ProtoCommandDraftOwnerInviteNotify            ProtoCommand = 5407
	ProtoCommandDraftGuestReplyInviteReq          ProtoCommand = 5421
	ProtoCommandDraftGuestReplyInviteRsp          ProtoCommand = 5403
	ProtoCommandDraftGuestReplyInviteNotify       ProtoCommand = 5490
	ProtoCommandDraftInviteResultNotify           ProtoCommand = 5473
	ProtoCommandDraftOwnerTwiceConfirmNotify      ProtoCommand = 5499
	ProtoCommandDraftGuestReplyTwiceConfirmReq    ProtoCommand = 5431
	ProtoCommandDraftGuestReplyTwiceConfirmRsp    ProtoCommand = 5475
	ProtoCommandDraftTwiceConfirmResultNotify     ProtoCommand = 5448
	ProtoCommandDraftGuestReplyTwiceConfirmNotify ProtoCommand = 5497
)

func (*DraftOwnerStartInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftOwnerStartInviteReq
}
func (*DraftOwnerStartInviteReq) ProtoMessageType() ProtoMessageType {
	return "DraftOwnerStartInviteReq"
}

func (*DraftOwnerStartInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftOwnerStartInviteRsp
}
func (*DraftOwnerStartInviteRsp) ProtoMessageType() ProtoMessageType {
	return "DraftOwnerStartInviteRsp"
}

func (*DraftOwnerInviteNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDraftOwnerInviteNotify }
func (*DraftOwnerInviteNotify) ProtoMessageType() ProtoMessageType { return "DraftOwnerInviteNotify" }

func (*DraftGuestReplyInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftGuestReplyInviteReq
}
func (*DraftGuestReplyInviteReq) ProtoMessageType() ProtoMessageType {
	return "DraftGuestReplyInviteReq"
}

func (*DraftGuestReplyInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftGuestReplyInviteRsp
}
func (*DraftGuestReplyInviteRsp) ProtoMessageType() ProtoMessageType {
	return "DraftGuestReplyInviteRsp"
}

func (*DraftGuestReplyInviteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftGuestReplyInviteNotify
}
func (*DraftGuestReplyInviteNotify) ProtoMessageType() ProtoMessageType {
	return "DraftGuestReplyInviteNotify"
}

func (*DraftInviteResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftInviteResultNotify
}
func (*DraftInviteResultNotify) ProtoMessageType() ProtoMessageType { return "DraftInviteResultNotify" }

func (*DraftOwnerTwiceConfirmNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftOwnerTwiceConfirmNotify
}
func (*DraftOwnerTwiceConfirmNotify) ProtoMessageType() ProtoMessageType {
	return "DraftOwnerTwiceConfirmNotify"
}

func (*DraftGuestReplyTwiceConfirmReq) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftGuestReplyTwiceConfirmReq
}
func (*DraftGuestReplyTwiceConfirmReq) ProtoMessageType() ProtoMessageType {
	return "DraftGuestReplyTwiceConfirmReq"
}

func (*DraftGuestReplyTwiceConfirmRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftGuestReplyTwiceConfirmRsp
}
func (*DraftGuestReplyTwiceConfirmRsp) ProtoMessageType() ProtoMessageType {
	return "DraftGuestReplyTwiceConfirmRsp"
}

func (*DraftTwiceConfirmResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftTwiceConfirmResultNotify
}
func (*DraftTwiceConfirmResultNotify) ProtoMessageType() ProtoMessageType {
	return "DraftTwiceConfirmResultNotify"
}

func (*DraftGuestReplyTwiceConfirmNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDraftGuestReplyTwiceConfirmNotify
}
func (*DraftGuestReplyTwiceConfirmNotify) ProtoMessageType() ProtoMessageType {
	return "DraftGuestReplyTwiceConfirmNotify"
}