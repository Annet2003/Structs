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
}

type House struct {
	Apartment Apartment
}

func CreateHouse() House {
	rooms := Apartment{
		Family:    apartment.CreateFamily(),
		Kitchen:   apartment.CreateKitchen(),
		Bedroom:   apartment.CreateBedroom(),
		Bathroom:  apartment.CreateBathroom(),
		Childroom: apartment.CreateChildroom()}

	house := House{Apartment: rooms}
	return house
}
