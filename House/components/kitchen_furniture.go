package components

import "fmt"

// мебель кухни

type CharacteristicKitchenFurniture struct {
	Type        string
	Amount      int
	Material    string
	Appointment string // назначение
	//NumberOfSeats      int
	MoistureResistance bool // влагоустойчивость
}

func (ckf CharacteristicKitchenFurniture) Print() {
	fmt.Print("\nтип: ", ckf.Type, "\nколичесвто: ", ckf.Amount, "\nматериал: ", ckf.Material, "\nназначение: ", ckf.Appointment, "\nвлагоустойчивость: ", ckf.MoistureResistance)
}

type Table struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
}

type Chairs struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
	//TypeConstruction: "обеденные - трансформеры, письменные - стул-кресло",
}

type Refrigerator struct {
	CharacteristicKitchenFurniture CharacteristicKitchenFurniture
}
