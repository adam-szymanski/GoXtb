package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/adam-szymanski/GoXtb/client"
)

var addr = flag.String("addr", "ws.xapi.pro", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/demo"}
	log.Printf("connecting to %s", u.String())

	c, err := client.NewClient(&u, true)
	if err != nil {
		panic(err)
	}
	c.Login(os.Getenv("XTB_LOGIN"), os.Getenv("XTB_PASSWORD"))
	c.GetAllSymbols()
	c.Logout()
	c.Close()
}
