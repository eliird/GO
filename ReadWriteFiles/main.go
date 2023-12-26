package main

import (
	"RWFiles/fileutils"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Reading Writing a File")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	
	root_path, _ := os.Getwd()
	file_path := root_path + "/data/test.txt"
	c, err := fileutils.ReadTextFile(file_path)//paths need to be absolute 
	if err == nil{
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		fileutils.WriteToFile(file_path+"-output.txt", newContent)
	}else{
		fmt.Printf("ERROR %v", err)
	}
	//Converting float to a strin
}
