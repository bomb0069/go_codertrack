package expense_day

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

	expense.read("expense_day_data.txt")

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

	expense.read("expense_first_with_20_food.txt")

	actual := expense.summaryByDay()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}
