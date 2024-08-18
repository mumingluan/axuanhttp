
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AddMatchUnitReq) },
		func() ProtoMessage { return new(AddMatchUnitRsp) },
		func() ProtoMessage { return new(MatchStartNotify) },
		func() ProtoMessage { return new(RemoveMatchUnitReq) },
		func() ProtoMessage { return new(RemoveMatchUnitRsp) },
		func() ProtoMessage { return new(MatchUnitRemovedNotify) },
		func() ProtoMessage { return new(GuestUnitMatchSuccNotify) },
		func() ProtoMessage { return new(ConfirmGuestUnitReq) },
		func() ProtoMessage { return new(ConfirmGuestUnitRsp) },
		func() ProtoMessage { return new(MatchPlayerJoinNotify) },
		func() ProtoMessage { return new(MatchPlayerQuitNotify) },
		func() ProtoMessage { return new(GuestUnitAgreedResultNotify) },
		func() ProtoMessage { return new(HostUnitMatchSuccNotify) },
		func() ProtoMessage { return new(SyncMatchUnitReq) },
		func() ProtoMessage { return new(SyncMatchUnitRsp) },
		func() ProtoMessage { return new(MatchserverWorkloadInfoNotify) },
		func() ProtoMessage { return new(MatchPlayerUpdateNotify) },
		func() ProtoMessage { return new(MatchIdentityNotify) },
		func() ProtoMessage { return new(GeneralMatchFinishNotify) },
	)
}

const (
	ProtoCommandAddMatchUnitReq               ProtoCommand = 10926
	ProtoCommandAddMatchUnitRsp               ProtoCommand = 10918
	ProtoCommandMatchStartNotify              ProtoCommand = 10925
	ProtoCommandRemoveMatchUnitReq            ProtoCommand = 10907
	ProtoCommandRemoveMatchUnitRsp            ProtoCommand = 10902
	ProtoCommandMatchUnitRemovedNotify        ProtoCommand = 10931
	ProtoCommandGuestUnitMatchSuccNotify      ProtoCommand = 10929
	ProtoCommandConfirmGuestUnitReq           ProtoCommand = 10922
	ProtoCommandConfirmGuestUnitRsp           ProtoCommand = 10944
	ProtoCommandMatchPlayerJoinNotify         ProtoCommand = 10949
	ProtoCommandMatchPlayerQuitNotify         ProtoCommand = 10920
	ProtoCommandGuestUnitAgreedResultNotify   ProtoCommand = 10945
	ProtoCommandHostUnitMatchSuccNotify       ProtoCommand = 10941
	ProtoCommandSyncMatchUnitReq              ProtoCommand = 10942
	ProtoCommandSyncMatchUnitRsp              ProtoCommand = 10914
	ProtoCommandMatchserverWorkloadInfoNotify ProtoCommand = 10947
	ProtoCommandMatchPlayerUpdateNotify       ProtoCommand = 10937
	ProtoCommandMatchIdentityNotify           ProtoCommand = 10923
	ProtoCommandGeneralMatchFinishNotify      ProtoCommand = 10935
)

func (*AddMatchUnitReq) ProtoCommand() ProtoCommand         { return ProtoCommandAddMatchUnitReq }
func (*AddMatchUnitReq) ProtoMessageType() ProtoMessageType { return "AddMatchUnitReq" }

func (*AddMatchUnitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAddMatchUnitRsp }
func (*AddMatchUnitRsp) ProtoMessageType() ProtoMessageType { return "AddMatchUnitRsp" }

func (*MatchStartNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMatchStartNotify }
func (*MatchStartNotify) ProtoMessageType() ProtoMessageType { return "MatchStartNotify" }

func (*RemoveMatchUnitReq) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveMatchUnitReq }
func (*RemoveMatchUnitReq) ProtoMessageType() ProtoMessageType { return "RemoveMatchUnitReq" }

func (*RemoveMatchUnitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveMatchUnitRsp }
func (*RemoveMatchUnitRsp) ProtoMessageType() ProtoMessageType { return "RemoveMatchUnitRsp" }

func (*MatchUnitRemovedNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMatchUnitRemovedNotify }
func (*MatchUnitRemovedNotify) ProtoMessageType() ProtoMessageType { return "MatchUnitRemovedNotify" }

func (*GuestUnitMatchSuccNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGuestUnitMatchSuccNotify
}
func (*GuestUnitMatchSuccNotify) ProtoMessageType() ProtoMessageType {
	return "GuestUnitMatchSuccNotify"
}

func (*ConfirmGuestUnitReq) ProtoCommand() ProtoCommand         { return ProtoCommandConfirmGuestUnitReq }
func (*ConfirmGuestUnitReq) ProtoMessageType() ProtoMessageType { return "ConfirmGuestUnitReq" }

func (*ConfirmGuestUnitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandConfirmGuestUnitRsp }
func (*ConfirmGuestUnitRsp) ProtoMessageType() ProtoMessageType { return "ConfirmGuestUnitRsp" }

func (*MatchPlayerJoinNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMatchPlayerJoinNotify }
func (*MatchPlayerJoinNotify) ProtoMessageType() ProtoMessageType { return "MatchPlayerJoinNotify" }

func (*MatchPlayerQuitNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMatchPlayerQuitNotify }
func (*MatchPlayerQuitNotify) ProtoMessageType() ProtoMessageType { return "MatchPlayerQuitNotify" }

func (*GuestUnitAgreedResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGuestUnitAgreedResultNotify
}
func (*GuestUnitAgreedResultNotify) ProtoMessageType() ProtoMessageType {
	return "GuestUnitAgreedResultNotify"
}

func (*HostUnitMatchSuccNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandHostUnitMatchSuccNotify
}
func (*HostUnitMatchSuccNotify) ProtoMessageType() ProtoMessageType { return "HostUnitMatchSuccNotify" }

func (*SyncMatchUnitReq) ProtoCommand() ProtoCommand         { return ProtoCommandSyncMatchUnitReq }
func (*SyncMatchUnitReq) ProtoMessageType() ProtoMessageType { return "SyncMatchUnitReq" }

func (*SyncMatchUnitRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSyncMatchUnitRsp }
func (*SyncMatchUnitRsp) ProtoMessageType() ProtoMessageType { return "SyncMatchUnitRsp" }

func (*MatchserverWorkloadInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMatchserverWorkloadInfoNotify
}
func (*MatchserverWorkloadInfoNotify) ProtoMessageType() ProtoMessageType {
	return "MatchserverWorkloadInfoNotify"
}

func (*MatchPlayerUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMatchPlayerUpdateNotify
}
func (*MatchPlayerUpdateNotify) ProtoMessageType() ProtoMessageType { return "MatchPlayerUpdateNotify" }

func (*MatchIdentityNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMatchIdentityNotify }
func (*MatchIdentityNotify) ProtoMessageType() ProtoMessageType { return "MatchIdentityNotify" }

func (*GeneralMatchFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGeneralMatchFinishNotify
}
func (*GeneralMatchFinishNotify) ProtoMessageType() ProtoMessageType {
	return "GeneralMatchFinishNotify"
}
