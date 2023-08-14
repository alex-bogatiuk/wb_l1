package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// создаем экземпляр WaitGroup
	wg := new(sync.WaitGroup)

	// увеличиваем счетчик WaitGroup на 1
	wg.Add(1)

	// запускаем горутину и передаем в нее ссылку на WaitGroup
	go worker(wg)

	// ждем, когда обнулится счетчик WaitGroup
	wg.Wait()
}

func worker(wg *sync.WaitGroup) {

	// Воркер начинает что-то делать
	fmt.Println("worker doing  something...")
	// Делаем слип
	time.Sleep(time.Second)
	// Воркер который доделал свою работу
	fmt.Println("worker done  something...")

	time.Sleep(5 * time.Second)
	fmt.Println("Time Out!")

	// Уменьшаем счетчик вэйт-группы
	wg.Done()

}
