package main

func main() {
	//var card string = "Ace of Diamonds"
	cards := newDeck()

	cards.shuffle()
	cards.print()
}
