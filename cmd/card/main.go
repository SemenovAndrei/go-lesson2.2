package main

import (
	"fmt"
	"github.com/i-hit/go-lesson2.2.git/pkg/card"
	"github.com/i-hit/go-lesson2.2.git/pkg/transfer"
)

func main() {
	cardSvc := card.NewService("Tinkoff")

	card1 := cardSvc.GetNewCard("visa", 1000, "RUB", "0001")
	card2 := cardSvc.GetNewCard("visa", 100, "RUB", "0002")

	list := transfer.NewService(cardSvc)

	fmt.Println(list.CardSvc)
	fmt.Println(card1, card2)

	fmt.Println(cardSvc.CheckNumber("0001"))
	fmt.Println(list.Card2Card("0001", "10002", 10000))

	fmt.Println(card1, card2)
}
