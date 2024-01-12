package components

import "fmt"

type CharacteristicBedroomFurniture struct {
	Type      string
	Amount    int
	Material  string
	Height    int
	Length    int
	Guarantee bool // гарантия
}

func (cbf CharacteristicBedroomFurniture) Print() {
	fmt.Print("\n- ", cbf.Type, "\nколичесвто: ", cbf.Amount, " || материал: ", cbf.Material, "\nвысота: ", cbf.Height, " м", " || длина: ", cbf.Length, " м", "\nгарантия: ", cbf.Guarantee)
}

type Bed struct {
	CharacteristicBedroomFurniture CharacteristicBedroomFurniture
}

type Closet struct {
	CharacteristicBedroomFurniture CharacteristicBedroomFurniture
}

type LadiesTable struct {
	CharacteristicBedroomFurniture CharacteristicBedroomFurniture
}
