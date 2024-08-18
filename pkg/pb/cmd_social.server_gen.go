
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AddAskFriendNotify) },
		func() ProtoMessage { return new(ServerGetPlayerFriendListReq) },
		func() ProtoMessage { return new(ServerGetPlayerFriendListRsp) },
		func() ProtoMessage { return new(ServerGetSocialDetailReq) },
		func() ProtoMessage { return new(ServerGetSocialDetailRsp) },
		func() ProtoMessage { return new(ServerAskAddFriendReq) },
		func() ProtoMessage { return new(ServerAskAddFriendRsp) },
		func() ProtoMessage { return new(ServerDealAddFriendReq) },
		func() ProtoMessage { return new(ServerDealAddFriendRsp) },
		func() ProtoMessage { return new(ServerDeleteFriendReq) },
		func() ProtoMessage { return new(ServerDeleteFriendRsp) },
		func() ProtoMessage { return new(ServerSetSignatureReq) },
		func() ProtoMessage { return new(ServerSetSignatureRsp) },
		func() ProtoMessage { return new(ServerGetPlayerFriendBriefReq) },
		func() ProtoMessage { return new(ServerGetPlayerFriendBriefRsp) },
		func() ProtoMessage { return new(SeverGetPS4FriendUidReq) },
		func() ProtoMessage { return new(SeverGetPS4FriendUidRsp) },
		func() ProtoMessage { return new(SyncPlayerBriefNotify) },
		func() ProtoMessage { return new(ServerAddBlacklistReq) },
		func() ProtoMessage { return new(ServerAddBlacklistRsp) },
		func() ProtoMessage { return new(ServerRemoveBlacklistReq) },
		func() ProtoMessage { return new(ServerRemoveBlacklistRsp) },
		func() ProtoMessage { return new(ServerGetRecentMpPlayerListReq) },
		func() ProtoMessage { return new(ServerGetRecentMpPlayerListRsp) },
		func() ProtoMessage { return new(ServerGetPlayerBlacklistReq) },
		func() ProtoMessage { return new(ServerGetPlayerBlacklistRsp) },
		func() ProtoMessage { return new(ServerPrivateChatReq) },
		func() ProtoMessage { return new(ServerPrivateChatRsp) },
		func() ProtoMessage { return new(ServerPullPrivateChatReq) },
		func() ProtoMessage { return new(ServerPullRecentChatReq) },
		func() ProtoMessage { return new(ServerUpdateActivitySocialDataNotify) },
		func() ProtoMessage { return new(ServerBlessingGetFriendPicListReq) },
		func() ProtoMessage { return new(ServerBlessingGetFriendPicListRsp) },
		func() ProtoMessage { return new(ServerGetFriendBriefReq) },
		func() ProtoMessage { return new(ServerGetFriendBriefRsp) },
		func() ProtoMessage { return new(ServerUpdateShowAvatarInfoNotify) },
		func() ProtoMessage { return new(ServerGetFriendShowAvatarInfoReq) },
		func() ProtoMessage { return new(ServerGetFriendShowAvatarInfoRsp) },
		func() ProtoMessage { return new(ServerReadPrivateChatReq) },
		func() ProtoMessage { return new(ServerGetFriendShowNameCardInfoReq) },
		func() ProtoMessage { return new(ServerGetFriendShowNameCardInfoRsp) },
		func() ProtoMessage { return new(ServerGetAskFriendBriefReq) },
		func() ProtoMessage { return new(ServerGetAskFriendBriefRsp) },
		func() ProtoMessage { return new(ServerAddPsnFriendReq) },
		func() ProtoMessage { return new(ServerAddPsnFriendRsp) },
		func() ProtoMessage { return new(ServerAddPsnBlackReq) },
		func() ProtoMessage { return new(ServerAddPsnBlackRsp) },
		func() ProtoMessage { return new(ServerFriendInfoChangeNotify) },
		func() ProtoMessage { return new(ServerPlantFlowerGetFriendFlowerDataReq) },
		func() ProtoMessage { return new(ServerPlantFlowerGetFriendFlowerDataRsp) },
		func() ProtoMessage { return new(ServerPlantFlowerGetFriendFlowerWishListReq) },
		func() ProtoMessage { return new(ServerPlantFlowerGetFriendFlowerWishListRsp) },
		func() ProtoMessage { return new(ServerWinterCampGetFriendItemDataReq) },
		func() ProtoMessage { return new(ServerWinterCampGetFriendItemDataRsp) },
		func() ProtoMessage { return new(ServerWinterCampGetFriendWishListReq) },
		func() ProtoMessage { return new(ServerWinterCampGetFriendWishListRsp) },
		func() ProtoMessage { return new(ServerGetCustomDungeonReq) },
		func() ProtoMessage { return new(ServerGetCustomDungeonRsp) },
		func() ProtoMessage { return new(ServerSaveCustomDungeonReq) },
		func() ProtoMessage { return new(ServerSaveCustomDungeonRsp) },
		func() ProtoMessage { return new(ServerPublishCustomDungeonReq) },
		func() ProtoMessage { return new(ServerPublishCustomDungeonRsp) },
		func() ProtoMessage { return new(ServerRemoveCustomDungeonReq) },
		func() ProtoMessage { return new(ServerRemoveCustomDungeonRsp) },
		func() ProtoMessage { return new(ServerUpdateCustomDungeonSocialNotify) },
		func() ProtoMessage { return new(ServerGetCustomDungeonBriefReq) },
		func() ProtoMessage { return new(ServerGetCustomDungeonBriefRsp) },
		func() ProtoMessage { return new(ServerGetRecommendCustomDungeonReq) },
		func() ProtoMessage { return new(ServerGetRecommendCustomDungeonRsp) },
		func() ProtoMessage { return new(ServerAddFriendByMuipReq) },
		func() ProtoMessage { return new(ServerAddFriendByMuipRsp) },
		func() ProtoMessage { return new(ServerDelFriendByMuipReq) },
		func() ProtoMessage { return new(ServerDelFriendByMuipRsp) },
		func() ProtoMessage { return new(ServerAddFriendAskByMuipReq) },
		func() ProtoMessage { return new(ServerAddFriendAskByMuipRsp) },
		func() ProtoMessage { return new(ServerDelFriendAskByMuipReq) },
		func() ProtoMessage { return new(ServerDelFriendAskByMuipRsp) },
		func() ProtoMessage { return new(ServerCustomDungeonCacheInvalidNotify) },
		func() ProtoMessage { return new(ServerCustomDungeonSocialGmNotify) },
		func() ProtoMessage { return new(ServerCustomDungeonFirstLikeNotify) },
		func() ProtoMessage { return new(ServerGetUgcReq) },
		func() ProtoMessage { return new(ServerGetUgcRsp) },
		func() ProtoMessage { return new(ServerGetUgcBriefReq) },
		func() ProtoMessage { return new(ServerGetUgcBriefRsp) },
		func() ProtoMessage { return new(ServerMultiGetUgcBriefReq) },
		func() ProtoMessage { return new(ServerMultiGetUgcBriefRsp) },
		func() ProtoMessage { return new(ServerSaveUgcReq) },
		func() ProtoMessage { return new(ServerSaveUgcRsp) },
		func() ProtoMessage { return new(ServerPublishUgcReq) },
		func() ProtoMessage { return new(ServerPublishUgcRsp) },
		func() ProtoMessage { return new(ServerCheckUgcUpdateReq) },
		func() ProtoMessage { return new(ServerCheckUgcUpdateRsp) },
		func() ProtoMessage { return new(ServerActivityGetFriendGiftDataReq) },
		func() ProtoMessage { return new(ServerActivityGetFriendGiftDataRsp) },
		func() ProtoMessage { return new(ServerActivityGetFriendGiftWishListReq) },
		func() ProtoMessage { return new(ServerActivityGetFriendGiftWishListRsp) },
		func() ProtoMessage { return new(SyncPlayerIpRegionNotify) },
	)
}

const (
	ProtoCommandAddAskFriendNotify                          ProtoCommand = 10872
	ProtoCommandServerGetPlayerFriendListReq                ProtoCommand = 10898
	ProtoCommandServerGetPlayerFriendListRsp                ProtoCommand = 10812
	ProtoCommandServerGetSocialDetailReq                    ProtoCommand = 10807
	ProtoCommandServerGetSocialDetailRsp                    ProtoCommand = 10821
	ProtoCommandServerAskAddFriendReq                       ProtoCommand = 10803
	ProtoCommandServerAskAddFriendRsp                       ProtoCommand = 10890
	ProtoCommandServerDealAddFriendReq                      ProtoCommand = 10873
	ProtoCommandServerDealAddFriendRsp                      ProtoCommand = 10899
	ProtoCommandServerDeleteFriendReq                       ProtoCommand = 10831
	ProtoCommandServerDeleteFriendRsp                       ProtoCommand = 10875
	ProtoCommandServerSetSignatureReq                       ProtoCommand = 10881
	ProtoCommandServerSetSignatureRsp                       ProtoCommand = 10805
	ProtoCommandServerGetPlayerFriendBriefReq               ProtoCommand = 10847
	ProtoCommandServerGetPlayerFriendBriefRsp               ProtoCommand = 10839
	ProtoCommandSeverGetPS4FriendUidReq                     ProtoCommand = 10853
	ProtoCommandSeverGetPS4FriendUidRsp                     ProtoCommand = 10822
	ProtoCommandSyncPlayerBriefNotify                       ProtoCommand = 10865
	ProtoCommandServerAddBlacklistReq                       ProtoCommand = 10804
	ProtoCommandServerAddBlacklistRsp                       ProtoCommand = 10893
	ProtoCommandServerRemoveBlacklistReq                    ProtoCommand = 10827
	ProtoCommandServerRemoveBlacklistRsp                    ProtoCommand = 10894
	ProtoCommandServerGetRecentMpPlayerListReq              ProtoCommand = 10888
	ProtoCommandServerGetRecentMpPlayerListRsp              ProtoCommand = 10826
	ProtoCommandServerGetPlayerBlacklistReq                 ProtoCommand = 10863
	ProtoCommandServerGetPlayerBlacklistRsp                 ProtoCommand = 10895
	ProtoCommandServerPrivateChatReq                        ProtoCommand = 10806
	ProtoCommandServerPrivateChatRsp                        ProtoCommand = 10834
	ProtoCommandServerPullPrivateChatReq                    ProtoCommand = 10850
	ProtoCommandServerPullRecentChatReq                     ProtoCommand = 10843
	ProtoCommandServerUpdateActivitySocialDataNotify        ProtoCommand = 10874
	ProtoCommandServerBlessingGetFriendPicListReq           ProtoCommand = 10876
	ProtoCommandServerBlessingGetFriendPicListRsp           ProtoCommand = 10846
	ProtoCommandServerGetFriendBriefReq                     ProtoCommand = 10841
	ProtoCommandServerGetFriendBriefRsp                     ProtoCommand = 10849
	ProtoCommandServerUpdateShowAvatarInfoNotify            ProtoCommand = 10891
	ProtoCommandServerGetFriendShowAvatarInfoReq            ProtoCommand = 10824
	ProtoCommandServerGetFriendShowAvatarInfoRsp            ProtoCommand = 10856
	ProtoCommandServerReadPrivateChatReq                    ProtoCommand = 10842
	ProtoCommandServerGetFriendShowNameCardInfoReq          ProtoCommand = 10830
	ProtoCommandServerGetFriendShowNameCardInfoRsp          ProtoCommand = 10867
	ProtoCommandServerGetAskFriendBriefReq                  ProtoCommand = 10858
	ProtoCommandServerGetAskFriendBriefRsp                  ProtoCommand = 10870
	ProtoCommandServerAddPsnFriendReq                       ProtoCommand = 10817
	ProtoCommandServerAddPsnFriendRsp                       ProtoCommand = 10802
	ProtoCommandServerAddPsnBlackReq                        ProtoCommand = 10819
	ProtoCommandServerAddPsnBlackRsp                        ProtoCommand = 10861
	ProtoCommandServerFriendInfoChangeNotify                ProtoCommand = 10829
	ProtoCommandServerPlantFlowerGetFriendFlowerDataReq     ProtoCommand = 10857
	ProtoCommandServerPlantFlowerGetFriendFlowerDataRsp     ProtoCommand = 10900
	ProtoCommandServerPlantFlowerGetFriendFlowerWishListReq ProtoCommand = 10816
	ProtoCommandServerPlantFlowerGetFriendFlowerWishListRsp ProtoCommand = 10887
	ProtoCommandServerWinterCampGetFriendItemDataReq        ProtoCommand = 10840
	ProtoCommandServerWinterCampGetFriendItemDataRsp        ProtoCommand = 10818
	ProtoCommandServerWinterCampGetFriendWishListReq        ProtoCommand = 10866
	ProtoCommandServerWinterCampGetFriendWishListRsp        ProtoCommand = 10868
	ProtoCommandServerGetCustomDungeonReq                   ProtoCommand = 10833
	ProtoCommandServerGetCustomDungeonRsp                   ProtoCommand = 10884
	ProtoCommandServerSaveCustomDungeonReq                  ProtoCommand = 10880
	ProtoCommandServerSaveCustomDungeonRsp                  ProtoCommand = 10889
	ProtoCommandServerPublishCustomDungeonReq               ProtoCommand = 10859
	ProtoCommandServerPublishCustomDungeonRsp               ProtoCommand = 10832
	ProtoCommandServerRemoveCustomDungeonReq                ProtoCommand = 10860
	ProtoCommandServerRemoveCustomDungeonRsp                ProtoCommand = 10814
	ProtoCommandServerUpdateCustomDungeonSocialNotify       ProtoCommand = 10892
	ProtoCommandServerGetCustomDungeonBriefReq              ProtoCommand = 10820
	ProtoCommandServerGetCustomDungeonBriefRsp              ProtoCommand = 10864
	ProtoCommandServerGetRecommendCustomDungeonReq          ProtoCommand = 10838
	ProtoCommandServerGetRecommendCustomDungeonRsp          ProtoCommand = 10811
	ProtoCommandServerAddFriendByMuipReq                    ProtoCommand = 10844
	ProtoCommandServerAddFriendByMuipRsp                    ProtoCommand = 10886
	ProtoCommandServerDelFriendByMuipReq                    ProtoCommand = 10801
	ProtoCommandServerDelFriendByMuipRsp                    ProtoCommand = 10828
	ProtoCommandServerAddFriendAskByMuipReq                 ProtoCommand = 10851
	ProtoCommandServerAddFriendAskByMuipRsp                 ProtoCommand = 10823
	ProtoCommandServerDelFriendAskByMuipReq                 ProtoCommand = 10871
	ProtoCommandServerDelFriendAskByMuipRsp                 ProtoCommand = 10845
	ProtoCommandServerCustomDungeonCacheInvalidNotify       ProtoCommand = 10808
	ProtoCommandServerCustomDungeonSocialGmNotify           ProtoCommand = 10809
	ProtoCommandServerCustomDungeonFirstLikeNotify          ProtoCommand = 10852
	ProtoCommandServerGetUgcReq                             ProtoCommand = 10885
	ProtoCommandServerGetUgcRsp                             ProtoCommand = 10825
	ProtoCommandServerGetUgcBriefReq                        ProtoCommand = 10878
	ProtoCommandServerGetUgcBriefRsp                        ProtoCommand = 10836
	ProtoCommandServerMultiGetUgcBriefReq                   ProtoCommand = 10854
	ProtoCommandServerMultiGetUgcBriefRsp                   ProtoCommand = 10883
	ProtoCommandServerSaveUgcReq                            ProtoCommand = 10869
	ProtoCommandServerSaveUgcRsp                            ProtoCommand = 10810
	ProtoCommandServerPublishUgcReq                         ProtoCommand = 10896
	ProtoCommandServerPublishUgcRsp                         ProtoCommand = 10862
	ProtoCommandServerCheckUgcUpdateReq                     ProtoCommand = 10815
	ProtoCommandServerCheckUgcUpdateRsp                     ProtoCommand = 10813
	ProtoCommandServerActivityGetFriendGiftDataReq          ProtoCommand = 10879
	ProtoCommandServerActivityGetFriendGiftDataRsp          ProtoCommand = 10855
	ProtoCommandServerActivityGetFriendGiftWishListReq      ProtoCommand = 10976
	ProtoCommandServerActivityGetFriendGiftWishListRsp      ProtoCommand = 10968
	ProtoCommandSyncPlayerIpRegionNotify                    ProtoCommand = 10975
)

func (*AddAskFriendNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAddAskFriendNotify }
func (*AddAskFriendNotify) ProtoMessageType() ProtoMessageType { return "AddAskFriendNotify" }

func (*ServerGetPlayerFriendListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetPlayerFriendListReq
}
func (*ServerGetPlayerFriendListReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetPlayerFriendListReq"
}

func (*ServerGetPlayerFriendListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetPlayerFriendListRsp
}
func (*ServerGetPlayerFriendListRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetPlayerFriendListRsp"
}

func (*ServerGetSocialDetailReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetSocialDetailReq
}
func (*ServerGetSocialDetailReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetSocialDetailReq"
}

func (*ServerGetSocialDetailRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetSocialDetailRsp
}
func (*ServerGetSocialDetailRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetSocialDetailRsp"
}

func (*ServerAskAddFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerAskAddFriendReq }
func (*ServerAskAddFriendReq) ProtoMessageType() ProtoMessageType { return "ServerAskAddFriendReq" }

func (*ServerAskAddFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerAskAddFriendRsp }
func (*ServerAskAddFriendRsp) ProtoMessageType() ProtoMessageType { return "ServerAskAddFriendRsp" }

func (*ServerDealAddFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerDealAddFriendReq }
func (*ServerDealAddFriendReq) ProtoMessageType() ProtoMessageType { return "ServerDealAddFriendReq" }

func (*ServerDealAddFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerDealAddFriendRsp }
func (*ServerDealAddFriendRsp) ProtoMessageType() ProtoMessageType { return "ServerDealAddFriendRsp" }

func (*ServerDeleteFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerDeleteFriendReq }
func (*ServerDeleteFriendReq) ProtoMessageType() ProtoMessageType { return "ServerDeleteFriendReq" }

func (*ServerDeleteFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerDeleteFriendRsp }
func (*ServerDeleteFriendRsp) ProtoMessageType() ProtoMessageType { return "ServerDeleteFriendRsp" }

func (*ServerSetSignatureReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerSetSignatureReq }
func (*ServerSetSignatureReq) ProtoMessageType() ProtoMessageType { return "ServerSetSignatureReq" }

func (*ServerSetSignatureRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerSetSignatureRsp }
func (*ServerSetSignatureRsp) ProtoMessageType() ProtoMessageType { return "ServerSetSignatureRsp" }

func (*ServerGetPlayerFriendBriefReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetPlayerFriendBriefReq
}
func (*ServerGetPlayerFriendBriefReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetPlayerFriendBriefReq"
}

func (*ServerGetPlayerFriendBriefRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetPlayerFriendBriefRsp
}
func (*ServerGetPlayerFriendBriefRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetPlayerFriendBriefRsp"
}

func (*SeverGetPS4FriendUidReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSeverGetPS4FriendUidReq
}
func (*SeverGetPS4FriendUidReq) ProtoMessageType() ProtoMessageType { return "SeverGetPS4FriendUidReq" }

func (*SeverGetPS4FriendUidRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSeverGetPS4FriendUidRsp
}
func (*SeverGetPS4FriendUidRsp) ProtoMessageType() ProtoMessageType { return "SeverGetPS4FriendUidRsp" }

func (*SyncPlayerBriefNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSyncPlayerBriefNotify }
func (*SyncPlayerBriefNotify) ProtoMessageType() ProtoMessageType { return "SyncPlayerBriefNotify" }

func (*ServerAddBlacklistReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerAddBlacklistReq }
func (*ServerAddBlacklistReq) ProtoMessageType() ProtoMessageType { return "ServerAddBlacklistReq" }

func (*ServerAddBlacklistRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerAddBlacklistRsp }
func (*ServerAddBlacklistRsp) ProtoMessageType() ProtoMessageType { return "ServerAddBlacklistRsp" }

func (*ServerRemoveBlacklistReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerRemoveBlacklistReq
}
func (*ServerRemoveBlacklistReq) ProtoMessageType() ProtoMessageType {
	return "ServerRemoveBlacklistReq"
}

func (*ServerRemoveBlacklistRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerRemoveBlacklistRsp
}
func (*ServerRemoveBlacklistRsp) ProtoMessageType() ProtoMessageType {
	return "ServerRemoveBlacklistRsp"
}

func (*ServerGetRecentMpPlayerListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetRecentMpPlayerListReq
}
func (*ServerGetRecentMpPlayerListReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetRecentMpPlayerListReq"
}

func (*ServerGetRecentMpPlayerListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetRecentMpPlayerListRsp
}
func (*ServerGetRecentMpPlayerListRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetRecentMpPlayerListRsp"
}

func (*ServerGetPlayerBlacklistReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetPlayerBlacklistReq
}
func (*ServerGetPlayerBlacklistReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetPlayerBlacklistReq"
}

func (*ServerGetPlayerBlacklistRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetPlayerBlacklistRsp
}
func (*ServerGetPlayerBlacklistRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetPlayerBlacklistRsp"
}

func (*ServerPrivateChatReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerPrivateChatReq }
func (*ServerPrivateChatReq) ProtoMessageType() ProtoMessageType { return "ServerPrivateChatReq" }

func (*ServerPrivateChatRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerPrivateChatRsp }
func (*ServerPrivateChatRsp) ProtoMessageType() ProtoMessageType { return "ServerPrivateChatRsp" }

func (*ServerPullPrivateChatReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPullPrivateChatReq
}
func (*ServerPullPrivateChatReq) ProtoMessageType() ProtoMessageType {
	return "ServerPullPrivateChatReq"
}

func (*ServerPullRecentChatReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPullRecentChatReq
}
func (*ServerPullRecentChatReq) ProtoMessageType() ProtoMessageType { return "ServerPullRecentChatReq" }

func (*ServerUpdateActivitySocialDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerUpdateActivitySocialDataNotify
}
func (*ServerUpdateActivitySocialDataNotify) ProtoMessageType() ProtoMessageType {
	return "ServerUpdateActivitySocialDataNotify"
}

func (*ServerBlessingGetFriendPicListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerBlessingGetFriendPicListReq
}
func (*ServerBlessingGetFriendPicListReq) ProtoMessageType() ProtoMessageType {
	return "ServerBlessingGetFriendPicListReq"
}

func (*ServerBlessingGetFriendPicListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerBlessingGetFriendPicListRsp
}
func (*ServerBlessingGetFriendPicListRsp) ProtoMessageType() ProtoMessageType {
	return "ServerBlessingGetFriendPicListRsp"
}

func (*ServerGetFriendBriefReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendBriefReq
}
func (*ServerGetFriendBriefReq) ProtoMessageType() ProtoMessageType { return "ServerGetFriendBriefReq" }

func (*ServerGetFriendBriefRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendBriefRsp
}
func (*ServerGetFriendBriefRsp) ProtoMessageType() ProtoMessageType { return "ServerGetFriendBriefRsp" }

func (*ServerUpdateShowAvatarInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerUpdateShowAvatarInfoNotify
}
func (*ServerUpdateShowAvatarInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ServerUpdateShowAvatarInfoNotify"
}

func (*ServerGetFriendShowAvatarInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendShowAvatarInfoReq
}
func (*ServerGetFriendShowAvatarInfoReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetFriendShowAvatarInfoReq"
}

func (*ServerGetFriendShowAvatarInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendShowAvatarInfoRsp
}
func (*ServerGetFriendShowAvatarInfoRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetFriendShowAvatarInfoRsp"
}

func (*ServerReadPrivateChatReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerReadPrivateChatReq
}
func (*ServerReadPrivateChatReq) ProtoMessageType() ProtoMessageType {
	return "ServerReadPrivateChatReq"
}

func (*ServerGetFriendShowNameCardInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendShowNameCardInfoReq
}
func (*ServerGetFriendShowNameCardInfoReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetFriendShowNameCardInfoReq"
}

func (*ServerGetFriendShowNameCardInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetFriendShowNameCardInfoRsp
}
func (*ServerGetFriendShowNameCardInfoRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetFriendShowNameCardInfoRsp"
}

func (*ServerGetAskFriendBriefReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetAskFriendBriefReq
}
func (*ServerGetAskFriendBriefReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetAskFriendBriefReq"
}

func (*ServerGetAskFriendBriefRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetAskFriendBriefRsp
}
func (*ServerGetAskFriendBriefRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetAskFriendBriefRsp"
}

func (*ServerAddPsnFriendReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerAddPsnFriendReq }
func (*ServerAddPsnFriendReq) ProtoMessageType() ProtoMessageType { return "ServerAddPsnFriendReq" }

func (*ServerAddPsnFriendRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerAddPsnFriendRsp }
func (*ServerAddPsnFriendRsp) ProtoMessageType() ProtoMessageType { return "ServerAddPsnFriendRsp" }

func (*ServerAddPsnBlackReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerAddPsnBlackReq }
func (*ServerAddPsnBlackReq) ProtoMessageType() ProtoMessageType { return "ServerAddPsnBlackReq" }

func (*ServerAddPsnBlackRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerAddPsnBlackRsp }
func (*ServerAddPsnBlackRsp) ProtoMessageType() ProtoMessageType { return "ServerAddPsnBlackRsp" }

func (*ServerFriendInfoChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerFriendInfoChangeNotify
}
func (*ServerFriendInfoChangeNotify) ProtoMessageType() ProtoMessageType {
	return "ServerFriendInfoChangeNotify"
}

func (*ServerPlantFlowerGetFriendFlowerDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPlantFlowerGetFriendFlowerDataReq
}
func (*ServerPlantFlowerGetFriendFlowerDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerPlantFlowerGetFriendFlowerDataReq"
}

func (*ServerPlantFlowerGetFriendFlowerDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPlantFlowerGetFriendFlowerDataRsp
}
func (*ServerPlantFlowerGetFriendFlowerDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerPlantFlowerGetFriendFlowerDataRsp"
}

func (*ServerPlantFlowerGetFriendFlowerWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPlantFlowerGetFriendFlowerWishListReq
}
func (*ServerPlantFlowerGetFriendFlowerWishListReq) ProtoMessageType() ProtoMessageType {
	return "ServerPlantFlowerGetFriendFlowerWishListReq"
}

func (*ServerPlantFlowerGetFriendFlowerWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPlantFlowerGetFriendFlowerWishListRsp
}
func (*ServerPlantFlowerGetFriendFlowerWishListRsp) ProtoMessageType() ProtoMessageType {
	return "ServerPlantFlowerGetFriendFlowerWishListRsp"
}

func (*ServerWinterCampGetFriendItemDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerWinterCampGetFriendItemDataReq
}
func (*ServerWinterCampGetFriendItemDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerWinterCampGetFriendItemDataReq"
}

func (*ServerWinterCampGetFriendItemDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerWinterCampGetFriendItemDataRsp
}
func (*ServerWinterCampGetFriendItemDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerWinterCampGetFriendItemDataRsp"
}

func (*ServerWinterCampGetFriendWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerWinterCampGetFriendWishListReq
}
func (*ServerWinterCampGetFriendWishListReq) ProtoMessageType() ProtoMessageType {
	return "ServerWinterCampGetFriendWishListReq"
}

func (*ServerWinterCampGetFriendWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerWinterCampGetFriendWishListRsp
}
func (*ServerWinterCampGetFriendWishListRsp) ProtoMessageType() ProtoMessageType {
	return "ServerWinterCampGetFriendWishListRsp"
}

func (*ServerGetCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetCustomDungeonReq
}
func (*ServerGetCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetCustomDungeonReq"
}

func (*ServerGetCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetCustomDungeonRsp
}
func (*ServerGetCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetCustomDungeonRsp"
}

func (*ServerSaveCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerSaveCustomDungeonReq
}
func (*ServerSaveCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ServerSaveCustomDungeonReq"
}

func (*ServerSaveCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerSaveCustomDungeonRsp
}
func (*ServerSaveCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ServerSaveCustomDungeonRsp"
}

func (*ServerPublishCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPublishCustomDungeonReq
}
func (*ServerPublishCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ServerPublishCustomDungeonReq"
}

func (*ServerPublishCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerPublishCustomDungeonRsp
}
func (*ServerPublishCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ServerPublishCustomDungeonRsp"
}

func (*ServerRemoveCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerRemoveCustomDungeonReq
}
func (*ServerRemoveCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ServerRemoveCustomDungeonReq"
}

func (*ServerRemoveCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerRemoveCustomDungeonRsp
}
func (*ServerRemoveCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ServerRemoveCustomDungeonRsp"
}

func (*ServerUpdateCustomDungeonSocialNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerUpdateCustomDungeonSocialNotify
}
func (*ServerUpdateCustomDungeonSocialNotify) ProtoMessageType() ProtoMessageType {
	return "ServerUpdateCustomDungeonSocialNotify"
}

func (*ServerGetCustomDungeonBriefReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetCustomDungeonBriefReq
}
func (*ServerGetCustomDungeonBriefReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetCustomDungeonBriefReq"
}

func (*ServerGetCustomDungeonBriefRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetCustomDungeonBriefRsp
}
func (*ServerGetCustomDungeonBriefRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetCustomDungeonBriefRsp"
}

func (*ServerGetRecommendCustomDungeonReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetRecommendCustomDungeonReq
}
func (*ServerGetRecommendCustomDungeonReq) ProtoMessageType() ProtoMessageType {
	return "ServerGetRecommendCustomDungeonReq"
}

func (*ServerGetRecommendCustomDungeonRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerGetRecommendCustomDungeonRsp
}
func (*ServerGetRecommendCustomDungeonRsp) ProtoMessageType() ProtoMessageType {
	return "ServerGetRecommendCustomDungeonRsp"
}

func (*ServerAddFriendByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerAddFriendByMuipReq
}
func (*ServerAddFriendByMuipReq) ProtoMessageType() ProtoMessageType {
	return "ServerAddFriendByMuipReq"
}

func (*ServerAddFriendByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerAddFriendByMuipRsp
}
func (*ServerAddFriendByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "ServerAddFriendByMuipRsp"
}

func (*ServerDelFriendByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerDelFriendByMuipReq
}
func (*ServerDelFriendByMuipReq) ProtoMessageType() ProtoMessageType {
	return "ServerDelFriendByMuipReq"
}

func (*ServerDelFriendByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerDelFriendByMuipRsp
}
func (*ServerDelFriendByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "ServerDelFriendByMuipRsp"
}

func (*ServerAddFriendAskByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerAddFriendAskByMuipReq
}
func (*ServerAddFriendAskByMuipReq) ProtoMessageType() ProtoMessageType {
	return "ServerAddFriendAskByMuipReq"
}

func (*ServerAddFriendAskByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerAddFriendAskByMuipRsp
}
func (*ServerAddFriendAskByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "ServerAddFriendAskByMuipRsp"
}

func (*ServerDelFriendAskByMuipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerDelFriendAskByMuipReq
}
func (*ServerDelFriendAskByMuipReq) ProtoMessageType() ProtoMessageType {
	return "ServerDelFriendAskByMuipReq"
}

func (*ServerDelFriendAskByMuipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerDelFriendAskByMuipRsp
}
func (*ServerDelFriendAskByMuipRsp) ProtoMessageType() ProtoMessageType {
	return "ServerDelFriendAskByMuipRsp"
}

func (*ServerCustomDungeonCacheInvalidNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCustomDungeonCacheInvalidNotify
}
func (*ServerCustomDungeonCacheInvalidNotify) ProtoMessageType() ProtoMessageType {
	return "ServerCustomDungeonCacheInvalidNotify"
}

func (*ServerCustomDungeonSocialGmNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCustomDungeonSocialGmNotify
}
func (*ServerCustomDungeonSocialGmNotify) ProtoMessageType() ProtoMessageType {
	return "ServerCustomDungeonSocialGmNotify"
}

func (*ServerCustomDungeonFirstLikeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCustomDungeonFirstLikeNotify
}
func (*ServerCustomDungeonFirstLikeNotify) ProtoMessageType() ProtoMessageType {
	return "ServerCustomDungeonFirstLikeNotify"
}

func (*ServerGetUgcReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGetUgcReq }
func (*ServerGetUgcReq) ProtoMessageType() ProtoMessageType { return "ServerGetUgcReq" }

func (*ServerGetUgcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGetUgcRsp }
func (*ServerGetUgcRsp) ProtoMessageType() ProtoMessageType { return "ServerGetUgcRsp" }

func (*ServerGetUgcBriefReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerGetUgcBriefReq }
func (*ServerGetUgcBriefReq) ProtoMessageType() ProtoMessageType { return "ServerGetUgcBriefReq" }

func (*ServerGetUgcBriefRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerGetUgcBriefRsp }
func (*ServerGetUgcBriefRsp) ProtoMessageType() ProtoMessageType { return "ServerGetUgcBriefRsp" }

func (*ServerMultiGetUgcBriefReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerMultiGetUgcBriefReq
}
func (*ServerMultiGetUgcBriefReq) ProtoMessageType() ProtoMessageType {
	return "ServerMultiGetUgcBriefReq"
}

func (*ServerMultiGetUgcBriefRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerMultiGetUgcBriefRsp
}
func (*ServerMultiGetUgcBriefRsp) ProtoMessageType() ProtoMessageType {
	return "ServerMultiGetUgcBriefRsp"
}

func (*ServerSaveUgcReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerSaveUgcReq }
func (*ServerSaveUgcReq) ProtoMessageType() ProtoMessageType { return "ServerSaveUgcReq" }

func (*ServerSaveUgcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerSaveUgcRsp }
func (*ServerSaveUgcRsp) ProtoMessageType() ProtoMessageType { return "ServerSaveUgcRsp" }

func (*ServerPublishUgcReq) ProtoCommand() ProtoCommand         { return ProtoCommandServerPublishUgcReq }
func (*ServerPublishUgcReq) ProtoMessageType() ProtoMessageType { return "ServerPublishUgcReq" }

func (*ServerPublishUgcRsp) ProtoCommand() ProtoCommand         { return ProtoCommandServerPublishUgcRsp }
func (*ServerPublishUgcRsp) ProtoMessageType() ProtoMessageType { return "ServerPublishUgcRsp" }

func (*ServerCheckUgcUpdateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCheckUgcUpdateReq
}
func (*ServerCheckUgcUpdateReq) ProtoMessageType() ProtoMessageType { return "ServerCheckUgcUpdateReq" }

func (*ServerCheckUgcUpdateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerCheckUgcUpdateRsp
}
func (*ServerCheckUgcUpdateRsp) ProtoMessageType() ProtoMessageType { return "ServerCheckUgcUpdateRsp" }

func (*ServerActivityGetFriendGiftDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerActivityGetFriendGiftDataReq
}
func (*ServerActivityGetFriendGiftDataReq) ProtoMessageType() ProtoMessageType {
	return "ServerActivityGetFriendGiftDataReq"
}

func (*ServerActivityGetFriendGiftDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerActivityGetFriendGiftDataRsp
}
func (*ServerActivityGetFriendGiftDataRsp) ProtoMessageType() ProtoMessageType {
	return "ServerActivityGetFriendGiftDataRsp"
}

func (*ServerActivityGetFriendGiftWishListReq) ProtoCommand() ProtoCommand {
	return ProtoCommandServerActivityGetFriendGiftWishListReq
}
func (*ServerActivityGetFriendGiftWishListReq) ProtoMessageType() ProtoMessageType {
	return "ServerActivityGetFriendGiftWishListReq"
}

func (*ServerActivityGetFriendGiftWishListRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandServerActivityGetFriendGiftWishListRsp
}
func (*ServerActivityGetFriendGiftWishListRsp) ProtoMessageType() ProtoMessageType {
	return "ServerActivityGetFriendGiftWishListRsp"
}

func (*SyncPlayerIpRegionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandSyncPlayerIpRegionNotify
}
func (*SyncPlayerIpRegionNotify) ProtoMessageType() ProtoMessageType {
	return "SyncPlayerIpRegionNotify"
}
