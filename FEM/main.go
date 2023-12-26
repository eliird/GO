package main

import (
	"fmt"
	"server/data"
	"time"
)

func main(){
	kyle := data.NewInstructor("kyle", "somthing")
	fmt.Println(kyle.FirstName)
	fmt.Println(kyle.Print())
	fmt.Println("________________________________________________----")

	goCourse := data.Course{Id:2, Name:"Go Fundamentals", Instuctor: kyle}
	//fmt.Printf("%v", goCourse)
	courseString := goCourse.String() 
	fmt.Println(courseString)

	swiftWS := data.Workshop{Course: data.Course{Name:"Swift for IOS", Instuctor:kyle}, Date: time.Now()}
	fmt.Println("%v", swiftWS)
	
	swiftWSEmbedded := data.EmbeddedWorkshop{}//embedded workshop can be used to access property with single dot than dpib;e dpt
/*
If I try to create an array trying to hold a course and a workshop I cannot do it without an interfac
Interface work using the signature of the function
data:= [2]data.Workshop
data[0] = goCourse //this will be invalid becuase they are different structs
data[1] = swiftWS
*/
	//Can put all the structs that have the signable function in an array
	data := [2]data.Signable{}
	data[0] = goCourse
	data[1] = swiftWS


} 
