package commands

type TradeCommand int8

const (
	BUY        TradeCommand = 0 // buy
	SELL       TradeCommand = 1 // sell
	BUY_LIMIT  TradeCommand = 2 // buy limit
	SELL_LIMIT TradeCommand = 3 // sell limit
	BUY_STOP   TradeCommand = 4 // buy stop
	SELL_STOP  TradeCommand = 5 // sell stop
	BALANCE    TradeCommand = 6 // Read only. Used in getTradesHistory (http://developers.xstore.pro/documentation/#getTradesHistory) for manager's deposit/withdrawal operations (profit>0 for deposit, profit<0 for withdrawal).
	CREDIT     TradeCommand = 7 // Read only
)

type TradeType uint8

const (
	OPEN TradeType iota
	PENDING
	CLOSE
	MODIFY
	DELETE
)
