// Code generated by "stringer -type=OperationType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opUnused-0]
	_ = x[opPing-1]
	_ = x[opJoin-2]
	_ = x[opCreateAccount-3]
	_ = x[opLogin-4]
	_ = x[opSendCrashLog-5]
	_ = x[opSendTraceRoute-6]
	_ = x[opSendVfxStats-7]
	_ = x[opSendGamePingInfo-8]
	_ = x[opCreateCharacter-9]
	_ = x[opDeleteCharacter-10]
	_ = x[opSelectCharacter-11]
	_ = x[opRedeemKeycode-12]
	_ = x[opGetGameServerByCluster-13]
	_ = x[opGetActiveSubscription-14]
	_ = x[opGetShopPurchaseUrl-15]
	_ = x[opGetBuyTrialDetails-16]
	_ = x[opGetReferralSeasonDetails-17]
	_ = x[opGetReferralLink-18]
	_ = x[opGetAvailableTrialKeys-19]
	_ = x[opGetShopTilesForCategory-20]
	_ = x[opMove-21]
	_ = x[opCastStart-22]
	_ = x[opCastCancel-23]
	_ = x[opTerminateToggleSpell-24]
	_ = x[opChannelingCancel-25]
	_ = x[opAttackBuildingStart-26]
	_ = x[opInventoryDestroyItem-27]
	_ = x[opInventoryMoveItem-28]
	_ = x[opInventoryRecoverItem-29]
	_ = x[opInventoryRecoverAllItems-30]
	_ = x[opInventorySplitStack-31]
	_ = x[opInventorySplitStackInto-32]
	_ = x[opGetClusterData-33]
	_ = x[opChangeCluster-34]
	_ = x[opConsoleCommand-35]
	_ = x[opChatMessage-36]
	_ = x[opReportClientError-37]
	_ = x[opRegisterToObject-38]
	_ = x[opUnRegisterFromObject-39]
	_ = x[opCraftBuildingChangeSettings-40]
	_ = x[opCraftBuildingTakeMoney-41]
	_ = x[opRepairBuildingChangeSettings-42]
	_ = x[opRepairBuildingTakeMoney-43]
	_ = x[opActionBuildingChangeSettings-44]
	_ = x[opHarvestStart-45]
	_ = x[opHarvestCancel-46]
	_ = x[opTakeSilver-47]
	_ = x[opActionOnBuildingStart-48]
	_ = x[opActionOnBuildingCancel-49]
	_ = x[opItemRerollQualityStart-50]
	_ = x[opItemRerollQualityCancel-51]
	_ = x[opInstallResourceStart-52]
	_ = x[opInstallResourceCancel-53]
	_ = x[opInstallSilver-54]
	_ = x[opBuildingFillNutrition-55]
	_ = x[opBuildingChangeRenovationState-56]
	_ = x[opBuildingBuySkin-57]
	_ = x[opBuildingClaim-58]
	_ = x[opBuildingGiveup-59]
	_ = x[opBuildingNutritionSilverStorageDeposit-60]
	_ = x[opBuildingNutritionSilverStorageWithdraw-61]
	_ = x[opBuildingNutritionSilverRewardSet-62]
	_ = x[opConstructionSiteCreate-63]
	_ = x[opPlaceableObjectPlace-64]
	_ = x[opPlaceableObjectPlaceCancel-65]
	_ = x[opPlaceableObjectPickup-66]
	_ = x[opFurnitureObjectUse-67]
	_ = x[opFarmableHarvest-68]
	_ = x[opFarmableFinishGrownItem-69]
	_ = x[opFarmableDestroy-70]
	_ = x[opFarmableGetProduct-71]
	_ = x[opTearDownConstructionSite-72]
	_ = x[opAuctionCreateRequest-73]
	_ = x[opAuctionCreateOffer-74]
	_ = x[opAuctionGetOffers-75]
	_ = x[opAuctionGetRequests-76]
	_ = x[opAuctionBuyOffer-77]
	_ = x[opAuctionAbortAuction-78]
	_ = x[opAuctionModifyAuction-79]
	_ = x[opAuctionAbortOffer-80]
	_ = x[opAuctionAbortRequest-81]
	_ = x[opAuctionSellRequest-82]
	_ = x[opAuctionGetFinishedAuctions-83]
	_ = x[opAuctionGetFinishedAuctionsCount-84]
	_ = x[opAuctionFetchAuction-85]
	_ = x[opAuctionGetMyOpenOffers-86]
	_ = x[opAuctionGetMyOpenRequests-87]
	_ = x[opAuctionGetMyOpenAuctions-88]
	_ = x[opUnknown90-89]
	_ = x[opAuctionGetItemAverageStats-90]
	_ = x[opAuctionGetItemAverageValue-91]
	_ = x[opContainerOpen-92]
	_ = x[opContainerClose-93]
	_ = x[opContainerManageSubContainer-94]
	_ = x[opRespawn-95]
	_ = x[opSuicide-96]
	_ = x[opJoinGuild-97]
	_ = x[opLeaveGuild-98]
	_ = x[opCreateGuild-99]
	_ = x[opInviteToGuild-100]
	_ = x[opDeclineGuildInvitation-101]
	_ = x[opKickFromGuild-102]
	_ = x[opDuellingChallengePlayer-103]
	_ = x[opDuellingAcceptChallenge-104]
	_ = x[opDuellingDenyChallenge-105]
	_ = x[opChangeClusterTax-106]
	_ = x[opClaimTerritory-107]
	_ = x[opGiveUpTerritory-108]
	_ = x[opChangeTerritoryAccessRights-109]
	_ = x[opGetMonolithInfo-110]
	_ = x[opGetClaimInfo-111]
	_ = x[opGetAttackInfo-112]
	_ = x[opGetTerritorySeasonPoints-113]
	_ = x[opGetAttackSchedule-114]
	_ = x[opScheduleAttack-115]
	_ = x[opGetMatches-116]
	_ = x[opGetMatchDetails-117]
	_ = x[opJoinMatch-118]
	_ = x[opLeaveMatch-119]
	_ = x[opChangeChatSettings-120]
	_ = x[opLogoutStart-121]
	_ = x[opLogoutCancel-122]
	_ = x[opClaimOrbStart-123]
	_ = x[opClaimOrbCancel-124]
	_ = x[opMatchLootChestOpeningStart-125]
	_ = x[opMatchLootChestOpeningCancel-126]
	_ = x[opDepositToGuildAccount-127]
	_ = x[opWithdrawalFromAccount-128]
	_ = x[opChangeGuildPayUpkeepFlag-129]
	_ = x[opChangeGuildTax-130]
	_ = x[opGetMyTerritories-131]
	_ = x[opMorganaCommand-132]
	_ = x[opGetServerInfo-133]
	_ = x[opInviteMercenaryToMatch-134]
	_ = x[opSubscribeToCluster-135]
	_ = x[opAnswerMercenaryInvitation-136]
	_ = x[opGetCharacterEquipment-137]
	_ = x[opGetCharacterSteamAchievements-138]
	_ = x[opGetCharacterStats-139]
	_ = x[opGetKillHistoryDetails-140]
	_ = x[opLearnMasteryLevel-141]
	_ = x[opReSpecAchievement-142]
	_ = x[opChangeAvatar-143]
	_ = x[opGetRankings-144]
	_ = x[opGetRank-145]
	_ = x[opGetGvgSeasonRankings-146]
	_ = x[opGetGvgSeasonRank-147]
	_ = x[opGetGvgSeasonHistoryRankings-148]
	_ = x[opGetGvgSeasonGuildMemberHistory-149]
	_ = x[opKickFromGvGMatch-150]
	_ = x[opGetChestLogs-151]
	_ = x[opGetAccessRightLogs-152]
	_ = x[opGetGuildAccountLogs-153]
	_ = x[opGetGuildAccountLogsLargeAmount-154]
	_ = x[opInviteToPlayerTrade-155]
	_ = x[opPlayerTradeCancel-156]
	_ = x[opPlayerTradeInvitationAccept-157]
	_ = x[opPlayerTradeAddItem-158]
	_ = x[opPlayerTradeRemoveItem-159]
	_ = x[opPlayerTradeAcceptTrade-160]
	_ = x[opPlayerTradeSetSilverOrGold-161]
	_ = x[opSendMiniMapPing-162]
	_ = x[opStuck-163]
	_ = x[opBuyRealEstate-164]
	_ = x[opClaimRealEstate-165]
	_ = x[opGiveUpRealEstate-166]
	_ = x[opChangeRealEstateOutline-167]
	_ = x[opGetMailInfos-168]
	_ = x[opGetMailCount-169]
	_ = x[opReadMail-170]
	_ = x[opSendNewMail-171]
	_ = x[opDeleteMail-172]
	_ = x[opMarkMailUnread-173]
	_ = x[opClaimAttachmentFromMail-174]
	_ = x[opUpdateLfgInfo-175]
	_ = x[opGetLfgInfos-176]
	_ = x[opGetMyGuildLfgInfo-177]
	_ = x[opGetLfgDescriptionText-178]
	_ = x[opLfgApplyToGuild-179]
	_ = x[opAnswerLfgGuildApplication-180]
	_ = x[opRegisterChatPeer-181]
	_ = x[opSendChatMessage-182]
	_ = x[opJoinChatChannel-183]
	_ = x[opLeaveChatChannel-184]
	_ = x[opSendWhisperMessage-185]
	_ = x[opSay-186]
	_ = x[opPlayEmote-187]
	_ = x[opStopEmote-188]
	_ = x[opGetClusterMapInfo-189]
	_ = x[opAccessRightsChangeSettings-190]
	_ = x[opMount-191]
	_ = x[opMountCancel-192]
	_ = x[opBuyJourney-193]
	_ = x[opSetSaleStatusForEstate-194]
	_ = x[opResolveGuildOrPlayerName-195]
	_ = x[opGetRespawnInfos-196]
	_ = x[opMakeHome-197]
	_ = x[opLeaveHome-198]
	_ = x[opResurrectionReply-199]
	_ = x[opAllianceCreate-200]
	_ = x[opAllianceDisband-201]
	_ = x[opAllianceGetMemberInfos-202]
	_ = x[opAllianceInvite-203]
	_ = x[opAllianceAnswerInvitation-204]
	_ = x[opAllianceCancelInvitation-205]
	_ = x[opAllianceKickGuild-206]
	_ = x[opAllianceLeave-207]
	_ = x[opAllianceChangeGoldPaymentFlag-208]
	_ = x[opAllianceGetDetailInfo-209]
	_ = x[opGetIslandInfos-210]
	_ = x[opAbandonMyIsland-211]
	_ = x[opBuyMyIsland-212]
	_ = x[opBuyGuildIsland-213]
	_ = x[opAbandonGuildIsland-214]
	_ = x[opUpgradeMyIsland-215]
	_ = x[opUpgradeGuildIsland-216]
	_ = x[opMoveMyIsland-217]
	_ = x[opMoveGuildIsland-218]
	_ = x[opTerritoryFillNutrition-219]
	_ = x[opTeleportBack-220]
	_ = x[opPartyInvitePlayer-221]
	_ = x[opPartyAnswerInvitation-222]
	_ = x[opPartyLeave-223]
	_ = x[opPartyKickPlayer-224]
	_ = x[opPartyMakeLeader-225]
	_ = x[opPartyChangeLootSetting-226]
	_ = x[opPartyMarkObject-227]
	_ = x[opPartySetRole-228]
	_ = x[opGetGuildMOTD-229]
	_ = x[opSetGuildMOTD-230]
	_ = x[opExitEnterStart-231]
	_ = x[opExitEnterCancel-232]
	_ = x[opQuestGiverRequest-233]
	_ = x[opGoldMarketGetBuyOffer-234]
	_ = x[opGoldMarketGetBuyOfferFromSilver-235]
	_ = x[opGoldMarketGetSellOffer-236]
	_ = x[opGoldMarketGetSellOfferFromSilver-237]
	_ = x[opGoldMarketBuyGold-238]
	_ = x[opGoldMarketSellGold-239]
	_ = x[opGoldMarketCreateSellOrder-240]
	_ = x[opGoldMarketCreateBuyOrder-241]
	_ = x[opGoldMarketGetInfos-242]
	_ = x[opGoldMarketCancelOrder-243]
	_ = x[opUnknown244-244]
	_ = x[opUnknown245-245]
	_ = x[opUnknown246-246]
	_ = x[opUnknown247-247]
	_ = x[opUnknown248-248]
	_ = x[opUnknown249-249]
	_ = x[opUnknown250-250]
	_ = x[opUnknown251-251]
	_ = x[opGoldMarketGetAverageInfo-252]
	_ = x[opSiegeCampClaimStart-253]
	_ = x[opSiegeCampClaimCancel-254]
	_ = x[opTreasureChestUsingStart-255]
	_ = x[opTreasureChestUsingCancel-256]
	_ = x[opUseLootChest-257]
	_ = x[opUseShrine-258]
	_ = x[opLaborerStartJob-259]
	_ = x[opLaborerTakeJobLoot-260]
	_ = x[opLaborerDismiss-261]
	_ = x[opLaborerMove-262]
	_ = x[opLaborerBuyItem-263]
	_ = x[opLaborerUpgrade-264]
	_ = x[opBuyPremium-265]
	_ = x[opBuyTrial-266]
	_ = x[opRealEstateGetAuctionData-267]
	_ = x[opRealEstateBidOnAuction-268]
	_ = x[opGetSiegeCampCooldown-269]
	_ = x[opFriendInvite-270]
	_ = x[opFriendAnswerInvitation-271]
	_ = x[opFriendCancelnvitation-272]
	_ = x[opFriendRemove-273]
	_ = x[opInventoryStack-274]
	_ = x[opInventorySort-275]
	_ = x[opEquipmentItemChangeSpell-276]
	_ = x[opExpeditionRegister-277]
	_ = x[opExpeditionRegisterCancel-278]
	_ = x[opJoinExpedition-279]
	_ = x[opDeclineExpeditionInvitation-280]
	_ = x[opVoteStart-281]
	_ = x[opVoteDoVote-282]
	_ = x[opRatingDoRate-283]
	_ = x[opEnteringExpeditionStart-284]
	_ = x[opEnteringExpeditionCancel-285]
	_ = x[opActivateExpeditionCheckPoint-286]
	_ = x[opArenaRegister-287]
	_ = x[opArenaRegisterCancel-288]
	_ = x[opArenaLeave-289]
	_ = x[opJoinArenaMatch-290]
	_ = x[opDeclineArenaInvitation-291]
	_ = x[opEnteringArenaStart-292]
	_ = x[opEnteringArenaCancel-293]
	_ = x[opArenaCustomMatch-294]
	_ = x[opArenaCustomMatchCreate-295]
	_ = x[opUpdateCharacterStatement-296]
	_ = x[opBoostFarmable-297]
	_ = x[opGetStrikeHistory-298]
	_ = x[opUseFunction-299]
	_ = x[opUsePortalEntrance-300]
	_ = x[opResetPortalBinding-301]
	_ = x[opQueryPortalBinding-302]
	_ = x[opClaimPaymentTransaction-303]
	_ = x[opChangeUseFlag-304]
	_ = x[opClientPerformanceStats-305]
	_ = x[opExtendedHardwareStats-306]
	_ = x[opClientLowMemoryWarning-307]
	_ = x[opTerritoryClaimStart-308]
	_ = x[opTerritoryClaimCancel-309]
	_ = x[opRequestAppStoreProducts-310]
	_ = x[opVerifyProductPurchase-311]
	_ = x[opQueryGuildPlayerStats-312]
	_ = x[opQueryAllianceGuildStats-313]
	_ = x[opTrackAchievements-314]
	_ = x[opSetAchievementsAutoLearn-315]
	_ = x[opDepositItemToGuildCurrency-316]
	_ = x[opWithdrawalItemFromGuildCurrency-317]
	_ = x[opAuctionSellSpecificItemRequest-318]
	_ = x[opFishingStart-319]
	_ = x[opFishingCasting-320]
	_ = x[opFishingCast-321]
	_ = x[opFishingCatch-322]
	_ = x[opFishingPull-323]
	_ = x[opFishingGiveLine-324]
	_ = x[opFishingFinish-325]
	_ = x[opFishingCancel-326]
	_ = x[opCreateGuildAccessTag-327]
	_ = x[opDeleteGuildAccessTag-328]
	_ = x[opRenameGuildAccessTag-329]
	_ = x[opFlagGuildAccessTagGuildPermission-330]
	_ = x[opAssignGuildAccessTag-331]
	_ = x[opRemoveGuildAccessTagFromPlayer-332]
	_ = x[opModifyGuildAccessTagEditors-333]
	_ = x[opRequestPublicAccessTags-334]
	_ = x[opChangeAccessTagPublicFlag-335]
	_ = x[opUpdateGuildAccessTag-336]
	_ = x[opSteamStartMicrotransaction-337]
	_ = x[opSteamFinishMicrotransaction-338]
	_ = x[opSteamIdHasActiveAccount-339]
	_ = x[opCheckEmailAccountState-340]
	_ = x[opLinkAccountToSteamId-341]
	_ = x[opBuyGvgSeasonBooster-342]
	_ = x[opChangeFlaggingPrepare-343]
	_ = x[opOverCharge-344]
	_ = x[opOverChargeEnd-345]
	_ = x[opRequestTrusted-346]
	_ = x[opChangeGuildLogo-347]
	_ = x[opPartyFinderRegisterForUpdates-348]
	_ = x[opPartyFinderUnregisterForUpdates-349]
	_ = x[opPartyFinderEnlistNewPartySearch-350]
	_ = x[opPartyFinderDeletePartySearch-351]
	_ = x[opPartyFinderChangePartySearch-352]
	_ = x[opPartyFinderChangeRole-353]
	_ = x[opPartyFinderApplyForGroup-354]
	_ = x[opPartyFinderAcceptOrDeclineApplyForGroup-355]
	_ = x[opPartyFinderGetEquipmentSnapshot-356]
	_ = x[opPartyFinderRegisterApplicants-357]
	_ = x[opPartyFinderUnregisterApplicants-358]
	_ = x[opPartyFinderFulltextSearch-359]
	_ = x[opPartyFinderRequestEquipmentSnapshot-360]
	_ = x[opGetPersonalSeasonTrackerData-361]
	_ = x[opUseConsumableFromInventory-362]
	_ = x[opClaimPersonalSeasonReward-363]
	_ = x[opEasyAntiCheatMessageToServer-364]
	_ = x[opSetNextTutorialState-365]
	_ = x[opAddPlayerToMuteList-366]
	_ = x[opRemovePlayerFromMuteList-367]
	_ = x[opProductShopUserEvent-368]
	_ = x[opGetVanityUnlocks-369]
	_ = x[opBuyVanityUnlocks-370]
	_ = x[opGetMountSkins-371]
	_ = x[opSetMountSkin-372]
	_ = x[opSetWardrobe-373]
	_ = x[opChangeCustomization-374]
	_ = x[opSetFavoriteIsland-375]
	_ = x[opGetGuildChallengePoints-376]
	_ = x[opTravelToHideout-377]
	_ = x[opSmartQueueJoin-378]
	_ = x[opSmartQueueLeave-379]
	_ = x[opSmartQueueSelectSpawnCluster-380]
	_ = x[opUpgradeHideout-381]
	_ = x[opInitHideoutAttackStart-382]
	_ = x[opInitHideoutAttackCancel-383]
	_ = x[opHideoutFillNutrition-384]
	_ = x[opHideoutGetInfo-385]
	_ = x[opHideoutGetOwnerInfo-386]
	_ = x[opHideoutSetTribute-387]
	_ = x[opOpenWorldAttackScheduleStart-388]
	_ = x[opOpenWorldAttackScheduleCancel-389]
	_ = x[opOpenWorldAttackConquerStart-390]
	_ = x[opOpenWorldAttackConquerCancel-391]
	_ = x[opGetOpenWorldAttackDetails-392]
	_ = x[opGetNextOpenWorldAttackScheduleTime-393]
	_ = x[opRecoverVaultFromHideout-394]
	_ = x[opGetGuildEnergyDrainInfo-395]
	_ = x[opChannelingUpdate-396]
}

const _OperationType_name = "opUnusedopPingopJoinopCreateAccountopLoginopSendCrashLogopSendTraceRouteopSendVfxStatsopSendGamePingInfoopCreateCharacteropDeleteCharacteropSelectCharacteropRedeemKeycodeopGetGameServerByClusteropGetActiveSubscriptionopGetShopPurchaseUrlopGetBuyTrialDetailsopGetReferralSeasonDetailsopGetReferralLinkopGetAvailableTrialKeysopGetShopTilesForCategoryopMoveopCastStartopCastCancelopTerminateToggleSpellopChannelingCancelopAttackBuildingStartopInventoryDestroyItemopInventoryMoveItemopInventoryRecoverItemopInventoryRecoverAllItemsopInventorySplitStackopInventorySplitStackIntoopGetClusterDataopChangeClusteropConsoleCommandopChatMessageopReportClientErroropRegisterToObjectopUnRegisterFromObjectopCraftBuildingChangeSettingsopCraftBuildingTakeMoneyopRepairBuildingChangeSettingsopRepairBuildingTakeMoneyopActionBuildingChangeSettingsopHarvestStartopHarvestCancelopTakeSilveropActionOnBuildingStartopActionOnBuildingCancelopItemRerollQualityStartopItemRerollQualityCancelopInstallResourceStartopInstallResourceCancelopInstallSilveropBuildingFillNutritionopBuildingChangeRenovationStateopBuildingBuySkinopBuildingClaimopBuildingGiveupopBuildingNutritionSilverStorageDepositopBuildingNutritionSilverStorageWithdrawopBuildingNutritionSilverRewardSetopConstructionSiteCreateopPlaceableObjectPlaceopPlaceableObjectPlaceCancelopPlaceableObjectPickupopFurnitureObjectUseopFarmableHarvestopFarmableFinishGrownItemopFarmableDestroyopFarmableGetProductopTearDownConstructionSiteopAuctionCreateRequestopAuctionCreateOfferopAuctionGetOffersopAuctionGetRequestsopAuctionBuyOfferopAuctionAbortAuctionopAuctionModifyAuctionopAuctionAbortOfferopAuctionAbortRequestopAuctionSellRequestopAuctionGetFinishedAuctionsopAuctionGetFinishedAuctionsCountopAuctionFetchAuctionopAuctionGetMyOpenOffersopAuctionGetMyOpenRequestsopAuctionGetMyOpenAuctionsopUnknown90opAuctionGetItemAverageStatsopAuctionGetItemAverageValueopContainerOpenopContainerCloseopContainerManageSubContaineropRespawnopSuicideopJoinGuildopLeaveGuildopCreateGuildopInviteToGuildopDeclineGuildInvitationopKickFromGuildopDuellingChallengePlayeropDuellingAcceptChallengeopDuellingDenyChallengeopChangeClusterTaxopClaimTerritoryopGiveUpTerritoryopChangeTerritoryAccessRightsopGetMonolithInfoopGetClaimInfoopGetAttackInfoopGetTerritorySeasonPointsopGetAttackScheduleopScheduleAttackopGetMatchesopGetMatchDetailsopJoinMatchopLeaveMatchopChangeChatSettingsopLogoutStartopLogoutCancelopClaimOrbStartopClaimOrbCancelopMatchLootChestOpeningStartopMatchLootChestOpeningCancelopDepositToGuildAccountopWithdrawalFromAccountopChangeGuildPayUpkeepFlagopChangeGuildTaxopGetMyTerritoriesopMorganaCommandopGetServerInfoopInviteMercenaryToMatchopSubscribeToClusteropAnswerMercenaryInvitationopGetCharacterEquipmentopGetCharacterSteamAchievementsopGetCharacterStatsopGetKillHistoryDetailsopLearnMasteryLevelopReSpecAchievementopChangeAvataropGetRankingsopGetRankopGetGvgSeasonRankingsopGetGvgSeasonRankopGetGvgSeasonHistoryRankingsopGetGvgSeasonGuildMemberHistoryopKickFromGvGMatchopGetChestLogsopGetAccessRightLogsopGetGuildAccountLogsopGetGuildAccountLogsLargeAmountopInviteToPlayerTradeopPlayerTradeCancelopPlayerTradeInvitationAcceptopPlayerTradeAddItemopPlayerTradeRemoveItemopPlayerTradeAcceptTradeopPlayerTradeSetSilverOrGoldopSendMiniMapPingopStuckopBuyRealEstateopClaimRealEstateopGiveUpRealEstateopChangeRealEstateOutlineopGetMailInfosopGetMailCountopReadMailopSendNewMailopDeleteMailopMarkMailUnreadopClaimAttachmentFromMailopUpdateLfgInfoopGetLfgInfosopGetMyGuildLfgInfoopGetLfgDescriptionTextopLfgApplyToGuildopAnswerLfgGuildApplicationopRegisterChatPeeropSendChatMessageopJoinChatChannelopLeaveChatChannelopSendWhisperMessageopSayopPlayEmoteopStopEmoteopGetClusterMapInfoopAccessRightsChangeSettingsopMountopMountCancelopBuyJourneyopSetSaleStatusForEstateopResolveGuildOrPlayerNameopGetRespawnInfosopMakeHomeopLeaveHomeopResurrectionReplyopAllianceCreateopAllianceDisbandopAllianceGetMemberInfosopAllianceInviteopAllianceAnswerInvitationopAllianceCancelInvitationopAllianceKickGuildopAllianceLeaveopAllianceChangeGoldPaymentFlagopAllianceGetDetailInfoopGetIslandInfosopAbandonMyIslandopBuyMyIslandopBuyGuildIslandopAbandonGuildIslandopUpgradeMyIslandopUpgradeGuildIslandopMoveMyIslandopMoveGuildIslandopTerritoryFillNutritionopTeleportBackopPartyInvitePlayeropPartyAnswerInvitationopPartyLeaveopPartyKickPlayeropPartyMakeLeaderopPartyChangeLootSettingopPartyMarkObjectopPartySetRoleopGetGuildMOTDopSetGuildMOTDopExitEnterStartopExitEnterCancelopQuestGiverRequestopGoldMarketGetBuyOfferopGoldMarketGetBuyOfferFromSilveropGoldMarketGetSellOfferopGoldMarketGetSellOfferFromSilveropGoldMarketBuyGoldopGoldMarketSellGoldopGoldMarketCreateSellOrderopGoldMarketCreateBuyOrderopGoldMarketGetInfosopGoldMarketCancelOrderopUnknown244opUnknown245opUnknown246opUnknown247opUnknown248opUnknown249opUnknown250opUnknown251opGoldMarketGetAverageInfoopSiegeCampClaimStartopSiegeCampClaimCancelopTreasureChestUsingStartopTreasureChestUsingCancelopUseLootChestopUseShrineopLaborerStartJobopLaborerTakeJobLootopLaborerDismissopLaborerMoveopLaborerBuyItemopLaborerUpgradeopBuyPremiumopBuyTrialopRealEstateGetAuctionDataopRealEstateBidOnAuctionopGetSiegeCampCooldownopFriendInviteopFriendAnswerInvitationopFriendCancelnvitationopFriendRemoveopInventoryStackopInventorySortopEquipmentItemChangeSpellopExpeditionRegisteropExpeditionRegisterCancelopJoinExpeditionopDeclineExpeditionInvitationopVoteStartopVoteDoVoteopRatingDoRateopEnteringExpeditionStartopEnteringExpeditionCancelopActivateExpeditionCheckPointopArenaRegisteropArenaRegisterCancelopArenaLeaveopJoinArenaMatchopDeclineArenaInvitationopEnteringArenaStartopEnteringArenaCancelopArenaCustomMatchopArenaCustomMatchCreateopUpdateCharacterStatementopBoostFarmableopGetStrikeHistoryopUseFunctionopUsePortalEntranceopResetPortalBindingopQueryPortalBindingopClaimPaymentTransactionopChangeUseFlagopClientPerformanceStatsopExtendedHardwareStatsopClientLowMemoryWarningopTerritoryClaimStartopTerritoryClaimCancelopRequestAppStoreProductsopVerifyProductPurchaseopQueryGuildPlayerStatsopQueryAllianceGuildStatsopTrackAchievementsopSetAchievementsAutoLearnopDepositItemToGuildCurrencyopWithdrawalItemFromGuildCurrencyopAuctionSellSpecificItemRequestopFishingStartopFishingCastingopFishingCastopFishingCatchopFishingPullopFishingGiveLineopFishingFinishopFishingCancelopCreateGuildAccessTagopDeleteGuildAccessTagopRenameGuildAccessTagopFlagGuildAccessTagGuildPermissionopAssignGuildAccessTagopRemoveGuildAccessTagFromPlayeropModifyGuildAccessTagEditorsopRequestPublicAccessTagsopChangeAccessTagPublicFlagopUpdateGuildAccessTagopSteamStartMicrotransactionopSteamFinishMicrotransactionopSteamIdHasActiveAccountopCheckEmailAccountStateopLinkAccountToSteamIdopBuyGvgSeasonBoosteropChangeFlaggingPrepareopOverChargeopOverChargeEndopRequestTrustedopChangeGuildLogoopPartyFinderRegisterForUpdatesopPartyFinderUnregisterForUpdatesopPartyFinderEnlistNewPartySearchopPartyFinderDeletePartySearchopPartyFinderChangePartySearchopPartyFinderChangeRoleopPartyFinderApplyForGroupopPartyFinderAcceptOrDeclineApplyForGroupopPartyFinderGetEquipmentSnapshotopPartyFinderRegisterApplicantsopPartyFinderUnregisterApplicantsopPartyFinderFulltextSearchopPartyFinderRequestEquipmentSnapshotopGetPersonalSeasonTrackerDataopUseConsumableFromInventoryopClaimPersonalSeasonRewardopEasyAntiCheatMessageToServeropSetNextTutorialStateopAddPlayerToMuteListopRemovePlayerFromMuteListopProductShopUserEventopGetVanityUnlocksopBuyVanityUnlocksopGetMountSkinsopSetMountSkinopSetWardrobeopChangeCustomizationopSetFavoriteIslandopGetGuildChallengePointsopTravelToHideoutopSmartQueueJoinopSmartQueueLeaveopSmartQueueSelectSpawnClusteropUpgradeHideoutopInitHideoutAttackStartopInitHideoutAttackCancelopHideoutFillNutritionopHideoutGetInfoopHideoutGetOwnerInfoopHideoutSetTributeopOpenWorldAttackScheduleStartopOpenWorldAttackScheduleCancelopOpenWorldAttackConquerStartopOpenWorldAttackConquerCancelopGetOpenWorldAttackDetailsopGetNextOpenWorldAttackScheduleTimeopRecoverVaultFromHideoutopGetGuildEnergyDrainInfoopChannelingUpdate"

var _OperationType_index = [...]uint16{0, 8, 14, 20, 35, 42, 56, 72, 86, 104, 121, 138, 155, 170, 194, 217, 237, 257, 283, 300, 323, 348, 354, 365, 377, 399, 417, 438, 460, 479, 501, 527, 548, 573, 589, 604, 620, 633, 652, 670, 692, 721, 745, 775, 800, 830, 844, 859, 871, 894, 918, 942, 967, 989, 1012, 1027, 1050, 1081, 1098, 1113, 1129, 1168, 1208, 1242, 1266, 1288, 1316, 1339, 1359, 1376, 1401, 1418, 1438, 1464, 1486, 1506, 1524, 1544, 1561, 1582, 1604, 1623, 1644, 1664, 1692, 1725, 1746, 1770, 1796, 1822, 1833, 1861, 1889, 1904, 1920, 1949, 1958, 1967, 1978, 1990, 2003, 2018, 2042, 2057, 2082, 2107, 2130, 2148, 2164, 2181, 2210, 2227, 2241, 2256, 2282, 2301, 2317, 2329, 2346, 2357, 2369, 2389, 2402, 2416, 2431, 2447, 2475, 2504, 2527, 2550, 2576, 2592, 2610, 2626, 2641, 2665, 2685, 2712, 2735, 2766, 2785, 2808, 2827, 2846, 2860, 2873, 2882, 2904, 2922, 2951, 2983, 3001, 3015, 3035, 3056, 3088, 3109, 3128, 3157, 3177, 3200, 3224, 3252, 3269, 3276, 3291, 3308, 3326, 3351, 3365, 3379, 3389, 3402, 3414, 3430, 3455, 3470, 3483, 3502, 3525, 3542, 3569, 3587, 3604, 3621, 3639, 3659, 3664, 3675, 3686, 3705, 3733, 3740, 3753, 3765, 3789, 3815, 3832, 3842, 3853, 3872, 3888, 3905, 3929, 3945, 3971, 3997, 4016, 4031, 4062, 4085, 4101, 4118, 4131, 4147, 4167, 4184, 4204, 4218, 4235, 4259, 4273, 4292, 4315, 4327, 4344, 4361, 4385, 4402, 4416, 4430, 4444, 4460, 4477, 4496, 4519, 4552, 4576, 4610, 4629, 4649, 4676, 4702, 4722, 4745, 4757, 4769, 4781, 4793, 4805, 4817, 4829, 4841, 4867, 4888, 4910, 4935, 4961, 4975, 4986, 5003, 5023, 5039, 5052, 5068, 5084, 5096, 5106, 5132, 5156, 5178, 5192, 5216, 5239, 5253, 5269, 5284, 5310, 5330, 5356, 5372, 5401, 5412, 5424, 5438, 5463, 5489, 5519, 5534, 5555, 5567, 5583, 5607, 5627, 5648, 5666, 5690, 5716, 5731, 5749, 5762, 5781, 5801, 5821, 5846, 5861, 5885, 5908, 5932, 5953, 5975, 6000, 6023, 6046, 6071, 6090, 6116, 6144, 6177, 6209, 6223, 6239, 6252, 6266, 6279, 6296, 6311, 6326, 6348, 6370, 6392, 6427, 6449, 6481, 6510, 6535, 6562, 6584, 6612, 6641, 6666, 6690, 6712, 6733, 6756, 6768, 6783, 6799, 6816, 6847, 6880, 6913, 6943, 6973, 6996, 7022, 7063, 7096, 7127, 7160, 7187, 7224, 7254, 7282, 7309, 7339, 7361, 7382, 7408, 7430, 7448, 7466, 7481, 7495, 7508, 7529, 7548, 7573, 7590, 7606, 7623, 7653, 7669, 7693, 7718, 7740, 7756, 7777, 7796, 7826, 7857, 7886, 7916, 7943, 7979, 8004, 8029, 8047}

func (i OperationType) String() string {
	if i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
