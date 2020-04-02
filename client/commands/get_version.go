package commands

type GetVersionReturnData struct {
	Version string
}

type GetVersionResult struct {
	Result
	ReturnData GetVersionReturnData
}

func NewGetVersionRequestCommand() *Command {
	return &Command{
		Command: "getVersion",
	}
}
