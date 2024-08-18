
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetPlayerFriendListReq) },
		func() ProtoMessage { return new(GetPlayerFriendListRsp) },
		func() ProtoMessage { return new(AskAddFriendReq) },
		func() ProtoMessage { return new(AskAddFriendRsp) },
		func() ProtoMessage { return new(DealAddFriendReq) },
		func() ProtoMessage { return new(DealAddFriendRsp) },
		func() ProtoMessage { return new(GetPlayerSocialDetailReq) },
		func() ProtoMessage { return new(GetPlayerSocialDetailRsp) },
		func() ProtoMessage { return new(DeleteFriendReq) },
		func() ProtoMessage { return new(DeleteFriendRsp) },
		func() ProtoMessage { return new(SetPlayerBirthdayReq) },
		func() ProtoMessage { return new(SetPlayerBirthdayRsp) },
		func() ProtoMessage { return new(SetPlayerSignatureReq) },
		func() ProtoMessage { return new(SetPlayerSignatureRsp) },
		func() ProtoMessage { return new(SetPlayerHeadImageReq) },
		func() ProtoMessage { return new(SetPlayerHeadImageRsp) },
		func() ProtoMessage { return new(UpdatePS4FriendListNotify) },
		func() ProtoMessage { return new(DeleteFriendNotify) },
		func() ProtoMessage { return new(AddFriendNotify) },
		func() ProtoMessage { return new(AskAddFriendNotify) },
		func() ProtoMessage { return new(SetNameCardReq) },
		func() ProtoMessage { return new(SetNameCardRsp) },
		func() ProtoMessage { return new(GetAllUnlockNameCardReq) },
		func() ProtoMessage { return new(GetAllUnlockNameCardRsp) },
		func() ProtoMessage { return new(AddBlacklistReq) },
		func() ProtoMessage { return new(AddBlacklistRsp) },
		func() ProtoMessage { return new(RemoveBlacklistReq) },
		func() ProtoMessage { return new(RemoveBlacklistRsp) },
		func() ProtoMessage { return new(UnlockNameCardNotify) },
		func() ProtoMessage { return new(GetRecentMpPlayerListReq) },
		func() ProtoMessage { return new(GetRecentMpPlayerListRsp) },
		func() ProtoMessage { return new(SocialDataNotify) },
		func() ProtoMessage { return new(TakeFirstShareRewardReq) },
		func() ProtoMessage { return new(TakeFirstShareRewardRsp) },
		func() ProtoMessage { return new(UpdatePS4BlockListReq) },
		func() ProtoMessage { return new(UpdatePS4BlockListRsp) },
		func() ProtoMessage { return new(GetPlayerBlacklistReq) },
		func() ProtoMessage { return new(GetPlayerBlacklistRsp) },
		func() ProtoMessage { return new(PlayerReportReq) },
		func() ProtoMessage { return new(PlayerReportRsp) },
		func() ProtoMessage { return new(SetFriendRemarkNameReq) },
		func() ProtoMessage { return new(SetFriendRemarkNameRsp) },
		func() ProtoMessage { return new(UpdatePlayerShowAvatarListReq) },
		func() ProtoMessage { return new(UpdatePlayerShowAvatarListRsp) },
		func() ProtoMessage { return new(GetFriendShowAvatarInfoReq) },
		func() ProtoMessage { return new(GetFriendShowAvatarInfoRsp) },
		func() ProtoMessage { return new(UpdatePlayerShowNameCardListReq) },
		func() ProtoMessage { return new(UpdatePlayerShowNameCardListRsp) },
		func() ProtoMessage { return new(GetFriendShowNameCardInfoReq) },
		func() ProtoMessage { return new(GetFriendShowNameCardInfoRsp) },
		func() ProtoMessage { return new(ForceAddPlayerFriendReq) },
		func() ProtoMessage { return new(ForceAddPlayerFriendRsp) },
		func() ProtoMessage { return new(ProfilePictureChangeNotify) },
		func() ProtoMessage { return new(PSNFriendListNotify) },
		func() ProtoMessage { return new(PSNBlackListNotify) },
		func() ProtoMessage { return new(GetPlayerAskFriendListReq) },
		func() ProtoMessage { return new(GetPlayerAskFriendListRsp) },
		func() ProtoMessage { return new(GetChatEmojiCollectionReq) },
		func() ProtoMessage { return new(GetChatEmojiCollectionRsp) },
		func() ProtoMessage { return new(SetChatEmojiCollectionReq) },
		func() ProtoMessage { return new(SetChatEmojiCollectionRsp) },
		func() ProtoMessage { return new(UpdatePS4FriendListReq) },
		func() ProtoMessage { return new(UpdatePS4FriendListRsp) },
		func() ProtoMessage { return new(FriendInfoChangeNotify) },
		func() ProtoMessage { return new(PlayerSignatureAuditDataNotify) },
		func() ProtoMessage { return new(PlayerSignatureNotify) },
		func() ProtoMessage { return new(SignatureAuditConfigNotify) },
		func() ProtoMessage { return new(ReadSignatureAuditReq) },
		func() ProtoMessage { return new(ReadSignatureAuditRsp) },
	)
}

const (
	ProtoCommandGetPlayerFriendListReq          ProtoCommand = 4072
	ProtoCommandGetPlayerFriendListRsp          ProtoCommand = 4098
	ProtoCommandAskAddFriendReq                 ProtoCommand = 4007
	ProtoCommandAskAddFriendRsp                 ProtoCommand = 4021
	ProtoCommandDealAddFriendReq                ProtoCommand = 4003
	ProtoCommandDealAddFriendRsp                ProtoCommand = 4090
	ProtoCommandGetPlayerSocialDetailReq        ProtoCommand = 4073
	ProtoCommandGetPlayerSocialDetailRsp        ProtoCommand = 4099
	ProtoCommandDeleteFriendReq                 ProtoCommand = 4031
	ProtoCommandDeleteFriendRsp                 ProtoCommand = 4075
	ProtoCommandSetPlayerBirthdayReq            ProtoCommand = 4048
	ProtoCommandSetPlayerBirthdayRsp            ProtoCommand = 4097
	ProtoCommandSetPlayerSignatureReq           ProtoCommand = 4081
	ProtoCommandSetPlayerSignatureRsp           ProtoCommand = 4005
	ProtoCommandSetPlayerHeadImageReq           ProtoCommand = 4082
	ProtoCommandSetPlayerHeadImageRsp           ProtoCommand = 4047
	ProtoCommandUpdatePS4FriendListNotify       ProtoCommand = 4039
	ProtoCommandDeleteFriendNotify              ProtoCommand = 4053
	ProtoCommandAddFriendNotify                 ProtoCommand = 4022
	ProtoCommandAskAddFriendNotify              ProtoCommand = 4065
	ProtoCommandSetNameCardReq                  ProtoCommand = 4004
	ProtoCommandSetNameCardRsp                  ProtoCommand = 4093
	ProtoCommandGetAllUnlockNameCardReq         ProtoCommand = 4027
	ProtoCommandGetAllUnlockNameCardRsp         ProtoCommand = 4094
	ProtoCommandAddBlacklistReq                 ProtoCommand = 4088
	ProtoCommandAddBlacklistRsp                 ProtoCommand = 4026
	ProtoCommandRemoveBlacklistReq              ProtoCommand = 4063
	ProtoCommandRemoveBlacklistRsp              ProtoCommand = 4095
	ProtoCommandUnlockNameCardNotify            ProtoCommand = 4006
	ProtoCommandGetRecentMpPlayerListReq        ProtoCommand = 4034
	ProtoCommandGetRecentMpPlayerListRsp        ProtoCommand = 4050
	ProtoCommandSocialDataNotify                ProtoCommand = 4043
	ProtoCommandTakeFirstShareRewardReq         ProtoCommand = 4074
	ProtoCommandTakeFirstShareRewardRsp         ProtoCommand = 4076
	ProtoCommandUpdatePS4BlockListReq           ProtoCommand = 4046
	ProtoCommandUpdatePS4BlockListRsp           ProtoCommand = 4041
	ProtoCommandGetPlayerBlacklistReq           ProtoCommand = 4049
	ProtoCommandGetPlayerBlacklistRsp           ProtoCommand = 4091
	ProtoCommandPlayerReportReq                 ProtoCommand = 4024
	ProtoCommandPlayerReportRsp                 ProtoCommand = 4056
	ProtoCommandSetFriendRemarkNameReq          ProtoCommand = 4042
	ProtoCommandSetFriendRemarkNameRsp          ProtoCommand = 4030
	ProtoCommandUpdatePlayerShowAvatarListReq   ProtoCommand = 4067
	ProtoCommandUpdatePlayerShowAvatarListRsp   ProtoCommand = 4058
	ProtoCommandGetFriendShowAvatarInfoReq      ProtoCommand = 4070
	ProtoCommandGetFriendShowAvatarInfoRsp      ProtoCommand = 4017
	ProtoCommandUpdatePlayerShowNameCardListReq ProtoCommand = 4002
	ProtoCommandUpdatePlayerShowNameCardListRsp ProtoCommand = 4019
	ProtoCommandGetFriendShowNameCardInfoReq    ProtoCommand = 4061
	ProtoCommandGetFriendShowNameCardInfoRsp    ProtoCommand = 4029
	ProtoCommandForceAddPlayerFriendReq         ProtoCommand = 4057
	ProtoCommandForceAddPlayerFriendRsp         ProtoCommand = 4100
	ProtoCommandProfilePictureChangeNotify      ProtoCommand = 4016
	ProtoCommandPSNFriendListNotify             ProtoCommand = 4087
	ProtoCommandPSNBlackListNotify              ProtoCommand = 4040
	ProtoCommandGetPlayerAskFriendListReq       ProtoCommand = 4018
	ProtoCommandGetPlayerAskFriendListRsp       ProtoCommand = 4066
	ProtoCommandGetChatEmojiCollectionReq       ProtoCommand = 4068
	ProtoCommandGetChatEmojiCollectionRsp       ProtoCommand = 4033
	ProtoCommandSetChatEmojiCollectionReq       ProtoCommand = 4084
	ProtoCommandSetChatEmojiCollectionRsp       ProtoCommand = 4080
	ProtoCommandUpdatePS4FriendListReq          ProtoCommand = 4089
	ProtoCommandUpdatePS4FriendListRsp          ProtoCommand = 4059
	ProtoCommandFriendInfoChangeNotify          ProtoCommand = 4032
	ProtoCommandPlayerSignatureAuditDataNotify  ProtoCommand = 4060
	ProtoCommandPlayerSignatureNotify           ProtoCommand = 4014
	ProtoCommandSignatureAuditConfigNotify      ProtoCommand = 4092
	ProtoCommandReadSignatureAuditReq           ProtoCommand = 4020
	ProtoCommandReadSignatureAuditRsp           ProtoCommand = 4064
)

func (*GetPlayerFriendListReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetPlayerFriendListReq }
func (*GetPlayerFriendListReq) ProtoMessageType() ProtoMessageType { return "GetPlayerFriendListReq" }

func (*GetPlayerFriendListRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetPlayerFriendListRsp }
func (*GetPlayerFriendListRsp) ProtoMessageType() ProtoMessageType { return "GetPlayerFriendListRsp" }

func (*AskAddFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandAskAddFriendReq }
func (*AskAddFriendReq) ProtoMessageType() ProtoMessageType { return "AskAddFriendReq" }

func (*AskAddFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAskAddFriendRsp }
func (*AskAddFriendRsp) ProtoMessageType() ProtoMessageType { return "AskAddFriendRsp" }

func (*DealAddFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandDealAddFriendReq }
func (*DealAddFriendReq) ProtoMessageType() ProtoMessageType { return "DealAddFriendReq" }

func (*DealAddFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDealAddFriendRsp }
func (*DealAddFriendRsp) ProtoMessageType() ProtoMessageType { return "DealAddFriendRsp" }

func (*GetPlayerSocialDetailReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerSocialDetailReq
}
func (*GetPlayerSocialDetailReq) ProtoMessageType() ProtoMessageType {
	return "GetPlayerSocialDetailReq"
}

func (*GetPlayerSocialDetailRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerSocialDetailRsp
}
func (*GetPlayerSocialDetailRsp) ProtoMessageType() ProtoMessageType {
	return "GetPlayerSocialDetailRsp"
}

func (*DeleteFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandDeleteFriendReq }
func (*DeleteFriendReq) ProtoMessageType() ProtoMessageType { return "DeleteFriendReq" }

func (*DeleteFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDeleteFriendRsp }
func (*DeleteFriendRsp) ProtoMessageType() ProtoMessageType { return "DeleteFriendRsp" }

func (*SetPlayerBirthdayReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerBirthdayReq }
func (*SetPlayerBirthdayReq) ProtoMessageType() ProtoMessageType { return "SetPlayerBirthdayReq" }

func (*SetPlayerBirthdayRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerBirthdayRsp }
func (*SetPlayerBirthdayRsp) ProtoMessageType() ProtoMessageType { return "SetPlayerBirthdayRsp" }

func (*SetPlayerSignatureReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerSignatureReq }
func (*SetPlayerSignatureReq) ProtoMessageType() ProtoMessageType { return "SetPlayerSignatureReq" }

func (*SetPlayerSignatureRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerSignatureRsp }
func (*SetPlayerSignatureRsp) ProtoMessageType() ProtoMessageType { return "SetPlayerSignatureRsp" }

func (*SetPlayerHeadImageReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerHeadImageReq }
func (*SetPlayerHeadImageReq) ProtoMessageType() ProtoMessageType { return "SetPlayerHeadImageReq" }

func (*SetPlayerHeadImageRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetPlayerHeadImageRsp }
func (*SetPlayerHeadImageRsp) ProtoMessageType() ProtoMessageType { return "SetPlayerHeadImageRsp" }

func (*UpdatePS4FriendListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdatePS4FriendListNotify
}
func (*UpdatePS4FriendListNotify) ProtoMessageType() ProtoMessageType {
	return "UpdatePS4FriendListNotify"
}

func (*DeleteFriendNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDeleteFriendNotify }
func (*DeleteFriendNotify) ProtoMessageType() ProtoMessageType { return "DeleteFriendNotify" }

func (*AddFriendNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAddFriendNotify }
func (*AddFriendNotify) ProtoMessageType() ProtoMessageType { return "AddFriendNotify" }

func (*AskAddFriendNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAskAddFriendNotify }
func (*AskAddFriendNotify) ProtoMessageType() ProtoMessageType { return "AskAddFriendNotify" }

func (*SetNameCardReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetNameCardReq }
func (*SetNameCardReq) ProtoMessageType() ProtoMessageType { return "SetNameCardReq" }

func (*SetNameCardRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetNameCardRsp }
func (*SetNameCardRsp) ProtoMessageType() ProtoMessageType { return "SetNameCardRsp" }

func (*GetAllUnlockNameCardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllUnlockNameCardReq
}
func (*GetAllUnlockNameCardReq) ProtoMessageType() ProtoMessageType { return "GetAllUnlockNameCardReq" }

func (*GetAllUnlockNameCardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetAllUnlockNameCardRsp
}
func (*GetAllUnlockNameCardRsp) ProtoMessageType() ProtoMessageType { return "GetAllUnlockNameCardRsp" }

func (*AddBlacklistReq) ProtoCommand() ProtoCommand         { return ProtoCommandAddBlacklistReq }
func (*AddBlacklistReq) ProtoMessageType() ProtoMessageType { return "AddBlacklistReq" }

func (*AddBlacklistRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAddBlacklistRsp }
func (*AddBlacklistRsp) ProtoMessageType() ProtoMessageType { return "AddBlacklistRsp" }

func (*RemoveBlacklistReq) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveBlacklistReq }
func (*RemoveBlacklistReq) ProtoMessageType() ProtoMessageType { return "RemoveBlacklistReq" }

func (*RemoveBlacklistRsp) ProtoCommand() ProtoCommand         { return ProtoCommandRemoveBlacklistRsp }
func (*RemoveBlacklistRsp) ProtoMessageType() ProtoMessageType { return "RemoveBlacklistRsp" }

func (*UnlockNameCardNotify) ProtoCommand() ProtoCommand         { return ProtoCommandUnlockNameCardNotify }
func (*UnlockNameCardNotify) ProtoMessageType() ProtoMessageType { return "UnlockNameCardNotify" }

func (*GetRecentMpPlayerListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRecentMpPlayerListReq
}
func (*GetRecentMpPlayerListReq) ProtoMessageType() ProtoMessageType {
	return "GetRecentMpPlayerListReq"
}

func (*GetRecentMpPlayerListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetRecentMpPlayerListRsp
}
func (*GetRecentMpPlayerListRsp) ProtoMessageType() ProtoMessageType {
	return "GetRecentMpPlayerListRsp"
}

func (*SocialDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSocialDataNotify }
func (*SocialDataNotify) ProtoMessageType() ProtoMessageType { return "SocialDataNotify" }

func (*TakeFirstShareRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeFirstShareRewardReq
}
func (*TakeFirstShareRewardReq) ProtoMessageType() ProtoMessageType { return "TakeFirstShareRewardReq" }

func (*TakeFirstShareRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeFirstShareRewardRsp
}
func (*TakeFirstShareRewardRsp) ProtoMessageType() ProtoMessageType { return "TakeFirstShareRewardRsp" }

func (*UpdatePS4BlockListReq) ProtoCommand() ProtoCommand         { return ProtoCommandUpdatePS4BlockListReq }
func (*UpdatePS4BlockListReq) ProtoMessageType() ProtoMessageType { return "UpdatePS4BlockListReq" }

func (*UpdatePS4BlockListRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUpdatePS4BlockListRsp }
func (*UpdatePS4BlockListRsp) ProtoMessageType() ProtoMessageType { return "UpdatePS4BlockListRsp" }

func (*GetPlayerBlacklistReq) ProtoCommand() ProtoCommand         { return ProtoCommandGetPlayerBlacklistReq }
func (*GetPlayerBlacklistReq) ProtoMessageType() ProtoMessageType { return "GetPlayerBlacklistReq" }

func (*GetPlayerBlacklistRsp) ProtoCommand() ProtoCommand         { return ProtoCommandGetPlayerBlacklistRsp }
func (*GetPlayerBlacklistRsp) ProtoMessageType() ProtoMessageType { return "GetPlayerBlacklistRsp" }

func (*PlayerReportReq) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerReportReq }
func (*PlayerReportReq) ProtoMessageType() ProtoMessageType { return "PlayerReportReq" }

func (*PlayerReportRsp) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerReportRsp }
func (*PlayerReportRsp) ProtoMessageType() ProtoMessageType { return "PlayerReportRsp" }

func (*SetFriendRemarkNameReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetFriendRemarkNameReq }
func (*SetFriendRemarkNameReq) ProtoMessageType() ProtoMessageType { return "SetFriendRemarkNameReq" }

func (*SetFriendRemarkNameRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetFriendRemarkNameRsp }
func (*SetFriendRemarkNameRsp) ProtoMessageType() ProtoMessageType { return "SetFriendRemarkNameRsp" }

func (*UpdatePlayerShowAvatarListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdatePlayerShowAvatarListReq
}
func (*UpdatePlayerShowAvatarListReq) ProtoMessageType() ProtoMessageType {
	return "UpdatePlayerShowAvatarListReq"
}

func (*UpdatePlayerShowAvatarListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdatePlayerShowAvatarListRsp
}
func (*UpdatePlayerShowAvatarListRsp) ProtoMessageType() ProtoMessageType {
	return "UpdatePlayerShowAvatarListRsp"
}

func (*GetFriendShowAvatarInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetFriendShowAvatarInfoReq
}
func (*GetFriendShowAvatarInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetFriendShowAvatarInfoReq"
}

func (*GetFriendShowAvatarInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetFriendShowAvatarInfoRsp
}
func (*GetFriendShowAvatarInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetFriendShowAvatarInfoRsp"
}

func (*UpdatePlayerShowNameCardListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdatePlayerShowNameCardListReq
}
func (*UpdatePlayerShowNameCardListReq) ProtoMessageType() ProtoMessageType {
	return "UpdatePlayerShowNameCardListReq"
}

func (*UpdatePlayerShowNameCardListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandUpdatePlayerShowNameCardListRsp
}
func (*UpdatePlayerShowNameCardListRsp) ProtoMessageType() ProtoMessageType {
	return "UpdatePlayerShowNameCardListRsp"
}

func (*GetFriendShowNameCardInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetFriendShowNameCardInfoReq
}
func (*GetFriendShowNameCardInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetFriendShowNameCardInfoReq"
}

func (*GetFriendShowNameCardInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetFriendShowNameCardInfoRsp
}
func (*GetFriendShowNameCardInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetFriendShowNameCardInfoRsp"
}

func (*ForceAddPlayerFriendReq) ProtoCommand() ProtoCommand {
	return ProtoCommandForceAddPlayerFriendReq
}
func (*ForceAddPlayerFriendReq) ProtoMessageType() ProtoMessageType { return "ForceAddPlayerFriendReq" }

func (*ForceAddPlayerFriendRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandForceAddPlayerFriendRsp
}
func (*ForceAddPlayerFriendRsp) ProtoMessageType() ProtoMessageType { return "ForceAddPlayerFriendRsp" }

func (*ProfilePictureChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandProfilePictureChangeNotify
}
func (*ProfilePictureChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ProfilePictureChangeNotify"
}

func (*PSNFriendListNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPSNFriendListNotify }
func (*PSNFriendListNotify) ProtoMessageType() ProtoMessageType { return "PSNFriendListNotify" }

func (*PSNBlackListNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPSNBlackListNotify }
func (*PSNBlackListNotify) ProtoMessageType() ProtoMessageType { return "PSNBlackListNotify" }

func (*GetPlayerAskFriendListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerAskFriendListReq
}
func (*GetPlayerAskFriendListReq) ProtoMessageType() ProtoMessageType {
	return "GetPlayerAskFriendListReq"
}

func (*GetPlayerAskFriendListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetPlayerAskFriendListRsp
}
func (*GetPlayerAskFriendListRsp) ProtoMessageType() ProtoMessageType {
	return "GetPlayerAskFriendListRsp"
}

func (*GetChatEmojiCollectionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetChatEmojiCollectionReq
}
func (*GetChatEmojiCollectionReq) ProtoMessageType() ProtoMessageType {
	return "GetChatEmojiCollectionReq"
}

func (*GetChatEmojiCollectionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetChatEmojiCollectionRsp
}
func (*GetChatEmojiCollectionRsp) ProtoMessageType() ProtoMessageType {
	return "GetChatEmojiCollectionRsp"
}

func (*SetChatEmojiCollectionReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetChatEmojiCollectionReq
}
func (*SetChatEmojiCollectionReq) ProtoMessageType() ProtoMessageType {
	return "SetChatEmojiCollectionReq"
}

func (*SetChatEmojiCollectionRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetChatEmojiCollectionRsp
}
func (*SetChatEmojiCollectionRsp) ProtoMessageType() ProtoMessageType {
	return "SetChatEmojiCollectionRsp"
}

func (*UpdatePS4FriendListReq) ProtoCommand() ProtoCommand         { return ProtoCommandUpdatePS4FriendListReq }
func (*UpdatePS4FriendListReq) ProtoMessageType() ProtoMessageType { return "UpdatePS4FriendListReq" }

func (*UpdatePS4FriendListRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUpdatePS4FriendListRsp }
func (*UpdatePS4FriendListRsp) ProtoMessageType() ProtoMessageType { return "UpdatePS4FriendListRsp" }

func (*FriendInfoChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandFriendInfoChangeNotify }
func (*FriendInfoChangeNotify) ProtoMessageType() ProtoMessageType { return "FriendInfoChangeNotify" }

func (*PlayerSignatureAuditDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerSignatureAuditDataNotify
}
func (*PlayerSignatureAuditDataNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerSignatureAuditDataNotify"
}

func (*PlayerSignatureNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerSignatureNotify }
func (*PlayerSignatureNotify) ProtoMessageType() ProtoMessageType { return "PlayerSignatureNotify" }

func (*SignatureAuditConfigNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSignatureAuditConfigNotify
}
func (*SignatureAuditConfigNotify) ProtoMessageType() ProtoMessageType {
	return "SignatureAuditConfigNotify"
}

func (*ReadSignatureAuditReq) ProtoCommand() ProtoCommand         { return ProtoCommandReadSignatureAuditReq }
func (*ReadSignatureAuditReq) ProtoMessageType() ProtoMessageType { return "ReadSignatureAuditReq" }

func (*ReadSignatureAuditRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReadSignatureAuditRsp }
func (*ReadSignatureAuditRsp) ProtoMessageType() ProtoMessageType { return "ReadSignatureAuditRsp" }
