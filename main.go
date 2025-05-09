package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your choice: ")
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	chances := 0
	switch choice {
	case "1":
		chances = 10
		fmt.Println("Great! You have selected the Easy difficulty level.")
	case "2":
		chances = 5
		fmt.Println("Great! You have selected the Medium difficulty level.")
	case "3":
		chances = 3
		fmt.Println("Great! You have selected the Hard difficulty level.")
	default:
		fmt.Println("Invalid choice. Defaulting to Medium difficulty.")
		chances = 5
	}

	fmt.Println("Let's start the game!")

	attempts := 0
	for attempts < chances {
		fmt.Printf("Enter your guess (%d remaining): ", chances-attempts)
		scanner.Scan()
		guessStr := strings.TrimSpace(scanner.Text())
		guess, err := strconv.Atoi(guessStr)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		attempts++
		if guess == target {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		} else if guess < target {
			fmt.Println("Incorrect! The number is greater than", guess)
		} else {
			fmt.Println("Incorrect! The number is less than", guess)
		}
	}

	fmt.Printf("Game over! You've used all your chances. The correct number was %d.\n", target)
}
