package commands

// Command describes command structure documented here: http://developers.xstore.pro/documentation/#input-data-format
type Command struct {
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments,omitempty"`
}

// Result describes result structure documented here: http://developers.xstore.pro/documentation/#output-data-format
type Result struct {
	Status     bool        `json:"status"`
	ErrorCode  string      `json:"errorCode,omitempty"`
	ErrorDescr string      `json:"errorDescr,omitempty"`
	ReturnData interface{} `json:"returnData,omitempty"`
}
