package main

// Import packages needed
import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

// Create a new deck type of slice of strings
type deck []string

// 1. Function of return type deck to create a new deck of 52 cards
func newDeck() deck {
	// Create a new deck type
	cards := deck{}
	// Create a slice of strings of the 4 main card suits
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// Create a slice of strings of the card values
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// Loop through the card suits
	// Note: not using the index so replacing it with an underscore
	for _, suit := range cardSuits {
		// Loop through the card values
		for _, value := range cardValues {
			// Assign/Append to the deck type of slice of strings the string value + of + suit
			// Example: Ace of Spades
			cards = append(cards, value + " of " + suit)
		}
	}
	// return the deck type slice of strings
	return cards
}

// 2. Receiver function to be called by a deck type
func (d deck) print() {
	// Loop through each index and card string in the deck type slice of strings
	for i, card := range d {
		// Print the index value and the card value (string)
		fmt.Println(i, card)
	}
}

// Receiver function to be called by a deck type that returns a string
func (d deck) toString() string {
	// Takes a slice of strings and casts them to a string separated by commas
	return strings.Join([]string(d), ",")
}

// 3. Receiver function to be called by a deck type that shuffles the current deck
func (d deck) shuffleCards() {
	// Get a new random source by passing in a new seed value generated by the current time the app is ran
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new Rand by passing in a new source
	r := rand.New(source)
	// Loop through each index of the deck type slice of strings
	for i := range d {
		// Create a new source position
		newPosition := r.Intn(len(d)-1)
		// Takes the value at d[newPosition] and replaces the value at d[i] with it
		// Takes the value at d[i] and replaces the value at d[newPosition] with it
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// 4. Receiver function to be called by a deck type that returns an ioutil error
func (d deck) saveToFile(filename string) error {
	// Writes the file to the current directory with the filename passed in
	// Converts the current deck type slice of strings to a string and then to a slice of bytes
	// Sets the file permissions
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// 5. Function of type deck to read from a file
func newDeckFromFile(filename string) deck {
	// returns a slice of bytes and an error (if there is one)
	bs, err := ioutil.ReadFile(filename)
	// Check if there is an error
	if err != nil {
		// Print the error
		fmt.Println("Error: ", err)
		// Exit the application
		os.Exit(1)
	}
	// Convert the slice of bytes to a string 
	// Convert the string to a slice of strings where each value is separated by a comma
	s := strings.Split(string(bs), ",")
	// Convert the slice of strings to a deck type of slice of strings
	// Return the deck type
	return deck(s)
}

// 6. Function to deal and return two values of type deck 
func deal(d deck, handSize int) (deck, deck) {
	// Return the values from 0 and up to but NOT including the value of handSize
	// Return the values starting at/including the value of handSize to the end of slice of strings
	return d[:handSize], d[handSize:]
}