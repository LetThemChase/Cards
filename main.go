package main

func main() {
	//var card string = "Ace of Diamonds"
	cards := newDeck()

	cards.saveToFile("hand")
	newDeckFromFile("hand").print()
}
