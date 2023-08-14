package main

import "fmt"

//////////////////////////////////////////////////
// Интерфейс, с которым должна работать система //
//////////////////////////////////////////////////

type BankStatementService interface {
	LoadPaymentData() error
}

type PaymentTable struct {
}

func (pt PaymentTable) LoadPaymentData() {
	fmt.Println("Загружаем список платежей из внешней системы")
}

//////////////////////////////////
// Payment адаптируемая система //
//////////////////////////////////

type Payment struct {
}

// Специфичная конвертация в нужный формат данных
func (p Payment) ConvertToPaymentTable() PaymentTable {
	fmt.Println("Конвертируем данные в нужный формат", p)
	return PaymentTable{}
}

///////////////////////////////
// PaymentAdapter - адаптер  //
///////////////////////////////

// Адаптер реализует интерфейс Payment и является адаптером
type PaymentAdapter struct {
	payment *Payment
}

// Адаптер реализует интерфейс BankStatementService
func (adapter PaymentAdapter) LoadPaymentData() error {
	adapter.payment.ConvertToPaymentTable()
	return nil
}

// Конструктор адаптера
func NewPaymentsAdapter(p *Payment) BankStatementService {
	return &PaymentAdapter{p}
}

func main() {

	// Типичный код нашей системы
	paymnt := Payment{}

	// Создаем переходник, который будет конвертировать данные
	paymntAdapt := NewPaymentsAdapter(&paymnt)

	fmt.Println(paymntAdapt.LoadPaymentData())

}
