package basic

var hra int = 5
var tax int = 2
var Basic int = 1

func Calculation() (allowance int, deduction int) {
	allowance = Basic * hra
	deduction = Basic * tax
	return
}
