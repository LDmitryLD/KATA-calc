package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Roman map[string]int

func main() {
	Roman = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите данные в формате: \"x + y\"")

		scanner.Scan()
		inp := strings.Split(scanner.Text(), " ")
		if len(inp) > 3 {
			fmt.Println("ошибка: формат математической операции не удовлетворяет заданию")
			break
		}

		if len(inp) < 3 {
			fmt.Println("ошибка: введённая строка не является математической операцией")
			break
		}

		var firstOperand string
		var operator string
		var secondOperand string

		firstOperand = inp[0]
		operator = inp[1]
		secondOperand = inp[2]

		ns, err := ArOrRom(firstOperand, secondOperand)
		if err != nil {
			fmt.Println(err)
			break
		}

		for key, value := range Roman {
			if key == firstOperand {
				firstOperand = strconv.Itoa(value)
			}
			if key == secondOperand {
				secondOperand = strconv.Itoa(value)
			}
		}

		x, _ := strconv.Atoi(firstOperand)

		y, _ := strconv.Atoi(secondOperand)

		if x > 10 || x < 1 || y > 10 || y < 1 {
			fmt.Println("ошибка: числа дложны быть от 1 до 10")
			break
		}

		if ns == "Arab" {
			result, err := CalculateAr(x, operator, y)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println(result)
		}

		if ns == "Roman" {
			result, err := CalculateRom(x, operator, y)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println(result)
		}

	}
}

func ArToRom(num int) (string, error) {

	var res string

	if num >= 40 && num < 50 {
		res += "XL"
		num -= 40
	} else if num >= 50 && num < 90 {
		res += "L"
		num -= 50
	} else if num >= 90 && num < 100 {
		res += "XC"
		num -= 90
	} else if num == 100 {
		return "C", nil
	} else if num > 100 {
		err := errors.New("ошибка: вводимое число слишком велико")
		return "", err
	} else if num == 0 {
		err := errors.New("ошибка: в римской системе счисления нету нуля")
		return "", err
	} else if num < 0 {
		err := errors.New("ошибка: в римской системе счисления нет отрицателных чисел")
		return "", err
	}

	tens := num / 10
	if tens > 0 {
		for i := 0; i < tens; i++ {
			res += "X"
		}
	}

	units := num % 10
	for rom, arab := range Roman {
		if units == arab {
			res += rom
		}
	}

	return res, nil
}

func CalculateAr(firstOperand int, operator string, secondOperand int) (res string, err error) {

	var resultOfCalculate int

	if operator == "+" {
		resultOfCalculate = firstOperand + secondOperand
		return strconv.Itoa(resultOfCalculate), nil

	} else if operator == "-" {
		resultOfCalculate = firstOperand - secondOperand
		return strconv.Itoa(resultOfCalculate), nil

	} else if operator == "*" {
		resultOfCalculate = firstOperand * secondOperand
		return strconv.Itoa(resultOfCalculate), nil

	} else if operator == "/" {
		resultOfCalculate = firstOperand / secondOperand
		return strconv.Itoa(resultOfCalculate), nil

	} else {
		err := errors.New("ошибка: некорректный оператор")
		return "", err
	}
}

func CalculateRom(firstOperand int, operator string, secondOperand int) (res string, err error) {

	var resultOfCalculate int

	if operator == "+" {
		resultOfCalculate = firstOperand + secondOperand
		resultRom, err := ArToRom(resultOfCalculate)
		if err != nil {
			return "", err
		}
		return resultRom, nil

	} else if operator == "-" {
		resultOfCalculate = firstOperand - secondOperand
		resultRom, err := ArToRom(resultOfCalculate)
		if err != nil {
			return "", err
		}
		return resultRom, nil

	} else if operator == "*" {
		resultOfCalculate = firstOperand * secondOperand
		resultRom, err := ArToRom(resultOfCalculate)
		if err != nil {
			return "", err
		}
		return resultRom, nil

	} else if operator == "/" {
		resultOfCalculate = firstOperand / secondOperand
		resultRom, err := ArToRom(resultOfCalculate)
		if err != nil {
			return "", err
		}
		return resultRom, nil

	} else {
		err := errors.New("ошибка: некорректный оператор")
		return "", err
	}
}

func ArOrRom(firstOperand, secondOperand string) (string, error) {

	var chek int

	_, ok := Roman[firstOperand]
	if !ok {
		chek += 3
	} else {
		chek += 2
	}

	_, ok = Roman[secondOperand]
	if !ok {
		chek += 3
	} else {
		chek += 2
	}

	if chek == 6 {
		return "Arab", nil
	} else if chek == 4 {
		return "Roman", nil
	} else if chek == 5 {
		return "", errors.New("ошибка: оба числа должны быть из одной системы счисления и от 1 до 10")
	}
	return "", errors.New("ошибка: некорректный ввод")
}
