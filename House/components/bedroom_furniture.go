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
	fmt.Print("\nтип: ", cbf.Type, "\nколичесвто: ", cbf.Amount, "\nматериал: ", cbf.Material, "\nвысота: ", cbf.Height, "\nдлина: ", cbf.Length, "\nгарантия: ", cbf.Guarantee)
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
