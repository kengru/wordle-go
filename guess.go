package main

import (
	"errors"
	"fmt"
	"strings"
)

type letterGuess struct {
	l string
	c color
}
type fullGuess struct {
	lg   []letterGuess
	word string
}
type guessList []fullGuess
type color string

const (
	Grey   color = "\033[37m"
	Yellow color = "\033[33m"
	Green  color = "\033[32m"
)

func validateGuess(d dict, g string, gl guessList) (string, error) {
	if len(g) != 5 {
		return "The guess needs to have 5 letters.", errors.New("guess_size_error")
	}
	if !d.isWordInDictionary(g) {
		return "That is not a word in our dictionary.", errors.New("doesnt_exist_error")
	}
	if gl.wordInGuessList(g) {
		return "You have already guessed that word.", errors.New("repeated_guess")
	}
	return "", nil
}

func checkGuess(a string, g string) fullGuess {
	gResult := []letterGuess{}
	for idx, letterRune := range g {
		var color color
		l := string(letterRune)
		la := string(a[idx])
		if l == la {
			color = Green
		} else if strings.Contains(a, l) {
			color = Yellow
		} else {
			color = Grey
		}

		newL := letterGuess{l, color}
		gResult = append(gResult, newL)
	}
	return fullGuess{gResult, g}
}

func (fg fullGuess) printGuess() {
	for _, l := range fg.lg {
		fmt.Print(l.c, l.l)
	}
	fmt.Println(Grey)
}

func (gl guessList) printPastGuesses() {
	for _, g := range gl {
		g.printGuess()
	}
}

func (gl guessList) wordInGuessList(w string) bool {
	for _, g := range gl {
		if g.word == w {
			return true
		}
	}
	return false
}
