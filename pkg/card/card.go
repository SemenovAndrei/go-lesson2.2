package card

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Issuer string
	Balance int64
	Currency string
	Number string
}

type Service struct {
	BankName string
	Cards []*Card
}

func NewService(bankname string) *Service  {
	return &Service {
		BankName : bankname,
	}
}

func (s *Service) GetNewCard(issuer string, balance int64, currency string, number string) *Card {
	card := &Card{
		Issuer: issuer,
		Balance: balance,
		Currency: currency,
		Number: number,
	}

	s.Cards = append(s.Cards, card)

	return card
}

// errors messages
var ErrCardNoValid = errors.New("wrong number")
var ErrCard = errors.New("card not found")

func (s *Service) CheckNumber(number, name string) (*Card, error) {

	str := strings.ReplaceAll(number," ", "")

	ok := isValid(str)
	if !ok {
		return nil, fmt.Errorf("%w. Card name: %s", ErrCardNoValid, name)
	}

	if strings.HasPrefix(str, "510621") {
		for _, c := range s.Cards {
			if strings.ReplaceAll(c.Number, " ", "") == str {
				return c, nil
			}
		}
	}

	return nil, fmt.Errorf("%w. Card name: %s", ErrCard, name)
}

func isValid(number string) bool {
	strToSlice := strings.Split(number, "")
	strToNumber := make( []int, len(strToSlice))

	for i, value := range strToSlice {
		r, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		strToNumber[i] = r
	}

	index := 0
	if len(strToNumber) % 2 != 0 {
		index = 1
	}

	for i := index; i < len(strToNumber); i += 2 {
		strToNumber[i] *= 2
		if strToNumber[i] > 9 {
			strToNumber[i] -= 9
		}
	}

	total := 0
	for _, value := range strToNumber {
		total += value
	}

	if total % 10 != 0 {
		return false
	}

	return true
}

