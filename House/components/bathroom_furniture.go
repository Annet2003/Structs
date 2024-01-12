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
	fmt.Print("\n- ", cbaf.Type, "\nцвет: ", cbaf.Color, " || материал: ", cbaf.Material, "\nобъем: ", cbaf.Volume, " л", " || габариты: ", cbaf.Dimensions, "\nгарантия: ", cbaf.Guarantee)
}

type WashingMachine struct {
	CharacteristicBathroomFurniture CharacteristicBathroomFurniture
}

type Sink struct {
	CharacteristicBathroomFurniture CharacteristicBathroomFurniture
}

type Bath struct {
	CharacteristicBathroomFurniture CharacteristicBathroomFurniture
}

type Toilet struct {
	CharacteristicBathroomFurniture CharacteristicBathroomFurniture
}

type Mirror struct {
	CharacteristicBathroomFurniture CharacteristicBathroomFurniture
}
