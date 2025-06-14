package main

import "fmt"

func main() {
	userInput := getUserInput()
	conversion(userInput, "eur", "rub")
	fmt.Print(userInput)
	const usdEur float64 = 1.1
	const usdRub float64 = 96.4
	eurRub := usdRub / usdEur
	fmt.Print(eurRub)
}

func getUserInput() float64 {
	var userInput float64
	fmt.Print("Ввод пользователя")
	fmt.Scan(&userInput)
	return userInput
}

func conversion(cash float64, original string, target string) {

}
