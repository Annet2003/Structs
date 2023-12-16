package apartment

import (
	"Structs/House/components"
)

type Bedroom struct {
	Bed    components.Bed
	Closet components.Closet
}

func CreateBedroom() Bedroom {

	bed := components.Bed{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "кровать",
		Amount:    5,
		Material:  "Дерево",
		Height:    1,
		Length:    2,
		Guarantee: true, // гарантия
		//Mattress: "беспружинный"}
	}}
	closet := components.Closet{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "шкаф",
		Amount:    3,
		Material:  "ДСП",
		Height:    2,
		Length:    1,
		Guarantee: true, // гарантия
		//Type:          "строенный",
		//DoorOpener:    "раздвижные",
		//AmountOfDoors: 2,
	}}

	bed.CharacteristicBedroomFurniture.Print()

	bedroom := Bedroom{Bed: bed, Closet: closet}

	return bedroom

}
