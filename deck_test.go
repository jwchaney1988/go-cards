package main

// Import packages needed
import (
	"os"
	"testing"
)

// Function to test the creation of a new deck
func TestNewDeck(t *testing.T) {
	// Get a new deck of cards
	d := newDeck()
	// Verify that there are 52 cards
	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d))
	}
	// Verify that the first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	// Verify that the last card is King of Clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

// Function to test the saving and retrieving a file with a deck of cards
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete the _decktesting file
	os.Remove("_decktesting")
	// Get a new deck of cards
	d := newDeck()
	// Save the deck to file
	d.saveToFile("_decktesting")
	// Retrieve the deck from file
	loadedDeck := newDeckFromFile("_decktesting")
	// Verify that there are 52 cards
	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(loadedDeck))
	}
	// Delete the _decktesting file
	os.Remove("_decktesting")
}