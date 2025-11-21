package service

import (
	"encoding/json"
	"errors"
	"gobus/data"
	"gobus/dto"
	"gobus/model"
)

func GetTicket(req dto.Request) (model.Ticket, error) {
	var destMap map[string]float64

	err := json.Unmarshal([]byte(data.Destination), &destMap)
	if err != nil {
		return model.Ticket{}, err
	}

	price, ok := destMap[req.Destination]
	if !ok {
		return model.Ticket{}, errors.New("tujuan tidak ditemukan")
	}

	ticket := model.Ticket{
		Passenger:   req.Name,
		Destination: req.Destination,
		Price:       price,
	}

	return ticket, nil
}
