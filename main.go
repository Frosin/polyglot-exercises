package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
	var finalForm = core.Form

	if r.URL.Path != "/" {
		return
	}
	translateText := r.FormValue("translate")
	resultText := r.FormValue("result")
	phraseText := r.FormValue("phrase")

	if resultText != "" && resultText != translateText {
		finalForm = strings.ReplaceAll(finalForm, "%phrase%", phraseText)
		finalForm = strings.ReplaceAll(finalForm, "%result%", resultText)
		finalForm = strings.ReplaceAll(finalForm, "%correct%", "Правильный перевод: "+resultText)
		fmt.Fprintf(w, "%s", finalForm)
		return
	}

	randomPhrases := arPhrases[rand.Intn(len(arPhrases))]
	randomWord := core.Words[rand.Intn(len(core.Words))]
	result := randomPhrases(randomWord, rusPronouns, engPronouns)
	randomPhrase := result[rand.Intn(len(result))]

	finalForm = strings.ReplaceAll(finalForm, "%phrase%", randomPhrase.Rus)
	finalForm = strings.ReplaceAll(finalForm, "%result%", randomPhrase.Eng)
	finalForm = strings.ReplaceAll(finalForm, "%correct%", "")
	finalForm = strings.ReplaceAll(finalForm, "is-invalid", "")
	fmt.Fprintf(w, "%s", finalForm)
}

func main() {

	rand.Seed(time.Now().Unix())
	port := os.Getenv("PORT")
	http.HandleFunc("/", getRandomResult)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
