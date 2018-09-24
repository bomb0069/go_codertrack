package expense_day

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var dayMapper = map[string]string{
	"5": "Mon",
	"6": "Tue",
	"7": "Wed",
	"1": "Thu",
	"2": "Fri",
	"3": "Sat",
	"4": "Sun",
}

type Expense struct {
	byDay map[string]int
}

func NewExpense() Expense {
	byDay := map[string]int{
		"Mon": 0,
		"Tue": 0,
		"Wed": 0,
		"Thu": 0,
		"Fri": 0,
		"Sat": 0,
		"Sun": 0,
	}
	return Expense{byDay}
}

func (e *Expense) read(fileName string) {
	data, _ := ioutil.ReadFile(fileName)
	splittedLine := strings.Split(string(data), "\r\n")
	for _, line := range splittedLine {
		expenseDate := strings.Fields(line)
		day := ""
		for index, word := range expenseDate {
			expenseCostString := ""
			if index == 0 {
				day = dayMapper[word]
			} else {
				expenseCostString = word[1:len(word)]
				expenseCost, _ := strconv.Atoi(expenseCostString)
				e.byDay[day] = e.byDay[day] + expenseCost
			}
		}

	}
}

func (e *Expense) summaryByDay() map[string]int {
	return e.byDay
}
