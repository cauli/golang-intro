package emoji

import (
	"math/rand"
	"time"
)

func Random() string {
	emojiSlice := healthyEmoji()

	rand.Seed(time.Now().UnixNano()) // Initialize pseudo random generator
	randomIndex := rand.Intn(len(emojiSlice))

	return emojiSlice[randomIndex]
}

func junkFoodEmoji() []string {
	return []string{"🌭", "🌮", "🍕"}
}

func healthyEmoji() []string {
	return []string{"🍇", "🍈", "🍉", "🍊", "🍋", "🍌", "🍍", "🍎", "🍏", "🍐", "🍑", "🍒", "🍓"}
}
