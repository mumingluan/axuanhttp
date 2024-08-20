
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetBlossomBriefInfoListReq) },
		func() ProtoMessage { return new(GetBlossomBriefInfoListRsp) },
		func() ProtoMessage { return new(BlossomBriefInfoNotify) },
		func() ProtoMessage { return new(WorldOwnerBlossomBriefInfoNotify) },
		func() ProtoMessage { return new(WorldOwnerBlossomScheduleInfoNotify) },
		func() ProtoMessage { return new(BlossomChestCreateNotify) },
		func() ProtoMessage { return new(OpenBlossomCircleCampGuideNotify) },
	)
}

const (
	ProtoCommandGetBlossomBriefInfoListReq          ProtoCommand = 2772
	ProtoCommandGetBlossomBriefInfoListRsp          ProtoCommand = 2798
	ProtoCommandBlossomBriefInfoNotify              ProtoCommand = 2712
	ProtoCommandWorldOwnerBlossomBriefInfoNotify    ProtoCommand = 2735
	ProtoCommandWorldOwnerBlossomScheduleInfoNotify ProtoCommand = 2707
	ProtoCommandBlossomChestCreateNotify            ProtoCommand = 2721
	ProtoCommandOpenBlossomCircleCampGuideNotify    ProtoCommand = 2703
)

func (*GetBlossomBriefInfoListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetBlossomBriefInfoListReq
}
func (*GetBlossomBriefInfoListReq) ProtoMessageType() ProtoMessageType {
	return "GetBlossomBriefInfoListReq"
}

func (*GetBlossomBriefInfoListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetBlossomBriefInfoListRsp
}
func (*GetBlossomBriefInfoListRsp) ProtoMessageType() ProtoMessageType {
	return "GetBlossomBriefInfoListRsp"
}

func (*BlossomBriefInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandBlossomBriefInfoNotify }
func (*BlossomBriefInfoNotify) ProtoMessageType() ProtoMessageType { return "BlossomBriefInfoNotify" }

func (*WorldOwnerBlossomBriefInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldOwnerBlossomBriefInfoNotify
}
func (*WorldOwnerBlossomBriefInfoNotify) ProtoMessageType() ProtoMessageType {
	return "WorldOwnerBlossomBriefInfoNotify"
}

func (*WorldOwnerBlossomScheduleInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWorldOwnerBlossomScheduleInfoNotify
}
func (*WorldOwnerBlossomScheduleInfoNotify) ProtoMessageType() ProtoMessageType {
	return "WorldOwnerBlossomScheduleInfoNotify"
}

func (*BlossomChestCreateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBlossomChestCreateNotify
}
func (*BlossomChestCreateNotify) ProtoMessageType() ProtoMessageType {
	return "BlossomChestCreateNotify"
}

func (*OpenBlossomCircleCampGuideNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandOpenBlossomCircleCampGuideNotify
}
func (*OpenBlossomCircleCampGuideNotify) ProtoMessageType() ProtoMessageType {
	return "OpenBlossomCircleCampGuideNotify"
}