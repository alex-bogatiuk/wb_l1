package main

import "fmt"

func main() {

	// Карта для хранения множеств
	// так как в задаче не определено какой тип данных у множества, допустим что это любой comparable тип
	items := map[any]bool{}

	set1 := []any{"1", 6, "2", "69", 17.5, "5", "3"}
	set2 := []any{"1", 2, "5", 17.5, "3", 6, "9"}

	// Добавляем ключи в карту
	for _, value := range set1 {
		items[value] = false
	}

	for _, value := range set2 {
		// Если ключ существует, значит найдено пересечение - ставим значение в ИСТИНА
		if _, ok := items[value]; ok {
			items[value] = true
			continue
		}

		// Значение не найдено :(
		items[value] = false
	}

	fmt.Println(items)

}
