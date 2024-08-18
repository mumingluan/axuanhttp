
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AnchorPointDataNotify) },
		func() ProtoMessage { return new(AnchorPointOpReq) },
		func() ProtoMessage { return new(AnchorPointOpRsp) },
		func() ProtoMessage { return new(SetUpLunchBoxWidgetReq) },
		func() ProtoMessage { return new(SetUpLunchBoxWidgetRsp) },
		func() ProtoMessage { return new(QuickUseWidgetReq) },
		func() ProtoMessage { return new(QuickUseWidgetRsp) },
		func() ProtoMessage { return new(WidgetCoolDownNotify) },
		func() ProtoMessage { return new(WidgetReportReq) },
		func() ProtoMessage { return new(WidgetReportRsp) },
		func() ProtoMessage { return new(ClientCollectorDataNotify) },
		func() ProtoMessage { return new(OneoffGatherPointDetectorDataNotify) },
		func() ProtoMessage { return new(SkyCrystalDetectorDataUpdateNotify) },
		func() ProtoMessage { return new(TreasureMapDetectorDataNotify) },
		func() ProtoMessage { return new(SetWidgetSlotReq) },
		func() ProtoMessage { return new(SetWidgetSlotRsp) },
		func() ProtoMessage { return new(WidgetSlotChangeNotify) },
		func() ProtoMessage { return new(GetWidgetSlotReq) },
		func() ProtoMessage { return new(GetWidgetSlotRsp) },
		func() ProtoMessage { return new(AllWidgetDataNotify) },
		func() ProtoMessage { return new(UseWidgetCreateGadgetReq) },
		func() ProtoMessage { return new(UseWidgetCreateGadgetRsp) },
		func() ProtoMessage { return new(UseWidgetRetractGadgetReq) },
		func() ProtoMessage { return new(UseWidgetRetractGadgetRsp) },
		func() ProtoMessage { return new(WidgetGadgetAllDataNotify) },
		func() ProtoMessage { return new(WidgetGadgetDataNotify) },
		func() ProtoMessage { return new(WidgetGadgetDestroyNotify) },
		func() ProtoMessage { return new(WidgetDoBagReq) },
		func() ProtoMessage { return new(WidgetDoBagRsp) },
		func() ProtoMessage { return new(WidgetActiveChangeNotify) },
		func() ProtoMessage { return new(WidgetUseAttachAbilityGroupChangeNotify) },
		func() ProtoMessage { return new(WidgetCaptureAnimalReq) },
		func() ProtoMessage { return new(WidgetCaptureAnimalRsp) },
		func() ProtoMessage { return new(WidgetUpdateExtraCDReq) },
		func() ProtoMessage { return new(WidgetUpdateExtraCDRsp) },
		func() ProtoMessage { return new(FireworksReformDataNotify) },
		func() ProtoMessage { return new(ReformFireworksReq) },
		func() ProtoMessage { return new(ReformFireworksRsp) },
		func() ProtoMessage { return new(LaunchFireworksReq) },
		func() ProtoMessage { return new(LaunchFireworksRsp) },
		func() ProtoMessage { return new(FireworksLaunchDataNotify) },
		func() ProtoMessage { return new(ChangeWidgetBackgroundActiveStateReq) },
		func() ProtoMessage { return new(ChangeWidgetBackgroundActiveStateRsp) },
		func() ProtoMessage { return new(AllWidgetBackgroundActiveStateNotify) },
		func() ProtoMessage { return new(RemotePlayerWidgetNotify) },
		func() ProtoMessage { return new(WidgetWeatherWizardDataNotify) },
	)
}

const (
	ProtoCommandAnchorPointDataNotify                   ProtoCommand = 4276
	ProtoCommandAnchorPointOpReq                        ProtoCommand = 4257
	ProtoCommandAnchorPointOpRsp                        ProtoCommand = 4252
	ProtoCommandSetUpLunchBoxWidgetReq                  ProtoCommand = 4272
	ProtoCommandSetUpLunchBoxWidgetRsp                  ProtoCommand = 4294
	ProtoCommandQuickUseWidgetReq                       ProtoCommand = 4299
	ProtoCommandQuickUseWidgetRsp                       ProtoCommand = 4270
	ProtoCommandWidgetCoolDownNotify                    ProtoCommand = 4295
	ProtoCommandWidgetReportReq                         ProtoCommand = 4291
	ProtoCommandWidgetReportRsp                         ProtoCommand = 4292
	ProtoCommandClientCollectorDataNotify               ProtoCommand = 4264
	ProtoCommandOneoffGatherPointDetectorDataNotify     ProtoCommand = 4297
	ProtoCommandSkyCrystalDetectorDataUpdateNotify      ProtoCommand = 4287
	ProtoCommandTreasureMapDetectorDataNotify           ProtoCommand = 4300
	ProtoCommandSetWidgetSlotReq                        ProtoCommand = 4259
	ProtoCommandSetWidgetSlotRsp                        ProtoCommand = 4277
	ProtoCommandWidgetSlotChangeNotify                  ProtoCommand = 4267
	ProtoCommandGetWidgetSlotReq                        ProtoCommand = 4253
	ProtoCommandGetWidgetSlotRsp                        ProtoCommand = 4254
	ProtoCommandAllWidgetDataNotify                     ProtoCommand = 4271
	ProtoCommandUseWidgetCreateGadgetReq                ProtoCommand = 4293
	ProtoCommandUseWidgetCreateGadgetRsp                ProtoCommand = 4290
	ProtoCommandUseWidgetRetractGadgetReq               ProtoCommand = 4286
	ProtoCommandUseWidgetRetractGadgetRsp               ProtoCommand = 4261
	ProtoCommandWidgetGadgetAllDataNotify               ProtoCommand = 4284
	ProtoCommandWidgetGadgetDataNotify                  ProtoCommand = 4266
	ProtoCommandWidgetGadgetDestroyNotify               ProtoCommand = 4274
	ProtoCommandWidgetDoBagReq                          ProtoCommand = 4255
	ProtoCommandWidgetDoBagRsp                          ProtoCommand = 4296
	ProtoCommandWidgetActiveChangeNotify                ProtoCommand = 4280
	ProtoCommandWidgetUseAttachAbilityGroupChangeNotify ProtoCommand = 4258
	ProtoCommandWidgetCaptureAnimalReq                  ProtoCommand = 4256
	ProtoCommandWidgetCaptureAnimalRsp                  ProtoCommand = 4289
	ProtoCommandWidgetUpdateExtraCDReq                  ProtoCommand = 5960
	ProtoCommandWidgetUpdateExtraCDRsp                  ProtoCommand = 6056
	ProtoCommandFireworksReformDataNotify               ProtoCommand = 6033
	ProtoCommandReformFireworksReq                      ProtoCommand = 6036
	ProtoCommandReformFireworksRsp                      ProtoCommand = 5929
	ProtoCommandLaunchFireworksReq                      ProtoCommand = 6090
	ProtoCommandLaunchFireworksRsp                      ProtoCommand = 6057
	ProtoCommandFireworksLaunchDataNotify               ProtoCommand = 5928
	ProtoCommandChangeWidgetBackgroundActiveStateReq    ProtoCommand = 5907
	ProtoCommandChangeWidgetBackgroundActiveStateRsp    ProtoCommand = 6060
	ProtoCommandAllWidgetBackgroundActiveStateNotify    ProtoCommand = 6092
	ProtoCommandRemotePlayerWidgetNotify                ProtoCommand = 5995
	ProtoCommandWidgetWeatherWizardDataNotify           ProtoCommand = 5942
)

func (*AnchorPointDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAnchorPointDataNotify }
func (*AnchorPointDataNotify) ProtoMessageType() ProtoMessageType { return "AnchorPointDataNotify" }

func (*AnchorPointOpReq) ProtoCommand() ProtoCommand         { return ProtoCommandAnchorPointOpReq }
func (*AnchorPointOpReq) ProtoMessageType() ProtoMessageType { return "AnchorPointOpReq" }

func (*AnchorPointOpRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAnchorPointOpRsp }
func (*AnchorPointOpRsp) ProtoMessageType() ProtoMessageType { return "AnchorPointOpRsp" }

func (*SetUpLunchBoxWidgetReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetUpLunchBoxWidgetReq }
func (*SetUpLunchBoxWidgetReq) ProtoMessageType() ProtoMessageType { return "SetUpLunchBoxWidgetReq" }

func (*SetUpLunchBoxWidgetRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetUpLunchBoxWidgetRsp }
func (*SetUpLunchBoxWidgetRsp) ProtoMessageType() ProtoMessageType { return "SetUpLunchBoxWidgetRsp" }

func (*QuickUseWidgetReq) ProtoCommand() ProtoCommand         { return ProtoCommandQuickUseWidgetReq }
func (*QuickUseWidgetReq) ProtoMessageType() ProtoMessageType { return "QuickUseWidgetReq" }

func (*QuickUseWidgetRsp) ProtoCommand() ProtoCommand         { return ProtoCommandQuickUseWidgetRsp }
func (*QuickUseWidgetRsp) ProtoMessageType() ProtoMessageType { return "QuickUseWidgetRsp" }

func (*WidgetCoolDownNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetCoolDownNotify }
func (*WidgetCoolDownNotify) ProtoMessageType() ProtoMessageType { return "WidgetCoolDownNotify" }

func (*WidgetReportReq) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetReportReq }
func (*WidgetReportReq) ProtoMessageType() ProtoMessageType { return "WidgetReportReq" }

func (*WidgetReportRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetReportRsp }
func (*WidgetReportRsp) ProtoMessageType() ProtoMessageType { return "WidgetReportRsp" }

func (*ClientCollectorDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientCollectorDataNotify
}
func (*ClientCollectorDataNotify) ProtoMessageType() ProtoMessageType {
	return "ClientCollectorDataNotify"
}

func (*OneoffGatherPointDetectorDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandOneoffGatherPointDetectorDataNotify
}
func (*OneoffGatherPointDetectorDataNotify) ProtoMessageType() ProtoMessageType {
	return "OneoffGatherPointDetectorDataNotify"
}

func (*SkyCrystalDetectorDataUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSkyCrystalDetectorDataUpdateNotify
}
func (*SkyCrystalDetectorDataUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "SkyCrystalDetectorDataUpdateNotify"
}

func (*TreasureMapDetectorDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTreasureMapDetectorDataNotify
}
func (*TreasureMapDetectorDataNotify) ProtoMessageType() ProtoMessageType {
	return "TreasureMapDetectorDataNotify"
}

func (*SetWidgetSlotReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetWidgetSlotReq }
func (*SetWidgetSlotReq) ProtoMessageType() ProtoMessageType { return "SetWidgetSlotReq" }

func (*SetWidgetSlotRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetWidgetSlotRsp }
func (*SetWidgetSlotRsp) ProtoMessageType() ProtoMessageType { return "SetWidgetSlotRsp" }

func (*WidgetSlotChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetSlotChangeNotify }
func (*WidgetSlotChangeNotify) ProtoMessageType() ProtoMessageType { return "WidgetSlotChangeNotify" }

func (*GetWidgetSlotReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetWidgetSlotReq }
func (*GetWidgetSlotReq) ProtoMessageType() ProtoMessageType { return "GetWidgetSlotReq" }

func (*GetWidgetSlotRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetWidgetSlotRsp }
func (*GetWidgetSlotRsp) ProtoMessageType() ProtoMessageType { return "GetWidgetSlotRsp" }

func (*AllWidgetDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAllWidgetDataNotify }
func (*AllWidgetDataNotify) ProtoMessageType() ProtoMessageType { return "AllWidgetDataNotify" }

func (*UseWidgetCreateGadgetReq) ProtoCommand() ProtoCommand {
	return ProtoCommandUseWidgetCreateGadgetReq
}
func (*UseWidgetCreateGadgetReq) ProtoMessageType() ProtoMessageType {
	return "UseWidgetCreateGadgetReq"
}

func (*UseWidgetCreateGadgetRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandUseWidgetCreateGadgetRsp
}
func (*UseWidgetCreateGadgetRsp) ProtoMessageType() ProtoMessageType {
	return "UseWidgetCreateGadgetRsp"
}

func (*UseWidgetRetractGadgetReq) ProtoCommand() ProtoCommand {
	return ProtoCommandUseWidgetRetractGadgetReq
}
func (*UseWidgetRetractGadgetReq) ProtoMessageType() ProtoMessageType {
	return "UseWidgetRetractGadgetReq"
}

func (*UseWidgetRetractGadgetRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandUseWidgetRetractGadgetRsp
}
func (*UseWidgetRetractGadgetRsp) ProtoMessageType() ProtoMessageType {
	return "UseWidgetRetractGadgetRsp"
}

func (*WidgetGadgetAllDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWidgetGadgetAllDataNotify
}
func (*WidgetGadgetAllDataNotify) ProtoMessageType() ProtoMessageType {
	return "WidgetGadgetAllDataNotify"
}

func (*WidgetGadgetDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetGadgetDataNotify }
func (*WidgetGadgetDataNotify) ProtoMessageType() ProtoMessageType { return "WidgetGadgetDataNotify" }

func (*WidgetGadgetDestroyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWidgetGadgetDestroyNotify
}
func (*WidgetGadgetDestroyNotify) ProtoMessageType() ProtoMessageType {
	return "WidgetGadgetDestroyNotify"
}

func (*WidgetDoBagReq) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetDoBagReq }
func (*WidgetDoBagReq) ProtoMessageType() ProtoMessageType { return "WidgetDoBagReq" }

func (*WidgetDoBagRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetDoBagRsp }
func (*WidgetDoBagRsp) ProtoMessageType() ProtoMessageType { return "WidgetDoBagRsp" }

func (*WidgetActiveChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWidgetActiveChangeNotify
}
func (*WidgetActiveChangeNotify) ProtoMessageType() ProtoMessageType {
	return "WidgetActiveChangeNotify"
}

func (*WidgetUseAttachAbilityGroupChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWidgetUseAttachAbilityGroupChangeNotify
}
func (*WidgetUseAttachAbilityGroupChangeNotify) ProtoMessageType() ProtoMessageType {
	return "WidgetUseAttachAbilityGroupChangeNotify"
}

func (*WidgetCaptureAnimalReq) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetCaptureAnimalReq }
func (*WidgetCaptureAnimalReq) ProtoMessageType() ProtoMessageType { return "WidgetCaptureAnimalReq" }

func (*WidgetCaptureAnimalRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetCaptureAnimalRsp }
func (*WidgetCaptureAnimalRsp) ProtoMessageType() ProtoMessageType { return "WidgetCaptureAnimalRsp" }

func (*WidgetUpdateExtraCDReq) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetUpdateExtraCDReq }
func (*WidgetUpdateExtraCDReq) ProtoMessageType() ProtoMessageType { return "WidgetUpdateExtraCDReq" }

func (*WidgetUpdateExtraCDRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWidgetUpdateExtraCDRsp }
func (*WidgetUpdateExtraCDRsp) ProtoMessageType() ProtoMessageType { return "WidgetUpdateExtraCDRsp" }

func (*FireworksReformDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFireworksReformDataNotify
}
func (*FireworksReformDataNotify) ProtoMessageType() ProtoMessageType {
	return "FireworksReformDataNotify"
}

func (*ReformFireworksReq) ProtoCommand() ProtoCommand         { return ProtoCommandReformFireworksReq }
func (*ReformFireworksReq) ProtoMessageType() ProtoMessageType { return "ReformFireworksReq" }

func (*ReformFireworksRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReformFireworksRsp }
func (*ReformFireworksRsp) ProtoMessageType() ProtoMessageType { return "ReformFireworksRsp" }

func (*LaunchFireworksReq) ProtoCommand() ProtoCommand         { return ProtoCommandLaunchFireworksReq }
func (*LaunchFireworksReq) ProtoMessageType() ProtoMessageType { return "LaunchFireworksReq" }

func (*LaunchFireworksRsp) ProtoCommand() ProtoCommand         { return ProtoCommandLaunchFireworksRsp }
func (*LaunchFireworksRsp) ProtoMessageType() ProtoMessageType { return "LaunchFireworksRsp" }

func (*FireworksLaunchDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandFireworksLaunchDataNotify
}
func (*FireworksLaunchDataNotify) ProtoMessageType() ProtoMessageType {
	return "FireworksLaunchDataNotify"
}

func (*ChangeWidgetBackgroundActiveStateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeWidgetBackgroundActiveStateReq
}
func (*ChangeWidgetBackgroundActiveStateReq) ProtoMessageType() ProtoMessageType {
	return "ChangeWidgetBackgroundActiveStateReq"
}

func (*ChangeWidgetBackgroundActiveStateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandChangeWidgetBackgroundActiveStateRsp
}
func (*ChangeWidgetBackgroundActiveStateRsp) ProtoMessageType() ProtoMessageType {
	return "ChangeWidgetBackgroundActiveStateRsp"
}

func (*AllWidgetBackgroundActiveStateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAllWidgetBackgroundActiveStateNotify
}
func (*AllWidgetBackgroundActiveStateNotify) ProtoMessageType() ProtoMessageType {
	return "AllWidgetBackgroundActiveStateNotify"
}

func (*RemotePlayerWidgetNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandRemotePlayerWidgetNotify
}
func (*RemotePlayerWidgetNotify) ProtoMessageType() ProtoMessageType {
	return "RemotePlayerWidgetNotify"
}

func (*WidgetWeatherWizardDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandWidgetWeatherWizardDataNotify
}
func (*WidgetWeatherWizardDataNotify) ProtoMessageType() ProtoMessageType {
	return "WidgetWeatherWizardDataNotify"
}
