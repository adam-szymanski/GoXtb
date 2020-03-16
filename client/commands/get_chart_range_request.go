package commands

import "time"

type ChartRangeInfoRecord struct {
	ChartLastInfoRecord
	End   Time `json:"end"`             // End of chart block (rounded down to the nearest interval and excluding)
	Ticks int  `json:"ticks,omitempty"` // Number of ticks needed, this field is optional
}

type getChartRangeRequest struct {
	Info ChartRangeInfoRecord `json:"info"`
}

func NewGetChartRangeRequestCommand(period Period, start time.Time, symbol string, end time.Time, ticks int) *Command {
	return &Command{
		Command: "getChartRangeRequest",
		Arguments: &getChartRangeRequest{
			Info: ChartRangeInfoRecord{
				ChartLastInfoRecord: ChartLastInfoRecord{
					Period: period,
					Start:  Time(start),
					Symbol: symbol,
				},
				End:   Time(end),
				Ticks: ticks,
			},
		},
	}
}
