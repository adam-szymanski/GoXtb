package commands

type TradeRecord struct {
	ClosePrice       float32 //Close price in base currency
	CloseTime        int     //Null if order is not closed
	CloseTimeString  string  //Null if order is not closed
	Closed           bool    // Closed
	Cmd              uint8   // Operation code
	Comment          string  // Comment
	Commission       float32 // Commission in account currency, null if not applicable
	CustomComment    string  // The value the customer may provide in order to retrieve it later.
	Digits           int     // Number of decimal places
	Expiration       int     // Null if order is not closed
	ExpirationString string  // Null if order is not closed
	MarginRate       float32 // Margin rate
	Offset           int     //Trailing offset
	OpenPrice        float32 //Open price in base currency
	OpenTime         int     //Open time
	OpenTimeString   string  //Open time string
	Order            int     // Order number for opened transaction
	Order2           int     // Order number for opened transaction
	Position         int     // Order number common both for opened and closed transaction
	Profit           float32 // Profit in account currency
	Sl               float32 // Zero if stop loss is not set (in base currency)
	Storage          float32 // order swaps in account currency
	Symbol           string  // symbol name or null for deposit/withdrawal operations
	Timestamp        int     // Timestamp
	Tp               float32 // Zero if take profit is not set (in base currency)
	Volume           float32 // Volume in lots
}

type GetTradesResult struct {
	Result
	ReturnData []TradeRecord
}

type getTradesRequest struct {
	OpenedOnly bool `json:"openedOnly"` // if true then only opened trades will be returned
}

func NewGetTradesRequestCommand(openedOnly bool) *Command {
	return &Command{
		Command: "getTrades",
		Arguments: &getTradesRequest{
			OpenedOnly: openedOnly,
		},
	}
}
