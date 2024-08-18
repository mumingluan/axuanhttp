
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetAllH5ActivityInfoReq) },
		func() ProtoMessage { return new(GetAllH5ActivityInfoRsp) },
		func() ProtoMessage { return new(H5ActivityIdsNotify) },
		func() ProtoMessage { return new(SetH5ActivityRedDotTimestampReq) },
		func() ProtoMessage { return new(SetH5ActivityRedDotTimestampRsp) },
	)
}

const (
	ProtoCommandGetAllH5ActivityInfoReq         ProtoCommand = 5668
	ProtoCommandGetAllH5ActivityInfoRsp         ProtoCommand = 5676
	ProtoCommandH5ActivityIdsNotify             ProtoCommand = 5675
	ProtoCommandSetH5ActivityRedDotTimestampReq ProtoCommand = 5657
	ProtoCommandSetH5ActivityRedDotTimestampRsp ProtoCommand = 5652
)

func (*GetAllH5ActivityInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllH5ActivityInfoReq
}
func (*GetAllH5ActivityInfoReq) ProtoMessageType() ProtoMessageType { return "GetAllH5ActivityInfoReq" }

func (*GetAllH5ActivityInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllH5ActivityInfoRsp
}
func (*GetAllH5ActivityInfoRsp) ProtoMessageType() ProtoMessageType { return "GetAllH5ActivityInfoRsp" }

func (*H5ActivityIdsNotify) ProtoCommand() ProtoCommand         { return ProtoCommandH5ActivityIdsNotify }
func (*H5ActivityIdsNotify) ProtoMessageType() ProtoMessageType { return "H5ActivityIdsNotify" }

func (*SetH5ActivityRedDotTimestampReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetH5ActivityRedDotTimestampReq
}
func (*SetH5ActivityRedDotTimestampReq) ProtoMessageType() ProtoMessageType {
	return "SetH5ActivityRedDotTimestampReq"
}

func (*SetH5ActivityRedDotTimestampRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetH5ActivityRedDotTimestampRsp
}
func (*SetH5ActivityRedDotTimestampRsp) ProtoMessageType() ProtoMessageType {
	return "SetH5ActivityRedDotTimestampRsp"
}
