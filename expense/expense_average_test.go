package expense

import (
	"math"
	"testing"
)

func Test_average_of_empty_file_is_0(t *testing.T) {
	expected := 0.00
	const TOLERANCE = 0.000001

	expense := NewExpense()

	expense.read("./expense_data/expense_day_data.txt")

	actual := expense.averagePerDay()

	if diff := math.Abs(expected - actual); diff > TOLERANCE {
		t.Errorf("%.2f != %.2f", expected, actual)
	}
}
