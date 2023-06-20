package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Введите первое число int16:")
	var firstNum int16

	for {
		var input string
		fmt.Scan(&input)

		value, err := strconv.Atoi(input) //проверка корректности введеного значения
		if value < -32768 || value > 32767 || err != nil {
			fmt.Println("Введите корректное число:")
			continue
		}

		firstNum = int16(value)
		break
	}

	fmt.Println("Введите второе число int16:")
	var secondNum int16

	for {
		var input string
		fmt.Scan(&input)

		value, err := strconv.Atoi(input) //проверка корректности введеного значения
		if value < -32768 || value > 32767 || err != nil {
			fmt.Println("Введите корректное число:")
			continue
		}

		secondNum = int16(value)
		break
	}

	var multiplied int32 = int32(firstNum) * int32(secondNum)

	if multiplied > 0 {
		if multiplied <= int32(math.MaxUint8) {
			converted := uint8(multiplied)
			fmt.Println("Результат умножения:", converted)
			fmt.Println("Минимальный тип данных для хранения результата умножения чисел: uint8")
		} else if multiplied <= int32(math.MaxUint16) {
			converted := uint16(multiplied)
			fmt.Println("Результат умножения:", converted)
			fmt.Println("Минимальный тип данных для хранения результата умножения чисел: uint16")
		} else {
			converted := uint32(multiplied)
			fmt.Println("Результат умножения:", converted)
			fmt.Println("Минимальный тип данных для хранения результата умножения чисел: uint32")
		}
	} else if multiplied < 0 {
		if multiplied <= int32(math.MaxInt8) {
			converted := int8(multiplied)
			fmt.Println("Результат умножения:", converted)
			fmt.Println("Минимальный тип данных для хранения результата умножения чисел: int8")
		} else if multiplied <= int32(math.MaxInt16) {
			converted := int16(multiplied)
			fmt.Println("Результат умножения:", converted)
			fmt.Println("Минимальный тип данных для хранения результата умножения чисел: int16")
		} else {
			converted := int32(multiplied)
			fmt.Println("Результат умножения:", converted)
			fmt.Println("Минимальный тип данных для хранения результата умножения чисел: int32")
		}
	} else {
		converted := int8(multiplied)
		fmt.Println("Результат умножения:", converted)
		fmt.Println("Минимальный тип данных для хранения результата умножения чисел: int8")
	}

}
