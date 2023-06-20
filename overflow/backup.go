package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Переполнение.")

	overflowesUint8 := 0
	overflowesUint16 := 0
	uint32Count := uint32(math.MaxUint32)

	for uint32Count > 0 {
		uint32Count = uint32Count - math.MaxUint8
		overflowesUint8++
	}

	fmt.Println("Количество переполнений чисел типа uint8 в диапазоне от 0 до uint32:", overflowesUint8)

	uint32Count = math.MaxUint32

	for uint32Count > 0 {
		uint32Count = uint32Count - math.MaxUint16
		overflowesUint16++
	}

	fmt.Println("Количество переполнений чисел типа uint16 в диапазоне от 0 до uint32:", overflowesUint16)
}
