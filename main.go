package main

import (
	"ExpenseTracker/data"
	"flag"
	"fmt"
	"os"
	"time"
)

var er data.ExpensesRepository

func main() {
	er.LoadExpenses()
	switch os.Args[1] {
	case "add":
		AddTask()
	case "delete":
		DeleteTask()
	case "summary":
		Summary()
	case "list":
		PresentTasks()
	case "update":
		UpdateTask()
	default:
		fmt.Println("none")
	}
}

func UpdateTask() {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	id := updateCmd.Int("id", -1, "id task")
	description := updateCmd.String("description", "-", "Description for expense")
	amount := updateCmd.Int("amount", 0, "Amount expense")
	category := updateCmd.Int("category", 0, "Categories: \n Undefine Category - 0\n Home - 1\n Food - 2\n	Clothes - 3\n Hobby - 4\n Healths - 5\n")

	updateCmd.Parse(os.Args[2:])

	if *id == -1 {
		fmt.Println("Введите пожалуйста id")
		return
	}

	updateExpense := data.Expense{
		Id: *id,
		Date: data.DateStruct{
			Day:   time.Now().Day(),
			Month: int(time.Now().Month()),
			Year:  time.Now().Year(),
		},
		Category:    data.Category(*category),
		Description: *description,
		Amount:      *amount,
	}

	er.UpdateById(updateExpense)

	fmt.Printf("Expense updated successfully (ID: %d)\n", *id)

}

func Summary() {

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	day := summaryCmd.Int("day", 0, "Day for filter")
	month := summaryCmd.Int("month", 0, "Month for filter")
	year := summaryCmd.Int("year", 0, "Year for filter")

	summaryCmd.Parse(os.Args[2:])

	er.Summary(data.DateStruct{
		Day:   *day,
		Month: *month,
		Year:  *year,
	})
}

func PresentTasks() {
	presentCmd := flag.NewFlagSet("list", flag.ExitOnError)

	day := presentCmd.Int("day", 0, "Day for filter")
	month := presentCmd.Int("month", 0, "Month for filter")
	year := presentCmd.Int("year", 0, "Year for filter")

	presentCmd.Parse(os.Args[2:])

	er.PresentTasks(data.DateStruct{
		Day:   *day,
		Month: *month,
		Year:  *year,
	})
}

func AddTask() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	description := addCmd.String("description", "-", "Description for expense")
	amount := addCmd.Int("amount", 0, "Amount expense")
	category := addCmd.Int("category", 0, "Categories: \n Undefine Category - 0\n Home - 1\n Food - 2\n Clothes - 3\n Hobby - 4\n Healths - 5")

	addCmd.Parse(os.Args[2:])

	lastId := er.GetLastId()

	newExpense := data.Expense{
		Id: lastId + 1,
		Date: data.DateStruct{
			Day:   time.Now().Day(),
			Month: int(time.Now().Month()),
			Year:  time.Now().Year(),
		},
		Category:    data.Category(*category),
		Description: *description,
		Amount:      *amount,
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.Id)

	er.AddNewExpense(newExpense)
}

func DeleteTask() {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	taskId := deleteCmd.Int("id", 0, "task id")

	deleteCmd.Parse(os.Args[2:])

	er.DeleteById(*taskId)

	fmt.Println("Expense deleted successfully")
}
