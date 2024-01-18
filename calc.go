package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	return strings.TrimSpace(line)
}

func convertToArabic(roman1 string, roman2 string) (int, int) {
	romanNums := map[string]int{
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

	roman1 = strings.ToUpper(roman1)
	roman2 = strings.ToUpper(roman2)

	if arabic1, found1 := romanNums[roman1]; found1 {
		if arabic2, found2 := romanNums[roman2]; found2 {
			return arabic1, arabic2
		} else {
			fmt.Println("Выдача паники, так как второй операнд не является римской или арабской цифрой")
			os.Exit(1)
			return 0, 0
		}
	} else {
		fmt.Println("Выдача паники, так как первый операнд не является римской или арабской цифрой")
		os.Exit(1)
		return 0, 0
	}
}

func convertToRoman(num int) string {
	romanNumerals := []struct {
		arabic int
		roman  string
	}{
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

	result := ""
	for _, pair := range romanNumerals {
		for num >= pair.arabic {
			result += pair.roman
			num -= pair.arabic
		}
	}

	return result
}

func convertLine() (int, int, string, bool) {
	line := strings.Split(getLine(), " ")
	flag := false

	if len(line) < 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		os.Exit(1)
	} else if len(line) > 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(1)
	}

	num1, err1 := strconv.Atoi(line[0])
	num2, err2 := strconv.Atoi(line[2])
	operator := line[1]

	if err1 == nil && err2 == nil {
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			fmt.Println("Выдача паники, так как число должно быть от 1 до 10")
			os.Exit(1)
		}
	}

	if err1 == nil && err2 != nil || err1 != nil && err2 == nil {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
		os.Exit(1)
	} else if err1 != nil && err2 != nil {
		flag := true
		arabic1, arabic2 := convertToArabic(line[0], line[2])
		return arabic1, arabic2, operator, flag
	}

	return num1, num2, operator, flag
}

func calculate(num1, num2 int, operator string, flag bool) interface{} {
	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Выдача паники, так как оператор не соответствует заданию.")
		return 0
	}

	if flag {
		if result < 1 {
			return fmt.Errorf("Выдача паники, так как значения вычисления римских чисел меньше 1")
		} else {
			return convertToRoman(result)
		}
	} else {
		return result
	}
}

func main() {
	num1, num2, operator, flag := convertLine()
	result := calculate(num1, num2, operator, flag)
	fmt.Println(result)
}
