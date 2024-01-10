package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Bathroom struct {
	WashingMachine components.WashingMachine
	Sink           components.Sink
	Bath           components.Bath
	Toilet         components.Toilet
	Mirror         components.Mirror
}

func CreateBathroom() Bathroom {

	washingmachine := components.WashingMachine{CharacteristicBathroomFurniture: components.CharacteristicBathroomFurniture{
		Type:       "Стиральная машина",
		Color:      "белый",
		Volume:     30,                   // объем
		Dimensions: "59.6 * 84.6 * 42.1", // габариты
		Material:   "нержавеющая сталь",
		Guarantee:  true, // гарантия
		//Method:                  "фронтальная",        // способ загрузки
		//AmountOfWashingPrograms: 11,                   // количство программа стирки
		//SpinMode:                true,                 // режим отжима
	}}

	sink := components.Sink{CharacteristicBathroomFurniture: components.CharacteristicBathroomFurniture{
		Type:       "Раковина",
		Color:      "мраморный",
		Volume:     10,             // объем
		Dimensions: "50 * 55 * 40", // габариты
		Material:   "керамика",
		Guarantee:  true, // гарантия
	}}

	bath := components.Bath{CharacteristicBathroomFurniture: components.CharacteristicBathroomFurniture{
		Type:       "Ванна",
		Color:      "белый",
		Volume:     165,             // объем
		Dimensions: "170 * 75 * 60", // габариты
		Material:   "акрил",
		Guarantee:  true, // гарантия
	}}

	toilet := components.Toilet{CharacteristicBathroomFurniture: components.CharacteristicBathroomFurniture{
		Type:       "Унитаз",
		Color:      "белый",
		Volume:     40,              // объем
		Dimensions: "100 * 60 * 40", // габариты
		Material:   "акрил",
		Guarantee:  true, // гарантия
	}}

	mirror := components.Mirror{CharacteristicBathroomFurniture: components.CharacteristicBathroomFurniture{
		Type:       "Зеркало",
		Color:      "с белой холодной подстветкой",
		Dimensions: "100 * 120", // габариты
		Material:   "амальгама",
		Guarantee:  true, // гарантия
	}}

	fmt.Print("\n\nВАННАЯ КОМНАТА")

	washingmachine.CharacteristicBathroomFurniture.Print()
	sink.CharacteristicBathroomFurniture.Print()
	bath.CharacteristicBathroomFurniture.Print()
	toilet.CharacteristicBathroomFurniture.Print()
	mirror.CharacteristicBathroomFurniture.Print()

	bathroom := Bathroom{WashingMachine: washingmachine, Sink: sink, Bath: bath, Toilet: toilet, Mirror: mirror}

	return bathroom

}
