package commands

type TradeTransInfo struct {
	Cmd           uint8   `json:"cmd"`           // Operation code
	CustomComment string  `json:"customComment"` // The value the customer may provide in order to retrieve it later.
	Expiration    int64   `json:"expiration"`    // Null if order is not closed
	Offset        int     `json:"offset"`        // Trailing offset
	Order         uint    `json:"order"`         // Order number for opened transaction
	Price         float32 `json:"price"`         // Trade price
	Sl            float32 `json:"sl"`            // Zero if stop loss is not set (in base currency)
	Symbol        string  `json:"symbol"`        // symbol name or null for deposit/withdrawal operations
	Tp            float32 `json:"tp"`            // Zero if take profit is not set (in base currency)
	Type          uint8   `json:"type"`          // Trade transaction type
	Volume        float32 `json:"volume"`        // Volume in lots
}

type TradeTransactionResultData struct {
	Order uint
}

type TradeTransactionResult struct {
	Result
	ReturnData TradeTransactionResultData
}

type tradeTransactionRequest struct {
	TradeTransInfo TradeTransInfo `json:"tradeTransInfo"` // if true then only opened trades will be returned
}

func NewTradeTransactionRequestCommand(cmd uint8, customComment string, expiration int64, offset int, order uint, price float32, sl float32, symbol string, tp float32, transactionType uint8, volume float32) *Command {
	return &Command{
		Command: "tradeTransaction",
		Arguments: &tradeTransactionRequest{
			TradeTransInfo: TradeTransInfo{
				Cmd:           cmd,
				CustomComment: customComment,
				Expiration:    expiration,
				Offset:        offset,
				Order:         order,
				Price:         price,
				Sl:            sl,
				Symbol:        symbol,
				Tp:            tp,
				Type:          transactionType,
				Volume:        volume,
			},
		},
	}
}
