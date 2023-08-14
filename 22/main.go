package main

import (
	"fmt"
	"math/big"
)

// В Go операции с большими числами выполняются с помощью пакета math/big
func main() {
	// Инициализируем числа как строки, так как int64 поддерживает не больше ~9.2 квинтиллионов
	a := "700000000000000000000"
	b := "35000000000000000000"

	// A
	bigA := new(big.Int)
	bigA.SetString(a, 10)

	//B
	bigB := new(big.Int)
	bigB.SetString(b, 10)

	//Создаем промежуточную переменную, с помощью которой осуществляем операции и выводим их на экран
	result := big.NewInt(0)
	fmt.Println("Сумма:", result.Add(bigA, bigB))
	fmt.Println("Разница:", result.Sub(bigA, bigB))
	fmt.Println("Умножение:", result.Mul(bigA, bigB))
	fmt.Println("Деление:", result.Div(bigA, bigB))

}
