package main

import (
	"fmt"
	"strings"
)

func someFunc() {
	// Экспертное мнение от Николая Тузова
	// 1. Локальные переменные создаются на стеке, а глобальные - на куче. Работа со стеком быстрее, чем с кучей,
	// поэтому лучше использовать локальные переменные, если это возможно.

	// 2. Строки в Go иммутабельны, то есть когда мы их изменяем, фактически создаём новую строку.
	// Поэтому, если мы хотим изменить строку, то лучше использовать []byte

	justString := []byte(createHugeString(1 << 10))
	justString = justString[:100]

	fmt.Println(string(justString))

}

func main() {
	someFunc()
}

func createHugeString(rowSize int) string {
	return strings.Repeat("g", rowSize)
}
