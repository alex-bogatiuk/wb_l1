package main

import "fmt"

func main() {

	set := []string{"cat", "cat", "dog", "cat", "tree"} // объявляем массив со строками

	items1 := make(map[string][]string) // карта строк с массивом строк в качестве ключа

	for _, value := range set { // проходим по массиву
		if _, ok := items1[value]; ok { // если элемент уже есть в карте - продолжаем
			continue
		}

		items1[value] = []string{"собственное", "множество"} // добавляем собственное множество
	}

	fmt.Println(items1)
}
