package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Bedroom struct {
	Bed    components.Bed
	Closet components.Closet
}

func CreateBedroom() Bedroom {

	bed := components.Bed{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "Кровать",
		Amount:    5,
		Material:  "дерево",
		Height:    1,
		Length:    2,
		Guarantee: true, // гарантия
		//Mattress: "беспружинный"}
	}}
	closet := components.Closet{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "Шкаф",
		Amount:    3,
		Material:  "ДСП",
		Height:    2,
		Length:    1,
		Guarantee: true, // гарантия
		//Type:          "строенный",
		//DoorOpener:    "раздвижные",
		//AmountOfDoors: 2,
	}}

	fmt.Print("СПАЛЬНЯ")
	bed.CharacteristicBedroomFurniture.Print()
	closet.CharacteristicBedroomFurniture.Print()

	bedroom := Bedroom{Bed: bed, Closet: closet}

	return bedroom

}
