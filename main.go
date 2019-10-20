package main

func main() {

	cards := newDeck()
	cards.shuffle()

	// fmt.Println(cards.toString())
	//cards.saveToFile("my_cards.txt")

	// cards := newDeckFromFile("my.txt")
	cards.print()
}
