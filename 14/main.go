package main

import (
	"fmt"
	"reflect"
)

func main() {

	number := 38
	str := "string"
	boolean := false
	ch := make(chan struct{})

	anyTypeArr := []any{number, str, boolean, ch}

	for _, v := range anyTypeArr {

		fmt.Println("Значение:", v)

		// Данный способ определения типа данных через switch
		switch v.(type) {
		case int:
			fmt.Println("Тип данных через switch: int")
		case string:
			fmt.Println("Тип данных через switch: string")
		case bool:
			fmt.Println("Тип данных через switch: bool")
		case chan struct{}:
			fmt.Println("Тип данных через switch: chan struct{}")
		default:
			fmt.Println("Тип данных через switch неопределен")
		}

		// Можно так же использовать пакет reflect. Но по возможности рекомендуется использовать switch
		fmt.Println("Тип данных через reflect:", reflect.TypeOf(v))
	}
}
