package commands

type GetServerTimeReturnData struct {
	Time       int // Server time
	TimeString string
}

type GetServerTimeResult struct {
	Result
	ReturnData GetServerTimeReturnData
}

func NewGetServerTimeRequestCommand() *Command {
	return &Command{
		Command: "getServerTime",
	}
}
