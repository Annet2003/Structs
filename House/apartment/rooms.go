package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Rooms struct {
	SBathroom  components.SBathroom
	SBedroom   components.SBedroom
	SChildroom components.SChildroom
	SKitchen   components.SKitchen
}

func CreateSquareRooms() Rooms {

	skitchen := components.SKitchen{RSquare: components.RSquare{
		Name:   "Кухня",
		Length: 6.5,
		Width:  5.0,
		Height: 3.0,
		Square: 32.5,
	}}

	sbathroom := components.SBathroom{RSquare: components.RSquare{
		Name:   "Ванная комната",
		Length: 4.0,
		Width:  3.0,
		Height: 3.0,
		Square: 12,
	}}
	sbedroom := components.SBedroom{RSquare: components.RSquare{
		Name:   "Спальня",
		Length: 7.4,
		Width:  5.0,
		Height: 3.0,
		Square: 37,
	}}
	schildroom := components.SChildroom{RSquare: components.RSquare{
		Name:   "Детская комната",
		Length: 6.0,
		Width:  3.0,
		Height: 3.0,
		Square: 18,
	}}

	fmt.Print("\n\nПостроился дом для семьи!\n")
	fmt.Print("\nРазмеры и площади комнат: ")

	skitchen.RSquare.Print()
	sbathroom.RSquare.Print()
	sbedroom.RSquare.Print()
	schildroom.RSquare.Print()

	rooms := Rooms{SKitchen: skitchen, SBathroom: sbathroom, SBedroom: sbedroom, SChildroom: schildroom}
	fmt.Print("\n\n Заполняем дом мебелью ...")
	return rooms

}
