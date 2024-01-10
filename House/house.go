package House

import (
	"Structs/House/apartment"
)

type Apartment struct {
	Kitchen   apartment.Kitchen
	Bedroom   apartment.Bedroom
	Bathroom  apartment.Bathroom
	Childroom apartment.Childroom
	Family    apartment.Family
	Rooms     apartment.Rooms
}

type House struct {
	Apartment Apartment
}

func CreateHouse() House {

	apart := Apartment{
		Family:    apartment.CreateFamily(),
		Rooms:     apartment.CreateSquareRooms(),
		Kitchen:   apartment.CreateKitchen(),
		Bedroom:   apartment.CreateBedroom(),
		Bathroom:  apartment.CreateBathroom(),
		Childroom: apartment.CreateChildroom()}

	house := House{Apartment: apart}
	return house
}
