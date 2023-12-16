package main

import "Structs/House"

func main() {
	House.CreateHouse()
}

/*func main() {

	// Family - семья

	mother := House.Mother{
		NameMother:     "Ольга",
		AgeMother:      35,
		HeightMother:   1.75,
		WEightMother:   60,
		PositionMother: "учитель",
	}
	fmt.Print(mother)

	father := House.Father{
		NameFather:     "Сергей",
		AgeFather:      38,
		HeightFather:   1.88,
		WEightFather:   76,
		PositionFather: "юрист",
	}
	fmt.Print(father)

	daughter := House.Daughter{
		NameDaughter:     "Настя",
		AgeDaughter:      12,
		HeightDaughter:   1.55,
		WeightDaughter:   30,
		PositionDaughter: "ученица",
	}
	fmt.Print(daughter)

	son := House.Son{
		NameSom:     "Антон",
		AgeSon:      14,
		HeightSon:   1.60,
		WeightSon:   50,
		PositionSon: "ученик",
	}
	fmt.Print(son)

	// Appliances - техника

	tv := House.TV{
		Amount:       4,
		Diagonal:     120,
		MaximumPower: 400, //максимальная потребляемая мощность
		Model:        "Samsung",
		ThreeDSound:  true,
		SmartTV:      true,
	}
	fmt.Print(tv)

	refrigerator := House.Refrigerator{
		Height:            2,
		Width:             1,
		Color:             "мраморный",
		Volume:            300,
		MaxMinTemperature: "-18 +5",
		Freezer:           true, // наличие морозилки
	}
	fmt.Print(refrigerator)

	washingmachine := House.WashingMachine{
		Method:                  "фронтальная",        // способ загрузки
		Volume:                  6,                    // объем
		AmountOfWashingPrograms: 11,                   // количство программа стирки
		Dimensions:              "59.6 * 84.6 * 42.1", // габариты
		SpinMode:                true,                 // режим отжима
	}

	fmt.Print(washingmachine)

	cooker := House.Cooker{
		Manufacturer:          "Candy",
		Type:                  "Электрическая",
		Size:                  "90 * 60",
		AdvancedFunctionality: true, // расширенный функционал
		Volume:                50,
	}
	fmt.Print(cooker)

	other := House.Other{
		ElectricKettle: true,
		Toaster:        true,
		Microwave:      true,
	}
	fmt.Print(other)

	// Furniture - мебель

	bed := House.Bed{
		Amount:    5,
		Material:  "Дерево",
		Mattress:  "беспружинный",
		Length:    2,
		Guarantee: true, // гарантия
	}
	fmt.Print(bed)

	table := House.Table{
		Amount:             4,
		Material:           "Дерево + стекло",
		Appointment:        "один обеденный, два письменных, один жернальный",
		NumberOfSeats:      4,
		MoistureResistance: true,
	}
	fmt.Print(table)

	closet := House.Closet{
		BodyMaterial:  "ДСП",
		Type:          "строенный",
		DoorOpener:    "раздвижные",
		AmountOfDoors: 2,
	}
	fmt.Print(closet)

	chairs := House.Chairs{
		Amount:      15,
		Appointment: "десять обеденных и пять для пиьменного стола ",
		Type:        "обеденные - трансформеры, письменные - стул-кресло",
	}
	fmt.Print(chairs)

	// Cultery - столовые приборы

	spoons := House.Spoons{
		Amount:          60,
		Material:        " ",
		NeedForVarietes: true,
		Variety:         "десертная, суповая, столовая, салатная, чайная, сервировочная",
	}
	fmt.Print(spoons)

	forks := House.Forks{
		Users:    5,
		Amount:   15,
		Material: "нержавеющая сталь",
	}
	fmt.Print(forks)

}*/
