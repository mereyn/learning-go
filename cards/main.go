package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// // save to a file
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cardss")
	// cards.print()
}
