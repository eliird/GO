package data

import "fmt"

func(i Instructor)Print() string{

	return fmt.Sprintf("%v, %v", i.FirstName,i.LastName)
}
