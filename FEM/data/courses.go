package data

import "fmt"

type Duration float32

type Course struct{
	Id int
	Name string
	Slug string
	Legacy bool 
	Duration Duration
	Instuctor Instructor
}

func (c Course) String() string{
	return fmt.Sprintf("____%v-----(%v)\n", c.Name, c.Instuctor)
}

func(c Course) SignUp() bool{
	return true

}
