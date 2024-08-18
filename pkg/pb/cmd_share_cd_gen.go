
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AllShareCDDataNotify) },
	)
}

const (
	ProtoCommandAllShareCDDataNotify ProtoCommand = 9072
)

func (*AllShareCDDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAllShareCDDataNotify }
func (*AllShareCDDataNotify) ProtoMessageType() ProtoMessageType { return "AllShareCDDataNotify" }
