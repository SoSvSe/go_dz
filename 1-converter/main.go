package main

import "fmt"

func main() {
	for {
		var moneyName string
		var moneyCOnvert string
		moneyName = getUserNameMoney()
		if moneyName != "u" && moneyName != "e" && moneyName != "p" {

			fmt.Println("Не правильно ввели валюту")
			continue
		}
		userInput := getUserInput()
		if userInput == 0 {
			fmt.Println("Введено неправильное или нулевое значение")
			continue
		}
		moneyCOnvert = moneyNameConvert(moneyName)
		// fmt.Println(moneyCOnvert)
		if moneyCOnvert != "u" && moneyCOnvert != "e" && moneyCOnvert != "p" {
			fmt.Println("Введено неправильное значение для конвертации")
			continue
		}
		cash := conversion(userInput, moneyName, moneyCOnvert)
		fmt.Println(cash)
		break
	}

	// conversion(userInput, "eur", "rub")
	// fmt.Print(userInput)
	// const usdEur float64 = 1.1
	// const usdRub float64 = 96.4
	// eurRub := usdRub / usdEur
	// fmt.Print(eurRub)

}

func getUserNameMoney() string {
	var moneyName string
	fmt.Println("Введите:")
	fmt.Println("u - если хотите конвертировать доллары")
	fmt.Println("e- если хотите конвертировать евро")
	fmt.Println("p - если хотите конвертировать рубли")
	fmt.Scan(&moneyName)
	// fmt.Println(moneyName)
	return moneyName
}

func getUserInput() float64 {
	var userInput float64
	fmt.Print("Ввод число для конвертации : ")
	fmt.Scan(&userInput)
	return userInput
}

func conversion(cash float64, original string, target string) float64 {
	var convertCash float64
	switch original {
	case "u":
		switch target {
		case "e":
			convertCash = cash * 0.9
		case "p":
			convertCash = cash * 100
		}
	case "e":
		switch target {
		case "u":
			convertCash = cash * 1.1
		case "p":
			convertCash = cash * 110
		}
	case "p":
		// fmt.Println("1")
		switch target {
		case "u":
			convertCash = cash * 100
		case "e":
			// fmt.Println("!")
			convertCash = cash * 110
		}
	}
	return convertCash
}

func moneyNameConvert(moneyName string) string {
	var moneyNameConvert string
	switch moneyName {
	case "u":
		fmt.Println("Введите валюту в которую нужно конвертировать: ")
		fmt.Println("E - евро")
		fmt.Println("Р - рубль")
		fmt.Scan(&moneyNameConvert)
		if moneyName == "e" && moneyName == "p" {
			moneyNameConvert = ""
		}
	case "e":
		fmt.Println("Введите валюту в которую нужно конвертировать: ")
		fmt.Println("U - долар")
		fmt.Println("Р - рубль")
		fmt.Scan(&moneyNameConvert)
		if moneyName == "u" && moneyName == "p" {
			moneyNameConvert = ""
		}
	case "p":
		fmt.Println("Введите валюту в которую нужно конвертировать: ")
		fmt.Println("E - евро")
		fmt.Println("U - долар")
		fmt.Scan(&moneyNameConvert)
		if moneyName == "e" && moneyName == "u" {
			moneyNameConvert = ""
		}
	}
	// fmt.Println(moneyNameConvert)
	return moneyNameConvert
}
