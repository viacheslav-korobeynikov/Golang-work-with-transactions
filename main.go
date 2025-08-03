package main

import "fmt"

// В цикле спрашиваем ввод транзакций
// Добавлять каждую введенную транзакцию в массив
// Вывести массив

func main() {
	transactions := []int{}
	for {
		transactions = append(transactions, getUserInput())
		isRepeatInput := checkUserChoice()
		fmt.Println(transactions)
		if !isRepeatInput {
			break
		} else {
			continue
		}
	}

}

func getUserInput() int {
	var transaction int
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
