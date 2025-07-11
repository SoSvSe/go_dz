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
		filtered, err := inputList()
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println(filtered)

		switch action {
		case "AVG":
			var result float64
			for _, value := range filtered {

				result += float64(value)

			}

			result = result / float64(len(filtered))

			// fmt.Println(result)
		case "SUM":
			var result float64
			for _, value := range filtered {

				result += float64(value)
			}
			fmt.Println(result)
		case "MED":

			numbers := make([]float64, len(filtered))

			for index, value := range filtered {

				numbers[index] = float64(value)
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

func inputList() ([]float64, error) {
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

	result := []float64{}
	for _, numStr := range filtered {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, errors.New("в списке есть недопустимое значение: " + numStr)
		}
		result = append(result, num)

	}
	return result, nil
}
