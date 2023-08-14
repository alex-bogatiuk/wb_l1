package main

import (
	"fmt"
)

func main() {
	phrase := "главрыба"
	fmt.Println(reverse(phrase))
}

func reverse(str string) string {
	// Преобразовываем строку в массив рун, так как строки иммутабельны
	runes := []rune(str)

	// Идем циклом до середины
	for i := 0; i < len(runes)/2; i++ {
		// смещение с конца
		j := len(runes) - i - 1
		// Меняем символы местами
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
