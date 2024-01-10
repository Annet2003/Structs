package components

import "fmt"

type CharacteristicChildroomFurniture struct {
	Type     string
	Material string
	Height   int
	Length   int
}

func (ccf CharacteristicChildroomFurniture) Print() {
	fmt.Print("\nтип: ", ccf.Type, "\nматериал: ", ccf.Material, "\nвысота: ", ccf.Height, "\nдлина: ", ccf.Length)
}

type Sofa struct {
	CharacteristicChildroomFurniture CharacteristicChildroomFurniture

	//Mattress string // пружинный или беспружинный
}

type Playhouse struct {
	CharacteristicChildroomFurniture CharacteristicChildroomFurniture
	//Type          string
	//DoorOpener    string
	//AmountOfDoors int
}

type Carpet struct {
	CharacteristicChildroomFurniture CharacteristicChildroomFurniture
}
