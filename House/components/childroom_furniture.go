package components

import "fmt"

type CharacteristicChildroomFurniture struct {
	Type     string
	Material string
	Height   int
	Length   int
}

func (ccf CharacteristicChildroomFurniture) Print() {
	fmt.Print("\n- ", ccf.Type, "\nматериал: ", ccf.Material, " || высота: ", ccf.Height, " м || длина: ", ccf.Length, " м")
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

type Desk struct {
	CharacteristicChildroomFurniture CharacteristicChildroomFurniture
}

type Deskchairs struct {
	CharacteristicChildroomFurniture CharacteristicChildroomFurniture
}
