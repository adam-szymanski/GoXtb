package commands

func NewGetAllSymbolsCommand() *Command {
	return &Command{
		Command: "getAllSymbols",
	}
}

type SymbolRecord struct {
	Ask                float64 // Ask price in base currency
	Bid                float64 // Bid price in base currency
	CategoryName       string  // Category name
	ContractSize       int     // Size of 1 lot
	Currency           string  // Currency
	CurrencyPair       bool    // Indicates whether the symbol represents a currency pair
	CurrencyProfit     string  // The currency of calculated profit
	Description        string  // Description
	Expiration         *Time   // Null if not applicable
	GroupName          string  // Symbol group name
	High               float64 // The highest price of the day in base currency
	InitialMargin      int     // nitial margin for 1 lot order, used for profit/margin calculation
	instantMaxVolume   int     // Maximum instant volume multiplied by 100 (in lots)
	Leverage           float64 // Symbol leverage
	LongOnly           bool    // Long only
	LotMax             float64 // Maximum size of trade
	LotMin             float64 // Minimum size of trade
	LotStep            float64 // A value of minimum step by which the size of trade can be changed (within lotMin - lotMax range)
	Low                float64 // The lowest price of the day in base currency
	MarginHedged       int     // Used for profit calculation
	MarginHedgedStrong bool    // For margin calculation
	MarginMaintenance  *int    // For margin calculation, null if not applicable
	MarginMode         int     // For margin calculation
	Percentage         float64 // Percentage
	PipsPrecision      int     // int of symbol's pip decimal places
	Precision          int     // Number of symbol's price decimal places
	ProfitMode         int     // For profit calculation
	QuoteID            int     // Source of price
	ShortSelling       bool    // Indicates whether short selling is allowed on the instrument
	SpreadRaw          float64 // The difference between raw ask and bid prices
	SpreadTable        float64 // Spread representation
	Starting           *Time   // Null if not applicable
	StepRuleID         int     // Appropriate step rule ID from getStepRules  command response
	StopsLevel         int     // Minimal distance (in pips) from the current price where the stopLoss/takeProfit can be set
	SwapRollover3Days  int     `json:"swap_rollover3days"` // Time when additional swap is accounted for weekend
	SwapEnable         bool    // Indicates whether swap value is added to position on end of day
	SwapLong           float64 // Swap value for long positions in pips
	SwapShort          float64 // Swap value for short positions in pips
	SwapType           int     // Type of swap calculated
	Symbol             string  // Symbol name
	TickSize           float64 // Smallest possible price change, used for profit/margin calculation, null if not applicable
	TickValue          float64 // Value of smallest possible price change (in base currency), used for profit/margin calculation, null if not applicable
	Time               Time    // Ask & bid tick time
	TimeString         string  // Time in String
	TrailingEnabled    bool    // Indicates whether trailing stop (offset) is applicable to the instrument.
	Type               int     // Instrument class number
}

type GetAllSymbolsResult struct {
	Result
	ReturnData []*SymbolRecord `json:"returnData"`
}
