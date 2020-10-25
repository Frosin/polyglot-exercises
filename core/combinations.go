package core

import (
	"fmt"

	"github.com/fatih/structs"
)

type (
	Combination struct {
		I    string
		You  string
		They string
		He   string
		She  string
	}

	Word struct {
		rusPresent Combination
		rusPast    Combination
		rusFuture  Combination

		engBase    string
		eng3Person string
		engPast    string
	}

	Phrases struct {
		Eng string
		Rus string
	}
)

func is3Person(key string) bool {
	return key == "He" || key == "She" || key == "They"
}

func GetQuestionPresent(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusPresent)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase
		link := "do"

		if is3Person(i) {
			verb = w.eng3Person
			link = "does"
		}

		enPhrase := fmt.Sprintf("%s %s %s?", link, pronoun, verb)
		ruPhrase := fmt.Sprintf("%s %s?", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetQuestionPast(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusPast)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase
		link := "did"

		enPhrase := fmt.Sprintf("%s %s %s?", link, pronoun, verb)
		ruPhrase := fmt.Sprintf("%s %s?", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetQuestionFuture(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusFuture)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase
		link := "will"

		enPhrase := fmt.Sprintf("%s %s %s?", link, pronoun, verb)
		ruPhrase := fmt.Sprintf("%s %s?", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetStatementPresent(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusPresent)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase

		if is3Person(i) {
			verb = w.eng3Person
		}

		enPhrase := fmt.Sprintf("%s %s", pronoun, verb)
		ruPhrase := fmt.Sprintf("%s %s", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetStatementPast(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusPast)

	for i, v := range engMap {
		pronoun := v
		verb := w.engPast

		enPhrase := fmt.Sprintf("%s %s", pronoun, verb)
		ruPhrase := fmt.Sprintf("%s %s", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetStatementFuture(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusFuture)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase
		link := "will"

		enPhrase := fmt.Sprintf("%s %s %s", pronoun, link, verb)
		ruPhrase := fmt.Sprintf("%s %s", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetNegativePresent(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusPresent)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase

		if is3Person(i) {
			verb = "doesn't " + verb
		} else {
			verb = "don't " + verb
		}

		enPhrase := fmt.Sprintf("%s %s", pronoun, verb)
		ruPhrase := fmt.Sprintf("%s не %s", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetNegativePast(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusPast)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase
		link := "did not"

		enPhrase := fmt.Sprintf("%s %s %s", pronoun, link, verb)
		ruPhrase := fmt.Sprintf("%s не %s", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}

func GetNegativeFuture(w Word, rusComb, engComb Combination) []Phrases {
	resPhrases := []Phrases{}

	rusMap := structs.Map(rusComb)
	engMap := structs.Map(engComb)
	rusWordMap := structs.Map(w.rusFuture)

	for i, v := range engMap {
		pronoun := v
		verb := w.engBase
		link := "will not"

		enPhrase := fmt.Sprintf("%s %s %s", pronoun, link, verb)
		ruPhrase := fmt.Sprintf("%s не %s", rusMap[i], rusWordMap[i])

		resPhrases = append(resPhrases, Phrases{
			Eng: enPhrase,
			Rus: ruPhrase,
		})
	}

	return resPhrases
}
