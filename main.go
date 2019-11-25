package main

func main() {
	// 1. Create a new deck type
	cards := newDeck()
	// 2. Print the new deck
	cards.print()
	// 3. Shuffle the cards and print the new order
	cards.shuffleCards()
	cards.print()
	// 4. Save the deck to a file
	cards.saveToFile("my_cards")
	// 5. Retrieve the deck from the file
	cards = newDeckFromFile("my_cards")
	// 6. Deal the cards and print the current hand and the reamining cards in the deck
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
