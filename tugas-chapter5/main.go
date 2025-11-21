package main

import (
	"bufio"
	"fmt"
	"gobus/dto"
	"gobus/handler"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== GoBus Ticket Checker ===")

	fmt.Print("Masukkan nama penumpang   : ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	fmt.Print("Masukkan kota tujuan      : ")
	destInput, _ := reader.ReadString('\n')
	destInput = strings.TrimSpace(destInput)

	req := dto.NewRequest(nameInput, destInput)

	resp := handler.ProcessTicket(req)

	handler.PrintResponse(resp)
}
