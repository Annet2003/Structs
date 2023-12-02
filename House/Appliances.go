// техника

package House

type TV struct {
	Amount       int
	Diagonal     int
	MaximumPower int //максимальная потребляемая мощность вТ
	Model        string
	ThreeDSound  bool
	SmartTV      bool
}

type Refrigerator struct { //холодильник
	Height            int
	Width             int
	Color             string
	Volume            int // объем
	MaxMinTemperature string
	Freezer           bool // наличие морозилки
}

type WashingMachine struct {
	Method                  string // способ загрузки
	Volume                  int    // объем
	AmountOfWashingPrograms int    // количство программа стирки
	Dimensions              string // габариты
	SpinMode                bool   // режим отжима
}

type Cooker struct { // плита, печь
	Manufacturer          string // производитель
	Type                  string // электоическая или газовая
	Size                  string
	AdvancedFunctionality bool // расширенный функционал
	Volume                int
}

type Other struct {
	ElectricKettle bool
	Toaster        bool
	Microwave      bool
}
