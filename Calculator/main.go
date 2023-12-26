package main

import "fmt"

func main(){

	var operation string
	var number1, number2 int

	fmt.Println("Calcutaot GO")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Which Operation do you want to perform?\n1-add\n2-sub\n3-mul\n4-div")
	fmt.Scanf("%s", &operation)//passed by the reference because everything copied by default
	fmt.Println("Enter the first Number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter the second Number")
	fmt.Scanf("%d", &number2)
	
	switch operation{
	case "add":
		fmt.Println(number1+number2)
	case "sub":
		fmt.Println(number1-number2)
	case "mul":
		fmt.Println(number1*number2)
	case "div":
		fmt.Println(number1/number2)
	}

}
