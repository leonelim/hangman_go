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
	var input string

	fmt.Print("Would you like to start a new game? (y/n): ")
	_, err := fmt.Scanf("%s", &input)
	for err != nil || !isCorrectInput(input) || !strings.ContainsAny(input, "yn") {
		fmt.Println("Wrong input! try again!")
		fmt.Print("Would you like to start a new game? (y/n): ")
		_, err = fmt.Scanf("%s", &input)
	}

	input = strings.ToLower(input)
	switch input {
	case "y":
		word, err := readWordFromFile()
		if err != nil {
			fmt.Println(err)
			return
		}
		startGame(word)
	case "n":
		break
	}
}

func isCorrectInput(input string) bool {
	input = strings.ToLower(input)
	guess, _ := utf8.DecodeRuneInString(input)
	return len(input) == 1 && guess >= 97 && guess < 123
}

func startGame(word string) {
	var hint []rune
	triedLetters := make(map[rune]bool)
	var guessStr string
	var guess rune
	mistakes := 0

	for i := 0; i < len(word); i++ {
		hint = append(hint, '_')
	}

	for !isGameOver(hint, mistakes) {
		fmt.Printf("Mistakes: %d", mistakes)
		fmt.Println(art[mistakes])

		fmt.Print("tried letters: ")
		for key := range triedLetters {
			fmt.Printf("%c", key)
		}
		fmt.Println()
		fmt.Println(string(hint))

		fmt.Print("guess?: ")

		_, err := fmt.Scanln(&guessStr)
		if err != nil || !isCorrectInput(guessStr) {
			fmt.Println("incorrect input! try again!")
			continue
		}
		guess = getFirstRune(guessStr)

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
		fmt.Printf("\n\n************\n")
	}
	if mistakes == maxAttempts {
		fmt.Println(art[maxAttempts])
		fmt.Println("You lose!")
	} else {
		fmt.Println("You win!")
	}
}

func isGameOver(hint []rune, mistakes int) bool {
	return slices.Contains(hint, '_') && mistakes == maxAttempts
}

func getFirstRune(str string) rune {
	var res rune
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
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
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
