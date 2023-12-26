package data

import "fmt"

func birthday(age int){
	age++
}
func birthdayRef(age *int){
	*age++
}
func init(){
	age:= 22
	fmt.Println(age)
	fmt.Println("__________________________-")
	birthday(age)
	fmt.Println("Birthday after fixed value",(age))
	birthdayRef(&age)
	fmt.Println("Birthday after reference value",(age))
}
