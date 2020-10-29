package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/mb-14/gomarkov"
	"github.com/michaeloverton/bots/internal/texts"
	"github.com/pkg/errors"
)

var order int
var ref int

func main() {
	flag.IntVar(&order, "order", 1, "order of markov chain")
	flag.IntVar(&ref, "ref", 1, "reference text")
	flag.Parse()

	fmt.Printf("markov chain order %d\n", order)

	var reference string
	switch ref {
	case 1:
		reference = texts.Revelation
	case 2:
		reference = texts.Genesis
	default:
		panic("invalid reference text")
	}

	switch order {
	case 1:
		text, err := orderOne(reference)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
		fmt.Println(len(text))
	case 2:
		text, err := orderTwo(reference)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
		fmt.Println(len(text))
	default:
		fmt.Println("invalid markov chain order")
	}
}

func orderOne(referenceText string) (string, error) {
	//Create a chain of order 1
	chain := gomarkov.NewChain(1)

	referenceWords := strings.Split(referenceText, " ")
	chain.Add(referenceWords)

	randSeed := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(randSeed)

	wordIndex := rander.Intn(len(referenceWords))
	word := referenceWords[wordIndex]

	text := word

	for i := 0; i < 30; i++ {
		nextWord, err := chain.Generate([]string{word})
		if err != nil {
			return "", errors.Wrap(err, "failed to generate order one next word")
		}
		text = fmt.Sprintf("%s %s", text, nextWord)
		word = nextWord
	}

	return text, nil
}

func orderTwo(referenceText string) (string, error) {
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

	for i := 0; i < 20; i++ {
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
