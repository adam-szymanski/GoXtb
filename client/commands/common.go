package commands

import (
	"strconv"
	"time"
)

// Command describes command structure documented here: http://developers.xstore.pro/documentation/#input-data-format
type Command struct {
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments,omitempty"`
}

// Result describes result structure documented here: http://developers.xstore.pro/documentation/#output-data-format
type Result struct {
	Status     bool   `json:"status"`
	ErrorCode  string `json:"errorCode,omitempty"`
	ErrorDescr string `json:"errorDescr,omitempty"`
}

type Time time.Time

func (t *Time) UnmarshalJSON(text []byte) error {
	timestamp, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(timestamp/1000, timestamp%1000))
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixNano()/1e6, 10)), nil
}
