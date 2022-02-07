package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type dict []string

func getDict() dict {
	bs, err := os.ReadFile("data/dict.txt")
	if err != nil {
		log.Fatal("Cannot read the dictionary.")
	}
	s := strings.Split(string(bs), ",")
	return dict(s)
}

func (d dict) randomWordOfDay() string {
	today := time.Now()
	seed := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	source := rand.NewSource(seed.UnixNano())
	r := rand.New(source)

	return d[r.Intn(len(d)-1)]
}

func (d dict) isWordInDictionary(s string) bool {
	for _, v := range d {
		if v == s {
			return true
		}
	}
	return false
}
