package main

func main() {
	//var card string = "Ace of Diamonds"
	cards := newDeck()

	hand, remaningCards := deal(cards, 5)

	hand.print()
	remaningCards.print()
	hand.saveToFile("hand")
	remaningCards.saveToFile("remainingCards")
}
