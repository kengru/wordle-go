package main

import (
	"fmt"
)

func main() {
	var guess string
	dict := getDict()
	answer := dict.randomWordOfDay()
	guesses := guessList{}
	totalGuessLeft := 5

	presentation()

	for totalGuessLeft > 0 {
		fmt.Print("Guess: ")
		fmt.Scanln(&guess)
		msg, err := validateGuess(dict, guess, guesses)
		if err != nil {
			fmt.Println(msg)
		} else {
			checkedGuess := checkGuess(answer, guess)
			guesses = append(guesses, checkedGuess)
			guesses.printPastGuesses()
			totalGuessLeft--
		}
	}
}

func presentation() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Try to guess the 5 letter word of the day: ")
}
