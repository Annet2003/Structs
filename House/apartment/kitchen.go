package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Kitchen struct {
	Table        components.Table
	Chairs       components.Chairs
	Refrigerator components.Refrigerator
	Cooker       components.Cooker
	Garniture    components.Garniture
}

func CreateKitchen() Kitchen {
	table := components.Table{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		Type:               "Столы",
		Amount:             2,
		Material:           "дерево + стекло",
		Appointment:        "один обеденный, один журнальный",
		MoistureResistance: true,
	}}

	chairs := components.Chairs{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		Type:               "Стулья",
		Amount:             15,
		Material:           "дерево",
		Appointment:        "10 обеденных и 5 для пиьменного стола ",
		MoistureResistance: true,
		//TypeConstruction:   "обеденные - трансформеры",
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

	cooker := components.Cooker{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		// плита, печь
		Type:               "Плита",
		Amount:             1,
		Material:           "стеклокерамика",
		Appointment:        "варочная поверхность + духовая печь",
		MoistureResistance: true,
	}}

	garniture := components.Garniture{CharacteristicKitchenFurniture: components.CharacteristicKitchenFurniture{
		Type:               "Гарнитура",
		Amount:             1,
		Material:           "ДСП",
		Appointment:        "столешница с рабочей зоной, шкафчики, сушилка для посуды",
		MoistureResistance: true,
	}}

	fmt.Print("\n\nКУХНЯ")

	table.CharacteristicKitchenFurniture.Print()
	chairs.CharacteristicKitchenFurniture.Print()
	refrigerator.CharacteristicKitchenFurniture.Print()
	cooker.CharacteristicKitchenFurniture.Print()
	garniture.CharacteristicKitchenFurniture.Print()

	kitchen := Kitchen{Table: table, Chairs: chairs, Refrigerator: refrigerator, Cooker: cooker, Garniture: garniture}
	return kitchen
}
