// Testing with Go // Checks all package tests
// Files must end with _test.go and can be used by typing in "go test" in the terminal
package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d, dLen := newDeck(), 54 // New Deck and expected length of deck (default 54)

	if len(d) != dLen {
		t.Errorf("Expected deck length of %v, but got %v", dLen, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-3] != "King of Diamonds" {
		t.Errorf("Expected third last card of 'King of Diamonds', but got %v", d[len(d)-3])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // Removes any file called _decktesting

	d, dLen := newDeck(), 54
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != dLen {
		t.Errorf("Expected deck length of %v, but got %v", dLen, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
