package main

import "fmt"

func main() {

	// Температурная последовательность
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту ключом-шагом 10 градусов
	temperatureGroup := make(map[int][]float32)

	for _, t := range temperature {
		// Определяем температурную группу
		group := (int(t) / 10) * 10

		// Добавляем значение в температурную группу
		temperatureGroup[group] = append(temperatureGroup[group], t)
	}

	fmt.Println(temperatureGroup)

}
