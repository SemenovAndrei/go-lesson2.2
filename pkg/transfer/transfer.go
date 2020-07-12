package transfer

import "github.com/i-hit/go-lesson2.2.git/pkg/card"



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

func (s *Service) Card2Card(from string, to string, amount int64) (int64, bool) {
	// TODO: ваш код
	commision := s.CommissionToTinkoff
	total := amount + commision

	cardTo, ok := s.CardSvc.CheckNumber(to)
	if !ok {
		cardFrom, ok := s.CardSvc.CheckNumber(from)
		if !ok {
			commision = s.CommissionOther * amount / 1000
			if commision < s.MinimumOther {
				commision = s.MinimumOther
			}
			total = amount + commision
			return total, true
		}

		commision = s.CommissionFromTinkoff * amount / 1000
		if commision < s.MinimumFromTinkoff {
			commision = s.MinimumFromTinkoff
		}
		total = amount + commision
		if cardFrom.Balance >= total {
			cardFrom.Balance -= total
			return total, true
		}
		return total, false
	}
	cardFrom, ok := s.CardSvc.CheckNumber(from)
	if ok {
		if cardFrom.Balance >= total {
			cardFrom.Balance -= total
			cardTo.Balance += total
			return total, true
		}
		return total, false
	}
	cardTo.Balance += total

	return total, true
}