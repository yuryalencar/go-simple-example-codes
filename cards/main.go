package main

func main() {
	cards := newDeckFromFile("my_cards.txt")
	cards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")
}
