
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetOpActivityInfoReq) },
		func() ProtoMessage { return new(GetOpActivityInfoRsp) },
		func() ProtoMessage { return new(OpActivityDataNotify) },
		func() ProtoMessage { return new(OpActivityUpdateNotify) },
	)
}

const (
	ProtoCommandGetOpActivityInfoReq   ProtoCommand = 5172
	ProtoCommandGetOpActivityInfoRsp   ProtoCommand = 5198
	ProtoCommandOpActivityDataNotify   ProtoCommand = 5112
	ProtoCommandOpActivityUpdateNotify ProtoCommand = 5135
)

func (*GetOpActivityInfoReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetOpActivityInfoReq }
func (*GetOpActivityInfoReq) ProtoMessageType() ProtoMessageType { return "GetOpActivityInfoReq" }

func (*GetOpActivityInfoRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetOpActivityInfoRsp }
func (*GetOpActivityInfoRsp) ProtoMessageType() ProtoMessageType { return "GetOpActivityInfoRsp" }

func (*OpActivityDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOpActivityDataNotify }
func (*OpActivityDataNotify) ProtoMessageType() ProtoMessageType { return "OpActivityDataNotify" }

func (*OpActivityUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandOpActivityUpdateNotify }
func (*OpActivityUpdateNotify) ProtoMessageType() ProtoMessageType { return "OpActivityUpdateNotify" }