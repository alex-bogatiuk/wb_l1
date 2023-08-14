package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("Засыпаем")
	// Время слипа в секундах
	Sleep(5)
	fmt.Println("Проснулись")
}
