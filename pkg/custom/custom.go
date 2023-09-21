package custom

// Custom file format, Updated: August 31, 2023 18:15
// https://support.cryptact.com/hc/en-us/articles/360002571312-Custom-File-for-any-other-trades

import (
	"math/big"
	"time"
)

type CustomData struct {
	// Trade datetime
	Timestamp time.Time
	// Transaction type
	Action Action
	// Data source name (exchange or shop name)
	Source string
	// Base currency
	// Supported cryptocurrency and fiat currencies
	// https://grid.cryptact.com/coins,
	// Unspported coins
	// https://support.cryptact.com/hc/en-us/articles/360002571312-Custom-File-for-any-other-trades#menu26
	Base string
	// Amount of change of base currency
	Volume *big.Float
	// Base currency price in Counter currency
	// When base currency is Supported Coins,
	// if thie price field is empty and counter currency is fiat,
	// crypact will look up the price instead if we have the one.
	Price *big.Float
	// Counter currency
	// Supported cryptocurrency and fiat currencies,
	// https://grid.cryptact.com/coins
	Counter string
	// Fee for the trade. Please input 0 if there was no fee for the trdes
	Fee *big.Float
	// Fee currency
	// Supported cryptocurrency and fiat currencies
	// https://grid.cryptact.com/coins
	FeeCcy string
	// Optional, Please use this field for your own purpose
	// Please do not use 「'」「"」「\」as these symbols cause errors.
	Comment string
}

// Transaction type
type Action string

// StringValue convert Action to string
func (a Action) StringValue() string {
	return string(a)
}

const (
	// Purchased at the price for the stated date and time
	ActionBuy Action = "BUY"
	// Sold at the price for the stated date and time
	ActionSell Action = "SELL"
	// Processed as a profit based on the price for the stated date and time
	ActionBonus Action = "BONUS"
	// Processed as a loss at the book price for the stated date and time
	ActionLoss Action = "LOSS"
	// Decrease the quantity without affecting the PNL
	ActionReduce Action = "REDUCE"
	// Processed as a profit based on the price for the stated date and time
	ActionStaking Action = "STAKING"
	// Processed as a profit based on the price for the stated date and time
	ActionLending Action = "LENDING"
	// Locks in the stated quantity
	ActionLend Action = "LEND"
	// Releases the stated quantity
	ActionRecover Action = "RECOVER"
	// Processed as a profit based on the price for the stated date and time
	ActionMining Action = "MINING"
	// Decrease the quantity without affecting the PNL
	ActionSendfee Action = "SENDFEE"
	// Processed as a sale at the price for the stated date and time
	ActionPay Action = "PAY"
	// Incorporate the price for the stated date and time into the book price
	ActionBorrow Action = "BORROW"
	// PNL is calculated as (price at the time of borrowing - book price) * quantity
	ActionReturn Action = "RETURN"
	// The book price for the NFT concerned is increased
	ActionLevelup Action = "LEVELUP"
	// PNL is calculated for the Fiat currency
	ActionCash Action = "CASH"
)
