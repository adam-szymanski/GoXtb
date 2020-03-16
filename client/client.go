package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/adam-szymanski/GoXtb/client/commands"
	"github.com/gorilla/websocket"
)

// Client is managing connection with XTB server.
type Client struct {
	url     *url.URL
	con     *websocket.Conn
	verbose bool
}

func (c *Client) executeCommand(command *commands.Command, response interface{}) error {
	if c.verbose {
		commandJSON, _ := json.Marshal(command)
		fmt.Printf("Executing command %s\n", string(commandJSON))
	}
	if err := c.con.WriteJSON(command); err != nil {
		if c.verbose {
			fmt.Println("Error while sending: ", err)
		}
		return err
	}
	if err := c.con.ReadJSON(response); err != nil {
		if c.verbose {
			fmt.Println("Error while reading: ", err)
		}
		return err
	}
	if c.verbose {
		resultJSON, _ := json.Marshal(response)
		fmt.Printf("Response: %s\n", string(resultJSON))
	}
	return nil
}

// NewClient is starting connection with given server.
func NewClient(u *url.URL, verbose bool) (*Client, error) {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &Client{url: u, con: c, verbose: verbose}, nil
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

// GetAllSymbols performs getAllSymbols request described here: http://developers.xstore.pro/documentation/#getAllSymbols
func (c *Client) GetAllSymbols() (*commands.GetAllSymbolsResult, error) {
	res := &commands.GetAllSymbolsResult{}
	err := c.executeCommand(commands.NewGetAllSymbolsCommand(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
