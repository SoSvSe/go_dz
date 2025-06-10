package main

import "fmt"

func main() {
	const usdEur float64 = 1.1
	const usdRub float64 = 96.4
	eurRub := usdRub / usdEur
	fmt.Print(eurRub)
}
