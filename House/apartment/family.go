package apartment

import (
	"Structs/House/components"
	"fmt"
)

type Family struct {
	Father   components.Father
	Mother   components.Mother
	Daughter components.Daughter
	Son      components.Son
}

func CreateFamily() Family {

	father := components.Father{CharacteristicFamilyData: components.CharacteristicFamilyData{
		Name:     "Сергей",
		Age:      38,
		Height:   1.88,
		Weight:   76,
		Position: "юрист",
	}}

	mother := components.Mother{CharacteristicFamilyData: components.CharacteristicFamilyData{
		Name:     "Ольга",
		Age:      35,
		Height:   1.75,
		Weight:   60,
		Position: "учитель",
	}}

	daughter := components.Daughter{CharacteristicFamilyData: components.CharacteristicFamilyData{
		Name:     "Настя",
		Age:      12,
		Height:   1.55,
		Weight:   30,
		Position: "ученица",
	}}

	son := components.Son{CharacteristicFamilyData: components.CharacteristicFamilyData{
		Name:     "Антон",
		Age:      14,
		Height:   1.60,
		Weight:   50,
		Position: "ученик",
	}}

	fmt.Print("СЕМЬЯ")

	father.CharacteristicFamilyData.Print()
	mother.CharacteristicFamilyData.Print()
	daughter.CharacteristicFamilyData.Print()
	son.CharacteristicFamilyData.Print()

	family := Family{Father: father, Mother: mother, Daughter: daughter, Son: son}

	return family
}
