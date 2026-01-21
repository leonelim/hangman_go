package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

const wordsFilePath string = "resources/words.txt"

func main() {
	var input string

	fmt.Print("Would you like to start a new game? (y/n): ")
	_, err := fmt.Scanf("%s", &input)
	for err != nil || !isCorrectInput(input) || !strings.ContainsAny(input, "yn") {
		fmt.Println("Wrong input! try again!")
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
	return len(input) == 1
}

func startGame(word string) {
	var hint []rune

	for i := 0; i < len(word); i++ {
		hint = append(hint, '_')
	}
	var guess rune
	var guessStr string
	mistakes := 0

	for slices.Contains(hint, '_') && mistakes < 6 {
		fmt.Println(string(hint))
		fmt.Printf("Mistakes: %d", mistakes)
		fmt.Println(art[mistakes])

		fmt.Print("guess?: ")

		_, err := fmt.Scanln(&guessStr)
		if err != nil {
			panic("error reading from stdin")
		}
		for _, char := range guessStr {
			guess = char
			break
		}

		if strings.ContainsRune(word, guess) {
			count := 0
			for _, char := range word {
				if char == guess {
					hint[count] = guess
				}
				count++
			}
		} else {
			mistakes++
		}
		if mistakes == 6 {
			fmt.Println(art[6])
		}
	}
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
