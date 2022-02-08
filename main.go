package main

import (
	"fmt"
	"os"
)

func main() {
	var guess string
	dict := getDict()
	answer := dict.randomWordOfDay()
	guesses := guessList{}
	totalGuessLeft := 6

	presentation()

	for totalGuessLeft > 0 {
		fmt.Print("Guess: ")
		fmt.Scanln(&guess)
		msg, err := validateGuess(dict, guess, guesses)
		if err != nil {
			fmt.Println(msg)
		} else {
			checkedGuess, guessed := checkGuess(answer, guess)
			if guessed {
				win(answer)
				os.Exit(0)
			}
			guesses = append(guesses, checkedGuess)
			guesses.printPastGuesses()
			totalGuessLeft--
		}
	}
	lost(answer)
}

func presentation() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Try to guess the 5 letter word of the day: ")
}

func win(g string) {
	fmt.Println(g, "was the answer!")
	fmt.Println("Good job!")
}

func lost(a string) {
	fmt.Println("You lost.")
}
