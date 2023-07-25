package models

import "time"

type RowExchange struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
}

type RowExchangeResponse struct {
	ID            uint    `json:"id"`
	Amount        float64 `json:"amount"`
	From_currency string  `json:"from_currency"`
	To_currency   string  `json:"to_currency"`
	Rate          float64 `json:"rate"`
}

type RegisterLog struct {
	ID            uint32 `gorm:"primaryKey"`
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
}

type Conversion struct {
	ConvertedAmount float64 `json:"valorConvertido"`
	CurrencySymbol  string  `json:"simboloMoeda"`
}
