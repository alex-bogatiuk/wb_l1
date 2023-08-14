package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создаем экземпляр контекста с тайм-аутом в 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// запускаем горутину и передаем ей контекст
	go worker(ctx)

	// Немедленное завершение работы программы по истечении тайм-аута
	<-ctx.Done()

	fmt.Println("Time Out!")
}

func worker(ctx context.Context) {
	for {
		select {
		// Немедленное завершение работы программы по истечении тайм-аута
		case <-ctx.Done():
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
