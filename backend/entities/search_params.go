package entities

type SearchParamter struct {
	What         string
	Kind         string
	Where        []string
	PriceStart   float64
	PriceEnd     float64
	Squaremeters int
	RoomNumber   float64
}
