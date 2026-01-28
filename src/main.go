package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
	"unicode/utf8"
)

const (
	wordsFilePath string = "resources/words.txt"
	maxAttempts   int    = 6
)

func main() {
	fmt.Print("Would you like to start a new game? (y/n): ")
	input, err := readInput()
	for err != nil || input != 'y' && input != 'n' {
		fmt.Println("Wrong input! try again!")
		fmt.Print("Would you like to start a new game? (y/n): ")
		input, err = readInput()
	}
	switch input {
	case 'y':
		word, err := readWordFromFile()
		if err != nil {
			fmt.Println("error reading from file")
			return
		}
		startGame(word)
	case 'n':
		return
	}
}

func readInput() (rune, error) {
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil || !isCorrectInput(input) {
		return 0, errors.New("incorrect input")
	}
	return getFirstRune(input), nil
}

func isCorrectInput(input string) bool {
	input = strings.ToLower(input)
	guess, _ := utf8.DecodeRuneInString(input)
	return guess >= 97 && guess < 123
}

func startGame(word string) {
	var hint []rune
	triedLetters := make(map[rune]bool)
	mistakes := 0

	for range word {
		hint = append(hint, '_')
	}

	for !isGameOver(hint, mistakes) {
		fmt.Printf("Mistakes: %d\n", mistakes)
		fmt.Println(art[mistakes])

		fmt.Print("tried letters: ")
		for key := range triedLetters {
			fmt.Printf("%c", key)
		}
		fmt.Printf("\n%s", string(hint))

		fmt.Print("guess?: ")

		guess, err := readInput()
		if err != nil {
			fmt.Println("incorrect input! try again!")
			continue
		}

		if strings.ContainsRune(word, guess) {
			hintIndex := 0
			for _, wordChar := range word {
				if wordChar == guess {
					hint[hintIndex] = guess
				}
				hintIndex++
			}
		} else if !triedLetters[guess] {
			mistakes++
			triedLetters[guess] = true
		}
		fmt.Printf("\n************\n")
	}
	if mistakes == maxAttempts {
		fmt.Println(art[maxAttempts])
		fmt.Println("You lose!")
	} else {
		fmt.Println("You win!")
	}
}

func isGameOver(hint []rune, mistakes int) bool {
	return !slices.Contains(hint, '_') || mistakes == maxAttempts
}

func getFirstRune(str string) rune {
	res := rune(0xFFFD)
	for _, char := range str {
		res = char
		break
	}
	return res
}

func readWordFromFile() (string, error) {
	file, err := os.Open(wordsFilePath)
	if err != nil {
		return "", errors.New("failed to open file")
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	i := true
	counter := 0
	for i {
		counter++
		i = scanner.Scan()
		lines = append(lines, scanner.Text())
	}
	randInt := rand.IntN(counter - 1)
	return lines[randInt], nil
}
