package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var action string
	var str string

	fmt.Println("AVG - среднее")
	fmt.Println("SUM - сумма")
	fmt.Println("MED - медианна")
	fmt.Print("Введите операцию которую хотите произвести: ")
	fmt.Scanln(&action)
	fmt.Println("Введите числа через запятую")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	number := strings.ReplaceAll(str, " ", "")
	stringNumbers := strings.Split(number, ",")

	fmt.Println(stringNumbers)

	switch action {
	case "AVG":
		var result float64
		for _, value := range stringNumbers {
			a, _ := strconv.ParseFloat(value, 64)

			result += a

		}

		result = result / float64(len(stringNumbers))

		fmt.Println(result)
	case "SUM":
		var result float64
		for _, value := range stringNumbers {
			a, _ := strconv.ParseFloat(value, 64)
			result += a
		}
		fmt.Println(result)
	case "MED":

		numbers := make([]float64, len(stringNumbers))
		// Нужно отсортировать массив
		for index, value := range stringNumbers {
			a, _ := strconv.ParseFloat(value, 64)
			numbers[index] = a
		}

		if len(numbers)%2 == 0 {

			fmt.Println((numbers[(len(numbers)/2)-1] + numbers[len(numbers)/2]) / 2.0)
		} else {
			fmt.Println(numbers[len(numbers)/2])
		}
	default:
		fmt.Println("Введено неправильное действие")
	}
}

func ParseFloat(value string, i int) (any, any) {
	panic("unimplemented")
}
