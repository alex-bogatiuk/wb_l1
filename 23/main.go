package main

import "fmt"

func main() {

	a := []string{"A", "B", "C", "D", "E"}

	fmt.Println(deleteElem(a, 4))

}

func deleteElem(a []string, i int) []string {
	a[i] = a[len(a)-1]  // копируем последний элемент в индекс i
	a[len(a)-1] = ""    // устанавливаем нулевое значение в последний элемент
	return a[:len(a)-1] // отрезаем последний элемент
}
