package main

func main() {
	cards := newDeck()
	cards = newDeckFromFile("myCards.txt")
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
