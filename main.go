package main

import "fmt"

// В цикле спрашиваем ввод транзакций
// Добавлять каждую введенную транзакцию в массив
// Вывести массив

func main() {
	transactions := []float64{}
	for {
		transaction := getUserInput()
		if transaction == 0 {
			fmt.Println("Транзакция не может быть равна 0")
			break
		} else {
			transactions = append(transactions, transaction)
		}
		isRepeatInput := checkUserChoice()
		if !isRepeatInput {
			break
		} else {
			continue
		}
	}
	fmt.Println(transactions)

}

func getUserInput() float64 {
	var transaction float64
	fmt.Println("Введите значение транзакции: ")
	fmt.Scan(&transaction)
	return transaction
}

func checkUserChoice() bool {
	var userChoice string
	fmt.Println("Вы хотите добавить еще транзакцию (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	} else {
		return false
	}
}
