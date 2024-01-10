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

	//Mattress string // пружинный или беспружинный
}

type Closet struct {
	CharacteristicBedroomFurniture CharacteristicBedroomFurniture
	//Type          string
	//DoorOpener    string
	//AmountOfDoors int
}

type LadiesTable struct {
	CharacteristicBedroomFurniture CharacteristicBedroomFurniture
}
