package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputFile = "../inputs/02.txt"

func main() {

	var oldPasswordPolicyMatches, corporatePasswordPolicyMatches int

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var first, second int
		var letter, password string
		_, err := fmt.Sscanf(line, "%d-%d %1s: %s\n", &first, &second, &letter, &password)
		if err != nil {
			fmt.Println(err)
		}
		if matchesSledPlaceDownTheRoad(first, second, letter, password) {
			oldPasswordPolicyMatches++
		}
		if matchesNewCorporatePolicy(first, second, letter, password) {
			corporatePasswordPolicyMatches++
		}
	}
	fmt.Printf("Old Company Passwords: %d \n", oldPasswordPolicyMatches)
	fmt.Printf("Corporate Company Passwords: %d \n", corporatePasswordPolicyMatches)
}

func matchesSledPlaceDownTheRoad(min, max int, letter, password string) bool {
	charactercount := strings.Count(password, letter)
	if charactercount < min || charactercount > max {
		return false
	}
	return true
}

func matchesNewCorporatePolicy(firstposition, secondposition int, letter, password string) bool {
	totalcount := 0
	if string(password[firstposition-1]) == letter {
		totalcount++
	}

	if string(password[secondposition-1]) == letter {
		totalcount++
	}
	return totalcount == 1
}
