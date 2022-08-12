package main

import "fmt"

func main() {

	//The Calculator Project was made it switch, now i will use if and else.

	var montante float64
	var increment float64
	var decrement float64
	var op int64

	fmt.Println("How much you have on your account? ")
	fmt.Scan(&montante)

	if montante <= 0 {
		fmt.Println("Sorry but you don't have any money...")
		return
	}

	fmt.Printf("\n What do you wanna do with your %f dollars?", montante)

	for aloop := true; aloop; aloop = !false {
		fmt.Println("\nPress 1 to see your wallet, 2 to deposit, 3 to withdrawal, 4 to leave the bank system, 5 or any other number besides (1,2,3,4) if you wanna contact us")
		fmt.Scan(&op)

		if op == 1 {
			fmt.Printf("\n Your wallet has: %f dollars", montante)

		} else if op == 2 {
			fmt.Println("How much do you wanna deposit on your account? : ")
			fmt.Scan(&increment)
			montante = montante + increment
			fmt.Println("Your deposit was sucessfull! Please, press 1 to see your total in U$")

		} else if op == 3 {
			fmt.Println("How much do you wanna withdrawal on your account? : ")
			fmt.Scan(&decrement)
			montante = montante - decrement
			fmt.Println("Your withdrawal was sucessfull! Please, press 1 to see your total in U$")

		} else if op == 4 {
			fmt.Println("Exiting the system...")
			aloop = false
			return

		} else {
			fmt.Println("Our contact 'Mocked' is (55) 30 3045-9928, or by the email 'Mocked' banksystem@mail.com.br")
		}

	}

}
