package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ruble float64

func (r ruble) String() string {
	return fmt.Sprintf("%v", strconv.FormatFloat(float64(r)/100, 'f', 2, 64))
}

var (
	energy int
	result ruble
)

func init() {
	flag.IntVar(&energy, "e", -1, "киловат в час в месяц")
	flag.Parse()
}

func main() {
	if energy > 0 {
		result = computed(energy)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)

		for {
			var err error

			fmt.Print("Введите ваши данные: ")
			scanner.Scan()

			energy, err = strconv.Atoi(scanner.Text())

			if err != nil {
				log.Printf("\n ошибка чтеиния: %v \n", err)
				continue
			}

			result = computed(energy)

			break
		}
	}
	fmt.Printf("Ваш счет: %v рублей", result)
}

type Iterator struct {
	cost   ruble
	energy int
}

func computed(e int) ruble {
	var r ruble

	iterator := [3]Iterator{{54, 200}, {70, 300}, {91, 1<<31 - 1}}

	for _, i := range iterator {
		if e > 0 {
			if e > i.energy {
				r += ruble(i.energy) * i.cost
			} else {
				r += ruble(e) * i.cost
			}
			e -= i.energy
		}
	}

	return r
}
