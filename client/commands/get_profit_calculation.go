package commands

type GetProfitCalculationReturnData struct {
	Profit float32 // Profit in account currency
}

type GetProfitCalculationResult struct {
	Result
	ReturnData GetProfitCalculationReturnData
}

type getProfitCalculationRequest struct {
	ClosePrice float32 `json:"closePrice"` // Period code
	Cmd        uint8   `json:"cmd"`        // Start of chart block (rounded down to the nearest interval and excluding)
	OpenPrice  float32 `json:"openPrice"`
	Symbol     string  `json:"symbol"` // Symbol
	Volume     float32 `json:"volume"`
}

func NewGetProfitCalculationRequestCommand(closePrice float32, cmd uint8, openPrice float32, symbol string, volume float32) *Command {
	return &Command{
		Command: "getProfitCalculation",
		Arguments: &getProfitCalculationRequest{
			ClosePrice: closePrice,
			Cmd:        cmd,
			OpenPrice:  openPrice,
			Symbol:     symbol,
			Volume:     volume,
		},
	}
}
