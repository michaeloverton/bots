package main

import (
	"flag"
	"fmt"

	"github.com/michaeloverton/bots/internal/markov"
	"github.com/michaeloverton/bots/internal/texts"
)

var order int
var ref int
var count int

func main() {
	flag.IntVar(&order, "order", 1, "order of markov chain")
	flag.IntVar(&ref, "ref", 1, "reference text")
	flag.IntVar(&count, "count", 20, "output word count")
	flag.Parse()

	var reference string
	switch ref {
	case 1:
		reference = texts.Revelation
	case 2:
		reference = texts.Genesis
	case 3:
		reference = texts.Nostradamus
	case 4:
		reference = texts.EmeraldTablet
	default:
		panic("invalid reference text")
	}

	switch order {
	case 1:
		text, err := markov.OrderOne(reference, count)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
		fmt.Println("char count:", len(text))
	case 2:
		text, err := markov.OrderTwo(reference, count)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
		fmt.Println("char count:", len(text))
	default:
		fmt.Println("invalid markov chain order")
	}
}
