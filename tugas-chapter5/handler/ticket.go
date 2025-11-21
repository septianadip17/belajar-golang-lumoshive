package handler

import (
	"fmt"
	"gobus/dto"
	"gobus/service"
)

func ProcessTicket(req dto.Request) dto.Response {
	ticket, err := service.GetTicket(req)
	if err != nil {
		return dto.Response{
			Passenger:   req.Name,
			Destination: req.Destination,
			Price:       0,
			Message:     err.Error(),
		}
	}

	return dto.Response{
		Passenger:   ticket.Passenger,
		Destination: ticket.Destination,
		Price:       ticket.Price,
		Message:     "OK",
	}
}

func PrintResponse(resp dto.Response) {
	fmt.Println()
	fmt.Println("=== Harga Tiket ===")
	fmt.Println("Penumpang :", resp.Passenger)
	fmt.Println("Tujuan    :", resp.Destination)

	if resp.Message != "OK" {
		fmt.Println("Info      :", resp.Message)
	} else {
		fmt.Printf("Harga     : Rp %.2f\n", resp.Price)
	}
	fmt.Println("===================")
}
