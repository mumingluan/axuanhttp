
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AddRedisMailReq) },
		func() ProtoMessage { return new(AddRedisMailRsp) },
		func() ProtoMessage { return new(DelRedisMailReq) },
		func() ProtoMessage { return new(DelRedisMailRsp) },
		func() ProtoMessage { return new(GetRedisMailReq) },
		func() ProtoMessage { return new(GetRedisMailRsp) },
		func() ProtoMessage { return new(RedisMailChangeNotify) },
		func() ProtoMessage { return new(SendOfflineMsgReq) },
		func() ProtoMessage { return new(SendOfflineMsgRsp) },
		func() ProtoMessage { return new(NewOfflineMsgNotify) },
		func() ProtoMessage { return new(GetOfflineMsgReq) },
		func() ProtoMessage { return new(GetOfflineMsgRsp) },
		func() ProtoMessage { return new(RemoveOfflineMsgNotify) },
		func() ProtoMessage { return new(ClearOfflineMsgNotify) },
		func() ProtoMessage { return new(UpdateRedisMailReq) },
		func() ProtoMessage { return new(UpdateRedisMailRsp) },
		func() ProtoMessage { return new(GetRedisMailTransactionStatusReq) },
		func() ProtoMessage { return new(GetRedisMailTransactionStatusRsp) },
		func() ProtoMessage { return new(ClearUpRedisMailNotify) },
		func() ProtoMessage { return new(CleanRedisMailNotify) },
		func() ProtoMessage { return new(ResendRedisMailReq) },
		func() ProtoMessage { return new(ResendRedisMailRsp) },
	)
}

const (
	ProtoCommandAddRedisMailReq                  ProtoCommand = 10372
	ProtoCommandAddRedisMailRsp                  ProtoCommand = 10398
	ProtoCommandDelRedisMailReq                  ProtoCommand = 10312
	ProtoCommandDelRedisMailRsp                  ProtoCommand = 10335
	ProtoCommandGetRedisMailReq                  ProtoCommand = 10307
	ProtoCommandGetRedisMailRsp                  ProtoCommand = 10321
	ProtoCommandRedisMailChangeNotify            ProtoCommand = 10303
	ProtoCommandSendOfflineMsgReq                ProtoCommand = 10390
	ProtoCommandSendOfflineMsgRsp                ProtoCommand = 10373
	ProtoCommandNewOfflineMsgNotify              ProtoCommand = 10399
	ProtoCommandGetOfflineMsgReq                 ProtoCommand = 10331
	ProtoCommandGetOfflineMsgRsp                 ProtoCommand = 10375
	ProtoCommandRemoveOfflineMsgNotify           ProtoCommand = 10348
	ProtoCommandClearOfflineMsgNotify            ProtoCommand = 10397
	ProtoCommandUpdateRedisMailReq               ProtoCommand = 10381
	ProtoCommandUpdateRedisMailRsp               ProtoCommand = 10305
	ProtoCommandGetRedisMailTransactionStatusReq ProtoCommand = 10382
	ProtoCommandGetRedisMailTransactionStatusRsp ProtoCommand = 10347
	ProtoCommandClearUpRedisMailNotify           ProtoCommand = 10339
	ProtoCommandCleanRedisMailNotify             ProtoCommand = 10353
	ProtoCommandResendRedisMailReq               ProtoCommand = 10322
	ProtoCommandResendRedisMailRsp               ProtoCommand = 10365
)

func (*AddRedisMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandAddRedisMailReq }
func (*AddRedisMailReq) ProtoMessageType() ProtoMessageType { return "AddRedisMailReq" }

func (*AddRedisMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAddRedisMailRsp }
func (*AddRedisMailRsp) ProtoMessageType() ProtoMessageType { return "AddRedisMailRsp" }

func (*DelRedisMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandDelRedisMailReq }
func (*DelRedisMailReq) ProtoMessageType() ProtoMessageType { return "DelRedisMailReq" }

func (*DelRedisMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDelRedisMailRsp }
func (*DelRedisMailRsp) ProtoMessageType() ProtoMessageType { return "DelRedisMailRsp" }

func (*GetRedisMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetRedisMailReq }
func (*GetRedisMailReq) ProtoMessageType() ProtoMessageType { return "GetRedisMailReq" }

func (*GetRedisMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetRedisMailRsp }
func (*GetRedisMailRsp) ProtoMessageType() ProtoMessageType { return "GetRedisMailRsp" }

func (*RedisMailChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRedisMailChangeNotify }
func (*RedisMailChangeNotify) ProtoMessageType() ProtoMessageType { return "RedisMailChangeNotify" }

func (*SendOfflineMsgReq) ProtoCommand() ProtoCommand         { return ProtoCommandSendOfflineMsgReq }
func (*SendOfflineMsgReq) ProtoMessageType() ProtoMessageType { return "SendOfflineMsgReq" }

func (*SendOfflineMsgRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSendOfflineMsgRsp }
func (*SendOfflineMsgRsp) ProtoMessageType() ProtoMessageType { return "SendOfflineMsgRsp" }

func (*NewOfflineMsgNotify) ProtoCommand() ProtoCommand         { return ProtoCommandNewOfflineMsgNotify }
func (*NewOfflineMsgNotify) ProtoMessageType() ProtoMessageType { return "NewOfflineMsgNotify" }

func (*GetOfflineMsgReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetOfflineMsgReq }
func (*GetOfflineMsgReq) ProtoMessageType() ProtoMessageType { return "GetOfflineMsgReq" }

func (*GetOfflineMsgRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetOfflineMsgRsp }
func (*GetOfflineMsgRsp) ProtoMessageType() ProtoMessageType { return "GetOfflineMsgRsp" }

func (*RemoveOfflineMsgNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveOfflineMsgNotify }
func (*RemoveOfflineMsgNotify) ProtoMessageType() ProtoMessageType { return "RemoveOfflineMsgNotify" }

func (*ClearOfflineMsgNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClearOfflineMsgNotify }
func (*ClearOfflineMsgNotify) ProtoMessageType() ProtoMessageType { return "ClearOfflineMsgNotify" }

func (*UpdateRedisMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandUpdateRedisMailReq }
func (*UpdateRedisMailReq) ProtoMessageType() ProtoMessageType { return "UpdateRedisMailReq" }

func (*UpdateRedisMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUpdateRedisMailRsp }
func (*UpdateRedisMailRsp) ProtoMessageType() ProtoMessageType { return "UpdateRedisMailRsp" }

func (*GetRedisMailTransactionStatusReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRedisMailTransactionStatusReq
}
func (*GetRedisMailTransactionStatusReq) ProtoMessageType() ProtoMessageType {
	return "GetRedisMailTransactionStatusReq"
}

func (*GetRedisMailTransactionStatusRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRedisMailTransactionStatusRsp
}
func (*GetRedisMailTransactionStatusRsp) ProtoMessageType() ProtoMessageType {
	return "GetRedisMailTransactionStatusRsp"
}

func (*ClearUpRedisMailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClearUpRedisMailNotify }
func (*ClearUpRedisMailNotify) ProtoMessageType() ProtoMessageType { return "ClearUpRedisMailNotify" }

func (*CleanRedisMailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCleanRedisMailNotify }
func (*CleanRedisMailNotify) ProtoMessageType() ProtoMessageType { return "CleanRedisMailNotify" }

func (*ResendRedisMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandResendRedisMailReq }
func (*ResendRedisMailReq) ProtoMessageType() ProtoMessageType { return "ResendRedisMailReq" }

func (*ResendRedisMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandResendRedisMailRsp }
func (*ResendRedisMailRsp) ProtoMessageType() ProtoMessageType { return "ResendRedisMailRsp" }
