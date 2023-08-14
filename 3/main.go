package main

import (
	"fmt"
)

func main() {

	// Массив чисел
	numbers := [5]int{2, 4, 6, 8, 10}

	/////////////////////////////
	// Реализация через каналы //
	/////////////////////////////
	// Создаем буферизированный канал для синхронизации выполнения между горутинами
	doneCh := make(chan int, len(numbers))
	defer close(doneCh)

	for _, n := range numbers {

		// Запускаем в горутине анонимную функцию, вычисляющую квадрат числа n и передающую его в канал
		go func(n int) {
			doneCh <- n * n
		}(n)
	}

	sum := 0
	for range numbers {
		sum += <-doneCh
	}

	fmt.Printf("Sum of the squares of the array = %d \n", sum)

	/////////////////////////////////////
	// Реализация через sync.WaitGroup //
	/////////////////////////////////////
	// Объявляем экземпляр WaitGroup, чтобы основная горутина не завершилась раньше времени
	//wg := sync.WaitGroup{}
	//
	//// Вносим длину numbers в счетчик WaitGroup
	//wg.Add(len(numbers))
	//
	//sum := 0
	//for _, n := range numbers {
	//	// Запускаем в горутине анонимную функцию, вычисляющую квадрат числа n и выводящую его в консоль
	//	go func(n int, sum *int) {
	//		*sum += n * n
	//
	//		// декрементируем счетчик WaitGroup
	//		wg.Done()
	//	}(n, &sum)
	//}
	//
	//// ждем, пока счетчик WaitGroup обнулится
	//wg.Wait()
	//
	//fmt.Printf("Sum of the squares of the array = %d \n", sum)

}
