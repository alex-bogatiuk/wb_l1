package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// инициализируем пустую карту
	m := make(map[int]string)

	// создаем WaitGroup
	wg := new(sync.WaitGroup)
	// и Mutex
	mutex := new(sync.Mutex)

	counter := 1

	// запускаем горутины и передаем в них мапу, ссылки на WaitGroup, Mutex и counter
	for i := 0; i < 10; i++ {
		// увеличиваем счетчик WaitGroup
		wg.Add(1)

		go worker(m, wg, mutex, &counter)
	}

	// ждем обнуления счетчика WaitGroup
	wg.Wait()
	fmt.Println(m)
}

func worker(m map[int]string, wg *sync.WaitGroup, mutex *sync.Mutex, counter *int) {
	// запускаем цикл, в котором будем записывать содержимое слайса в мапу
	for i := 0; i < 10; i++ {
		// Генерируем случайный символ
		randomInt := rand.Intn(100)
		// блокируем Mutex, чтобы остальные горутины не могли в него писать
		mutex.Lock()
		// добавляем в мапу запись: ключ - число i, значение - i-й элемент в слайсе
		m[*counter] = string(randomInt)
		*counter++
		// разблокируем Mutex
		mutex.Unlock()
	}
	// уменьшаем счетчик WaitGroup
	wg.Done()
}
