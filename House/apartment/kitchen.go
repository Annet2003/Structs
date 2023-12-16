package apartment

import (
	"Structs/House/components"
)

type Kitchen struct {
	Table        components.Table
	Chairs       components.Chairs
	Refrigerator components.Refrigerator
}

func CreateKitchen() Kitchen {
	table := components.Table{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		Type:        "Стол",
		Amount:      4,
		Material:    "Дерево + стекло",
		Appointment: "один обеденный, два письменных, один жернальный",
		//NumberOfSeats:      4,
		MoistureResistance: true,
	}}

	chairs := components.Chairs{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		Type:               "Стулья",
		Amount:             15,
		Material:           "Дерево",
		Appointment:        "десять обеденных и пять для пиьменного стола ",
		MoistureResistance: true,
		//TypeConstruction: "обеденные - трансформеры, письменные - стул-кресло",
	}}

	refrigerator := components.Refrigerator{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		Type:               "Холодильник",
		Amount:             1,
		Material:           "листовая сталь",
		Appointment:        "холодильная + морозильная камера",
		MoistureResistance: true,
		//Height:            2,
		//Width:             1,
		//Color:             "мраморный",
		//Volume:            300,
		//MaxMinTemperature: "-18 +5",

	}}

	table.CharacteristicKitchenFurniture.Print()
	kitchen := Kitchen{Table: table, Chairs: chairs, Refrigerator: refrigerator}
	return kitchen
}
