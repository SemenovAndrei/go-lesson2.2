package main

import (
	"fmt"
	"github.com/i-hit/go-lesson2.2.git/pkg/card"
	"github.com/i-hit/go-lesson2.2.git/pkg/transfer"
)

func main() {
	cardSvc := card.NewService("Tinkoff")

	card1 := cardSvc.GetNewCard("visa", 1000, "RUB", "5106 2100 0000 0007")
	card2 := cardSvc.GetNewCard("visa", 100, "RUB", "5106 2100 0000 0000 6")

	list := transfer.NewService(cardSvc)

	fmt.Println(list.CardSvc)
	fmt.Println(card1, card2)

	fmt.Println(list.Card2Card("5106 2100 0000 0007", "5106 2100 0000 0000 6", 100))

	fmt.Println(card1, card2)
}
