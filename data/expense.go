package data

import "fmt"

type Expense struct {
	Id          int        `json:"id"`
	Date        DateStruct `json:"date"`
	Description string     `json:"description"`
	Amount      int        `json:"amount"`
}

func (ex Expense) String() string {
	return fmt.Sprintf("#%d %s %s %d₽", ex.Id, ex.Date, ex.Description, ex.Amount)
}
