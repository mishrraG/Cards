package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func newDeck() deck {
	cards := deck{}
	var cardValues deck
	var cardSuits deck
	cardSuits = []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues = []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
func newDeckFromFile(fileName string) deck {
	bs, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error hain bhaisahab : ", error)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
func randomIntn(min, max int) int {
	return min + rand.Intn(max-min)
}
