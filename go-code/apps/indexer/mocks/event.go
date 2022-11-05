package mocks

import (
	"time"

	"dao-exchange/internal/cache"
	"dao-exchange/internal/models"
	"dao-exchange/internal/models/indexer"
)

// EventDefCache mock event definiton cache
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
		ContractAddress: "0x37bbbec6cd25bb128887b39b1d9c922cddbd0987",
		EventSignature:  "0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985",
		ContractType:    "d1verse-dex",
		EventName:       "OrderCancelled",
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
		ContractAddress: "0x37bbbec6cd25bb128887b39b1d9c922cddbd0987",
		EventSignature:  "0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936888",
		ContractType:    "d1verse-dex",
		EventName:       "AllOrdersCancelled",
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
		ContractAddress: "0xc81c6ced11a5d9b2a80f00267a4646bc6d7e4e95",
		EventSignature:  "0x47354e3d476d837f131469d5dcc8adf1e40a983ba4c457e15b41c178238425c9",
		ContractType:    "proxy",
		EventName:       "ProxyOfUser",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)
	cache.NewLocalEventCache().InsertCache(defs)
}
