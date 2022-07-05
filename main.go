package main

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
    return d.name
}

func (c Cat) String() string {
    return c.name
}

func (c Cow) String() string {
    return c.name
}

func CalculateFoodConsumption(animals []Animal) int {
    var foodConsumption int
    for _, animal := range animals {
        foodConsumption += animal.GetFoodConsumption()
    }
    return foodConsumption
}

func main() {

}
