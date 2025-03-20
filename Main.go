package main

import (
	variablen "Taschenrechner/Variablen"
	"fmt"
)

func main() {

	var number1 float64
	var number2 float64
	var operator string
	var exit bool = true
	var stop int

	for exit {

		fmt.Println("*********************")
		fmt.Println("***Taschenrechner***")
		fmt.Println("*********************")
		fmt.Println("Taschenrechner(1) Beenden(2)")
		fmt.Scan(&stop)
		if stop == 2 {
			fmt.Println("Aufwiedersehen")
			break
		}
		fmt.Println("Erste Zahl")
		fmt.Scan(&number1)
		fmt.Println("Zweite Zahl")
		fmt.Scan(&number2)
		fmt.Println("Wähle den Operator (+ - * /)")
		fmt.Scan(&operator)
		switch operator {
		case "+":
			fmt.Println(variablen.Addition(number1, number2))
		case "-":
			fmt.Println(variablen.Subtraction(number1, number2))
		case "*":
			fmt.Println(variablen.Multiplikation(number1, number2))
		case "/":
			fmt.Println(variablen.Divison(number1, number2))
		default:
			fmt.Println("Ungültiger Operateor ")

		}

	}

}
