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

  InventoryDestroyItem         OperationType = 23
  InventoryMoveItem            OperationType = 24 // quick equip is move
  InventorySplitStack          OperationType = 25

  Repair                       OperationType = 40
  Claim                        OperationType = 50
  GiveUp                       OperationType = 51
  Place                        OperationType = 56
  PickUp                       OperationType = 58

  AuctionGetOffers             OperationType = 70
  AuctionGetRequests           OperationType = 71
  AuctionBuyOffer              OperationType = 72
  AuctionAbortAuction          OperationType = 73
  AuctionAbortOffer            OperationType = 74
  AuctionAbortRequest          OperationType = 75
  AuctionSellRequest           OperationType = 76
  AuctionGetFinishedAuctions   OperationType = 77
  AuctionFetchAuction          OperationType = 78
  AuctionGetMyOpenOffers       OperationType = 79
  AuctionGetMyOpenRequests     OperationType = 80
  AuctionGetMyOpenAuctions     OperationType = 81
  AuctionGetItemsAverage       OperationType = 82

  GetMyGuildInfo               OperationType = 105

  DepositToGuildAccount        OperationType = 116
  WithdrawalFromAccount        OperationType = 117
  ChangeGuildTax               OperationType = 119

  ReadMail                     OperationType = 150
  SendNewMail                  OperationType = 151
  DeleteMail                   OperationType = 152

  GetClusterMapInfo            OperationType = 169

  GoldMarketGetAverageInfo     OperationType = 220

  RealEstateGetAuctionData     OperationType = 232 // ?
  RealEstateBidOnAuction       OperationType = 233 // ?

  FriendInvitationSend         OperationType = 239
  FriendInvitationResponseSend OperationType = 240
  FriendInvitationReceive      OperationType = 241
  FriendRemove                 OperationType = 242

  Stack                        OperationType = 243
  Sort                         OperationType = 244
)
