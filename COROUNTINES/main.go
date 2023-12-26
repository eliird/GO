package main

import "fmt"

type Duration float64

func printMessage(text string){
	for i:= 0; i<10; i++{
		fmt.Println(text)
		time.Sleep(time)
	}
}


func main(){
	printMessage("Go is Great!")
}:
