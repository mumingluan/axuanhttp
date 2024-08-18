
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PrivateChatReq) },
		func() ProtoMessage { return new(PrivateChatRsp) },
		func() ProtoMessage { return new(PrivateChatNotify) },
		func() ProtoMessage { return new(PullPrivateChatReq) },
		func() ProtoMessage { return new(PullPrivateChatRsp) },
		func() ProtoMessage { return new(PullRecentChatReq) },
		func() ProtoMessage { return new(PullRecentChatRsp) },
		func() ProtoMessage { return new(ReadPrivateChatReq) },
		func() ProtoMessage { return new(ReadPrivateChatRsp) },
		func() ProtoMessage { return new(ChatChannelUpdateNotify) },
		func() ProtoMessage { return new(ChatChannelDataNotify) },
		func() ProtoMessage { return new(ChatChannelShieldNotify) },
		func() ProtoMessage { return new(ChatChannelInfoNotify) },
	)
}

const (
	ProtoCommandPrivateChatReq          ProtoCommand = 5022
	ProtoCommandPrivateChatRsp          ProtoCommand = 5048
	ProtoCommandPrivateChatNotify       ProtoCommand = 4962
	ProtoCommandPullPrivateChatReq      ProtoCommand = 4971
	ProtoCommandPullPrivateChatRsp      ProtoCommand = 4953
	ProtoCommandPullRecentChatReq       ProtoCommand = 5040
	ProtoCommandPullRecentChatRsp       ProtoCommand = 5023
	ProtoCommandReadPrivateChatReq      ProtoCommand = 5049
	ProtoCommandReadPrivateChatRsp      ProtoCommand = 4981
	ProtoCommandChatChannelUpdateNotify ProtoCommand = 5025
	ProtoCommandChatChannelDataNotify   ProtoCommand = 4998
	ProtoCommandChatChannelShieldNotify ProtoCommand = 5047
	ProtoCommandChatChannelInfoNotify   ProtoCommand = 5031
)

func (*PrivateChatReq) ProtoCommand() ProtoCommand         { return ProtoCommandPrivateChatReq }
func (*PrivateChatReq) ProtoMessageType() ProtoMessageType { return "PrivateChatReq" }

func (*PrivateChatRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPrivateChatRsp }
func (*PrivateChatRsp) ProtoMessageType() ProtoMessageType { return "PrivateChatRsp" }

func (*PrivateChatNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPrivateChatNotify }
func (*PrivateChatNotify) ProtoMessageType() ProtoMessageType { return "PrivateChatNotify" }

func (*PullPrivateChatReq) ProtoCommand() ProtoCommand         { return ProtoCommandPullPrivateChatReq }
func (*PullPrivateChatReq) ProtoMessageType() ProtoMessageType { return "PullPrivateChatReq" }

func (*PullPrivateChatRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPullPrivateChatRsp }
func (*PullPrivateChatRsp) ProtoMessageType() ProtoMessageType { return "PullPrivateChatRsp" }

func (*PullRecentChatReq) ProtoCommand() ProtoCommand         { return ProtoCommandPullRecentChatReq }
func (*PullRecentChatReq) ProtoMessageType() ProtoMessageType { return "PullRecentChatReq" }

func (*PullRecentChatRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPullRecentChatRsp }
func (*PullRecentChatRsp) ProtoMessageType() ProtoMessageType { return "PullRecentChatRsp" }

func (*ReadPrivateChatReq) ProtoCommand() ProtoCommand         { return ProtoCommandReadPrivateChatReq }
func (*ReadPrivateChatReq) ProtoMessageType() ProtoMessageType { return "ReadPrivateChatReq" }

func (*ReadPrivateChatRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReadPrivateChatRsp }
func (*ReadPrivateChatRsp) ProtoMessageType() ProtoMessageType { return "ReadPrivateChatRsp" }

func (*ChatChannelUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChatChannelUpdateNotify
}
func (*ChatChannelUpdateNotify) ProtoMessageType() ProtoMessageType { return "ChatChannelUpdateNotify" }

func (*ChatChannelDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChatChannelDataNotify }
func (*ChatChannelDataNotify) ProtoMessageType() ProtoMessageType { return "ChatChannelDataNotify" }

func (*ChatChannelShieldNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandChatChannelShieldNotify
}
func (*ChatChannelShieldNotify) ProtoMessageType() ProtoMessageType { return "ChatChannelShieldNotify" }

func (*ChatChannelInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChatChannelInfoNotify }
func (*ChatChannelInfoNotify) ProtoMessageType() ProtoMessageType { return "ChatChannelInfoNotify" }
