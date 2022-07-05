package main

import (
	"fmt"
	"strconv"
)

type Dog struct {
	weight int
	name   string
}

type Cat struct {
	weight int
	name   string
}

type Cow struct {
	weight int
	name   string
}

func (d Dog) GetFoodConsumption() int {
	return d.weight * 2
}

func (d Cat) GetFoodConsumption() int {
	return d.weight * 7
}

func (d Cow) GetFoodConsumption() int {
	return d.weight * 25
}

type Animal interface {
	GetFoodConsumption() int
}

func (d Dog) String() string {
	return "my name: " + d.name + " my weight: " + strconv.Itoa(d.weight) + " food for me: " +
		strconv.Itoa(d.GetFoodConsumption())
}

func (c Cat) String() string {
	return "my name: " + c.name + " my weight: " + strconv.Itoa(c.weight) + " food for me: " +
		strconv.Itoa(c.GetFoodConsumption())
}

func (c Cow) String() string {
	return "my name: " + c.name + " my weight: " + strconv.Itoa(c.weight) + " food for me: " +
		strconv.Itoa(c.GetFoodConsumption())
}

func CalculateFoodConsumption(animals []Animal) int {
	var foodConsumption int
	for _, animal := range animals {
		fmt.Println(animal)
		foodConsumption += animal.GetFoodConsumption()
	}
	return foodConsumption
}

func main() {
	var animals []Animal
	animals = append(animals, Dog{weight: 10, name: "dog"})
	animals = append(animals, Cat{weight: 5, name: "cat"})
	animals = append(animals, Cow{weight: 100, name: "cow"})
	fmt.Println(CalculateFoodConsumption(animals))
}
