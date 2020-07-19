package transfer

import (
	"errors"
	"github.com/i-hit/go-lesson2.2.git/pkg/card"
)

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

var ErrFromCardBalance = errors.New("transfer amount is greater than balance")


func (s *Service) Card2Card(from string, to string, amount int64) error {
	commission := s.CommissionToTinkoff
	total := amount + commission


	cardFrom, err := s.CardSvc.CheckNumber(from, "cardFrom")
	if err != nil {
		return err
	}

	cardTo, err := s.CardSvc.CheckNumber(to, "cardTo")
	if err != nil {
		return err
	}

	if cardFrom.Balance < total {
		return ErrFromCardBalance
	}

	cardFrom.Balance -= total
	cardTo.Balance += total

	return nil
}