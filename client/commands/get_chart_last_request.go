package commands

import "time"

type RateInfoRecord struct {
	Close     float64 // Value of close price (shift from open price)
	Ctm       Time    // Candle start time in CET / CEST time zone (see Daylight Saving Time, DST)
	CtmString string  // String representation of the 'ctm' field
	High      float64 // Highest value in the given period (shift from open price)
	Low       float64 // Lowest value in the given period (shift from open price)
	Open      float64 // Open price (in base currency * 10 to the power of digits)
	Vol       float64 // Volume in lots
}

type GetCharResultReturnData struct {
	Digits    int
	RateInfos []RateInfoRecord
}

type GetChartResult struct {
	Result
	ReturnData GetCharResultReturnData
}

type ChartLastInfoRecord struct {
	Period Period `json:"period"` // Period code
	Start  Time   `json:"start"`  // Start of chart block (rounded down to the nearest interval and excluding)
	Symbol string `json:"symbol"` // Symbol
}

type getChartLastRequest struct {
	Info ChartLastInfoRecord `json:"info"`
}

type Period int

const (
	PERIOD_M1  Period = 1
	PERIOD_M5  Period = 5
	PERIOD_M15 Period = 15
	PERIOD_M30 Period = 30
	PERIOD_H1  Period = 60
	PERIOD_H4  Period = 240
	PERIOD_D1  Period = 1440
	PERIOD_W1  Period = 10080
	PERIOD_MN1 Period = 43200
)

func NewGetChartLastRequestCommand(period Period, start time.Time, symbol string) *Command {
	return &Command{
		Command: "getChartLastRequest",
		Arguments: &getChartLastRequest{
			Info: ChartLastInfoRecord{
				Period: period,
				Start:  Time(start),
				Symbol: symbol,
			},
		},
	}
}
