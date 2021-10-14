package main

import (
	"os"
	"testing"
)

func Test_newDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected the first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs"{
		t.Errorf("Expected the first card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func Test_saveToDeck_and_newDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, bit got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}