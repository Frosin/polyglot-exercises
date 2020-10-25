package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Frosin/polyglot-exercises/core"
)

var (
	engPronouns = core.Combination{
		I:    "i",
		You:  "you",
		They: "they",
		He:   "he",
		She:  "she",
	}

	rusPronouns = core.Combination{
		I:    "я",
		You:  "ты",
		They: "они",
		He:   "он",
		She:  "она",
	}

	arPhrases = []func(core.Word, core.Combination, core.Combination) []core.Phrases{
		core.GetQuestionPresent,
		core.GetQuestionPast,
		core.GetQuestionFuture,
		core.GetStatementPresent,
		core.GetStatementPast,
		core.GetStatementFuture,
		core.GetNegativePresent,
		core.GetNegativePast,
		core.GetNegativeFuture,
	}
)

func getRandomResult(w http.ResponseWriter, r *http.Request) {
	randomPhrases := arPhrases[rand.Intn(len(arPhrases))]
	randomWord := core.Words[rand.Intn(len(core.Words))]
	result := randomPhrases(randomWord, rusPronouns, engPronouns)
	randomPhrase := result[rand.Intn(len(result))]

	fmt.Fprintf(w, "%s: %s", randomPhrase.Eng, randomPhrase.Rus)
}

func main() {

	rand.Seed(time.Now().Unix())
	port := os.Getenv("PORT")
	http.HandleFunc("/", getRandomResult)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
