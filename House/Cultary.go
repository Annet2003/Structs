package House

// столовые приборы

type Spoons struct { // ложки
	Amount          int
	Material        string
	NeedForVarietes bool
	Variety         string // десертная, суповая, столовая, салатная, чайная, сервировочная
}

type Forks struct { // вилки
	Users    int
	Amount   int
	Material string
}
