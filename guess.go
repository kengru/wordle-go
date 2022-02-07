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
type guessList []letterGuess
type color string

const (
	Grey   color = "\033[37m"
	Yellow color = "\033[33m"
	Green  color = "\033[32m"
)

func validateGuess(d dict, g string) (string, error) {
	if len(g) != 5 {
		return "The guess needs to have 5 letters.", errors.New("guess_size_error")
	}
	if !d.isWordInDictionary(g) {
		return "That is not a word in our dictionary.", errors.New("doesnt_exist_error")
	}
	return "", nil
}

func checkGuess(a string, g string) guessList {
	gResult := guessList{}
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
	return gResult
}

func (lg guessList) printGuess() {
	for _, l := range lg {
		fmt.Print(l.c, l.l)
	}
	fmt.Println(Grey)
}
