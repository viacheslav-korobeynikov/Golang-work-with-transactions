package main

import "fmt"

// В цикле спрашиваем ввод транзакций
// Добавлять каждую введенную транзакцию в массив
// Вывести массив
// Вывести сумму баланса в консоль

func main() {

	//tr1 := []int{1, 2, 3}
	//
	//for _, value := range tr1 {
	//	fmt.Println(value)
	//}

	transactions := []float64{}
	for {
		transaction := getUserInput()
		if transaction == 0 {
			break
		} else {
			transactions = append(transactions, transaction)

		}
	}
	fmt.Println(transactions)
	sum := calculateTransactionSum(transactions)
	fmt.Printf("Сумма транзакций составляет: %.2f", sum)

}

func getUserInput() float64 {
	var transaction float64
	fmt.Println("Введите значение транзакции (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateTransactionSum(arr []float64) float64 {
	var sum float64
	sum = 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
