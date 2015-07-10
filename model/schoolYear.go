package model

import "github.com/mholt/binding"

type SchoolYear struct {
	ID    uint   `json:"id"`
	Start int    `json:"start"`
	End   int    `json:"end"`
	Terms []Term `json:"terms"`
}

func (t *SchoolYear) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.ID:    "id",
		&t.Start: "start",
		&t.End:   "end",
		&t.Terms: "terms",
	}
}
