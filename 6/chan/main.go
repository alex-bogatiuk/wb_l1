package main

import (
	"fmt"
	"time"
)

func main() {

	// Объявляем канал выхода
	exitCh := make(chan struct{})
	defer close(exitCh) // закрываем канал

	// Создаем воркера
	go worker(exitCh)

	// Ожидаем 5 секунд после чего завершение программы
	select {
	case <-time.After(5 * time.Second):
		exitCh <- struct{}{}
		fmt.Println("Time Out!")
	}
}

func worker(exitCh chan struct{}) {
	for {
		select {
		// Немедленное завершение работы программы по истечении тайм-аута
		case <-exitCh:
			return
		default:
			// Воркер начинает что-то делать
			fmt.Println("worker doing  something...")
			// Делаем слип
			time.Sleep(time.Second)
			// Воркер который доделал свою работу
			fmt.Println("worker done  something...")
		}
	}
}
