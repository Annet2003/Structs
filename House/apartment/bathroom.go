package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Bathroom struct {
	WashingMachine components.WashingMachine
}

func CreateBathroom() Bathroom {

	washingmachine := components.WashingMachine{CharacteristicBathroomFurniture: components.CharacteristicBathroomFurniture{
		Type:       "Стиральная машина",
		Color:      "белый",
		Volume:     6,                    // объем
		Dimensions: "59.6 * 84.6 * 42.1", // габариты
		Material:   "нержавеющая сталь",
		Guarantee:  true, // гарантия
		//Method:                  "фронтальная",        // способ загрузки
		//AmountOfWashingPrograms: 11,                   // количство программа стирки
		//SpinMode:                true,                 // режим отжима
	}}

	fmt.Print("ВАННАЯ КОМНАТА")
	washingmachine.CharacteristicBathroomFurniture.Print()

	bathroom := Bathroom{WashingMachine: washingmachine}

	return bathroom

}
