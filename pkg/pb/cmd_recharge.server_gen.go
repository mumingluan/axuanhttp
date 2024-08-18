
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(RechargeOrderNotify) },
		func() ProtoMessage { return new(MarkOrderFinishedReq) },
		func() ProtoMessage { return new(MarkOrderFinishedRsp) },
		func() ProtoMessage { return new(GetUnfinishedOrderReq) },
		func() ProtoMessage { return new(GetUnfinishedOrderRsp) },
		func() ProtoMessage { return new(GetSpecificUnfinishedReq) },
		func() ProtoMessage { return new(GetSpecificUnfinishedRsp) },
	)
}

const (
	ProtoCommandRechargeOrderNotify      ProtoCommand = 10726
	ProtoCommandMarkOrderFinishedReq     ProtoCommand = 10718
	ProtoCommandMarkOrderFinishedRsp     ProtoCommand = 10725
	ProtoCommandGetUnfinishedOrderReq    ProtoCommand = 10707
	ProtoCommandGetUnfinishedOrderRsp    ProtoCommand = 10702
	ProtoCommandGetSpecificUnfinishedReq ProtoCommand = 10731
	ProtoCommandGetSpecificUnfinishedRsp ProtoCommand = 10729
)

func (*RechargeOrderNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRechargeOrderNotify }
func (*RechargeOrderNotify) ProtoMessageType() ProtoMessageType { return "RechargeOrderNotify" }

func (*MarkOrderFinishedReq) ProtoCommand() ProtoCommand         { return ProtoCommandMarkOrderFinishedReq }
func (*MarkOrderFinishedReq) ProtoMessageType() ProtoMessageType { return "MarkOrderFinishedReq" }

func (*MarkOrderFinishedRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMarkOrderFinishedRsp }
func (*MarkOrderFinishedRsp) ProtoMessageType() ProtoMessageType { return "MarkOrderFinishedRsp" }

func (*GetUnfinishedOrderReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetUnfinishedOrderReq }
func (*GetUnfinishedOrderReq) ProtoMessageType() ProtoMessageType { return "GetUnfinishedOrderReq" }

func (*GetUnfinishedOrderRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetUnfinishedOrderRsp }
func (*GetUnfinishedOrderRsp) ProtoMessageType() ProtoMessageType { return "GetUnfinishedOrderRsp" }

func (*GetSpecificUnfinishedReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetSpecificUnfinishedReq
}
func (*GetSpecificUnfinishedReq) ProtoMessageType() ProtoMessageType {
	return "GetSpecificUnfinishedReq"
}

func (*GetSpecificUnfinishedRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetSpecificUnfinishedRsp
}
func (*GetSpecificUnfinishedRsp) ProtoMessageType() ProtoMessageType {
	return "GetSpecificUnfinishedRsp"
}
