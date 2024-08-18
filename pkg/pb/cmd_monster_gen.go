
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(MonsterSummonTagNotify) },
	)
}

const (
	ProtoCommandMonsterSummonTagNotify ProtoCommand = 1372
)

func (*MonsterSummonTagNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMonsterSummonTagNotify }
func (*MonsterSummonTagNotify) ProtoMessageType() ProtoMessageType { return "MonsterSummonTagNotify" }
