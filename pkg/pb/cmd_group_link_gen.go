
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GroupLinkAllNotify) },
		func() ProtoMessage { return new(GroupLinkChangeNotify) },
		func() ProtoMessage { return new(GroupLinkDeleteNotify) },
		func() ProtoMessage { return new(GroupLinkMarkUpdateNotify) },
	)
}

const (
	ProtoCommandGroupLinkAllNotify        ProtoCommand = 5776
	ProtoCommandGroupLinkChangeNotify     ProtoCommand = 5768
	ProtoCommandGroupLinkDeleteNotify     ProtoCommand = 5775
	ProtoCommandGroupLinkMarkUpdateNotify ProtoCommand = 5757
)

func (*GroupLinkAllNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGroupLinkAllNotify }
func (*GroupLinkAllNotify) ProtoMessageType() ProtoMessageType { return "GroupLinkAllNotify" }

func (*GroupLinkChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGroupLinkChangeNotify }
func (*GroupLinkChangeNotify) ProtoMessageType() ProtoMessageType { return "GroupLinkChangeNotify" }

func (*GroupLinkDeleteNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGroupLinkDeleteNotify }
func (*GroupLinkDeleteNotify) ProtoMessageType() ProtoMessageType { return "GroupLinkDeleteNotify" }

func (*GroupLinkMarkUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandGroupLinkMarkUpdateNotify
}
func (*GroupLinkMarkUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "GroupLinkMarkUpdateNotify"
}