package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Переполнение.")

	overflowesUint8 := 0
	overflowesUint16 := 0

	countUint8 := uint8(0)
	countUint16 := uint16(0)

	for i := uint32(0); i < math.MaxUint32; i++ {
		countUint8++
		countUint16++

		result1 := int(countUint8) - 1
		result2 := int(countUint16) - 1

		if result1 < 0 {
			overflowesUint8++
		}
		if result2 < 0 {
			overflowesUint16++
		}
	}

	fmt.Println("Количество переполнений чисел типа uint8 в диапазоне от 0 до uint32:", overflowesUint8)
	fmt.Println("Количество переполнений чисел типа uint16 в диапазоне от 0 до uint32:", overflowesUint16)
}
