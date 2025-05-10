package game

import "number-guessing-game/models"

func ParseDifficulty(choice string) models.Difficulty {
	switch choice {
	case "1":
		return models.Difficulty{Name: "Easy", Attempts: 10}
	case "2":
		return models.Difficulty{Name: "Medium", Attempts: 5}
	case "3":
		return models.Difficulty{Name: "Hard", Attempts: 3}
	default:
		return models.Difficulty{Name: "Medium", Attempts: 5}
	}
}
