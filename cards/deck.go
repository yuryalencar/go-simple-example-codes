package main

import "fmt"

type deck []string

func newDeck() deck {
    cards := deck{}

    cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string {"Ace", "Two", "Three", "Four"}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)
        }
    }

    return cards
}

func (d deck) print() {
    fmt.Println("CARDS LIST")
    for _, card := range d {
        fmt.Println(card)
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}