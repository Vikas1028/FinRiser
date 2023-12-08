package main

import (
	"time"
)

type ClientHoldings struct {
	BuyAvg   int    `json:"buyAvg"`
	ClientID string `json:"clientId"`
	Qty      int    `json:"qty"`
	ScripID  string `json:"scripId"`
}

type Clients struct {
	ClientID  string    `json:"clientId"`
	Contact   string    `json:"contact"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Suspended bool      `json:"suspended"`
}

type Dealers struct {
	Contact   string    `json:"contact"`
	DealerID  string    `json:"dealerId"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Ranking   int       `json:"ranking"`
	Suspended bool      `json:"suspended"`
}

type OrderDetails struct {
	ClientID string    `json:"clientId"`
	DateTime time.Time `json:"dateTime"`
	OrderID  string    `json:"orderId"`
	Price    int       `json:"price"`
	Qty      int       `json:"qty"`
	ScripID  string    `json:"scripId"`
	Type     bool      `json:"type"`
	UserID   string    `json:"userId"`
}

type Position struct {
	ClientID string    `json:"clientId"`
	DateTime time.Time `json:"dateTime"`
	OrderID  string    `json:"orderId"`
	Qty      int       `json:"qty"`
	ScripID  string    `json:"scripId"`
	TradeID  string    `json:"tradeId"`
	Type     bool      `json:"type"`
	UserID   string    `json:"userId"`
}

type Trade struct {
	ClientID string    `json:"clientId"`
	DateTime time.Time `json:"dateTime"`
	OrderID  string    `json:"orderId"`
	Price    int       `json:"price"`
	Qty      int       `json:"qty"`
	ScripID  string    `json:"scripId"`
	TradeID  string    `json:"tradeId"`
	Type     bool      `json:"type"`
}

type UserFund struct {
	AvailableCash int    `json:"availableCash"`
	UserID        string `json:"userId"`
	UtilizedCash  int    `json:"utilizedCash"`
}

type UserDetails struct {
	AccountOpenDate time.Time `json:"accountOpenDate"`
	Address         string    `json:"address"`
	Adhar           string    `json:"adhar"`
	FatherName      string    `json:"fatherName"`
	PAN             string    `json:"pan"`
	UserID          string    `json:"userId"`
}

/*
// using []byte in place of []byte


type ClientHoldings struct {
	BuyAvg   int    `json:"buyAvg"`
	ClientID []byte `json:"clientId"`
	Qty      int    `json:"qty"`
	ScripID  []byte `json:"scripId"`
}

type Clients struct {
	ClientID  []byte    `json:"clientId"`
	Contact   []byte    `json:"contact"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Name      []byte    `json:"name"`
	Password  []byte    `json:"password"`
	Suspended bool      `json:"suspended"`
}

type Dealers struct {
	Contact   []byte    `json:"contact"`
	DealerID  []byte    `json:"dealerId"`
	EKYC      time.Time `json:"eKYC"`
	LastLogin time.Time `json:"lastLogin"`
	Name      []byte    `json:"name"`
	Password  []byte    `json:"password"`
	Ranking   int       `json:"ranking"`
	Suspended bool      `json:"suspended"`
}

type OrderDetails struct {
	ClientID []byte    `json:"clientId"`
	DateTime time.Time `json:"dateTime"`
	OrderID  []byte    `json:"orderId"`
	Price    int       `json:"price"`
	Qty      int       `json:"qty"`
	ScripID  []byte    `json:"scripId"`
	Type     bool      `json:"type"`
	UserID   []byte    `json:"userId"`
}

type Position struct {
	ClientID []byte    `json:"clientId"`
	DateTime time.Time `json:"dateTime"`
	OrderID  []byte    `json:"orderId"`
	Qty      int       `json:"qty"`
	ScripID  []byte    `json:"scripId"`
	TradeID  []byte    `json:"tradeId"`
	Type     bool      `json:"type"`
	UserID   []byte    `json:"userId"`
}

type Trade struct {
	ClientID []byte    `json:"clientId"`
	DateTime time.Time `json:"dateTime"`
	OrderID  []byte    `json:"orderId"`
	Price    int       `json:"price"`
	Qty      int       `json:"qty"`
	ScripID  []byte    `json:"scripId"`
	TradeID  []byte    `json:"tradeId"`
	Type     bool      `json:"type"`
}

type UserFund struct {
	AvailableCash int    `json:"availableCash"`
	UserID        []byte `json:"userId"`
	UtilizedCash  int    `json:"utilizedCash"`
}

type UsersDetails struct {
	AccountOpenDate time.Time `json:"accountOpenDate"`
	Address         []byte    `json:"address"`
	Adhar           []byte    `json:"adhar"`
	FatherName      []byte    `json:"fatherName"`
	PAN             []byte    `json:"pan"`
	UserID          []byte    `json:"userId"`
}

*/
