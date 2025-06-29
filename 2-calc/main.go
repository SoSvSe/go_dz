package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {

		action, err := inputOperation()
		if err != nil {
			fmt.Println("Введена неправильная операция")
			break
		}
		bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println("Введите числа через запятую")
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		number := strings.ReplaceAll(str, " ", "")
		stringNumbers := strings.Split(number, ",")
		filtered := []string{}
		for _, val := range stringNumbers {
			if val != "" {
				filtered = append(filtered, val)
			}
		}
		fmt.Println(filtered)

		switch action {
		case "AVG":
			var result float64
			for _, value := range filtered {
				a, _ := strconv.ParseFloat(value, 64)

				result += a

			}

			result = result / float64(len(filtered))

			fmt.Println(result)
		case "SUM":
			var result float64
			for _, value := range filtered {
				a, err := strconv.ParseFloat(value, 64)
				if err != nil {
					a = 0
				}
				result += a
			}
			fmt.Println(result)
		case "MED":

			numbers := make([]float64, len(filtered))

			for index, value := range filtered {
				a, err := strconv.ParseFloat(value, 64)
				if err != nil {
					break
				}
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
		break
	}

}

func inputOperation() (string, error) {
	var action string
	fmt.Println("AVG - среднее")
	fmt.Println("SUM - сумма")
	fmt.Println("MED - медианна")
	fmt.Print("Введите операцию которую хотите произвести: ")
	fmt.Scan(&action)
	if action == "AVG" || action == "SUM" || action == "MED" {
		return action, nil
	}
	return action, errors.New("Неправильный ввод")
}
