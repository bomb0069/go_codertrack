package expense

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var dayMapper = map[string]string{
	"5": "Mon",
	"6": "Tue",
	"0": "Wed",
	"1": "Thu",
	"2": "Fri",
	"3": "Sat",
	"4": "Sun",
}

var catMapper = map[string]string{
	"f": "food",           // (อาหาร)
	"g": "game",           // (เล่นเกม)
	"m": "movie",          // (ดูหนัง)
	"s": "stationery",     // (เครื่องเขียน)
	"t": "transportation", // (ค่าเดินทาง)
}

type Payment struct {
	date     int
	day      string
	cat      string
	category string
	price    int
}

type Expense struct {
	byDay    map[string]int
	byCat    map[string]int
	payments []Payment
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
	byCat := map[string]int{
		"TOTAL": 0,
	}
	payments := []Payment{}
	return Expense{byDay, byCat, payments}
}

func (e *Expense) read(fileName string) {
	data, _ := ioutil.ReadFile(fileName)
	splittedLine := strings.Split(string(data), "\r\n")
	for _, line := range splittedLine {
		expenseDate := strings.Fields(line)
		day := ""
		dayRead := 0
		dayString := ""
		for index, word := range expenseDate {
			expenseCostString := ""
			if index == 0 {
				dayRead, _ = strconv.Atoi(word)
				dayNum := dayRead % 7
				day = strconv.Itoa(dayNum)
				dayString = dayMapper[day]
				if len(expenseDate) == 1 {
					payment := Payment{
						date:     dayRead,
						day:      dayString,
						cat:      "",
						category: "",
						price:    0,
					}
					e.payments = append(e.payments, payment)
				}
			} else {
				catShort := word[0:1]
				catDescription := catMapper[catShort]
				expenseCostString = word[1:len(word)]
				expenseCost, _ := strconv.Atoi(expenseCostString)
				payment := Payment{
					date:     dayRead,
					day:      dayString,
					cat:      catShort,
					category: catDescription,
					price:    expenseCost,
				}
				e.payments = append(e.payments, payment)
			}
		}

	}
}

func (e *Expense) summaryByDay() map[string]int {
	for _, payment := range e.payments {
		dayString := payment.day
		expenseCost := payment.price
		e.byDay[dayString] = e.byDay[dayString] + expenseCost
	}

	return e.byDay
}

func (e *Expense) summaryByCat() map[string]int {
	total := 0
	for _, payment := range e.payments {
		catString := payment.category
		expenseCost := payment.price
		total = total + expenseCost
		if catString != "" {
			e.byCat[catString] = e.byCat[catString] + expenseCost
		}
	}
	e.byCat["TOTAL"] = total
	return e.byCat
}

func (e *Expense) averagePerDay() float64 {
	if len(e.payments) == 0 {
		return 0.00
	}

	byDay := map[int]int{}
	for _, payment := range e.payments {
		byDay[payment.date] = byDay[payment.date] + payment.price
	}

	allCost := 0
	for _, prices := range byDay {
		allCost = allCost + prices
	}
	return float64(allCost) / float64(len(byDay))
}
