package expense

import (
	"reflect"
	"testing"
)

func Test_empty_file_will_return_all_0(t *testing.T) {
	expected := map[string]int{
		"Mon": 0,
		"Tue": 0,
		"Wed": 0,
		"Thu": 0,
		"Fri": 0,
		"Sat": 0,
		"Sun": 0,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_day_data.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_expense_first_with_20_food(t *testing.T) {
	expected := map[string]int{
		"Mon": 0,
		"Tue": 0,
		"Wed": 0,
		"Thu": 20,
		"Fri": 0,
		"Sat": 0,
		"Sun": 0,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_first_with_20_food.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}

func Test_first_day(t *testing.T) {
	expected := map[string]int{
		"Mon": 0,
		"Tue": 0,
		"Wed": 0,
		"Thu": 227,
		"Fri": 0,
		"Sat": 0,
		"Sun": 0,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_first_day.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}

func Test_first_week(t *testing.T) {
	expected := map[string]int{
		"Mon": 541,
		"Tue": 237,
		"Wed": 62,
		"Thu": 227,
		"Fri": 453,
		"Sat": 20,
		"Sun": 54,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_first_week.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}

func Test_more_than_one_week(t *testing.T) {
	expected := map[string]int{
		"Mon": 541,
		"Tue": 237,
		"Wed": 62,
		"Thu": 227,
		"Fri": 460,
		"Sat": 20,
		"Sun": 54,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_more_than_one_week.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}

func Test_all_data(t *testing.T) {
	expected := map[string]int{
		"Mon": 1587,
		"Tue": 915,
		"Wed": 932,
		"Thu": 1565,
		"Fri": 1184,
		"Sat": 287,
		"Sun": 98,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_all_data.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}
