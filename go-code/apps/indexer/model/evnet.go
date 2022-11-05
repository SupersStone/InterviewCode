package model

// EventModel ETH event
type EventModel struct {
	BlockTimestamp int64    `json:"blockTimestamp"`
	ChainID        int64    `json:"chainId"`
	ChainName      string   `json:"chainName"`
	Log            LogModel `json:"log"`
}

// LogModel log detail
type LogModel struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

// GetContractAddr get event log address
func (e *EventModel) GetContractAddr() string {
	return e.Log.Address
}

// GetFunctionSign get event log function signature
func (e *EventModel) GetFunctionSign() string {
	return e.Log.Topics[0]
}
