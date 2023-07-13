package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	TurnOn() bool
	TurnOff() bool
	Move(direction int, speed int) bool
	Stop()
}

type Gas interface {
	FillUp()
}

//mesmo se estar explicitamente dizendo, ao implementar todos os métodos de Vehicle
//seria com se Car extendesse de vehicle, é necessário implementar todos métodos.
type Car struct {
	model string
}

func (c Car) TurnOn() bool {
	fmt.Println("Turning the car ON!")
	return true
}

func (c Car) TurnOff() bool {
	fmt.Println("Turning the car OFF!")
	return true

}

func (c Car) Move(direction int, speed int) bool {
	fmt.Println("Moving the car")
	return true
}

func (c Car) Stop() {
	fmt.Println("Stoping the car")
}

func (c Car) FillUp() {
	fmt.Println("Filling the car's tank")
}

type Truck struct {
	model string
}

func (t Truck) TurnOn() bool {
	fmt.Println("Turning the truck ON!")
	return true
}

func (t Truck) TurnOff() bool {
	fmt.Println("Turning the truck OFF!")
	return true

}

func (t Truck) Move(direction int, speed int) bool {
	fmt.Println("Moving the truck")
	return true
}

func (t Truck) Stop() {
	fmt.Println("Stoping the truck")
}

type Computer interface {
	TurnOn() bool
	TurnOff() bool
}

type Laptop struct {
	model string
	on    bool
}

func (l Laptop) TurnOn() bool {
	fmt.Println("Turning on Laptop")
	l.on = true
	return true
}

func (l Laptop) TurnOff() bool {
	fmt.Println("Turning off Laptop")
	l.on = false
	return true
}

func main() {
	newCar := Car{"Mustang"}
	fmt.Println(newCar, reflect.TypeOf(newCar))
	newCar.TurnOn()
	newCar.Move(1, 1)
	newCar.Stop()
	newCar.TurnOff()

	fmt.Println("----")

	var newVehicle Vehicle
	fmt.Println(newVehicle, reflect.TypeOf(newVehicle))
	newVehicle = Car{"Camaro"}
	fmt.Println(newVehicle, reflect.TypeOf(newVehicle))
	newVehicle.TurnOn()
	newVehicle.Move(1, 1)
	newVehicle.Stop()
	newVehicle.TurnOff()

	fmt.Println("----")

	newTruck := Truck{"Scania"}
	fmt.Println(newTruck, reflect.TypeOf(newTruck))
	newTruck.TurnOn()
	newTruck.Move(1, 1)
	newTruck.Stop()
	newTruck.TurnOff()

	fmt.Println("----")

	var aVehicle interface{}
	fmt.Println(aVehicle, reflect.TypeOf(aVehicle))
	aVehicle = newTruck
	fmt.Println(aVehicle, reflect.TypeOf(aVehicle))

	fmt.Println("----")

	//nessa linha quero saber se newVehicle implementa Truck, porém newVehicle é um Car
	value, answer := newVehicle.(Truck)
	fmt.Println(value, answer)

	fmt.Println("----")

	//newVehicle = newCar

	//aqui eu checo se newVehicle implementar Car(sim) e tambem Gas que é implementado por Car
	v2, a2 := newVehicle.(Car)
	fmt.Println(v2, a2)
	v3, a3 := newVehicle.(Gas)
	fmt.Println(v3, a3)

	fmt.Println("----")

	// var c Computer = Laptop{"Asus", false}
	var c Computer = &Laptop{"Asus", false}
	fmt.Println(c)
	c.TurnOn()
}
