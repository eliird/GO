package fileutils

import "os"

func ReadTextFile(filename string)(string, error){
	content, err := os.ReadFile(filename)
	if err!=nil{
		//error in the read
		return "", err
	}else{
		//file read
		return string(content), err
	}

}
