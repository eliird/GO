package main

func printData(){
	print("Hello ")
	print("World!")
}

func add(a int, b int) int{
	return a+b
}    

func addAndSubtract(a int, b int)(int, int){
	return a+b, a-b
}

func increment(x *int){
	*x++
}
//Design Pattern when writing a go functions
/*

func readUser(id int)(user, err){
	if ok{
		return user, nil
	}else{
		return nil, errorDetails
	}
}
*/
