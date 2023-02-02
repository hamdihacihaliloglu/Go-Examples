package structs

import "fmt"

type student struct {
	firstName string
	lastName  string
	schoolId  int
	age       int
}

func (s student) save() {
	fmt.Println("Save : ", s.firstName)
}

func Workshop2() {
	s := student{firstName: "Hamdi", lastName: "Hho", schoolId: 117, age: 23}
	s.save()
}
