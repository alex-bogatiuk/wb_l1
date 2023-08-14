package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Получение количества секунд через которые программа должна завершиться
	secondsValue := getSecondsToEnd()

	// Добавляем отменяемый конекст. Отмена этого контекста высвобождает связанные с ним ресурсы,
	// поэтому код должен вызвать отмену, как только операции, выполняемые в этом контексте, завершатся.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(secondsValue))
	defer cancel() // отложенное закрытие

	// Объявляем буферизированный канал с произвольными данными
	tasksCh := make(chan any)
	defer close(tasksCh) // отложенное закрытие

	// Объявляем канал выхода
	exitCh := make(chan struct{})
	defer close(exitCh) // отложенное закрытие

	// Создаем горутину читающую из канала
	go worker(ctx, tasksCh, exitCh)

	// Передаем произвольные данные в канал
	go func() {
		for j := 1; ; j++ {
			tasksCh <- j
		}
	}()

	<-exitCh
}

func worker(ctx context.Context, tasks <-chan any, exitCh chan<- struct{}) {
	for t := range tasks {
		// Воркер начинает что-то делать
		fmt.Println("worker doing  task", t)
		// Делаем слип, чтобы нагляднее увидеть окончание работы горутин
		time.Sleep(2 * time.Second)
		// Воркер который доделал свою работу
		fmt.Println("worker done  task", t)

		select {
		case <-ctx.Done():
			fmt.Println("worker completed it`s work before the end of the program")
			time.Sleep(500 * time.Millisecond)
			exitCh <- struct{}{}
			return
		default:
		}
	}
}

func getSecondsToEnd() int {

	fmt.Print("Seconds to end the program: ")
	n := 0 // ожидаем integer при вводе
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	return n
}
