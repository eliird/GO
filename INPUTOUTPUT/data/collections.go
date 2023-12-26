package data

var Countries [10]string
var Slice []int
var slice []int //cannot be accessed from outside the package 
var Codes map[string]int
//Init functions execute when the package is built they are run first before the main is run
//Can be used to initializethe variables
//The init in this file executes when something gets called from this file
//Can have multiple inits across multiple files
func init(){
	Countries[0] = "Pakistan"
	Countries[1] = "Somehting"
	Countries[8] = "Something else"
}
func init(){
	Slice = append(Slice, 9)
}
