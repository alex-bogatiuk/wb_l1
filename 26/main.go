package main

import (
	"fmt"
	"strings"
)

func main() {

	set1 := "abcd"
	//set2 := "abCdefAaf "
	//set3 := "aabcd"

	fmt.Println(isUnique(set1))
}

func isUnique(str string) bool {

	// Строку в руны и нижний регистр так как функция проверки должна быть регистронезависимой
	runes := []rune(strings.ToLower(str))

	// Карта для хранения множеств
	items := map[rune]uint8{}

	// Добавляем ключи в карту
	for _, value := range runes {
		if _, ok := items[value]; ok { // если элемент уже есть в карте - строка имеет не уникальные символы
			return false
		}
	}

	return true
}
