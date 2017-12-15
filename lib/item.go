package lib

type EquipmentItem struct {
	ItemIndex  int64  `json:"ItemIndex"`
	Quantity   int16  `json:"Quantity"`
	Creator    string `json:"Creator"`
	Quality    byte   `json:"Quality"`
	Durability int64  `json:"Durability"` // was multiplied by 10000
}

type StackableItem struct {
	ItemIndex int64 `json:"ItemIndex"`
	Quantity  int16 `json:"Quantity"`
}

type FurnitureItem struct {
	ItemIndex  int64  `json:"ItemIndex"`
	Quantity   int16  `json:"Quantity"`
	Creator    string `json:"Creator"`
	Durability int64  `json:"Durability"` // was multiplied by 10000
}

type JournalItem struct {
	ItemIndex  int64  `json:"ItemIndex"`
	Quantity   int16  `json:"Quantity"`
	Creator    string `json:"Creator"`
	FameStored int64  `json:"FameStored"` // was multiplied by 10000
}

const (
	ItemType_EquipmentItem = "EquipmentItem"
	ItemType_FurnitureItem = "FurnitureItem"
	ItemType_JournalItem   = "JournalItem"
	ItemType_StackableItem = "StackableItem"
)

type ItemContainer struct {
	Item     interface{} `json:"Item"`
	ItemType string      `json:"ItemType"`
}

func (container *ItemContainer) AsEquipmentItem(itemIndex int64, quantity int16, creator string, quality byte, durability int64) {
	container.Item = EquipmentItem{ItemIndex: itemIndex, Quantity: quantity, Creator: creator, Quality: quality, Durability: durability}
	container.ItemType = ItemType_EquipmentItem
}

func (container *ItemContainer) AsStackableItem(itemIndex int64, quantity int16) {
	container.Item = StackableItem{ItemIndex: itemIndex, Quantity: quantity}
	container.ItemType = ItemType_StackableItem
}

func (container *ItemContainer) AsFurnitureItem(itemIndex int64, quantity int16, creator string, durability int64) {
	container.Item = FurnitureItem{ItemIndex: itemIndex, Quantity: quantity, Creator: creator, Durability: durability}
	container.ItemType = ItemType_FurnitureItem
}

func (container *ItemContainer) AsJournalItem(itemIndex int64, quantity int16, creator string, fameStored int64) {
	container.Item = JournalItem{ItemIndex: itemIndex, Quantity: quantity, Creator: creator, FameStored: fameStored}
	container.ItemType = ItemType_JournalItem
}
