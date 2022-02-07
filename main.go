package main

import (
	"errors"
	"fmt"
)

func main() {
	dict := getDict()
	answer := dict.randomWordOfDay()
	var guess string

	presentation()
	fmt.Scanln(&guess)

	msg, err := validateGuess(dict, guess)

	if err != nil {
		fmt.Println(msg)
	}

	fmt.Println(answer)
}

func presentation() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Try to guess the 5 letter word of the day: ")
}

func validateGuess(d dict, g string) (string, error) {
	if len(g) != 5 {
		return "The guess needs to have 5 letters.", errors.New("guess_size_error")
	}
	if !d.isWordInDictionary(g) {
		return "That is not a word in our dictionary.", errors.New("doesnt_exist_error")
	}
	return "", nil
}
