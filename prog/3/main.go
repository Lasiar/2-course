package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	radius        string = "1"
	diameter      string = "2"
	circumference string = "3"
)

type Circle interface {
	area() float64
}

type Radius float64

func (r Radius) area() float64 {
	return math.Pi * float64(r*r)
}

type Diameter float64

func (d Diameter) area() float64 {
	return math.Pi * float64(d*d) / 4
}

type Circumference float64

func (c Circumference) area() float64 {
	return float64(c*c) / (4 * math.Pi)
}

func main() {

	fmt.Printf("Площадь круга: %v", computed(scanUserValue(getUserType())))
}

func computed(value interface{}) float64 {
	if v, ok := value.(Circle); ok {
		return v.area()
	}
	return 0
}

func getUserType() interface{} {

	var rawString string

	for {
		fmt.Print("1. Радиус \n2. Диаметр \n3. Длина окружности \n > ")

		if _, err := fmt.Scan(&rawString); err != nil {
			fmt.Println("Ошибка чтения: %v", err)
			continue
		}

		switch rawString {
		case radius:
			return Radius(0)
		case diameter:
			return Diameter(0)
		case circumference:
			return Circumference(0)
		default:
			fmt.Println("Введите только значение от 1 до 3")

		}

	}
}

func scanUserValue(userValue interface{}) interface{} {
	var rawString string

	for {
		fmt.Print("Введите значение: \n > ")

		if _, err := fmt.Scan(&rawString); err != nil {
			fmt.Println("Ошибка чтения: %м")
			continue
		}

		value, err := strconv.ParseFloat(strings.Replace(rawString, ",", ".", 0), 64)
		if err != nil || value == 0 {
			fmt.Printf("\n Ошибка чтения значения: %v \n", err)
			continue
		}

		if value <= 0 {
			fmt.Println("\n Значение должно быть больше еденицы")
			continue
		}

		switch userValue.(type) {
		case Radius:
			return Radius(value)
		case Diameter:
			return Diameter(value)
		case Circumference:
			return Circumference(value)
		}
	}

}
