package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbersdict = map[string]string{

	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func combineNumbers(first int, second int) int {
	strA := strconv.Itoa(first)
	strB := strconv.Itoa(second)
	combinedStr := strA + strB
	combinedInt, err := strconv.Atoi(combinedStr)
	if err != nil {
		panic(err)
	}

	return combinedInt
}

func keepNumbers(s string) []int {
	var sb []int

	for _, r := range s {
		if unicode.IsDigit(r) {
			sb = append(sb, int(r-'0'))
		}
	}

	return sb
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var sum int
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Text = ", text)
		text = convertNumberStringsToInts(text)
		fmt.Println("Text after string to num = ", text)
		numbers := keepNumbers(text)
		if len(numbers) == 1 {
			number := combineNumbers(numbers[0], numbers[0])
			sum += number
			fmt.Println("Number = ", number)
		} else {
			number := combineNumbers(numbers[0], numbers[len(numbers)-1])
			sum += number
			fmt.Println("Number = ", number)
		}

		fmt.Println("Numbers = ", numbers)

		fmt.Println("Sum = ", sum)
	}
	fmt.Println("Sum = ", sum)
	err = file.Close()
	if err != nil {
		panic(err)
		return
	}
}

func convertNumberStringsToInts(text string) string {
	for key, value := range numbersdict {
		text = strings.Replace(text, key, value, -1)
	}
	return text
}
