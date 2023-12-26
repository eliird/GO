package main

import (
	"firstModule/data"
	"fmt"
)
var url = "https://frontendmasters.com" //cannot declrare the variable globally with := notation need to specify var or const

func main(){

	var x int = 10 //var keyword assigned the null values null value for int is 0
	var y int //gets assigned a 0 value
	
	const z int = 2 //cannot be mutated in runtime
	defer fmt.Println("Bye!!")
	defer fmt.Println("Good!")
	//The above statement defers the execution of the statement to the end of the function
	//defer statement execute even when panic is hit
	print("__________________________\n--")
	fmt.Println("Printing using fmt")// the capital at the start of the function name is used for functions that need to be exported outside packages i.e. is Public
	//print does not work for all systems so use the fmt package
	fmt.Println("Hello from the module!\n", x, y, z, url)
	fmt.Println(data.MaxSpeed)
	fmt.Printf("Hello %v World %s, %v\n", x, "***",url)	
	
	/* 

	panic("Something Paniced")//stops execution at this point
	[5]int declares the array
	[]int declares a slice, slice is basically a dynamicaaly array
	map[key]value is dict/map
	*/
}
