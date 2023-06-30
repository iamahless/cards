package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck lenght of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Clubs, but go %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeskTestFromFile(t *testing.T) {
	_ = os.Remove("_decktesting_sd")

	deck := newDeck()
	_ = deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck lenght of 20, but got %v", len(loadedDeck))
	}

	_ = os.Remove("_decktesting_sd")
}
