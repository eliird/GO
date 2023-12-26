package data

import "time"

type Workshop struct{
	Course Course
	Date time.Time

}

type EmbeddedWorkshop struct{
	Course
	Name string //if you ask for name you will get this name rather than the name from course
	Date time.Time
	Instructor Instructor
}

func NewWorkshop(name string, instructor Instructor) EmbeddedWorkshop{

	w:= EmbeddedWorkshop{}
	w.Name = name
	w.Instructor = instructor
	w.Course.Name = "Sometjing got named something"
	return w
}
func (w Workshop)SignUp()bool{
	return true
}
