package commands

// Login describes JSON login command documented here: http://developers.xstore.pro/documentation/#login
type loginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	AppName  string `json:"appName,omitempty"`
}

type LoginResult struct {
	Result
	StreamSessionId string `json:"streamSessionId"`
}

func NewLoginCommand(user, password string) *Command {
	return &Command{
		Command: "login",
		Arguments: &loginRequest{
			UserId:   user,
			Password: password,
			AppName:  "GoXTB",
		},
	}
}
