package expense

import (
	"reflect"
	"testing"
)

func Test_empty_file_will_return_total_0(t *testing.T) {
	expected := map[string]int{
		"TOTAL": 0,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_day_data.txt")

	actual := expense.summaryByCat()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_expense_first_with_20_food_by_cat(t *testing.T) {
	expected := map[string]int{
		"food":  20,
		"TOTAL": 20,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_first_with_20_food.txt")

	actual := expense.summaryByCat()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}

func Test_all_data_by_cat(t *testing.T) {
	expected := map[string]int{
		"transportation":279,
		"movie":1680,
		"game":1220,
		"food":3057,
		"stationery":332,
		"TOTAL":6568,
	}

	expense := NewExpense()

	expense.read("./expense_by_day/expense_all_data.txt")

	actual := expense.summaryByCat()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}

}
