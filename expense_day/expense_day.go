package expense_day

type Expense struct {
}

func NewExpense() Expense {
	return Expense{}
}

func (e *Expense) read(fileName string) {

}

func (e *Expense) summaryByDay() map[string]int {
	return map[string]int{
		"Mon": 0,
		"Tue": 0,
		"Wed": 0,
		"Thu": 0,
		"Fri": 0,
		"Sat": 0,
		"Sun": 0,
	}
}
