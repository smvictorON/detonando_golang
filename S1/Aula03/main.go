package main

import (
	"fmt"
	"math"
	"reflect"
)

type Student struct {
	name string
}

type Course struct {
	description string
	students    []Student
}

//usando ponteiro aqui no Course o que mudarmos aqui dentro vale
//para essa struct fora tambem
func (c *Course) register(s Student) {
	c.students = append(c.students, s)
}

type EadCourse struct {
	course  Course
	website string
}

type NewFloat float64

func (nf NewFloat) roundBy(places float64) NewFloat {
	nplaces := math.Pow(10.00, places)
	return NewFloat(math.Round(float64(nf)*nplaces) / nplaces)
}

func main() {
	var newStudent Student
	newStudent.name = "Victor"
	//reflect pega o tipo do dado
	fmt.Println(newStudent, reflect.TypeOf(newStudent))

	//passando por ponteiro
	pst := &newStudent
	//as duas formas abaixo printam a mesma coisa
	fmt.Println(pst.name, reflect.TypeOf(newStudent))
	fmt.Println((*pst).name, reflect.TypeOf(newStudent))

	var other Student
	other.name = "Other Student"
	fmt.Println(pst.name, other.name)

	p := Student{"Lamara"}
	fmt.Println(p.name)

	engineering := Course{"Engineering", make([]Student, 0)}
	engineering.register(newStudent)
	engineering.register(p)
	fmt.Println(engineering)

	var n NewFloat = 5.0293019384
	fmt.Println(n.roundBy(2))
	fmt.Println(n.roundBy(3))

	ead := EadCourse{Course{"New Ead", make([]Student, 0)}, "ead.com"}
	ead.course.register(newStudent)
	fmt.Println(ead)
}
