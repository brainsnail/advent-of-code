package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputFile = "../inputs/02.txt"

func main() {

	var validpasswords int

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	// 9-12 q: qqqxhnhdmqqqqjz

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var min, max int
		var letter, password string
		_, err := fmt.Sscanf(line, "%d-%d %1s: %s\n", &min, &max, &letter, &password)
		if err != nil {
			fmt.Println(err)
		}
		if validpassword(min, max, letter, password) {
			validpasswords++
		}
	}
	fmt.Println(validpasswords)
}

func validpassword(min, max int, letter, password string) bool {
	charactercount := strings.Count(password, letter)
	if charactercount < min || charactercount > max {
		return false
	}
	return true
}
