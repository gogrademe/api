package model

type SchoolYear struct {
	ID    uint
	Start int
	End   int
	Terms []Term
}
