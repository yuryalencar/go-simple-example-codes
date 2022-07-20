package main

func main() {
    cards := deck{"Ace of Diamonds", newCard()}
    cards = append(cards, "Ace of Hearts")

    cards.print();
}

func newCard() string {
    return "Ace of Spades"
}