package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type Istilah struct {
	Nama string
}

func createAkronim(phrase string) string {
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return unicode.IsSpace(r) || r == '-'
	})

	akronim := ""

	for _, word := range words {
		akronim += strings.ToUpper(string(word[0]))
	}

	return akronim
}

func Akronim(istilahs []Istilah) {
	var wg sync.WaitGroup

	for _, istilah := range istilahs {
		wg.Add(1)
		go func(istilah Istilah) {
			defer wg.Done()
			acronym := createAkronim(istilah.Nama)
			fmt.Printf("%s - %s\n", istilah.Nama, acronym)
		}(istilah)
	}

	wg.Wait()
}
