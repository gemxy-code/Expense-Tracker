package data

import (
	"fmt"
)

type DateStruct struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

func (d DateStruct) String() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}
