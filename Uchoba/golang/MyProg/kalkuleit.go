package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Print("Привет, это мой калькулятор, введи число, оператор и число.\nВведи первое чилсо: ")
	var inpunUserNumA int
	fmt.Scan(&inpunUserNumA)
	fmt.Print("Введи оператор (+ - * /): ")
	var inpunUserOperator string
	fmt.Scan(&inpunUserOperator)
	fmt.Print("Введи второе число: ")
	var inpunUserNumB int 
	fmt.Scan(&inpunUserNumB)

	var resul int
	switch inpunUserOperator {
	case "+":
		resul = kalkuleit(inpunUserNumA, inpunUserNumB, inpunUserOperator)
	case "-":
		resul = kalkuleit(inpunUserNumA, inpunUserNumB, inpunUserOperator)
	case "*":
		resul = kalkuleit(inpunUserNumA, inpunUserNumB, inpunUserOperator)
	case "/":
		if inpunUserOperator == "/" {
			if inpunUserNumA == 0 || inpunUserNumB == 0 {
			fmt.Println("Деление на 0 невозможно")
			os.Exit(0)
			}
		}
		resul = kalkuleit(inpunUserNumA, inpunUserNumB, inpunUserOperator)
	default:
		fmt.Println("Не верный ввод рператора")
		os.Exit(0)
	}


	fmt.Printf("%d %s %d = %d\n", inpunUserNumA, inpunUserOperator, inpunUserNumB, resul)
}

func kalkuleit(numA, numB int, operator string) int {
	switch operator {
	case "+":
		return numA + numB
	case "-":
		return numA - numB
	case "*":
		return numA * numB
	default:
		return numA / numB
	}
}