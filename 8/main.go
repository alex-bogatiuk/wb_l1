package main

import (
	"fmt"
	"strconv"
)

func setBit(n int64, i uint, bitValue int) int64 {
	mask := int64(1) << i
	if bitValue == 0 {
		return n &^ mask
	}
	return n | mask
}

func main() {
	var num int64
	var i uint
	var bitValue int

	fmt.Print("Введите число int64: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Некорректные данные. Тип данных int64 принимает значения не больше 9223372036854775807")
		return
	}

	fmt.Print("Введите положение бита для изменения: ")
	_, err = fmt.Scanln(&i)
	if err != nil || i < 0 || i > 63 { // максимальное количество разрядов для int64 - 63
		fmt.Println("Некорректные данные. Положение бита для изменения должно быть в диапозоне от 1 до 63")
		return
	}

	fmt.Print("Введите какое значение бита нужно установить (0 or 1): ")
	_, err = fmt.Scanln(&bitValue)
	if err != nil || (bitValue != 0 && bitValue != 1) {
		fmt.Println("Некорректные данные. Значение бита может быть только 1 или 0")
		return
	}

	result := setBit(num, i, bitValue)
	fmt.Printf("Результат: %d\n", result)

	fmt.Println("Представление числа", num, "в двоичной системе:", strconv.FormatInt(num, 2))
	fmt.Println("Представление числа", result, "в двоичной системе:", strconv.FormatInt(result, 2))

}
