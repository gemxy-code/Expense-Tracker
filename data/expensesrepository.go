package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gosuri/uitable"
)

type ExpensesRepository struct {
	AllExpenses []Expense
}

func (er ExpensesRepository) FilterByDate(date DateStruct) ([]Expense, error) {
	result := make([]Expense, 0)

	for i, val := range er.AllExpenses {
		if (date.Day == val.Date.Day || date.Day == 0) &&
			(date.Month == val.Date.Month || date.Month == 0) &&
			(date.Year == val.Date.Year || date.Year == 0) {
			result = append(result, er.AllExpenses[i])
		}
	}

	if len(result) > 0 {
		return result, nil
	}

	return nil, errors.New("нет данных, соответствующих выбранной дате")
}

func (er ExpensesRepository) Summary(date DateStruct) {

	filteredTasks, err := er.FilterByDate(date)

	if err != nil {
		fmt.Printf("Summary: error in filter: %v\n", err.Error())
		return
	}

	sum := 0

	for _, val := range filteredTasks {
		sum += val.Amount
	}

	fmt.Printf("# Total expenses: %d\n", sum)
}

func (er ExpensesRepository) PresentTasks(date DateStruct) {

	filteredTasks, err := er.FilterByDate(date)

	if err != nil {
		fmt.Printf("PresentTasks: error in filter: %v\n", err.Error())
		return
	}

	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("ID", "Date ", "Description", "Amount")
	for _, val := range filteredTasks {
		table.AddRow(val.Id, val.Date, val.Description, val.Amount)
	}
	fmt.Println(table)
}

func (er *ExpensesRepository) DeleteById(taskId int) error {

	for index, val := range er.AllExpenses {
		if val.Id == taskId {
			if index == len(er.AllExpenses)-1 {
				er.AllExpenses = er.AllExpenses[:index]
			} else if index == 0 {
				er.AllExpenses = er.AllExpenses[index+1:]
			} else {
				er.AllExpenses = append(er.AllExpenses[:index], er.AllExpenses[index+1:]...)
			}
			er.saveExpenses()
			return nil
		}
	}

	return errors.New("отсутствует задача с таким ID :(")

}

func (er *ExpensesRepository) AddNewExpense(newExpense Expense) {
	er.AllExpenses = append(er.AllExpenses, newExpense)
	er.saveExpenses()
}

func (er *ExpensesRepository) LoadExpenses() {
	buf, err := os.ReadFile("temp/expenses.json")
	if err != nil && err != io.EOF {
		fmt.Printf("LoadExpenses read file error: %s\n", err.Error())
		return
	}

	if len(buf) == 0 {
		fmt.Println("Нет данных в файле")
		return
	}

	err = json.Unmarshal(buf, &er.AllExpenses)
	if err != nil {
		fmt.Printf("LoadExpenses unmarshal json error: %s\n", err.Error())
		return
	}
}

func (er ExpensesRepository) GetLastId() int {
	lastId := -1
	if len(er.AllExpenses) > 0 {
		lastId = er.AllExpenses[len(er.AllExpenses)-1].Id
	}
	return lastId
}

func (er ExpensesRepository) saveExpenses() {
	b, err := json.Marshal(er.AllExpenses)
	if err != nil {
		fmt.Printf("SaveExpenses marshal json error: %s\n", err.Error())
		return
	}

	file, err := os.OpenFile("temp/expenses.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Printf("SaveExpenses open file error: %s\n", err.Error())
		return
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		fmt.Printf("SaveExpenses write in file error: %s\n", err.Error())
		return
	}
}
