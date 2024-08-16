package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
)

type GameState struct {
	words      []string
	chosenWord string
	showWord   string
}

type Request struct {
	Letter string `json:"letter"`
}

var gameState GameState

func getWord(c echo.Context) error {
	wrd := fmt.Sprintf("word: %s", gameState.showWord)
	return c.String(http.StatusOK, wrd)
}

func insertWord(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "BAD REQUEST - Erro ao processar requisição")
	}

	processLetter(req.Letter)

	word := fmt.Sprintf("word: %s", gameState.showWord)
	return c.String(http.StatusOK, word)
}

func newGame(c echo.Context) error {
	gameState.showWord = ""

	initGame()

	word := fmt.Sprintf("word: %s", gameState.showWord)
	return c.String(http.StatusOK, word)
}

func main() {
	initGame()
	e := echo.New()
	e.GET("/word", getWord)
	e.GET("/word/new", newGame)
	e.POST("/word", insertWord)
	e.Logger.Fatal(e.Start(":1323"))
}

func initGame() {
	gameState.words = []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince", "raspberry", "strawberry", "tangerine", "ugli", "vanilla", "watermelon", "xigua", "yellow", "zucchini"}

	// Gera indice aleatório
	reader := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	index := reader.Intn(len(gameState.words))

	gameState.chosenWord = gameState.words[index]

	for range gameState.chosenWord {
		gameState.showWord += "_"
	}
}

func processLetter(letter string) {
	letter = strings.ToLower(letter)
	if strings.Contains(gameState.chosenWord, letter) {
		for i, c := range gameState.chosenWord {
			if string(c) == letter {
				gameState.showWord = gameState.showWord[:i] + letter + gameState.showWord[i+1:]
			}
		}
	}

	if !strings.Contains(gameState.showWord, "_") {
		txt := "Parabéns! Você acertou a palavra!\n"
		txt += "A palavra era " + gameState.chosenWord
		gameState.showWord = txt
	}
}
