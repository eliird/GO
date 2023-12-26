package data

type Instructor struct{
	Id int //the small letter one cannot be used from outside this package
	FirstName string
	LastName string
	Score int
}

func NewInstructor(name string, lastname string) Instructor{
	return Instructor{FirstName: name, LastName: lastname}
}
