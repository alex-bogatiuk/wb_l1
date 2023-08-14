package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// Методы структуры Human

func (h Human) PrintName() {
	fmt.Printf("My name is %s \n", h.Name)
}

func (h Human) PrintAge() {
	fmt.Printf("%s is %d years old \n", h.Name, h.Age)
}

// Встраивание структуры Human в структуру Action

type Action struct {
	Human
}

// Произвольный метод структуры Action

func (h Human) Run() {
	fmt.Printf("%s is running \n", h.Name)
}

func main() {
	action := Action{Human{"Alexandr", 32}}
	// Вызов методов Human из структуры Human
	action.PrintName()
	action.PrintAge()
}
