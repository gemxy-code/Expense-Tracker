package data

import "fmt"

type Category int

const (
	Undefine Category = iota
	Home
	Food
	Clothes
	Hobby
	Health
)



func (c Category) ToString() string {
	switch c {
	case Home:
		return "Дом"
	case Food:
		return "Продукты"
	case Clothes:
		return "Вещи"
	case Hobby:
		return "Хобби"
	case Health:
		return "Здоровье"
	default:
		return "Неопределённая"
	}
}

type Expense struct {
	Id          int        `json:"id"`
	Date        DateStruct `json:"date"`
	Category    Category   `json:"category"`
	Description string     `json:"description"`
	Amount      int        `json:"amount"`
}

func (ex Expense) String() string {
	return fmt.Sprintf("#%d %s %s %d₽", ex.Id, ex.Date, ex.Description, ex.Amount)
}
