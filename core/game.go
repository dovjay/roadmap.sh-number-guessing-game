package core

import (
	"fmt"
	"number-guessing-game/game"
)

func Start() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	choice := game.ScanInput("Enter your choice: ")
	difficulty := game.ParseDifficulty(choice)

	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty.Name)
	fmt.Println("Let's start the game!")

	target := game.GenerateRandomNumber(1, 100)
	attempts := 0

	for attempts < difficulty.Attempts {
		remaining := difficulty.Attempts - attempts
		guessStr := game.ScanInput(fmt.Sprintf("Enter your guess (%d remaining): ", remaining))
		guess, err := game.ParseInt(guessStr)
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
