package main

func main() {
	cards := newDeck()

	cards.shuffle()

	_ = cards.saveToFile("my_cards")

	cards.print()
}
