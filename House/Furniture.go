// мебель

package House

type Bed struct {
	Amount    int
	Material  string
	Mattress  string // пружинный или беспружинный
	Length    int
	Guarantee bool // гарантия
}

type Table struct {
	Amount             int
	Material           string
	Appointment        string // назначение
	NumberOfSeats      int
	MoistureResistance bool // влагоустойчивость
}

type Closet struct { // шкаф
	BodyMaterial  string // материал корпуса: дерево, пластик, стекло, ДСП
	Type          string // встроенный или отдельностоящий
	DoorOpener    string // механизм открывания дверей: распашные или раздвижные
	AmountOfDoors int    //
}

type Chairs struct {
	Amount      int
	Appointment string // назначение: обеденные, к письменному столу
	Type        string // трансформер, стандартные Х, стул-кресло
}
