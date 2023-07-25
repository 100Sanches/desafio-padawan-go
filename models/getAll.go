package models

import (
	"challenge/db"
)

func (RowExchange) TableName() string {
	return "conversionsLogs"
}

func FetchAllConversions() ([]RowExchangeResponse, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	var conversions []RowExchange
	err = conn.Find(&conversions).Error
	if err != nil {
		return nil, err
	}

	var responseConversions []RowExchangeResponse
	for _, c := range conversions {
		responseConversions = append(responseConversions, RowExchangeResponse{
			ID:            c.ID,
			Amount:        c.Amount,
			From_currency: c.From_currency,
			To_currency:   c.To_currency,
			Rate:          c.Rate,
		})
	}

	return responseConversions, nil
}

func GetAllConversions() ([]RowExchangeResponse, error) {
	conversions, err := FetchAllConversions()
	if err != nil {
		return nil, err
	}

	return conversions, nil
}
