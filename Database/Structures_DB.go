package models

import (
	"database/sql"
	"time"
)

type ClientHoldingsDB struct {
	BuyAvg   sql.NullInt64
	ClientID sql.NullString
	Qty      sql.NullInt64
	ScripID  sql.NullString
}

type ClientDetailsDB struct {
	ClientID  sql.NullString
	Contact   sql.NullString
	EKYC      sql.NullTime
	LastLogin sql.NullTime
	Name      sql.NullString
	Password  sql.NullString
	Suspended bool
}

type DealerDetailsDB struct {
	Contact   sql.NullString
	DealerID  sql.NullString
	EKYC      sql.NullTime
	LastLogin sql.NullTime
	Name      sql.NullString
	Password  sql.NullString
	Ranking   sql.NullInt64
	Suspended bool
}

type OrderDetailsDB struct {
	ClientID sql.NullString
	DateTime sql.NullTime
	OrderID  sql.NullString
	Price    sql.NullInt64
	Qty      sql.NullInt64
	ScripID  sql.NullString
	Type     bool
	UserID   sql.NullString
}

type PositionDetailsDB struct {
	ClientID sql.NullString
	DateTime sql.NullTime
	OrderID  sql.NullString
	Qty      sql.NullInt64
	ScripID  sql.NullString
	TradeID  sql.NullString
	Type     bool
	UserID   sql.NullString
}

type TradeDetailsDB struct {
	ClientID sql.NullString
	DateTime sql.NullTime
	OrderID  sql.NullString
	Price    sql.NullInt64
	Qty      sql.NullInt64
	ScripID  sql.NullString
	TradeID  sql.NullString
	Type     bool
}

type UserFundDetailsDB struct {
	AvailableCash sql.NullInt64
	UserID        sql.NullString
	UtilizedCash  sql.NullInt64
}

type UserDetailsDB struct {
	AccountOpenDate sql.NullTime
	Address         sql.NullString
	Adhar           sql.NullString
	FatherName      sql.NullString
	PAN             sql.NullString
	UserID          sql.NullString
}
