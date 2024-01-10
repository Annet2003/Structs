package components

import "fmt"

// мебель кухни

type CharacteristicKitchenFurniture struct {
	Type               string
	Amount             int
	Material           string
	Appointment        string // назначение
	MoistureResistance bool   // влагоустойчивость
}

func (ckf CharacteristicKitchenFurniture) Print() {
	fmt.Print("\n- ", ckf.Type, "\nколичество: ", ckf.Amount, " || материал: ", ckf.Material, " \nназначение: ", ckf.Appointment, " || влагоустойчивость: ", ckf.MoistureResistance)
}

type Table struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
}

type Chairs struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
	//TypeConstruction               string
	//: "обеденные - трансформеры, письменные - стул-кресло",
}

//func (c Chairs) Print() {
//	fmt.Print("Тип сборки", c.TypeConstruction)
//}

type Refrigerator struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
}

type Cooker struct { // плита, печь
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
}

type Garniture struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
}
