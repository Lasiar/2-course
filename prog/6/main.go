package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	arr := newArray()
	fmt.Println("Несортированный массив: \t", getString(arr))
	sort.Ints(arr)
	fmt.Println("Сортированный массив: \t",getString(arr))
}

func getString(arr []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(arr), " "), ", "), "[]")
}

func newArray() []int {
	rand.Seed(time.Now().UTC().UnixNano())

	array := make([]int, 20, 20)

	for i := range array {
		randInt := rand.Int31n(3)
		array[i] = int(randInt)
	}

	return array
}
