package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected deck length of %v, but got %v", len(d), len(loadedDeck))
	}

	if loadedDeck[0] != d[0] {
		t.Errorf("Expected first card of %v, but got %v", d[0], loadedDeck[0])
	}

	if loadedDeck[len(d)-1] != d[len(d)-1] {
		t.Errorf("Expected last card of %v, but got %v", d[0], loadedDeck[0])
	}

	os.Remove("_decktesting")
}
