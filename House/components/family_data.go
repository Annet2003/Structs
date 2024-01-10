package components

import "fmt"

type CharacteristicFamilyData struct {
	Type     string
	Name     string
	Age      int
	Height   float32
	Weight   int
	Position string
}

func (cfd CharacteristicFamilyData) Print() {
	fmt.Print("\n- ", cfd.Type, " - ", cfd.Name, " -", "\nвозраст: ", cfd.Age, " лет; рост: ", cfd.Height, " м; вес: ", cfd.Weight, " кг; должность: ", cfd.Position)
}

type Mother struct {
	CharacteristicFamilyData CharacteristicFamilyData
}

type Father struct {
	CharacteristicFamilyData CharacteristicFamilyData
}

type Daughter struct {
	CharacteristicFamilyData CharacteristicFamilyData
}

type Son struct {
	CharacteristicFamilyData CharacteristicFamilyData
}
