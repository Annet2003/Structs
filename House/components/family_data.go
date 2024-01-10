package components

import "fmt"

type CharacteristicFamilyData struct {
	Name     string
	Age      int
	Height   float32
	Weight   int
	Position string
}

func (cfd CharacteristicFamilyData) Print() {
	fmt.Print("\nимя: ", cfd.Name, "\nвозраст: ", cfd.Age, "\nрост: ", cfd.Height, "\nвес: ", cfd.Weight, "\nдолжность: ", cfd.Position)
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
