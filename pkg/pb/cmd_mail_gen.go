
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(MailChangeNotify) },
		func() ProtoMessage { return new(ReadMailNotify) },
		func() ProtoMessage { return new(GetMailItemReq) },
		func() ProtoMessage { return new(GetMailItemRsp) },
		func() ProtoMessage { return new(DelMailReq) },
		func() ProtoMessage { return new(DelMailRsp) },
		func() ProtoMessage { return new(GetAuthkeyReq) },
		func() ProtoMessage { return new(GetAuthkeyRsp) },
		func() ProtoMessage { return new(ClientNewMailNotify) },
		func() ProtoMessage { return new(GetAllMailReq) },
		func() ProtoMessage { return new(GetAllMailRsp) },
		func() ProtoMessage { return new(ChangeMailStarNotify) },
		func() ProtoMessage { return new(GetAllMailNotify) },
		func() ProtoMessage { return new(GetAllMailResultNotify) },
	)
}

const (
	ProtoCommandMailChangeNotify       ProtoCommand = 1498
	ProtoCommandReadMailNotify         ProtoCommand = 1412
	ProtoCommandGetMailItemReq         ProtoCommand = 1435
	ProtoCommandGetMailItemRsp         ProtoCommand = 1407
	ProtoCommandDelMailReq             ProtoCommand = 1421
	ProtoCommandDelMailRsp             ProtoCommand = 1403
	ProtoCommandGetAuthkeyReq          ProtoCommand = 1490
	ProtoCommandGetAuthkeyRsp          ProtoCommand = 1473
	ProtoCommandClientNewMailNotify    ProtoCommand = 1499
	ProtoCommandGetAllMailReq          ProtoCommand = 1431
	ProtoCommandGetAllMailRsp          ProtoCommand = 1475
	ProtoCommandChangeMailStarNotify   ProtoCommand = 1448
	ProtoCommandGetAllMailNotify       ProtoCommand = 1497
	ProtoCommandGetAllMailResultNotify ProtoCommand = 1481
)

func (*MailChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMailChangeNotify }
func (*MailChangeNotify) ProtoMessageType() ProtoMessageType { return "MailChangeNotify" }

func (*ReadMailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandReadMailNotify }
func (*ReadMailNotify) ProtoMessageType() ProtoMessageType { return "ReadMailNotify" }

func (*GetMailItemReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetMailItemReq }
func (*GetMailItemReq) ProtoMessageType() ProtoMessageType { return "GetMailItemReq" }

func (*GetMailItemRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetMailItemRsp }
func (*GetMailItemRsp) ProtoMessageType() ProtoMessageType { return "GetMailItemRsp" }

func (*DelMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandDelMailReq }
func (*DelMailReq) ProtoMessageType() ProtoMessageType { return "DelMailReq" }

func (*DelMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDelMailRsp }
func (*DelMailRsp) ProtoMessageType() ProtoMessageType { return "DelMailRsp" }

func (*GetAuthkeyReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetAuthkeyReq }
func (*GetAuthkeyReq) ProtoMessageType() ProtoMessageType { return "GetAuthkeyReq" }

func (*GetAuthkeyRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetAuthkeyRsp }
func (*GetAuthkeyRsp) ProtoMessageType() ProtoMessageType { return "GetAuthkeyRsp" }

func (*ClientNewMailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClientNewMailNotify }
func (*ClientNewMailNotify) ProtoMessageType() ProtoMessageType { return "ClientNewMailNotify" }

func (*GetAllMailReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetAllMailReq }
func (*GetAllMailReq) ProtoMessageType() ProtoMessageType { return "GetAllMailReq" }

func (*GetAllMailRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetAllMailRsp }
func (*GetAllMailRsp) ProtoMessageType() ProtoMessageType { return "GetAllMailRsp" }

func (*ChangeMailStarNotify) ProtoCommand() ProtoCommand         { return ProtoCommandChangeMailStarNotify }
func (*ChangeMailStarNotify) ProtoMessageType() ProtoMessageType { return "ChangeMailStarNotify" }

func (*GetAllMailNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGetAllMailNotify }
func (*GetAllMailNotify) ProtoMessageType() ProtoMessageType { return "GetAllMailNotify" }

func (*GetAllMailResultNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGetAllMailResultNotify }
func (*GetAllMailResultNotify) ProtoMessageType() ProtoMessageType { return "GetAllMailResultNotify" }
