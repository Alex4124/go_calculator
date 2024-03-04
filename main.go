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
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	return romanNumerals[roman]
}

// Преобразует арабскую цифру в римскую
func arabicToRoman(arabic int) string {
	arabicNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
	}
	if arabic <= 0 {
		panic("Римское число не может быть нулем или отрицательным")
	}
	return arabicNumerals[arabic]
}

func main() {

	romanNums := []string{
		"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	}

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

	} else if contains(romanNums, firstNumStr) && contains(romanNums, secondNumStr) {

		firstNum := romanToArabic(firstNumStr)
		secondNum := romanToArabic(secondNumStr) // Преобразование строки в число
		result := arabicToRoman(action(firstNum, secondNum, operation))
		fmt.Println(result)

	} else {

		panic("Неверные числа")
	}

}
