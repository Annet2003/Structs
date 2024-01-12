package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Bedroom struct {
	Bed         components.Bed
	Closet      components.Closet
	LadiesTable components.LadiesTable
}

func CreateBedroom() Bedroom {

	bed := components.Bed{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "Кровать",
		Amount:    5,
		Material:  "дерево",
		Height:    1,
		Length:    2,
		Guarantee: true, // гарантия
	}}
	closet := components.Closet{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "Шкаф",
		Amount:    3,
		Material:  "ДСП",
		Height:    2,
		Length:    1,
		Guarantee: true, // гарантия
	}}

	ladiesTable := components.LadiesTable{CharacteristicBedroomFurniture: components.CharacteristicBedroomFurniture{
		Type:      "Дамский столик",
		Amount:    1,
		Material:  "ДСП + зеркало",
		Height:    2,
		Length:    2,
		Guarantee: true, // гарантия
	}}

	fmt.Print("\n\nСПАЛЬНЯ")
	bed.CharacteristicBedroomFurniture.Print()
	closet.CharacteristicBedroomFurniture.Print()
	ladiesTable.CharacteristicBedroomFurniture.Print()

	bedroom := Bedroom{Bed: bed, Closet: closet, LadiesTable: ladiesTable}

	return bedroom

}
