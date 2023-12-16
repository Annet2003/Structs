package House

import (
	"Structs/House/apartment"
)

type Apartment struct {
	Kitchen  apartment.Kitchen
	Bedroom  apartment.Bedroom
	Bathroom apartment.Bathroom
}

type House struct {
	Apartment Apartment
}

func CreateHouse() House {
	rooms := Apartment{Kitchen: apartment.CreateKitchen(),
		Bedroom:  apartment.CreateBedroom(),
		Bathroom: apartment.CreateBathroom()}

	house := House{Apartment: rooms}
	return house
}
