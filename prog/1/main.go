package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	weight float64
	result float64
)

const diff float64 = 0.73

func init() {
	flag.Float64Var(&weight, "w", -1, "укажите ваш вес")
	flag.Parse()
}

func main() {
	if weight > 0 {
		result = computed(weight, diff)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)

		for {
			var err error

			fmt.Print("Введите ваш вес: ")
			scanner.Scan()
			rowString := scanner.Text()

			rowString = strings.Replace(rowString,",", ".", 1)

			weight, err = strconv.ParseFloat(rowString, 64)

			if err != nil {
				log.Printf("\n ошибка чтеиния, введите цифры, разделяя десятчиную часть точкой: %v \n", err)
				continue
			}

			result = computed(weight, diff)

			break
		}
	}
	fmt.Printf("Ваш вес на луне: %v", result)
}

func computed(weigth, diff float64) float64 {
	return weigth * diff
}
