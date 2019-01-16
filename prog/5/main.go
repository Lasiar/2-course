package main

import "fmt"

func main() {
	str := "ФФФФФФФ"

	for i := range str {
		fmt.Println(str[i:])
	}

}
