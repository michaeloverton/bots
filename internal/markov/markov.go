package markov

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/mb-14/gomarkov"
	"github.com/pkg/errors"
)

func OrderOne(referenceText string, wordCount int) (string, error) {
	//Create a chain of order 1
	chain := gomarkov.NewChain(1)

	referenceWords := strings.Split(referenceText, " ")
	chain.Add(referenceWords)

	randSeed := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(randSeed)

	wordIndex := rander.Intn(len(referenceWords))
	word := referenceWords[wordIndex]

	text := word

	for i := 0; i < wordCount; i++ {
		nextWord, err := chain.Generate([]string{word})
		if err != nil {
			return "", errors.Wrap(err, "failed to generate order one next word")
		}
		text = fmt.Sprintf("%s %s", text, nextWord)
		word = nextWord
	}

	return text, nil
}

func OrderTwo(referenceText string, wordCount int) (string, error) {
	//Create a chain of order 2
	chain := gomarkov.NewChain(2)

	referenceWords := strings.Split(referenceText, " ")
	chain.Add(referenceWords)

	randSeed := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(randSeed)

	wordIndex := rander.Intn(len(referenceWords) - 1)
	word := referenceWords[wordIndex]
	otherWord := referenceWords[wordIndex+1]

	text := fmt.Sprintf("%s %s", word, otherWord)

	for i := 0; i < wordCount; i++ {
		nextWord, err := chain.Generate([]string{word, otherWord})
		if err != nil {
			return "", errors.Wrap(err, "failed to generate order two next word")
		}

		nextNextWord, err := chain.Generate([]string{otherWord, nextWord})
		if err != nil {
			return "", errors.Wrap(err, "failed to generate order two next word")
		}

		text = fmt.Sprintf("%s %s %s", text, nextWord, nextNextWord)

		word = nextWord
		otherWord = nextNextWord
	}

	return text, nil
}
