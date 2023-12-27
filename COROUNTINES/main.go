package main

import (
	"fmt"
	"time"
)

type Duration float64


func printMessageChannel(text string, channel chan string){	
	for i:= 0; i<10; i++{
		fmt.Println(text)
		time.Sleep(8888888800*time.Millisecond)
	}
	channel <- "Done!"
}
func printMessage(text string){
	for i:= 0; i<10; i++{
		fmt.Println(text)
		time.Sleep(400*time.Millisecond)
	}
}

func linearExecution(){
	
	printMessage("Go is Great!")
	printMessage("We miss you!")
}
func routineExection(){
	//parallel execution of the code
	go printMessage("Go is Great!")
	go printMessage("FME Rocks")
	go printMessage("We miss you")
	for{}//you need to keep the execution thread busy to see the output because if main ends the threads keep running but out console gets disconnetd	
}

func channelExecution(){
	var channel chan string
	go printMessageChannel("FME Rocks!", channel)
	response:= <-channel
	fmt.Println(response)
}
//How to communicate data between two routines
//using channels
func main(){
//	routineExection()
	channelExecution()
	//Buffered Channel, Capable of holding more than one values
	logs:= make(chan string, 2)
	logs<-"Hello"
	logs<-"World!"
	//To avoud deadlocks you have to close the channels before ending the program with close chan
}
