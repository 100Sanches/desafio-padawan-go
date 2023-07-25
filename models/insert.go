package models

import (
	"challenge/db"
	"fmt"
	"strconv"
)

func (RegisterLog) TableName() string {
	return "conversionsLogs"
}

func RegisterConversion(registerLog RegisterLog) (uint32, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	err = conn.Create(&registerLog).Error
	if err != nil {
		return 0, err
	}

	return registerLog.ID, nil
}

func ExchangeHandlerService(params map[string]string) (Conversion, error) {
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	from := params["from"]
	to := params["to"]
	rate, _ := strconv.ParseFloat(params["rate"], 64)

	var symbol string
	switch to {
	case "USD":
		symbol = "$"
	case "BRL":
		symbol = "R$"
	case "EUR":
		symbol = "€"
	}

	convertedAmount := amount * rate
	conversion := Conversion{ConvertedAmount: convertedAmount, CurrencySymbol: symbol}
	registerLog := RegisterLog{Amount: amount, From_currency: from, To_currency: to, Rate: rate}

	idInsert, err := RegisterConversion(registerLog)
	if err != nil {
		return Conversion{}, err
	}

	fmt.Printf("Conversion record inserted successfully! Id: %d\n", idInsert)
	return conversion, nil
}
