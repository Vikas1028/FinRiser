package main

import (
	"time"
)

type Clients struct {
	ClientID  string    `json:"clientId"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Suspended bool      `json:"suspended"`
}

type UsersDetails struct {
	UserID          string    `json:"userId"`
	Adhar           string    `json:"adhar"`
	PAN             string    `json:"pan"`
	FatherName      string    `json:"fatherName"`
	Address         string    `json:"address"`
	Photo           []byte    `json:"photo,omitempty"`
	AccountOpenDate time.Time `json:"accountOpenDate"`
}

type Dealers struct {
	DealerID  string    `json:"dealerId"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Suspended bool      `json:"suspended"`
	Ranking   int       `json:"ranking"`
}

type ClientHoldings struct {
	ClientID string `json:"clientId"`
	ScripID  string `json:"scripId"`
	Qty      int    `json:"qty"`
	BuyAvg   int    `json:"buyAvg"`
}

type UserFund struct {
	UserID        string `json:"userId"`
	AvailableCash int    `json:"availableCash"`
	UtilizedCash  int    `json:"utilizedCash"`
}

type OrderDetails struct {
	OrderID  string    `json:"orderId"`
	UserID   string    `json:"userId"`
	Type     bool      `json:"type"`
	Qty      int       `json:"qty"`
	DateTime time.Time `json:"dateTime"`
	ClientID string    `json:"clientId"`
	ScripID  string    `json:"scripId"`
	Price    int       `json:"price"`
}

type Trade struct {
	TradeID  string    `json:"tradeId"`
	OrderID  string    `json:"orderId"`
	Type     bool      `json:"type"`
	Qty      int       `json:"qty"`
	DateTime time.Time `json:"dateTime"`
	ClientID string    `json:"clientId"`
	ScripID  string    `json:"scripId"`
	Price    int       `json:"price"`
}

type Position struct {
	UserID   string    `json:"userId"`
	Type     bool      `json:"type"`
	Qty      int       `json:"qty"`
	DateTime time.Time `json:"dateTime"`
	ClientID string    `json:"clientId"`
	ScripID  string    `json:"scripId"`
	TradeID  string    `json:"tradeId"`
	OrderID  string    `json:"orderId"`
}

/*
// using []byte in place of string

type Clients struct {
	ClientID  []byte    `json:"clientId"`
	Password  []byte    `json:"password"`
	Name      []byte    `json:"name"`
	Contact   []byte    `json:"contact"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Suspended bool      `json:"suspended"`
}

type UsersDetails struct {
	UserID          []byte    `json:"userId"`
	Adhar           []byte    `json:"adhar"`
	PAN             []byte    `json:"pan"`
	FatherName      []byte    `json:"fatherName"`
	Address         []byte    `json:"address"`
	Photo           []byte    `json:"photo,omitempty"`
	AccountOpenDate time.Time `json:"accountOpenDate"`
}

type Dealers struct {
	DealerID  []byte    `json:"dealerId"`
	Password  []byte    `json:"password"`
	Name      []byte    `json:"name"`
	Contact   []byte    `json:"contact"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Suspended bool      `json:"suspended"`
	Ranking   int       `json:"ranking"`
}

type ClientHoldings struct {
	ClientID []byte `json:"clientId"`
	ScripID  []byte `json:"scripId"`
	Qty      int    `json:"qty"`
	BuyAvg   int    `json:"buyAvg"`
}

type UserFund struct {
	UserID        []byte `json:"userId"`
	AvailableCash int    `json:"availableCash"`
	UtilizedCash  int    `json:"utilizedCash"`
}

type OrderDetails struct {
	OrderID  []byte    `json:"orderId"`
	UserID   []byte    `json:"userId"`
	Type     bool      `json:"type"`
	Qty      int       `json:"qty"`
	DateTime time.Time `json:"dateTime"`
	ClientID []byte    `json:"clientId"`
	ScripID  []byte    `json:"scripId"`
	Price    int       `json:"price"`
}

type Trade struct {
	TradeID  []byte    `json:"tradeId"`
	OrderID  []byte    `json:"orderId"`
	Type     bool      `json:"type"`
	Qty      int       `json:"qty"`
	DateTime time.Time `json:"dateTime"`
	ClientID []byte    `json:"clientId"`
	ScripID  []byte    `json:"scripId"`
	Price    int       `json:"price"`
}

type Position struct {
	UserID   []byte    `json:"userId"`
	Type     bool      `json:"type"`
	Qty      int       `json:"qty"`
	DateTime time.Time `json:"dateTime"`
	ClientID []byte    `json:"clientId"`
	ScripID  []byte    `json:"scripId"`
	TradeID  []byte    `json:"tradeId"`
	OrderID  []byte    `json:"orderId"`
}
*/
