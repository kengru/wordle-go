package main

import (
	"fmt"
)

func main() {
	var guess string
	dict := getDict()
	answer := dict.randomWordOfDay()
	totalGuessLeft := 5

	presentation()

	for totalGuessLeft > 0 {
		fmt.Scanln(&guess)
		msg, err := validateGuess(dict, guess)
		if err != nil {
			fmt.Println(msg)
		} else {
			guessCheck := checkGuess(answer, guess)
			guessCheck.printGuess()
			totalGuessLeft--
		}
	}

	fmt.Println(answer)
}

func presentation() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Try to guess the 5 letter word of the day: ")
}
