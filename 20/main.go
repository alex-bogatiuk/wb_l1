package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "Lorem ipsum dolor sit amet consectetur adipiscing elit Fusce ligula diam hendrerit sed vestibulum " +
		"non dapibus et orci Nam pretium eros malesuada condimentum dignissim nulla dui mollis arcu nec aliquet"

	fmt.Println(reverse(phrase))

}

func reverse(str string) string {

	// Преобразовываем строку в слайс
	strSlice := strings.Split(str, " ")

	// Идем циклом до середины
	for i := 0; i < len(strSlice)/2; i++ {
		// смещение с конца
		j := len(strSlice) - i - 1
		// Меняем слова местами
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}

	return strings.Join(strSlice, " ")
}
