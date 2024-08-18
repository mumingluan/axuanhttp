
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AbilityInvocationFixedNotify) },
		func() ProtoMessage { return new(AbilityInvocationsNotify) },
		func() ProtoMessage { return new(ClientAbilityInitBeginNotify) },
		func() ProtoMessage { return new(ClientAbilityInitFinishNotify) },
		func() ProtoMessage { return new(AbilityInvocationFailNotify) },
		func() ProtoMessage { return new(ClientAbilitiesInitFinishCombineNotify) },
		func() ProtoMessage { return new(WindSeedClientNotify) },
		func() ProtoMessage { return new(AbilityChangeNotify) },
		func() ProtoMessage { return new(ClientAbilityChangeNotify) },
		func() ProtoMessage { return new(ServerUpdateGlobalValueNotify) },
		func() ProtoMessage { return new(ServerGlobalValueChangeNotify) },
		func() ProtoMessage { return new(ClientAIStateNotify) },
		func() ProtoMessage { return new(ServerCombatEndNotify) },
		func() ProtoMessage { return new(ClientRemoveCombatEndModifierNotify) },
		func() ProtoMessage { return new(PerformOperationNotify) },
	)
}

const (
	ProtoCommandAbilityInvocationFixedNotify           ProtoCommand = 1172
	ProtoCommandAbilityInvocationsNotify               ProtoCommand = 1198
	ProtoCommandClientAbilityInitBeginNotify           ProtoCommand = 1112
	ProtoCommandClientAbilityInitFinishNotify          ProtoCommand = 1135
	ProtoCommandAbilityInvocationFailNotify            ProtoCommand = 1107
	ProtoCommandClientAbilitiesInitFinishCombineNotify ProtoCommand = 1103
	ProtoCommandWindSeedClientNotify                   ProtoCommand = 1199
	ProtoCommandAbilityChangeNotify                    ProtoCommand = 1131
	ProtoCommandClientAbilityChangeNotify              ProtoCommand = 1175
	ProtoCommandServerUpdateGlobalValueNotify          ProtoCommand = 1148
	ProtoCommandServerGlobalValueChangeNotify          ProtoCommand = 1197
	ProtoCommandClientAIStateNotify                    ProtoCommand = 1181
	ProtoCommandServerCombatEndNotify                  ProtoCommand = 1105
	ProtoCommandClientRemoveCombatEndModifierNotify    ProtoCommand = 1182
	ProtoCommandPerformOperationNotify                 ProtoCommand = 1147
)

func (*AbilityInvocationFixedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAbilityInvocationFixedNotify
}
func (*AbilityInvocationFixedNotify) ProtoMessageType() ProtoMessageType {
	return "AbilityInvocationFixedNotify"
}

func (*AbilityInvocationsNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAbilityInvocationsNotify
}
func (*AbilityInvocationsNotify) ProtoMessageType() ProtoMessageType {
	return "AbilityInvocationsNotify"
}

func (*ClientAbilityInitBeginNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientAbilityInitBeginNotify
}
func (*ClientAbilityInitBeginNotify) ProtoMessageType() ProtoMessageType {
	return "ClientAbilityInitBeginNotify"
}

func (*ClientAbilityInitFinishNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientAbilityInitFinishNotify
}
func (*ClientAbilityInitFinishNotify) ProtoMessageType() ProtoMessageType {
	return "ClientAbilityInitFinishNotify"
}

func (*AbilityInvocationFailNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAbilityInvocationFailNotify
}
func (*AbilityInvocationFailNotify) ProtoMessageType() ProtoMessageType {
	return "AbilityInvocationFailNotify"
}

func (*ClientAbilitiesInitFinishCombineNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientAbilitiesInitFinishCombineNotify
}
func (*ClientAbilitiesInitFinishCombineNotify) ProtoMessageType() ProtoMessageType {
	return "ClientAbilitiesInitFinishCombineNotify"
}

func (*WindSeedClientNotify) ProtoCommand() ProtoCommand         { return ProtoCommandWindSeedClientNotify }
func (*WindSeedClientNotify) ProtoMessageType() ProtoMessageType { return "WindSeedClientNotify" }

func (*AbilityChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAbilityChangeNotify }
func (*AbilityChangeNotify) ProtoMessageType() ProtoMessageType { return "AbilityChangeNotify" }

func (*ClientAbilityChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientAbilityChangeNotify
}
func (*ClientAbilityChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ClientAbilityChangeNotify"
}

func (*ServerUpdateGlobalValueNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerUpdateGlobalValueNotify
}
func (*ServerUpdateGlobalValueNotify) ProtoMessageType() ProtoMessageType {
	return "ServerUpdateGlobalValueNotify"
}

func (*ServerGlobalValueChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGlobalValueChangeNotify
}
func (*ServerGlobalValueChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ServerGlobalValueChangeNotify"
}

func (*ClientAIStateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClientAIStateNotify }
func (*ClientAIStateNotify) ProtoMessageType() ProtoMessageType { return "ClientAIStateNotify" }

func (*ServerCombatEndNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerCombatEndNotify }
func (*ServerCombatEndNotify) ProtoMessageType() ProtoMessageType { return "ServerCombatEndNotify" }

func (*ClientRemoveCombatEndModifierNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandClientRemoveCombatEndModifierNotify
}
func (*ClientRemoveCombatEndModifierNotify) ProtoMessageType() ProtoMessageType {
	return "ClientRemoveCombatEndModifierNotify"
}

func (*PerformOperationNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPerformOperationNotify }
func (*PerformOperationNotify) ProtoMessageType() ProtoMessageType { return "PerformOperationNotify" }
