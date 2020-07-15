package transfer

import (
	"errors"
	"github.com/i-hit/go-lesson2.2.git/pkg/card"
)

var ErrFromCardBalance = errors.New("transfer amount is greater than balance")
var ErrCardFrom = errors.New("card from not found")
var ErrCardTo = errors.New("card to not found")

type Service struct {
	CardSvc *card.Service
	// поля для хранения:
	// комиссий в процентах и минимума в рублях*
	CommissionToTinkoff int64
	CommissionFromTinkoff int64
	MinimumFromTinkoff int64
	CommissionOther int64
	MinimumOther int64
}

func NewService(cardSvc *card.Service) *Service {
	return &Service{
		CardSvc: cardSvc,
		CommissionToTinkoff: 0,
		CommissionFromTinkoff: 5,
		MinimumFromTinkoff: 10,
		CommissionOther: 15,
		MinimumOther: 30,
	}
}

func (s *Service) Card2Card(from string, to string, amount int64) error {
	// TODO: ваш код
	commision := s.CommissionToTinkoff
	total := amount + commision


	cardFrom, ok := s.CardSvc.CheckNumber(from)
	if !ok {
		return ErrCardFrom
	}
	cardTo, ok := s.CardSvc.CheckNumber(to)
	if !ok {
		return ErrCardTo
	}

	if cardFrom.Balance < total {
		return ErrFromCardBalance
	}

	cardFrom.Balance -= total
	cardTo.Balance += total

	return nil
}