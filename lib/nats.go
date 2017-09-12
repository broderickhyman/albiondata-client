package lib

const (
	// Public Topics
	NatsGoldPricesIngest    = "goldprices.ingest"
	NatsGoldPricesDeduped   = "goldprices.deduped"
	NatsMarketOrdersIngest  = "marketorders.ingest"
	NatsMarketOrdersDeduped = "marketorders.deduped"
	NatsMapDataIngest       = "mapdata.ingest"
	NatsMapDataDeduped      = "mapdata.deduped"

	// Private Topics
	NatsSkillData           = "skills"
	NatsMarketNotifications = "marketnotifications"
)
