package main

import (
	"fmt"
	"github.com/MattSilvaa/journeygenius/server/internal/ai"
	"os"
)

func main() {
	key := os.Getenv("AI_API_KEY")

	fmt.Println("Hello World")
	ai.SetupClient(key)
	fmt.Println("Goodbye World")
}
