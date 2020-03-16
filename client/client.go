package client

import (
	"net/url"

	"github.com/adam-szymanski/GoXtb/client/commands"
	"github.com/gorilla/websocket"
)

// Client is managing connection with XTB server.
type Client struct {
	url *url.URL
	con *websocket.Conn
}

func (c *Client) executeCommand(command *commands.Command, response interface{}) error {
	if err := c.con.WriteJSON(command); err != nil {
		return err
	}
	if err := c.con.ReadJSON(response); err != nil {
		return err
	}
	return nil
}

// NewClient is starting connection with given server.
func NewClient(u *url.URL) (*Client, error) {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &Client{url: u, con: c}, nil
}

// Close is closing connection with server.
func (c *Client) Close() error {
	return c.con.Close()
}

// COMMANDS

// Login performs login request described here: http://developers.xstore.pro/documentation/#login
func (c *Client) Login(user, password string) (*commands.LoginResult, error) {
	res := &commands.LoginResult{}
	err := c.executeCommand(commands.NewLoginCommand(user, password), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Login performs loout request described here: http://developers.xstore.pro/documentation/#logout
func (c *Client) Logout() (*commands.Result, error) {
	res := &commands.Result{}
	err := c.executeCommand(commands.NewLogoutCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
