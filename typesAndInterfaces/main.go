package main

import "fmt"

type distanceKm float64
type distanceM float64

type User struct{
	id int
	name string
}

func (d distanceM) toKm()distanceKm{

	return distanceKm(d*1.60935)
}
func (d distanceKm) toMiles() distanceM{
	return distanceM(d/1.60935)
}

func main(){
	distanceK := distanceKm(100)
	fmt.Println("Distance in Km: ", distanceK)
	distanceM := distanceK.toMiles()
	fmt.Println("Distance in Miles: ", distanceM)
	
	var u1 User
	u1 = User{id: 1, name:"A"}
	u2:= User{id: 2, name:"B"}
	fmt.Println(u1.id, u2.name)
}
