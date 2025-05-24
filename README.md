# 🎯 Number Guessing Game (CLI)

A simple CLI-based number guessing game written in Go. The computer randomly selects a number between 1 and 100, and you have to guess it within a limited number of attempts based on your selected difficulty.

📌 [Number Guessing Game on roadmap.sh](https://roadmap.sh/projects/number-guessing-game)

---

## 🧩 Features

- Random number generation
- Selectable difficulty levels: Easy, Medium, Hard
- Interactive CLI input
- Feedback on whether your guess is too high or too low
- Victory and game-over states

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go 1.18 or newer (recommended Go 1.21+)
- Terminal or PowerShell access

---

### 📦 Installation

```bash
git clone https://github.com/yourusername/number-guessing-game.git
cd number-guessing-game
go build
```

---

### ▶️ Usage

Run the compiled binary:

```bash
./number-guessing-game
```

Sample interaction:

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 30
Congratulations! You guessed the correct number in 2 attempts.
```
