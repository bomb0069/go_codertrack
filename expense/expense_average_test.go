package expense

import (
	"math"
	"testing"
)

const TOLERANCE = 0.000001

func Test_average_of_empty_file_is_0(t *testing.T) {
	expected := 0.00

	expense := NewExpense()

	expense.read("./expense_data/expense_day_data.txt")

	actual := expense.averagePerDay()

	if diff := math.Abs(expected - actual); diff > TOLERANCE {
		t.Errorf("%.2f != %.2f", expected, actual)
	}
}

func Test_average_20_food_is_20(t *testing.T) {
	expected := 20.00

	expense := NewExpense()

	expense.read("./expense_data/expense_first_with_20_food.txt")

	actual := expense.averagePerDay()

	if diff := math.Abs(expected - actual); diff > TOLERANCE {
		t.Errorf("%.2f != %.2f", expected, actual)
	}
}

func Test_average_for_first_day(t *testing.T) {
	expected := 227.00

	expense := NewExpense()

	expense.read("./expense_data/expense_first_day.txt")

	actual := expense.averagePerDay()

	if diff := math.Abs(expected - actual); diff > TOLERANCE {
		t.Errorf("%.2f != %.2f", expected, actual)
	}
}

func Test_average_for_first_week(t *testing.T) {
	expected := 227.7142857

	expense := NewExpense()

	expense.read("./expense_data/expense_first_week.txt")

	actual := expense.averagePerDay()

	if diff := math.Abs(expected - actual); diff > TOLERANCE {
		t.Errorf("%.2f != %.2f", expected, actual)
	}
}
