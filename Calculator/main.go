package main

import (
	"fmt"
	"os"
	"strconv"
)

func roman_convert(roman_num string) int {

	var num = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	out := 0
	length := len(roman_num)

	for i := 0; i < length; i++ {
		c := string(roman_num[i])
		vc, err := num[c]
		if !err {
			fmt.Print(fmt.Errorf("error occured: roman num has to be wrote in right order"))
			os.Exit(2)
		}
		if i < length-1 {
			next := string(roman_num[i+1])
			vcnext := num[next]

			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}

		} else {
			out += vc
		}

	}

	return out
}

func num_to_roman(num int) string {

	var num_convert = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	out := ""
	for num > 0 {
		max := max_decimal(num)
		out += num_convert[max]
		num -= max
	}

	return out
}

func max_decimal(num int) int {
	var decimal_table = []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	for _, v := range decimal_table {
		if v <= num {
			return v

		}
	}
	return 1
}

func math_err(x, y int, operand string) {
	convert_x := num_to_roman(x)
	convert_y := num_to_roman(y)
	fmt.Print(fmt.Errorf("error: \n %s canot be %s by %s", convert_x, operand, convert_y))
	os.Exit(1)
}

func divide(x, y int, param string) {

	res := "deviding: "

	divide_result := x / y

	switch param {
	case "rom":

		if x < y {
			math_err(x, y, "devided")
		}
		converted_result := num_to_roman(divide_result)
		fmt.Print(res, converted_result)
		return

	case "num":

		fmt.Print(res, divide_result)
	}
}

func multiplicate(x, y int, param string) {

	res := "multiplication: "
	multiplication := x * y

	switch param {
	case "rom":
		convert := num_to_roman(multiplication)
		fmt.Print(res, convert)
		return

	case "num":
		fmt.Print(res, multiplication)
	}
}

func add(x, y int, param string) {

	sum := x + y
	res := "sum: "
	switch param {
	case "rom":
		convert := num_to_roman(sum)
		fmt.Println(res, convert)
		return
	case "num":
		fmt.Print(res, sum)
	}
}

func substract(x, y int, param string) {

	res := "sub: "
	substraction := x - y

	switch param {
	case "rom":
		if x <= y {
			math_err(x, y, "substracted")
		}
		convert := num_to_roman(substraction)
		fmt.Print(res, convert)
		return

	case "num":
		fmt.Print(res, substraction)
	}
}

func check(x, y int) bool {

	if x > 10 || y > 10 {
		fmt.Print(fmt.Errorf("error, expected x and y to be lower than 10"))
		return false
	}

	if x < 0 || y < 0 {
		fmt.Println(fmt.Errorf("error, expected x and y to be greater than 0"))
		return false
	}

	return true
}

func check_get(x, y int, operator, param string) any {

	error_check := check(x, y)
	var answer string

	switch error_check {

	case false:
		os.Exit(0)

	}
	switch operator {
	case "+":
		add(x, y, param)

	case "-":
		substract(x, y, param)

	case "*":
		multiplicate(x, y, param)

	case "/":
		divide(x, y, param)
	default:
		fmt.Print(fmt.Errorf("error, operator %s not supported", operator))
		os.Exit(0)
	}
	fmt.Print("\ncontinue? (y / n)")
	fmt.Scan(&answer)
	switch answer {
	case "y":
		main()
	case "n":
		os.Exit(0)
	}
	return "end of call"
}

func check_for_int(x, y, operator string) {

	switch operator {
	case "":
		fmt.Print(fmt.Errorf("error, expected operator to be not empty"))
		os.Exit(0)
	}

	int_x, errx := strconv.Atoi(x)
	int_y, erry := strconv.Atoi(y)

	if errx == nil && erry == nil {
		param := "num"
		check_get(int_x, int_y, operator, param)

	}

	if errx != nil && erry != nil {
		param := "rom"
		check_for_roman(x, y, operator, param)

	} else {
		fmt.Println(fmt.Errorf("error ocured, expected x and y are both same type"))
		os.Exit(0)
	}
}

func check_for_roman(x, y, operator, param string) {

	x_convert := roman_convert(x)
	y_convert := roman_convert(y)
	check_get(x_convert, y_convert, operator, param)

}

func main() {
	var x string
	var operator string
	var y string
	fmt.Print("Type expression, operands are separated by spaces: \n")
	fmt.Scanln(&x, &operator, &y)
	if len(x) != 0 && len(y) != 0 && len(operator) != 0 {
		check_for_int(x, y, operator)

	} else {
		fmt.Println("error, expected x, y, and operator to be not empty")
		os.Exit(0)
	}
}
