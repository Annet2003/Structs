package components

import "fmt"

type RSquare struct {
	Name   string
	Length float32
	Width  float32
	Height float32
	Square float32
}

func (rs RSquare) Print() {
	fmt.Print("\n - ", rs.Name, " - длина: ", rs.Length, " м", " // ширина: ", rs.Width, " м", " // высота: ", rs.Height, " м", " // площаль: ", rs.Square, " м^2")
}

type SKitchen struct {
	RSquare RSquare
}

type SBathroom struct {
	RSquare RSquare
}

type SBedroom struct {
	RSquare RSquare
}

type SChildroom struct {
	RSquare RSquare
}
