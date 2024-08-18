
package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AvatarAddNotify) },
		func() ProtoMessage { return new(AvatarDelNotify) },
		func() ProtoMessage { return new(SetUpAvatarTeamReq) },
		func() ProtoMessage { return new(SetUpAvatarTeamRsp) },
		func() ProtoMessage { return new(ChooseCurAvatarTeamReq) },
		func() ProtoMessage { return new(ChooseCurAvatarTeamRsp) },
		func() ProtoMessage { return new(ChangeAvatarReq) },
		func() ProtoMessage { return new(ChangeAvatarRsp) },
		func() ProtoMessage { return new(AvatarPromoteReq) },
		func() ProtoMessage { return new(AvatarPromoteRsp) },
		func() ProtoMessage { return new(SpringUseReq) },
		func() ProtoMessage { return new(SpringUseRsp) },
		func() ProtoMessage { return new(RefreshBackgroundAvatarReq) },
		func() ProtoMessage { return new(RefreshBackgroundAvatarRsp) },
		func() ProtoMessage { return new(AvatarTeamUpdateNotify) },
		func() ProtoMessage { return new(AvatarDataNotify) },
		func() ProtoMessage { return new(AvatarUpgradeReq) },
		func() ProtoMessage { return new(AvatarUpgradeRsp) },
		func() ProtoMessage { return new(AvatarDieAnimationEndReq) },
		func() ProtoMessage { return new(AvatarDieAnimationEndRsp) },
		func() ProtoMessage { return new(AvatarChangeElementTypeReq) },
		func() ProtoMessage { return new(AvatarChangeElementTypeRsp) },
		func() ProtoMessage { return new(AvatarFetterDataNotify) },
		func() ProtoMessage { return new(AvatarExpeditionDataNotify) },
		func() ProtoMessage { return new(AvatarExpeditionAllDataReq) },
		func() ProtoMessage { return new(AvatarExpeditionAllDataRsp) },
		func() ProtoMessage { return new(AvatarExpeditionStartReq) },
		func() ProtoMessage { return new(AvatarExpeditionStartRsp) },
		func() ProtoMessage { return new(AvatarExpeditionCallBackReq) },
		func() ProtoMessage { return new(AvatarExpeditionCallBackRsp) },
		func() ProtoMessage { return new(AvatarExpeditionGetRewardReq) },
		func() ProtoMessage { return new(AvatarExpeditionGetRewardRsp) },
		func() ProtoMessage { return new(ChangeMpTeamAvatarReq) },
		func() ProtoMessage { return new(ChangeMpTeamAvatarRsp) },
		func() ProtoMessage { return new(ChangeTeamNameReq) },
		func() ProtoMessage { return new(ChangeTeamNameRsp) },
		func() ProtoMessage { return new(SceneTeamUpdateNotify) },
		func() ProtoMessage { return new(FocusAvatarReq) },
		func() ProtoMessage { return new(FocusAvatarRsp) },
		func() ProtoMessage { return new(AvatarSatiationDataNotify) },
		func() ProtoMessage { return new(AvatarWearFlycloakReq) },
		func() ProtoMessage { return new(AvatarWearFlycloakRsp) },
		func() ProtoMessage { return new(AvatarFlycloakChangeNotify) },
		func() ProtoMessage { return new(AvatarGainFlycloakNotify) },
		func() ProtoMessage { return new(AvatarEquipAffixStartNotify) },
		func() ProtoMessage { return new(AvatarFetterLevelRewardReq) },
		func() ProtoMessage { return new(AvatarFetterLevelRewardRsp) },
		func() ProtoMessage { return new(AddNoGachaAvatarCardNotify) },
		func() ProtoMessage { return new(AvatarPromoteGetRewardReq) },
		func() ProtoMessage { return new(AvatarPromoteGetRewardRsp) },
		func() ProtoMessage { return new(AvatarChangeCostumeReq) },
		func() ProtoMessage { return new(AvatarChangeCostumeRsp) },
		func() ProtoMessage { return new(AvatarChangeCostumeNotify) },
		func() ProtoMessage { return new(AvatarGainCostumeNotify) },
		func() ProtoMessage { return new(AvatarChangeAnimHashReq) },
		func() ProtoMessage { return new(AvatarChangeAnimHashRsp) },
		func() ProtoMessage { return new(PersistentDungeonSwitchAvatarReq) },
		func() ProtoMessage { return new(PersistentDungeonSwitchAvatarRsp) },
		func() ProtoMessage { return new(AddBackupAvatarTeamReq) },
		func() ProtoMessage { return new(AddBackupAvatarTeamRsp) },
		func() ProtoMessage { return new(DelBackupAvatarTeamReq) },
		func() ProtoMessage { return new(DelBackupAvatarTeamRsp) },
		func() ProtoMessage { return new(AvatarTeamAllDataNotify) },
		func() ProtoMessage { return new(AvatarRenameInfoNotify) },
		func() ProtoMessage { return new(ItemRenameAvatarReq) },
		func() ProtoMessage { return new(ItemRenameAvatarRsp) },
	)
}

const (
	ProtoCommandAvatarAddNotify                  ProtoCommand = 1769
	ProtoCommandAvatarDelNotify                  ProtoCommand = 1773
	ProtoCommandSetUpAvatarTeamReq               ProtoCommand = 1690
	ProtoCommandSetUpAvatarTeamRsp               ProtoCommand = 1646
	ProtoCommandChooseCurAvatarTeamReq           ProtoCommand = 1796
	ProtoCommandChooseCurAvatarTeamRsp           ProtoCommand = 1661
	ProtoCommandChangeAvatarReq                  ProtoCommand = 1640
	ProtoCommandChangeAvatarRsp                  ProtoCommand = 1607
	ProtoCommandAvatarPromoteReq                 ProtoCommand = 1664
	ProtoCommandAvatarPromoteRsp                 ProtoCommand = 1639
	ProtoCommandSpringUseReq                     ProtoCommand = 1748
	ProtoCommandSpringUseRsp                     ProtoCommand = 1642
	ProtoCommandRefreshBackgroundAvatarReq       ProtoCommand = 1743
	ProtoCommandRefreshBackgroundAvatarRsp       ProtoCommand = 1800
	ProtoCommandAvatarTeamUpdateNotify           ProtoCommand = 1706
	ProtoCommandAvatarDataNotify                 ProtoCommand = 1633
	ProtoCommandAvatarUpgradeReq                 ProtoCommand = 1770
	ProtoCommandAvatarUpgradeRsp                 ProtoCommand = 1701
	ProtoCommandAvatarDieAnimationEndReq         ProtoCommand = 1610
	ProtoCommandAvatarDieAnimationEndRsp         ProtoCommand = 1694
	ProtoCommandAvatarChangeElementTypeReq       ProtoCommand = 1785
	ProtoCommandAvatarChangeElementTypeRsp       ProtoCommand = 1651
	ProtoCommandAvatarFetterDataNotify           ProtoCommand = 1782
	ProtoCommandAvatarExpeditionDataNotify       ProtoCommand = 1771
	ProtoCommandAvatarExpeditionAllDataReq       ProtoCommand = 1722
	ProtoCommandAvatarExpeditionAllDataRsp       ProtoCommand = 1648
	ProtoCommandAvatarExpeditionStartReq         ProtoCommand = 1715
	ProtoCommandAvatarExpeditionStartRsp         ProtoCommand = 1719
	ProtoCommandAvatarExpeditionCallBackReq      ProtoCommand = 1752
	ProtoCommandAvatarExpeditionCallBackRsp      ProtoCommand = 1726
	ProtoCommandAvatarExpeditionGetRewardReq     ProtoCommand = 1623
	ProtoCommandAvatarExpeditionGetRewardRsp     ProtoCommand = 1784
	ProtoCommandChangeMpTeamAvatarReq            ProtoCommand = 1708
	ProtoCommandChangeMpTeamAvatarRsp            ProtoCommand = 1753
	ProtoCommandChangeTeamNameReq                ProtoCommand = 1603
	ProtoCommandChangeTeamNameRsp                ProtoCommand = 1666
	ProtoCommandSceneTeamUpdateNotify            ProtoCommand = 1775
	ProtoCommandFocusAvatarReq                   ProtoCommand = 1654
	ProtoCommandFocusAvatarRsp                   ProtoCommand = 1681
	ProtoCommandAvatarSatiationDataNotify        ProtoCommand = 1693
	ProtoCommandAvatarWearFlycloakReq            ProtoCommand = 1737
	ProtoCommandAvatarWearFlycloakRsp            ProtoCommand = 1698
	ProtoCommandAvatarFlycloakChangeNotify       ProtoCommand = 1643
	ProtoCommandAvatarGainFlycloakNotify         ProtoCommand = 1656
	ProtoCommandAvatarEquipAffixStartNotify      ProtoCommand = 1662
	ProtoCommandAvatarFetterLevelRewardReq       ProtoCommand = 1653
	ProtoCommandAvatarFetterLevelRewardRsp       ProtoCommand = 1606
	ProtoCommandAddNoGachaAvatarCardNotify       ProtoCommand = 1655
	ProtoCommandAvatarPromoteGetRewardReq        ProtoCommand = 1696
	ProtoCommandAvatarPromoteGetRewardRsp        ProtoCommand = 1683
	ProtoCommandAvatarChangeCostumeReq           ProtoCommand = 1778
	ProtoCommandAvatarChangeCostumeRsp           ProtoCommand = 1645
	ProtoCommandAvatarChangeCostumeNotify        ProtoCommand = 1644
	ProtoCommandAvatarGainCostumeNotify          ProtoCommand = 1677
	ProtoCommandAvatarChangeAnimHashReq          ProtoCommand = 1711
	ProtoCommandAvatarChangeAnimHashRsp          ProtoCommand = 1647
	ProtoCommandPersistentDungeonSwitchAvatarReq ProtoCommand = 1684
	ProtoCommandPersistentDungeonSwitchAvatarRsp ProtoCommand = 1768
	ProtoCommandAddBackupAvatarTeamReq           ProtoCommand = 1687
	ProtoCommandAddBackupAvatarTeamRsp           ProtoCommand = 1735
	ProtoCommandDelBackupAvatarTeamReq           ProtoCommand = 1731
	ProtoCommandDelBackupAvatarTeamRsp           ProtoCommand = 1729
	ProtoCommandAvatarTeamAllDataNotify          ProtoCommand = 1749
	ProtoCommandAvatarRenameInfoNotify           ProtoCommand = 1680
	ProtoCommandItemRenameAvatarReq              ProtoCommand = 1750
	ProtoCommandItemRenameAvatarRsp              ProtoCommand = 1635
)

func (*AvatarAddNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarAddNotify }
func (*AvatarAddNotify) ProtoMessageType() ProtoMessageType { return "AvatarAddNotify" }

func (*AvatarDelNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarDelNotify }
func (*AvatarDelNotify) ProtoMessageType() ProtoMessageType { return "AvatarDelNotify" }

func (*SetUpAvatarTeamReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetUpAvatarTeamReq }
func (*SetUpAvatarTeamReq) ProtoMessageType() ProtoMessageType { return "SetUpAvatarTeamReq" }

func (*SetUpAvatarTeamRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetUpAvatarTeamRsp }
func (*SetUpAvatarTeamRsp) ProtoMessageType() ProtoMessageType { return "SetUpAvatarTeamRsp" }

func (*ChooseCurAvatarTeamReq) ProtoCommand() ProtoCommand         { return ProtoCommandChooseCurAvatarTeamReq }
func (*ChooseCurAvatarTeamReq) ProtoMessageType() ProtoMessageType { return "ChooseCurAvatarTeamReq" }

func (*ChooseCurAvatarTeamRsp) ProtoCommand() ProtoCommand         { return ProtoCommandChooseCurAvatarTeamRsp }
func (*ChooseCurAvatarTeamRsp) ProtoMessageType() ProtoMessageType { return "ChooseCurAvatarTeamRsp" }

func (*ChangeAvatarReq) ProtoCommand() ProtoCommand         { return ProtoCommandChangeAvatarReq }
func (*ChangeAvatarReq) ProtoMessageType() ProtoMessageType { return "ChangeAvatarReq" }

func (*ChangeAvatarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandChangeAvatarRsp }
func (*ChangeAvatarRsp) ProtoMessageType() ProtoMessageType { return "ChangeAvatarRsp" }

func (*AvatarPromoteReq) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarPromoteReq }
func (*AvatarPromoteReq) ProtoMessageType() ProtoMessageType { return "AvatarPromoteReq" }

func (*AvatarPromoteRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarPromoteRsp }
func (*AvatarPromoteRsp) ProtoMessageType() ProtoMessageType { return "AvatarPromoteRsp" }

func (*SpringUseReq) ProtoCommand() ProtoCommand         { return ProtoCommandSpringUseReq }
func (*SpringUseReq) ProtoMessageType() ProtoMessageType { return "SpringUseReq" }

func (*SpringUseRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSpringUseRsp }
func (*SpringUseRsp) ProtoMessageType() ProtoMessageType { return "SpringUseRsp" }

func (*RefreshBackgroundAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshBackgroundAvatarReq
}
func (*RefreshBackgroundAvatarReq) ProtoMessageType() ProtoMessageType {
	return "RefreshBackgroundAvatarReq"
}

func (*RefreshBackgroundAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandRefreshBackgroundAvatarRsp
}
func (*RefreshBackgroundAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "RefreshBackgroundAvatarRsp"
}

func (*AvatarTeamUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarTeamUpdateNotify }
func (*AvatarTeamUpdateNotify) ProtoMessageType() ProtoMessageType { return "AvatarTeamUpdateNotify" }

func (*AvatarDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarDataNotify }
func (*AvatarDataNotify) ProtoMessageType() ProtoMessageType { return "AvatarDataNotify" }

func (*AvatarUpgradeReq) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarUpgradeReq }
func (*AvatarUpgradeReq) ProtoMessageType() ProtoMessageType { return "AvatarUpgradeReq" }

func (*AvatarUpgradeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarUpgradeRsp }
func (*AvatarUpgradeRsp) ProtoMessageType() ProtoMessageType { return "AvatarUpgradeRsp" }

func (*AvatarDieAnimationEndReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarDieAnimationEndReq
}
func (*AvatarDieAnimationEndReq) ProtoMessageType() ProtoMessageType {
	return "AvatarDieAnimationEndReq"
}

func (*AvatarDieAnimationEndRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarDieAnimationEndRsp
}
func (*AvatarDieAnimationEndRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarDieAnimationEndRsp"
}

func (*AvatarChangeElementTypeReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarChangeElementTypeReq
}
func (*AvatarChangeElementTypeReq) ProtoMessageType() ProtoMessageType {
	return "AvatarChangeElementTypeReq"
}

func (*AvatarChangeElementTypeRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarChangeElementTypeRsp
}
func (*AvatarChangeElementTypeRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarChangeElementTypeRsp"
}

func (*AvatarFetterDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarFetterDataNotify }
func (*AvatarFetterDataNotify) ProtoMessageType() ProtoMessageType { return "AvatarFetterDataNotify" }

func (*AvatarExpeditionDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionDataNotify
}
func (*AvatarExpeditionDataNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionDataNotify"
}

func (*AvatarExpeditionAllDataReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionAllDataReq
}
func (*AvatarExpeditionAllDataReq) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionAllDataReq"
}

func (*AvatarExpeditionAllDataRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionAllDataRsp
}
func (*AvatarExpeditionAllDataRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionAllDataRsp"
}

func (*AvatarExpeditionStartReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionStartReq
}
func (*AvatarExpeditionStartReq) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionStartReq"
}

func (*AvatarExpeditionStartRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionStartRsp
}
func (*AvatarExpeditionStartRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionStartRsp"
}

func (*AvatarExpeditionCallBackReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionCallBackReq
}
func (*AvatarExpeditionCallBackReq) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionCallBackReq"
}

func (*AvatarExpeditionCallBackRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionCallBackRsp
}
func (*AvatarExpeditionCallBackRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionCallBackRsp"
}

func (*AvatarExpeditionGetRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionGetRewardReq
}
func (*AvatarExpeditionGetRewardReq) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionGetRewardReq"
}

func (*AvatarExpeditionGetRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarExpeditionGetRewardRsp
}
func (*AvatarExpeditionGetRewardRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarExpeditionGetRewardRsp"
}

func (*ChangeMpTeamAvatarReq) ProtoCommand() ProtoCommand         { return ProtoCommandChangeMpTeamAvatarReq }
func (*ChangeMpTeamAvatarReq) ProtoMessageType() ProtoMessageType { return "ChangeMpTeamAvatarReq" }

func (*ChangeMpTeamAvatarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandChangeMpTeamAvatarRsp }
func (*ChangeMpTeamAvatarRsp) ProtoMessageType() ProtoMessageType { return "ChangeMpTeamAvatarRsp" }

func (*ChangeTeamNameReq) ProtoCommand() ProtoCommand         { return ProtoCommandChangeTeamNameReq }
func (*ChangeTeamNameReq) ProtoMessageType() ProtoMessageType { return "ChangeTeamNameReq" }

func (*ChangeTeamNameRsp) ProtoCommand() ProtoCommand         { return ProtoCommandChangeTeamNameRsp }
func (*ChangeTeamNameRsp) ProtoMessageType() ProtoMessageType { return "ChangeTeamNameRsp" }

func (*SceneTeamUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSceneTeamUpdateNotify }
func (*SceneTeamUpdateNotify) ProtoMessageType() ProtoMessageType { return "SceneTeamUpdateNotify" }

func (*FocusAvatarReq) ProtoCommand() ProtoCommand         { return ProtoCommandFocusAvatarReq }
func (*FocusAvatarReq) ProtoMessageType() ProtoMessageType { return "FocusAvatarReq" }

func (*FocusAvatarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandFocusAvatarRsp }
func (*FocusAvatarRsp) ProtoMessageType() ProtoMessageType { return "FocusAvatarRsp" }

func (*AvatarSatiationDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarSatiationDataNotify
}
func (*AvatarSatiationDataNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarSatiationDataNotify"
}

func (*AvatarWearFlycloakReq) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarWearFlycloakReq }
func (*AvatarWearFlycloakReq) ProtoMessageType() ProtoMessageType { return "AvatarWearFlycloakReq" }

func (*AvatarWearFlycloakRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarWearFlycloakRsp }
func (*AvatarWearFlycloakRsp) ProtoMessageType() ProtoMessageType { return "AvatarWearFlycloakRsp" }

func (*AvatarFlycloakChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarFlycloakChangeNotify
}
func (*AvatarFlycloakChangeNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarFlycloakChangeNotify"
}

func (*AvatarGainFlycloakNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarGainFlycloakNotify
}
func (*AvatarGainFlycloakNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarGainFlycloakNotify"
}

func (*AvatarEquipAffixStartNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarEquipAffixStartNotify
}
func (*AvatarEquipAffixStartNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarEquipAffixStartNotify"
}

func (*AvatarFetterLevelRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarFetterLevelRewardReq
}
func (*AvatarFetterLevelRewardReq) ProtoMessageType() ProtoMessageType {
	return "AvatarFetterLevelRewardReq"
}

func (*AvatarFetterLevelRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarFetterLevelRewardRsp
}
func (*AvatarFetterLevelRewardRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarFetterLevelRewardRsp"
}

func (*AddNoGachaAvatarCardNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAddNoGachaAvatarCardNotify
}
func (*AddNoGachaAvatarCardNotify) ProtoMessageType() ProtoMessageType {
	return "AddNoGachaAvatarCardNotify"
}

func (*AvatarPromoteGetRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarPromoteGetRewardReq
}
func (*AvatarPromoteGetRewardReq) ProtoMessageType() ProtoMessageType {
	return "AvatarPromoteGetRewardReq"
}

func (*AvatarPromoteGetRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarPromoteGetRewardRsp
}
func (*AvatarPromoteGetRewardRsp) ProtoMessageType() ProtoMessageType {
	return "AvatarPromoteGetRewardRsp"
}

func (*AvatarChangeCostumeReq) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarChangeCostumeReq }
func (*AvatarChangeCostumeReq) ProtoMessageType() ProtoMessageType { return "AvatarChangeCostumeReq" }

func (*AvatarChangeCostumeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarChangeCostumeRsp }
func (*AvatarChangeCostumeRsp) ProtoMessageType() ProtoMessageType { return "AvatarChangeCostumeRsp" }

func (*AvatarChangeCostumeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarChangeCostumeNotify
}
func (*AvatarChangeCostumeNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarChangeCostumeNotify"
}

func (*AvatarGainCostumeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarGainCostumeNotify
}
func (*AvatarGainCostumeNotify) ProtoMessageType() ProtoMessageType { return "AvatarGainCostumeNotify" }

func (*AvatarChangeAnimHashReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarChangeAnimHashReq
}
func (*AvatarChangeAnimHashReq) ProtoMessageType() ProtoMessageType { return "AvatarChangeAnimHashReq" }

func (*AvatarChangeAnimHashRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarChangeAnimHashRsp
}
func (*AvatarChangeAnimHashRsp) ProtoMessageType() ProtoMessageType { return "AvatarChangeAnimHashRsp" }

func (*PersistentDungeonSwitchAvatarReq) ProtoCommand() ProtoCommand {
	return ProtoCommandPersistentDungeonSwitchAvatarReq
}
func (*PersistentDungeonSwitchAvatarReq) ProtoMessageType() ProtoMessageType {
	return "PersistentDungeonSwitchAvatarReq"
}

func (*PersistentDungeonSwitchAvatarRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandPersistentDungeonSwitchAvatarRsp
}
func (*PersistentDungeonSwitchAvatarRsp) ProtoMessageType() ProtoMessageType {
	return "PersistentDungeonSwitchAvatarRsp"
}

func (*AddBackupAvatarTeamReq) ProtoCommand() ProtoCommand         { return ProtoCommandAddBackupAvatarTeamReq }
func (*AddBackupAvatarTeamReq) ProtoMessageType() ProtoMessageType { return "AddBackupAvatarTeamReq" }

func (*AddBackupAvatarTeamRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAddBackupAvatarTeamRsp }
func (*AddBackupAvatarTeamRsp) ProtoMessageType() ProtoMessageType { return "AddBackupAvatarTeamRsp" }

func (*DelBackupAvatarTeamReq) ProtoCommand() ProtoCommand         { return ProtoCommandDelBackupAvatarTeamReq }
func (*DelBackupAvatarTeamReq) ProtoMessageType() ProtoMessageType { return "DelBackupAvatarTeamReq" }

func (*DelBackupAvatarTeamRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDelBackupAvatarTeamRsp }
func (*DelBackupAvatarTeamRsp) ProtoMessageType() ProtoMessageType { return "DelBackupAvatarTeamRsp" }

func (*AvatarTeamAllDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarTeamAllDataNotify
}
func (*AvatarTeamAllDataNotify) ProtoMessageType() ProtoMessageType { return "AvatarTeamAllDataNotify" }

func (*AvatarRenameInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarRenameInfoNotify }
func (*AvatarRenameInfoNotify) ProtoMessageType() ProtoMessageType { return "AvatarRenameInfoNotify" }

func (*ItemRenameAvatarReq) ProtoCommand() ProtoCommand         { return ProtoCommandItemRenameAvatarReq }
func (*ItemRenameAvatarReq) ProtoMessageType() ProtoMessageType { return "ItemRenameAvatarReq" }

func (*ItemRenameAvatarRsp) ProtoCommand() ProtoCommand         { return ProtoCommandItemRenameAvatarRsp }
func (*ItemRenameAvatarRsp) ProtoMessageType() ProtoMessageType { return "ItemRenameAvatarRsp" }
