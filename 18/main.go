package main

import (
	"fmt"
	"sync"
)

// counter хранить в себе сам счетчик, мьютекс для конкурентного доступа к счетчику
// и вэйт группа для синхронизации с вызывающей горутиной
type counter struct {
	count int
	mutex sync.Mutex
	wg    sync.WaitGroup
}

func main() {

	c := newCounter()

	// Запускаем горутины
	for i := 5; i > 0; i-- {
		c.wg.Add(1)
		go c.increase()
	}

	// ждем, пока счетчик WaitGroup обнулится
	c.wg.Wait()

	fmt.Println(c.count)
}

// newCounter возвращает новый объект counter
func newCounter() *counter {
	return &counter{}
}

// increase конкурентно увеличивает счетчик на единицу
func (c *counter) increase() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
	c.wg.Done()
}
