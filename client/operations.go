package client

type operation interface {
	Process(state *albionState)
}

//go:generate stringer -type=OperationType
type OperationType uint16
const (
	// Account Operations
  Join                         OperationType = 2
  CreateAccount                OperationType = 3
  Login                        OperationType = 4
  CreateCharacter              OperationType = 6
  DeleteCharacter              OperationType = 7
  SelectCharacter              OperationType = 8
  RedeemKeycode                OperationType = 9

  GetGameServerByCluster       OperationType = 10


  Move                         OperationType = 17
  CastStart                    OperationType = 19
  CastCancel                   OperationType = 20
  ChannelingCancel             OperationType = 21

  InventoryDestroyItem         OperationType = 24
  InventoryMoveItem            OperationType = 25 // quick equip is move
  InventorySplitStack          OperationType = 26

  ChangeZone                   OperationType = 27
  ReportPlayer                 OperationType = 28

  RegisterToObject             OperationType = 31 // building, instanciate chest
  UnRegisterToObject           OperationType = 32 // building, instanciate chest

  Repair                       OperationType = 40
  Claim                        OperationType = 51
  GiveUp                       OperationType = 52

  FarmableHarvest              OperationType = 57
  FarmableGetProduct           OperationType = 61
  FarmableUseFoca              OperationType = 271

  AuctionGetOffers             OperationType = 71 // JSON
  AuctionGetRequests           OperationType = 72 // JSON
  AuctionBuyOffer              OperationType = 73
  AuctionAbortAuction          OperationType = 74
  AuctionAbortOffer            OperationType = 75
  AuctionAbortRequest          OperationType = 76
  AuctionSellRequest           OperationType = 77
  AuctionGetFinishedAuctions   OperationType = 78
  AuctionFetchAuction          OperationType = 79
  AuctionGetMyOpenOffers       OperationType = 80
  AuctionGetMyOpenRequests     OperationType = 81
  AuctionGetMyOpenAuctions     OperationType = 82
  AuctionGetItemsAverage       OperationType = 83

  ContainerOpen                OperationType = 84
  ContainerClose               OperationType = 85

  Respawn                      OperationType = 86
  Suicide                      OperationType = 87

  JoinGuild                    OperationType = 88
  LeaveGuild                   OperationType = 89
  CreateGuild                  OperationType = 90
  InviteToGuild                OperationType = 91
  DeclineGuildInvitation       OperationType = 92
  KickFromGuild                OperationType = 93

  GetAttackSchedule            OperationType = 105
  GetMatches                   OperationType = 107

  DepositToGuildAccount        OperationType = 117
  WithdrawalFromAccount        OperationType = 118
  ChangeGuildTax               OperationType = 120

  GetMyTerritories             OperationType = 122
  GetWorldMap                  OperationType = 126
  GetMyGuildInfo               OperationType = 129

  PromotePlayer                OperationType = 133
  DemotePlayer                 OperationType = 134

  ReadMail                     OperationType = 151
  SendNewMail                  OperationType = 152
  DeleteMail                   OperationType = 153

  GetClusterMapInfo            OperationType = 174
  AccessRightsChangeSettings   OperationType = 175
  Mount                        OperationType = 176
  MountCancel                  OperationType = 177
  BuyJourney                   OperationType = 178

  MakeHome                     OperationType = 182
  LeaveHome                    OperationType = 183

  GetIslandInfos               OperationType = 195

  GoldMarketGetAverageInfo     OperationType = 221

  RealEstateGetAuctionData     OperationType = 233 // ?
  RealEstateBidOnAuction       OperationType = 234 // ?

  FriendInvitationSend         OperationType = 240
  FriendInvitationResponseSend OperationType = 241
  FriendInvitationReceive      OperationType = 242
  FriendRemove                 OperationType = 243

  Stack                        OperationType = 244
  Sort                         OperationType = 245

  SendGraphicSettings          OperationType = 278
  SendHardwareConfiguration    OperationType = 279
)

