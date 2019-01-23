package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 1 {
		log.Println("Укажите файл")
		return
	}
	pathFile := os.Args[1]


}


func processed(path string) error {
	file,err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	regexp.MustCompile("")


	for scanner.Scan() {
		text := scanner.Text()
	}

}

func validString(string) {
	return
}