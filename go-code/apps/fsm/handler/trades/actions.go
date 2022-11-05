package trades

import "dao-exchange/internal/models/indexer"

// Actions interface
type Actions interface {
	HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*OrderStatusHandler, error)
}
