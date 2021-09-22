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
	_ = x[opAttackStart-22]
	_ = x[opCastStart-23]
	_ = x[opCastCancel-24]
	_ = x[opTerminateToggleSpell-25]
	_ = x[opChannelingCancel-26]
	_ = x[opAttackBuildingStart-27]
	_ = x[opInventoryDestroyItem-28]
	_ = x[opInventoryMoveItem-29]
	_ = x[opInventoryRecoverItem-30]
	_ = x[opInventoryRecoverAllItems-31]
	_ = x[opInventorySplitStack-32]
	_ = x[opInventorySplitStackInto-33]
	_ = x[opGetClusterData-34]
	_ = x[opChangeCluster-35]
	_ = x[opConsoleCommand-36]
	_ = x[opChatMessage-37]
	_ = x[opReportClientError-38]
	_ = x[opRegisterToObject-39]
	_ = x[opUnRegisterFromObject-40]
	_ = x[opCraftBuildingChangeSettings-41]
	_ = x[opCraftBuildingTakeMoney-42]
	_ = x[opRepairBuildingChangeSettings-43]
	_ = x[opRepairBuildingTakeMoney-44]
	_ = x[opActionBuildingChangeSettings-45]
	_ = x[opHarvestStart-46]
	_ = x[opHarvestCancel-47]
	_ = x[opTakeSilver-48]
	_ = x[opActionOnBuildingStart-49]
	_ = x[opActionOnBuildingCancel-50]
	_ = x[opItemRerollQualityStart-51]
	_ = x[opItemRerollQualityCancel-52]
	_ = x[opInstallResourceStart-53]
	_ = x[opInstallResourceCancel-54]
	_ = x[opInstallSilver-55]
	_ = x[opBuildingFillNutrition-56]
	_ = x[opBuildingChangeRenovationState-57]
	_ = x[opBuildingBuySkin-58]
	_ = x[opBuildingClaim-59]
	_ = x[opBuildingGiveup-60]
	_ = x[opBuildingNutritionSilverStorageDeposit-61]
	_ = x[opBuildingNutritionSilverStorageWithdraw-62]
	_ = x[opBuildingNutritionSilverRewardSet-63]
	_ = x[opConstructionSiteCreate-64]
	_ = x[opPlaceableObjectPlace-65]
	_ = x[opPlaceableObjectPlaceCancel-66]
	_ = x[opPlaceableObjectPickup-67]
	_ = x[opFurnitureObjectUse-68]
	_ = x[opFarmableHarvest-69]
	_ = x[opFarmableFinishGrownItem-70]
	_ = x[opFarmableDestroy-71]
	_ = x[opFarmableGetProduct-72]
	_ = x[opFarmableFill-73]
	_ = x[opTearDownConstructionSite-74]
	_ = x[opCastleGateUse-75]
	_ = x[opAuctionCreateOffer-76]
	_ = x[opAuctionCreateRequest-77]
	_ = x[opAuctionGetOffers-78]
	_ = x[opAuctionGetRequests-79]
	_ = x[opAuctionBuyOffer-80]
	_ = x[opAuctionAbortAuction-81]
	_ = x[opAuctionModifyAuction-82]
	_ = x[opAuctionAbortOffer-83]
	_ = x[opAuctionAbortRequest-84]
	_ = x[opAuctionSellRequest-85]
	_ = x[opAuctionGetFinishedAuctions-86]
	_ = x[opAuctionGetFinishedAuctionsCount-87]
	_ = x[opAuctionFetchAuction-88]
	_ = x[opAuctionGetMyOpenOffers-89]
	_ = x[opAuctionGetMyOpenRequests-90]
	_ = x[opAuctionGetMyOpenAuctions-91]
	_ = x[opAuctionGetItemAverageStats-92]
	_ = x[opAuctionGetItemAverageValue-93]
	_ = x[opContainerOpen-94]
	_ = x[opContainerClose-95]
	_ = x[opContainerManageSubContainer-96]
	_ = x[opRespawn-97]
	_ = x[opSuicide-98]
	_ = x[opJoinGuild-99]
	_ = x[opLeaveGuild-100]
	_ = x[opCreateGuild-101]
	_ = x[opInviteToGuild-102]
	_ = x[opDeclineGuildInvitation-103]
	_ = x[opKickFromGuild-104]
	_ = x[opDuellingChallengePlayer-105]
	_ = x[opDuellingAcceptChallenge-106]
	_ = x[opDuellingDenyChallenge-107]
	_ = x[opChangeClusterTax-108]
	_ = x[opClaimTerritory-109]
	_ = x[opGiveUpTerritory-110]
	_ = x[opChangeTerritoryAccessRights-111]
	_ = x[opGetMonolithInfo-112]
	_ = x[opGetClaimInfo-113]
	_ = x[opGetAttackInfo-114]
	_ = x[opGetTerritorySeasonPoints-115]
	_ = x[opGetAttackSchedule-116]
	_ = x[opScheduleAttack-117]
	_ = x[opGetMatches-118]
	_ = x[opGetMatchDetails-119]
	_ = x[opJoinMatch-120]
	_ = x[opLeaveMatch-121]
	_ = x[opChangeChatSettings-122]
	_ = x[opLogoutStart-123]
	_ = x[opLogoutCancel-124]
	_ = x[opClaimOrbStart-125]
	_ = x[opClaimOrbCancel-126]
	_ = x[opMatchLootChestOpeningStart-127]
	_ = x[opMatchLootChestOpeningCancel-128]
	_ = x[opDepositToGuildAccount-129]
	_ = x[opWithdrawalFromAccount-130]
	_ = x[opChangeGuildPayUpkeepFlag-131]
	_ = x[opChangeGuildTax-132]
	_ = x[opGetMyTerritories-133]
	_ = x[opMorganaCommand-134]
	_ = x[opGetServerInfo-135]
	_ = x[opInviteMercenaryToMatch-136]
	_ = x[opSubscribeToCluster-137]
	_ = x[opAnswerMercenaryInvitation-138]
	_ = x[opGetCharacterEquipment-139]
	_ = x[opGetCharacterSteamAchievements-140]
	_ = x[opGetCharacterStats-141]
	_ = x[opGetKillHistoryDetails-142]
	_ = x[opLearnMasteryLevel-143]
	_ = x[opReSpecAchievement-144]
	_ = x[opChangeAvatar-145]
	_ = x[opGetRankings-146]
	_ = x[opGetRank-147]
	_ = x[opGetGvgSeasonRankings-148]
	_ = x[opGetGvgSeasonRank-149]
	_ = x[opGetGvgSeasonHistoryRankings-150]
	_ = x[opGetGvgSeasonGuildMemberHistory-151]
	_ = x[opKickFromGvGMatch-152]
	_ = x[opGetChestLogs-153]
	_ = x[opGetAccessRightLogs-154]
	_ = x[opGetGuildAccountLogs-155]
	_ = x[opGetGuildAccountLogsLargeAmount-156]
	_ = x[opInviteToPlayerTrade-157]
	_ = x[opPlayerTradeCancel-158]
	_ = x[opPlayerTradeInvitationAccept-159]
	_ = x[opPlayerTradeAddItem-160]
	_ = x[opPlayerTradeRemoveItem-161]
	_ = x[opPlayerTradeAcceptTrade-162]
	_ = x[opPlayerTradeSetSilverOrGold-163]
	_ = x[opSendMiniMapPing-164]
	_ = x[opStuck-165]
	_ = x[opBuyRealEstate-166]
	_ = x[opClaimRealEstate-167]
	_ = x[opGiveUpRealEstate-168]
	_ = x[opChangeRealEstateOutline-169]
	_ = x[opGetMailInfos-170]
	_ = x[opGetMailCount-171]
	_ = x[opReadMail-172]
	_ = x[opSendNewMail-173]
	_ = x[opDeleteMail-174]
	_ = x[opMarkMailUnread-175]
	_ = x[opClaimAttachmentFromMail-176]
	_ = x[opUpdateLfgInfo-177]
	_ = x[opGetLfgInfos-178]
	_ = x[opGetMyGuildLfgInfo-179]
	_ = x[opGetLfgDescriptionText-180]
	_ = x[opLfgApplyToGuild-181]
	_ = x[opAnswerLfgGuildApplication-182]
	_ = x[opRegisterChatPeer-183]
	_ = x[opSendChatMessage-184]
	_ = x[opJoinChatChannel-185]
	_ = x[opLeaveChatChannel-186]
	_ = x[opSendWhisperMessage-187]
	_ = x[opSay-188]
	_ = x[opPlayEmote-189]
	_ = x[opStopEmote-190]
	_ = x[opGetClusterMapInfo-191]
	_ = x[opAccessRightsChangeSettings-192]
	_ = x[opMount-193]
	_ = x[opMountCancel-194]
	_ = x[opBuyJourney-195]
	_ = x[opSetSaleStatusForEstate-196]
	_ = x[opResolveGuildOrPlayerName-197]
	_ = x[opGetRespawnInfos-198]
	_ = x[opMakeHome-199]
	_ = x[opLeaveHome-200]
	_ = x[opResurrectionReply-201]
	_ = x[opAllianceCreate-202]
	_ = x[opAllianceDisband-203]
	_ = x[opAllianceGetMemberInfos-204]
	_ = x[opAllianceInvite-205]
	_ = x[opAllianceAnswerInvitation-206]
	_ = x[opAllianceCancelInvitation-207]
	_ = x[opAllianceKickGuild-208]
	_ = x[opAllianceLeave-209]
	_ = x[opAllianceChangeGoldPaymentFlag-210]
	_ = x[opAllianceGetDetailInfo-211]
	_ = x[opGetIslandInfos-212]
	_ = x[opAbandonMyIsland-213]
	_ = x[opBuyMyIsland-214]
	_ = x[opBuyGuildIsland-215]
	_ = x[opAbandonGuildIsland-216]
	_ = x[opUpgradeMyIsland-217]
	_ = x[opUpgradeGuildIsland-218]
	_ = x[opMoveMyIsland-219]
	_ = x[opMoveGuildIsland-220]
	_ = x[opTerritoryFillNutrition-221]
	_ = x[opTeleportBack-222]
	_ = x[opPartyInvitePlayer-223]
	_ = x[opPartyAnswerInvitation-224]
	_ = x[opPartyLeave-225]
	_ = x[opPartyKickPlayer-226]
	_ = x[opPartyMakeLeader-227]
	_ = x[opPartyChangeLootSetting-228]
	_ = x[opPartyMarkObject-229]
	_ = x[opPartySetRole-230]
	_ = x[opGetGuildMOTD-231]
	_ = x[opSetGuildMOTD-232]
	_ = x[opExitEnterStart-233]
	_ = x[opExitEnterCancel-234]
	_ = x[opQuestGiverRequest-235]
	_ = x[opGoldMarketGetBuyOffer-236]
	_ = x[opGoldMarketGetBuyOfferFromSilver-237]
	_ = x[opGoldMarketGetSellOffer-238]
	_ = x[opGoldMarketGetSellOfferFromSilver-239]
	_ = x[opGoldMarketBuyGold-240]
	_ = x[opGoldMarketSellGold-241]
	_ = x[opGoldMarketCreateSellOrder-242]
	_ = x[opGoldMarketCreateBuyOrder-243]
	_ = x[opGoldMarketGetInfos-244]
	_ = x[opGoldMarketCancelOrder-245]
	_ = x[opUnknown246-246]
	_ = x[opUnknown247-247]
	_ = x[opGoldMarketGetAverageInfo-248]
	_ = x[opSiegeCampClaimStart-249]
	_ = x[opSiegeCampClaimCancel-250]
	_ = x[opTreasureChestUsingStart-251]
	_ = x[opTreasureChestUsingCancel-252]
	_ = x[opUseLootChest-253]
	_ = x[opUseShrine-254]
	_ = x[opLaborerStartJob-255]
	_ = x[opLaborerTakeJobLoot-256]
	_ = x[opLaborerDismiss-257]
	_ = x[opLaborerMove-258]
	_ = x[opLaborerBuyItem-259]
	_ = x[opLaborerUpgrade-260]
	_ = x[opBuyPremium-261]
	_ = x[opBuyTrial-262]
	_ = x[opRealEstateGetAuctionData-263]
	_ = x[opRealEstateBidOnAuction-264]
	_ = x[opGetSiegeCampCooldown-265]
	_ = x[opFriendInvite-266]
	_ = x[opFriendAnswerInvitation-267]
	_ = x[opFriendCancelnvitation-268]
	_ = x[opFriendRemove-269]
	_ = x[opInventoryStack-270]
	_ = x[opInventorySort-271]
	_ = x[opEquipmentItemChangeSpell-272]
	_ = x[opExpeditionRegister-273]
	_ = x[opExpeditionRegisterCancel-274]
	_ = x[opJoinExpedition-275]
	_ = x[opDeclineExpeditionInvitation-276]
	_ = x[opVoteStart-277]
	_ = x[opVoteDoVote-278]
	_ = x[opRatingDoRate-279]
	_ = x[opEnteringExpeditionStart-280]
	_ = x[opEnteringExpeditionCancel-281]
	_ = x[opActivateExpeditionCheckPoint-282]
	_ = x[opArenaRegister-283]
	_ = x[opArenaRegisterCancel-284]
	_ = x[opArenaLeave-285]
	_ = x[opJoinArenaMatch-286]
	_ = x[opDeclineArenaInvitation-287]
	_ = x[opEnteringArenaStart-288]
	_ = x[opEnteringArenaCancel-289]
	_ = x[opArenaCustomMatch-290]
	_ = x[opArenaCustomMatchCreate-291]
	_ = x[opUpdateCharacterStatement-292]
	_ = x[opBoostFarmable-293]
	_ = x[opGetStrikeHistory-294]
	_ = x[opUseFunction-295]
	_ = x[opUsePortalEntrance-296]
	_ = x[opResetPortalBinding-297]
	_ = x[opQueryPortalBinding-298]
	_ = x[opClaimPaymentTransaction-299]
	_ = x[opChangeUseFlag-300]
	_ = x[opClientPerformanceStats-301]
	_ = x[opExtendedHardwareStats-302]
	_ = x[opClientLowMemoryWarning-303]
	_ = x[opTerritoryClaimStart-304]
	_ = x[opTerritoryClaimCancel-305]
	_ = x[opRequestAppStoreProducts-306]
	_ = x[opVerifyProductPurchase-307]
	_ = x[opQueryGuildPlayerStats-308]
	_ = x[opQueryAllianceGuildStats-309]
	_ = x[opTrackAchievements-310]
	_ = x[opSetAchievementsAutoLearn-311]
	_ = x[opDepositItemToGuildCurrency-312]
	_ = x[opWithdrawalItemFromGuildCurrency-313]
	_ = x[opAuctionSellSpecificItemRequest-314]
	_ = x[opFishingStart-315]
	_ = x[opFishingCasting-316]
	_ = x[opFishingCast-317]
	_ = x[opFishingCatch-318]
	_ = x[opFishingPull-319]
	_ = x[opFishingGiveLine-320]
	_ = x[opFishingFinish-321]
	_ = x[opFishingCancel-322]
	_ = x[opCreateGuildAccessTag-323]
	_ = x[opDeleteGuildAccessTag-324]
	_ = x[opRenameGuildAccessTag-325]
	_ = x[opFlagGuildAccessTagGuildPermission-326]
	_ = x[opAssignGuildAccessTag-327]
	_ = x[opRemoveGuildAccessTagFromPlayer-328]
	_ = x[opModifyGuildAccessTagEditors-329]
	_ = x[opRequestPublicAccessTags-330]
	_ = x[opChangeAccessTagPublicFlag-331]
	_ = x[opUpdateGuildAccessTag-332]
	_ = x[opSteamStartMicrotransaction-333]
	_ = x[opSteamFinishMicrotransaction-334]
	_ = x[opSteamIdHasActiveAccount-335]
	_ = x[opCheckEmailAccountState-336]
	_ = x[opLinkAccountToSteamId-337]
	_ = x[opBuyGvgSeasonBooster-338]
	_ = x[opChangeFlaggingPrepare-339]
	_ = x[opOverCharge-340]
	_ = x[opOverChargeEnd-341]
	_ = x[opRequestTrusted-342]
	_ = x[opChangeGuildLogo-343]
	_ = x[opPartyFinderRegisterForUpdates-344]
	_ = x[opPartyFinderUnregisterForUpdates-345]
	_ = x[opPartyFinderEnlistNewPartySearch-346]
	_ = x[opPartyFinderDeletePartySearch-347]
	_ = x[opPartyFinderChangePartySearch-348]
	_ = x[opPartyFinderChangeRole-349]
	_ = x[opPartyFinderApplyForGroup-350]
	_ = x[opPartyFinderAcceptOrDeclineApplyForGroup-351]
	_ = x[opPartyFinderGetEquipmentSnapshot-352]
	_ = x[opPartyFinderRegisterApplicants-353]
	_ = x[opPartyFinderUnregisterApplicants-354]
	_ = x[opPartyFinderFulltextSearch-355]
	_ = x[opPartyFinderRequestEquipmentSnapshot-356]
	_ = x[opGetPersonalSeasonTrackerData-357]
	_ = x[opUseConsumableFromInventory-358]
	_ = x[opClaimPersonalSeasonReward-359]
	_ = x[opEasyAntiCheatMessageToServer-360]
	_ = x[opSetNextTutorialState-361]
	_ = x[opAddPlayerToMuteList-362]
	_ = x[opRemovePlayerFromMuteList-363]
	_ = x[opProductShopUserEvent-364]
	_ = x[opGetVanityUnlocks-365]
	_ = x[opBuyVanityUnlocks-366]
	_ = x[opGetMountSkins-367]
	_ = x[opSetMountSkin-368]
	_ = x[opSetWardrobe-369]
	_ = x[opChangeCustomization-370]
	_ = x[opSetFavoriteIsland-371]
	_ = x[opGetGuildChallengePoints-372]
	_ = x[opTravelToHideout-373]
	_ = x[opSmartQueueJoin-374]
	_ = x[opSmartQueueLeave-375]
	_ = x[opSmartQueueSelectSpawnCluster-376]
	_ = x[opUpgradeHideout-377]
	_ = x[opInitHideoutAttackStart-378]
	_ = x[opInitHideoutAttackCancel-379]
	_ = x[opHideoutFillNutrition-380]
	_ = x[opHideoutGetInfo-381]
	_ = x[opHideoutGetOwnerInfo-382]
	_ = x[opHideoutSetTribute-383]
	_ = x[opOpenWorldAttackScheduleStart-384]
	_ = x[opOpenWorldAttackScheduleCancel-385]
	_ = x[opOpenWorldAttackConquerStart-386]
	_ = x[opOpenWorldAttackConquerCancel-387]
	_ = x[opGetOpenWorldAttackDetails-388]
	_ = x[opGetNextOpenWorldAttackScheduleTime-389]
	_ = x[opRecoverVaultFromHideout-390]
	_ = x[opGetGuildEnergyDrainInfo-391]
	_ = x[opChannelingUpdate-392]
}

const _OperationType_name = "opUnusedopPingopJoinopCreateAccountopLoginopSendCrashLogopSendTraceRouteopSendVfxStatsopSendGamePingInfoopCreateCharacteropDeleteCharacteropSelectCharacteropRedeemKeycodeopGetGameServerByClusteropGetActiveSubscriptionopGetShopPurchaseUrlopGetBuyTrialDetailsopGetReferralSeasonDetailsopGetReferralLinkopGetAvailableTrialKeysopGetShopTilesForCategoryopMoveopAttackStartopCastStartopCastCancelopTerminateToggleSpellopChannelingCancelopAttackBuildingStartopInventoryDestroyItemopInventoryMoveItemopInventoryRecoverItemopInventoryRecoverAllItemsopInventorySplitStackopInventorySplitStackIntoopGetClusterDataopChangeClusteropConsoleCommandopChatMessageopReportClientErroropRegisterToObjectopUnRegisterFromObjectopCraftBuildingChangeSettingsopCraftBuildingTakeMoneyopRepairBuildingChangeSettingsopRepairBuildingTakeMoneyopActionBuildingChangeSettingsopHarvestStartopHarvestCancelopTakeSilveropActionOnBuildingStartopActionOnBuildingCancelopItemRerollQualityStartopItemRerollQualityCancelopInstallResourceStartopInstallResourceCancelopInstallSilveropBuildingFillNutritionopBuildingChangeRenovationStateopBuildingBuySkinopBuildingClaimopBuildingGiveupopBuildingNutritionSilverStorageDepositopBuildingNutritionSilverStorageWithdrawopBuildingNutritionSilverRewardSetopConstructionSiteCreateopPlaceableObjectPlaceopPlaceableObjectPlaceCancelopPlaceableObjectPickupopFurnitureObjectUseopFarmableHarvestopFarmableFinishGrownItemopFarmableDestroyopFarmableGetProductopFarmableFillopTearDownConstructionSiteopCastleGateUseopAuctionCreateOfferopAuctionCreateRequestopAuctionGetOffersopAuctionGetRequestsopAuctionBuyOfferopAuctionAbortAuctionopAuctionModifyAuctionopAuctionAbortOfferopAuctionAbortRequestopAuctionSellRequestopAuctionGetFinishedAuctionsopAuctionGetFinishedAuctionsCountopAuctionFetchAuctionopAuctionGetMyOpenOffersopAuctionGetMyOpenRequestsopAuctionGetMyOpenAuctionsopAuctionGetItemAverageStatsopAuctionGetItemAverageValueopContainerOpenopContainerCloseopContainerManageSubContaineropRespawnopSuicideopJoinGuildopLeaveGuildopCreateGuildopInviteToGuildopDeclineGuildInvitationopKickFromGuildopDuellingChallengePlayeropDuellingAcceptChallengeopDuellingDenyChallengeopChangeClusterTaxopClaimTerritoryopGiveUpTerritoryopChangeTerritoryAccessRightsopGetMonolithInfoopGetClaimInfoopGetAttackInfoopGetTerritorySeasonPointsopGetAttackScheduleopScheduleAttackopGetMatchesopGetMatchDetailsopJoinMatchopLeaveMatchopChangeChatSettingsopLogoutStartopLogoutCancelopClaimOrbStartopClaimOrbCancelopMatchLootChestOpeningStartopMatchLootChestOpeningCancelopDepositToGuildAccountopWithdrawalFromAccountopChangeGuildPayUpkeepFlagopChangeGuildTaxopGetMyTerritoriesopMorganaCommandopGetServerInfoopInviteMercenaryToMatchopSubscribeToClusteropAnswerMercenaryInvitationopGetCharacterEquipmentopGetCharacterSteamAchievementsopGetCharacterStatsopGetKillHistoryDetailsopLearnMasteryLevelopReSpecAchievementopChangeAvataropGetRankingsopGetRankopGetGvgSeasonRankingsopGetGvgSeasonRankopGetGvgSeasonHistoryRankingsopGetGvgSeasonGuildMemberHistoryopKickFromGvGMatchopGetChestLogsopGetAccessRightLogsopGetGuildAccountLogsopGetGuildAccountLogsLargeAmountopInviteToPlayerTradeopPlayerTradeCancelopPlayerTradeInvitationAcceptopPlayerTradeAddItemopPlayerTradeRemoveItemopPlayerTradeAcceptTradeopPlayerTradeSetSilverOrGoldopSendMiniMapPingopStuckopBuyRealEstateopClaimRealEstateopGiveUpRealEstateopChangeRealEstateOutlineopGetMailInfosopGetMailCountopReadMailopSendNewMailopDeleteMailopMarkMailUnreadopClaimAttachmentFromMailopUpdateLfgInfoopGetLfgInfosopGetMyGuildLfgInfoopGetLfgDescriptionTextopLfgApplyToGuildopAnswerLfgGuildApplicationopRegisterChatPeeropSendChatMessageopJoinChatChannelopLeaveChatChannelopSendWhisperMessageopSayopPlayEmoteopStopEmoteopGetClusterMapInfoopAccessRightsChangeSettingsopMountopMountCancelopBuyJourneyopSetSaleStatusForEstateopResolveGuildOrPlayerNameopGetRespawnInfosopMakeHomeopLeaveHomeopResurrectionReplyopAllianceCreateopAllianceDisbandopAllianceGetMemberInfosopAllianceInviteopAllianceAnswerInvitationopAllianceCancelInvitationopAllianceKickGuildopAllianceLeaveopAllianceChangeGoldPaymentFlagopAllianceGetDetailInfoopGetIslandInfosopAbandonMyIslandopBuyMyIslandopBuyGuildIslandopAbandonGuildIslandopUpgradeMyIslandopUpgradeGuildIslandopMoveMyIslandopMoveGuildIslandopTerritoryFillNutritionopTeleportBackopPartyInvitePlayeropPartyAnswerInvitationopPartyLeaveopPartyKickPlayeropPartyMakeLeaderopPartyChangeLootSettingopPartyMarkObjectopPartySetRoleopGetGuildMOTDopSetGuildMOTDopExitEnterStartopExitEnterCancelopQuestGiverRequestopGoldMarketGetBuyOfferopGoldMarketGetBuyOfferFromSilveropGoldMarketGetSellOfferopGoldMarketGetSellOfferFromSilveropGoldMarketBuyGoldopGoldMarketSellGoldopGoldMarketCreateSellOrderopGoldMarketCreateBuyOrderopGoldMarketGetInfosopGoldMarketCancelOrderopUnknown246opUnknown247opGoldMarketGetAverageInfoopSiegeCampClaimStartopSiegeCampClaimCancelopTreasureChestUsingStartopTreasureChestUsingCancelopUseLootChestopUseShrineopLaborerStartJobopLaborerTakeJobLootopLaborerDismissopLaborerMoveopLaborerBuyItemopLaborerUpgradeopBuyPremiumopBuyTrialopRealEstateGetAuctionDataopRealEstateBidOnAuctionopGetSiegeCampCooldownopFriendInviteopFriendAnswerInvitationopFriendCancelnvitationopFriendRemoveopInventoryStackopInventorySortopEquipmentItemChangeSpellopExpeditionRegisteropExpeditionRegisterCancelopJoinExpeditionopDeclineExpeditionInvitationopVoteStartopVoteDoVoteopRatingDoRateopEnteringExpeditionStartopEnteringExpeditionCancelopActivateExpeditionCheckPointopArenaRegisteropArenaRegisterCancelopArenaLeaveopJoinArenaMatchopDeclineArenaInvitationopEnteringArenaStartopEnteringArenaCancelopArenaCustomMatchopArenaCustomMatchCreateopUpdateCharacterStatementopBoostFarmableopGetStrikeHistoryopUseFunctionopUsePortalEntranceopResetPortalBindingopQueryPortalBindingopClaimPaymentTransactionopChangeUseFlagopClientPerformanceStatsopExtendedHardwareStatsopClientLowMemoryWarningopTerritoryClaimStartopTerritoryClaimCancelopRequestAppStoreProductsopVerifyProductPurchaseopQueryGuildPlayerStatsopQueryAllianceGuildStatsopTrackAchievementsopSetAchievementsAutoLearnopDepositItemToGuildCurrencyopWithdrawalItemFromGuildCurrencyopAuctionSellSpecificItemRequestopFishingStartopFishingCastingopFishingCastopFishingCatchopFishingPullopFishingGiveLineopFishingFinishopFishingCancelopCreateGuildAccessTagopDeleteGuildAccessTagopRenameGuildAccessTagopFlagGuildAccessTagGuildPermissionopAssignGuildAccessTagopRemoveGuildAccessTagFromPlayeropModifyGuildAccessTagEditorsopRequestPublicAccessTagsopChangeAccessTagPublicFlagopUpdateGuildAccessTagopSteamStartMicrotransactionopSteamFinishMicrotransactionopSteamIdHasActiveAccountopCheckEmailAccountStateopLinkAccountToSteamIdopBuyGvgSeasonBoosteropChangeFlaggingPrepareopOverChargeopOverChargeEndopRequestTrustedopChangeGuildLogoopPartyFinderRegisterForUpdatesopPartyFinderUnregisterForUpdatesopPartyFinderEnlistNewPartySearchopPartyFinderDeletePartySearchopPartyFinderChangePartySearchopPartyFinderChangeRoleopPartyFinderApplyForGroupopPartyFinderAcceptOrDeclineApplyForGroupopPartyFinderGetEquipmentSnapshotopPartyFinderRegisterApplicantsopPartyFinderUnregisterApplicantsopPartyFinderFulltextSearchopPartyFinderRequestEquipmentSnapshotopGetPersonalSeasonTrackerDataopUseConsumableFromInventoryopClaimPersonalSeasonRewardopEasyAntiCheatMessageToServeropSetNextTutorialStateopAddPlayerToMuteListopRemovePlayerFromMuteListopProductShopUserEventopGetVanityUnlocksopBuyVanityUnlocksopGetMountSkinsopSetMountSkinopSetWardrobeopChangeCustomizationopSetFavoriteIslandopGetGuildChallengePointsopTravelToHideoutopSmartQueueJoinopSmartQueueLeaveopSmartQueueSelectSpawnClusteropUpgradeHideoutopInitHideoutAttackStartopInitHideoutAttackCancelopHideoutFillNutritionopHideoutGetInfoopHideoutGetOwnerInfoopHideoutSetTributeopOpenWorldAttackScheduleStartopOpenWorldAttackScheduleCancelopOpenWorldAttackConquerStartopOpenWorldAttackConquerCancelopGetOpenWorldAttackDetailsopGetNextOpenWorldAttackScheduleTimeopRecoverVaultFromHideoutopGetGuildEnergyDrainInfoopChannelingUpdate"

var _OperationType_index = [...]uint16{0, 8, 14, 20, 35, 42, 56, 72, 86, 104, 121, 138, 155, 170, 194, 217, 237, 257, 283, 300, 323, 348, 354, 367, 378, 390, 412, 430, 451, 473, 492, 514, 540, 561, 586, 602, 617, 633, 646, 665, 683, 705, 734, 758, 788, 813, 843, 857, 872, 884, 907, 931, 955, 980, 1002, 1025, 1040, 1063, 1094, 1111, 1126, 1142, 1181, 1221, 1255, 1279, 1301, 1329, 1352, 1372, 1389, 1414, 1431, 1451, 1465, 1491, 1506, 1526, 1548, 1566, 1586, 1603, 1624, 1646, 1665, 1686, 1706, 1734, 1767, 1788, 1812, 1838, 1864, 1892, 1920, 1935, 1951, 1980, 1989, 1998, 2009, 2021, 2034, 2049, 2073, 2088, 2113, 2138, 2161, 2179, 2195, 2212, 2241, 2258, 2272, 2287, 2313, 2332, 2348, 2360, 2377, 2388, 2400, 2420, 2433, 2447, 2462, 2478, 2506, 2535, 2558, 2581, 2607, 2623, 2641, 2657, 2672, 2696, 2716, 2743, 2766, 2797, 2816, 2839, 2858, 2877, 2891, 2904, 2913, 2935, 2953, 2982, 3014, 3032, 3046, 3066, 3087, 3119, 3140, 3159, 3188, 3208, 3231, 3255, 3283, 3300, 3307, 3322, 3339, 3357, 3382, 3396, 3410, 3420, 3433, 3445, 3461, 3486, 3501, 3514, 3533, 3556, 3573, 3600, 3618, 3635, 3652, 3670, 3690, 3695, 3706, 3717, 3736, 3764, 3771, 3784, 3796, 3820, 3846, 3863, 3873, 3884, 3903, 3919, 3936, 3960, 3976, 4002, 4028, 4047, 4062, 4093, 4116, 4132, 4149, 4162, 4178, 4198, 4215, 4235, 4249, 4266, 4290, 4304, 4323, 4346, 4358, 4375, 4392, 4416, 4433, 4447, 4461, 4475, 4491, 4508, 4527, 4550, 4583, 4607, 4641, 4660, 4680, 4707, 4733, 4753, 4776, 4788, 4800, 4826, 4847, 4869, 4894, 4920, 4934, 4945, 4962, 4982, 4998, 5011, 5027, 5043, 5055, 5065, 5091, 5115, 5137, 5151, 5175, 5198, 5212, 5228, 5243, 5269, 5289, 5315, 5331, 5360, 5371, 5383, 5397, 5422, 5448, 5478, 5493, 5514, 5526, 5542, 5566, 5586, 5607, 5625, 5649, 5675, 5690, 5708, 5721, 5740, 5760, 5780, 5805, 5820, 5844, 5867, 5891, 5912, 5934, 5959, 5982, 6005, 6030, 6049, 6075, 6103, 6136, 6168, 6182, 6198, 6211, 6225, 6238, 6255, 6270, 6285, 6307, 6329, 6351, 6386, 6408, 6440, 6469, 6494, 6521, 6543, 6571, 6600, 6625, 6649, 6671, 6692, 6715, 6727, 6742, 6758, 6775, 6806, 6839, 6872, 6902, 6932, 6955, 6981, 7022, 7055, 7086, 7119, 7146, 7183, 7213, 7241, 7268, 7298, 7320, 7341, 7367, 7389, 7407, 7425, 7440, 7454, 7467, 7488, 7507, 7532, 7549, 7565, 7582, 7612, 7628, 7652, 7677, 7699, 7715, 7736, 7755, 7785, 7816, 7845, 7875, 7902, 7938, 7963, 7988, 8006}

func (i OperationType) String() string {
	if i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
