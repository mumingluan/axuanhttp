
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(SendOfflineOpReq) },
		func() ProtoMessage { return new(SendOfflineOpRsp) },
		func() ProtoMessage { return new(GetOfflineOpReq) },
		func() ProtoMessage { return new(GetOfflineOpRsp) },
		func() ProtoMessage { return new(NewOfflineOpNotify) },
		func() ProtoMessage { return new(RemoveOfflineOpReq) },
		func() ProtoMessage { return new(RemoveOfflineOpRsp) },
		func() ProtoMessage { return new(SendGCGOfflineOpReq) },
		func() ProtoMessage { return new(SendGCGOfflineOpRsp) },
	)
}

const (
	ProtoCommandSendOfflineOpReq    ProtoCommand = 12226
	ProtoCommandSendOfflineOpRsp    ProtoCommand = 12218
	ProtoCommandGetOfflineOpReq     ProtoCommand = 12225
	ProtoCommandGetOfflineOpRsp     ProtoCommand = 12207
	ProtoCommandNewOfflineOpNotify  ProtoCommand = 12202
	ProtoCommandRemoveOfflineOpReq  ProtoCommand = 12231
	ProtoCommandRemoveOfflineOpRsp  ProtoCommand = 12229
	ProtoCommandSendGCGOfflineOpReq ProtoCommand = 12222
	ProtoCommandSendGCGOfflineOpRsp ProtoCommand = 12244
)

func (*SendOfflineOpReq) ProtoCommand() ProtoCommand         { return ProtoCommandSendOfflineOpReq }
func (*SendOfflineOpReq) ProtoMessageType() ProtoMessageType { return "SendOfflineOpReq" }

func (*SendOfflineOpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSendOfflineOpRsp }
func (*SendOfflineOpRsp) ProtoMessageType() ProtoMessageType { return "SendOfflineOpRsp" }

func (*GetOfflineOpReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetOfflineOpReq }
func (*GetOfflineOpReq) ProtoMessageType() ProtoMessageType { return "GetOfflineOpReq" }

func (*GetOfflineOpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetOfflineOpRsp }
func (*GetOfflineOpRsp) ProtoMessageType() ProtoMessageType { return "GetOfflineOpRsp" }

func (*NewOfflineOpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandNewOfflineOpNotify }
func (*NewOfflineOpNotify) ProtoMessageType() ProtoMessageType { return "NewOfflineOpNotify" }

func (*RemoveOfflineOpReq) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveOfflineOpReq }
func (*RemoveOfflineOpReq) ProtoMessageType() ProtoMessageType { return "RemoveOfflineOpReq" }

func (*RemoveOfflineOpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveOfflineOpRsp }
func (*RemoveOfflineOpRsp) ProtoMessageType() ProtoMessageType { return "RemoveOfflineOpRsp" }

func (*SendGCGOfflineOpReq) ProtoCommand() ProtoCommand         { return ProtoCommandSendGCGOfflineOpReq }
func (*SendGCGOfflineOpReq) ProtoMessageType() ProtoMessageType { return "SendGCGOfflineOpReq" }

func (*SendGCGOfflineOpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSendGCGOfflineOpRsp }
func (*SendGCGOfflineOpRsp) ProtoMessageType() ProtoMessageType { return "SendGCGOfflineOpRsp" }
