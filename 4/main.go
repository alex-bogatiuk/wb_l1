package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Получение количества воркеров
	workerNum := getWorkerNum()

	// Добавляем отменяемый конекст. Отмена этого контекста высвобождает связанные с ним ресурсы,
	// поэтому код должен вызвать отмену, как только операции, выполняемые в этом контексте, завершатся.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // излишний кэнсел, но лучше перестраховаться

	// Объявляем буферизированный канал с произвольными данными
	tasksCh := make(chan any, 100)

	// Объявляем канал выхода
	exitCh := make(chan struct{})

	// Создаем воркер-пулл
	for w := workerNum; w > 0; w-- {
		go worker(ctx, w, tasksCh, exitCh)
	}

	// Передаем произвольные данные в канал
	go func() {
		for j := 1; ; j++ {
			tasksCh <- j
		}
	}()
	defer close(tasksCh)

	// Канал для перехвата сигнала прерывания от операционной системы
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh: // первый вызов отмены контекста позволяет работающим горутинам закончить свою работу
			cancel()
		}
		<-signalCh // повторный вызов сигнала вызывает незамедлительное завершение
		os.Exit(1)
	}()

	<-exitCh
}

func worker(ctx context.Context, id int, tasks <-chan any, exitCh chan<- struct{}) {
	for t := range tasks {
		// Воркер начинает что-то делать
		fmt.Println("worker", id, "doing  task", t)
		// Делаем слип, чтобы нагляднее увидеть окончание работы горутин
		time.Sleep(2 * time.Second)
		// Воркер который доделал свою работу
		fmt.Println("worker", id, "done  task", t)

		select {
		// Если контекст завершился, то вызываем return и завершаем горутину
		case <-ctx.Done():
			fmt.Println("worker", id, "completed it`s work before the end of the program")
			time.Sleep(500 * time.Millisecond)
			exitCh <- struct{}{}
			return
		default:
		}
	}
}

func getWorkerNum() int {

	fmt.Print("Number of workers: ")
	n := 0 // ожидаем integer при вводе
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	return n
}
