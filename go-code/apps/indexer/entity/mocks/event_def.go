package mocks

import (
	"time"

	"dao-exchange/internal/cache"
	"dao-exchange/internal/models"
	"dao-exchange/internal/models/indexer"
)

// EventDefCache init local cache
func EventDefCache() {
	defs := []indexer.EventDefinition{}
	def := indexer.EventDefinition{
		Base: models.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ContractAddress: "0x37bbbec6cd25bb128887b39b1d9c922cddbd0987",
		EventSignature:  "0xd03d2bbbdcfad01f71197db5bfbb7072d418388655e881877a85fbfbca864f6b",
		ContractType:    "d1verse-dex",
		EventName:       "FixedPriceOrderMatched",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)

	def = indexer.EventDefinition{
		Base: models.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ContractAddress: "0x83be785152a581e8F801B85e67CeC3820a614d40",
		EventSignature:  "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		ContractType:    "erc721",
		EventName:       "Transfer",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)
	cache.NewLocalEventCache().InsertCache(defs)
}
