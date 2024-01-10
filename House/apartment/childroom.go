package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Childroom struct {
	Sofa       components.Sofa
	Playhouse  components.Playhouse
	Carpet     components.Carpet
	Desk       components.Desk
	Deskchairs components.Deskchairs
}

func CreateChildroom() Childroom {

	sofa := components.Sofa{CharacteristicChildroomFurniture: components.CharacteristicChildroomFurniture{
		Type:     "Кровать",
		Material: "дерево",
		Height:   1,
		Length:   2,
	}}
	playhouse := components.Playhouse{CharacteristicChildroomFurniture: components.CharacteristicChildroomFurniture{
		Type:     "Игровой домик",
		Material: "ДСП, ткань",
		Height:   1,
		Length:   1,
	}}

	carpet := components.Carpet{CharacteristicChildroomFurniture: components.CharacteristicChildroomFurniture{
		Type:     "Ковер",
		Material: "шерсть",
		Height:   2,
		Length:   2,
	}}

	desk := components.Desk{CharacteristicChildroomFurniture: components.CharacteristicChildroomFurniture{
		Type:     "Письменный стол",
		Material: "дерево",
		Height:   1,
		Length:   2,
	}}

	deskchairs := components.Deskchairs{CharacteristicChildroomFurniture: components.CharacteristicChildroomFurniture{
		Type:     "Стулья",
		Material: "дерево",
		Height:   1,
		Length:   1,
	}}

	fmt.Print("\n\nДЕТСКАЯ КОМНАТА")

	sofa.CharacteristicChildroomFurniture.Print()
	playhouse.CharacteristicChildroomFurniture.Print()
	carpet.CharacteristicChildroomFurniture.Print()
	desk.CharacteristicChildroomFurniture.Print()
	deskchairs.CharacteristicChildroomFurniture.Print()

	childroom := Childroom{Sofa: sofa, Playhouse: playhouse, Carpet: carpet, Desk: desk, Deskchairs: deskchairs}

	return childroom

}
