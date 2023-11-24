package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		panic("Pembagian oleh nol tidak diizinkan")
	}
	return a / b, nil
}

func performOperation(operator string, a, b float64) (float64, error) {
	switch operator {
	case "a":
		return add(a, b), nil
	case "b":
		return subtract(a, b), nil
	case "c":
		return multiply(a, b), nil
	case "d":
		return divide(a, b)
	default:
		panic("Operasi aritmatika tidak valid")
	}
}

func getUserInput(prompt string) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	inputStr := scanner.Text()

	value, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		panic("Input tidak valid")
	}

	return value
}

func calculator() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	fmt.Println("Pilih operasi aritmatika:")
	fmt.Println("a> Penjumlahan (+)")
	fmt.Println("b> Pengurangan (-)")
	fmt.Println("c> Perkalian (*)")
	fmt.Println("d> Pembagian (/)")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan pilihan operasi: ")
	scanner.Scan()
	operator := strings.ToLower(scanner.Text())

	a := getUserInput("Masukkan angka pertama: ")
	b := getUserInput("Masukkan angka kedua: ")

	result, err := performOperation(operator, a, b)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Hasil %s %.2f dan %.2f adalah %.2f\n", operator, a, b, result)
}
