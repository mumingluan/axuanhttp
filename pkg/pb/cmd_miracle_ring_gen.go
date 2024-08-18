
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(UseMiracleRingReq) },
		func() ProtoMessage { return new(UseMiracleRingRsp) },
		func() ProtoMessage { return new(MiracleRingDataNotify) },
		func() ProtoMessage { return new(MiracleRingTakeRewardReq) },
		func() ProtoMessage { return new(MiracleRingTakeRewardRsp) },
		func() ProtoMessage { return new(MiracleRingDropResultNotify) },
		func() ProtoMessage { return new(MiracleRingDeliverItemReq) },
		func() ProtoMessage { return new(MiracleRingDeliverItemRsp) },
		func() ProtoMessage { return new(MiracleRingDestroyNotify) },
	)
}

const (
	ProtoCommandUseMiracleRingReq           ProtoCommand = 5226
	ProtoCommandUseMiracleRingRsp           ProtoCommand = 5218
	ProtoCommandMiracleRingDataNotify       ProtoCommand = 5225
	ProtoCommandMiracleRingTakeRewardReq    ProtoCommand = 5207
	ProtoCommandMiracleRingTakeRewardRsp    ProtoCommand = 5202
	ProtoCommandMiracleRingDropResultNotify ProtoCommand = 5231
	ProtoCommandMiracleRingDeliverItemReq   ProtoCommand = 5229
	ProtoCommandMiracleRingDeliverItemRsp   ProtoCommand = 5222
	ProtoCommandMiracleRingDestroyNotify    ProtoCommand = 5244
)

func (*UseMiracleRingReq) ProtoCommand() ProtoCommand         { return ProtoCommandUseMiracleRingReq }
func (*UseMiracleRingReq) ProtoMessageType() ProtoMessageType { return "UseMiracleRingReq" }

func (*UseMiracleRingRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUseMiracleRingRsp }
func (*UseMiracleRingRsp) ProtoMessageType() ProtoMessageType { return "UseMiracleRingRsp" }

func (*MiracleRingDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMiracleRingDataNotify }
func (*MiracleRingDataNotify) ProtoMessageType() ProtoMessageType { return "MiracleRingDataNotify" }

func (*MiracleRingTakeRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMiracleRingTakeRewardReq
}
func (*MiracleRingTakeRewardReq) ProtoMessageType() ProtoMessageType {
	return "MiracleRingTakeRewardReq"
}

func (*MiracleRingTakeRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMiracleRingTakeRewardRsp
}
func (*MiracleRingTakeRewardRsp) ProtoMessageType() ProtoMessageType {
	return "MiracleRingTakeRewardRsp"
}

func (*MiracleRingDropResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMiracleRingDropResultNotify
}
func (*MiracleRingDropResultNotify) ProtoMessageType() ProtoMessageType {
	return "MiracleRingDropResultNotify"
}

func (*MiracleRingDeliverItemReq) ProtoCommand() ProtoCommand {
	return ProtoCommandMiracleRingDeliverItemReq
}
func (*MiracleRingDeliverItemReq) ProtoMessageType() ProtoMessageType {
	return "MiracleRingDeliverItemReq"
}

func (*MiracleRingDeliverItemRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandMiracleRingDeliverItemRsp
}
func (*MiracleRingDeliverItemRsp) ProtoMessageType() ProtoMessageType {
	return "MiracleRingDeliverItemRsp"
}

func (*MiracleRingDestroyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMiracleRingDestroyNotify
}
func (*MiracleRingDestroyNotify) ProtoMessageType() ProtoMessageType {
	return "MiracleRingDestroyNotify"
}
