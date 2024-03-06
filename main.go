package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Проверяет есть ли элемент в массиве
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Выполняет арифметическое действие
func action(n1, n2 int, operation string) int {
	var result int
	switch operation {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		if n2 == 0 {
			panic("Деление на ноль")
		}
		result = n1 / n2
	default:
		panic("Неверный оператор")
	}
	return result
}

// Преобразует римскую цифру в арабскую
func romanToArabic(roman string) int {
	allRoman := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	dec := 0

	for i, r := range allRoman {
		for strings.HasPrefix(roman, r) {
			dec += i
			roman = roman[len(r):]
		}
	}

	return dec
}

// Преобразует арабскую цифру в римскую
func arabicToRoman(arabic int) string {
	allRoman := []struct {
		i int
		r string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	if arabic <= 0 {
		panic("Римское число не может быть нулем или отрицательным")
	}

	roman := ""

	for arabic > 0 {
		for _, r := range allRoman {
			for arabic >= r.i {
				roman += r.r
				arabic -= r.i
			}
		}
	}

	return roman
}

// Проверяет входит ли число в римскую систему счисления
func isRoman(roman string) bool {
	romanNums := []string{
		"I", "V", "X", "L", "C", "D", "M",
	}

	splitString := strings.Split(roman, ``)

	for i := 0; i < len(splitString); i++ {
		if !contains(romanNums, splitString[i]) {
			return false
		}
	}

	return true

}

func main() {

	arabicNums := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите математическую операцию: ")
	task, _ := reader.ReadString('\n') // Принимает данные от пользователя
	task = strings.TrimSpace(task)
	splitString := strings.Split(task, ` `) // Разделение строки

	// Проверяет пральность формата вводимых данных
	if len(splitString) != 3 {
		panic(errors.New("Неверный формат математической операции"))

	}

	firstNumStr := splitString[0]
	secondNumStr := splitString[2]
	var operation = splitString[1]

	// Проверка на совпадение двух различных систем счисления
	if contains(arabicNums, firstNumStr) && contains(arabicNums, secondNumStr) {

		firstNum, _ := strconv.Atoi(firstNumStr)
		secondNum, _ := strconv.Atoi(secondNumStr) // Преобразование строки в число
		result := action(firstNum, secondNum, operation)
		fmt.Println(result)

	} else if isRoman(firstNumStr) && isRoman(secondNumStr) {

		firstNum := romanToArabic(firstNumStr)
		secondNum := romanToArabic(secondNumStr) // Преобразование строки в число
		result := arabicToRoman(action(firstNum, secondNum, operation))
		fmt.Println(result)

	} else {

		panic("Неверные числа")
	}

}
