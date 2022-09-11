package main

func main() {
	cards := newDeck()

	// hand, reamainingCard := deal(cards, 5)

	// hand.print()
	// reamainingCard.print()
	// cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards[0])
}
