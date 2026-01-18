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

func startGame() error {
	word, err := readWordFromFile()
	if err != nil {
		return errors.New("failed to open file")
	}
	var hint []rune

	for i := 0; i < len(word); i++ {
		hint = append(hint, '_')
	}
	var guess rune
	mistakes := 0

	for slices.Contains(hint, '_') && mistakes < 6 {
		fmt.Println(string(hint))
		fmt.Printf("Mistakes: %d", mistakes)

		fmt.Print("guess?: ")

		_, err = fmt.Scanf("%c\n", &guess)
		if err != nil {
			panic("error reading from stdin")
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
	}
	return nil
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
	for i {
		i = scanner.Scan()
		lines = append(lines, scanner.Text())
	}
	randInt := rand.Int() % len(lines)
	return lines[randInt], nil
}

func main() {
	var input string
	fmt.Println("Would you like to start a new game? (y/n)")
	_, err := fmt.Scanf("%s", &input)
	for err != nil {
		fmt.Println("Wrong input! try again!")
		_, err = fmt.Scanf("%d", &input)
	}
	switch input {
	case "Y":
	case "y":
		err := startGame()
		if err != nil {
			return
		}
	case "N":
	case "n":
		break
	}
}
