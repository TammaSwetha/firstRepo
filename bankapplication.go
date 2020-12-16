package main

import "fmt"

var accountNo int

type account struct {
	firstName, LastName string
	balance             int
	salary              int
}

var m map[int]account //  int account number
func main() {
	m = make(map[int]account)
	for {
		var i int
		fmt.Println("Welcome to Bank application")
		s := `1. addAccount,
		2. updateAccount,
		3. transactionAccount,
		4. viewAccount`
		fmt.Println("choose one option")
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			addAccount()
		case 2:
			fmt.Println("updateAccount()")
		case 3:
			fmt.Println("(transactionAccount()")
		case 4:
			viewAccount()
		}
	}
}
func addAccount() {
	v := account{}
	fmt.Println("What is your account number:")
	fmt.Scan(&accountNo)
	fmt.Println("First Name:")
	fmt.Scan(&v.firstName)
	fmt.Println("Last Name:")
	fmt.Scan(&v.LastName)
	fmt.Println("Balance: ")
	fmt.Scan(&v.balance)

	m[accountNo] = account{v.firstName, v.LastName, v.balance}
	fmt.Println(m[accountNo])

}

// func updateAccount() {

// }
func viewAccount() {
	fmt.Println(m[accountNo])
}
