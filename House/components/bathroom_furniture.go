package components

import "fmt"

type CharacteristicBathroomFurniture struct {
	Type       string
	Color      string
	Volume     int    // объем
	Dimensions string // габариты
	Material   string
	Guarantee  bool // гарантия
}

func (cbaf CharacteristicBathroomFurniture) Print() {
	fmt.Print("\nтип", cbaf.Type, "\nцвет: ", cbaf.Color, "\nобъем: ", cbaf.Volume, "\nгабариты: ", cbaf.Dimensions, "\nматериал: ", cbaf.Material, "\nгарантия: ", cbaf.Guarantee)
}

type WashingMachine struct {
	CharacteristicBathroomFurniture CharacteristicBathroomFurniture
	//Method                  string // способ загрузки
	//AmountOfWashingPrograms int    // количство программа стирки
	//Dimensions              string // габариты
	//SpinMode                bool   // режим отжима
}
