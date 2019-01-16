package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {

	text := getUserNumber()

	counter := computed(text)

	fmt.Printf("%v раз вчтречается \"3\"", counter)
}

func getUserNumber() string {
	for {

		var rawText string
		fmt.Print("Введите число \n > ")

		if _, err := fmt.Scan(&rawText); err != nil {
			log.Println("\nОшибка чтения")
			continue
		}

		if !isDigit(rawText) {
			log.Printf("Введите число")
			continue
		}

		return rawText
	}
}

func computed(text string) int {
	counter := 0

	for _, v := range text {
		if v == '3' {
			counter++
		}
	}
	return counter
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}