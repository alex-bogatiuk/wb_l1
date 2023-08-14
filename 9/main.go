package main

import (
	"fmt"
	"sync"
)

func main() {

	// создаем слайс с числами
	numbers := []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 21, 22, 23}

	// создаем канал для чисел
	numCh := make(chan int)
	// и канал для результата умножения
	multiplicationCh := make(chan int)

	// создаем экземпляр WaitGroup
	wg := new(sync.WaitGroup)

	// увеличиваем счетчик WaitGroup на 2
	wg.Add(2)

	// Анонимная функция пишет числа в канал
	go func() {
		// запускаем цикл и перебираем слайс с числами
		for _, n := range numbers {
			// передаем числа в канал для чисел
			numCh <- n
		}
		close(numCh)
	}()

	go multiplyTheNumbers(wg, numCh, multiplicationCh)
	go printNumbers(wg, multiplicationCh)

	// ждем обнуления счетчика WaitGroup
	wg.Wait()
}

func multiplyTheNumbers(wg *sync.WaitGroup, numCh <-chan int, squareCh chan int) {
	// в цикле читаем канал с числами
	for n := range numCh {
		// полученное число умножаем на 2
		s := n * 2
		// и передаем в канал для результата умножения
		squareCh <- s
	}
	// после закрытия канала с числами закрываем канал для результатов умножения
	close(squareCh)
	// и уменьшаем счетчик WaitGroup
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup, ch <-chan int) {
	// в цикле читаем канал с результатами умножения
	for n := range ch {
		// полученные числа выводим в консоль
		fmt.Println(n)
	}
	// после закрытия канала уменьшаем счетчик WaitGroup
	wg.Done()
}
